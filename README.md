# CmdBox Default `version` Module

[![GoDoc](https://godoc.org/cmdbox-version?status.svg)](https://godoc.org/cmdbox-version)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report
Card](https://goreportcard.com/badge/cmdbox-version)](https://goreportcard.com/report/cmdbox-version)
[![Coverage](https://gocover.io/_badge/github.com/rwxrob/cmdbox-version)](https://gocover.io/github.com/rwxrob/cmdbox-version)

This module provides the default `version` command to be used in [CmdBox]
composite (monolith) programs. Hence, it has no stand-alone variation.
CmdBox developers are free to use whatever `version` module they choose or
to omit it entirely when minimal executable sizes are desired.

[CmdBox]: <https://github.com/rwxrob/cmdbox>

## Install 

Use `import` with a blank identifier to be composed:

```go
import (
  "github.com/rwxrob/cmdbox"
  _ "github.com/rwxrob/cmdbox-version"
)
```

## Embedded Documentation

See the [`cmd.go`](cmd.go) file itself for additional embedded
documentation about this command.

## Usage

```
foo version
```

## Legal

Copyright (c) 2021 Robert S. Muhlestein
Released under the [Apache 2.0](LICENSE)

Contributors and project participants implicitly accept the 
[Developer Certificate of Authenticity (DCO)](DCO).

