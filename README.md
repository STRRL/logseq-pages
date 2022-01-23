# logseq-pages

A cli tool for list all pages in logseq repo, marked with public or private.

When I using [logseq](https://github.com/logseq/logseq) to build my knowledge base and publish publicly, I found that "
Make it public for publishing" and "Make it private" could only be accessed on the application. So I have to manually
walk though all the pages and make sure that all I wanted to be published pages have been marked as public.

That is terrible experience.

So I would write the cli tools to:

- list all the pages
- [WIP] mark pages as public or private with interactive terminal ui

## Installation

You could use `go install github.com/strrl/logseq-pages/cmd/logseq-pages@latest` to install the tool.

Or clone this repo, and run `make` then you could find the binary `bin/logseq-pages`.

## Overview

```text
$ logseq-pages list --work-directory ~/playground/github/whatiknown        
+-------+-----------------------------------+--------+-----------------------+-----------------------------------------+
|     # | NAME                              | PUBLIC | ALIAS                 | PATH                                    |
+-------+-----------------------------------+--------+-----------------------+-----------------------------------------+
|     0 | 2022_01_04.md                     |        |                       | journals/2022_01_04.md                  |
|     1 | 2022_01_10.md                     |        |                       | journals/2022_01_10.md                  |
|     2 | 2022_01_20.md                     |        |                       | journals/2022_01_20.md                  |
|     3 | 2022_01_21.md                     |        |                       | journals/2022_01_21.md                  |
|     4 | About Me.md                       | *      | strrl, STRRL          | pages/About Me.md                       |
|     5 | Alfred.md                         | *      |                       | pages/Alfred.md                         |
|     6 | Arch Linux.md                     | *      |                       | pages/Arch Linux.md                     |
|     7 | Chaos Engineering.md              | *      |                       | pages/Chaos Engineering.md              |
|     8 | Chaos Mesh.md                     | *      | chaos-mesh, ChaosMesh | pages/Chaos Mesh.md                     |
|     9 | Cloud Native.md                   | *      |                       | pages/Cloud Native.md                   |
|    10 | Committer.md                      | *      |                       | pages/Committer.md                      |
|    11 | Container Orchestration.md        | *      |                       | pages/Container Orchestration.md        |
|    12 | Espanso.md                        | *      |                       | pages/Espanso.md                        |
|    13 | Fedora.md                         | *      |                       | pages/Fedora.md                         |
|    14 | Git Repository.md                 | *      | Git 仓库              | pages/Git Repository.md                 |
|    15 | I want to create my first page.md | *      |                       | pages/I want to create my first page.md |
|    16 | Kiu.md                            | *      |                       | pages/Kiu.md                            |
|    17 | Kubernetes.md                     | *      |                       | pages/Kubernetes.md                     |
|    18 | Linux Container.md                | *      |                       | pages/Linux Container.md                |
|    19 | Linux.md                          | *      |                       | pages/Linux.md                          |
|    20 | MacOS.md                          | *      |                       | pages/MacOS.md                          |
|    21 | Nonlinear System.md               | *      |                       | pages/Nonlinear System.md               |
|    22 | SemVer 2.0.md                     | *      |                       | pages/SemVer 2.0.md                     |
|    23 | Ubuntu Server.md                  | *      |                       | pages/Ubuntu Server.md                  |
|    24 | XinHua Dictionary.md              | *      |                       | pages/XinHua Dictionary.md              |
|    25 | chaos-mesh.dev.md                 | *      |                       | pages/chaos-mesh.dev.md                 |
|    26 | chaos-mesh.md                     | *      |                       | pages/chaos-mesh.md                     |
|    27 | contents.md                       | *      |                       | pages/contents.md                       |
|    28 | de facto standard.md              | *      |                       | pages/de facto standard.md              |
|    29 | favorites.md                      | *      |                       | pages/favorites.md                      |
|    30 | homelab.md                        | *      |                       | pages/homelab.md                        |
|    31 | logseq.md                         | *      |                       | pages/logseq.md                         |
|    32 | pkg.go.dev.md                     | *      |                       | pages/pkg.go.dev.md                     |
|    33 | text expander.md                  | *      |                       | pages/text expander.md                  |
|    34 | 非线性系统.md                     | *      |                       | pages/非线性系统.md                     |
+-------+-----------------------------------+--------+-----------------------+-----------------------------------------+
| TOTAL | 35                                |        |                       |                                         |
+-------+-----------------------------------+--------+-----------------------+-----------------------------------------+

```

With filter `private` and output `json`, beatified by `jq`:

```
$ logseq-pages list --output json --work-directory ~/playground/github/whatiknown --filter private | jq
[
  {
    "Name": "2022_01_04.md",
    "Alias": [],
    "Path": "journals/2022_01_04.md",
    "Public": false
  },
  {
    "Name": "2022_01_10.md",
    "Alias": [],
    "Path": "journals/2022_01_10.md",
    "Public": false
  },
  {
    "Name": "2022_01_20.md",
    "Alias": [],
    "Path": "journals/2022_01_20.md",
    "Public": false
  },
  {
    "Name": "2022_01_21.md",
    "Alias": [],
    "Path": "journals/2022_01_21.md",
    "Public": false
  }
]
```