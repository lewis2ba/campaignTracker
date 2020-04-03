package models

type Campaign struct{
	title string
	description []byte 
	playerCharacter struct {
		playerName string
		character struct {
			name string
			healthPoints int
			strength int
		}
	session struct{
		number int
		date string 
		}
	}
}