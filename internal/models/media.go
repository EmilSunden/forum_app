package models

// Models to serve as the data structure of media e.g Movies, Shows, etc

// Movie model
type Movie struct {
	ID          int
	Title       string
	Description string
	Year        int
	Poster      string
}

// Show model
type Show struct {
	ID          int
	Title       string
	Description string
	Year        int
	Poster      string
}
