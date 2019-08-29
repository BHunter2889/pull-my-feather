build:
	gofmt -s -w ./
	go mod tidy -v
	go mod verify
	env GOOS=linux go build -v -ldflags="-d -s -w" -gcflags="-m=2" -x -a -tags netgo -installsuffix netgo -o bin/pull-my-feather ./main.go ./apl.go ./pillow.go

.PHONY: clean
clean:
	rm -rf ./bin

.PHONY: deploy
deploy: clean build
	sls deploy --verbose