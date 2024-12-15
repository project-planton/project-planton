package apidocs

import (
	"github.com/pkg/errors"
	gendoc "github.com/pseudomuto/protoc-gen-doc"
	"google.golang.org/protobuf/proto"
)

// GetMessageDocs finds and returns the documentation for the given proto.Message from the loaded docs template.
// It searches by the message's fully-qualified name.
func GetMessageDocs(msg proto.Message) (*gendoc.Message, error) {
	apiDocsJson, err := getApiDocsJson()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get API docs")
	}

	messageDescriptor := msg.ProtoReflect().Descriptor()
	fullName := string(messageDescriptor.FullName())

	for _, f := range apiDocsJson.Files {
		for _, m := range f.Messages {
			if m.FullName == fullName {
				return m, nil
			}
		}
	}

	return nil, errors.Errorf("documentation not found for message: %s", fullName)
}
