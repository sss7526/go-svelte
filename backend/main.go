package main

import (
	"fmt"
	"go-svelte/storage"
	"log"
	"net/http"
	"os"
	"path/filepath"

	// "encoding/json"
	"time"

	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var store *sessions.CookieStore

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	

	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		log.Fatal("SESSION_KEY not set in environment")
	}

	store = sessions.NewCookieStore([]byte(sessionKey))
}


// var mockUsername = "admin"
// var mockPassword = "password"

const sessionExpiration = time.Minute

// type Credentials struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func main() {
	_, err := storage.InitSQLiteDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	fs := http.FileServer(http.Dir("../frontend/build"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join("../frontend/build", r.URL.Path)

		if fileExists(path) {
			fs.ServeHTTP(w, r)
			return
		}

		http.ServeFile(w, r, filepath.Join("../frontend/build", "index.html"))
	})

	http.HandleFunc("/api/login", loginHandler)
	http.HandleFunc("/api/signup", signupHandler)
	http.HandleFunc("/api/check-auth", checkAuthHandler)
	http.HandleFunc("/api/dashboard", dashboardHandler)
	http.HandleFunc("/api/logout", logoutHandler)

	log.Println("Server started on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

// func loginHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
// 		return
// 	}

// 	var creds Credentials
// 	err := json.NewDecoder(r.Body).Decode(&creds)
// 	if err != nil {
// 		http.Error(w, "Bad Request", http.StatusBadRequest)
// 		return
// 	}

// 	if creds.Username == mockUsername && creds.Password == mockPassword {
// 		session, _ := store.Get(r, "session")

// 		session.Values["authenticated"] = true
// 		session.Values["username"] = creds.Username
// 		session.Options = &sessions.Options{
// 			Path:		"/",
// 			MaxAge: 	int(sessionExpiration.Seconds()), 
// 			HttpOnly: 	true,
// 			SameSite:	http.SameSiteStrictMode,
// 		}

// 		err := session.Save(r, w)
// 		if err != nil {
// 			log.Println("Error saving session:", err)
// 			http.Error(w, "Unable to save session", http.StatusInternalServerError)
// 			return
// 		}
		
// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintf(w, `{"message": "Login successful"}`)
// 	} else {
// 		http.Error(w, "Unauthorized", http.StatusUnauthorized)
// 	}
// }

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

func checkAuthHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	auth, ok := session.Values["authenticated"].(bool)
	if !ok || !auth {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}
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