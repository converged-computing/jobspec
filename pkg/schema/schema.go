package schema

import (
	"fmt"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"sigs.k8s.io/yaml"
)

var (
	SchemaUrl = "https://raw.githubusercontent.com/compspec/jobspec-go/main/pkg/schema/schema.json"
)

// Validate assumes a yaml file, since jobspecs are typically in yaml
func Validate(jobspec []byte, schemaUrl, schema string) (bool, error) {
	s, err := jsonschema.CompileString(schemaUrl, schema)
	if err != nil {
		return false, err
	}

	var v interface{}
	err = yaml.Unmarshal(jobspec, &v)
	if err != nil {
		return false, err
	}
	fmt.Println(v)

	if err = s.Validate(v); err != nil {
		return false, err
	}

	err = s.Validate(v)
	if err != nil {
		return false, err
	}
	return true, nil
}
