fmt: vet tidy
	gofumpt -l -w .

tidy:
	go list -f '{{{{.Dir}}' -m | xargs -I '{}' sh -c 'cd "{}" && go mod tidy'

vet:
	go list -f '{{{{.Dir}}' -m | xargs golangci-lint run
