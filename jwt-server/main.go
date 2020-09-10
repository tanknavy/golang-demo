package main

import (
	//"time"
	"fmt"
	"net/http"
	"log"
	jwt "github.com/dgrijalva/jwt-go"
)

func homePage(w http.ResponseWriter, r * http.Request){

	fmt.Fprintf(w, "token verified, can see the secret information")//格式化输出响应request
}

var mySigningKey = []byte("mySecretPhrase")

//javascript闭包，python的decorator, java里面可使用的拦截器
func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler{//返回一个接口类型

	return http.HandlerFunc(func(w http.ResponseWriter, r * http.Request){
		
		if r.Header["Token"] != nil{ //原始request的header中是否包含token
			
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error){
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok{ //签名方法不符
					return nil, fmt.Errorf("There was an error")
				}
				return mySigningKey, nil //内部返回签名给token，实际中是提取这个token
			})

			if err != nil{ //解签出错
				fmt.Fprintf(w, err.Error()) //给网页返回错误
			}
			
			if token.Valid{ //解签ok,还会检测有效期哦！
				fmt.Println(token) //测试用，发现token携带了加密的信息(可定期更换)，还有加入的map,包含过期信息
				//生产中还要取出map中的值，比如username
				fmt.Println(token.Claims.(jwt.MapClaims)["user"]) //取得用户名可以千人前面
				endTime := token.Claims.(jwt.MapClaims)["exp"] //1.599700804e+09,float64
				fmt.Printf("%v,%T", endTime, endTime)
			
				endpoint(w, r) //可以正常返回网页了
			}

		}else{
			fmt.Fprintf(w, "Not Authorized!")//给网页返回没有授权
		}
	})
}


func handleRequests(){
	//http.HandleFunc("/", isAuthorized(homePage))
	//使用isAuthorized先拦截下来，这是不能用HanleFunc了，要用Handle
	http.Handle("/", isAuthorized(homePage))//函数里面包含函数，闭包
	fmt.Println("server start at 8030...")
	log.Fatal(http.ListenAndServe(":8030", nil))//Fatl就是pring错误然后os.Exit()
}

func main(){
	fmt.Println("My Simple Server")

	handleRequests()
}