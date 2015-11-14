# drone-go-mysql

This is a simple command line utility for managing a todo list, backed by a mysql database. This project demonstrates how to build, test and setup continuous integration using [Drone CI](https://github.com/drone/drone).

## Build and Test

Compiling the code:

```
go get
go build -o todo
```

## Usage

Adding todo items:

```
$ todo add "download drone"
$ todo add "setup continuous delivery"
$ todo add "profit"
```

Listing todo items:

```
$ todo ls

1 download drone
2 setup continuous delivery
3 profit
```

Removing todo items:

```
$ todo rm 1
```