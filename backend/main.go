package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/AndersStigsson/whisky-calendar/delivery/router"
	"github.com/AndersStigsson/whisky-calendar/whisky/controller"
	"github.com/AndersStigsson/whisky-calendar/whisky/repository"
	"github.com/AndersStigsson/whisky-calendar/whisky/usecase"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	ctx := context.Background()
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	db, err := sqlx.Connect("mysql", connString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err.Error())
	}
	whiskyRepository := repository.NewMySQLWhiskyRepository(db)
	whiskyUseCase := usecase.NewWhiskyUseCase(&whiskyRepository)
	router := router.NewMuxRouter()
	whiskyController := controller.NewWhiskyController(whiskyUseCase, ctx, router)
	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	router.GET("/whiskies", whiskyController.GetAllWhiskies)
	router.GET("/whisky/:id", whiskyController.GetWhiskyByID)

	router.SERVE(":42069")
}
