package kit

import "fmt"

func LdapLookup(url, bindDN, bindPassword, baseDN, userFilter string, tls bool) *LDAPClient {
	client := &LDAPClient{
		URL:          url,
		Base:         baseDN,
		UseTLS:       tls, // TLS认证
		BindDN:       bindDN,
		BindPassword: bindPassword,
		UserFilter:   fmt.Sprintf("(%s=%%s)", userFilter),
		Attributes:   []string{},
	}
	return client
}
