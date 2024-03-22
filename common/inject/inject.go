package inject

import (
	"assignment-movie/common/database"
	"assignment-movie/repositories/postgres"
	"assignment-movie/routes"
	"assignment-movie/services"
	"github.com/facebookgo/inject"
	"log"
)

type InjectData struct {
	Routes *routes.RouteService
}

func DependencyInjection(liq InjectData) {
	db, err := database.InitDatabase()
	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// POSTGRES
	moviesPostgres := postgres.NewMoviesPostgres(db)

	// SERVICES
	moviesService := services.NewMoviesService(moviesPostgres)

	dependencies := []*inject.Object{
		{Value: moviesService, Name: "movies_service"},
	}

	if liq.Routes != nil {
		dependencies = append(dependencies,
			&inject.Object{Value: liq.Routes, Name: "routes"},
			&inject.Object{Value: liq.Routes.MovieService, Name: "controller_movie_master"},
		)
	}

	// DEPENDENCY INJECTION
	var g inject.Graph
	if err = g.Provide(dependencies...); err != nil {
		log.Fatal("Failed Inject Dependencies", err)
	}

	if err = g.Populate(); err != nil {
		log.Fatal("Failed Populate Inject Dependencies", err)
	}

}
