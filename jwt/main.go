package main
import (
	"time"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("mySecretPhrase")

func GenerateJWT() (string, error){
	token := jwt.New(jwt.SigningMethodHS256) //加密算法

	//Claims是一个有Valid方法的interface, jwt.MapClaims类型实现了valid方法
	//既然实现了interface的方法，就可以被Claims引用
	claims := token.Claims.(jwt.MapClaims)//使用map[string]interface{}对JSON decoding 
	//fmt.Println(token.Claims.([]int)) //测试
	fmt.Printf("claims: %v %T\n", claims, claims)
	claims["authorized"] = true
	claims["user"] = "Bob Cheng"
	claims["exp"] = time.Now().Add(time.Minute *30).Unix() //30分钟有效期

	tokenString, err := token.SignedString(mySigningKey) //签名产生token

	if err != nil{
		fmt.Errorf("something went wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil

}

func main(){
	fmt.Println("simple client")

	tokenString, err := GenerateJWT()
	if err != nil{
		fmt.Println("error generating token string")
	}
	fmt.Println(tokenString)
}