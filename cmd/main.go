package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/Hanqur/todo_app"
	"github.com/Hanqur/todo_app/pkg/handler"
	"github.com/Hanqur/todo_app/pkg/repository"
	"github.com/Hanqur/todo_app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	// задаём формат логов 
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatal("error initializing configs:", err.Error())
	}

	// подгружаем .env файл
	if err := godotenv.Load(); err != nil {
		logrus.Fatal("error loading env variable: ", err.Error())
	}

	// запускаем базу данных
	db, err := repository.NewPostrgesDB(repository.Config{
		Host: viper.GetString("db.host"),
		Port: viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: viper.GetString("db.dbname"),
		SSLMode: viper.GetString("db.sslmode"),
	})
	if err != nil {
		logrus.Fatal("failed to initialize db: ", err.Error())
	}

	// устанавливаем зависимости
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)


	// запускаем сервер
	srv := new(todo_app.Server)
	go func () {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatal("error with running http server:", err.Error())
	
		}
	}()

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit
	
	logrus.Print("TodoApp shootting down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occurred on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occurred on db connection close: %s", err.Error())
	}
	

}

// читаем config
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
