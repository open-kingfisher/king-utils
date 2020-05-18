package kit

func LdapLookup() *LDAPClient {
	client := &LDAPClient{
		Base:         "dc=kingfisher,dc=com",
		Host:         "ldap.kingfisher.com",
		Port:         389,
		UseTLS:       false, // TLS认证
		BindDN:       "cn=admin,dc=kingfisher,dc=com",
		BindPassword: "kingfisher",
		UserFilter:   "(cn=%s)",
		Attributes:   []string{"cn", "mail"},
	}
	return client
}
