package client

import (
  "fmt"
  "log"

  "github.com/andreykaipov/goobs"
  "github.com/spf13/viper"
)

var c *goobs.Client

func NewClient() *goobs.Client {
  if c != nil {
    return c
  }

  var err error
  host := fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("port"))
  log.Printf("Host: %s\n", host)
  if viper.IsSet("password") {
    log.Println("Connecting with password...")
    c, err = goobs.New(host, goobs.WithPassword(viper.GetString("password")))
    if err != nil {
      log.Fatal(err)
    }
  } else {
    log.Println("Connecting without password...")
    c, err = goobs.New(host)
    if err != nil {
      log.Fatal(err)
    }
  }

  return c
}
