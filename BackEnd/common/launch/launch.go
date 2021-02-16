// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package launch provides generic functions that can be helpful for
// service launchers.
package launch

import (
  "flag"
  "log"
  "github.com/go-chi/chi/middleware"
  "github.com/go-chi/chi"
  "net/http"
  "time"
  "strconv"
)

// Fills a Mux with a good base middleware stack, recommended by go-chi.
func BaseMiddle(r *chi.Mux){
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  r.Use(middleware.Timeout(time.Minute))
}

// Launch the service/services from r, using the port specified at command line
// with option -p <port> or 443 if no port was specified, and printing
// service_name in logs.
func Launch(service_name string, r *chi.Mux) {
  var p int
  flag.IntVar(&p, "p", 443, "Specify port. Default is 443")
  flag.Parse()

  log.SetPrefix(service_name+": ")
  log.SetFlags(0)
  log.Printf("Starting %v\n", service_name)

  port := ":"+strconv.Itoa(p)

  log.Printf("Port %v\n", port)

  cert := "cert.pem"
  key := "key.pem"

  log.Fatal(http.ListenAndServeTLS(port, cert, key, r))

}
