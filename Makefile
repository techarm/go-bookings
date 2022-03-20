# ANSI color
RED=\033[31m
GREEN=\033[32m
RESET=\033[0m

COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''

run:
	go run cmd/web/main.go cmd/web/middleware.go cmd/web/routers.go

test:
	go test -v ./... | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

coverage:
	go test ./... -coverprofile=coverage.out && go tool cover -html=coverage.out && rm -f coverage.out