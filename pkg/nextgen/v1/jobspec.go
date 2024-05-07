package v1

import (
	"encoding/json"
	"os"
	"reflect"

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

// GetSlots returns all slots in resources
func (js *Jobspec) GetSlots() []Resource {

	emptyResource := Resource{}
	slots := []Resource{}
	fauxSlots := []Resource{}

	// We first look explicitly for slots
	for name, resource := range js.Resources {

		// A jobspec resource with no slot is assumed to have
		// a slot at the top level. We wrap in a faux slot
		fauxSlots = append(fauxSlots, generateFauxSlot(name, resource))

		// Slot at the top level already!
		if resource.Type == "slot" {
			slots = append(slots, resource)
		}
		for _, with := range resource.With {
			slot := getSlots(with)
			if !reflect.DeepEqual(emptyResource, slot) {
				slots = append(slots, slot)
			}
		}
	}

	// If we found no slots, assume all top level resources are slots
	if len(slots) == 0 {
		return fauxSlots
	}
	return slots
}

// GetScheduledSlots returns all slots marked for scheduling
// If none are marked, we assume they all are
func (js *Jobspec) GetScheduledSlots() []Resource {

	slots := js.GetSlots()
	scheduled := []Resource{}

	allTrue := true
	for _, slot := range slots {
		if slot.Schedule {
			allTrue = false
			scheduled = append(scheduled, slot)
		}
	}
	// If none marked for scheduling, they all should be
	if allTrue {
		return slots
	}
	return scheduled
}

// GetScheduledNamedSlots returns slots as a lookup by name
func (js *Jobspec) GetScheduledNamedSlots() map[string]Resource {

	slots := js.GetScheduledSlots()
	named := map[string]Resource{}
	for _, slot := range slots {
		named[slot.Label] = slot
	}
	return named
}

// A fauxSlot will only be use if we don't have any actual slots
func generateFauxSlot(name string, resource Resource) Resource {
	return Resource{
		Type:     "slot",
		Label:    name,
		Count:    1,
		Schedule: resource.Schedule,
		With:     []Resource{resource},
	}
}

// getSlots is a recursive helper function that takes resources explicitly
func getSlots(resource Resource) Resource {

	emptyResource := Resource{}

	// If we find a slot, stop here
	if resource.Type == "slot" {
		return resource
	}
	for _, with := range resource.With {
		slot := getSlots(with)

		// If we find a slot
		if !reflect.DeepEqual(emptyResource, slot) {
			return slot
		}
	}
	return emptyResource
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
