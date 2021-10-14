package main

type Character struct {
	Name        string `json:"name" tag:"foo"`
	Surname     string `json:"surname"`
	Job         string `json:"job,omitempty"`
	YearOfBirth int    `json:"year_of_birth,omitempty"`
}
