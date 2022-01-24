package googleAuthenticationHandler

import (
	"context"
	"encoding/json"
	"net/http"
	"os"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"

	"github.com/coding-CEO/go-backend-test/utils"
	"github.com/coding-CEO/go-backend-test/utils/httpUtils"
)

//FIXME: I am not very happy with structure of ignoring error,
// but it will crash in the beginning if there is any error, so it works fine for now.
var (
	_ = godotenv.Load() //ignored err
	clientID     = os.Getenv("GOOGLE_OAUTH2_CLIENT_ID")
	clientSecret = os.Getenv("GOOGLE_OAUTH2_CLIENT_SECRET")

	provider, _ = oidc.NewProvider(context.TODO(), "https://accounts.google.com") //ignored err
	oidcConfig = &oidc.Config{
		ClientID: clientID,
	}
	verifier = provider.Verifier(oidcConfig)
	config = oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email"},
	}
)

// TODO: implement the challenge_code and challenge_code_methon system

func GoogleGenerateUserOAuthCode(w http.ResponseWriter, r *http.Request) {
	
	httpUtils.AddAuthenticationRouteHeaders(w, r)

	redirectUri := r.URL.Query().Get("redirect_uri")
	if len(redirectUri) <= 0 {
		http.Error(w, "redirect_uri in parameters is empty", http.StatusBadRequest)
		return
	}
	config.RedirectURL = redirectUri;

	state, err := utils.RandomString(16)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	nonce, err := utils.RandomString(16)
	if err != nil {
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	
	utils.SetCallbackCookie(w, r, "state", state)
	utils.SetCallbackCookie(w, r, "nonce", nonce)
	
	w.Write([]byte(config.AuthCodeURL(state, oidc.Nonce(nonce))))
}

func GoogleVerifyUserOAuthCode(w http.ResponseWriter, r *http.Request) {

	httpUtils.AddAuthenticationRouteHeaders(w, r)

	state, err := r.Cookie("state")
	if err != nil {
		http.Error(w, "state not found in cookie", http.StatusBadRequest)
		return
	}
	if r.URL.Query().Get("state") != state.Value {
		http.Error(w, "state did not match", http.StatusBadRequest)
		return
	}

	oauth2Token, err := config.Exchange(context.TODO(), r.URL.Query().Get("code"))
	if err != nil {
		http.Error(w, "Failed to exchange token: " + err.Error(), http.StatusInternalServerError)
		return
	}
	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		http.Error(w, "No id_token field in oauth2 token.", http.StatusInternalServerError)
		return
	}
	idToken, err := verifier.Verify(context.TODO(), rawIDToken)
	if err != nil {
		http.Error(w, "Failed to verify ID Token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	nonce, err := r.Cookie("nonce")
	if err != nil {
		http.Error(w, "nonce not found in cookie", http.StatusBadRequest)
		return
	}
	if idToken.Nonce != nonce.Value {
		http.Error(w, "nonce did not match", http.StatusBadRequest)
		return
	}

	oauth2Token.AccessToken = "*REDACTED*" // hide the access token, because you don't need it anymore now

	resp := struct {
		OAuth2Token   *oauth2.Token
		IDTokenClaims *json.RawMessage // ID Token payload is just JSON.
	}{oauth2Token, new(json.RawMessage)}

	if err := idToken.Claims(&resp.IDTokenClaims); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)
}