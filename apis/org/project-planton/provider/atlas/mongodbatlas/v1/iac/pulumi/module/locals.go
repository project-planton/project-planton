package module

import (
	mongodbatlasv1 "github.com/project-planton/project-planton/apis/org/project-planton/provider/atlas/mongodbatlas/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Locals struct {
	MongodbAtlas *mongodbatlasv1.MongodbAtlas
}

func initializeLocals(ctx *pulumi.Context, stackInput *mongodbatlasv1.MongodbAtlasStackInput) *Locals {
	locals := &Locals{}

	//assign value for the locals variable to make it available across the project
	locals.MongodbAtlas = stackInput.Target

	return locals
}
