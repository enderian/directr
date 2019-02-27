package delegation

import (
	"github.com/enderian/directrd/types"
	"testing"
)

func Test_authenticateUsingLdap(t *testing.T) {
	conf = &types.Configuration{
		LDAP: &types.LdapConfiguration{
			Address:              "ds.aueb.gr",
			Port:                 636,
			UseTLS:               true,
			BaseDN:               "dc=aueb,dc=gr",
			BindUsername:         "uid=roService,ou=Services,dc=aueb,dc=gr",
			BindPassword:         "r0s-@gu31",
			SearchClass:          "eduPerson",
			UsernameAttribute:    "uid",
			FullNameAttribute:    "cn",
			AffiliationAttribute: "eduPersonPrimaryAffiliation",
		},
	}
	if err := AuthenticateLdap(&types.User{Username: "p3150133"}, "xxx"); err != nil {
		t.Fatal(err)
	}
}
