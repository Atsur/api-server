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

2. **Authorize google auth to read bucket file**

```shell script
gcloud auth application-default login
```

3. **Set Enviornment Variables**

```shell script
go run env.go
```

```shell script
chmod +x ./scripts/run-locally
./scripts/run-locally
```

4. **Run Locally**

```shell script
go run env.go
```

```shell script
./scripts/run-locally
```

-- Format codebase
```shell script
gofmt -s -w .
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

```shell script - gcloud deploy
gcloud builds submit --region=us-west2 --config cloudbuild.yaml   
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
