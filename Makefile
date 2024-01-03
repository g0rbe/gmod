.PHONE: test
test:
	go test ./...

.PHONE: benchmark
benchmark:
	go test ./... -bench=. -benchmem