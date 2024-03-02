package main

import (
	"flag"
	"fmt"
	"os"

	v1 "github.com/compspec/jobspec-go/pkg/jobspec/v1"
)

func main() {
	fmt.Println("This example reads, parses, and validates a Jobspec")

	// Assumes running from the root
	fileName := flag.String("json", "examples/v1/example2/jobspec.yaml", "yaml file")
	flag.Parse()

	yamlFile := *fileName
	if yamlFile == "" {
		flag.Usage()
		os.Exit(0)
	}
	js, err := v1.LoadJobspecYaml(yamlFile)
	if err != nil {
		fmt.Printf("error reading %s:%s\n", yamlFile, err)
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
		fmt.Printf("error marshalling %s:%s\n", yamlFile, err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
