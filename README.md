# Project: GITOps64

- [Project: GITOps64](#project-gitops64)
  - [Overview](#overview)
    - [Directory structure](#directory-structure)
    - [Applications](#applications)
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

GITOps64 is a working implementation of the GITOps methodoly intended to serve as a boilerplate for automated Kubernetes Application deployment projects.
The main purpose is to provide ready-to-use deployment modules for popular Kubernetes tools.

### Directory structure

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

### Applications

Applications are organized based on type:

- kubernetes: optional kubernetes components and APIs (e.g.: metrics-server, etc.)
- infrastructure: infrastructure providing services to applications (e.g.: cert-manager, sealed-secrets, etc.)
- applications: end-user applications
- resources: non-application objects consumed by application and infrastructure (e.g.: certificates, storage, etc/)

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
  - GitHub CLI
- Kubernetes
  - KubeCTL
  - Minikube or Kind
  - Helm CLI
- GitOps
  - FluxCD or ArgoCD CLI

#### Infrastructure

- Upstream repository: used to create and maintain the main GITOps repository
- Kubernetes cluster: target cluster that will be managed by GitOps

### Installation

- Login to GitHub using the GH CLI

```shell
gh auth login
```

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

- (FluxCD only) Turn off the following filter from `.gitignore` file:

```shell
# FluxCD / Upstream only: do not save FluxCD deployment
var/fluxcd/*/flux-system
```

- Review and update dev-environment configuration as needed: `etc/<ENVIRONMENT>`, in particular values marked with replacement tags `X_..._X`
  - [FluxCD](etc/dev/infrastructure/argocd/bl64-nodeport/gitops64.yaml)
  - [ArgoCD](etc/dev/infrastructure/fluxcd/bl64-minikube/gitops64.yaml)
  - [GitHub](etc/dev/infrastructure/github/bl64-default/gitops64.yaml)

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

- Create dev/test environment
  - Using Minikube

```shell
./src/kubernetes/minikube/bl64/control -e dev -p kvm-medium -c &&
./src/kubernetes/minikube/bl64/control -e dev -p kvm-medium -s
```

  - Using Kind

```shell
./src/kubernetes/kind/bl64/control -e dev -p medium -c &&
./src/kubernetes/kind/bl64/control -e dev -p medium -s
```

- Deploy GitOps service to Kubernetes
  - Using ArgoCD:

```shell
./src/infrastructure/argocd/bl64/control-service -e dev -p nodeport -c &&
./src/infrastructure/argocd/bl64/control-service -e dev -p nodeport -l &&
./src/infrastructure/argocd/bl64/control-application -e dev -p nodeport -c
```

  - Using FluxCD:

```shell
./src/infrastructure/fluxcd/bl64/control -e dev -p github -c
```

- (optional) Start NGINX to proxy MetalLB. This will allow local connections from the workstation to exposed cluster services of LoadBalancer type

```shell
./src/infrastructure/nginx/bl64/control -e dev -p k8s -c
```

## Operation

### Remove NGINX proxy

Use to stop and remove the NGIX container

```shell
./src/infrastructure/nginx/bl64/control -e dev -p k8s -d
```

### Remove Kind K8S cluster

Use to stop and **destroy** the cluster and deployed applications

```shell
./src/kubernetes/kind/bl64/control -e dev -p medium -d
```

## Contributing

Help on implementing new features and maintaining the code base is welcomed.

- [Guidelines](CONTRIBUTING.md)
- [Contributor Covenant Code of Conduct](CODE_OF_CONDUCT.md)

## License

[Apache-2.0](https://www.apache.org/licenses/LICENSE-2.0.txt)

## Author

- [SerDigital64](https://github.com/serdigital64)
