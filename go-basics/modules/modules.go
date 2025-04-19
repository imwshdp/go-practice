package modules

import (
	"fmt"

	fiber "github.com/gofiber/fiber"
	// fiber2 "github.com/gofiber/fiber/v2"
)

func Modules() {
	// run go mod tidy for exclude unused packages

	fmt.Println(fiber.New())
	// fmt.Println(fiber2.New())

}
