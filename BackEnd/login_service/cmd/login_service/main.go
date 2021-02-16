// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package that will launch the service
package main

import (
  "github.com/go-chi/chi"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/login_service/internal/app/service"
  "gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/launch"
)

var service_name = "login_service"

// Main will launch the service
func main() {

  r := chi.NewRouter()

  launch.BaseMiddle(r)
  r.Post("/accesstokens", service.Login)

  launch.Launch(service_name, r)
}
