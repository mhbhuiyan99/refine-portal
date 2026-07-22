package requests

import (
	"fmt"
	"strings"

	"github.com/beego/beego/v2/server/web"
)

// GetURLFromConfig returns a URL value from the application configuration.
//
// Responsibilities:
//   - Read the URL from the application configuration.
//   - Validate that the URL is not empty.
//   - Return the configured URL.
func GetURLFromConfig(
	configKey string,
) (string, error) {

	url, err := web.AppConfig.String(configKey)
	if err != nil {
		return "", fmt.Errorf("failed to read %s: %w", configKey, err)
	}

	if strings.TrimSpace(url) == "" {
		return "", fmt.Errorf("%s is empty", configKey)
	}

	return url, nil
}