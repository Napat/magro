package magro

import (
	"fmt"
	"path/filepath"
	"runtime"
)

// Function return a string containing the function name of the call stack
func Function(depthList ...int) string {
	var skip int
	if depthList == nil {
		skip = 1
	} else {
		skip = depthList[0]
	}
	parentFunc, _, _, _ := runtime.Caller(skip)
	return runtime.FuncForPC(parentFunc).Name()
}

// File return a string containing the filename of the call stack
func File(depthList ...int) string {
	var skip int
	if depthList == nil {
		skip = 1
	} else {
		skip = depthList[0]
	}
	_, fullpath, _, _ := runtime.Caller(skip)
	filename := filepath.Base(fullpath)
	return filename
}

// FileFullpath return a string containing the filename(fullpath) of the call stack
func FileFullpath(depthList ...int) string {
	var skip int
	if depthList == nil {
		skip = 1
	} else {
		skip = depthList[0]
	}
	_, fullpath, _, _ := runtime.Caller(skip)
	return fullpath
}

// Line return a line number of the call stack
func Line(depthList ...int) int {
	var skip int
	if depthList == nil {
		skip = 1
	} else {
		skip = depthList[0]
	}
	_, _, line, _ := runtime.Caller(skip)
	return line
}

// Where return a string information containing the file name, function name
// and the line number of the call stack
func Where(depthList ...int) string {
	var skip int
	if depthList == nil {
		skip = 1
	} else {
		skip = depthList[0]
	}
	parentFunc, fullpath, line, _ := runtime.Caller(skip)
	return fmt.Sprintf("File: %s Function: %s Line: %d", filepath.Base(fullpath), runtime.FuncForPC(parentFunc).Name(), line)
}

// Info return a file name, function name
// and the line number of the call stack
func Info(depthList ...int) (filename string, funcName string, line int) {
	var skip int
	if depthList == nil {
		skip = 1
	} else {
		skip = depthList[0]
	}
	parentFunc, fullpath, line, _ := runtime.Caller(skip)
	filename = filepath.Base(fullpath)
	funcName = runtime.FuncForPC(parentFunc).Name()
	return filename, funcName, line
}
