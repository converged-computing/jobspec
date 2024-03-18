package experimental

var (
	jobspecVersion = 2
)

type Jobspec struct {
	Version    int        `json:"version" yaml:"version"`
	Resources  []Resource `json:"resources,omitempty" yaml:"resources,omitempty"`
	Task       Task       `json:"task,omitempty" yaml:"task,omitempty"`
	Attributes Attributes `json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

type Task struct {
	Command   []string               `json:"command,omitempty" yaml:"command,omitempty"`
	Slot      string                 `json:"slot,omitempty" yaml:"slot,omitempty"`
	Count     Count                  `json:"count,omitempty" yaml:"count,omitempty"`
	Resources map[string]interface{} `json:"resources,omitempty" yaml:"resources,omitempty"`
	Scripts   []Script               `json:"scripts,omitempty" yaml:"scripts,omitempty"`
	Transform []interface{}          `json:"transform,omitempty" yaml:"transform,omitempty"`
}

type Script struct {
	Name    string `json:"name" yaml:"name"`
	Content string `json:"content" yaml:"content"`
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

type Environment struct {
	Home string `json:"HOME"`
}
type System struct {
	Duration    int               `yaml:"duration,omitempty" json:"duration,omitempty"`
	Cwd         string            `yaml:"cwd,omitempty" json:"cwd,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty" json:"environment,omitempty"`
}
type Attributes struct {
	System System `yaml:"system,omitempty" json:"system"`
}
