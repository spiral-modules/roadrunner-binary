test:
	go test -v -race ./...
build:
	CGO_ENABLED=0 go build -buildmode=pie -trimpath -ldflags "-s" -o rr qcmd/rr/main.go
