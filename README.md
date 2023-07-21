# Backend for Source Academy Stories

[![Coverage Status](https://coveralls.io/repos/github/source-academy/stories-backend/badge.svg?branch=main)](https://coveralls.io/github/source-academy/stories-backend?branch=main)

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

### Setting up git hooks

We use pre-commit and pre-push hooks to ensure that code is formatted and linted before committing and pushing. To set up the hooks, run the following commands:

```bash
make hooks
```

**Notes:**

* Hooks only work on macOS/Linux. Support for Windows is coming soon.
* You will need to install golangci-lint locally for git hooks to work. See [Running linter](#running-linter) for instructions.

### Setting up the database

#### Creating the database

Run the following command to create the database:

```bash
make db_create
```

#### Migrating the database

For this step onwards, make sure you have your environment variables set up correctly. See [Setting up environment variables](#setting-up-environment-variables) for more details.

Run the migration target to migrate the database:

```bash
make db_migrate
```

By default, the migration target will run all pending migrations (visible using `make db_status` command). To only run a certain number of migrations forward, use the `steps` argument:

```bash
make db_migrate steps=1
```

#### Rolling back migrations

To roll back the most recent migration, run the following command:

```bash
make db_rollback
```

To roll back a certain number of migrations, use the `steps` argument:

```bash
make db_rollback steps=1
```

### Setting up environment variables

Copy `.env.example` to `.env` and fill in/modify the required values as needed.

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

### Testing your code

By convention, test files are named `*_test.go` and are placed in the same directory as the code they are testing.

To run all tests:

```bash
make test
```

To run all tests and view test coverage:

```bash
make coverage
```

## Building for production

```bash
make build
```
