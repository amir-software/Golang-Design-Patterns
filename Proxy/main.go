package main

func main() {
	postGresClient := dataBaseProxy{user: "AllowedUser"}
	postGresClient.exceute_query()

	postGresClient2 := dataBaseProxy{user: "AnotherUser"}
	postGresClient2.exceute_query()
}
