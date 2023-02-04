package client

import (
  "fmt"
  "log"

  "github.com/andreykaipov/goobs"
  "github.com/spf13/viper"
)
func NewClient() *goobs.Client {
  var c *goobs.Client
  var err error
  if viper.IsSet("password") {
    c, err = goobs.New(fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("port")), goobs.WithPassword(viper.GetString("password")))
    if err != nil {
      log.Fatal(err)
    }
  } else {
    c, err = goobs.New(fmt.Sprintf("%s:%d", viper.GetString("host"), viper.GetInt("port")))
    if err != nil {
      log.Fatal(err)
    }
  }

  return c
}
