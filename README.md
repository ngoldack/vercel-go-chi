# vercel-go-chi

Demo application using vercel serverless functions to create a simple API with go using chi router.

## About

This template uses common go api practices for developing apis for deployment on vercel.
When deployed on vercel, every time the serverless function is invoked, a new router is created and the current request is matched against it.

### Who this is for

If you have a small api with not a lot of traffic, this template and serverless functions in general are right for you.

### Who this is not for

If you have a lot of traffic or have a lot of external dependencies (Databases, external APIs, ...) then you might use a real server.

For example if you have are connecting to a database,
vvery time the serverless function is invoked, a new connection to the database will be created and destroyed afterwards. This can create a bottleneck in a high traffic environment.

## Technologies used

- [Go](https://go.dev/) - Programming language
- [Chi](https://go-chi.io/) - Router
- [Vercel](https://vercel.com/) - Serverless functions
- [Magefile](https://magefile.dev/) - Task executor
- [Dagger](https://dagger.io/) - CI
- [Github Actions](https://github.com/features/actions) - CI executor

## Using this template

To use this template, you can click on the "Use this template" button or clone it using git:

```bash
git clone https://github.com/ngoldack/vercel-go-chi.git
```

## Pre-requisites

To run this application locally, you need to have [Go](https://golang.org/) and [Magefile](https://magefile.dev) installed.

If you want to run the CI locally, you also need [Docker](https://www.docker.com/) and [Dagger](https://dagger.io) installed.

## Tasks

### Dev server

```bash
mage -v dev
```

### Test

```bash
mage -v test
```

### CI (build, lint, test)

Requires `docker` and `dagger` installed.

```bash
dagger run go run ci/main.go
```
