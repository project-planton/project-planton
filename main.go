package main

import (
	"github.com/plantonhq/project-planton/cmd/project-planton"
	clipanic "github.com/plantonhq/project-planton/internal/cli/panic"
)

func main() {
	finished := new(bool)
	defer clipanic.Handle(finished)
	project_planton.Execute()
	*finished = true
}
