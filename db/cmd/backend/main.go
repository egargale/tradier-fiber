package main

import (
	"fmt"
	"log"
	"tradier-fiber/internals/redis"
	"tradier-fiber/internals/rest"
	"tradier-fiber/internals/util"
	// "github.com/gofiber/adaptor/v2"
	// "github.com/davecgh/go-spew/spew"
)

func main() {

	// viper.AddConfigPath(".")
	// viper.SetConfigName("app")
	// viper.SetConfigType("env")
	// viper.AutomaticEnv()
	// err := viper.ReadInConfig()
	// if err != nil {
	// 	log.Fatalf("Error reading config file, %s", err)
	// }
	// err = viper.Unmarshal(&MyConfig)
	// if err != nil {
	// 	log.Fatalf("Unable to decode the config file, %v", err)
	// }
	// log.Printf("database uri is %s", MyConfig.TradierKey)
	if err := util.LoadConfig("./"); err != nil {
		panic(fmt.Errorf("invalid application configuration: %s", err))
	}
	// MyConfig, err := util.LoadConfig("./")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }
	log.Printf("Tradier Key: %s", util.MyConfig.TradierKey)
	log.Printf("Tradier Account: %s", util.MyConfig.TradierAccount)
	redis.TestRedis()

	// Get Postions
	// rest.MyTradier_Market()
	rest.MyTradierSandboxPosition()
	rest.MyTradierProductionPosition()

	// Setup endpoints
	// app := fiber.New()
	// // app.Get("/sse", adaptor.HTTPHandler(handler(dashboardHandler)))
	// app.Get("/mm", tradier_market)
	// app.Get("/market", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello, World!") // send text
	// })
	// app.Listen(":3300")
}
