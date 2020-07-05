package projectpath

// internal package to use local project path art runtime, I find helpful with debugging/logging
// as suggested by Xeoncross in https://stackoverflow.com/a/58294680/4326287

import (
	"path/filepath"
	"runtime"
)

var (
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	Root = filepath.Join(filepath.Dir(b), "../..")
)
