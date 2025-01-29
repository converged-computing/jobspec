package v1

var (
	jobspecVersion = 1
)

type Jobspec struct {
	Version    int        `json:"version" yaml:"version"`
	Resources  []Resource `json:"resources,omitempty" yaml:"resources,omitempty"`
	Tasks      []Tasks    `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Attributes Attributes `json:"attributes" yaml:"attributes"`
}

type Constraints struct {
	Hostlist []string `json:"hostlist,omitempty" yaml:"hostlist,omitempty"`
}

type Tasks struct {
	Command []string `json:"command,omitempty" yaml:"command,omitempty"`
	Slot    string   `json:"slot,omitempty" yaml:"slot,omitempty"`
	Count   Count    `json:"count,omitempty" yaml:"count,omitempty"`
}

type Count struct {
	PerSlot int32 `json:"per_slot,omitempty" yaml:"per_slot,omitempty"`
	Total   int32 `json:"total,omitempty" yaml:"total,omitemtpy"`
}

type Resource struct {
	Type      string     `yaml:"type,omitempty" json:"type,omitempty"`
	Unit      string     `yaml:"unit,omitempty" json:"unit,omitempty"`
	Count     int32      `yaml:"count,omitempty" json:"count,omitempty"`
	With      []Resource `yaml:"with,omitempty" json:"with,omitempty"`
	Label     string     `yaml:"label,omitempty" json:"label,omitempty"`
	Exclusive bool       `yaml:"exclusive,omitempty" json:"exclusive,omitempty"`
}

type System struct {
	Duration    int               `yaml:"duration,omitempty" json:"duration,omitempty"`
	Cwd         string            `yaml:"cwd,omitempty" json:"cwd,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty" json:"environment,omitempty"`
	Constraints Constraints       `yaml:"constraints,omitempty" json:"constraints,omitempty"`
}

type Attributes struct {
	System System `yaml:"system,omitempty" json:"system,omitempty"`
}
