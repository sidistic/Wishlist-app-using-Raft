# 

## Instructions To Run App
```bash
    cd welcome-app/src/go.etcd.io/etcd/contrib/raftexample
    goreman start

    cd welcome-app
    bash setup.sh
    go run main.go
    go run server.go
```
## Instructions To Run Tests

Must run setup.sh first before running the following tests.

```bash
    cd welcome-app/login
    go test -v

    cd welcome-app/user
    go test -v

    cd welcome-app/feed
    go test -v    
```

## To Access the App

Open the browser and go to 

```
    http://localhost:8080/welcome
```