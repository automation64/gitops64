# Project: GITOps64

- [Project: GITOps64](#project-gitops64)
  - [Overview](#overview)
  - [Design Requirements](#design-requirements)
  - [IaC Repository Organization](#iac-repository-organization)
    - [Main directory structure](#main-directory-structure)
    - [Modules types](#modules-types)
    - [Modules Organization](#modules-organization)
  - [Deployment](#deployment)
    - [Requirements](#requirements)
      - [Tools](#tools)
      - [Infrastructure](#infrastructure)
    - [Installation](#installation)
  - [Operation](#operation)
    - [Remove NGINX proxy](#remove-nginx-proxy)
    - [Remove Kind K8S cluster](#remove-kind-k8s-cluster)
  - [Contributing](#contributing)
  - [License](#license)
  - [Author](#author)

## Overview

GITOps64 is a working implementation of the GITOps methodology, intended to serve as a boilerplate for automated Kubernetes Application deployment projects.
The main purpose is to provide ready-to-use deployment modules for popular Kubernetes tools.

## Design Requirements

- DR01: separate CICD code from module code
- DR02: support both dev-time and run-time data and code
- DR03: use Infrastructure-as-Code (IaC) as much as possible
- DR04: support multiple IaC and not IaC tools
- DR05: keep code, configuration and run-time separated
- DR06: use environments to allow multiple instances of configuration and run-time for the same module code
- DR07: use profiles for code and configuration to allow multiple personalities for the same module
- DR08: use dedicated git repositories for module code as much as possible, for cases where the IaC tools supports remote modules
- DR09: keep the main branch stable (trunk branching strategy)
- DR09: do not use long-lived branches except for release ones
- DR10: group IaC tool code by module

## IaC Repository Organization

### Main directory structure

- `bin`: continuous integration, repository, development scripts
- `data`: location for persistent data files
- `docs`: application deployment documentation
- `etc`: deployment configuration
- `lib`: location for temporary development time libraries (excluded from GIT)
- `src`: deployment source code (Bash, Kustomization, etc)
- `test`: test scripts for deployment modules
- `tmp`: location for temporary content (excluded from GIT)
- `var`: location for persistent runtime data (included in GIT)
- `vault`: location for runtime secrets (excluded from GIT)

### Modules types

Modules are organized based by type:

- kubernetes: optional kubernetes components and APIs (e.g.: metrics-server, etc.)
- infrastructure: infrastructure providing services to applications (e.g.: cert-manager, sealed-secrets, etc.)
- applications: end-user applications
- resources: non-application objects consumed by application and infrastructure (e.g.: certificates, storage, etc/)

### Modules Organization

- 'etc/ENVIROMENT/MODULE_TYPE/MODULE_NAME/IAC_TOOL/PROFILE'
- 'src/MODULE_TYPE/MODULE_NAME/IAC_TOOL/PROFILE'

## Deployment

Use the following procedure to deploy GitOps64 for testing purposes on a Minikube cluster and forked GitHub repository.

### Requirements

#### Tools

- Virtualization
  - Docker and/or KVM
- Tools
  - GIT
  - Bash
  - Curl
  - YamlQuery (yq)
  - (\*) GitHub CLI
- Kubernetes
  - KubeCTL
  - Minikube or Kind
  - Helm CLI
- GitOps
  - (\*) FluxCD or ArgoCD CLI

Tools marked with (\*) are not needed locally if using the lab container (`dev-lab-kind`)

#### Infrastructure

- Upstream repository: used to create and maintain the main GITOps repository
- Kubernetes cluster: target cluster that will be managed by GitOps

### Installation

- Fork this repository to your GitHub account
- Clone the forked repository to your workstation

  ```shell
  git clone <FORKED_REPOSITORY>
  cd gitops64
  ```

- Initialize dev time resources

  ```shell
  ./bin/dev-lib-local &&
  ./bin/dev-lib-base
  ```

- Review and update dev-environment configuration as needed: `etc/<ENVIRONMENT>`, in particular values marked with replacement tags `X_..._X`

  - [FluxCD](etc/dev/infrastructure/argocd/bl64/nodeport/gitops64.yaml)
  - [ArgoCD](etc/dev/infrastructure/fluxcd/bl64/github/gitops64.yaml)
  - [GitHub](etc/dev/infrastructure/github/bl64/default/gitops64.yaml)
  - Additional files can be found using the command: `find etc -type f -exec grep -l -E 'X_[A-Z0-9_]+_X' {} \;`

- Review and update enabled modules: `var/<ENVIRONMENT>`:

  - FluxCD
    - [kubernetes](var/dev/fluxcd/kubernetes/)
    - [infrastructure](var/dev/fluxcd/infrastructure/)
    - [applications](var/dev/fluxcd/applications/)
  - ArgoCD
    - [kubernetes](var/dev/argocd/kubernetes/kustomization.yaml)
    - [infrastructure](var/dev/argocd/infrastructure/kustomization.yaml)
    - [resources](var/dev/argocd/resources/kustomization.yaml)
    - [applications](var/dev/argocd/applications/kustomization.yaml)

- Create dev/test kubernetes cluster

  - Using Minikube

  ```shell
  ./src/kubernetes/minikube/bl64/default/main -e dev -p kvm-medium -c &&
  ./src/kubernetes/minikube/bl64/default/main -e dev -p kvm-medium -s
  ```

- Using Kind

```shell
./src/kubernetes/kind/bl64/default/main -e dev -p medium -c &&
./src/kubernetes/kind/bl64/default/main -e dev -p medium -s
```

- (optional) Open lab container environment for running gitops tasks

```shell
./bin/dev-lab-kind
```

- Setup private GitHub repository access

  ```shell
  ./src/infrastructure/github/bl64/default/main -e dev -l &&
  ./src/infrastructure/github/bl64/default/main -e dev -a
  ```

- Deploy GitOps service to Kubernetes

  - Using ArgoCD:

  ```shell
  ./src/infrastructure/argocd/bl64/default/main -e dev -p nodeport -c &&
  ./src/infrastructure/argocd/bl64/default/main -e dev -p nodeport -l &&
  ./src/infrastructure/argocd/bl64/default/main -e dev -p nodeport -k
  ```

  - Using FluxCD:

    - Remove the following line from .`gitignore`: `var/fluxcd/*/flux-system`
    - Deploy FluxCD:

    ```shell
    ./src/infrastructure/fluxcd/bl64/default/main -e dev -p github -c
    ```

- (optional) Start NGINX to proxy MetalLB. This will allow local connections from the workstation to exposed cluster services of LoadBalancer type

```shell
./src/infrastructure/nginx/bl64/default/main -e dev -p k8s -c
```

## Operation

### Remove NGINX proxy

Use to stop and remove the NGIX container

```shell
./src/infrastructure/nginx/bl64/default/main -e dev -p k8s -d
```

### Remove Kind K8S cluster

Use to stop and **destroy** the cluster and deployed applications

```shell
./src/kubernetes/kind/bl64/default/main -e dev -p medium -d
```

## Contributing

Help on implementing new features and maintaining the code base is welcomed.

- [Guidelines](CONTRIBUTING.md)
- [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md)

## License

[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0.txt)

## Author

- [SerDigital64](https://github.com/serdigital64)
