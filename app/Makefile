.PHONY: build
build:
	GOARCH=amd64 GOOS=darwin go build -o ./dist/a ./a
	GOARCH=amd64 GOOS=darwin go build -o ./dist/a2 ./a2
	GOARCH=amd64 GOOS=darwin go build -o ./dist/b ./b
	GOARCH=amd64 GOOS=darwin go build -o ./dist/control ./control
	ls -al ./dist
	shasum -a 256 ./dist/*
