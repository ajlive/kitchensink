fmt: vet tidy
	gofumpt -l -w .
	npx prettier -w .

install:
	npm install

tidy:
	go list -f '{{{{.Dir}}' -m | xargs -I '{}' sh -c 'cd "{}" && go mod tidy'

vet:
	go list -f '{{{{.Dir}}' -m | xargs golangci-lint run
