package database

// cribbed from lstoll

import (
	"database/sql"
	"net/url"
	"os"
	"runtime"

	"github.com/pkg/errors"
)

// New returns a sql.DB set up for this application.
func New(defaultURI string) (*sql.DB, error) {
	client := os.Getenv("DYNO")
	uri := os.Getenv("DATABASE_URL")
	if uri == "" {
		uri = defaultURI
	}
	db, err := sql.Open("postgres", formatURL(uri, client))
	if err != nil {
		return nil, errors.Wrap(err, "Error connecting to database")
	}
	db.SetMaxOpenConns(runtime.NumCPU())
	return db, nil
}

// formatURL returns new database URL with application_name set.
func formatURL(uri, client string) string {
	url, err := url.Parse(uri)
	if err != nil {
		return uri
	}
	v := url.Query()
	v.Set("application_name", client)
	url.RawQuery = v.Encode()
	return url.String()
}
