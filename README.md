# sensu-output-mutator

Creating releases for sensu is handled by GitHub Actions.

- run `./publish_release.sh v1.x.y` (and commit interactively)
- apply `sensu/mutator.yaml` via sensuctl
