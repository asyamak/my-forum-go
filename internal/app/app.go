package app

import (
	"fmt"
	"forum/config"
	"forum/internal/controller"
	"forum/internal/repository"
	"forum/internal/usecase"
	"forum/pkg/database"
	"forum/pkg/httpserver"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	config *config.Config
	// db         *sql.DB
}

func NewApp(conf *config.Config) *App {
	return &App{
		config: conf,
		// db: db,
	}
}

func (a *App) Start() {
	// initialise repository
	fmt.Println("app start - db init")
	db, err := database.InitDB(a.config)
	if err != nil {
		log.Fatalf("app - start - repository init error: %v\n", err)
	}
	// should defer to closure be in server.Shutdown??? or be in dbinit
	defer db.Close()

	// initialise tables - what the best way to place creation of tables?
	if err := database.CreateTables(db); err != nil {
		log.Fatalf("app - start - create tables: %v\n", err)
	}

	// repository layer
	userRepository := repository.NewUser(db) // vynesty v otdelnuy funkciy kostruktory
	post := repository.NewPost(db)
	// usecase layer
	useCase := usecase.NewUseCase(userRepository, post)
	// handler
	handler := controller.NewHandler(useCase.User, useCase.Post)
	router := controller.SetupRouter(handler)
	// http server
	server := httpserver.NewServer(a.config, router)

	//
	//
	//
	//
	//
	//

	// waiting signal for graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	// checking if we receiving signal to shut down server
	select {
	case s := <-interrupt:
		log.Printf("app - start - signal: " + s.String())
	case err = <-server.Notify():
		log.Printf("app -start - server.Notify: %v", err)
	}
	// shutdown server
	err = server.Shutdown()
	if err != nil {
		log.Printf("app -start - server.Shutdown: %v\n", err)
	}
}
