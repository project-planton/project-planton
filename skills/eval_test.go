package skills

import (
	"io/fs"
	"os"
	"regexp"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

// evalCheck pins one piece of a recorded answer to the skill text that
// teaches it: Pattern, an RE2 regex, must match the whole of File inside
// the skill named Skill.
type evalCheck struct {
	Skill   string `yaml:"skill"`
	File    string `yaml:"file"`
	Pattern string `yaml:"pattern"`
}

// evalQuestion is one question a person asks an agent that loaded a skill,
// with the answer the platform actually gives.
type evalQuestion struct {
	ID       string      `yaml:"id"`
	Persona  string      `yaml:"persona"`
	Question string      `yaml:"question"`
	Answer   string      `yaml:"answer"`
	Checks   []evalCheck `yaml:"checks"`
}

type evalBank struct {
	Questions []evalQuestion `yaml:"questions"`
}

// The eval bank's answers are worth something only while the skill still
// teaches them. Every check is read against the embedded skill -- the bytes
// agents load -- so an edit that drops a sentence an answer rests on fails
// here, naming the question, instead of surfacing as a wrong answer in
// front of a person.
func TestSkillEvalBankIsFresh(t *testing.T) {
	raw, err := os.ReadFile("eval_questions.yaml")
	if err != nil {
		t.Fatalf("reading the eval bank: %v", err)
	}
	var bank evalBank
	if err := yaml.Unmarshal(raw, &bank); err != nil {
		t.Fatalf("parsing the eval bank: %v", err)
	}
	if len(bank.Questions) == 0 {
		t.Fatal("the eval bank has no questions")
	}

	seen := map[string]bool{}
	for _, q := range bank.Questions {
		if q.ID == "" || q.Persona == "" || strings.TrimSpace(q.Question) == "" || strings.TrimSpace(q.Answer) == "" {
			t.Errorf("question %q: id, persona, question and answer are all required", q.ID)
			continue
		}
		if seen[q.ID] {
			t.Errorf("question %q: duplicate id", q.ID)
			continue
		}
		seen[q.ID] = true
		if len(q.Checks) == 0 {
			t.Errorf("question %q: an answer needs at least one check pinning it to the skill", q.ID)
			continue
		}
		for i, c := range q.Checks {
			if c.File != "SKILL.md" && !(strings.HasPrefix(c.File, "references/") && strings.HasSuffix(c.File, ".md")) {
				t.Errorf("question %q check %d: file must be SKILL.md or references/<name>.md (got %q)", q.ID, i, c.File)
				continue
			}
			skill, ok := Skill(c.Skill)
			if !ok {
				t.Errorf("question %q check %d: no skill %q", q.ID, i, c.Skill)
				continue
			}
			re, err := regexp.Compile(c.Pattern)
			if err != nil || c.Pattern == "" {
				t.Errorf("question %q check %d: bad pattern %q: %v", q.ID, i, c.Pattern, err)
				continue
			}
			content, err := fs.ReadFile(skill, c.File)
			if err != nil {
				t.Errorf("question %q check %d: %s/%s is not in the skill: %v", q.ID, i, c.Skill, c.File, err)
				continue
			}
			if !re.Match(content) {
				t.Errorf("question %q check %d: %s/%s no longer teaches %q -- restore the sentence, or update the question's answer and checks in eval_questions.yaml",
					q.ID, i, c.Skill, c.File, c.Pattern)
			}
		}
	}
}
