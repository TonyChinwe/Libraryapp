package init

import (
	dbs "github.com/TonyChinwe/libraryapp/config/db"
	"github.com/TonyChinwe/libraryapp/pkg/books"
	"github.com/TonyChinwe/libraryapp/pkg/books/api"
	"github.com/TonyChinwe/libraryapp/pkg/books/repository"
	"github.com/TonyChinwe/libraryapp/pkg/books/service"
	"github.com/TonyChinwe/libraryapp/pkg/library"
	"github.com/TonyChinwe/libraryapp/pkg/users"

	"github.com/gorilla/mux"
)

func InitBook(dbs *dbs.MongoDB, router *mux.Router, libryService library.LibraryServiceImpl, userService users.UserServiceImpl) books.BookServiceImpl {
	bookRepo := repository.NewBookRepo(dbs)
	bookService := service.NewBookService(bookRepo, libryService)
	api.NewBookApi(bookService, libryService, userService, router)
	return bookService
}
