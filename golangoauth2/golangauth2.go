package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

const (
	clientID     = "fdab27asfs38yourown9b"
	clientSecret = "b4740adskfjdd254ddad4833eyourownf13d38c7"
	redirectURL  = "http://localhost:3000/auth/callback"
	state        = "random-string"
)

var (
	oauthConf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"user:email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
	// Some random string, random for each request
	oauthStateString = "random"
)

func main() {

	fmt.Println("# 42Snippets")
	fmt.Println("## Golang OAuth2")

	/*
		OAuth 2.0 (Open Authorization) is an open standard for authorization that enables third-party applications to obtain limited access to a user's resources, such as a user's data in a web service. It works by allowing a user to grant a third-party application access to their resources without sharing their credentials.

		OAuth 2.0 is commonly used as a way for users to log in to third-party applications using their existing accounts (such as their Github accounts). It is also often used to authorize API requests, allowing users to grant limited access to their data to other applications or developers.

		To use OAuth 2.0, a user must first authenticate with the service that hosts their resources (the "resource server"). The user is then redirected to the service's authorization server, where they are prompted to grant access to their resources to a third-party application. If the user grants access, the authorization server sends an authorization code back to the application. The application can then use this code to obtain an access token, which it can use to access the user's resources on the resource server.

		This code defines a few HTTP handlers to demonstrate this:

			- handleHome displays a link that the user can click to log in with GitHub.
			- handleLogin redirects the user to the GitHub OAuth 2.0 authorization URL.
			- handleGitHubCallback exchanges the authorization code for an access token and fetches the authenticated user's profile information.

		The user is redirected to the /login URL, which redirects them to the GitHub OAuth 2.0 authorization URL. After the user grants access, they are redirected back to the /auth/callback URL with an authorization code. This code is exchanged for an access token and the user's profile information is fetched using the access token.

		ClientID and secret - https://github.com/settings/applications/new
	*/

	http.HandleFunc("/", handleHome)
	http.HandleFunc("/login", handleLogin)
	http.HandleFunc("/auth/callback", handleGitHubCallback)
	http.ListenAndServe(":3000", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	var htmlIndex = `<html><body>
            <a href="/login">Log in with GitHub</a>
        </body></html>`
	fmt.Fprintf(w, htmlIndex)
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	url := oauthConf.AuthCodeURL(oauthStateString, oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleGitHubCallback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	oauthClient := oauthConf.Client(context.Background(), token)
	resp, err := oauthClient.Get("https://api.github.com/user")
	if err != nil {
		fmt.Printf("oauthClient.Get() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	var user User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		fmt.Printf("json.NewDecoder().Decode() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	fmt.Printf("Logged in as GitHub user: %s\n", user.Login)

	// Display a message on the page to indicate that the user has successfully logged in
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Welcome, %s! You have successfully logged in.", user.Login)
}

type User struct {
	Login string `json:"login"`
}
