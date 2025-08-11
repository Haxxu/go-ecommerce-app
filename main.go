package main

import (
	"fmt"
	_ "go-ecommerce-app/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//var myFamily [3]string
	//myFamily[0] = "A"
	//myFamily[1] = "B"
	//myFamily[2] = "C"

	//myFamily := [3]string{"A", "B", "C"}

	var myFriends []string
	myFriends = append(myFriends, "John", "Jane")

	myCources := [][]string{
		{"Go", "NodeJS"},
		{"AWS", "GCP"},
	}

	fmt.Println(myCources)

	app.Listen(":9000")
}
