package main

import (
	"context"
	"fmt"
	"github.com/bondhan/godddnews/domain/constants"
	"github.com/bondhan/godddnews/infrastructure"
	"github.com/bondhan/godddnews/infrastructure/client"
	errorcodes "github.com/bondhan/godddnews/infrastructure/error"
	"github.com/bondhan/godddnews/infrastructure/persistence"
	"github.com/bondhan/godddnews/interfaces/handlers/database"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

const (
	Production = "PRODUCTION"
)

func main() {
	env := os.Getenv("ENV")
	logName := os.Getenv("LOG_NAME")
	isProd := false
	if env == Production {
		isProd = true
	}

	//create instance of logger client
	logger := client.NewLogger(
		client.LogName(logName),
		client.IsProduction(isProd))
	logger.Info("logger client created")

	//instantiate database sql client
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
	sqlDriver := os.Getenv("DB_DRIVER")
	sqlHandler, err := persistence.NewSQLHandler(sqlDriver, psqlInfo)
	if err != nil {
		logger.Fatalf("err: %s", err)
	}
	logger.Info("database connected")

	//close the db connection gracefully when exit
	defer func(dbHandler database.SQLHandler, l *logrus.Logger) {
		l.Info("closing database connection")
		err := dbHandler.Close()
		if err != nil {
			l.Errorf("err: %s", err)
		}
	}(sqlHandler, logger)

	//respClient is an object to send rest response
	respClient := client.NewResp(errorcodes.ErrorCodesPayment)

	//instantiate manager for common object manager
	manager := client.NewManager()
	manager.SetObject(constants.Logger, logger)
	manager.SetObject(constants.SqlHandler, sqlHandler)
	manager.SetObject(constants.RespondClient, respClient)
	logger.Info("manager created")

	//init all layers and routers
	r := infrastructure.InitApplicationAndRouters(manager)

	/************************************/
	//start the web server
	/************************************/
	appPort := os.Getenv("APPLICATION_PORT")
	server := &http.Server{Addr: ":" + appPort, Handler: r}
	go func() {
		logger.Info("application started at port:", appPort)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			logger.Fatal(err)
		}
	}()
	/************************************/

	// Setting up a channel to capture system signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	//wait forever until 1 of signals above are received
	<-stop

	// send warning that we are closing
	logger.Warnf("got signal: %v, closing DB connection gracefully", stop)

	ctxShutdown, cancelWTimeout := context.WithTimeout(context.Background(), 10*time.Second)
	defer func() {
		cancelWTimeout()
	}()
	//try to shut down the server
	logger.Warn("shutting down http server")
	if err := server.Shutdown(ctxShutdown); err != nil {
		logger.Error(err)
	}
}
