package app

import (
	"context"
	"dbhushan9/tech-talk-backend/db"
	"dbhushan9/tech-talk-backend/domain"
	"dbhushan9/tech-talk-backend/handlers"
	"dbhushan9/tech-talk-backend/middlewares"
	"dbhushan9/tech-talk-backend/services"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	Router *mux.Router
	DB     *mongo.Database
	Logger *log.Logger
	//services
	//dao
}

type Service struct {
	TechTalk *services.TechTalkService
}

type DAO struct {
	TechTalk *domain.TechTalkDAO
}

// ConfigAndRunApp will create and initialize App structure. App factory function.
func ConfigAndRunApp() {
	godotenv.Load()

	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	serverHost := fmt.Sprintf("0.0.0.0:%s", port)

	app := new(App)
	app.Initialize(mongoURI, dbName)
	app.Run(serverHost)
}

// Initialize initialize the app with
func (app *App) Initialize(mongoURI string, dbName string) {
	app.DB = db.CreateConnection(dbName, mongoURI)
	app.Router = mux.NewRouter()
	app.Logger = log.New(os.Stdout, "products-api ", log.LstdFlags)

	app.setRouteHandler()
	app.setMiddlewares()
}

func (app *App) setRouteHandler() {
	sr := app.Router.PathPrefix("/tech-talks").Subrouter()

	tt := handlers.New(app.Logger, app.DB)
	sr.HandleFunc("/", tt.Get).Methods(http.MethodGet, http.MethodOptions)
	sr.HandleFunc("/{id}", tt.GetByID).Methods(http.MethodGet, http.MethodOptions)
	sr.HandleFunc("/", tt.Create).Methods(http.MethodPost, http.MethodOptions)
	sr.HandleFunc("/{id}", tt.Update).Methods(http.MethodPut, http.MethodOptions)

	sr2 := app.Router.PathPrefix("/healthcheck").Subrouter()
	h := handlers.NewHealthcheck(app.Logger)
	sr2.HandleFunc("/", h.Get).Methods(http.MethodGet)

}

// setMiddlewares will set global middleware in router
func (app *App) setMiddlewares() {
	app.Router.Use(middlewares.CORS)
	app.Router.Use(middlewares.JSONContentTypeMiddleware)
}

func (app *App) Run(serverHost string) {
	srv := &http.Server{
		Handler:      app.Router,
		Addr:         serverHost,
		ErrorLog:     app.Logger,        // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		app.Logger.Println("Starting server")

		err := srv.ListenAndServe()
		if err != nil {
			app.Logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
		app.Logger.Println("Server started")

	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	log.Println("Stoping MongoDB Connection...")
	app.DB.Client().Disconnect(context.Background())

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	srv.Shutdown(ctx)

}
