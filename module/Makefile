CGO_LDFLAGS="-L$$GOPATH/src/github.com/nmarley/bls-signatures/build"
CGO_CXXFLAGS="-I$$GOPATH/src/github.com/nmarley/bls-signatures/src -I$$GOPATH/src/github.com/nmarley/bls-signatures/contrib/relic/include -I$$GOPATH/src/github.com/nmarley/bls-signatures/build/contrib/relic/include"

.PHONY: default fmt run

default:  fmt run ## Run default target : fmt + run

fmt:  ## Run go fmt to format Go files
	go fmt

build:  ## Build the program
	CGO_CXXFLAGS=$(CGO_CXXFLAGS) CGO_LDFLAGS=$(CGO_LDFLAGS) go build
