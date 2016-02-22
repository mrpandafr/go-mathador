package main

import (
	"fmt"

	"github.com/mrpandafr/go-mathador/gomathadorredis"
	"github.com/mrpandafr/go-mathador/gomathadorserver"
)

const (
	//Version : application version
	Version = "0.0.1"
	//Name ; application Name
	Name = "go-mathador"
	//URL : application URL
	URL = "https://github.com/mrpandafr/go-mathador"
	//AuthorName : application author
	AuthorName = "Jean-Sébastien Saillard"
	//AuthorEmail : author email
	AuthorEmail = "jeansebastien.saillard@gmail.com"
)

var (
	printVersion   bool
	silent         bool
	sqliteDatabase string
	redisHost      string
	redisPort      string
)

func appInfo() {
	fmt.Println("Name    : ", Name)
	fmt.Println("Version : ", Version)
	fmt.Println("URL     : ", URL)
	fmt.Println("Author  : ", AuthorName, "(", AuthorEmail, ")")
}

func appGoodbye() {
	fmt.Println("Merci d'avoir utilisé cette application")
}

func main() {
	defer appGoodbye()

	appInfo()

	gomathadorredis.Channel()

	fmt.Print("run web and rest server")
	gomathadorserver.Routes()
}
