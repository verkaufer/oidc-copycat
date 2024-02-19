package oauth2

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Defines the OAuth /authorize endpoint

func authorizeEndpoint(c *gin.Context) {
	ctx := c.Request.Context()

	ar, err := oauth2.NewAuthorizeRequest(ctx, c.Request)
	if err != nil {
		log.Printf("failed to NewAuthorizeRequest: %+v", err)
		oauth2.WriteAuthorizeError(ctx, c.Writer, ar, err)
		return
	}

	// TODO: read and present
	for _, s := range ar.GetRequestedScopes() {
		log.Printf("Asking for scope: %s\n", s)
	}

	// BEGIN: user has not granted access, return form instead
	if c.PostForm("userEmail") != "example@example.com" {
		log.Println("Did not submit userEmail with expected value. Returning form.")
		c.HTML(http.StatusOK, "auth/index.tmpl", gin.H{"title": "Authorize"})
		return
	}

	// END: user not granted access

	// let's see what scopes the user gave consent to
	for _, scope := range c.QueryArray("scopes") {
		ar.GrantScope(scope)
	}

	// Now that the user is authorized, we set up a session:
	mySessionData := newSession("peter")

	resp, err := oauth2.NewAuthorizeResponse(ctx, ar, mySessionData)
	if err != nil {
		log.Printf("failed to NewAuthorizeResponse: %+v", err)
		oauth2.WriteAuthorizeError(ctx, c.Writer, ar, err)
		return
	}

	oauth2.WriteAuthorizeResponse(ctx, c.Writer, ar, resp)
}
