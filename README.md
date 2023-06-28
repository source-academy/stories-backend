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

## Starting the server

### Development

```bash
make dev
```

### Building for production

```bash
make build
```

## Testing your code

By convention, test files are named `*_test.go` and are placed in the same directory as the code they are testing.

To run all tests:

```bash
make test
```

To run all tests and view test coverage:

```bash
make coverage
```
