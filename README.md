# Test service in go

This service serves no other purpose other than being
packaged into a docker image.


## Endpoints

1. `/string`: displays a greeting to the server admin as a regular string.
2. `/json`: same as `/string` but in JSON format.

## Building the source code (compilation)

1. Service is written in Go 1.18 – make sure it is installed.
2. Run `go build -o /path/to/output/binary /path/to/source/code/directory`


## Running

To run the service just run the output binary that you specified
during the compilation.

## Configuration

The service reads its configuration from the following environment
variables:

1. `SERVICE_PORT`: required. Specifies the port the service should listen on.
2. `USER_NAME`: optional. Specifies server admin's name. `Roman` by default.