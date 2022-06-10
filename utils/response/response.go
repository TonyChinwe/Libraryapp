package response

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/handlers"
)

type Response struct {
	Statuscode int         `json:"status"`
	Message    string      `json:"msg"`
	Data       interface{} `json:"data"`
}

func BaseResponse(w http.ResponseWriter, status int, result interface{}) {
	baseResponse := Response{
		Statuscode: status,
		Data:       result,
	}
	switch status {
	case http.StatusCreated:
		baseResponse.Message = "New record created successfully"
	case http.StatusBadRequest:
		baseResponse.Message = "Bad request"
	case http.StatusInternalServerError:
		baseResponse.Message = "Internal Server Error"
	case http.StatusNotFound:
		baseResponse.Message = "No result found"
	case http.StatusOK:
		baseResponse.Message = "Result found"
	case http.StatusAccepted:
		baseResponse.Message = "Successful"
	default:
		baseResponse.Message = "Internal Server Error"
	}
	//Send header, status code and output to writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(baseResponse)
}

func Headers(r http.Handler) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodOptions})
	return handlers.CORS(headersOk, originsOk, methodsOk)(r)
}
