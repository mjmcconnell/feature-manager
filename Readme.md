# Feature Manager

This project is setup as a sample project of golang based architecture rather than an actual application purpose.

One aim is to keep as much of the third-party libraries contained to specific modules or the internal package.
The developer should have the ability to swap out any of the third-party packages without impacting any of the other services.

For example, currently the `echo` package is used to handle http routing and middleware.
As it is only directly references with the `internal/routes.go` file, it can be swapped out of another provider, and as long as the `Router` methods do not change, modules that use http do not need to be updated.

Another aim is to build out functionality in a modular fashion. This is achived by keeping any common functionality within the `internal/` package, and then keeping functional specific code within the modules.

For example, a module that adds functionality for a to-do list, will define all routes and models within the module, but leverage the internal package for the http router, and model driver

## Adding a new module

All modules should live under the `modules/` directory.

The module should contain a file matching the name of the module at the root of the module, which impliments the module interface.

For an example, check out `modules/hello/hello.go`

To add the module into the app, simply add the reference into the `imports` section of `cmd/feature-manager/main.go`

Example

    import (
        "github.com/mjmcconnell/feature-manager/internal"
        // modules
        ....
        _ "github.com/mjmcconnell/feature-manager/modules/my-new-module"
    )

## Local development

### Native

#### Server

To start the local dev server from your local environment;

    $ cd server && make run

#### UI

### Docker

With [docker](https://docs.docker.com/engine/install/) and [docker-compose](https://docs.docker.com/compose/install/) installed, you can get started by running:

    $ docker-compose up

### Running

Once the local development environment is up and running, you should be able to navigate to http://localhost:8080 in your browser to see the API running

The project uses [air](https://github.com/cosmtrek/air) to support live-reloading of the server as you develop

## Key packages

* Echo - Http handling
  * https://echo.labstack.com/
  * https://pkg.go.dev/github.com/labstack/echo
* Gorm - ORM library
  * https://gorm.io/
  * https://pkg.go.dev/github.com/go-gorm/gorm
* Atlas - Inspecting, planning, linting and executing schema changes
  * https://atlasgo.io/
  * https://pkg.go.dev/github.com/ariga/atlas
* jwt-go - JWT based authentication
  * https://golang-jwt.github.io/jwt/
  * https://pkg.go.dev/github.com/golang-jwt/jwt
