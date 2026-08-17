package password

import "testing"

func TestBug2ChecksAllBlacklistEntries(t *testing.T) {
	if !Blacklisted("secret", []string{"admin", "secret"}) {
		t.Fatal("a match in a later blacklist entry must be detected")
	}
}
