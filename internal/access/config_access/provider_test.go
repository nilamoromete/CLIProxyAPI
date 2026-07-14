package configaccess

import (
	"testing"

	sdkconfig "github.com/router-for-me/CLIProxyAPI/v6/sdk/config"
)

func TestNewProviderRejectsPlaceholderAPIKey(t *testing.T) {
	cfg := &sdkconfig.AccessProvider{Name: "config-test", APIKeys: []string{"your-api-key-1"}}

	provider, err := newProvider(cfg, nil)
	if err == nil {
		t.Fatal("expected placeholder API key to be rejected")
	}
	if provider != nil {
		t.Fatal("expected no provider when placeholder API key is configured")
	}
}

func TestNewProviderAcceptsNonPlaceholderAPIKey(t *testing.T) {
	cfg := &sdkconfig.AccessProvider{Name: "config-test", APIKeys: []string{"local-unique-test-key"}}

	provider, err := newProvider(cfg, nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if provider == nil {
		t.Fatal("expected provider for non-placeholder key")
	}
}
