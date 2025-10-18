package main

import (
	"buf.build/go/bufplugin/check"
	"github.com/project-planton/project-planton/buf/lint/planton/rules"
)

func main() {
	check.Main(&check.Spec{
		Rules: []*check.RuleSpec{
			rules.DefaultRequiresOptionalRule,
		},
	})
}

