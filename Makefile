
start:
	go run main.go

build:
	GOOS=linux GOARCH=amd64 go build -o dist/bootstrap main.go

unit-test:
	go test $(shell go list ./... | grep -v integrationtests)

integration-test:
	go test ./integrationtests 

infra-init:
	terraform -chdir=tf init
	terraform -chdir=tf workspace select dev || terraform -chdir=tf workspace new dev 
	terraform -chdir=tf state pull

infra-deploy:
	$(MAKE) build
	terraform -chdir=tf apply

infra-destroy:
	terraform -chdir=tf destroy

# disallow any parallelism (-j) for Make. This is necessary since some
# commands during the build process create temporary files that collide
# under parallel conditions.
.NOTPARALLEL: