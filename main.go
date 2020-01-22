package main

import (
	"github.com/kondo1018008/golangPractice/db"
	"github.com/kondo1018008/golangPractice/server"
)

func main() {
	db.Init()
	server.Init()
}