package init

import (
	dbs "github.com/TonyChinwe/libraryapp/config/db"
	"github.com/TonyChinwe/libraryapp/pkg/library"
	"github.com/TonyChinwe/libraryapp/pkg/library/api"
	"github.com/TonyChinwe/libraryapp/pkg/library/repository"
	"github.com/TonyChinwe/libraryapp/pkg/library/service"

	"github.com/gorilla/mux"
)

func InitLibrary(dbs *dbs.MongoDB, router *mux.Router) library.LibraryServiceImpl {
	libRepo := repository.NewLibraryRepo(dbs)
	libService := service.NewLibraryService(libRepo)
	api.NewLibraryApi(libService, router)
	return libService
}
