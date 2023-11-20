
include .env

## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'


# ==================================================================================== # 
# DEVELOPMENT
# ==================================================================================== #

## run: run in development mode
.PHONY: run
run:
	@go run ./cmd
