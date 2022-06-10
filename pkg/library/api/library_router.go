package api

import (
	"net/http"
)

func (api LibraryApi) LibraryRouters() {
	libraryRoute := api.router.PathPrefix("/api/v1/library").Subrouter()
	libraryRoute.HandleFunc("", api.CreateLibrary).Methods(http.MethodPost)
	libraryRoute.HandleFunc("", api.GetAllLibrary).Methods(http.MethodGet)
	libraryRoute.HandleFunc("/search", api.SearchLibrary).Methods(http.MethodGet)
	libraryRoute.HandleFunc("/{id}", api.GetLibrary).Methods(http.MethodGet)
	libraryRoute.HandleFunc("/{id}", api.UpdateLibrary).Methods(http.MethodPut)
	libraryRoute.HandleFunc("/{id}", api.DeleteLibrary).Methods(http.MethodDelete)
}
