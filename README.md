# Create Go Backend

Creates a starter project folder for development of Go backends using a single command.

## Prerequisites

The following has to be installed prior to running the script:
- Git
- Go
- Bash

## Creating the project folder

Run the following command to create a new Go project folder:
```bash
git clone git@github.com:David-The-Programmer/create-go-backend.git /tmp/create-go-backend/ && /tmp/create-go-backend/create.sh
```

## Using docker (optional)

Once the project is generated, the `Dockerfile` and `compose.yml` files can be deleted if usage of docker is not required.

Run the following command to start running the Go backend:
```bash
docker compose up --watch
```
If errors popup while building the Docker image, run the following command after fixing the error to force `docker compose` to rebuild the Docker image and ignore the cache.
```bash
docker compose build
```
Subsequently run `docker compose up --watch` to start the container.

The interesting thing is that `--watch` flag will cause `docker compose` to watch any changes made to the Go code and rebuild upon change.

This solves the issue of constantly manually rebuilding the code, which is a pain :(

## Using VSCode configuration files (optional)

The `.vscode` folder and the `.devcontainer.json` file can be deleted if VSCode would not be used as the editor of choice.

## TODOS
- [ ] Compatibility of `Dockerfile` and `compose.yml` with `.devcontainer.json` (maybe a separate way to make sure vscode uses workspace specific configs, lang version and extensions?)
- [ ] Running `go mod init` via the docker container instead of on the local system to ensure go version in `go.mod` file matches the go version in `Dockerfile`?
- [ ] Automated tests to ensure script is create project folder properly (including `docker compose up --watch`?)
- [ ] Improving folder name and module path prompts to ensure user does not give invalid inputs (e.g, no input, input with only spaces, module name not matching folder name, etc)


