# vercel-go-chi

Demo application using vercel serverless functions to create a simple API with go using chi router.

## Technologies used

- [Go](https://go.dev/) - Programming language
- [Chi](https://go-chi.io/) - Router
- [Vercel](https://vercel.com/) - Serverless functions
- [Dagger](https://dagger.io/) - CI
- [Github Actions](https://github.com/features/actions) - CI Executor

## Using this template

To use this template, you can click on the "Use this template" button or clone it using git:

```bash
git clone https://github.com/ngoldack/vercel-go-chi.git
```

## Pre-requisites

To run this application locally, you need to have [Go](https://golang.org/) installed.

If you want to run the CI, you also need [Docker](https://www.docker.com/) and [Dagger](https://dagger.io) installed.

## Running locally

### CI (build, lint, test)

```bash
dagger run go run ci/main.go
```

### Dev server

```bash
go run ./cmd/server/main.go
```
