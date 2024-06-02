# Create Go Backend

Creates a starter project folder for development of Go backends using a single command.

## Prerequisites

The following has to be installed prior to running the script:
- Go

## Installation

Run the following command to install this tool
```
go install github.com/David-The-Programmer/create-go-backend@latest
```

## Creating the project folder

Run the following command to create a new Go project folder:
```bash
cgb create
```

You would be prompted for:
1. The path of the new project folder

    Take note that you can use either:
    - A relative path (to your current directory). For example, `./project` or `../project`.
    - A absolute path. For example, `~/project`.

2. The go module path

    Please read the go documentation on specifying the [go module path](https://go.dev/doc/modules/managing-dependencies#naming_module).

3. The go version

    Only the version number needs to be entered, for e.g, `1.17` or `1.21.7`.

    All go release versions can be found [here](https://go.dev/doc/devel/release).


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

## Changing the Go version after creation

If the Go version has to be changed after the project folder has been created, take note to change:
1. The `.env` file. This file contains the `GO_VERSION` environment variable, which is used by `docker compose` to know which Go version to use when building the docker image. Only changing the value of the `GO_VERSION` variable is required, there is no need to export the variable manually, `docker compose` automatically interpolates the variable value from the `.env` file.
2. The `go.mod` file. Run `go mod edit -go=<new go version>`. This will correct the Go version in the `go.mod` file.

Subsequently, run `docker compose build` to force rebuilding of the docker image before running `docker compose up --watch`. 

This is because `docker compose up --watch` uses the cached image by default, hence the change in Go version would not be reflected.

## TODOS
- [ ] Able to execute command cgb create to actual create project folder
- [ ] Package CLI tool to be go install-able
- [ ] Versioning
- [ ] Update README sections on installation and creating project folder
- [ ] Automated tests to ensure script creates project folder properly (including `docker compose up --watch`?)
- [ ] Improving folder name and module path prompts to ensure user does not give invalid inputs (e.g, no input, input with only spaces, module name not matching folder name, etc)
    - [ ] Warn user of directory override when project folder path already exists
    - [ ] Re-prompt user if user does not want to override existing folder in project folder path
    - [ ] Account for relative path vs absolute paths for prompt of project path
- [ ] Add timeouts when downloading template folder content files
- [ ] Pipe error logs to file instead of just slog (allow both options)
- [ ] Inclusion of git hooks for testing and linting
- [ ] Compatibility of `Dockerfile` and `compose.yml` with `.devcontainer.json` (maybe a separate way to make sure vscode uses workspace specific configs, lang version and extensions?)
