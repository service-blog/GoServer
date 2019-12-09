package service
import (
	"net/http"
	"time"
	"blog/model"
	"io/ioutil"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
)


func getArticleList(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if _,ok :=  checkToken(r); !ok {
		log.Println("token wrong")
		return
	}
	ok, articleList := model.GetArticleList(vars["username"])
	if !ok {
		log.Println("get list wrong")
	}
	JsonResponse(articleList, w, http.StatusOK)
}

func getAllArticleById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	if _,ok :=  checkToken(r); !ok {
		log.Println("token wrong")
		return
	}
	id := vars["id"]
	name := vars["Name"]
	ok, article := model.GetAllArticleById(name, id)
	if ok != true {
		log.Println("get art by id wrong")
		return
	}
	JsonResponse(article, w, http.StatusOK)
}

func createArticle(w http.ResponseWriter, r *http.Request){
	var artreq model.ArtCreateReq
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}
	if err = json.Unmarshal(body, &artreq); err != nil {
		log.Println(err)
		return
	}
	username := artreq.Name
	model.CreateArticle(username, model.Article{
		Id: model.GetNextId(username),
		Name:	username,
		Title:	artreq.Title,
		Date:	time.Now().Format("2006/01/02 15:04:05"),
		Content: artreq.Content,
	})
	JsonResponse("", w, http.StatusOK)
}