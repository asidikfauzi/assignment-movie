package main

import (
	"assignment-movie/common/inject"
	"assignment-movie/routes"
)

func main() {
	route := routes.InitPackage()
	inject.DependencyInjection(inject.InjectData{
		Routes: route,
	})

	route.InitRouter()
}
