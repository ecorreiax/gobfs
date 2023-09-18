# gobfs

[![Go Report Card](https://goreportcard.com/badge/github.com/ecorreiax/gobfs)](https://goreportcard.com/report/github.com/ecorreiax/gobfs)
![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/ecorreiax/gobfs)
![GitHub](https://img.shields.io/github/license/ecorreiax/gobfs)
![GitHub issues](https://img.shields.io/github/issues/ecorreiax/gobfs)

Bloom Filter Structure implementation in Go

## What's a Bloom Filter

> A Bloom filter is a space-efficient probabilistic data structure, conceived by Burton Howard Bloom in 1970, that is used to test whether an element is a member of a set. False positive matches are possible, but false negatives are not – in other words, a query returns either "possibly in set" or "definitely not in set".
>
> <cite>[Source from Wikipedia][1]</cite>

[1]: https://en.wikipedia.org/wiki/Bloom_filter

## Install

With a [correctly configured](https://go.dev/doc/install#testing) Go toolchain:

```
go get -u github.com/ecorreiax/gobfs
```

## Examples

Create a hash for a specif string and then add this index into the bitset array

```go
func myFunction() {
  h := sha1.New()
  a := [3]string{"apple", "grape", "mango"}
  for _, s := range a {
    idx, _ := hash.CreateHash(h, s)
    hash.AddHash(idx)
  }
}
```

Check if a specific string is definitely not in the bitset array

```go
func checkString() {
  s := "coconut"
  found := gobfs.Check(u)
  if found {
   // possibly in bitset
  }
  // definitely not in bitset
}
```

_Possibly in bitset means that is not possible to be sure, since differents words after hashed can result in the same index._

## Core team

- [@ecorreiax](https://github.com/ecorreiax)


## License

This project is under [MIT License](https://github.com/ecorreiax/gobfs/blob/main/LICENSE).
