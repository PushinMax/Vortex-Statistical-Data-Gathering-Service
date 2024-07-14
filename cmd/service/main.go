package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vortex-statistical-data-gathering-service/internal/handler"
	"vortex-statistical-data-gathering-service/internal/repository"
	"vortex-statistical-data-gathering-service/internal/server"
	"vortex-statistical-data-gathering-service/internal/service"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatal(err)
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatal(err)
	}
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.Init()); err != nil {
			log.Fatal(err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)

	<-ch
	_ = srv.Shutdown(context.Background())
	_ = db.Close()
}

func InitConfig() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
