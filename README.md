# Flux Jobspec (Go)

This is a simple library that provides go structures for:

 - the Flux Framework [Jobspec](https://flux-framework.readthedocs.io/projects/flux-rfc/en/latest/spec_25.html) (package [jobspec](pkg/jobspec))
 - the [Next Generation Jobspec](https://compspec.github.io/jobspec) (package [nextgen](pkg/nextgen/))

Note for nextgen, since Go is more strict with typing, we accept a parsed JobSpec, meaning that all resources have been defined in the top level named section,
and are referenced by name in tasks. We will start assuming that a request for the resource groups should be satisfied within the same cluster, and each is a separate
match request.

## Usage

Build the examples:

```bash
make
```

Run all tests at once:

```bash
make test
```

Here is an example of usage. Note that this isn't a full program, but is intended to show helper functions.
In this small program, we load in a jobspec (from yaml) and then validate and serialize to each of json and yaml.

```go
package main

import (
	"fmt"
	"os"

	v1 "github.com/compspec/jobspec-go/pkg/jobspec/v1"
)

func main() {

  // Example 1: reading from file
  yamlFile := "examples/v1/example1/jobspec.yaml"

  // This is how to read from a YAML file
  js, err := v1.LoadJobspecYaml(yamlFile)
  // Validate the jobspec
  valid, err := js.Validate()

  // Convert back to YAML (print out as string(out))
  out, err := js.JobspecToYaml()

  // Convert back into JSON (also print string(out))
  out, err = js.JobspecToJson()

  // Example 2: creating from scratch
  var nodes int32 = 2
  var tasks int32 = 12
  js, err := v1.NewSimpleJobspec("myjobspec", "echo hello world", nodes, tasks)
  // proceed with equivalent functions above!
}
```

For full examples, see the [examples](examples/v1) directory.

### Version 1

You can run any example (and view the code) to see how it works!

```bash
./examples/v1/bin/example1
```
```console
This example reads, parses, and validates a Jobspec
map[attributes:map[system:map[cwd:/home/flux duration:3600 environment:map[HOME:/home/flux]]] resources:[map[count:4 type:node with:[map[count:1 label:default type:slot with:[map[count:2 type:core]]]]]] tasks:[map[command:[app] count:map[per_slot:1] slot:default]] version:1]
schema is valid
{1 [{node  4 [{slot  1 [{core  2 []  false}] default false}]  false}] [{[app] default {1 0}}] {{3600 /home/flux map[HOME:/home/flux]}}}
attributes:
  system:
    cwd: /home/flux
    duration: 3600
    environment:
      HOME: /home/flux
resources:
- count: 4
  type: node
  with:
  - count: 1
    label: default
    type: slot
    with:
    - count: 2
      type: core
tasks:
- command:
  - app
  count:
    per_slot: 1
  slot: default
version: 1
```

And that's mostly it! This library will eventually go into other Go projects that need Jobspec, and for now
just provides the basic types and validation.


## License

HPCIC DevTools is distributed under the terms of the MIT license.
All new contributions must be made under this license.

See [LICENSE](https://github.com/compspec/jobspec-go/blob/main/LICENSE),
[COPYRIGHT](https://github.com/compspec/jobspec-go/blob/main/COPYRIGHT), and
[NOTICE](https://github.com/compspec/jobspec-go/blob/main/NOTICE) for details.

SPDX-License-Identifier: (MIT)

LLNL-CODE- 842614
