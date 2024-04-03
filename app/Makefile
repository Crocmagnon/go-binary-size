.PHONY: build
build:
	mkdir -p ./dist
	GOARCH=amd64 GOOS=darwin go build -o ./dist/a ./a
	GOARCH=amd64 GOOS=darwin go build -o ./dist/a2 ./a2
	GOARCH=amd64 GOOS=darwin go build -o ./dist/b ./b
	GOARCH=amd64 GOOS=darwin go build -o ./dist/control ./control

.PHONY: deptree
deptree:
	mkdir -p ./deptree
	depth ./a > ./deptree/a
	depth ./a2 > ./deptree/a2
	depth ./b > ./deptree/b
	depth ./control > ./deptree/control
