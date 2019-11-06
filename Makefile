.PHONY: default fmt run

default:  fmt run ## Run default target : fmt + run

fmt:  ## Run go fmt to format Go files
	go fmt

run:  ## Run the program
	go run main.go
