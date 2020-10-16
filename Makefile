include .env

blah: dev

start:
	@go run .

dev: dependencies
	@air -c ${AIR_FILE}

dependencies:
	@go mod download
	@go get ./...

commit:
	@git add .
	@git commit -am ${COMMIT_MESSAGE}

deploy: commit
	@git push heroku develop:master

push: commit
	@git push
