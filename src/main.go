package main

import (
	"crud/src/api"
	"crud/src/database"
)

func main() {
	database.OpenDB()
	api.RunAPI()

}
