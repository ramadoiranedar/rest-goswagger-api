# .PHONY: Foo

# make folder /gen
make-gen:
	if [ ! -d gen ]; then mkdir gen; fi

# remove folder /gen
remove-gen:
	if [ -d gen ]; then rm -R gen; fi

# remove folder /pkg
remove-pkg:
	if [ -d pkg ]; then rm -R pkg; fi

# remove folder /cmd
remove-cmd:	
	if [ -d cmd ]; then rm -R cmd; fi

# remove file executable server at root directory /
remove-executable-server:
	rm -rf rest-goswagger-api-server 

 # validation swagger file yaml OpenAPI
 validate:
	swagger validate ./api/swagger.yml

# [ONCE] init for the first time generate server with default main.go file
init-server: validate make-gen
	swagger generate server --main-package=../../cmd/rest-goswagger-api-server -A rest-goswagger-api-server -t ./gen  -f ./api/swagger.yml --principal models.Principal
	cp main.go.dist cmd/rest-goswagger-api-server/main.go

# generate server excluding main.go file for development
gen-server: validate make-gen
	swagger generate server --exclude-main --regenerate-configureapi -A rest-goswagger-api-server -t ./gen  -f ./api/swagger.yml --principal models.Principal

# generate client after generate server for this app, we can add more than one client for this
# Example:
#		swagger generate client -A rest-goswagger-api-server -f ./api/swagger.yml -c pkg/client -m ./gen /models --principal models.Principal
# 		swagger generate client -A for-example1-server -f ./api/example1.yml -c pkg/client -m ./gen /model
# 		swagger generate client -A for-example2-server -f ./api/example2.yml -c pkg/client -m ./gen /model
gen-client: validate	
	swagger generate client -A rest-goswagger-api-server -f ./api/swagger.yml -c pkg/client/rest_goswagger_api_client -m ./gen/models --principal models.Principal
	swagger generate client -A themoviedb -f ./api/themoviedb.yml -c pkg/client/themoviedb_client -m ./gen/model --principal models.Principal

# cleaning or remove unused object, generated, cached files
clean: remove-executable-server remove-gen remove-pkg
	go clean -i .

# generate server and client after run the init-server for development
gen-dev: clean gen-server gen-client 

# generate build file executable server
build:
	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix cgo ./cmd/rest-goswagger-api-server

# run the executable file server
run:
	./rest-goswagger-api-server --port=8080 --host=0.0.0.0 --config="./configs/app.yaml"

# run the server for development without build the file executable server
run-dev:
	go run ./cmd/rest-goswagger-api-server/main.go --config=./configs/app.yaml --host=0.0.0.0 --port=8080

# serve UI of API documenations Swagger OpenAPI
api-doc: validate
	swagger serve api/swagger.yml --no-open --host=0.0.0.0 --port=8081 --base-path=/

# for easy generate all and run server for development
all-dev: gen-dev build run

# TODO: make for development/stagging/production
# . . .

# Documentation HELPER
help:
	@echo "make: compile packages and dependencies"

	@echo "make init-server: [ONCE] init for the first time generate server with main.go file"
	@echo "make gen-server: generate server excluding main.go file for development"
	@echo "make client-server: generate client after generate server for this app, we can add more than one client for this"
	@echo "make gen-dev: generate server and client after run the init-server for development"
	@echo "make validate: validation Swagger file yaml OpenAPI"
	@echo "make clean: cleaning or remove unused object, generated, cached files"
	@echo "make build: generate build file executable server"
	@echo "make api-doc: serve UI of API documenations"
	@echo "make run: serve REST API by the executable file server"
	@echo "make run-dev: run the server for development without build the file executable server"
	@echo "make all-dev: for easy generate all and run server for development"
	@echo "make make-gen: make folder /gen"
	@echo "make remove-gen: remove folder /gen"
	@echo "make remove-pkg: remove folder /pkg"
	@echo "make remove-cmd: remove folder /cmd"
	@echo "make remove-executable-server: remove file executable server at root directory /"
