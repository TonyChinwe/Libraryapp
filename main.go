package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/TonyChinwe/libraryapp/middleware"
	"github.com/TonyChinwe/libraryapp/middleware/jwt"
	"github.com/TonyChinwe/libraryapp/utils/logging"
	"github.com/urfave/negroni"

	dbs "github.com/TonyChinwe/libraryapp/config/db"
	"github.com/TonyChinwe/libraryapp/config/env"
	booksInit "github.com/TonyChinwe/libraryapp/pkg/books/init"
	libraryInit "github.com/TonyChinwe/libraryapp/pkg/library/init"
	usersInit "github.com/TonyChinwe/libraryapp/pkg/users/init"
	"github.com/gorilla/mux"
)

func main() {
	env.LoadEnv()
	port := env.GetEnvWithKey("SERVER_PORT")
	routers := mux.NewRouter().StrictSlash(true)
	dbConn := dbs.NewMongoDBConnection()

	registerServices(dbConn, routers)

	handler := registerMiddleWares(routers)

	srv := &http.Server{
		Handler:      handler,
		Addr:         port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("Application is running on Port %v....\n", port)
	log.Fatal(srv.ListenAndServe())

}

func registerMiddleWares(router *mux.Router) *negroni.Negroni {
	logging.Logger()
	n := negroni.Classic()
	n.Use(middleware.Cors())
	router.Use(jwt.ProtectApi)
	n.UseHandler(router)
	router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("server running"))
		if err != nil {
			return
		}
	})
	return n
}

func registerServices(dbConn *dbs.MongoDB, routers *mux.Router) {
	userServe := usersInit.InitUser(dbConn, routers)
	libraryServe := libraryInit.InitLibrary(dbConn, routers)
	booksInit.InitBook(dbConn, routers, libraryServe, userServe)
}
