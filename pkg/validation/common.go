package validation

import (
	"fmt"
	"os"

	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/apis/options"
)

const multipleValuesForSecretSource = "multiple values specified for secret source: specify either value, fromEnv of fromFile"

func validateSecretSource(source options.SecretSource) string {
	switch {
	case len(source.Value) > 0 && source.FromEnv == "" && source.FromFile == "":
		return ""
	case len(source.Value) == 0 && source.FromEnv != "" && source.FromFile == "":
		return validateSecretSourceEnv(source.FromEnv)
	case len(source.Value) == 0 && source.FromEnv == "" && source.FromFile != "":
		return validateSecretSourceFile(source.FromFile)
	default:
		return multipleValuesForSecretSource
	}
}

func validateSecretSourceEnv(key string) string {
	if value := os.Getenv(key); value == "" {
		return fmt.Sprintf("error loading secret from environent: no value for for key %q", key)
	}
	return ""
}

func validateSecretSourceFile(path string) string {
	if _, err := os.Stat(path); err != nil {
		return fmt.Sprintf("error loadig secret from file: %v", err)
	}
	return ""
}
