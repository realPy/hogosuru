package initinterface

import "syscall/js"

var initInterface []func() js.Value

func RegisterInterface(f func() js.Value) {
	initInterface = append(initInterface, f)
}

func Init() {
	for _, f := range initInterface {
		f()
	}

	initInterface = make([]func() js.Value, 0)
}
