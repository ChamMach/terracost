# Contributing Guidelines

Cycloid Team is glad to see you contributing to this project ! In this document, we will provide you some guidelines in order to help get your contribution accepted.

## Reporting an issue

### Issues

When you find a bug in Terracost, it should be reported using [Github issues](https://github.com/cycloidio/terracost/issues). Please provide key information like your Operating System (OS), Go version and finally the version of the library that you're using.

## Submit a contribution

### Setup your git repository

If you want to contribute to an existing issue, you can start by _forking_ this repository, then clone your fork on your machine.

```shell
$ git clone https://github.com/<your-username>/terracost.git
$ cd terracost
```

In order to stay updated with the upstream, it's highly recommended to add `cycloidio/terracost` as a remote upstream.

```shell
$ git remote add upstream https://github.com/cycloidio/terracost.git
```

Do not forget to frequently update your fork with the upstream.

```shell
$ git fetch upstream --prune
$ git rebase upstream/master
```

### Test your submission

First-time setup of the development environment requires running the database migrations:

```shell
$ make db-migrate
```

Run all the tests:

```shell
$ make test
```

## Adding a new resource

### General

* On the `{provider}/terraform` resources info

### AWS

To add more AWS resources first read the [documentation](docs/aws.md) we have about it.

### Google

To add more Google resources first read the [documentation](docs/google.md) we have about it.

### AzureRM

To add more AzureRm resources first read the [documentation](docs/azurerm.md) we have about it.

## Adding a new provider/backend

**Please be aware that, at the moment, Cycloid only supports MySQL as a backend.** Based on this, please refrain from making contributions that add a new backend or cloud provider as we cannot guarantee they'd be merged and/or supported. To make improvements in this area, please instead open an appropriate issue so that we can discuss it and provide any necessary guidance.
