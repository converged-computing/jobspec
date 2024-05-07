package v1

var (
	jobspecVersion = 1
)

// The JobSpec is what the user writes to describe their work
type Jobspec struct {
	Version   int        `json:"version" yaml:"version"`
	Name      string     `json:"name,omitempty" yaml:"name,omitempty"`
	Resources Resources  `json:"resources,omitempty" yaml:"resources,omitempty"`
	Tasks     Tasks      `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Groups    Groups     `json:"groups,omitempty" yaml:"groups,omitempty"`
	Requires  []Requires `json:"requires,omitempty" yaml:"requires,omitempty"`
}

type Environment map[string]string
type Resources map[string]Resource
type Requires map[string]string
type Tasks []Task
type Groups []Group

type Task struct {
	Group     string     `json:"group,omitempty" yaml:"group,omitempty"`
	Name      string     `json:"name,omitempty" yaml:"name,omitempty"`
	Replicas  int        `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	Resources string     `json:"resources,omitempty" yaml:"resources,omitempty"`
	Command   []string   `json:"command,omitempty" yaml:"command,omitempty"`
	Requires  []Requires `json:"requires,omitempty" yaml:"requires,omitempty"`
}

type Group struct {
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Tasks     Tasks  `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Groups    Groups `json:"groups,omitempty" yaml:"groups,omitempty"`
	Resources string `json:"resources,omitempty" yaml:"resources,omitempty"`
}

type Resource struct {
	Type       string              `yaml:"type,omitempty" json:"type,omitempty"`
	Unit       string              `yaml:"unit,omitempty" json:"unit,omitempty"`
	Count      int32               `yaml:"count,omitempty" json:"count,omitempty"`
	With       []Resource          `yaml:"with,omitempty" json:"with,omitempty"`
	Label      string              `yaml:"label,omitempty" json:"label,omitempty"`
	Exclusive  bool                `yaml:"exclusive,omitempty" json:"exclusive,omitempty"`
	Requires   map[string]Requires `json:"requires,omitempty" yaml:"requires,omitempty"`
	Attributes Attributes          `json:"attributes,omitempty" yaml:"attributes,omitempty"`
	Schedule   bool                `yaml:"schedule,omitempty" json:"schedule,omitempty"`
}

type Attributes struct {
	Duration    string      `yaml:"duration,omitempty" json:"duration,omitempty"`
	Cwd         string      `yaml:"cwd,omitempty" json:"cwd,omitempty"`
	Environment Environment `yaml:"environment,omitempty" json:"environment,omitempty"`
}

// Temporary holder type for named resource and paired requires
/*type ScheduledSlot struct {
	Resource Resource
	Requires Requires
}

// A workload transforms a jobspec into workload units, each a schedulable unit
// This means that this is namespaced by the slots that we have, which are parsed
// here with GetScheduledNamedSlots
type Workload struct {

	// A slot can be for a task or group - it does not matter
	Slots map[string]Slot `yaml:"slots,omitempty" json:"slots,omitempty"`
}

// Slot is a unit of work with a top level resource request,
// a top level requirements specification (for an individual group or task
// at the top level)
type Slot struct {
	Resources Resource     `yaml:"resource,omitempty" json:"resource,omitempty"`
	Requires  SlotRequires `yaml:"requires,omitempty" json:"requires,omitempty"`
}

// WorkloadRequires contins task / group level requirements for easy parsing
type SlotRequires struct {

	// Global describes the requirements for the entire task or group
	Global Requires `yaml:"global,omitempty" json:"global,omitempty"`

	// Lookup of tasks and group requirements, if needed
	Tasks  map[string]Requires `yaml:"tasks,omitempty" json:"tasks,omitempty"`
	Groups map[string]Requires `yaml:"groups,omitempty" json:"groups,omitempty"`
}

// The WorkSpec describes the jobspec in schedulable units
/*type Workspec struct {
	Version   int                 `json:"version" yaml:"version"`
	Name      string              `json:"name,omitempty" yaml:"name,omitempty"`
	Resources Resources           `json:"resources,omitempty" yaml:"resources,omitempty"`
	Tasks     Tasks               `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Groups    Groups              `json:"groups,omitempty" yaml:"groups,omitempty"`
	Requires  map[string]Requires `json:"requires,omitempty" yaml:"requires,omitempty"`
}*/
