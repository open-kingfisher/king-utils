package kit

func LdapLookup() *LDAPClient {
	client := &LDAPClient{
		Base:         "dc=kingfisher,dc=com",
		Host:         "ldap.kingfisher.com",
		Port:         389,
		UseSSL:       false,
		BindDN:       "cn=admin,ou=king,dc=kingfisher,dc=com",
		BindPassword: "kingfisher",
		UserFilter:   "(userName=%s@kingfisher.com)",
		Attributes:   []string{"mail", "telephoneNumber", "mobile", "name"},
	}
	return client
}
