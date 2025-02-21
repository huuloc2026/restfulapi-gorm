## E-commerce REST API in Go

### Installation

There are a few tools that you need to install to run the project.
So make sure you have the following tools installed on your machine.

## Running the project

Firstly make sure you have a MySQL database running on your machine or just swap for any storage you like under `/db`.

Then create a database with the name you want _(`ecom` is the default)_ and run the migrations.

```bash
make migrate-up
```

After that, you can run the project with the following command:

```bash
make run
```

## Running the tests

To run the tests, you can use the following command:

```bash
make test
```
