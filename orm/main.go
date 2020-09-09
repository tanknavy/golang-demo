package main
import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	//以下导入自定包，包不能是package main, 可以取别名，格式alise "path/dir"
	user "godemo/modules" //路径文件中不是package main，否则is a program, not an importable package
)

func handleRequest(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", hellworld).Methods("GET")
	
	//user一个endpoint下面不同http method
	myRouter.HandleFunc("/users", user.AllUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", user.NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", user.DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", user.UpdateUser).Methods("PUT")
	
	log.Fatal(http.ListenAndServe(":8091",myRouter))
}

func hellworld(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello world orm")
}

func main(){
	fmt.Println("Golang ORM")
	
	user.InitialMigration()
	
	handleRequest()
}