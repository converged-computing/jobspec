package v1

import (
	"fmt"
	"strings"
)

// NewSimpleJobSpec generates a simple jobspec for nodes, command, tasks, and (optionally) a name
func NewSimpleJobspec(name, command string, nodes, tasks int32) (*Jobspec, error) {

	// If no name provided for the slot, use the first
	// work of the command
	if name == "" {
		parts := strings.Split(command, " ")
		name = strings.ToLower(parts[0])
	}
	if nodes < 1 {
		return nil, fmt.Errorf("nodes for the job must be >= 1")
	}
	if command == "" {
		return nil, fmt.Errorf("a command must be provided")
	}

	// The node resource is what we are asking for
	nodeResource := Resource{
		Type:  "node",
		Count: nodes,
	}

	// But we put it under the slot of a rack
	rackResource := Resource{
		Type:     "rack",
		Replicas: 1,
		Label:    name,
	}

	// If tasks are defined, this is total tasks across the nodes
	// We add to the slot
	if tasks != 0 {
		taskResource := Resource{
			Type:  "core",
			Count: tasks,
		}
		nodeResource.With = []Resource{taskResource}
	}

	// Resource name matches resources to named set
	rackResource.With = []Resource{nodeResource}
	resourceName := "task-resources"

	// Tasks reference the slot and command
	// Note: if we need better split can use "github.com/google/shlex"
	cmd := strings.Split(command, " ")
	taskResource := Task{
		Command:   cmd,
		Replicas:  1,
		Resources: resourceName,
	}
	tasklist := []Task{taskResource}

	return &Jobspec{
		Version:   jobspecVersion,
		Tasks:     tasklist,
		Resources: Resources{resourceName: rackResource},
	}, nil
}
