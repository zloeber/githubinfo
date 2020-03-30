package version

import (
	"fmt"
	"runtime"
)

// GitCommit returns the git commit that was compiled. This will be filled in by the compiler.
var GitCommit string

// Version returns the main version number that is being run at the moment. This will be filled in by the compiler.
var Version string

// Application name. This will be filled in by the compiler.
var AppName string

// BuildDate returns the date the binary was built.  This will be filled in by the compiler.
var BuildDate string

// GoVersion returns the version of the go runtime used to compile the binary
var GoVersion = runtime.Version()

// OsArch returns the os and arch used to build the binary
var OsArch = fmt.Sprintf("%s %s", runtime.GOOS, runtime.GOARCH)
