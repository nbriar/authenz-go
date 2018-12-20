package configuration

import "os"

// AuthTokenSecret is set from the environment variable AUTH_TOKEN_SECRET
// it defaults to 'secret' if the env variable is not set
func AuthTokenSecret() string {
	AuthTokenSecret := os.Getenv("AUTH_TOKEN_SECRET")
	if AuthTokenSecret == "" {
		AuthTokenSecret = "secret"
	}

	return AuthTokenSecret
}

// HTTP header constants. Please keep these lines alphabetized.
// IMPORTANT: All values here must be in "canonical" form, which means that all letters are
// lowercase except the first letter and any letter immediately following a dash.  (Just like
// https://golang.org/pkg/net/http/#CanonicalHeaderKey.)
const (
	HTTP_HEADER_ACCESS_CONTROL_ALLOW_CREDENTIALS = "Access-Control-Allow-Credentials"
	HTTP_HEADER_ACCESS_CONTROL_ALLOW_HEADERS     = "Access-Control-Allow-Headers"
	HTTP_HEADER_ACCESS_CONTROL_ALLOW_METHODS     = "Access-Control-Allow-Methods"
	HTTP_HEADER_ACCESS_CONTROL_ALLOW_ORIGIN      = "Access-Control-Allow-Origin"
	HTTP_HEADER_ACCESS_CONTROL_REQUEST_HEADERS   = "Access-Control-Request-Headers"
	HTTP_HEADER_ACCESS_CONTROL_REQUEST_METHOD    = "Access-Control-Request-Method"
	HTTP_HEADER_AUTHORIZATION                    = "Authorization"
	HTTP_HEADER_CONTENT_TYPE                     = "Content-Type"
	HTTP_HEADER_ORIGIN                           = "Origin"
	HTTP_HEADER_VARY                             = "Vary"
	HTTP_HEADER_X_REAL_IP                        = "X-Real-Ip"
)
