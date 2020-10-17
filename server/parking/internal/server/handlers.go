package server

import (
	"encoding/json"
	"log"
	"net/http"
	"parking/internal/auth"
	"parking/internal/spaceCounter"
	"parking/internal/tools"
	"parking/model"
)

func (s *Server) SendFrame(w http.ResponseWriter, r *http.Request) {
	body, closeFunc, err := tools.ReadRequestBodyJson(r, &model.InternalRequest{})
	if err != nil {
		log.Printf("Can`t read json body: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonReq := body.(*model.InternalRequest)
	defer closeFunc()

	prediction, err := s.Predictor.CarDetector(jsonReq.Content)
	if err != nil {
		log.Printf("Error to get prediction: %s\n", err)
	}

	log.Println(prediction)

	log.Println(spaceCounter.SpaceCounter.GetSpaceCount(prediction))
}

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, closeFunc, err := tools.ReadRequestBodyJson(r, &model.LoginRequest{})
	if err != nil {
		log.Printf("Can`t read json body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		response := model.LoginResponse{Err: err.Error()}
		response.Err = err.Error()
		jsonResp, _ := json.Marshal(response)

		w.Write(jsonResp)
		return
	}

	jsonReq := body.(*model.LoginRequest)
	defer closeFunc()

	response := model.LoginResponse{}

	user, err := s.Database.GetUser(jsonReq.Login)
	if err != nil {
		w.WriteHeader(401)
		response.Err = err.Error()
		jsonResp, _ := json.Marshal(response)

		w.Write(jsonResp)
		return
	}

	if err := auth.Hash.ComparePassword(user.Password, jsonReq.Password); err != nil {
		w.WriteHeader(400)
		response.Err = "Incorrect password!"
		jsonResp, _ := json.Marshal(response)

		w.Write(jsonResp)
		return
	}

	token, err := auth.CreateNewToken(user.Id)
	if err != nil {
		w.WriteHeader(500)
		response.Err = err.Error()
		jsonResp, _ := json.Marshal(response)

		w.Write(jsonResp)
		return
	}

	user.Token = token
	response.User = model.UserToUserResponse(user)

	jsonResp, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(500)
		response.Err = err.Error()
		jsonResp, _ := json.Marshal(response)

		w.Write(jsonResp)
	}

	w.Write(jsonResp)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri := r.RequestURI
		log.Println(uri)

		if uri == "/login" {
			next.ServeHTTP(w, r)
			return
		}
		//TODO: remove
		next.ServeHTTP(w, r)
		return
		//

		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		token := c.Value
		code, ok := auth.ParseToken(token)
		if !ok {
			w.WriteHeader(code)
			return
		}

		next.ServeHTTP(w, r)
	})
}