// Package syllabus checks the K-5 syllabus: its graph and its agreement
// with the topics already written.
package syllabus

import (
	"os"
	"regexp"
	"slices"
	"strings"
	"testing"

	"go.yaml.in/yaml/v3"
)

type topic struct {
	ID        string   `yaml:"id"`
	Title     string   `yaml:"title"`
	Grade     string   `yaml:"grade"`
	Strand    string   `yaml:"strand"`
	Requires  []string `yaml:"requires"`
	Standards []string `yaml:"standards"`
	Needs     []string `yaml:"needs"`
}

type file struct {
	Topics []topic               `yaml:"topics"`
	Needs  map[string]string     `yaml:"needs"`
	Strand []struct{ ID string } `yaml:"strands"`
}

var (
	grades   = []string{"K", "1", "2", "3", "4", "5"}
	idRE     = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)
	stdRE    = regexp.MustCompile(`^(K|[1-5])\.(CC|OA|NBT|NF|MD|G)\.[A-D]\.[0-9]+[a-e]?$`)
	strandIn = map[string]bool{}
)

func load(t *testing.T, path string, v any) {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := yaml.Unmarshal(b, v); err != nil {
		t.Fatalf("%s: %v", path, err)
	}
}

func TestSyllabus(t *testing.T) {
	var s, pack file
	load(t, "syllabus.yaml", &s)
	load(t, "../topics.yaml", &pack)
	for _, st := range pack.Strand {
		strandIn[st.ID] = true
	}

	byID := map[string]topic{}
	gradeOf := func(g string) int { return slices.Index(grades, g) }
	perGrade := map[string]int{}
	for _, tp := range s.Topics {
		if !idRE.MatchString(tp.ID) {
			t.Errorf("%s: not a kebab-case ID", tp.ID)
		}
		if _, dup := byID[tp.ID]; dup {
			t.Errorf("%s: listed twice", tp.ID)
		}
		byID[tp.ID] = tp
		perGrade[tp.Grade]++
		if gradeOf(tp.Grade) < 0 {
			t.Errorf("%s: grade %q", tp.ID, tp.Grade)
		}
		if !strandIn[tp.Strand] {
			t.Errorf("%s: strand %q is not in topics.yaml", tp.ID, tp.Strand)
		}
		if tp.Title == "" || strings.HasSuffix(tp.Title, ".") || strings.HasSuffix(tp.Title, "?") {
			t.Errorf("%s: title %q", tp.ID, tp.Title)
		}
		if len(tp.Standards) == 0 {
			t.Errorf("%s: no standard", tp.ID)
		}
		for _, sd := range tp.Standards {
			if !stdRE.MatchString(sd) {
				t.Errorf("%s: standard %q", tp.ID, sd)
			}
		}
		for _, n := range tp.Needs {
			if _, ok := s.Needs[n]; !ok {
				t.Errorf("%s: need %q is not described", tp.ID, n)
			}
		}
	}

	// Edges: to known topics, never to a later grade, no cycles; one root.
	roots := 0
	for _, tp := range s.Topics {
		if len(tp.Requires) == 0 {
			roots++
		}
		for _, r := range tp.Requires {
			req, ok := byID[r]
			if !ok {
				t.Errorf("%s requires %s, which is not in the syllabus", tp.ID, r)
				continue
			}
			if gradeOf(req.Grade) > gradeOf(tp.Grade) {
				t.Errorf("%s (grade %s) requires %s from a later grade %s", tp.ID, tp.Grade, r, req.Grade)
			}
		}
	}
	if roots != 1 {
		t.Errorf("%d topics require nothing; want one root", roots)
	}
	state := map[string]int{} // 1 visiting, 2 done
	var visit func(id string, path []string)
	visit = func(id string, path []string) {
		switch state[id] {
		case 1:
			t.Errorf("cycle: %s", strings.Join(append(path, id), " → "))
			return
		case 2:
			return
		}
		state[id] = 1
		for _, r := range byID[id].Requires {
			if _, ok := byID[r]; ok {
				visit(r, append(path, id))
			}
		}
		state[id] = 2
	}
	for _, tp := range s.Topics {
		visit(tp.ID, nil)
	}

	// Written topics agree with the syllabus.
	for _, w := range pack.Topics {
		sy, ok := byID[w.ID]
		if !ok {
			t.Errorf("%s is written but not in the syllabus", w.ID)
			continue
		}
		if sy.Grade != w.Grade || sy.Strand != w.Strand || sy.Title != w.Title {
			t.Errorf("%s: syllabus says %s/%s/%q, topics.yaml %s/%s/%q", w.ID, sy.Grade, sy.Strand, sy.Title, w.Grade, w.Strand, w.Title)
		}
		// While its grade is unwritten, a written topic may require an
		// earlier topic directly that the syllabus reaches through new
		// ones in between; it must still be one of its ancestors.
		for _, r := range w.Requires {
			if !ancestor(byID, w.ID, r) {
				t.Errorf("%s: topics.yaml requires %s, which the syllabus doesn't lead to", w.ID, r)
			}
		}
	}
	t.Logf("%d topics: K %d, 1 %d, 2 %d, 3 %d, 4 %d, 5 %d; %d written", len(s.Topics),
		perGrade["K"], perGrade["1"], perGrade["2"], perGrade["3"], perGrade["4"], perGrade["5"], len(pack.Topics))
}

// ancestor reports whether a requires b, directly or through other topics.
func ancestor(byID map[string]topic, a, b string) bool {
	seen := map[string]bool{}
	var walk func(id string) bool
	walk = func(id string) bool {
		for _, r := range byID[id].Requires {
			if r == b {
				return true
			}
			if !seen[r] {
				seen[r] = true
				if walk(r) {
					return true
				}
			}
		}
		return false
	}
	return walk(a)
}
