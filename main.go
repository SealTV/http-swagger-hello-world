package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/SealTV/http-swagger-hw/docs" // docs is generated by Swag CLI, you have to import it.

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	httpSwagger "github.com/SealTV/http-swagger"
)

var (
	host = flag.String("host", "0.0.0.0:1323", "server address to listen")
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1323
// @BasePath /api/v1
func main() {
	flag.Parse()

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", ping)
	})

	r.Route("/swagger", func(r chi.Router) {
		r.Use(hostChanger)
		r.Get("/*", httpSwagger.Handler(
			httpSwagger.URL("./swagger/doc.json"), //The url pointing to API definition
		))
	})

	log.Fatal(http.ListenAndServe(*host, r))
}

// pin godoc
// @Summary Show a 'pong' message
// @Description get string by ID
// @ID get-string-by-int
// @Accept  json
// @Produce  json
// @Success 200 {string} string c
// @Header 200 {string} string c
// @Router /ping [get]
func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprint(w, "pong")
}

// this middleware change Host value for swagger docs info object
func hostChanger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		docs.SwaggerInfo.Host = r.URL.Hostname()
		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
