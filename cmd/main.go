package main

import (
	"QuizAppApi"
	"QuizAppApi/pkg/handler"
	"QuizAppApi/pkg/repository"
	"QuizAppApi/pkg/service"
	"context"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initialiation of configs: %s", err.Error())
	}

	//if err := godotenv.Load(); err != nil {
	//	log.Fatalf("error loading env variables: %s", err.Error())
	//}

	db, err := repository.NewPostgresDB(repository.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: os.Getenv("DB_USER"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to init db: %s", err.Error())
	}
	repo := repository.NewRepository(db)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	err = services.Authorization.AdminInit("admin", os.Getenv("ADMIN_KEY"))
	if err != nil {
		log.Printf("failed to init admin : %s", err.Error())
	}

	srv := new(QuizAppApi.Server)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("failed to start server : %s", err.Error())
		}
	}()

	log.Println("QuizAppApi Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("QuizAppApi Stopped")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}

	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("error occured on database connection stopping: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	return viper.ReadInConfig()
}
