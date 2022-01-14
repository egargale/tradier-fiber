package main

import (
	// "context"
	"log"
	"strconv"

	// "github.com/egargale/tradier-fiber/internals/postgresql"
	"github.com/egargale/tradier-fiber/internals/rest"
	"github.com/egargale/tradier-fiber/internals/util"

	// "github.com/go-redis/redis/internal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	// "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var MyConfig util.Config

type Todo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []*Todo{
	{Id: 1, Name: "Walk the dog", Completed: false},
	{Id: 2, Name: "Walk the cat", Completed: false},
}

func main() {

	// Parsing Config with Viper
	//
	conferr := util.LoadConfig("./")
	if conferr != nil {
		log.Fatal("cannot load config:", conferr)
	}
	log.Printf("Tradier Key: %s", util.MyConfig.TradierKey)
	log.Printf("Tradier Account: %s", util.MyConfig.TradierAccount)
	log.Printf("DB Driver: %s", util.MyConfig.DBDriver)

	// DB initialization
	//
	// db, dberr := sql.Open(MyConfig.DBDriver, MyConfig.DBSource)
	// db, dberr := sql.Open("pgx", MyConfig.DBSource)
	// dbconn, dberr := pgx.Connect(context.Background(), MyConfig.DBSource)
	// if dberr != nil {
	// 	log.Fatal("cannot connect to db:", dberr)
	// }
	// defer dbconn.Close(context.Background())

	// store := postgresql.NewRepo(dbconn)

	// get sessionid
	rest.MyTradier_Stream()

	// Start Fiber App
	//
	app := fiber.New()
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello world")
	})

	SetupApiV1(app)

	err := app.Listen(":3300")
	if err != nil {
		panic(err)
	}
}

func SetupApiV1(app *fiber.App) {
	v1 := app.Group("/v1")

	SetupTodosRoutes(v1)
}

func SetupTodosRoutes(grp fiber.Router) {
	todosRoutes := grp.Group("/todos")
	todosRoutes.Get("/", GetTodos)
	todosRoutes.Post("/", CreateTodo)
	todosRoutes.Get("/:id", GetTodo)
	todosRoutes.Delete("/:id", DeleteTodo)
	todosRoutes.Patch("/:id", UpdateTodo)
}

func UpdateTodo(ctx *fiber.Ctx) error {
	type request struct {
		Name      *string `json:"name"`
		Completed *bool   `json:"completed"`
	}

	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	var body request
	err = ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse body",
		})
	}

	var todo *Todo

	for _, t := range todos {
		if t.Id == id {
			todo = t
			break
		}
	}

	if todo == nil {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	if body.Name != nil {
		todo.Name = *body.Name
	}

	if body.Completed != nil {
		todo.Completed = *body.Completed
	}

	return ctx.Status(fiber.StatusOK).JSON(todo)
}

func DeleteTodo(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[0:i], todos[i+1:]...)
			return ctx.SendStatus(fiber.StatusNoContent)
		}
	}

	return ctx.SendStatus(fiber.StatusNotFound)
}

func GetTodo(ctx *fiber.Ctx) error {
	paramsId := ctx.Params("id")
	id, err := strconv.Atoi(paramsId)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse id",
		})
	}

	for _, todo := range todos {
		if todo.Id == id {
			return ctx.Status(fiber.StatusOK).JSON(todo)
		}
	}

	return ctx.SendStatus(fiber.StatusNotFound)
}

func CreateTodo(ctx *fiber.Ctx) error {
	type request struct {
		Name string `json:"name"`
	}

	var body request

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
	}

	todo := &Todo{
		Id:        len(todos) + 1,
		Name:      body.Name,
		Completed: false,
	}

	todos = append(todos, todo)

	return ctx.Status(fiber.StatusCreated).JSON(todo)
}

func GetTodos(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(todos)
}
