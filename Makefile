BUILD_BINARY=.\bin\f1-results-api.exe
GO_BINARY=f1-results-api.exe
DOCKER_NAME=f1-results-api
DOCKER_RUN_NAME=f1-results-api-run

test:
	@echo start tests
	@go test .\cmd\api\ -v -count=1
	@echo - done

build:
	@echo start build
	@go build -o .\bin\f1-results-api.exe .\cmd\api
	@echo - done

start: build
	@echo start run
	@start /B ${BUILD_BINARY} &
	@echo - done

stop:
	@echo killing proces
	@taskkill /IM ${GO_BINARY} /F
	@echo - done

restart: stop start

swag:
	@echo creating swagger
	@swag init -g .\cmd\api\main.go
	@echo - done

swagr: swag restart

dbuild:
	@echo building from Dockerfile
	@docker build -t ${DOCKER_NAME} .
	@echo - done

dstart:
	@echo running docker
	@docker run --rm -p 8080:80 -d --name ${DOCKER_RUN_NAME} ${DOCKER_NAME}
	@echo - done

dstop:
	@echo stopping docker
	@docker stop ${shell docker ps -q --filter "name=$(DOCKER_RUN_NAME)"}
	@echo - done
