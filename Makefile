.PHONY: gateway
gateway:
	go run ./cmd/gateway

.PHONY: msq
msq:
	go run ./cmd/msq
