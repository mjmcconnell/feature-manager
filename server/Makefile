help:		## list available cmds.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

run: 		## install local packages and start the dev server with live reload
	go mod tidy
	air --build.cmd "go build -o bin/api cmd/feature-manager/main.go" --build.bin "./bin/api"
