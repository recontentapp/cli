# Recontent.app CLI

When pulling updates, make sure to include ones from submodules using:

```sh
git pull --recurse-submodules
```

## Build the client

Prerequisites: 
- [`swagger-cli`](https://github.com/APIDevTools/swagger-cli)
- [`oapi-codegen`](https://github.com/deepmap/oapi-codegen)

```sh
make build-client
```
