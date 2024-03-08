# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [4.0.0]

### Added

- HTTPBin
  - argocd: added LoadBalancer profile
- Kind
  - bl64/control: added default network ranges

### Changed

- ArgoCD
  - bl64/control: **breaking change** normalize parameters: -c: create, -s: show, -d: destroy
- FluxCD
  - bl64/control: **breaking change** normalize parameters: -c: create, -s: show, -d: destroy
- NGINX
  - bl64/control: **breaking change** normalize parameters: -c: create, -s: show, -d: destroy
- Kind
  - bl64/control: **breaking change** normalize parameters: -c: create, -s: show, -d: destroy
  - bl64/control: added default network ranges
- Minikube
  - bl64/control: **breaking change** normalize parameters: -c: create, -s: show, -d: destroy
- GitOps64
  - Minimized list of services enabled by default to simplify initial setup

## [3.2.0]

### Changed

- ArgoCD
  - Migrated to modular bl64
- FluxCD
  - Migrated to modular bl64
- GitHub
  - Migrated to modular bl64
- GitOps64
  - Migrated to modular bl64
- NGINX
  - Migrated to modular bl64
- Kind
  - Migrated to modular bl64
- Minikube
  - Migrated to modular bl64

## [3.1.0]

### Added

- Modules
  - NGINX
- kind/control
  - Show node IP range

### Fixed

- argocd/control-service
  - Password generator

## [3.0.0]

### Added

- Modules
  - MySQL Operator
  - Kiali
  - Wordpress
  - Kind
  - HTTP-Echo
  - MetalLB
- CertManager
  - Added Trust-Manager component
- Resources
  - Private CA
  - InnoDB cluster
  - Gateway
- istio-system
  - Added ArgoCD support
- istio-ingress
  - Added ArgoCD support
- httpbin
  - Added ArgoCD support

### Changed

- all: **breaking change** Reorganized directory structure to group modules based on type: application, infrastructure, kubernetes, resource

## [2.0.0]

### Added

- Modules
  - K8S CSI Secrets Store
  - Vault
  - Sealed Secrets
  - FluentBit
  - OpenSearch
  - Grafana
  - Loki
- Minikube/addons
  - csi-hostpath-driver
  - volumesnapshots
- GitOps64/setup
  - **breaking change** gitops64_setup_load: added profile parameter
- ArgoCD/control
  - Added profile support
- Minikube/control
  - Added profile support
- FluxCD/control
  - Added profile support

### Changed

- Minikube/control
  - **breaking change** Addons configuration variables (cluster.yaml):
    - minikube.cluster.addons.install: variable removed
    - minikube.cluster.addons: is now a string variable that contains a space separated list of minikube addons
- Prometheus
  - Added ArgoCD support

## [1.0.0]

### Added

- Modules
  - ArgoCD
  - GitHub

### Changed

- GitOps64
  - Reorganized directory structure to support module profiles. Profiles are a group of configuration and manifest files that represents a specific scenario. For example: metrics-server with profiles for argocd and fluxcd
  - Configuration: added base configuration file: etc/all/gitops/setup.yaml
  - Library: added gitops64_setup_load() function to import module YAML configuration
- Metrics Server
  - ArgoCD support
- Gateway-API
  - ArgoCD support
- Cert-Manager
  - ArgoCD support

## [0.2.0]

### Added

- Modules:
  - CertManager
  - Rook
  - Istio-System
  - Gateway-API
  - Flagger
  - Prometheus
- Demo Applications
  - HTTPBin
- FluxCD/control
  - Uninstall feature
- Minikube/control
  - Delete feature

## [0.1.0]

### Added

- Modules:
  - Metrics Server

## [0.0.1]

### Added

- Initial release

[4.0.0]: https://github.com/automation64/gitops64/compare/3.2.0...4.0.0
[3.2.0]: https://github.com/automation64/gitops64/compare/3.1.0...3.2.0
[3.1.0]: https://github.com/automation64/gitops64/compare/3.0.0...3.1.0
[3.0.0]: https://github.com/automation64/gitops64/compare/2.0.0...3.0.0
[2.0.0]: https://github.com/automation64/gitops64/compare/1.0.0...2.0.0
[1.0.0]: https://github.com/automation64/gitops64/compare/0.2.0...1.0.0
[0.2.0]: https://github.com/automation64/gitops64/compare/0.1.0...0.2.0
[0.1.0]: https://github.com/automation64/gitops64/compare/0.0.1...0.1.0
[0.0.1]: https://github.com/automation64/gitops64/releases/tag/0.0.1
