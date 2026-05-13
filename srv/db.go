package main

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	// _ "github.com/lib/pq" // To register the driver.
)

var db *sql.DB

func openDB(url string, port string) {
	var err error
	db, err = sql.Open("sqlite3", "stuff.db")

	if err != nil {
		log.Fatal("DB NOT FOUND!")
	}
}

func readUser(id string) *User {
	row := db.QueryRow("SELECT userId, emoji FROM users WHERE userId = ?;", id)
	if row.Err() != nil {
		log.Fatal("USER NOT FOUND, BITCH!")
	}
	user := &User{}
	row.Scan(&user.Id, &user.Emoji)

	if user.Id == "" {
		log.Fatal("Unexpected error, user: ", user)
	}

	return user
}

func readAvailableUserEmojis() []uint16 {
	rows, err := db.Query("SELECT emoji FROM users WHERE userId = \"\";")
	if err != nil {
		log.Fatal("SELECT WHERE USERID = \"\"")
	}
	available := []uint16{}
	for rows.Next() {
		var emoji uint16
		if err := rows.Scan(&emoji); err != nil {
			log.Fatal(err)
		}
		available = append(available, emoji)
	}
	return available
}

func addUser(emoji uint16) error {
	newUserId, err := uuid.NewUUID()
	if err != nil {
		panic("UUID ERROR")
	}
	result := db.QueryRow("INSERT INTO users (userId, emoji) VALUES (?, ?);", newUserId.String(), emoji)
	if result.Err() != nil {
		log.Fatal("INSERT USER ERROR")
	}

	var resultValue int
	result.Scan(&resultValue)
	return nil
}

func addMood(userId string, moodEmoji uint16) error {
	result := db.QueryRow("INSERT INTO moods (userId, mood) VALUES (?, ?);", userId, moodEmoji)
	if result.Err() != nil {
		log.Fatal("INSERT MOOD ERROR")
	}

	var resultValue int
	result.Scan(&resultValue)
	return nil
}
