GOOS=linux go build -o ../../docker/go-hello/hello
docker rmi -f ariefdarmawan/go-hello
docker build --rm -t ariefdarmawan/go-hello:latest ../../docker/go-hello
docker push ariefdarmawan/go-hello