package main

import (
	"monitoring/internal"
	"flag"
	"fmt"
	"os"
	"monitoring/core"
	"monitoring/collector"
	"sync"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
	"monitoring/manager"
)


func main() {
	var wg sync.WaitGroup
	fp, _ := os.OpenFile("./echo.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	wg.Add(1)
	config := internal.SetConfigFile()
	//cfg := readCommand()
	system := core.New()
	client := collector.New(system, config)

	go client.Start(&wg)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output:fp,
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	// Routes
	e.PUT("/push-metric", hello)
	// Start server
	e.Logger.Fatal(e.Start(":8001"))

	wg.Wait()
}

// Handler
func hello(c echo.Context) error {
	res := &manager.Response{}

	if res == nil {
		//
	}
	res.Code = 200
	res.Message = "success"
	c.Response().Header().Set("Cache-Control", "no-cache, no-store, ")
	c.Response().Header().Set("Server", "Goms-manager/0.1")
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)

	return c.JSONPretty(http.StatusCreated, res, "\t")
	//return json.NewEncoder(c.Response()).Encode(body)
}

func readCommand() {
	application := new(string)
	flag.StringVar(application, "app","system","-app is flags that purpose of this server\n")

	flag.Parse()

	// command line argument 의 갯수가 0개 이거나 설정하지 않은 남은 argument 가 있다면 return
	if flag.NFlag() == 0 || flag.NArg() != 0{
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println(*application)
}