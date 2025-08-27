package doauth2

import "golang.org/x/oauth2"

const TokenRevocationURL = "https://discord.com/api/oauth2/token/revoke"

var Endpoint = oauth2.Endpoint{
	AuthURL:  "https://discord.com/oauth2/authorize",
	TokenURL: "https://discord.com/api/oauth2/token",
}
