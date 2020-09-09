package lambda
// aws-lambda
import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

type Request struct { //struct映射到json
	ID float64 `json:"id"`
	Value string `json:"value"`
}

type Response struct {
	Message string `json:"message`
	Ok bool `json:"ok"`
}

//测试时{"id":12345, "value":"some-value"}
func Handler(request Requst) (Response,error){
	return Response{
		Message: fmt.Sprintf("Process Request ID %f", request.ID)
		Ok: true,
	}, nil
}

fun main(){
	lambda.Start(Hanlder)
}