
include .env

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'


# ==================================================================================== # 
# DEVELOPMENT
# ==================================================================================== #

## dev: run in development mode
.PHONY: dev
dev:
	@go run ./cmd
