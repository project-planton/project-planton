package convertstringmaps

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ConvertGoStringMapToPulumiStringMap converts a Golang string map to a Pulumi String map.
func ConvertGoStringMapToPulumiStringMap(goStringMap map[string]string) pulumi.StringMap {
	pulumiStringMap := make(pulumi.StringMap)
	for k, v := range goStringMap {
		pulumiStringMap[k] = pulumi.String(v)
	}
	return pulumiStringMap
}

// ConvertGoStringMapToPulumiMap converts a Golang string map to a Pulumi map.
func ConvertGoStringMapToPulumiMap(goStringMap map[string]string) pulumi.Map {
	pulumiMap := make(pulumi.Map)
	for k, v := range goStringMap {
		pulumiMap[k] = pulumi.String(v)
	}
	return pulumiMap
}
