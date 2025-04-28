package main

import (
	"api_book/helper"
)

func main() {

	server := InitializedServer()
	errServer := server.ListenAndServe()
	helper.PanicIfError(errServer)
}