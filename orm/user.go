package main //要导出的包不能是main，除非放到go下面

import (
	"fmt"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "all users endpoint hit")
}

func NewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "new users endpoint hit")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "delete users endpoint hit")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "update users endpoint hit")
}
