package main

import (
	"squirrel/teste/db"
)

func main() {
	db.Connection()
	db.Create()
	db.FindAll()

}
