package main

import (
	"log"

	"github.com/inouttt/stkt/test_2/internal/config"
	ll "github.com/inouttt/stkt/test_2/internal/log"
	"github.com/inouttt/stkt/test_2/internal/movie"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	configuration := config.New()
	args := configuration.SetArgs()
	dbConf := config.GetDbConf(args)

	mysqlCon, err := config.NewMysqlSession(dbConf)
	if err != nil {
		log.Panicln("Failed initialize DB")
	}

	logRepo := ll.NewRepository(mysqlCon)
	logger := ll.Newservices(logRepo)
	movie := movie.NewMovie(args.UrlTarget, logger)

	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	movie.Handle(app)

	err = app.Listen(":" + args.Port)
	if err != nil {
		panic(err)
	}

}
