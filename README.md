# hexagonal-architecture




# Testing
```bash
go test
go test -cover
go test -coverprofile=coverage.out
go tool cover -func=coverage.out
go tool cover -html=coverage.out
```

## Mockgen and gomock

```bash
go install github.com/golang/mock/mockgen@v1.6.0
go get -u github.com/golang/mock/gomock
cd app
mockgen -destination=mocks/app.go -source=product.go

```