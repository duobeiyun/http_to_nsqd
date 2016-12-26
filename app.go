package main

import (
	"github.com/jessevdk/go-flags"
	"os"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var Opts struct {
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`
	HTTPAddress string `long:"host" description:"HTTP Address to listen" required:"true" default:"127.0.0.1:3000"`
	NSQDAddress string `long:"nsqd" description:"NSQD Address for publish" required:"true" default:"127.0.0.1:4150"`
	TopicMapping map[string]string `long:"urltopic" description:"URL and Topic mapping." required:"true"`
	Param string `long:"param" description:"Param in GET or POST" required:"true" default:"result"`
}

func main() {
	_, err := flags.Parse(&Opts)
	if err != nil {
		os.Exit(1)
	}
	e := echo.New()
	InitNsq(Opts.NSQDAddress)
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{Level: 6}))
	e.Use(middleware.Recover())
	e.POST("/*", Postlog, Validate, NsqPublish)
	e.GET("/*", Getlog, Validate, NsqPublish)
	e.Logger.Fatal(e.Start(Opts.HTTPAddress))
}
