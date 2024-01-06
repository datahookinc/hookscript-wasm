package main

import (
	"bytes"
	"fmt"
	"syscall/js"

	hs "github.com/datahookinc/hookscript"
)

func main() {

	js.Global().Set("RunScript", RunScript())

	// Prevent the application from ending
	select {}
}

func RunScript() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		// The function receives the promise's (res, rej) arguments
		handler := js.FuncOf(func(this js.Value, promiseArgs []js.Value) any {

			res := promiseArgs[0]
			rej := promiseArgs[1]

			// index out of range with length 0
			if len(args) != 1 {
				rej.Invoke("Invalid number of arguments passed to function")
			}

			script := args[0].String()

			// print() statements send their values to the buffer to be passed to the console
			var b bytes.Buffer
			rt := hs.NewWASMRuntime()
			err := hs.NewRuntime(script, &b, rt)
			if err != nil {
				rej.Invoke(fmt.Sprintf("Error running script: %s", err.Error()))
				return nil
			}
			// return the value from the bytes as a string
			res.Invoke(b.String())
			return nil
		})

		promiseConstructor := js.Global().Get("Promise")
		return promiseConstructor.New(handler)
	})
}
