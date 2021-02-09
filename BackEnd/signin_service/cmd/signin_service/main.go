package main

import (
  "github.com/go-chi/chi"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/signin_service/internal/app/service"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/launch"
)

var service_name = "signin_service"

func main() {

  r := chi.NewRouter()

  launch.BaseMiddle(r)
  r.Post("/users", service.RegisterUser)

  launch.Launch(service_name, r)
}
