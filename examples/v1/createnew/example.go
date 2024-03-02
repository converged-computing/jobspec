package main

import (
	"fmt"
	"os"

	v1 "github.com/compspec/jobspec-go/pkg/jobspec/v1"
)

func main() {
	fmt.Println("This example creates a new Jobspec")

	var nodes int32 = 2
	var tasks int32 = 12
	js, err := v1.NewSimpleJobspec("myjobspec", "echo hello world", nodes, tasks)
	if err != nil {
		fmt.Printf("error creating jobspec: %s\n", err)
		os.Exit(1)
	}

	// Validate the jobspec
	valid, err := js.Validate()
	if !valid || err != nil {
		fmt.Printf("schema is not valid:%s\n", err)
		os.Exit(1)
	} else {
		fmt.Println("schema is valid")
	}
	fmt.Println(js)

	out, err := js.JobspecToYaml()
	if err != nil {
		fmt.Printf("error marshalling jobspec: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(out))

	// One example of json
	out, err = js.JobspecToJson()
	if err != nil {
		fmt.Printf("error marshalling jobspec: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
