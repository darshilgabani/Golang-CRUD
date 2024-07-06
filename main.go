package main

import (
	"database/sql"
	"fmt"

	"net/http"

	"goguru/config"
	"goguru/controller"
	"goguru/helper"
	"goguru/model"
	"goguru/repository"
	"goguru/router"
	"goguru/service"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

func main() {
	// insertData()
	// tempFunc()
	// runServer()
	// routes.Run(":2002")

	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("users").AutoMigrate(&model.Users{})

	// Repository
	usersRepository := repository.NewUsersRepositoryImpl(db)

	// Service
	usersService := service.NewUsersServiceImpl(usersRepository, validate)

	// Controller
	usersController := controller.NewUsersController(usersService)

	// Router
	routes := router.NewRouter(usersController)

	server := &http.Server{
		Addr:    ":8889",
		Handler: routes,
	}

	err := server.ListenAndServe()

	helper.ErrorPanic(err)
}

func insertData() {
	dsn := "root:Dar@1611#@tcp(localhost:3306)/DarshilDb"

	db, err := sql.Open("mysql", dsn)
	helper.ErrorPanic(err)
	defer db.Close()
	fmt.Println("Successfully Connected to MySQL database")

	// insert, err := db.Query("INSERT INTO users VALUES('Darshil Mukeshbhai Gabani')")
	// helper.ErrorPanic(err)
	// defer insert.Close()
	// fmt.Println("Successfully inserted database")

	data, err := db.Query("SELECT * FROM users")
	helper.ErrorPanic(err)
	fmt.Println(data.Columns())
	setAndRunServer(data)
	defer data.Close()
}

func setAndRunServer(row *sql.Rows) {
	log.Info().Msg("Started Seever!")
	routes := gin.Default()

	routes.GET("/ping", func(ctx *gin.Context) {
		handlerFunction(ctx, row)
	})

	server := &http.Server{
		Addr:    ":8889",
		Handler: routes,
	}

	err := server.ListenAndServe()

	helper.ErrorPanic(err)
}

func handlerFunction(context *gin.Context, row *sql.Rows) {
	// context.JSON(http.StatusOK, gin.H{
	// 	"Hello": "How are you!?",
	// })
	// context.JSON(http.StatusOK, "Hello!, How are you all?")
	context.JSON(http.StatusOK, row)
}

// func runServer() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello World from Darshil!")
// 	})

// 	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hellow")
// 	})

// 	port := ":5005"

// 	fmt.Println("Server is running on port" + port)

// 	log.Fatal(http.ListenAndServe(port, nil))
// }
