package main

import (
	_ "go-ecommerce-app/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//myWishlist := make(map[string]string)
	//myWishlist["first"] = "Health"
	//myWishlist["second"] = "100 Million Dollar"
	//myWishlist["third"] = "Beautiful Wife"
	//myWishlist["fourth"] = "Beautiful Wife"
	//
	//delete(myWishlist, "fourth")
	//
	//fmt.Printf("My wishlist is %v\n", myWishlist)

	type Product struct {
		Name  string
		Price float64
	}

	app.Listen(":9000")
}
