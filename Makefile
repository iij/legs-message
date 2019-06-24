setup:
	go get -u golang.org/x/lint/golint
	go get -u golang.org/x/tools/cmd/goimports
lint:
	go vet ./
	golint -set_exit_status ./...
fmt: lint
	goimports -w ./
vet:
	go tool vet ./
test:
	env ENV=test go test -cover -count=1 ./...
