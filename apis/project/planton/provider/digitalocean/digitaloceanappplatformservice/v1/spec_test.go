package digitaloceanappplatformservicev1

import (
	"testing"

	foreignkeyv1 "github.com/project-planton/project-planton/apis/project/planton/shared/foreignkey/v1"

	"buf.build/go/protovalidate"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/project-planton/project-planton/apis/project/planton/shared"
	"github.com/project-planton/project-planton/apis/project/planton/provider/digitalocean"
)

func TestDigitalOceanAppPlatformServiceSpec(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DigitalOceanAppPlatformServiceSpec Custom Validation Tests")
}

var _ = Describe("DigitalOceanAppPlatformServiceSpec Custom Validation Tests", func() {

	Describe("When valid input is passed", func() {
		Context("digitalocean_app_platform_service", func() {

			It("should not return a validation error for minimal valid fields with git source", func() {
				input := &DigitalOceanAppPlatformService{
					ApiVersion: "digital-ocean.project-planton.org/v1",
					Kind:       "DigitalOceanAppPlatformService",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-app-service",
					},
					Spec: &DigitalOceanAppPlatformServiceSpec{
						ServiceName:      "test-service",
						Region:           digitalocean.DigitalOceanRegion_nyc3,
						ServiceType:      DigitalOceanAppPlatformServiceType_web_service,
						InstanceSizeSlug: DigitalOceanAppPlatformInstanceSize_basic_xxs,
						Source: &DigitalOceanAppPlatformServiceSpec_GitSource{
							GitSource: &DigitalOceanAppPlatformGitSource{
								RepoUrl: "https://github.com/example/repo.git",
								Branch:  "main",
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should not return a validation error for minimal valid fields with image source", func() {
				input := &DigitalOceanAppPlatformService{
					ApiVersion: "digital-ocean.project-planton.org/v1",
					Kind:       "DigitalOceanAppPlatformService",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-app-service",
					},
					Spec: &DigitalOceanAppPlatformServiceSpec{
						ServiceName:      "test-service",
						Region:           digitalocean.DigitalOceanRegion_nyc3,
						ServiceType:      DigitalOceanAppPlatformServiceType_web_service,
						InstanceSizeSlug: DigitalOceanAppPlatformInstanceSize_basic_xxs,
						Source: &DigitalOceanAppPlatformServiceSpec_ImageSource{
							ImageSource: &DigitalOceanAppPlatformRegistrySource{
								Registry: &foreignkeyv1.StringValueOrRef{
									LiteralOrRef: &foreignkeyv1.StringValueOrRef_Value{Value: "registry.digitalocean.com/myregistry"},
								},
								Repository: "myapp",
								Tag:        "latest",
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})

			It("should return a validation error when no source is provided", func() {
				input := &DigitalOceanAppPlatformService{
					ApiVersion: "digital-ocean.project-planton.org/v1",
					Kind:       "DigitalOceanAppPlatformService",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-app-service",
					},
					Spec: &DigitalOceanAppPlatformServiceSpec{
						ServiceName:      "test-service",
						Region:           digitalocean.DigitalOceanRegion_nyc3,
						ServiceType:      DigitalOceanAppPlatformServiceType_web_service,
						InstanceSizeSlug: DigitalOceanAppPlatformInstanceSize_basic_xxs,
						// No source provided - should fail validation
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("Either git_source or image_source must be specified"))
			})

			It("should return a validation error when autoscale is enabled but min/max not set", func() {
				input := &DigitalOceanAppPlatformService{
					ApiVersion: "digital-ocean.project-planton.org/v1",
					Kind:       "DigitalOceanAppPlatformService",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-app-service",
					},
					Spec: &DigitalOceanAppPlatformServiceSpec{
						ServiceName:      "test-service",
						Region:           digitalocean.DigitalOceanRegion_nyc3,
						ServiceType:      DigitalOceanAppPlatformServiceType_web_service,
						InstanceSizeSlug: DigitalOceanAppPlatformInstanceSize_basic_xxs,
						EnableAutoscale:  true,
						// Missing min_instance_count and max_instance_count
						Source: &DigitalOceanAppPlatformServiceSpec_GitSource{
							GitSource: &DigitalOceanAppPlatformGitSource{
								RepoUrl: "https://github.com/example/repo.git",
								Branch:  "main",
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).NotTo(BeNil())
				Expect(err.Error()).To(ContainSubstring("min_instance_count and max_instance_count must be set"))
			})

			It("should not return a validation error when autoscale is properly configured", func() {
				input := &DigitalOceanAppPlatformService{
					ApiVersion: "digital-ocean.project-planton.org/v1",
					Kind:       "DigitalOceanAppPlatformService",
					Metadata: &shared.ApiResourceMetadata{
						Name: "test-app-service",
					},
					Spec: &DigitalOceanAppPlatformServiceSpec{
						ServiceName:      "test-service",
						Region:           digitalocean.DigitalOceanRegion_nyc3,
						ServiceType:      DigitalOceanAppPlatformServiceType_web_service,
						InstanceSizeSlug: DigitalOceanAppPlatformInstanceSize_basic_xxs,
						EnableAutoscale:  true,
						MinInstanceCount: 2,
						MaxInstanceCount: 5,
						Source: &DigitalOceanAppPlatformServiceSpec_GitSource{
							GitSource: &DigitalOceanAppPlatformGitSource{
								RepoUrl: "https://github.com/example/repo.git",
								Branch:  "main",
							},
						},
					},
				}
				err := protovalidate.Validate(input)
				Expect(err).To(BeNil())
			})
		})
	})
})
