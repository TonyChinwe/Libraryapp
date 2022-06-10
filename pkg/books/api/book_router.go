package api

import (
	"net/http"
)

func (api BookApi) BookRouters() {
	bookRoute := api.router.PathPrefix("/api/v1/book").Subrouter()
	bookRoute.HandleFunc("", api.CreateBook).Methods(http.MethodPost)
	bookRoute.HandleFunc("/library/{name}", api.GetAllBooksInLibrary).Methods(http.MethodGet)
	bookRoute.HandleFunc("/search", api.SearchBook).Methods(http.MethodGet)
	bookRoute.HandleFunc("/{id}", api.GetBook).Methods(http.MethodGet)
	bookRoute.HandleFunc("/{id}", api.UpdateBook).Methods(http.MethodPut)
	bookRoute.HandleFunc("/{id}", api.DeleteBook).Methods(http.MethodDelete)
}
