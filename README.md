# zigrt-curriculum

The curriculum for [Zigrt](https://github.com/chinny-labs/zigrt): the K-5 topic graph, its lessons and their problem templates, released as one versioned content pack. Licensed CC BY-SA 4.0.

This release is a 30-topic slice of Arithmetic K-5 that exercises every answer type, layout and figure. The rest of the graph is written strand by strand on top of it.

## Layout

The repository root is the pack (format: [Writing content](https://github.com/chinny-labs/zigrt/blob/main/docs/content/templates.md)).

| File | Holds |
| --- | --- |
| `manifest.yaml` | Pack ID, calendar version, `min_app_version`, courses |
| `topics.yaml` | Strands, every topic in teaching order, `requires` and `encompasses` edges, transitions |
| `topics/<id>.yaml` | One lesson per topic: knowledge points and their templates |
| `released-ids.json` | Every topic, knowledge point and template ID ever released |

## Working on it

`zigrt-content` is pinned by the `tool` line in `go.mod`, so it is always the `problem` and `content` packages of the Zigrt release it targets (Zigrt ADR 10). It needs Go 1.26.

```sh
go tool zigrt-content check                            # every gate, as CI runs it
go tool zigrt-content lint add-within-20               # one topic, fast
go tool zigrt-content sample add-within-20.kp1.t1 -seeds 1-20
go tool zigrt-content graph -strand fractions
go tool zigrt-content build -out dist                  # the pack file and its checksum
```

Read `sample` output, not just `check`: the gates catch broken templates, only reading catches a clumsy sentence.

To move to a newer Zigrt release: `go get -tool github.com/chinny-labs/zigrt/cmd/zigrt-content@vX.Y.Z`.

## Rules

- **IDs are permanent.** A released topic, knowledge point or template ID is never renamed, reused or deleted; a split or merge is a `transitions` entry. The `ids` gate enforces it against `released-ids.json`.
- **A released template keeps generating the same problems.** Fix its text freely; change what it generates and it needs a new template ID.
- Explanations stay short: the longest sentence allowed is 10 words in kindergarten, then 12, 14, 16, 18 and 20 by grade 5.
- Everything ships in the pack: no links outside it.

## Releasing

Packs use calendar versions (`2026.11.1`); `min_app_version` says which Zigrt can load them.

1. In a PR, set `version` in `manifest.yaml` and run `go tool zigrt-content ids -write` to record every new ID in `released-ids.json`. Merge it.
2. Tag the merge commit with the version and push the tag: `git tag 2026.11.1 && git push origin 2026.11.1`.

The release job checks that the tag matches the manifest and that every ID is recorded, runs every gate, and publishes `zigrt-k5-<version>.zip` and its `.sha256` as release assets. The Zigrt app pins one pack by version and checksum and embeds it.
