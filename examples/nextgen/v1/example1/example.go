package main

import (
	"flag"
	"fmt"
	"os"

	v1 "github.com/compspec/jobspec-go/pkg/nextgen/v1"
)

func main() {
	fmt.Println("This example reads, parses, and validates a Jobspec")

	// Assumes running from the root
	fileName := flag.String("json", "examples/nextgen/v1/example1/jobspec.yaml", "yaml file")
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

	out, err := js.JobspecToYaml()
	if err != nil {
		fmt.Printf("error marshalling %s:%s\n", yamlFile, err)
		os.Exit(1)
	}
	fmt.Println(string(out))

	// One example of json
	out, err = js.JobspecToJson()
	if err != nil {
		fmt.Printf("error marshalling %s:%s\n", yamlFile, err)
		os.Exit(1)
	}
	fmt.Println(string(out))

	// Test getting slots
	slots, err := js.GetSlots()
	if err != nil {
		fmt.Printf("error getting slots:%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d slots\n", len(slots))
	for _, slot := range slots {
		fmt.Println(slot)
	}

	slots, err = js.GetScheduledSlots()
	if err != nil {
		fmt.Printf("error getting scheduled slots:%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d scheduled slots\n", len(slots))
	for _, slot := range slots {
		fmt.Println(slot)
	}

	named, err := js.GetScheduledNamedSlots()
	if err != nil {
		fmt.Printf("error getting scheduled named slots:%s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d scheduled named slots\n", len(slots))
	for name, _ := range named {
		fmt.Println(name)
	}
}
