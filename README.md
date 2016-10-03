<h1 align="center">
  <br>
  <a href="https://github.com/crissilvaeng/goseed"><img src="https://golang.org/doc/gopher/appenginegophercolor.jpg" alt="Go Seed" width="200"></a>
  <br>
  Go Seed
  <br>
  <br>
</h1>

<h4 align="center">A seed project</h4>
<h4 align="center">For Restfull Web APIs</h4>
<h5 align="center">Built with Golang and Gin</h5>

<p align="center">
  <a href="https://travis-ci.org/crissilvaeng/goseed"><img src="https://travis-ci.org/crissilvaeng/goseed.svg?branch=master" alt="Build Status"></a>
  <a href="https://coveralls.io/github/crissilvaeng/goseed?branch=HEAD"><img src="https://coveralls.io/repos/github/crissilvaeng/goseed/badge.svg?branch=HEAD" alt="Coverage Status"></a>
  <a href="https://goreportcard.com/report/github.com/crissilvaeng/goseed"><img src="https://goreportcard.com/badge/github.com/crissilvaeng/goseed" alt="Go Report Card"></a>
  <a href="https://github.com/crissilvaeng/goseed/releases"><img src="https://img.shields.io/github/release/crissilvaeng/goseed.svg" alt="Latest Release"></a>
  <a href="https://godoc.org/github.com/crissilvaeng/goseed"><img src="https://godoc.org/github.com/crissilvaeng/goseed?status.svg" alt="GoDoc"></a>
  <a href="https://golang.org/doc/devel/release.html"><img src="https://img.shields.io/badge/go%20version-1.7-orange.svg" alt="Go Version"></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License"></a>
</p>

<p align="center">
  <a href="https://golang.org/doc/gopher/">Gopher Images</a> by <a href="http://www.reneefrench.com/">Renée French</a> is licensed under <a href="https://creativecommons.org/licenses/by/3.0/">Creative Commons Attributions</a>
</p>

## Proposal

This repository is a seed for API projects developed in Go. It contains the dependencies, basic structure, test scripts and coverage. It is possible to easily integrate it with other services like [Travis Ci](https://travis-ci.org/), [Heroku](https://www.heroku.com/) and [Coveralls](https://coveralls.io/).

## Use from the source

Assuming a recent version installed of Go and the variables `GOROOT` and `GOPATH` already defined, the set of commands below should recreate this project locally on a machine. Otherwise, [this installation guide](https://golang.org/doc/install) available in official of the [Golang](https://golang.org) site will help.

In addition, to manage possible dependencies this project use a vendor folder. Is possible restore dependencies using [govendor](https://github.com/kardianos/govendor) whose [installation instructions are available here](https://github.com/kardianos/govendor) or [Godep](https://github.com/tools/godep) whose [installations instructions are available here](https://github.com/tools/godep).

Using govendor the following instructions should work fine, otherwise check the dependency manager documentation.

**On Windows**

```bash
> go get -u github.com/crissilvaeng/goseed
> cd %GOPATH%\github.com\crissilvaeng\goseed
> govendor sync
> go get
> goseed
```

**On Linux**

```bash
$ go get -u github.com/crissilvaeng/goseed
$ cd $GOPATH/github.com/crissilvaeng/goseed
$ govendor sync
$ go get
$ goseed
```

To test, run this command `PORT=5000 go run main.go`.

Then check the results with  `curl -i http://localhost:5000/v1/hello`. If all went well, the following results should be obtained:

```bash
HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Sun, 02 Oct 2016 07:46:36 GMT
Content-Length: 69

{"message":"The answer to life, the universe and everything is 42."}
 ```

## CI and Deploy

This solution is designed to use the following services: [Travis CI](https://travis-ci.org/), [Coveralls](https://coveralls.io/) and [Heroku](https://www.heroku.com/).

### [Coveralls](https://coveralls.io/)

Connect this repository to a account in Coveralls](https://coveralls.io/) service. Save the `repo_token`. Replace the URL of the README badge.

### [Travis CI](https://travis-ci.org/)

Connect this repository to a account in [Travis CI](https://travis-ci.org/) service. Set two environment variables. One called `COVERALLS_TOKEN` whit the `repo_token` from [Coveralls](https://coveralls.io/). The second called  `CI_SERVICE` with the name of the service.

### [Heroku](https://www.heroku.com/)

Create a new app. Next, in deploy option connect with Github and configure to deploy this repository after build on [Travis CI](https://travis-ci.org/).

## Change Log

To see the changes in this project, see the [change log file](CHANGELOG.md).

## Contribution

To contribute to this project, see the file [contribution tips](CONTRIBUTING.md) and [code of conduct](CONDUCT.md).

## License

The MIT License (MIT). To see the details of this license, see the [license file](LICENSE.md).

:octocat: