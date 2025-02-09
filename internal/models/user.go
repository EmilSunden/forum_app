package models

type User struct {
	Username       string
	ProfilePicture ProfilePicture
}

type ProfilePicture struct {
	Image []byte
}
