package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"parking/model"
	"parking/tools"
)

func SendFrame(w http.ResponseWriter, r *http.Request) {
	body, closeFunc, err := tools.ReadRequestBodyJson(r, &model.InternalRequest{})
	if err != nil {
		log.Printf("Can`t read json body: %s", err)
		w.WriteHeader(500)
		return
	}

	jsonReq := body.(*model.InternalRequest)
	defer closeFunc()

	fmt.Println(jsonReq)
}

func Login(w http.ResponseWriter, r *http.Request) {
	body, closeFunc, err := tools.ReadRequestBodyJson(r, &model.LoginRequest{})
	if err != nil {
		log.Printf("Can`t read json body: %s", err)
		w.WriteHeader(500)
		return
	}

	jsonReq := body.(*model.LoginRequest)
	defer closeFunc()


	response := model.LoginResponse{
		Token: "erfn-234ns-2134",
		Err:   "Incorrect password",
	}
	jsonResp, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(500)
	}

	w.Write(jsonResp)
	fmt.Println(jsonReq)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}