package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Print("Starting StatMoods server...")

	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	http.HandleFunc("/moods", moods)
	http.HandleFunc("/mood", mood)

	openDB("", "")
	log.Printf("Listening...")
	http.ListenAndServe(":8080", nil)

	fmt.Print("Server shut down.")
}

func user(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		id := r.URL.Query().Get("id")
		user := readUser(id)
		log.Printf("Read User: %+v\n", user)
	default:
		return
	}
}

func mood(w http.ResponseWriter, r *http.Request) {

}

func users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		log.Print("GET /users")

	case "POST":
		b := make([]byte, 4)
		r.Body.Read(b)
		str := string(b)
		log.Printf("Body: %s, %x", str, str)

		data := binary.BigEndian.Uint32(b)

		log.Printf("Emoji value: %d", data)

		// addUser(emojiVal)
	}
}

func moods(w http.ResponseWriter, r *http.Request) {

}
