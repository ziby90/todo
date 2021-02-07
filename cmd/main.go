package main

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/ziby90/todo-app"
	"gitlab.com/ziby90/todo-app/pkg/handler"
	"gitlab.com/ziby90/todo-app/pkg/repository"
	"gitlab.com/ziby90/todo-app/pkg/repository/sql"
	"gitlab.com/ziby90/todo-app/pkg/service"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initialization config : %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading variables: %s", err.Error())
	}
	db, err := sql.NewPostgresDB(sql.Config{
		Host:     viper.GetString("maindb.host"),
		Port:     viper.GetString("maindb.port"),
		Username: viper.GetString("maindb.user"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("maindb.dbname"),
		SSLMode:  viper.GetString("maindb.sslmode"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialization db : %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server :%s  ", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
