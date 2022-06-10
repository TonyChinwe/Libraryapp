package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TonyChinwe/libraryapp/config/constant"
	"github.com/TonyChinwe/libraryapp/middleware/jwt"
	"github.com/TonyChinwe/libraryapp/pkg/users"
	"github.com/TonyChinwe/libraryapp/pkg/users/model"
	"github.com/TonyChinwe/libraryapp/utils/helpers"
	"github.com/TonyChinwe/libraryapp/utils/response"
	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

type UserApi struct {
	UserService users.UserServiceImpl
	router      *mux.Router
}

func NewUserApi(userServiceImp users.UserServiceImpl, router *mux.Router) *UserApi {
	userApi := &UserApi{UserService: userServiceImp, router: router}
	userApi.UserRouters()
	userApi.UserAuthRouters()
	return userApi
}

func (api UserApi) AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	credentials := new(model.LoginReq)
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	authenticateUser, err := api.UserService.AuthenticateUser(credentials, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	tokenDetails, err := jwt.GenerateJwtToken(authenticateUser)
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusCreated, tokenDetails)
}

func (api UserApi) CreateUser(w http.ResponseWriter, r *http.Request) {
	var us model.User
	err := json.NewDecoder(r.Body).Decode(&us)
	fmt.Println("1 address body: ", us.ID)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	_, er := api.UserService.GetUserByUsername(us.Username, r.Context())
	if er == nil {
		response.BaseResponse(w, http.StatusBadRequest, "User with username already exists")
		return
	}

	passwordBytes := []byte(us.Password)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return
	}
	us.Role = constant.UserRole
	us.IsActive = true
	us.CreatedAT = helpers.CurrentDate()
	us.Password = string(hash)
	result, err := api.UserService.CreateUser(us, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		fmt.Println(" ERROR COMING FROM HERE")
		return
	}
	response.BaseResponse(w, http.StatusCreated, result)
}

func (api UserApi) CreateUserAuthor(w http.ResponseWriter, r *http.Request) {
	var us model.User
	err := json.NewDecoder(r.Body).Decode(&us)
	fmt.Println("1 address body: ", us.ID)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	_, er := api.UserService.GetUserByUsername(us.Username, r.Context())
	if er == nil {
		response.BaseResponse(w, http.StatusBadRequest, "User with username already exists")
		return
	}
	passwordBytes := []byte(us.Password)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return
	}
	us.Role = constant.AuthorRole
	us.IsActive = true
	us.CreatedAT = helpers.CurrentDate()
	us.Password = string(hash)

	result, err := api.UserService.CreateUser(us, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusCreated, result)
}

func (api UserApi) CreateUserAdmin(w http.ResponseWriter, r *http.Request) {
	var us model.User
	err := json.NewDecoder(r.Body).Decode(&us)
	fmt.Println("1 address body: ", us.ID)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	existing, er := api.UserService.GetUserByUsername(us.Username, r.Context())
	fmt.Println("Creating from here" + existing.Username)
	if er == nil {
		response.BaseResponse(w, http.StatusBadRequest, "User with username already exists")
		return
	}
	passwordBytes := []byte(us.Password)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return
	}
	us.Role = constant.AdminRole
	us.IsActive = true
	us.CreatedAT = helpers.CurrentDate()
	us.Password = string(hash)

	result, err := api.UserService.CreateUser(us, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusCreated, result)
}

func (api UserApi) GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id := params["id"]
	result, err := api.UserService.GetUser(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}

func (api UserApi) DeleteUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	res, err := api.UserService.DeleteUser(id, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)

}

func (api UserApi) UpdateUser(w http.ResponseWriter, r *http.Request) {

	var us model.User
	params := mux.Vars(r)
	id := params["id"]

	err := json.NewDecoder(r.Body).Decode(&us)
	if err != nil {
		response.BaseResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := api.UserService.UpdateUser(id, us, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusAccepted, res)
}

func (api UserApi) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	result, err := api.UserService.GetAllUsers(r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}

func (api UserApi) SearchUser(w http.ResponseWriter, r *http.Request) {
	var filter interface{}
	query := r.URL.Query().Get("q")
	if query != "" {
		err := json.Unmarshal([]byte(query), &filter)
		if err != nil {
			response.BaseResponse(w, http.StatusBadRequest, err.Error())
			return
		}
	}

	result, err := api.UserService.SearchUser(filter, r.Context())
	if err != nil {
		response.BaseResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	response.BaseResponse(w, http.StatusOK, result)
}
