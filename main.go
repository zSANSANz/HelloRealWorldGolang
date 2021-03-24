package main

import (
	"project/config"
	"project/routes"

	//"github.com/iswanulumam/go-training-restful/routes"
	//m "github.com/iswanulumam/go-training-restful/middleware"
)

func main() {
	config.InitDB()
	e := routes.New()
	//e.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}

