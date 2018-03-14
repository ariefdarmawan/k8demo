GOOS=linux GOARCH=arm go build -o ../../docker/go-hello/hello
docker build --rm -t ariefdarmawan/go-hello:latest ../../docker/go-hello
docker push ariefdarmawan/go-hello