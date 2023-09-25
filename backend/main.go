package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	calendarDateController "github.com/AndersStigsson/whisky-calendar/calendardate/controller"
	calendarDateRepository "github.com/AndersStigsson/whisky-calendar/calendardate/repository"
	calendarDateUseCase "github.com/AndersStigsson/whisky-calendar/calendardate/usecase"
	commentController "github.com/AndersStigsson/whisky-calendar/comment/controller"
	commentRepository "github.com/AndersStigsson/whisky-calendar/comment/repository"
	commentUseCase "github.com/AndersStigsson/whisky-calendar/comment/usecase"
	"github.com/AndersStigsson/whisky-calendar/delivery/router"
	distilleryController "github.com/AndersStigsson/whisky-calendar/distillery/controller"
	distilleryRepository "github.com/AndersStigsson/whisky-calendar/distillery/repository"
	distilleryUseCase "github.com/AndersStigsson/whisky-calendar/distillery/usecase"
	"github.com/AndersStigsson/whisky-calendar/middlewares"
	regionController "github.com/AndersStigsson/whisky-calendar/region/controller"
	regionRepository "github.com/AndersStigsson/whisky-calendar/region/repository"
	regionUseCase "github.com/AndersStigsson/whisky-calendar/region/usecase"
	userController "github.com/AndersStigsson/whisky-calendar/user/controller"
	userRepository "github.com/AndersStigsson/whisky-calendar/user/repository"
	userUseCase "github.com/AndersStigsson/whisky-calendar/user/usecase"
	whiskyController "github.com/AndersStigsson/whisky-calendar/whisky/controller"
	whiskyRepository "github.com/AndersStigsson/whisky-calendar/whisky/repository"
	whiskyUseCase "github.com/AndersStigsson/whisky-calendar/whisky/usecase"
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
	router := router.NewMuxRouter()

	whiskyRepository := whiskyRepository.NewMySQLWhiskyRepository(db)
	whiskyUseCase := whiskyUseCase.NewWhiskyUseCase(&whiskyRepository)
	whiskyController := whiskyController.NewWhiskyController(whiskyUseCase, ctx, router)

	calendarDateRepository := calendarDateRepository.NewMySQLCalendarDateRepository(db)
	cduc := calendarDateUseCase.NewCalendarDateUseCase(&calendarDateRepository)
	cdc := calendarDateController.NewWhiskyController(cduc, ctx, router)

	regionRepo := regionRepository.New(db)
	ruc := regionUseCase.New(&regionRepo)
	rc := regionController.New(ruc, ctx, router)

	dr := distilleryRepository.New(db)
	duc := distilleryUseCase.New(&dr)
	dc := distilleryController.New(duc, ctx, router)

	cr := commentRepository.New(db)
	cuc := commentUseCase.New(&cr)
	cc := commentController.New(cuc, ctx, router)

	ur := userRepository.NewMySQL(db)
	uuc := userUseCase.New(&ur)
	uc := userController.New(uuc, ctx, router)

	router.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})

	router.GET("/whiskies", middlewares.VerifyJWTToken(whiskyController.GetAllWhiskies))
	router.GET("/whisky/{id}", whiskyController.GetWhiskyByID)
	router.GET("/calendar", cdc.GetAllDates)
	router.GET("/calendar/{day}", cdc.GetDateByDayOfMonth)
	router.GET("/region/{id}", rc.GetRegionByID)
	router.GET("/distillery/{id}", dc.GetDistilleryByID)
	router.GET("/comments/{id}", cc.GetCommentByID)
	router.GET("/whisky/{whiskyId}/comments", cc.GetCommentsByWhiskyID)
	router.POST("/comments", cc.StoreComment)
	router.POST("/user/register", uc.Register)
	router.POST("/user/login", uc.Login)

	router.SERVE(":42069")
}
