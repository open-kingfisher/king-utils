package kit

import (
	"testing"
)

func TestLdapLookup(t *testing.T) {
	client := LdapLookup()
	ok, user, err := client.Authenticate("xiaoming", "xxxxx")
	if err != nil {
		t.Errorf("Error authenticating user %s: %+v", "username", err)
	}
	if !ok {
		t.Errorf("Authenticating failed for user %s", "username")
	}
	t.Logf("User: %+v", user)

	users, err := client.GetUser("xiaoming")
	if err != nil {
		t.Errorf("Get user info err: %s", err)
	} else {
		t.Logf("User: %+v", users)
	}
}
