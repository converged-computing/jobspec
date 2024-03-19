package experimental

import (
	"encoding/json"
	"os"

	"sigs.k8s.io/yaml"

	"github.com/compspec/jobspec-go/pkg/schema"
)

// LoadJobspecYaml loads a jobspec from a yaml file path
func LoadJobspecYaml(yamlFile string) (*Jobspec, error) {
	js := Jobspec{}
	file, err := os.ReadFile(yamlFile)
	if err != nil {
		return &js, err
	}

	err = yaml.Unmarshal([]byte(file), &js)
	if err != nil {
		return &js, err
	}
	return &js, nil
}

// JobspectoYaml convets back to yaml (as string)
func (js *Jobspec) JobspecToYaml() (string, error) {
	out, err := yaml.Marshal(js)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// JobspectoJson convets back to json string
func (js *Jobspec) JobspecToJson() (string, error) {
	out, err := json.MarshalIndent(js, "", " ")
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// Validate converts to bytes and validate with jsonschema
func (js *Jobspec) Validate() (bool, error) {

	// Get back into bytes form
	out, err := yaml.Marshal(js)
	if err != nil {
		return false, err
	}
	// Validate the jobspec
	return schema.Validate(out, schema.SchemaUrl, Schema)

}

// Helper function to get a job name, derived from the command
func (js *Jobspec) GetJobName() string {

	// Generic name to fall back tp
	name := "app"

	// If we have tasks, we can get from the command
	// This entire set of checks is meant to be conservative
	// and avoid any errors with nil / empty arrays, etc.
	command := js.Task.Command
	if len(command) > 0 {
		name = command[0]
	}
	return name
}
