package main

import (
	"tempApi/config"
	"tempApi/dto"
	"tempApi/helper"
	"tempApi/router"
	"tempApi/server"
)

func main() {

	//init envVars
	//init db connections
	//init certificates

	srv := server.New()
	router.Routes(srv)
	helper.Val = dto.Info{
		Name:        "",
		Designation: "",
		Branch:      "",
	}
	srv.Logger.Fatal(srv.Start(":" + config.ServerPort))
}
