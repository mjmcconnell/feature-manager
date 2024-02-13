docker.build:
	docker build -t server:local -f docker/server.Dockerfile .

docker.run:
	docker run -p 8080:8080 server:local
