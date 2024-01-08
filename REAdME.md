# format, lint and test
go fmt ./...
golangci-lint run ./... -v
go mod tidy
go test ./... -v -count=1
# go build
git add .
git commit -m 'module vx.x.x'
git tag vx.x.x
git push origin vx.x.x
# go module
$env:GOPROXY="goproxy.cn"
go list -m github.com/clarkhao/go-func@vx.x.x