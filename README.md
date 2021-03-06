# semver [![GoDoc](https://godoc.org/github.com/typerandom/semver?status.png)](http://godoc.org/github.com/typerandom/semver) [![Build Status](https://travis-ci.org/typerandom/semver.svg?branch=master)](https://travis-ci.org/typerandom/semver)

Semantic versioning (http://semver.org/) for Go. Full support of the 2.0 spec.

## Install

Just use go get.

    go get gopkg.in/typerandom/semver.v1
    
And then just import the package into your own code.

    import (
        "gopkg.in/typerandom/semver.v1"
    )

## Getting started

```go
package main

import (
	"fmt"
	"gopkg.in/typerandom/semver.v1"
)

func main() {
	version := semver.New("1.0.0")
	fmt.Printf("Version is: v%d.%d.%d", version.Major(), version.Minor(), version.Patch())
}
```

## Documentation

For full documentation [see GoDoc](https://godoc.org/github.com/typerandom/semver).

## Licensing

Semver is licensed under the MIT license. See LICENSE for the full license text.
