# Test Movie
Test the Xsis Movie. This application was built using Golang 1.21.6, PostgreSQL latest, and Docker.

## Link Documentation Postman
[Click Link Postman](https://movies-assignment-test.postman.co/workspace/Team-Workspace~9b151ce6-888e-4056-83c5-bdd8512bef68/collection/33807534-9b237978-51bd-4477-81fa-a4ea70811ddd?action=share&creator=33807534&active-environment=33807534-2426e1ae-cfa5-4610-ba25-18b3858f5659)


## Installation
Use the package manager docker to install app Postgres latest.
```bash
docker-compose up -d --build
```

Create a new Postgres latest database with name `movies`.

## Mac OS / Linux users
Run the `make migrate` command to populate the database tables.
```bash
make migrate
```

Run the `make seed` command to fill data in several tables.
```bash
make seed
```

Run the `make all` command to run the application.
```bash
make all
```

## Windows users
Run the following command to populate the database table.
``` bash
go mod vendor -v
rm -f cmd/migrate/migrate
go build -o cmd/migrate/migrate cmd/migrate/migrate.go
./cmd/migrate/migrate
```

Run the following command to fill data in several tables.
``` bash
go mod vendor -v
rm -f cmd/seed/seed
go build -o cmd/seed/seed cmd/seed/seed.go
./cmd/seed/seed
```
Run the following command to run the application.
``` bash
go mod vendor -v
rm -f cmd/app/app
go build -o cmd/app/app cmd/app/app.go
./cmd/app/app
```