test-controllers:
	go test -v ./app/controllers -cover -coverprofile=coverage.out && go tool cover -func coverage.out