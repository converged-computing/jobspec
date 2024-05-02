package v1

import (
	"encoding/json"
	"fmt"
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

// GetResources to get resource groups across the jobspec
// this is intended for graph scheduling
func (js *Jobspec) GetResources(data []byte) ([]Resource, error) {

	// We assume every discovered resource is a unique satisfy
	resources := []Resource{}

	// Make sure all task and group resources are known
	for _, task := range js.Tasks {
		r, err := js.getResources(task)
		if err != nil {
			return resources, err
		}
		resources = append(resources, r)
	}
	for _, group := range js.Groups {
		r, err := js.getResources(group)
		if err != nil {
			return resources, err
		}
		resources = append(resources, r)
	}
	return resources, nil
}

// getResources unwraps resources. If there is a named string, we assume
// in reference to a named resource group. We will need a strategy to combine
// these intelligently when we ask for a match - right now just assuming
// separate groups
func (js *Jobspec) getResources(resources interface{}) (Resource, error) {
	resource := Resource{}
	switch v := resources.(type) {
	case string:
		resourceKey := resources.(string)
		spec, ok := js.Resources[resourceKey]
		if !ok {
			return resource, fmt.Errorf("task is missing resource")
		}
		return spec, nil
	case Resource:
		return resources.(Resource), nil
	default:
		return resource, fmt.Errorf("type %s is unknown", v)
	}
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
	if js.Name != "" {
		name = js.Name
	}
	return name
}
