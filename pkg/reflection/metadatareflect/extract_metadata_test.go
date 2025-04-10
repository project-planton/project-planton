package metadatareflect

import (
	awss3bucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awss3bucket/v1"
	"testing"

	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"google.golang.org/protobuf/proto"
)

func TestExtractMetadata(t *testing.T) {
	tests := []struct {
		name  string
		input proto.Message
		want  *shared.ApiResourceMetadata
	}{
		{
			name: "when metadata is set should return the metadata from input",
			input: &awss3bucketv1.AwsS3Bucket{
				Metadata: &shared.ApiResourceMetadata{
					Id: "test-id",
				},
			},
			want: &shared.ApiResourceMetadata{Id: "test-id"},
		}, {
			name: "when metadata object is empty in input, should return empty metadata object",
			input: &awss3bucketv1.AwsS3Bucket{
				Metadata: &shared.ApiResourceMetadata{},
			},
			want: &shared.ApiResourceMetadata{},
		},
		{
			name:  "when metadata object is the input, nil should be returned",
			input: &shared.ApiResourceLifecycleAndAuditStatus{},
			want:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExtractMetadata(tt.input)
			if !proto.Equal(got, tt.want) {
				t.Errorf("Extractmetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}
