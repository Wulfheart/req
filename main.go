package main

import (
	"github.com/Wulfheart/req/cmd"
	"io/ioutil"
)
import v8 "rogchap.com/v8go"

func main() {
	cmd.Execute()

	x, _ := ioutil.ReadFile("./requester/js/main.js")
	s := string(x)
	ctx, _ := v8.NewContext()             // creates a new V8 context with a new Isolate aka VM
	_, err := ctx.RunScript(s, "main.js") // executes a script on the global context

	if err != nil {
		panic(err)
	}
	// val, err := ctx.RunScript("result", "main.js") // return a value in JavaScript back to Go
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("addition result: %s\n", val)

}
