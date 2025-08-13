package awscloudfrontv1

import (
    "testing"

    "github.com/bufbuild/protovalidate-go"
    . "github.com/onsi/ginkgo/v2"
    . "github.com/onsi/gomega"

    "github.com/project-planton/project-planton/apis/project/planton/shared"
)

func TestAwsCloudFront(t *testing.T) {
    RegisterFailHandler(Fail)
    RunSpecs(t, "AwsCloudFront Suite")
}

var _ = Describe("AwsCloudFront Validation", func() {
    It("accepts a minimal valid resource", func() {
        input := &AwsCloudFront{
            ApiVersion: "aws.project-planton.org/v1",
            Kind:       "AwsCloudFront",
            Metadata: &shared.ApiResourceMetadata{ Name: "cf-basic" },
            Spec: &AwsCloudFrontSpec{
                DefaultCacheBehavior: &DefaultCacheBehavior{
                    OriginId:             "origin-1",
                    ViewerProtocolPolicy: ViewerProtocolPolicy_REDIRECT_TO_HTTPS,
                },
                Origins: []*Origin{{ Id: "origin-1", DomainName: "example.s3.amazonaws.com" }},
            },
        }
        err := protovalidate.Validate(input)
        Expect(err).To(BeNil())
    })
})


