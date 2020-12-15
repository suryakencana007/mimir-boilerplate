# {{cookiecutter.repo_name}}

> Description Project

[![pipeline status](https://{{cookiecutter.repo_path}}/{{cookiecutter.repo_name}}/badges/master/pipeline.svg)](https://{{cookiecutter.repo_path}}/{{cookiecutter.repo_name}}/commits/master) [![coverage report](https://{{cookiecutter.repo_path}}/{{cookiecutter.repo_name}}/badges/master/coverage.svg)](https://{{cookiecutter.repo_path}}/{{cookiecutter.repo_name}}/commits/master)

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com) [![forthebadge](https://forthebadge.com/images/badges/built-with-love.svg)](https://forthebadge.com)

### Prerequisites

- Go 1.14+
- Docker (Developed with version 19.03)
- [Goose](https://github.com/pressly/goose) DB migrations tool
- [Jaeger](https://hub.docker.com/r/jaegertracing/all-in-one) All-in-one image can be used for development
- PostgreSql (Developed with version 11.0)
- Export your `~/.ssh/id.rsa` into `GIT_PRIVATE_KEY` in your `.zshrc` or `.bashrc` or similar (export GIT_PRIVATE_KEY=\`cat ~/.ssh/id_rsa\`)

### Install

Goose:

```
$ go get -u github.com/pressly/goose/cmd/goose
```

### Create Config File

Create an env file development by copying the given example.

```
$ cp config/.env.defaultexample config/.env
```

### Run DB migrations

Use Goose to run migrations:

```
$ make migrate DSN="host=localhost port=54322 user=root password=root1234 dbname={{cookiecutter.repo_name}} sslmode=disable" CMD="up"
```

Use Goose to adding migrate:

```
$ make migrate DSN="host=localhost port=54322 user=root password=root1234 dbname={{cookiecutter.repo_name}} sslmode=disable"  CMD="create add_table_db sql"
```

### Run The Project Using Docker

#### Build The Project

```
$ docker build --build-arg "SSH_PRIVATE_KEY=${GIT_PRIVATE_KEY}" -t {{cookiecutter.repo_name}}:latest .
```

### Run The Docker Image

To run the docker image, you need to use volume to mount the config:

Run: rest application

```
$ docker run -p 8009:8009 -v /full/path/to/this/project/config:/go/bin/config {{cookiecutter.repo_name}} rest
```

### Structure Project

```md
+--
|-- api # docs swagger openapi
|-- cmd
|-- |-- {{cookiecutter.repo_name}}
|-- |-- |-- main.go # main package application
|-- config # configuration file and setup
|-- |-- .env # environment variable for development
|-- |-- .env.default.example # environment variable for development
|-- |-- app.config.yml # config properties
|-- |-- config.go # config structure
|-- db # provider setup i.e postgres, mariadb, redis dkk
|-- entity # struct for domain model
|-- migrations # sql migration database
|-- rest # http adapter
|-- repository # repository local package
|-- security # package lib for security
|-- service # services local package
|-- worker # worker service package
|-- coverage.sh # coverage generator file
|-- Dockerfile # docker file build image
|-- go.mod # go11+ module package management
|-- gomod.sh # initial go module
|-- Makefile # makefile cli to build and deploy
|-- README.md
```
