package module

import (
	"fmt"
	"net"
	"sort"
	"strconv"

	"github.com/pkg/errors"
	"github.com/project-planton/project-planton/apis/org/project_planton/shared/cloudresourcekind"

	awsvpcv1 "github.com/project-planton/project-planton/apis/org/project_planton/provider/aws/awsvpc/v1"
	"github.com/project-planton/project-planton/pkg/iac/pulumi/pulumimodule/provider/aws/awstagkeys"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubnetName string
type SubnetCidr string
type AvailabilityZone string

type Locals struct {
	AwsVpc             *awsvpcv1.AwsVpc
	AwsTags            map[string]string
	PrivateAzSubnetMap map[AvailabilityZone]map[SubnetName]SubnetCidr
	PublicAzSubnetMap  map[AvailabilityZone]map[SubnetName]SubnetCidr
}

func initializeLocals(ctx *pulumi.Context, stackInput *awsvpcv1.AwsVpcStackInput) (*Locals, error) {
	locals := &Locals{}
	locals.AwsVpc = stackInput.Target

	_, vpcCidrBlock, err := net.ParseCIDR(locals.AwsVpc.Spec.VpcCidr)
	if err != nil {
		return nil, errors.Wrapf(err, "invalid vpc cidr %s", locals.AwsVpc.Spec.VpcCidr)
	}
	vpcMaskSize, _ := vpcCidrBlock.Mask.Size()
	if vpcMaskSize > int(locals.AwsVpc.Spec.SubnetSize) {
		return nil, errors.Errorf("spec.subnetSize /%d cannot be bigger than the VPC /%d",
			locals.AwsVpc.Spec.SubnetSize, vpcMaskSize)
	}

	locals.AwsTags = map[string]string{
		awstagkeys.Resource:     strconv.FormatBool(true),
		awstagkeys.Organization: locals.AwsVpc.Metadata.Org,
		awstagkeys.Environment:  locals.AwsVpc.Metadata.Env,
		awstagkeys.ResourceKind: cloudresourcekind.CloudResourceKind_AwsVpc.String(),
		awstagkeys.ResourceId:   locals.AwsVpc.Metadata.Id,
	}

	locals.PrivateAzSubnetMap = GetPrivateAzSubnetMap(locals.AwsVpc)
	locals.PublicAzSubnetMap = getPublicAzSubnetMap(locals.AwsVpc)

	return locals, nil
}

func GetPrivateAzSubnetMap(awsVpc *awsvpcv1.AwsVpc) map[AvailabilityZone]map[SubnetName]SubnetCidr {
	privateAzSubnetMap := make(map[AvailabilityZone]map[SubnetName]SubnetCidr, 0)

	for azIndex, az := range awsVpc.Spec.AvailabilityZones {
		for subnetIndex := 0; subnetIndex < int(awsVpc.Spec.SubnetsPerAvailabilityZone); subnetIndex++ {
			privateSubnetName := fmt.Sprintf("private-subnet-%s-%d", az, subnetIndex)
			privateSubnetCidr := fmt.Sprintf("10.0.%d.0/%d", 100+azIndex*10+subnetIndex, awsVpc.Spec.SubnetSize)

			if privateAzSubnetMap[AvailabilityZone(az)] == nil {
				privateAzSubnetMap[AvailabilityZone(az)] = make(map[SubnetName]SubnetCidr)
			}

			privateAzSubnetMap[AvailabilityZone(az)][SubnetName(privateSubnetName)] = SubnetCidr(privateSubnetCidr)
		}
	}
	return privateAzSubnetMap
}

func getPublicAzSubnetMap(awsVpc *awsvpcv1.AwsVpc) map[AvailabilityZone]map[SubnetName]SubnetCidr {
	publicAzSubnetMap := make(map[AvailabilityZone]map[SubnetName]SubnetCidr, 0)

	for azIndex, az := range awsVpc.Spec.AvailabilityZones {
		for subnetIndex := 0; subnetIndex < int(awsVpc.Spec.SubnetsPerAvailabilityZone); subnetIndex++ {
			publicSubnetName := fmt.Sprintf("public-subnet-%s-%d", az, subnetIndex)
			publicSubnetCidr := fmt.Sprintf("10.0.%d.0/%d", azIndex*10+subnetIndex, awsVpc.Spec.SubnetSize)

			if publicAzSubnetMap[AvailabilityZone(az)] == nil {
				publicAzSubnetMap[AvailabilityZone(az)] = make(map[SubnetName]SubnetCidr)
			}
			publicAzSubnetMap[AvailabilityZone(az)][SubnetName(publicSubnetName)] = SubnetCidr(publicSubnetCidr)
		}
	}
	return publicAzSubnetMap
}

func getSortedAzKeys(azSubnetMap map[AvailabilityZone]map[SubnetName]SubnetCidr) []string {
	keys := make([]string, 0, len(azSubnetMap))
	for k := range azSubnetMap {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	return keys
}

func getSortedSubnetNameKeys(subnetMap map[SubnetName]SubnetCidr) []string {
	keys := make([]string, 0, len(subnetMap))
	for k := range subnetMap {
		keys = append(keys, string(k))
	}
	sort.Strings(keys)
	return keys
}
