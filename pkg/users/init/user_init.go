package init

import (
	dbs "github.com/TonyChinwe/libraryapp/config/db"
	"github.com/TonyChinwe/libraryapp/pkg/users"
	"github.com/TonyChinwe/libraryapp/pkg/users/api"
	"github.com/TonyChinwe/libraryapp/pkg/users/repository"
	"github.com/TonyChinwe/libraryapp/pkg/users/service"
	"github.com/gorilla/mux"
)

func InitUser(dbs *dbs.MongoDB, router *mux.Router) users.UserServiceImpl {
	userRepo := repository.NewUserRepo(dbs)
	userService := service.NewUserService(userRepo)
	api.NewUserApi(userService, router)
	return userService
}
