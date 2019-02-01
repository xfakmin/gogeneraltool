package main

import (
	"fmt"

	check "github.com/xfakmin/gogeneraltool/checkparam"
)

type Params struct {
	Param1 *string `check:"param1"`
	Param2 *string `check:"param2"`
}

func CheckParams1(p interface{}) error {
	param := p.(*string)
	if len(*param) < 10 {
		fmt.Printf("this param length less that 10. [param=%+v]\n", *param)
	} else if len(*param) > 10 {
		fmt.Printf("this param length greater that 10. [param=%+v]\n", *param)
	} else {
		fmt.Printf("nothing.\n")
	}

	return nil
}

func CheckParams2(p interface{}) error {
	param := p.(*string)
	if len(*param) < 10 {
		fmt.Printf(" in the checkparam2 [param=%+v]\n", *param)
	}
	return nil
}

func main() {
	check.RegisterCheckFunc("param1", check.CheckFunc(CheckParams1))
	check.RegisterCheckFunc("param2", check.CheckFunc(CheckParams1))

	// 每个参数可以注册多个校验函数
	check.RegisterCheckFunc("param1", check.CheckFunc(CheckParams2))
	check.RegisterCheckFunc("param2", check.CheckFunc(CheckParams2))

	param1 := "abcdefg"
	param2 := "abcdefghijk"

	test := &Params{
		Param1: &param1,
		Param2: &param2,
	}

	err := check.CheckParams(test)
	if err != nil {
		fmt.Printf("The parameter format is incorrect.\n")
	}
}
