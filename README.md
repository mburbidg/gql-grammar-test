# grmtest
Parse a GQL-program in a file, or parse all the GQL-programs in a directory.

# Build instructions

Install Go 1.22 or higher: https://go.dev/doc/install

From the root directory, execute the following command `go build`.

# Running grmtest

To run a single GQL-program through the parser use the -file parameter as in the following example.

```
./grmtest -file=/Users/tom/pgm.gql
```

To run all GQL-programs in a directory through the parser use the -dir parameter as in the following example.

```
./grmtest -dir=/Users/tom/gql-programs
```


