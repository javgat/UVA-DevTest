package main

import (
  "log"
  "github.com/go-chi/chi"
  "github.com/go-chi/chi/middleware"
  "net/http"
  "time"

  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/signin_service/internal/app/service"
)

var service_name = "signin_service"

func main() {
  log.SetPrefix(service_name+": ")
  log.SetFlags(0)

  log.Printf("Starting %v\n", service_name)

  port := ":443"
  r := chi.NewRouter()

  // A good base middleware stack, recommended by go-chi
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  r.Use(middleware.Timeout(time.Minute))

  r.Post("/users", service.RegisterUser)

  cert := "cert.pem"
  key := "key.pem"

  log.Fatal(http.ListenAndServeTLS(port, cert, key, r))

}
