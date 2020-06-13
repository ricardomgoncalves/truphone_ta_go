# Truphone Assessment

## Introduction



### Contents
* Building
* Running
* Testing
* Postman

## Building

Running the service requires to build the docker image locally. 
To do this just run the command:

```shell script
make build_app
```

There you go, you now have truphone_go:latest docker image.

## Running

To run the service, as well its database just run:

```shell script
make run
```

This will run one instance of the postgres and one of the truphone_go.

## Testing

### Unit Testing

To run the unit tests and see the coverage (76.2% of files and 85.3% of statements with 485 tests) just run:

```shell script
make coverage
```

This will run the unit tests and will open in the browser the coverage report.

## Postman

A postman collection included in the file: `Truphone.postman_collection.json`

