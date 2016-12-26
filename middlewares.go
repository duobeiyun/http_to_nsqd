package main

import (
	"github.com/labstack/echo"
	"github.com/nsqio/go-nsq"
	"github.com/labstack/gommon/log"
)

var producer *nsq.Producer

func InitNsq(addr string) {
	cfg := nsq.NewConfig()
	var err error
	producer, err = nsq.NewProducer(addr, cfg)
	if err != nil {
		log.Fatalf("failed to create nsq.Producer - %s", err)
	}
}

func NsqPublish(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			topic, ok := c.Get("topic").(string)
			if !ok {
				log.Error("topic is not string.")
				log.Error(topic)
				return
			}
			messageStr, ok := c.Get("logMessage").(string)
			processed := Process(messageStr, c.RealIP())
			message := []byte(processed)
			if !ok {
				log.Error("message is not string.")
				log.Error(message)
				return
			}
			err := producer.Publish(topic, message)
			if err != nil {
				log.Errorf("nsq send failed. %s", err)
			}
		}()
		return next(c)
	}
}


func Validate(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		req := c.Request()
		host := req.Host
		uri := req.URL.Path
		if topic, ok := Opts.TopicMapping[uri]; ok {
			c.Set("topic", topic)
			return next(c)
		}
		log.Infof("HOST=%s. 404 not found %s.", host, uri)
		return echo.ErrNotFound
	}
}