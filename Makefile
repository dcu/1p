test:
	go test github.com/dcu/1p/keychain -cover -coverprofile=keychain_coverage.out -short
	go tool cover -func=keychain_coverage.out
	go test github.com/dcu/1p/cli -cover -coverprofile=cli_coverage.out -short
	go tool cover -func=cli_coverage.out
	@rm *.out

