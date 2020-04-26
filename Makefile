
all: build

build:
	docker build --progress=plain -t appleboy/docker-demo -f Dockerfile .

buildkit:
	docker build --progress=plain -t appleboy/docker-buildkit -f Dockerfile.buildkit .
