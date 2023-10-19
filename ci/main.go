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

	if err = ci.build(ctx); err != nil {
		slog.Error("an error occurred", "error", err)
		os.Exit(1)
	}

	if err = ci.lint(ctx); err != nil {
		slog.Error("an error occurred", "error", err)
		os.Exit(1)
	}

	if err = ci.test(ctx); err != nil {
		slog.Error("an error occurred", "error", err)
		os.Exit(1)
	}

	if v, ok := os.LookupEnv("CI"); ok && (v == "true" || v == "1") {
		slog.Info("CI detected, uploading coverage")
		if err = ci.codecov(ctx); err != nil {
			slog.Error("an error occurred", "error", err)
			os.Exit(1)
		}

		slog.Info("GitHub Actions detected, starting deployment")
		if err = ci.deploy(ctx); err != nil {
			slog.Error("an error occurred", "error", err)
			os.Exit(1)
		}
	}

	slog.Info("CI finished successfully")
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
	golangci = golangci.WithExec([]string{"golangci-lint", "run", "./..."})
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
	golang = golang.WithExec([]string{"go", "test", "-coverprofile", BuildPath + "cover.out", "./..."})

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

	// copy coverage file to host
	output := golang.File(BuildPath + "cover.out")
	_, err = output.Export(ctx, BuildPath+"cover.out")
	if err != nil {
		return err
	}

	return nil
}

func (ci CI) codecov(ctx context.Context) error {
	client := ci.client.Pipeline("upload-coverage")

	codecov := client.Container().From("alpine:latest")
	path := client.Host().Directory(BuildPath)
	codecov = codecov.WithDirectory(BuildPath, path).WithWorkdir(BuildPath)

	// install curl
	codecov = codecov.WithExec([]string{"apk", "add", "--no-cache", "curl"})

	// Install codecov
	// curl -Os https://uploader.codecov.io/latest/alpine/codecov chmod +x codecov ./codecov
	codecov = codecov.WithExec([]string{
		"curl", "-Os", "https://uploader.codecov.io/latest/alpine/codecov",
	})
	codecov = codecov.WithExec([]string{"chmod", "+x", "codecov"})
	codecov = codecov.WithExec([]string{"./codecov", "-t", os.Getenv("CODECOV_TOKEN")})

	out, err := codecov.Stdout(ctx)
	if err != nil {
		slog.Error("codecov error", "out", out, "error", err)
		return err
	}

	out, err = codecov.Stderr(ctx)
	if err != nil {
		slog.Error("codecov error", "out", out, "error", err)
		return err
	}

	return nil
}

func (ci CI) deploy(ctx context.Context) error {
	_ = ci.client.Pipeline("deploy")

	slog.Info("Vercel auto-deployment configured. Skipping deployment.")

	return nil
}
