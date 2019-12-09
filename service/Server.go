package service

import (
	// "net/http"
	"os"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func NewServer() *negroni.Negroni {
	formatter := render.New()

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("web server directory error")
		} else {
			webRoot = root
		}
	}
	// mx.HandleFunc("/", index).Methods("GET")
	mx.HandleFunc("/signin", signIn).Methods("POST")
	mx.HandleFunc("/signup", signUp).Methods("POST")
	mx.HandleFunc("/list", getArticleList).Methods("GET")
	mx.HandleFunc("/getarticle", getAllArticleById).Methods("GET")
	mx.HandleFunc("/createarticle", createArticle).Methods("POST")

	// mx.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
}
