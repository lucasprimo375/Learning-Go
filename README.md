# Learning-Go
Repository to Learn Go from Udemy course

# How to run a Go file
```bash
go run <filename>
```

# How to init module
```bash
go mod init <modulename>
```

## To build a module
This will generate an executable

```bash
go build
```

## To install module
This will generate an executable in the go installation folder

```bash
go install
```

## To get external module

```bash
go get <link-to-external-module>
```

## To remove all Unused dependencies

```bash
go mod tidy
```

# Testing

## Test file name
All test files must end with _test.go, i.e. mytest_test.go

## Test function name
All test functions must start with Test, i.e. TestMyFunction

## Running a test
Run the following command in the same folder as the file test:
```bash
go test
```