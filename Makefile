build:
	@ echo "Building application... "
	@ go build -trimpath -o ./bin/engine ./app/
	@ echo "Done!"

build-migrate:
	@ echo "Building migrate... "
	@ go build -trimpath -o ./bin/migrate ./command/migrate.db.go
	@ echo "Done!"

build-image:
	@ echo "Bualding docker amage..."
	@ docker build \
		--file ./infra/docker/prod.Dockerfile \
		--tag xeff09/decauth-be \
		.
	@ echo "Done!"

# container env (default)

up:
	@ docker compose up

follow:
	@ docker logs decauth-app-1 --follow

migrate:
	@ docker exec -it decauth-app-1 go run ./command/migrate.db.go 

# version control

.PHONY: release

TYPE ?= patch

release:
	@ echo "Attaching with Version=$(TYPE)..."
	@ echo "$(TYPE)" > VERSION_TYPE 
