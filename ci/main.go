package main

import (
	"context"
	"log/slog"
	"os"

	"dagger.io/dagger"
)

const (
	// Path to the build directory
	BuildPath = "build/"
	SrcPath   = "./"
)

type CI struct {
	client *dagger.Client
}

func main() {
	ctx := context.Background()
	ci, err := NewCI(ctx)
	if err != nil {
		slog.Error("an error occurred", "error", err)
		os.Exit(1)
	}
	defer ci.client.Close()

	err = ci.build(ctx)
	if err != nil {
		slog.Error("an error occurred", "error", err)
		os.Exit(1)
	}

	err = ci.lint(ctx)
	if err != nil {
		slog.Error("an error occurred", "error", err)
		os.Exit(1)
	}

	err = ci.test(ctx)
	if err != nil {
		slog.Error("an error occurred", "error", err)
		os.Exit(1)
	}

	slog.Info("CI passed")
}

func NewCI(ctx context.Context) (*CI, error) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		return nil, err
	}

	return &CI{
		client: client,
	}, nil
}

func (ci CI) build(ctx context.Context) error {
	client := ci.client.Pipeline("build")

	src := client.Host().Directory(".")
	golang := client.Container().From("golang:latest")
	golang = golang.WithDirectory("/src", src).WithWorkdir("/src")
	golang = golang.WithExec([]string{"go", "build", "-o", BuildPath, "./cmd/..."})
	output := golang.Directory(BuildPath)
	_, err := output.Export(ctx, BuildPath)
	if err != nil {
		return err
	}

	return nil
}

func (ci CI) lint(ctx context.Context) error {
	client := ci.client.Pipeline("lint")

	src := client.Host().Directory(SrcPath)
	golangci := client.Container().From("golangci/golangci-lint:latest")
	golangci = golangci.WithDirectory("/src", src).WithWorkdir("/src")
	golangci = golangci.WithExec([]string{"golangci-lint", "run", "-v", "./..."})
	out, err := golangci.Stdout(ctx)

	if err != nil {
		slog.Error("lint error", "out", out, "error", err)
		return err
	}

	return nil
}

func (ci CI) test(ctx context.Context) error {
	client := ci.client.Pipeline("test")

	src := client.Host().Directory(SrcPath)
	golang := client.Container().From("golang:latest")
	golang = golang.WithDirectory("/src", src).WithWorkdir("/src")
	golang = golang.WithExec([]string{"go", "test", "./..."})
	out, err := golang.Stdout(ctx)
	if err != nil {
		slog.Error("test error", "out", out, "error", err)
		return err
	}

	out, err = golang.Stderr(ctx)
	if err != nil {
		slog.Error("test error", "out", out, "error", err)
		return err
	}

	return nil
}
