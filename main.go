//author: bondhan.novandy@gmail.com/Depok/Indonesia

package main

import (
	"fmt"
	"os"

	"github.com/bondhan/godddnews/infrastructure/driver"
	"github.com/bondhan/godddnews/infrastructure/manager"
	"github.com/bondhan/godddnews/interfaces"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sirupsen/logrus"
)

func main() {
	_, isProd := os.LookupEnv("PRODUCTION_ENV")
	if isProd {
		fmt.Println("PRODUCTION_ENV TRUE", os.Getenv("PRODUCTION_ENV"))
		driver.NewLogDriver(os.Getenv("LOG_NAME"), logrus.ErrorLevel).InitLog()
	} else {
		fmt.Println("PRODUCTION_ENV FALSE", os.Getenv("PRODUCTION_ENV"))
		driver.NewLogDriver(os.Getenv("LOG_NAME"), logrus.TraceLevel).InitLog()
	}

	manager.GetContainer()
	logrus.Info("manager was called")

	logrus.Info("application started at port:", os.Getenv("APPLICATION_PORT"))
	interfaces.Init(os.Getenv("APPLICATION_PORT"))
}
