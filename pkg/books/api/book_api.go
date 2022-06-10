package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TonyChinwe/libraryapp/config/constant"
	//"github.com/TonyChinwe/libraryapp/constant"
	"github.com/TonyChinwe/libraryapp/middleware/jwt"
	"github.com/TonyChinwe/libraryapp/pkg/books"
	"github.com/TonyChinwe/libraryapp/pkg/books/model"
	"github.com/TonyChinwe/libraryapp/pkg/library"
	"github.com/TonyChinwe/libraryapp/pkg/users"

	"github.com/TonyChinwe/libraryapp/utils/response"
	"github.com/gorilla/mux"
)

type BookApi struct {
	BookService books.BookServiceImpl
	LibService  library.LibraryServiceImpl
	UserService users.UserServiceImpl
	router      *mux.Router
}

func NewBookApi(bookServiceImp books.BookServiceImpl, libService library.LibraryServiceImpl, userService users.UserServiceImpl, router *mux.Router) *BookApi {
	bookApi := &BookApi{BookService: bookServiceImp, LibService: libService, UserService: userService, router: router}
	bookApi.BookRouters()
	return bookApi
}

func (api BookApi) CreateBook(w http.ResponseWriter, r *http.Request) {
	username, _ := jwt.ContextEmail(r.Context())
	fmt.Println("username ::: " + username)
	res, er := api.UserService.GetUserByUsername(username, r.Context())
	fmt.Println("user role :: " + res.Username)
	if er != nil {
		fmt.Println("I am here")
		response.BaseResponse(w, http.StatusBadRequest, er.Error())
		return
	}

	if res.Role != constant.AuthorRole {
		response.BaseResponse(w, http.StatusForbidden, "You are not an author")
		return
	}

	var buk model.Book
	err := json.NewDecoder(r.Body).Decode(&buk)
	fmt.Println("1 address body: ", buk.ID)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if buk.LibName == "" {
		response.BaseResponse(w, http.StatusBadRequest, "Provide the library name you want the book to be store in")
		return
	}

	libResult, er := api.LibService.GetLibraryByname(buk.LibName, r.Context())
	if er != nil {
		response.BaseResponse(w, http.StatusBadRequest, "The library not found")
		return
	}
	buk.Library = &libResult
	buk.LibName = libResult.Name
	buk.Author = &res
	result, err := api.BookService.CreateBook(buk, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusCreated, result)
}

func (api BookApi) GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	result, err := api.BookService.GetBook(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}

func (api BookApi) DeleteBook(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	res, err := api.BookService.DeleteBook(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)

}

func (api BookApi) UpdateBook(w http.ResponseWriter, r *http.Request) {

	var buk model.Book
	params := mux.Vars(r)
	id := params["id"]

	err := json.NewDecoder(r.Body).Decode(&buk)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := api.BookService.UpdateBook(id, buk, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)
}

func (api BookApi) GetAllBooksInLibrary(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	name := params["name"]
	result, err := api.BookService.GetAllBooksInLibrary(name, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}

func (api BookApi) SearchBook(w http.ResponseWriter, r *http.Request) {
	var filter interface{}
	query := r.URL.Query().Get("q")
	fmt.Println("query:  " + query)
	if query != "" {
		err := json.Unmarshal([]byte(query), &filter)
		if err != nil {
			response.BaseResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	result, err := api.BookService.SearchBook(filter, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}
