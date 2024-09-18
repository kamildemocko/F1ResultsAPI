BUILD_BINARY=.\bin\f1-results-api.exe
GO_BINARY=f1-results-api.exe

test:
	@echo start tests
	@go test .\cmd\api\ -v -count=1
	@echo - done

build:
	@echo start build
	@go build -o .\bin\f1-results-api.exe .\cmd\api
	@echo - done

run: build
	@echo start run
	@start /B ${BUILD_BINARY} &
	@echo - done

stop:
	@echo killing proces
	@taskkill /IM ${GO_BINARY} /F
	@echo - done

restart: stop run

swag:
	@echo creating swagger
	@swag init -g .\cmd\api\main.go
	@echo - done

swagr: swag restart
