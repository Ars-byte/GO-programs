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

## Run with Go

1. Clone the repository.

   ```sh
   git clone https://github.com/Ars-byte/GO-programs
   cd GO-programs
   ```

2. Run any program by its folder name.

   ```sh
   go run ./hello-world
   ```

   Replace `hello-world` with another folder from the table above.

## Run a release binary

1. Download a program (example: `hello-world`).

   ```sh
   curl -fsSL https://github.com/Ars-byte/GO-programs/releases/latest/download/hello-world -o hello-world
   ```

2. Make it executable.

   ```sh
   chmod +x hello-world
   ```

3. Run it.

   ```sh
   ./hello-world
   ```

All binaries are attached to the
[latest release](https://github.com/Ars-byte/GO-programs/releases/latest), with
checksums in `SHA256SUMS`.
