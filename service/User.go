package service

import (
	"net/http"
	"blog/model"
	"io/ioutil"
	"log"
	"encoding/json"
)
func signIn(w http.ResponseWriter, r *http.Request) {
	var user model.User;
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(body, &user); err != nil {
		log.Println(err)
		return
	}

	ok, user := model.GetUserByName(user.Username)
	if !ok {
		log.Println("dont ")
		return
	}
	if user.Password != user.Password {
		log.Println("wrong password")
		return
	}
	log.Println("login sucess")
	if _,ok := checkToken(r); !ok {
		//需要重建token
		token, err := createToken(user.Username)
		log.Println("rebuilt: ", token)
		if err != nil {
			return
		} else {
			JsonResponse(token, w, http.StatusOK)
			return
		}

	} else {
		log.Println("rebuilt: ", r.Header.Get("Token"))
		JsonResponse(model.Token{Token: r.Header.Get("Token")}, w, http.StatusOK)
	}
}

func signUp(w http.ResponseWriter, r *http.Request) {
	var user model.User;
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	if err := json.Unmarshal(body, &user); err != nil {
		log.Println(err)
		return
	}
	if ok, _ := model.GetUserByName(user.Username); ok{
		log.Println("duplicate")
		return
	}
	log.Println("create sucess")
	model.AddUser(model.User{
		Username: user.Username,
		Password: user.Password,
	})
	JsonResponse("", w, http.StatusOK)
}