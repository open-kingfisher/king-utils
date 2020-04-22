package kit

func LdapLookup() *LDAPClient {
	client := &LDAPClient{
		Base:         "dc=example,dc=kingfisher,dc=com",
		Host:         "ldap.kingfisher.com",
		Port:         389,
		UseSSL:       false,
		BindDN:       "cn=admin,ou=ldap,OU=admin,dc=example,dc=kingfisher,dc=com",
		BindPassword: "password",
		UserFilter:   "(userName=%s@kingfisher.com)",
		Attributes:   []string{"mail", "telephoneNumber", "mobile", "name"},
	}
	return client
}
