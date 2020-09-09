package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	handleRequest()
}

//Article 解析（encode/decode） 的时候，使用 `sname`，而不是 `Field`
//https://www.cnblogs.com/fengxm/p/9917686.html
type Article struct { 
	Title   string `json:"Title"` //struct映射到json
	Desc    string `json:"desc"` // 在json中显示key为desc
	Content string `json:"content"`
}

//Articles 定义其中元素为Article的slice
type Articles []Article //定义一个其中元素为Article的slice

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{ //slice中添加数据
		Article{Title: "test title", Desc: "test desc", Content: "hello world!"},
		Article{Title: "test title2", Desc: "test desc2", Content: "hello Golang!"},
	}

	fmt.Fprint(w, "Endpoint Hit: all articles endpoint")
	json.NewEncoder(w).Encode(articles) //json格式写出字符串
}

//主页响应
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit") //写出
}

func postArticle(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "test post endpoint worked") //写出
}

//路由处理request
func handleRequest() {

	// http.HandleFunc("/", homePage) //手工handler
	// http.HandleFunc("/articles", allArticles)
	//log.Fatal(http.ListenAndServe(":8090",nil))

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage) //手工handler
	myRouter.HandleFunc("/articles", allArticles).Methods("GET") //相同endpoint不同方法
	myRouter.HandleFunc("/articles", postArticle).Methods("POST")
	log.Fatal(http.ListenAndServe(":8090", myRouter))

}
