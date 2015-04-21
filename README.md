# 1p A command line interface for 1Password

1p is a CLI for 1Password. For now it is read-only.

## Installation
The application is written in `go`. To install you need to use `go get` as follows:

    go get github.com/dcu/1p

and the command `1p` will be available in `$GOPATH/bin` so make sure it is in your `$PATH`

## Usage

The CLI support the following sub commands:

### Getting help

The command `1p help` will display the available commands and usage.

### Quering an item

To get information about a item type the following command:

    1p query <pattern>

For example

    1p query gmail

If several items are found `1p` will ask you to select one.
The following aliases are supported: `q`


### Copying password to clipboard

To copy a password to the clipboard use the following syntax:

    1p copy <pattern>

For example

    1p copy gmail

If several items are found `1p` will ask you to select one
The following aliases are supported: `c`, `cp`

## Running the tests

To run the test and coverage just type following command:

    make test

It'll print the test results and coverage.

