package main
//_test.go是测试文件命名规范，go test testing/ 运行testing包下面的测试
//项目可以放到GOPATH下面,

import (
	"testing"
	//"godemo/test-go"
)

func TestCalculate(t *testing.T){ //单个测试
	if Calculate(2) !=5 {
		t.Error("Eexpect 3+2 to equal 5!")
	}
}

func TestTableCalculate(t *testing.T){ //测试列表
	var tests = []struct{ //匿名struct
		input int
		expected int
	}{ //往tests slice中添加struct元素
		{input:1, expected: 3},
		{2, 4},
		{-1, 1},
		{10,13},
		{0,2},
		{999, 1001},
	}

	for _, test := range tests {
		if output := Calculate(test.input); output != test.expected{
			t.Error("Test failed:{} inputted, {} expected, received:{}", test.input,test.expected,output)
		}
	}

}