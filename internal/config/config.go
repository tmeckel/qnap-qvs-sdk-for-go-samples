package config

import b64 "encoding/base64"

var (
	clientID     string
	clientSecret string
	tenantID     string
)

// ClientID is the QNAP client ID aka username.
func ClientID() string {
	if clientID == "" {
		panic("ClientID not configured")
	}
	return clientID
}

// ClientSecret is the QNAP client secret aka password
func ClientSecret() string {
	if clientSecret == "" {
		panic("ClientSecret not configured")
	}
	return b64.StdEncoding.EncodeToString([]byte(clientSecret))
}

// TenantID is the QNAP device url to which this client belongs.
func TenantID() string {
	if tenantID == "" {
		panic("TenantID not configured")
	}
	return tenantID
}
