all:
	go install golang.org/x/tools/cmd/goimports
	goimports -d -w $$(find . -type f -name '*.go' -not -path '*/mock.go')
	go test ./...
	go mod tidy

mockgen:
	go install github.com/golang/mock/mockgen
	mockgen \
		-destination internal/test/mock/compute/mock.go \
		-package mockcompute github.com/nokamoto/demo20-apis/cloud/compute/v1alpha ComputeClient
