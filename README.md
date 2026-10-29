# GO programs

Small Go practice programs, one folder each.

## Programs

| Program | Description |
| :--- | :--- |
| `hello-world` | Prints `Hello, World!` |
| `calculator` | Basic calculator (`+`, `-`, `*`, `/`) |
| `age-checker` | Tells whether the age is under or over 18 |
| `variables` | Explains variables and types in Go |
| `operators` | Arithmetic operators demo |
| `fibonacci` | Prints the Fibonacci sequence up to `N` |
| `read-files` | Reads a text file and prints it |
| `date-and-hour` | Prints the current date and time |

## Usage

1. Clone the repository.

   ```sh
   git clone https://github.com/Ars-byte/GO-programs
   cd GO-programs
   ```

2. Run a program by its folder name.

   ```sh
   go run ./hello-world
   ```

   Replace `hello-world` with another folder from the table above.

3. Optional: build a binary.

   ```sh
   go build -o hello-world ./hello-world
   ./hello-world
   ```

`read-files` reads a sample file, so run it from its own folder.

```sh
cd read-files
go run .
```

Prebuilt binaries are also available on the
[latest release](https://github.com/Ars-byte/GO-programs/releases/latest).
