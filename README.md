# aliens-invasion-go

### Aliens Invasion GO implementation

Project structure:

cmd/aliensapp - app running the aliens invasion simulation  
pkg/alieninvasion - package that contains all the aliens invasion logic  
tests/data - test data files (consistent/inconsistent)  
docs/README.md - steps to generate documenation  

### Prerequisites

You'll need `go` [installed](https://golang.org/doc/install) on your machine

### Install using `go get` and Run

```
go get github.com/artem-brazhnikov/aliens-invasion-go/cmd/aliensapp
aliensapp
```

### Get Source Code and Run

```
mkdir -p $GOPATH/src/github.com/artem-brazhnikov
cd $GOPATH/src/github.com/artem-brazhnikov
git clone https://github.com/artem-brazhnikov/aliens-invasion-go.git
cd aliens-invasion-go
go run ./cmd/aliensapp
```

### Documentation

Documentation can be found in the [docs](docs/README.md).

### Running the tests

`go test ./...`

### Authors

* **Artem B.** - [artem](https://github.com/artem-brazhnikov)

### License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details
