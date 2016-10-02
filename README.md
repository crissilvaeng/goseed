# Goddamned

[![Build Status](https://travis-ci.org/crissilvaeng/goddamned.svg?branch=master)](https://travis-ci.org/crissilvaeng/goddamned) [![Coverage Status](https://coveralls.io/repos/github/crissilvaeng/goddamned/badge.svg?branch=HEAD)](https://coveralls.io/github/crissilvaeng/goddamned?branch=HEAD) ![go-version](https://img.shields.io/badge/go%20version-1.7-orange.svg) [![GoDoc](https://godoc.org/github.com/crissilvaeng/goddamned?status.svg)](https://godoc.org/github.com/crissilvaeng/goddamned) ![license](https://img.shields.io/badge/license-MIT-blue.svg)

The objective of this project is to practice writing applications Go. Covering aspects such as testing, documentation, restful API and deploys in cloud environments.

## Use from the source

Assuming a recent version installed of Go and the variables `GOROOT` and `GOPATH` already defined, the set of commands below should recreate this project locally on a machine. Otherwise, [this installation guide](https://golang.org/doc/install) available in official of the [Golang](https://golang.org) site will help.

In addition, to manage possible dependencies this project use a vendor folder. Is possible restore dependencies using [govendor](https://github.com/kardianos/govendor) whose [installation instructions are available here](https://github.com/kardianos/govendor) or [Godep](https://github.com/tools/godep) whose [installations instructions are available here](https://github.com/tools/godep).

Using govendor the following instructions should work fine, otherwise check the dependency manager documentation.

**On Windows**

```bash
> go get -u github.com/crissilvaeng/goddamned
> cd %GOPATH%\github.com\crissilvaeng\goddamned
> govendor sync
> go get
> goddamned
```

**On Linux**

```bash
$ go get -u github.com/crissilvaeng/goddamned
$ cd $GOPATH/github.com/crissilvaeng/goddamned
$ govendor sync
$ go get
$ goddamned
```

## Change Log

To see the changes in this project, see the [change log file](CHANGELOG.md).

## Contribution

To contribute to this project, see the file [contribution tips](CONTRIBUTING.md) and [code of conduct](CONDUCT.md).

## License

The MIT License (MIT). To see the details of this license, see the [license file](LICENSE.md).

:octocat: