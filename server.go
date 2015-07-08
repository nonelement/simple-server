package main

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "syscall"
)

type config struct {
  Port string
  Serve string
}

func (c *config) getPort() string {
  return fmt.Sprintf(":%s", c.Port)
}

var (
  settings config
)

func init() {
  settings = readConfig("settings.json")
  setHTTPHandlers()
}

func readConfig(file string) (settings config) {
  defer (func() {
    if r := recover(); r!=nil {
      syscall.Exit(2)
    }
  })()
  contents, err := ioutil.ReadFile(file)
  if err==nil {
    err = json.Unmarshal(contents, &settings)
  }
  if err!=nil {
      panic("Unable to read configuration file. Cannot start server.")
  }
  return
}

func setHTTPHandlers() {
  if settings.Serve != "" {
    http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(settings.Serve))))
  }
}

func main() {
  http.ListenAndServe(settings.getPort(), nil)
}
