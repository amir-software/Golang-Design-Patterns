package main

import "fmt"

type dataBaseProxy struct {
	db   PostgesDb
	user string
}

func (proxy *dataBaseProxy) exceute_query() {
	if proxy.user == "AllowedUser" {
		proxy.db.exceute_query() // Concrete Exectuing Query
	} else {
		fmt.Println("Not Allowed")
	}
}
