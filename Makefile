
all:
	@echo "TODO: Docker up, down, build, run"
	@$(MAKE) run

run:
	@go run cmd/main.go

test:
	@go test -count=1 -p 1 -cover ./...
