package util_test

import (
	"testing"

	"github.com/nasa9084/salamander/salamander/util"
)

func TestPassword(t *testing.T) {
	candidates := []struct {
		inPassword string
		inSalt     string
		expected   string
	}{
		{"password", "", "$$sha512$30$b852faf38923d6e7d3977237e52ab36b29b8a08f02bfd0b1fe5696a4f5965c0115362536234cb531240394ef9fe52ef1d94f64eb8586c3472bd22dba99e79440"},
		{"somethinganythingnothing", "", "$$sha512$30$9055c42bd36ee0a0d8b1127b863d306bb7f30c841abfa93e96445055466ad7a3ea2c6ec852a3106f174ef4d8552132ea24cea87e16d3846027e91171dd07fbd1"},
		{"password", "userID", "$userID$sha512$30$5732df155b77754b631b6d8324115dd2d3c5af4f86bcbde21690e06f7a9920fa73f33dba139530cfa82aefdbe2d8c693a0564abdfea73d27ceea4e1c53cf3780"},
	}

	for _, c := range candidates {
		hashed := util.Password(c.inPassword, c.inSalt)
		if hashed != c.expected {
			t.Errorf(`"%s" != "%s"`, hashed, c.expected)
			return
		}
	}
}
