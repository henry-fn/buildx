# docker buildx history inspect

<!---MARKER_GEN_START-->
Inspect a build

### Subcommands

| Name                                                 | Description                |
|:-----------------------------------------------------|:---------------------------|
| [`attachment`](buildx_history_inspect_attachment.md) | Inspect a build attachment |


### Options

| Name                  | Type     | Default  | Description                              |
|:----------------------|:---------|:---------|:-----------------------------------------|
| `--builder`           | `string` |          | Override the configured builder instance |
| `-D`, `--debug`       | `bool`   |          | Enable debug logging                     |
| [`--format`](#format) | `string` | `pretty` | Format the output                        |


<!---MARKER_GEN_END-->

## Description

Inspect a build record to view metadata such as duration, status, build inputs,
platforms, outputs, and attached artifacts. You can also use flags to extract
provenance, SBOMs, or other detailed information.

## Examples

### Inspect the most recent build

```console
$ docker buildx history inspect
Name:           buildx (binaries)
Context:        .
Dockerfile:     Dockerfile
VCS Repository: https://github.com/crazy-max/buildx.git
VCS Revision:   f15eaa1ee324ffbbab29605600d27a84cab86361
Target:         binaries
Platforms:      linux/amd64
Keep Git Dir:   true

Started:        2025-02-07 11:56:24
Duration:       1m  1s
Build Steps:    16/16 (25% cached)

Image Resolve Mode:     local

Materials:
URI                                                             DIGEST
pkg:docker/docker/dockerfile@1                                  sha256:93bfd3b68c109427185cd78b4779fc82b484b0b7618e36d0f104d4d801e66d25
pkg:docker/golang@1.23-alpine3.21?platform=linux%2Famd64        sha256:2c49857f2295e89b23b28386e57e018a86620a8fede5003900f2d138ba9c4037
pkg:docker/tonistiigi/xx@1.6.1?platform=linux%2Famd64           sha256:923441d7c25f1e2eb5789f82d987693c47b8ed987c4ab3b075d6ed2b5d6779a3

Attachments:
DIGEST                                                                  PLATFORM        TYPE
sha256:217329d2af959d4f02e3a96dcbe62bf100cab1feb8006a047ddfe51a5397f7e3                 https://slsa.dev/provenance/v0.2
```

### Inspect a specific build

```console
# Using a build ID
docker buildx history inspect qu2gsuo8ejqrwdfii23xkkckt

# Or using a relative offset
docker buildx history inspect ^1
```

### <a name="format"></a> Format the output (--format)

The formatting options (`--format`) pretty-prints the output to `pretty` (default),
`json` or using a Go template.

**Pretty output**

```console
$ docker buildx history inspect
Name:           buildx (binaries)
Context:        .
Dockerfile:     Dockerfile
VCS Repository: https://github.com/crazy-max/buildx.git
VCS Revision:   f15eaa1ee324ffbbab29605600d27a84cab86361
Target:         binaries
Platforms:      linux/amd64
Keep Git Dir:   true

Started:        2025-02-07 11:56:24
Duration:       1m  1s
Build Steps:    16/16 (25% cached)

Image Resolve Mode:     local

Materials:
URI                                                             DIGEST
pkg:docker/docker/dockerfile@1                                  sha256:93bfd3b68c109427185cd78b4779fc82b484b0b7618e36d0f104d4d801e66d25
pkg:docker/golang@1.23-alpine3.21?platform=linux%2Famd64        sha256:2c49857f2295e89b23b28386e57e018a86620a8fede5003900f2d138ba9c4037
pkg:docker/tonistiigi/xx@1.6.1?platform=linux%2Famd64           sha256:923441d7c25f1e2eb5789f82d987693c47b8ed987c4ab3b075d6ed2b5d6779a3

Attachments:
DIGEST                                                                  PLATFORM        TYPE
sha256:217329d2af959d4f02e3a96dcbe62bf100cab1feb8006a047ddfe51a5397f7e3                 https://slsa.dev/provenance/v0.2

Print build logs: docker buildx history logs g9808bwrjrlkbhdamxklx660b
```
**JSON output**

```console
$ docker buildx history inspect --format json
{
  "Name": "buildx (binaries)",
  "Ref": "5w7vkqfi0rf59hw4hnmn627r9",
  "Context": ".",
  "Dockerfile": "Dockerfile",
  "VCSRepository": "https://github.com/crazy-max/buildx.git",
  "VCSRevision": "f15eaa1ee324ffbbab29605600d27a84cab86361",
  "Target": "binaries",
  "Platform": [
    "linux/amd64"
  ],
  "KeepGitDir": true,
  "StartedAt": "2025-02-07T12:01:05.75807272+01:00",
  "CompletedAt": "2025-02-07T12:02:07.991778875+01:00",
  "Duration": 62233706155,
  "Status": "completed",
  "NumCompletedSteps": 16,
  "NumTotalSteps": 16,
  "NumCachedSteps": 4,
  "Config": {
    "ImageResolveMode": "local"
  },
  "Materials": [
    {
      "URI": "pkg:docker/docker/dockerfile@1",
      "Digests": [
        "sha256:93bfd3b68c109427185cd78b4779fc82b484b0b7618e36d0f104d4d801e66d25"
      ]
    },
    {
      "URI": "pkg:docker/golang@1.23-alpine3.21?platform=linux%2Famd64",
      "Digests": [
        "sha256:2c49857f2295e89b23b28386e57e018a86620a8fede5003900f2d138ba9c4037"
      ]
    },
    {
      "URI": "pkg:docker/tonistiigi/xx@1.6.1?platform=linux%2Famd64",
      "Digests": [
        "sha256:923441d7c25f1e2eb5789f82d987693c47b8ed987c4ab3b075d6ed2b5d6779a3"
      ]
    }
  ],
  "Attachments": [
    {
      "Digest": "sha256:450fdd2e6b868fecd69e9891c2c404ba461aa38a47663b4805edeb8d2baf80b1",
      "Type": "https://slsa.dev/provenance/v0.2"
    }
  ]
}
```

**Go template output**

```console
$ docker buildx history inspect --format "{{.Name}}: {{.VCSRepository}} ({{.VCSRevision}})"
buildx (binaries): https://github.com/crazy-max/buildx.git (f15eaa1ee324ffbbab29605600d27a84cab86361)
```
