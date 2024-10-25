package main

import (
	// "fmt"
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

const sessionExpiration = time.Minute

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
	// http.Handle("/_app/", http.StripPrefix("/_app/", http.FileServer(http.Dir("../frontend/build/_app"))))

	log.Println("Server started on :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}