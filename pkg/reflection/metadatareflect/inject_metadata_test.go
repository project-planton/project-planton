package metadatareflect

import (
	awss3bucketv1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awss3bucket/v1"
	"testing"

	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"google.golang.org/protobuf/proto"
)

func TestInjectMetadata(t *testing.T) {
	tests := []struct {
		name  string
		input proto.Message
		meta  *shared.ApiResourceMetadata
		want  proto.Message
	}{
		{
			name:  "when metadata is injected it should appear in the output",
			input: &awss3bucketv1.AwsS3Bucket{},
			meta:  &shared.ApiResourceMetadata{Id: "test-id"},
			want: &awss3bucketv1.AwsS3Bucket{
				Metadata: &shared.ApiResourceMetadata{Id: "test-id"},
			},
		},
		{
			name:  "when meta is nil the message must stay unchanged",
			input: &awss3bucketv1.AwsS3Bucket{},
			meta:  nil,
			want:  &awss3bucketv1.AwsS3Bucket{},
		},
		{
			name:  "when message has no metadata field the call is a no‑op",
			input: &shared.ApiResourceLifecycleAndAuditStatus{},
			meta:  &shared.ApiResourceMetadata{Id: "test-id"},
			want:  &shared.ApiResourceLifecycleAndAuditStatus{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := InjectMetadata(proto.Clone(tt.input), tt.meta) // clone to avoid mutating test data
			if !proto.Equal(got, tt.want) {
				t.Errorf("InjectMetadata() = %v, want %v", got, tt.want)
			}
		})
	}
}
