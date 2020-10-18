package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"parking/internal/auth"
	"parking/internal/mail"
	"parking/internal/spaceCounter"
	"parking/internal/tools"
	"parking/model"
	"time"
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

	count, _ := spaceCounter.SpaceCounter.GetSpaceCount(prediction)

	err = s.Database.SaveFrame(jsonReq.Id, count, jsonReq.Content)
	if err != nil {
		log.Printf("Error to insert frame: %s\n", err)
	}

	log.Println(count)
	//err = s.Database.UpdateFreeSpace(jsonReq.Id, points)
	//if err != nil {
	//	log.Printf("Error to insert frame: %s\n", err)
	//}
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

	s.setTokenInCookie(w, user.Token)
	w.Write(jsonResp)
}

func (s *Server) setTokenInCookie(w http.ResponseWriter, token string) {
	cookie := http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: time.Now().Add(3 * time.Hour),
	}
	http.SetCookie(w, &cookie)
}

func (s *Server) GetFrame(w http.ResponseWriter, r *http.Request) {
	body, closeFunc, err := tools.ReadRequestBodyJson(r, &model.FrameRequest{})
	if err != nil {
		log.Printf("Can`t read json body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonReq := body.(*model.FrameRequest)
	defer closeFunc()

	frame, err := s.Database.GetFrame(jsonReq.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	response := model.FrameResponse{
		Image: frame,
	}
	jsonResp, err := json.Marshal(response)
	if err != nil {
		w.Write(jsonResp)
	}

	w.Write(jsonResp)
}

func (s *Server) GetCameras(w http.ResponseWriter, r *http.Request) {
	cameras, err := s.Database.GetAllCameras()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	jsonResp, err := json.Marshal(cameras)
	if err != nil {
		w.WriteHeader(500)
	}

	w.Write(jsonResp)
}

func (s *Server) GetUserInfo(w http.ResponseWriter, r *http.Request) {
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
}

func (s *Server) SendEmail(w http.ResponseWriter, r *http.Request) {
	body, closeFunc, err := tools.ReadRequestBodyJson(r, &model.EmailRequest{})
	if err != nil {
		log.Printf("Can`t read json body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	jsonReq := body.(*model.EmailRequest)
	defer closeFunc()

	//log.Println(jsonReq)
	if err := mail.Send(jsonReq.Email, jsonReq.ErrorCode, jsonReq.Comment); err != nil {
		log.Printf("Error to send email: %s\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Server) GetFreeSpace(w http.ResponseWriter, r *http.Request) {
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
	id, ok := auth.ParseToken(token)
	if ok {
		w.WriteHeader(id)
		return
	}

	id = 1

	responses, err := s.Database.GetParkingSpace(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := model.SpacesResponse{
		Spaces: responses,
	}

	jsonResp, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(500)
	}

	w.Write(jsonResp)
}

func (s Server) SetPolygon(w http.ResponseWriter, r *http.Request) {
	//body, closeFunc, err := tools.ReadRequestBodyJson(r, map[string]interface{}{})
	//if err != nil {
	//	log.Printf("Can`t read json body: %s", err)
	//	w.WriteHeader(http.StatusBadRequest)
	//	//return
	//}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}
	log.Println(body)
	//jsonReq := body.(map[string]interface{})
	//defer closeFunc()

	//fmt.Println(jsonReq)

}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri := r.RequestURI
		log.Println(uri)

		w.Header().Set("content-type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, token")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")

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
