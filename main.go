package main

import (
	"fmt"
	"os"

	"github.com/mrpandafr/go-mathador/gomathadorredis"
	"github.com/mrpandafr/go-mathador/gomathadorserver"
)

func init() {
	//url de la base redis
	os.Setenv("REDIS_URL", urlRedis)
}

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
	//urlRedis : url de la base redis
	urlRedis = "http://192.168.99.100:6379"
)

var (
	printVersion   bool
	silent         bool
	sqliteDatabase string
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

	webChannel, err := gomathadorredis.NewChannel("mathador", "web")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print("run web and rest server")
	gomathadorserver.NewServer(webChannel)

}
