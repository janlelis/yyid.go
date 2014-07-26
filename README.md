# YYID

Generates random tokens that look like [type 4 UUIDs](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_.28random.29): `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`

In contrast to [RFC 4122](https://tools.ietf.org/rfc/rfc4122.txt), it uses all digits (128bit)

Source of randomness: [crypto/rand](http://golang.org/pkg/crypto/rand/)


## Install

Add dependency to your go file:

```go
import (
  "github.com/janlelis/yyid.go"
)
```

Then install with:

```
$ go install
```


## Usage

```go
yyid.New() // => "b4d254ec-198e-dbca-2ca3-bd3551180b50"
```

## Also Available

- YYID for [Ruby](https://github.com/janlelis/yyid.rb)
- YYID for [JavaScript](https://github.com/janlelis/yyid.js)
- YYID for [Elixir](https://github.com/janlelis/yyid.ex)

