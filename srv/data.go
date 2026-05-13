package main

type (
	User struct {
		Id    string
		Emoji uint16
	}

	Users []User

	Mood struct {
		UserId    string
		MoodEmoji uint16
	}

	Moods []Mood
)
