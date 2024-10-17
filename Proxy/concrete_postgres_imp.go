package main

import "fmt"

type PostgesDb struct {
	DataBase
}

func (pg PostgesDb) exceute_query() {
	fmt.Println("Exceuting query by the Postgres database")
}
