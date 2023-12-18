# Project: GITOps64

- [Project: GITOps64](#project-gitops64)
  - [Overview](#overview)
    - [Directory structure](#directory-structure)
    - [Modules](#modules)
  - [Deployment](#deployment)
    - [Requirements](#requirements)
      - [Tools](#tools)
      - [Infrastructure](#infrastructure)
    - [Installation](#installation)
  - [Contributing](#contributing)
  - [License](#license)
  - [Author](#author)

## Overview

GITOps64 is a working implementation of the GITOps methodoly intended to serve as a boilerplate for automated Kubernetes Application deployment projects.
The main purpose is to provide ready-to-use deployment modules for popular Kubernetes tools.

### Directory structure

- `bin`: continous integration, repository, development scripts
- `data`: location for persistent data files
- `docs`: application deployment documentation
- `etc`: deployment configuration
- `lib`: location for temporary development time libraries (excluded from GIT)
- `src`: deployment source code (Bash, Kustomization, etc)
- `test`: test scripts for deployment modules
- `tmp`: location for temporary content (excluded from GIT)
- `var`: location for persistent runtime data (included in GIT)
- `vault`: location for runtime secrets (excluded from GIT)

### Modules

- GitOps
  - FluxCD
  - ArgoCD
  - Minikube
- Infrastructure
  - Certificate Manager
  - Flagger
  - GitHub
  - Istio
  - Metrics Server
  - Prometheus
  - Rook
- Demo Applications
  - HTTPBin

## Deployment

Use the following procedure to deploy GitOps64 for testing purposes on a Minikube cluster and forked GitHub repository.

### Requirements

#### Tools

- GIT
- KubeCTL
- FluxCD
- GitHub CLI
- Minikube

#### Infrastructure

- Upstream repository: used to create and maintain the main GITOps repository
- Kubernetes cluster: target cluster that will be managed by GitOps

### Installation

- Fork this repository to your GitHub account
- Clone the forked repository to your workstation

```shell
git clone <FORKED_REPOSITORY>
```

- Initialize dev time resources

```shell
cd gitops64
./bin/dev-lib
```

- Turn off the following filter from `.gitignore` file:

```shell
# FluxCD / Upstream only: do not save FluxCD deployment
var/fluxcd/*/flux-system
```

- Review and update dev-environment configuration as needed: `etc/dev`
- Deploy Minikube cluster

```shell
./src/minikube/bl64/control -s -e dev &&
./src/minikube/bl64/control -t -e dev
```

- Login to GitHub using the GH CLI

```shell
gh auth login
```

- Deploy FluxCD to Minikube using GitHub

```shell
./src/fluxcd/bl64/control -b -e dev
```

## Contributing

Help on implementing new features and maintaining the code base is welcomed.

- [Guidelines](CONTRIBUTING.md)
- [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md)

## License

[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0.txt)

## Author

- [SerDigital64](https://github.com/serdigital64)
