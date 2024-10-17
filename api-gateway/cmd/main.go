package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	app "yir/api-gateway/internal/app/gateway"
	"yir/api-gateway/internal/controller"
)

func main() {
	// логгер, который еще и ведет запись в файл!
	// конфиг
	// подключение к сервису {1..N} по grpc clientу
	
	// Инициализация гейтвея
	s := app.New("localhost", "8080", controller.InitRouter())

	// Запуск гейтвея
	go s.Run()

	// Gracefull Shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	s.Shutdown(context.Background())

}
