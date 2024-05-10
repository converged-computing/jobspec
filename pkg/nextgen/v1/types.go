package v1

var (
	jobspecVersion = 1
)

// The JobSpec is what the user writes to describe their work
type Jobspec struct {
	Version    int        `json:"version" yaml:"version"`
	Name       string     `json:"name,omitempty" yaml:"name,omitempty"`
	Resources  Resources  `json:"resources,omitempty" yaml:"resources,omitempty"`
	Tasks      Tasks      `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Groups     Groups     `json:"groups,omitempty" yaml:"groups,omitempty"`
	Requires   []Requires `json:"requires,omitempty" yaml:"requires,omitempty"`
	Attributes Attributes `json:"attributes,omitempty" yaml:"attributes,omitempty"`
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
	Type       string     `yaml:"type,omitempty" json:"type,omitempty"`
	Unit       string     `yaml:"unit,omitempty" json:"unit,omitempty"`
	Replicas   int32      `yaml:"replicas,omitempty" json:"replicas,omitempty"`
	Count      int32      `yaml:"count,omitempty" json:"count,omitempty"`
	With       []Resource `yaml:"with,omitempty" json:"with,omitempty"`
	Label      string     `yaml:"label,omitempty" json:"label,omitempty"`
	Exclusive  bool       `yaml:"exclusive,omitempty" json:"exclusive,omitempty"`
	Requires   []Requires `json:"requires,omitempty" yaml:"requires,omitempty"`
	Attributes Attributes `json:"attributes,omitempty" yaml:"attributes,omitempty"`
	Schedule   bool       `yaml:"schedule,omitempty" json:"schedule,omitempty"`
}

type Attributes struct {
	Duration    string      `yaml:"duration,omitempty" json:"duration,omitempty"`
	Cwd         string      `yaml:"cwd,omitempty" json:"cwd,omitempty"`
	Environment Environment `yaml:"environment,omitempty" json:"environment,omitempty"`
}
