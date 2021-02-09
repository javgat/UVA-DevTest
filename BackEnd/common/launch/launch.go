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

func BaseMiddle(r *chi.Mux){

  // A good base middleware stack, recommended by go-chi
  r.Use(middleware.RequestID)
  r.Use(middleware.RealIP)
  r.Use(middleware.Logger)
  r.Use(middleware.Recoverer)
  r.Use(middleware.Timeout(time.Minute))
}

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
