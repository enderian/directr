package types

type Configuration struct {
	Database string `yaml:"database"`
	Redis    string `yaml:"redis"`
	Logs     string `yaml:"logs"`

	API    APIConfiguration    `yaml:"api"`
	User   UserConfiguration   `yaml:"user"`
	Radius RadiusConfiguration `yaml:"radius"`
	LDAP   LdapConfiguration   `yaml:"ldap"`
}

type AuthenticationStrategy string

const (
	AuthenticationLDAP   AuthenticationStrategy = "ldap"
	AuthenticationCached AuthenticationStrategy = "cached"
	AuthenticationLocal  AuthenticationStrategy = "local"
	AuthenticationOp     AuthenticationStrategy = "os"
)

type APIConfiguration struct {
	Disabled bool `yaml:"disabled"`

	RestAddr    string `yaml:"rest_addr"`
	ServiceAddr string `yaml:"service_addr"`
}

type UserConfiguration struct {
	Authentication []AuthenticationStrategy `yaml:"authentication"`

	Authorization      bool  `yaml:"authorization"`
	AuthorizationRules Rules `yaml:"authorization_rules"`

	AutoRegister      bool  `yaml:"auto_register"`
	AutoRegisterRules Rules `yaml:"auto_register_rules"`

	RegistrationFilter string `yaml:"registration_filter"`
}

type RadiusConfiguration struct {
	DisabledAuth       bool `yaml:"disabled_auth"`
	DisabledAccounting bool `yaml:"disabled_accounting"`

	SharedSecret string `yaml:"shared_secret"`
}

type LdapConfiguration struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
	UseTLS  bool   `yaml:"use_tls"`

	BaseDN       string `yaml:"base_dn"`
	BindUsername string `yaml:"bind_username"`
	BindPassword string `yaml:"bind_password"`

	SearchClass string `yaml:"search_class"`

	UsernameAttribute    string `yaml:"username_attribute"`
	FullNameAttribute    string `yaml:"full_name_attribute"`
	AffiliationAttribute string `yaml:"affiliation_attribute"`
	MobileAttribute      string `yaml:"mobile_attribute"`

	ExtraAttributes map[string]string `yaml:"extra_attributes"`
	FixFullNameCase bool              `yaml:"fix_full_name_case"`
}
