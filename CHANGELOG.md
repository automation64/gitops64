# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

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

[1.0.0]: https://github.com/automation64/gitops64/compare/0.2.0...1.0.0
[0.2.0]: https://github.com/automation64/gitops64/compare/0.1.0...0.2.0
[0.1.0]: https://github.com/automation64/gitops64/compare/0.0.1...0.1.0
[0.0.1]: https://github.com/automation64/gitops64/releases/tag/0.0.1
