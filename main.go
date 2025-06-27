// package main

// import "fmt"

// func main ()  {
// 	fmt.Pritln("hello world")
// 	app := fiber.New()
// 	app.Get("/abc",func(c *fiber.ctx) error{
// 			return c. sendstarting("hello world")

//		})
//		app.Listen(":3000")
//	}

package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main(){
	fmt.Printf("hello")

	     app:=fiber.New()
		app.Get("/abc",func(c *fiber.Ctx) error{
			return c. SendString("hello world")

		})
		app.Listen(":3000")
	
} 
