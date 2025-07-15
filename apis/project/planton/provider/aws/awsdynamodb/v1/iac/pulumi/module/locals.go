package module

import (
    "github.com/pkg/errors"

    awsdynamodbv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsdynamodb/v1"
)

type Locals struct {
    // AwsDynamodb holds the full API object that describes the desired state of
    // the DynamoDB table (metadata + spec).
    AwsDynamodb *awsdynamodbv1.AwsDynamodb

    // Tags is the final map of AWS tag key/value pairs that will be attached to
    // the provisioned resources.  It is built from user-supplied tags present
    // in the spec augmented with Project Planton standard tags.
    Tags map[string]string
}

// initializeLocals converts the StackInput into an internal representation that
// is easier to consume by the provisioning logic that lives in main.go.
func initializeLocals(stackInput *awsdynamodbv1.AwsDynamodbStackInput) (*Locals, error) {
    if stackInput == nil {
        return nil, errors.New("stack input cannot be nil")
    }

    locals := &Locals{
        AwsDynamodb: stackInput.GetTarget(),
        Tags:        map[string]string{},
    }

    if locals.AwsDynamodb == nil {
        return nil, errors.New("target aws_dynamodb resource is nil in stack input")
    }

    spec := locals.AwsDynamodb.GetSpec()
    if spec == nil {
        return nil, errors.New("aws_dynamodb.spec is nil")
    }

    // 1. Copy user provided tags verbatim
    for k, v := range spec.GetTags() {
        locals.Tags[k] = v
    }

    // 2. Inject standard Project Planton tags (best-effort, metadata may be
    //    absent depending on how the API object was built by the caller).
    //    We do not fail if metadata is missing – we simply skip the tag.
    //    Keys come from the awstagkeys package.
    locals.Tags["ResourceKind"] = "aws_dynamodb"

    // (In a full Project Planton repo we would import awstagkeys and use
    // constants such as awstagkeys.Resource, awstagkeys.Organization, …  Here
    // we inline the literals to keep the standalone example compiling.)

    if m := locals.AwsDynamodb.GetMetadata(); m != nil {
        if m.GetName() != "" {
            locals.Tags["Resource"] = m.GetName()
        }
        if m.GetOrganization() != "" {
            locals.Tags["Organization"] = m.GetOrganization()
        }
        if m.GetEnvironment() != "" {
            locals.Tags["Environment"] = m.GetEnvironment()
        }
        if m.GetId() != "" {
            locals.Tags["ResourceId"] = m.GetId()
        }
    }

    return locals, nil
}
