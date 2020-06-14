[![codecov](https://codecov.io/gh/garrettsparks/idgen/branch/master/graph/badge.svg)](https://codecov.io/gh/garrettsparks/idgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/garrettsparks/idgen)](https://goreportcard.com/report/github.com/garrettsparks/idgen)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=garrettsparks_idgen&metric=alert_status)](https://sonarcloud.io/dashboard?id=garrettsparks_idgen)

# idgen
--
    import "github.com/garrettsparks/idgen"

Package idgen uses crypt/rand to generate a random string of a given length.

## Usage

```go
const (
        // DefaultIDChars defines the default set of characters as alphanumeric plus uppercase.
        DefaultIDChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

        // DefaultIDLength defines the default ID length as 8.
        DefaultIDLength = 8
)
```

#### type IDBuilder

```go
type IDBuilder struct {
}
```

IDBuilder is a builder for an ID.

#### func  New

```go
func New() *IDBuilder
```
New creates a new IDBuilder with a specified length.

#### func (*IDBuilder) BuildID

```go
func (builder *IDBuilder) BuildID() string
```
BuildID builds an ID.

#### func (*IDBuilder) WithCharset

```go
func (builder *IDBuilder) WithCharset(chars string) *IDBuilder
```
WithCharset sets the string of available characters for the builder to use when
generating an ID. The default is alphanumeric plus uppercase.

#### func (*IDBuilder) WithLength

```go
func (builder *IDBuilder) WithLength(length int) *IDBuilder
```
WithLength sets the length of the randomly generated ID. The default length is
8.
