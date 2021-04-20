package main

import (
	"fmt"
	_ "wrapError/common"
	"wrapError/service"
)

func main() {
	_, err := service.NewUserService("").GetUserName()
	if err != nil {
		fmt.Printf("%+v\n", err)
	}

	if _, err = service.NewUserService("2").GetUserName(); err != nil {
		fmt.Printf("%+v\n", err)
	}

	if err = service.NewUserService("1").DeleteUser(); err != nil {
		fmt.Printf("%+v\n", err)
	}
}
