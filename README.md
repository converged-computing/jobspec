# Flux JobSpec (Go)

This is a simple library that provides go structures for the Flux Framework [Jobspec](https://flux-framework.readthedocs.io/projects/flux-rfc/en/latest/spec_25.html) for use in other projects. 

## Usage

Build the examples:

```bash
make
```

Run all tests at once:

```bash
make test
```

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
