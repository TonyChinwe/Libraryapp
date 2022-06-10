package api

import "net/http"

func (api UserApi) UserAuthRouters() {
	subRouter := api.router.PathPrefix("/api/user/auth").Subrouter()
	subRouter.HandleFunc("/login", api.AuthenticateUser).Methods(http.MethodPost)
	subRouter.HandleFunc("/register-user", api.CreateUser).Methods(http.MethodPost)
	subRouter.HandleFunc("/register-author", api.CreateUserAuthor).Methods(http.MethodPost)
	subRouter.HandleFunc("/register-admin", api.CreateUserAdmin).Methods(http.MethodPost)

}
