# go_http_sample

### usage

```
git clone https://github.com/rsym/go_http_sample.git
cd ./go_http_sample
go mod tidy

docker run -itd --name redis -p 6379:6379 redis
go run main.go
```

### test

```
go test -v ./
```
