package auth

import "testing"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("hashing password: %s", err)
	}

	if hash == "" {
		t.Errorf("expected: hash, got: empty string")
	}

	if hash == "password" {
		t.Error("expected: hash to be different from password, got: password and hash are the same")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashPassword("password")
	if err != nil {
		t.Errorf("hashing password: %s", err)
	}

	arePasswordsMatched := ComparePasswords(hash, []byte("password"))
	if !arePasswordsMatched {
		t.Errorf("expected: true, got: false")
	}

	arePasswordsMatched = ComparePasswords(hash, []byte("notpassword"))
	if arePasswordsMatched {
		t.Errorf("expected: false, got: true")
	}
}
