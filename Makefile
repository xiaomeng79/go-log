#


.PHONY : fmt
fmt :
	@echo "格式化代码"
	@gofmt -l -w ./



.PHONY : test
test :
	@echo "检查代码"
	@go vet ./...
	@echo "测试"
	@go test -race -coverprofile=coverage.txt -covermode=atomic ./...