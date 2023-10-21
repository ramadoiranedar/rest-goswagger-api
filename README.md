# Template GoSwagger V2 RESTful API

This package is a simple yet powerful representation of your RESTful API contains **GoSwagger 2.0** with customization.

## Documentation

### Quickstart

1. Clone the repository `git clone https://github.com/ramadoiranedar/rest-goswagger-api.git`
1. Initialize the server `make init-server`
1. Generate server and Client `gen-dev`
1. Serve the server API `make run`
1. Now a simple RESTful API is serving at: `http::/localhost:8080`
1. Next? you can customize this RESTful API based on [GoSwagger](https://goswagger.io/)

### Base Structure
```

├── api                 // for swagger api files
├── cmd                 // for main server files and customization files eg. cron job
├── configs             // for configuration files
├── gen                 // for swagger generated files (eg. models, server) configuration swagger
│   ├── models
│   └── restapi
│       ├── operations
├── internal            // for customization RESTful API files
│   ├── handlers
│   └── rest
├── main.go.dist        // this is a overwritten (Customized) main file from default generate by the GoSwagger. This file will auto move to main server folder
├── makefile            // for CLI Commands to this RESTful API files
├── pkg                 // for client generated files by swagger
│   └── client
│       ├── app
└── runtime.go          // for runtime (eg. database, cache, client service, routing, Authorization, etc) to work with Rest API and middlewares
```

### Features
From [GoSwagger](https://github.com/go-swagger/go-swagger): `go-swagger`` brings to the go community a complete suite of fully-featured, high-performance, API components to work with a Swagger API: server, client and data model.

- Generates a server from a swagger specification
- Generates a client from a swagger specification
- Generates a CLI (command line tool) from a swagger specification (alpha stage)
- Supports most features offered by jsonschema and swagger, including polymorphism
- Generates a swagger specification from annotated go code
- Additional tools to work with a swagger spec
- Great customization features, with vendor extensions and customizable templates

### Customization

- Base structure
- Runtime
- Authorization
- Routing
- Database client mysql
- Configuration file (eg. environtment or secret, credentials)
- Generated files and folders
- ...

### Makefile
```
make                            : Foo
make init-server                : [ONCE] init for the first time generate server with main.go file
make gen-server                 : generate server excluding main.go file for development
make client-server              : generate client after generate server for this app, we can add more than one client for this
make gen-dev                    : generate server and client after run the init-server for development
make validate                   : validation Swagger file yaml OpenAPI
make clean                      : cleaning or remove unused object, generated, cached files
make build                      : generate build file executable server
make api-doc                    : serve UI of API documenations
make run                        : serve REST API by the executable file server
make run-dev                    : run the server for development without build the file executable server
make all-dev                    : for easy generate all and run server for development
make make-gen                   : make folder /gen
make remove-gen                 : remove folder /gen
make remove-pkg                 : remove folder /pkg
make remove-cmd                 : remove folder /cmd
make remove-executable-server   : remove file executable server at root directory /
```

## Postman Collections

Example postman collection is [Here](https://api.postman.com/collections/27566000-8cf3cc54-e44b-428d-9749-26fc15eacad7?access_key=PMAT-01HD96P71AWTAQ832P96G3VEGM)

## Project Status
- Made with ♡.
- Usages setup RESTful API using golang for learning and portfolio.

## Released Version

- version 0.1 (21 October 2023)

## Thanks
Special Thanks to [GoSwagger](https://github.com/go-swagger/go-swagger)