package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func printMemoryUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	alloc := m.Alloc
	totalAlloc := m.TotalAlloc

	fmt.Printf("Alloc: %d KB | TotalAlloc: %d KB \n", bytesToKB(alloc), bytesToKB(totalAlloc))
}

func bytesToKB(b uint64) uint64 {
	return b / 1000
}

func main() {
	go func() {
		start := time.Now()
		countdown := 300 * time.Second
		end := start.Add(countdown)

		for {
			now := time.Now()
			if now.After(end) {
				break
			}
			printMemoryUsage()
			time.Sleep(time.Second)
		}
	}()

	app := fiber.New()
	limit := limiter.New(limiter.Config{
		KeyGenerator: func(c *fiber.Ctx) string {
			return fmt.Sprint(time.Now().UnixNano())
		},
		Max:        5,
		Expiration: time.Second,
	})

	app.Get("/test", limit, func(c *fiber.Ctx) error {
		time.Sleep(time.Second)
		return c.SendStatus(200)
	})

	app.Listen(":9999")
}
