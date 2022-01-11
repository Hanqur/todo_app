package main

import (
	"log"

	"github.com/Hanqur/todo_app"
	"github.com/Hanqur/todo_app/pkg/handler"
	"github.com/Hanqur/todo_app/pkg/repository"
	"github.com/Hanqur/todo_app/pkg/service"
	"github.com/spf13/viper"
	_ "github.com/lib/pq"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal("error initializing configs:", err.Error())
	}

	db, err := repository.NewPostrgesDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		log.Fatal("failed to initialize db: ", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo_app.Server)
	if err := srv.Run(viper.GetString("8080"), handlers.InitRoutes()); err != nil {
		log.Fatal("error with running http server:", err.Error())

	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
