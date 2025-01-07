package main

import (
	"fmt"
	"github.com/lishimeng/app-starter/buildscript"
)

func main() {
	err := buildscript.Generate(buildscript.Project{
		Namespace: "lishimeng",
	},
		buildscript.Application{
			Name:    "passport-go",
			AppPath: "cmd/passport",
			HasUI:   true,
		},
		buildscript.Application{
			Name:    "passport-profile",
			AppPath: "cmd/profile",
			HasUI:   true,
		},
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
