package main

import (
	"github.com/geejjoo/lamoda_task/internal/handler"
	"github.com/geejjoo/lamoda_task/internal/repository"
	"github.com/geejjoo/lamoda_task/internal/service"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println(new(log.Logger))

	if err := initConfig(); err != nil {
		log.Fatalf("Failed to init configs: %s", err.Error())
	}

	cfg := &repository.Config{
		StorageTablePath: viper.GetString("tables.storage_table"),
		ProductTablePath: viper.GetString("tables.product_table"),
		ReserveTablePath: viper.GetString("tables.reserve_table"),
		DriverName:       viper.GetString("db.driver_name"),
		Host:             viper.GetString("db.host"),
		Port:             viper.GetString("db.port"),
		Username:         viper.GetString("db.username"),
		Password:         viper.GetString("db.password"),
		DBName:           viper.GetString("db.dbname"),
		SSLMode:          viper.GetString("db.sslmode"),
	}

	db, err := repository.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to init db: %s", err.Error())
	}

	repository := repository.NewRepository(db, cfg.StorageTablePath, cfg.ProductTablePath, cfg.ReserveTablePath)
	services := service.NewService(repository)
	handlers := handler.NewHandler(services)

	httpServer := &http.Server{
		Addr:           ":" + viper.GetString("port"),
		Handler:        handlers.InitRoutes(),
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal("Failed to while running http server %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
