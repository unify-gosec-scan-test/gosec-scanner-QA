package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os/exec"

	_ "github.com/mattn/go-sqlite3"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the vulnerable Go app!")
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	target := r.URL.Query().Get("target")
	cmd := exec.Command("sh", "-c", "ping -c 1 "+target) // Vulnerable to command injection
	out, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Ping failed: "+err.Error(), 500) // Verbose error
		return
	}
	fmt.Fprintf(w, "Output:\n%s", out)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	password := r.URL.Query().Get("password")

	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		http.Error(w, "DB error", 500)
		return
	}
	defer db.Close()

	// SQL Injection
	query := "SELECT * FROM users WHERE username = '" + username + "' AND password = '" + password + "'"
	row := db.QueryRow(query)

	var id int
	var user, pass string
	err = row.Scan(&id, &user, &pass)
	if err != nil {
		http.Error(w, "Login failed: "+err.Error(), 403)
		return
	}
	fmt.Fprintf(w, "Welcome %s!", user)
}

func DebugHandler(w http.ResponseWriter, r *http.Request) {
	// Exposed debug info
	fmt.Fprintf(w, "DEBUG MODE ENABLED: UserAgent: %s", r.UserAgent())
}
