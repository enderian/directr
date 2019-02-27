package delegation

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/enderian.directrd/types"
	"gopkg.in/ldap.v3"
	"strings"
)

func initLdapConnection() (*ldap.Conn, error) {
	var l *ldap.Conn
	var err error

	if conf == nil || conf.LDAP.Address == "" {
		return nil, errors.New("ldap has not been enabled")
	}

	if conf.LDAP.UseTLS {
		l, err = ldap.DialTLS("tcp4",
			fmt.Sprintf("%s:%d", conf.LDAP.Address, conf.LDAP.Port),
			&tls.Config{ServerName: conf.LDAP.Address, InsecureSkipVerify: true})
	} else {
		l, err = ldap.Dial("tcp4", fmt.Sprintf("%s:%d", conf.LDAP.Address, conf.LDAP.Port))
	}

	if err != nil {
		return nil, err
	}
	return l, l.Bind(conf.LDAP.BindUsername, conf.LDAP.BindPassword)
}

func AuthenticateLdap(user *types.User, password string) error {
	l, err := initLdapConnection()
	if err != nil {
		return err
	}
	defer l.Close()

	// Search for the given username
	filter := fmt.Sprintf("(&(objectClass=%s)(%s=%s))",
		conf.LDAP.SearchClass, conf.LDAP.UsernameAttribute, user.Username)
	searchRequest := ldap.NewSearchRequest(
		conf.LDAP.BaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,
		0, 0, false, filter, []string{"dn"}, nil)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return err
	}

	if len(sr.Entries) != 1 {
		return errors.New("user does not exist or too many entries returned")
	}

	targetDN := sr.Entries[0].DN
	// Bind as the user to verify their password
	err = l.Bind(targetDN, password)
	if err != nil {
		return err
	}

	//Fill the rest of the data and do the rest of the work.
	return FillLdap(user)
}

func FillLdap(user *types.User) error {
	l, err := initLdapConnection()
	if err != nil {
		return err
	}
	defer l.Close()

	// Search for the given username
	filter := fmt.Sprintf("(&(objectClass=%s)(%s=%s))",
		conf.LDAP.SearchClass, conf.LDAP.UsernameAttribute, user.Username)
	searchRequest := ldap.NewSearchRequest(
		conf.LDAP.BaseDN, ldap.ScopeWholeSubtree, ldap.NeverDerefAliases,
		0, 0, false, filter, nil, nil)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return err
	}
	if len(sr.Entries) != 1 {
		return errors.New("user does not exist or too many entries returned")
	}
	if conf.LDAP.ExtraAttributes != nil {
		user.Extras = make(map[string]string)
	}

	user.DN = sr.Entries[0].DN

	//Setup all the other attributes
	for _, entry := range sr.Entries[0].Attributes {
		switch entry.Name {
		case conf.LDAP.FullNameAttribute:
			{
				if conf.LDAP.FixFullNameCase {
					user.FullName = strings.Title(strings.ToLower(entry.Values[0]))
				} else {
					user.FullName = entry.Values[0]
				}
			}
		case conf.LDAP.AffiliationAttribute:
			user.Affiliation = entry.Values[0]
		default:
			{
				if conf.LDAP.ExtraAttributes != nil {
					if extName, ok := conf.LDAP.ExtraAttributes[entry.Name]; ok {
						user.Extras[extName] = entry.Values[0]
					}
				}
			}
		}
	}
	return nil
}
