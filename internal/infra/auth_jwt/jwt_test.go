package auth_jwt

import (
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	secret := AccessTokenSecret
	claims := &Claims{
		UserID:    "123",
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(time.Hour),
	}

	token, err := GenerateToken(secret, claims.UserID, claims)
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}

	if token == "" {
		t.Fatalf("GenerateToken returned empty token")
	}
}

func TestParseToken(t *testing.T) {
	secret := AccessTokenSecret
	claims := &Claims{
		UserID:    "abc",
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(2 * time.Hour),
	}

	// generate
	token, err := GenerateToken(secret, claims.UserID, claims)
	if err != nil {
		t.Fatalf("GenerateToken error: %v", err)
	}

	// parse
	parsed, err := ParseToken(secret, token)
	if err != nil {
		t.Fatalf("ParseToken error: %v", err)
	}

	// compare
	if parsed.UserID != claims.UserID {
		t.Errorf("user_id mismatch: got %s, want %s", parsed.UserID, claims.UserID)
	}

	if !parsed.IssuedAt.Equal(claims.IssuedAt.Truncate(time.Second)) {
		t.Errorf("issued_at mismatch")
	}

	if !parsed.ExpiresAt.Equal(claims.ExpiresAt.Truncate(time.Second)) {
		t.Errorf("expires_at mismatch")
	}
}

func TestParseTokenInvalid(t *testing.T) {
	secret := AccessTokenSecret
	// totally invalid token string
	_, err := ParseToken(secret, "not-a-token")
	if err == nil {
		t.Fatalf("expected error for invalid token")
	}
}
