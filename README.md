# GO programs

A collection of small Go practice programs, each in its own folder.

## Programs

| Folder | Description |
| :--- | :--- |
| `hello-world` | Prints `Hello, World!` |
| `calculator` | Basic calculator (`+`, `-`, `*`, `/`) |
| `age-checker` | Tells whether the age is under or over 18 |
| `variables` | Variables and types demo |
| `operators` | Arithmetic operators demo |
| `fibonacci` | Prints the Fibonacci sequence up to `N` |

## Build and run

Each folder is a standalone `main` package.

```sh
go build -o /tmp/hello-world ./hello-world
/tmp/hello-world
```

Or with `go run`:

```sh
go run ./hello-world
```

## Releases

On every push, a GitHub Actions workflow builds all programs and publishes the
binaries, tarballs and a `SHA256SUMS` file to the
[latest release](https://github.com/Ars-byte/GO-programs/releases/latest).
