help:		## list available cmds.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

docker.build: 	## build local server dockerfile
	docker build -t server:local -f docker/server.Dockerfile .

docker.run: 	## run local server dockerfile
	docker run -p 8080:8080 server:local
