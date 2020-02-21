package ipv4

import (
	"testing"
)

func TestIsPublic(t *testing.T) {

	ip := "62.3.44.2"
	if !IsPublic(ip) {
		t.Errorf("Wrong output for IsPublic for %s, got: %v, want: %v.", ip, false, true)
	}

	ip = "192.168.1.2"
	if IsPublic(ip) {
		t.Errorf("Wrong output for IsPublic for %s, got: %v, want: %v.", ip, true, false)
	}

	ip = "100.121.0.3"
	if IsPublic(ip) {
		t.Errorf("Wrong output for IsPublic for %s, got: %v, want: %v.", ip, true, false)
	}
}

func TestIsPrivate(t *testing.T) {
	ip := "192.168.1.2"
	if !IsPrivate(ip) {
		t.Errorf("Wrong output for IsPrivate for %s, got: %v, want: %v.", ip, false, true)
	}

	ip = "200.32.22.12"
	if IsPrivate(ip) {
		t.Errorf("Wrong output for IsPrivate for %s, got: %v, want: %v.", ip, true, false)
	}

	ip = "100.64.0.1"
	if !IsPrivate(ip) {
		t.Errorf("Wrong output for IsPublic for %s, got: %v, want: %v.", ip, false, true)
	}
}

func TestIsCGN(t *testing.T) {
	ip := "100.64.0.1"
	if !IsCGN(ip) {
		t.Errorf("Wrong output for IsCGN for %s, got: %v, want: %v.", ip, false, true)
	}

	ip = "192.168.1.2"
	if IsCGN(ip) {
		t.Errorf("Wrong output for IsPrivate for %s, got: %v, want: %v.", ip, true, false)
	}

	ip = "200.32.22.12"
	if IsCGN(ip) {
		t.Errorf("Wrong output for IsPrivate for %s, got: %v, want: %v.", ip, true, false)
	}
}
