package routes
import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
)
// Helper Functions
func filter(ss []string, test func(string) bool) (ret []string) {
    for _, s := range ss {
        if test(s) {
            ret = append(ret, s)
        }
    }
    return
}
type Article struct {
	Id		string 	`json:"Id"`
	Title 	string 	`json:"Title"`
	Desc 	string 	`json:"desc"`
	Content string 	`json:"content"`
}


var Articles []Article

//var Articles []Article
func HomePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homepage")
}

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}
func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// fmt.Fprintf(w, "Key: " + key)
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}
func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}
func CreateNewArticleArgBased(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	paramArticle := Article{
		Id:vars["id"],
		Title:vars["title"],
		Desc:vars["desc"],
		Content:vars["content"],
	}

	Articles = append(Articles, paramArticle)
	json.NewEncoder(w).Encode(paramArticle)
}
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	json.Unmarshal(reqBody, &article)

	for _, article := range Articles {
		if article.Id == key {
			Articles = append(Articles, article)
		}
	}
	json.NewEncoder(w).Encode(article)
}

func Main() {
	Articles = []Article {
		Article{Id:"1", Title: "Hello", Desc:"Article Description", Content:"Article Content"},
		Article{Id:"2", Title: "Hello 2", Desc:"Article Description", Content:"Article Content"},
	}
}
