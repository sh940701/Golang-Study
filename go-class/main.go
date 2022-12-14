package main

import (
	"fmt"
	controller "go-class/controller"
	model "go-class/model"
)

func main() {
	ctl := controller.NewCTL()

	fmt.Println(ctl.CalcSum(2, 3))
	fmt.Println(ctl.CalcMul(2, 3))
	fmt.Println(ctl.CalcDiv(2, 3))
	fmt.Println(ctl.CalcSub(2, 3))


	mod := model.NewModel()

	fmt.Println(mod.Run("run"))
	fmt.Println(mod.Jump("jump"))
	fmt.Println(mod.Sleep("sleep"))
	fmt.Println(mod.Walk("walk"))
	fmt.Println(mod.Fly("fly"))
}