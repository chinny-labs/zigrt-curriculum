# The K-5 syllabus

[`syllabus.yaml`](syllabus.yaml) lists every topic the finished pack will have: 255, confirming plan 4.4's estimate of "roughly 300" a little lower. Each topic has its grade, strand, prerequisites, Common Core standards, and any figure the app does not draw yet. `go test ./syllabus` checks:

- **The graph:** one root, no cycles, no edge to a missing topic or a later grade.
- **Written topics:** each one in `topics.yaml` agrees with the syllabus on grade, strand and title, and requires only its syllabus ancestors.

## Shape

| Strand | K | 1 | 2 | 3 | 4 | 5 | All |
| --- | --- | --- | --- | --- | --- | --- | --- |
| Counting and place value | 10 | 6 | 9 | 1 | 5 | 1 | 32 |
| Adding and taking away | 7 | 18 | 14 | 3 | 2 | | 44 |
| Multiplying | | | 2 | 11 | 12 | 6 | 31 |
| Dividing | | | | 6 | 7 | 4 | 17 |
| Fractions | | 2 | 1 | 6 | 14 | 16 | 39 |
| Decimals | | | | | 6 | 14 | 20 |
| Measuring (and data) | 2 | 5 | 13 | 13 | 8 | 7 | 48 |
| Shapes and space | 5 | 2 | 2 | 2 | 10 | 3 | 24 |
| **All** | **24** | **33** | **41** | **42** | **64** | **51** | **255** |

All 255 are written (finished 2026-11-18, grade by grade from K). Grade 4 is the largest because Common Core puts most of fractions, multi-digit multiplication and division, and angles there.

## Order of work

Lessons are written **grade by grade**, K first, so a younger child has a complete course early and the placement check is sound for each finished grade. A topic moves from the syllabus into `topics.yaml` when its lesson is written. Its ID, grade, strand and edges stay the same, and `encompasses` weights are set as each lesson is written.

## Figures the app needs first

65 topics need a picture the app's figure catalogue does not draw yet. Each is app work, in the Apache-2.0 `figure` package, done before the grade that first needs it:

| Grade | New figures |
| --- | --- |
| K | dots, ten frame, flat shapes, solid shapes, small scenes (above, below, beside), comparison pictures |
| 1 | base-ten blocks, ruler, table or tally chart |
| 2 | coins, picture graph, bar graph, line plot |
| 3 | rectilinear polygon with labelled sides, scales, split rectangle (area model) |
| 4 | lines and rays, angle, protractor, decimal number line |
| 5 | joined boxes |

Writing the lessons turned up a few more, added with their grade: `cut-shape` (1), `grid` (2), `symmetry` and `triangle` (4), and `plot` (5). The decimal number line needed nothing new.

`needs` at the end of `syllabus.yaml` describes each one.

## Reading

Kindergarten and early grade 1 children mostly can't read the prompts yet. Decided (2026-11-17): a parent reads along; Zigrt doesn't read aloud. So K and 1 lessons lean on pictures, keep prompts to a few words, and offer answers as short phrases a parent can read out.
