package v1

var (
	jobspecVersion = 1
)

type Jobspec struct {
	Version   int               `json:"version" yaml:"version"`
	Name      string            `json:"name,omitempty" yaml:"name,omitempty"`
	Resources Resources         `json:"resources,omitempty" yaml:"resources,omitempty"`
	Tasks     Tasks             `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Groups    Groups            `json:"groups,omitempty" yaml:"groups,omitempty"`
	Requires  map[string]string `json:"requires,omitempty" yaml:"requires,omitempty"`
}

type Environment map[string]string
type Resources map[string]Resource
type Tasks []Task
type Groups []Group

type Task struct {
	Group      string     `json:"group,omitempty" yaml:"group,omitempty"`
	Name       string     `json:"name,omitempty" yaml:"name,omitempty"`
	Replicas   int        `json:"replicas,omitempty" yaml:"replicas,omitempty"`
	Resources  string     `json:"resources,omitempty" yaml:"resources,omitempty"`
	Command    []string   `json:"command,omitempty" yaml:"command,omitempty"`
	Attributes Attributes `json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

type Group struct {
	Name       string     `json:"name,omitempty" yaml:"name,omitempty"`
	Resources  string     `json:"resources,omitempty" yaml:"resources,omitempty"`
	Tasks      Tasks      `json:"tasks,omitempty" yaml:"tasks,omitempty"`
	Attributes Attributes `json:"attributes,omitempty" yaml:"attributes,omitempty"`
}

type Resource struct {
	Type      string     `yaml:"type,omitempty" json:"type,omitempty"`
	Unit      string     `yaml:"unit,omitempty" json:"unit,omitempty"`
	Count     int32      `yaml:"count,omitempty" json:"count,omitempty"`
	With      []Resource `yaml:"with,omitempty" json:"with,omitempty"`
	Label     string     `yaml:"label,omitempty" json:"label,omitempty"`
	Exclusive bool       `yaml:"exclusive,omitempty" json:"exclusive,omitempty"`
	Schedule  bool       `yaml:"schedule,omitempty" json:"schedule,omitempty"`
}

type Attributes struct {
	Duration    string      `yaml:"duration,omitempty" json:"duration,omitempty"`
	Cwd         string      `yaml:"cwd,omitempty" json:"cwd,omitempty"`
	Environment Environment `yaml:"environment,omitempty" json:"environment,omitempty"`
}
