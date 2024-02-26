package main

import (
	"flag"
	"fmt"
	"os"

	"sigs.k8s.io/yaml"

	"github.com/compspec/jobspec-go/pkg/schema"

	v1 "github.com/compspec/jobspec-go/pkg/jobspec/v1"
)

func main() {
	fmt.Println("This example reads, parses, and validates a Jobspec")

	// Assumes running from the root
	fileName := flag.String("json", "examples/v1/example5/jobspec.yaml", "yaml file")
	flag.Parse()

	yamlFile := *fileName
	if yamlFile == "" {
		flag.Usage()
		os.Exit(0)
	}
	file, err := os.ReadFile(yamlFile)
	if err != nil {
		fmt.Printf("error reading %s:%s\n", yamlFile, err)
		os.Exit(1)
	}

	// Validate the jobspec
	valid, err := schema.Validate(file, schema.SchemaUrl, v1.Schema)
	if !valid || err != nil {
		fmt.Printf("schema is not valid:%s\n", err)
		os.Exit(1)
	} else {
		fmt.Println("schema is valid")
	}

	js := v1.Jobspec{}
	err = yaml.Unmarshal([]byte(file), &js)
	if err != nil {
		fmt.Printf("error unmarshalling %s:%s\n", yamlFile, err)
		os.Exit(1)
	}
	fmt.Println(js)

	out, err := yaml.Marshal(js)
	if err != nil {
		fmt.Printf("error marshalling %s:%s\n", yamlFile, err)
		os.Exit(1)
	}
	fmt.Println(string(out))
}
