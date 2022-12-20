package secondfile

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err != nil {
		t.Error("hello is not an email")
	}

	_, err = IsEmail("simar@gmail.com")
	if err == nil {
		t.Errorf("simar@gmail.com is a valid email")
	}
}
