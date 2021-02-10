## greet service
.PHONY: compile-greet
compile-greet: ## compiles the proto file
	protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

.PHONY: greet-server
greet-server: ## builds and runs server
	go run greet/greet_server/server.go

.PHONY: greet-client
greet-client: ## builds and runs client
	go run greet/greet_client/client.go

## calculator service
.PHONY: compile-calc
compile-calc: ## compiles the proto file
	protoc calculator/calcpb/calc.proto --go_out=plugins=grpc:.

.PHONY: calc-server
calc-server: ## builds and runs server
	go run calculator/calculator_server/server.go

.PHONY: calc-client
calc-client: ## builds and runs client
	go run calculator/calculator_client/client.go

