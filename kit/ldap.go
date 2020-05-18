package kit

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/go-ldap/ldap"
	"github.com/open-kingfisher/king-utils/common/log"
)

type LDAPClient struct {
	Attributes         []string
	Base               string
	BindDN             string
	BindPassword       string
	GroupFilter        string
	Host               string
	ServerName         string
	UserFilter         string
	Conn               *ldap.Conn
	Port               int
	InsecureSkipVerify bool
	UseTLS             bool
	SkipTLSVerify      bool
	ClientCertificates []tls.Certificate // Adding client certificates
}

// 连接LDAP服务
func (lc *LDAPClient) Connect() error {
	if lc.Conn == nil {
		var l *ldap.Conn
		var err error
		address := fmt.Sprintf("ldap://%s:%d", lc.Host, lc.Port)

		l, err = ldap.DialURL(address)
		if err != nil {
			return err
		}
		if lc.UseTLS {
			// 带TLS的连接
			err = l.StartTLS(&tls.Config{InsecureSkipVerify: true})
			if err != nil {
				return err
			}
		}

		lc.Conn = l
	}
	return nil
}

// 关闭连接
func (lc *LDAPClient) Close() {
	if lc.Conn != nil {
		lc.Conn.Close()
		lc.Conn = nil
	}
}

// 认证
func (lc *LDAPClient) Authenticate(username, password string) (bool, map[string]string, error) {
	err := lc.Connect()
	if err != nil {
		return false, nil, err
	}
	// 绑定查询用户
	if lc.BindDN != "" && lc.BindPassword != "" {
		err := lc.Conn.Bind(lc.BindDN, lc.BindPassword)
		if err != nil {
			log.Errorf("ldap bind dn:%s password:%s error: %s", lc.BindDN, lc.BindPassword, err)
			return false, nil, err
		}
	}
	// 搜索指定用户
	searchRequest := ldap.NewSearchRequest(
		lc.Base,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(lc.UserFilter, username),
		lc.Attributes,
		nil,
	)
	sr, err := lc.Conn.Search(searchRequest)
	if err != nil {
		log.Errorf("ldap search username:%s, userFilter:%s, error: %s", lc.UserFilter, username, err)
		return false, nil, err
	}
	if len(sr.Entries) < 1 {
		return false, nil, errors.New("user does not exist")
	}

	if len(sr.Entries) > 1 {
		return false, nil, errors.New("too many entries returned")
	}

	userDN := sr.Entries[0].DN
	user := map[string]string{}
	for _, attr := range lc.Attributes {
		user[attr] = sr.Entries[0].GetAttributeValue(attr)
	}
	// 验证用户命名密码
	err = lc.Conn.Bind(userDN, password)
	if err != nil {
		log.Errorf("ldap bind userDN:%s, error: %s", userDN)
		return false, user, err
	}

	// 重新绑定查询用户
	if lc.BindDN != "" && lc.BindPassword != "" {
		err = lc.Conn.Bind(lc.BindDN, lc.BindPassword)
		if err != nil {
			return true, user, err
		}
	}

	return true, user, nil
}

// GetGroupsOfUser returns the group for a user.
func (lc *LDAPClient) GetUser(username string) (map[string]string, error) {
	err := lc.Connect()
	if err != nil {
		return nil, err
	}

	// First bind with a read only user
	if lc.BindDN != "" && lc.BindPassword != "" {
		err := lc.Conn.Bind(lc.BindDN, lc.BindPassword)
		if err != nil {
			return nil, err
		}
	}

	attributes := append(lc.Attributes, "dn")
	// Search for the given username
	searchRequest := ldap.NewSearchRequest(
		lc.Base,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(lc.UserFilter, username),
		attributes,
		nil,
	)

	sr, err := lc.Conn.Search(searchRequest)
	if err != nil {
		return nil, err
	}

	if len(sr.Entries) < 1 {
		return nil, errors.New("User does not exist")
	}

	if len(sr.Entries) > 1 {
		return nil, errors.New("Too many entries returned")
	}

	user := map[string]string{}
	for _, attr := range lc.Attributes {
		user[attr] = sr.Entries[0].GetAttributeValue(attr)
	}

	return user, nil
}
