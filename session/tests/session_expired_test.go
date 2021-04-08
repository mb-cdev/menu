package tests

import (
	"menu/session/session"
	"testing"
	"time"
)

func TestIsSessionExpired(t *testing.T) {
	sExpired := session.New()
	sNotExpired := session.New()

	sExpired.Expire_at = time.Now().Unix() - 10

	if !sExpired.IsExpired() {
		t.Error("Session expired is tagged as not expired")
	}

	if sNotExpired.IsExpired() {
		t.Error("Session not expired is tagged as expired")
	}

	sExpired.Drop()
	sNotExpired.Drop()
}
