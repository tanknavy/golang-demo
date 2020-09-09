package main

import (
	"github.com/golang/protobuf/proto"
	"fmt"
	"log"
)

func main(){
	fmt.Println("hello golang")

	//protoc generate 手动从proto文件产生Person的struct
	bob := &Person{//protocol-buffer 
		Name : "Bob",
		Age :24,
	}

	data,err := proto.Marshal(bob) //encoding
	if err != nil {
		log.Fatal("Marshalling error:", err)
	
	}
	fmt.Println(data)

	newBob := &Person{}

	err = proto.Unmarshal(data, newBob) //deconde
	if err != nil {
		log.Fatal("Marshalling error:", err)
	
	}
	fmt.Println(newBob.GetName())
	fmt.Println(newBob.GetAge())

}