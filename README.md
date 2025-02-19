# Go Api Service  

## Project Technologies:  

[go](https://go.dev/)  
[net/http](https://pkg.go.dev/net/http)  
[docker](https://docs.docker.com/)

## Usage:  

1. Build and run the app:  
Docker  
```
docker build -t apiservice .
```
```
docker run -p 5000:5000 apiservice
```  
Go
```
go build -o apiservice cmd/apiservice/main.go
```
```
./apiservice
```    

2. Go to address: [http://localhost:5000/api](http://localhost:5000/api)
