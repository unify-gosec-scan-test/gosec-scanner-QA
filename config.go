package main

import "os"

func LoadSecrets() (string, string) {
	// Hardcoded fallback
	user := os.Getenv("DB_USER")
	if user == "" {
		user = "admin"
	}
	pass := os.Getenv("DB_PASS")
	if pass == "" {
		pass = "password123"
	}
	return user, pass
}
