# api-server
A Go (Golang) API REST built with Gin Framework


[![Go Report Card](https://goreportcard.com/badge/github.com/antonioalfa22/go-rest-template)](https://goreportcard.com/report/github.com/antonioalfa22/go-rest-template)
[![Open Source Love](https://badges.frapsoft.com/os/mit/mit.svg?v=102)](https://github.com/ellerbrock/open-source-badge/)
[![Build Status](https://travis-ci.com/antonioalfa22/go-rest-template.svg?branch=master)](https://travis-ci.com/antonioalfa22/go-rest-template)


## Run locally

1. **Create Secret Keys**
```shell script
go run env.go
```

2. **Set Enviornment Variables**

```shell script
go run env.go
```

```shell script
chmod +x ./scripts/run-locally
./scripts/run-locally
```
_______

## Possible Errors
- oauth2: "invalid_grant" "Account has been deleted" error
-- Check google service account and make sure secretes manager accessor permission has been explicitly added to the service account
-- 


_______

## 1. Run with Docker

1. **Build**

```shell script
make build
docker build . -t api-rest
```

2. **Run**

```shell script
docker run -p 8080:8080 api-rest
```

3. **Test**

```shell script
go test -v ./test/...
```

_______

## 2. Generate Docs

```shell script
# Get swag
go get -u github.com/swaggo/swag/cmd/swag

# Generate docs
swag init --dir cmd/api --parseDependency --output docs
```

Run and go to **http://localhost:8080/docs/index.html**
