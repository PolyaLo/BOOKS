package models

type User struct {
	Name  string `json: "name" validate:required`
	Login string `json: "login" validate:required,email`
	Pass  string `json: "pass" validate:required`
}

type Book struct {
	BID    string `json: "b_id" validate:required`
	Label  string `json: "laber" validate:required`
	Author string `json: "author" validate:required`
}
