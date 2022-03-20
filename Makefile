NAME = companies

build:
		go build -o ${NAME} cmd/${NAME}/*.go

build_mac:
		GOOS=darwin GOARCH=amd64 go build -o ${NAME} cmd/${NAME}/*.go

build_windows:
		GOOS=windows GOARCH=amd64 go build -o ${NAME} cmd/${NAME}/*.go

run:
		go run cmd/${DIR_NAME}/*.go

test:
	go test ./... -covermode=atomic -coverpkg=./... -coverprofile ./$gcoverage.out
