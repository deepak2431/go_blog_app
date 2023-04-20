package main_test

import (
	"log"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM blogs")
	a.DB.Exec("ALTER SEQUENCE blogs_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS blogs  
(
    id SERIAL,
    title TEXT NOT NULL,  
	post_content TEXT NOT NULL,  
	author TEXT NOT NULL,
    CONSTRAINT blog_pkey PRIMARY KEY (id)  
)`
