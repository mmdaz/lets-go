package main

import "lets_go/interface_practice/modules"

func main() {
	mainModule := modules.NewMainModule()
	dependentModule := modules.NewDependentInterface(mainModule)
	dependentModule.DoPractice("hello!")
}
