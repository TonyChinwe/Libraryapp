package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TonyChinwe/libraryapp/config/constant"
	"github.com/TonyChinwe/libraryapp/middleware/jwt"

	"github.com/TonyChinwe/libraryapp/pkg/library"
	"github.com/TonyChinwe/libraryapp/pkg/library/model"
	"github.com/TonyChinwe/libraryapp/utils/response"
	"github.com/gorilla/mux"
)

type LibraryApi struct {
	LibraryService library.LibraryServiceImpl
	router         *mux.Router
}

func NewLibraryApi(libraryServiceImp library.LibraryServiceImpl, router *mux.Router) *LibraryApi {
	libraryApi := &LibraryApi{LibraryService: libraryServiceImp, router: router}
	libraryApi.LibraryRouters()
	return libraryApi
}

func (api LibraryApi) CreateLibrary(w http.ResponseWriter, r *http.Request) {

	username, _ := jwt.ContextEmail(r.Context())
	fmt.Println("username ::: " + username)
	res, er := api.LibraryService.GetAUser(username, r.Context())
	fmt.Println("user role :: " + res.Username)
	if er != nil {
		fmt.Println("I am here")
		response.BaseResponse(w, http.StatusBadRequest, er.Error())
		return
	}
	if res.Role != constant.AdminRole {
		response.BaseResponse(w, http.StatusForbidden, "You cannot create library because you are not an Admin")
		return
	}

	var lib model.Library
	err := json.NewDecoder(r.Body).Decode(&lib)
	fmt.Println("1 address body: ", lib.ID)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	result, err := api.LibraryService.CreateLibrary(lib, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusCreated, result)
}

func (api LibraryApi) GetLibrary(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	result, err := api.LibraryService.GetLibrary(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}

func (api LibraryApi) DeleteLibrary(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	res, err := api.LibraryService.DeleteLibrary(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)

}

func (api LibraryApi) UpdateLibrary(w http.ResponseWriter, r *http.Request) {

	var lib model.Library
	params := mux.Vars(r)
	id := params["id"]

	err := json.NewDecoder(r.Body).Decode(&lib)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := api.LibraryService.UpdateLibrary(id, lib, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)
}

func (api LibraryApi) GetAllLibrary(w http.ResponseWriter, r *http.Request) {
	result, err := api.LibraryService.GetAllLibrary(r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}

func (api LibraryApi) SearchLibrary(w http.ResponseWriter, r *http.Request) {
	var filter interface{}
	query := r.URL.Query().Get("q")
	if query != "" {
		err := json.Unmarshal([]byte(query), &filter)
		if err != nil {
			response.BaseResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	result, err := api.LibraryService.SearchLibrary(filter, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}
