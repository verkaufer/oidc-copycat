package oauth2

import (
	"crypto/rand"
	"crypto/rsa"
	"time"

	"github.com/ory/fosite"
	"github.com/ory/fosite/compose"
	"github.com/ory/fosite/handler/openid"
	"github.com/ory/fosite/storage"
	"github.com/ory/fosite/token/jwt"
)

var (
	// TODO: read from env
	secret = []byte("some-cool-secret-that-is-32bytes")

	config = &fosite.Config{
		AccessTokenLifespan:         time.Minute * 30,
		IDTokenIssuer:               "https://copycatoidc.com",
		EnforcePKCEForPublicClients: true,

		// TODO: *EndpointHandlers are called before the respective endpoints are served
		// AuthorizeEndpointHandlers: nil, // called before /authorize served
		// TokenEndpointHandlers:     nil, // called before /token served
	}

	// TODO replace with custom storage
	store         = storage.NewExampleStore()
	privateKey, _ = rsa.GenerateKey(rand.Reader, 2048)
)

// Build a fosite instance with all OAuth2 and OpenID Connect handlers enabled, plugging in our configurations as specified above.
var oauth2 = compose.ComposeAllEnabled(config, store, privateKey)

// A session is passed from the `/auth` to the `/token` endpoint. You probably want to store data like: "Who made the request",
func newSession(user string) *openid.DefaultSession {

	return &openid.DefaultSession{
		Claims: &jwt.IDTokenClaims{
			Subject:     user,
			Audience:    []string{"https://my-client.my-application.com"},
			ExpiresAt:   time.Now().Add(time.Hour * 6),
			IssuedAt:    time.Now(),
			RequestedAt: time.Now(),
			AuthTime:    time.Now(),
		},
		Headers: &jwt.Headers{
			Extra: make(map[string]interface{}),
		},
	}
}
