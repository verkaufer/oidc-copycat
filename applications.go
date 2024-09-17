package oidc_copycat

type OIDCApplication struct {
	Name         string   `form:"name" binding:"required" validate:"required,alphanumunicode"`
	AppType      string   `form:"visibility" binding:"required" validate:"oneof=public private"`
	RedirectURIs []string `form:"redirect_uri" binding:"gt=0,dive,uri"`
}
