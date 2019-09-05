/*
Package psqlxtest provides utilities for creating isolated postgres
database instances with sqlx connections for testing purposes.
*/
package psqlxtest

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// keepDB around for debugging purposes?
var keepDB, _ = strconv.ParseBool(os.Getenv("TEST_KEEP_DB"))

// TmpDB creates a temporary database and returns a function to remove it.
func TmpDB(t *testing.T) (*sqlx.DB, func()) {
	u := dbURL(t)
	name := randDBName(t)

	// Create original db connection to create database.
	db0, err := sqlx.Connect("postgres", u.String())
	if err != nil {
		t.Fatal("unable to connect to localhost db:", err)
	}
	if _, err := db0.Exec("CREATE DATABASE " + name + ";"); err != nil {
		t.Fatal("unable to create database:", err)
	}

	return db0, func() {
		if !keepDB {
			if _, err := db0.Exec("DROP DATABASE " + name + ";"); err != nil {
				t.Fatalf("unable to drop database: %v", err)
			}
		}
		if err := db0.Close(); err != nil {
			t.Fatalf("unable to close database: %v", err)
		}
	}
}

func dbURL(t *testing.T) url.URL {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword("postgres", "postgres"),
		Host:   "localhost:5432",
		Path:   "postgres",
		RawQuery: (url.Values{
			"sslmode":  []string{"disable"},
			"timezone": []string{"UTC"},
		}).Encode(),
	}

	if h, ok := os.LookupEnv("TEST_DB_HOST"); ok {
		t.Logf("using TEST_DB_HOST=%q", h)
		u.Host = h
	}

	return u
}

func randDBName(t *testing.T) string {
	return fmt.Sprintf("%s_%v", strings.ToLower(t.Name()), time.Now().Unix())
}
