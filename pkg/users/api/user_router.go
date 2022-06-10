package api

import (
	"net/http"
)

func (api UserApi) UserRouters() {
	userRoute := api.router.PathPrefix("/api/v1/user").Subrouter()
	// userRoute.HandleFunc("", api.CreateUser).Methods(http.MethodPost)
	// userRoute.HandleFunc("/admin", api.CreateUserAdmin).Methods(http.MethodPost)
	// userRoute.HandleFunc("/author", api.CreateUserAuthor).Methods(http.MethodPost)
	userRoute.HandleFunc("", api.GetAllUsers).Methods(http.MethodGet)
	userRoute.HandleFunc("/search", api.SearchUser).Methods(http.MethodGet)
	userRoute.HandleFunc("/{id}", api.GetUser).Methods(http.MethodGet)
	userRoute.HandleFunc("/{id}", api.UpdateUser).Methods(http.MethodPut)
	userRoute.HandleFunc("/{id}", api.DeleteUser).Methods(http.MethodDelete)
}
