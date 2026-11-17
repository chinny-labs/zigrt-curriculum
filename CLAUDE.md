# zigrt-curriculum

The Zigrt K-5 curriculum, as a content pack. The repo root is the pack. The format is `docs/content/templates.md` in `chinny-labs/zigrt`; the design is that repo's `docs/design/plan.md` sections 3 and 4.

- Run `go tool zigrt-content check` before pushing, and read `go tool zigrt-content sample <template> -seeds 1-20` for every template you write or change.
- IDs in `released-ids.json` are permanent: never rename, reuse or delete one; add a `transitions` entry instead. Changing what a released template generates needs a new template ID.
- Student-facing text is short and at the topic's grade: the readability gate caps sentence length per grade. Prefer figures and worked steps over prose.
- Make multiple-choice distractors real mistakes (hands read the wrong way round, a forgotten carry).
- Build numbers from their parts instead of filtering random draws; keep `where` clauses cheap.
- Releases: see README. CI runs `check` and `build` on every PR.
