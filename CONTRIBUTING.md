# Contributing

## Code Style and Code Conventions

The code written in [Golang](https://golang.org/) in this repository tries to follow the rules set by the community for this topic.

A summary of the introductory recommendations can be found in [CodeReviewComments wiki topic](https://github.com/golang/go/wiki/CodeReviewComments), a more complete version can be found in [Effective Go article](https://golang.org/doc/effective_go.html). Additional guidance on how to write good code can be found in the article [How to Write Go Code](https://golang.org/doc/code.html).

## Dependency Management

In addition, to manage possible dependencies this project uses [govendor](https://github.com/kardianos/govendor). Modifications on the dependencies must be submitted with the pull request on file `vendor/vendor.json`. Instructions as govendor works can be found [here](https://github.com/kardianos/govendor).

## Commit message pattern

- Use imperatively ("Add feature" no "Adding feature" or "Added feature").
- First line should have a maximum of 72 characters.
- Consider describe in detail in the body commits.
- Consider using an emoji at the start of the commit message.

### Emoji table for commits

Emoji | Code | Commit Type
------------ | ------------- | -------------
:tada: | `:tada:` | Initial commit.
:art: | `:art:` | Improve structure or code format.
:racehorse: | `:racehorse:` | Improves performance.
:memo: | `:memo:` | Write documentation.
:bug: | `:bug:` | Fix a bug.
:fire: | `:fire:` | Remove code or file.
:green_heart: | `:green_heart:` | Fixes a building in CI.
:white_check_mark: | `:white_check_mark:` | Add tests.
:lock: | `:lock:` | Improves security.
:arrow_up: | `:arrow_up:` | Dependencies upgrade.
:arrow_down: | `:arrow_down:` | Dependencies downgrade.
:poop: | `:poop:` | Deprecated.
:construction: | `:construction:` | Work in progress.
:rocket: | `:rocket:` | New feature.
:see_no_evil: | `:see_no_evil:` | Ugly hack.
:package: | `:package:` | New release.
:eyeglasses: | `:eyeglasses:` | Modified GUI.
:wrench: | `:wrench:` | Modified configuration files.
:bulb: | `:bulb:` | When removing linter warnings.

### Example
```bash
git commit -m ":memo: Add contribution instructions.
>
> It was created a file (CONTRIBUTING.md) with instructions
> on how to make a good commit."
```

## Testing

This repository emphasizes the test coverage. Covering all the features and components with unit tests and performing integration tests whenever possible.

Guidelines on how to produce good unit tests can be found [here](http://geosoft.no/development/unittesting.html) and [here](http://artofunittesting.com/unit-testing-review-guidelines/). Resources and tools about tests Go can be found [here](https://github.com/avelino/awesome-go#testing).

There is an excellent post on the Go Lang blog called [The cover story](https://blog.golang.org/cover) that illustrates how to get the current project test coverage and other relevant data.

## Contribution workflow

This repository adopts workflow known as [GitHub Flow](https://guides.github.com/introduction/flow/) created by Github. Details on the operation of branchs model can be found [here](https://guides.github.com/introduction/flow/).

![gitflow-diagram](images/github-workflow.png)