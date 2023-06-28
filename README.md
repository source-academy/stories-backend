# Backend for Source Academy Stories

## Setup

### Installing languages and tools required

Install Golang and PostgresSQL. We recommend using a version manager like [asdf](https://asdf-vm.com/) to manage your installations. Use the versions of the tools listed in the `.tool-versions` file.

* (For `asdf` only) if not already installed, install the necessary plugins:

  ```bash
  asdf plugin add golang
  asdf plugin add postgres
  ```

* Install the required versions of the tools:

  ```bash
  asdf install golang 1.20.5
  asdf install postgres 14.8
  ```

To work with and debug the API, we also strongly recommend installing tools like [Postman](https://www.postman.com/downloads/).

### Setting up the database

**TODO:** Add instructions for setting up the database.

## Development

### Starting the server

```bash
make dev
```

### Running linter

Step 1: Install [golangci-lint](https://golangci-lint.run/usage/install/#local-installation) locally.

Step 2: Run the lint commmand:

```bash
make lint
```

## Building for production

```bash
make build
```
