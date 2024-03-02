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

	// The slot is where we are doing an assessment for scheduling
	slot := Resource{
		Type:  "slot",
		Count: int32(1),
		Label: name,
	}

	// If tasks are defined, this is total tasks across the nodes
	// We add to the slot
	if tasks != 0 {
		taskResource := Resource{
			Type:  "core",
			Count: tasks,
		}
		slot.With = []Resource{taskResource}
	}

	// And then the entire resource spec is added to the top level node resource
	nodeResource.With = []Resource{slot}

	// Tasks reference the slot and command
	// Note: if we need better split can use "github.com/google/shlex"
	cmd := strings.Split(command, " ")
	taskResource := []Tasks{
		{
			Command: cmd,
			Slot:    name,
			Count:   Count{PerSlot: int32(1)},
		}}

	// Attributes are for the system, we aren't going to add them yet
	// attributes:
	// system:
	//   duration: 3600.
	//   cwd: "/home/flux"
	//   environment:
	// 	HOME: "/home/flux"
	// This is verison 1 as defined by v1 above
	return &Jobspec{
		Version:   jobspecVersion,
		Resources: []Resource{nodeResource},
		Tasks:     taskResource,
	}, nil
}
