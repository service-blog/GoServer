package service

import (
    "net/http"
    "log"
    "blog/model"
    "encoding/json"
    "time"
    "github.com/dgrijalva/jwt-go"
    "github.com/dgrijalva/jwt-go/request"
)
var (
    SecretKey string = "blog"
)
func JsonResponse(response interface{}, w http.ResponseWriter, code int) {
    json, err := json.Marshal(response)
    if err != nil {
        log.Fatal(err)
        return
    }

    w.Header().Set("Access-Control-Allow-Methods","PUT,POST,GET,DELETE,OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With,Content-Type,Authorization")
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
    w.Write(json)
}

func createToken(username string) (model.Token, error){
    token := jwt.New(jwt.SigningMethodHS256)
    claims := make(jwt.MapClaims)
    claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
    claims["iat"] = time.Now().Unix()
    token.Claims = claims

    tokenString, err := token.SignedString([]byte(SecretKey))
    if err != nil {
        log.Fatal(err)
        return model.Token{}, err
    }
    return model.Token{Token:tokenString},nil
}

func checkToken(r *http.Request) (string, bool){
    log.Println("getNameByToken: ",r.Header.Get("Token"))
    token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
        func(token *jwt.Token) (interface{}, error) {
            return []byte(SecretKey), nil
        })
    if err != nil {
        log.Println("token invailid")
        return "", false
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        log.Println("token")
        return string(claims["sub"].(string)), true
    }
    return "", false

}