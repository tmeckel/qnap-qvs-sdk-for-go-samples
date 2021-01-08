package config

import (
	"os"
)

// ParseEnvironment loads configured environment variables
func ParseEnvironment() error {
	clientID = os.Getenv("QNAPQVS_CLIENT_ID")
	clientSecret = os.Getenv("QNAPQVS_CLIENT_SECRET")
	tenantID = os.Getenv("QNAPQVS_TENANT_ID")

	return nil
}
