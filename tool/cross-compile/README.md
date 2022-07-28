### 参数
GOOS: darwin、freebsd、linux、windows
GOARCH: amd64、arm
CGO_ENABLED 交叉编译不支持CGO,所以需要禁用它

### 在Mac下编译Linux和Windows64位可执行程序
```bash
# 交叉编译-在Mac下编译Linux和Windows64位可执行程序
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

### 在Windows下编译Mac和Windows64位可执行程序
```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

