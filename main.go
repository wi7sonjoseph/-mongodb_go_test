package main

import "github.com/wi7sonjoseph/mongodb_go_test/service"

func main() {
	client := service.SetupDB()
	service.CrudFunctionality(client)
}
