# Autoversioned

Reference project with for versioning approach reference

## Components

* Simple golang webserver packaged into docker image
* Helm chart with helmfile and skaffold wrappers on top for deployment and local development respectively

## How to run

### Install helm package

* `helm repo add autoversioned https://mtrqq.github.io/autoversioned`
* `helm install auto autoversioned/autoversioned [--version <version>]`

### Local run via skaffold

* `skaffold dev --port-forward`

### Deploy using helmfile based on the environment

* `helmfile apply [-e (dev|staging|prod)]`

### Run using docker

* `docker run -p 8080:8080 mtrqq/autoversioned:<version>`

### Run using golang

* `go mod download`
* `go run api/serve.go`

## Versioning

### Commits

Versioning is based on [conventional commits](https://www.conventionalcommits.org/en/v1.0.0/) specification and powered by [commitizen](https://commitizen-tools.github.io/commitizen/) (note: there's a bunch of projects with conflicting names so it's better to use direct link). The only difference for developer when working with such kind of repository is that you have to follow certain commit naming convension, here's some examples

* `feat(core): implement ability to fetch cluster statistics`
* `fix: decrease memory footprint for statistics service`
* `chore: bump build-and-push action version`
* `refactor: migrate from log to zerolog`

If doesn't affect developers' life in case team is using squash and merge strategy for PRs as the only thing to track in this case is merge commit naming. Otherwise if you decided to go hard way with rebase strategy - you should check each commit message which is going into remote (commit-msg `pre-commit` hooks may come inhandy in this case).

### Release process

When we decide to release application at certain moment of development - we should manually trigger [release pipeline](.github/workflows/release.yaml) pipeline, most of the time you don't have to tune any parameters as it's automatical strategy is already pretty good, but anyway they are present in case you need them. Here's what this workflow does in brief (1.2.3 version is only for reference purposes):

* Creates `release/1.2.3` with already patched versions and changelog
* Creates pre-approved PR from `release/1.2.3` branch so that you can review all the changes done
* Creates draft release for tag `v1.2.3` with incremental changelog based on commits
* Builds following artifacts:
  * Linux based executables for API server on a most used architectures and attaches them to a release
  * API server docker image, pushes it to dockerhub repository, its' `imageid` is attached to release as a file
  * Helm chart package is created and added to a [GitHub Pages based repo index](https://mtrqq.github.io/autoversioned)
