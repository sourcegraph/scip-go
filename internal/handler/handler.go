package handler

import "fmt"

var devMode bool = false

func SetDev(dev bool) {
	devMode = dev
}

func ErrOrPanic(format string, a ...any) error {
	if devMode {
		panic(fmt.Sprintf(format, a...))
	}

	return fmt.Errorf(format, a...)
}

func Println(a ...any) {
	if devMode {
		fmt.Println(a...)
	}
}
