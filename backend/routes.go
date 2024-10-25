package main

import (
	"encoding/json"
	"net/http"
	"go-svelte/storage"
	"fmt"
	"log"
	// "errors"
	"github.com/gorilla/sessions"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)

	err := storage.CreateUser(creds.Username, creds.Password)
	if err != nil {
		log.Println("Error creating user:", err)
		http.Error(w, "Error: Could not creaete user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created"))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	json.NewDecoder(r.Body).Decode(&creds)

	err := storage.AuthenticateUser(creds.Username, creds.Password)
	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	session, _ := store.Get(r, "session")
	session.Values["authenticated"] = true
	session.Values["username"] = creds.Username
	session.Options = &sessions.Options{
		Path:		"/",
		MaxAge: 	int(sessionExpiration.Seconds()),
		HttpOnly:	true,
		SameSite:	http.SameSiteStrictMode,
	}


	err = session.Save(r, w)
	if err != nil {
		log.Println("Error saving session:", err)
		http.Error(w, "Unable to save session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Login successful"))
	// fmt.Fprintf(w, `{"message": "Login successful"}`)
}

func checkAuthHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Request: %+v\n", r)
	session, _ := store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		log.Println("Auth failed")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	log.Println("Auth success")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"authenticated": true}`)
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "Hello from the Go Dashboard!"}`))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options.MaxAge = -1
	if err := session.Save(r, w); err != nil {
		http.Error(w, "Failed to logout", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"message": "Logged out successfully"}`)
}