.PHONY: build
build:
	go build -ldflags="-s -w"
	
.PHONY: run
run: build
	rm gorm.db
	./game-api
