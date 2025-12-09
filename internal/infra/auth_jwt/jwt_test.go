package auth_jwt

import (
	"testing"
	"time"

	application "todoDB/internal/application/user"
	"todoDB/internal/domain/auth"
)

func TestGenerateToken(t *testing.T) {
	secret := application.AccessTokenSecret
	claims := &auth.Claims{
		UserName:  "123",
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(time.Hour),
	}

	token, err := GenerateToken(secret, claims.UserName, claims)
	if err != nil {
		t.Fatalf("GenerateToken returned error: %v", err)
	}

	if token == "" {
		t.Fatalf("GenerateToken returned empty token")
	}
}

func TestParseToken(t *testing.T) {
	secret := application.AccessTokenSecret
	claims := &auth.Claims{
		UserName:  "abc",
		IssuedAt:  time.Now(),
		ExpiresAt: time.Now().Add(2 * time.Hour),
	}

	// generate
	token, err := GenerateToken(secret, claims.UserName, claims)
	if err != nil {
		t.Fatalf("GenerateToken error: %v", err)
	}

	// parse
	parsed, err := ParseToken(secret, token)
	if err != nil {
		t.Fatalf("ParseToken error: %v", err)
	}

	// fmt.Println(parsed.UserName)
	// fmt.Printf("parsed: %v\n", parsed)

	// compare
	if parsed.UserName != claims.UserName {
		t.Errorf("user_id mismatch: got %s, want %s", parsed.UserName, claims.UserName)
	}

	if !parsed.IssuedAt.Equal(claims.IssuedAt.Truncate(time.Second)) {
		t.Errorf("issued_at mismatch")
	}

	if !parsed.ExpiresAt.Equal(claims.ExpiresAt.Truncate(time.Second)) {
		t.Errorf("expires_at mismatch")
	}
}

func TestParseTokenInvalid(t *testing.T) {
	secret := application.AccessTokenSecret
	// totally invalid token string
	_, err := ParseToken(secret, "not-a-token")
	if err == nil {
		t.Fatalf("expected error for invalid token")
	}
}
