# CmdBox Default Version Module (`cmdbox-version`)

[![GoDoc](https://godoc.org/cmdbox-version?status.svg)](https://godoc.org/cmdbox-version)
[![License](https://img.shields.io/badge/license-Apache2-brightgreen.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/cmdbox-version)](https://goreportcard.com/report/cmdbox-version)

This module provides the default `version` command to be used in [CmdBox]
composite programs. Hence, it has no stand-alone variation.
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
foo version other
```

## Legal

Copyright (c) 2021 Robert S. Muhlestein
Released under the [Apache 2.0](LICENSE)

Contributors and project participants implicitly accept the 
[Developer Certificate of Authenticity (DCO)](DCO).

"CmdBox" and "cmdbox" are legal trademarks of Robert S. Muhlestein but
can be used freely to refer to the CmdBox project
<https://github.com/rwxrob/cmdbox> without limitation. To avoid
potential developer confusion, intentionally using these trademarks to
refer to other projects --- free or proprietary --- is prohibited.
