package module

import (
	"github.com/pkg/errors"
	civo "github.com/pulumi/pulumi-civo/sdk/v2/go/civo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// volume provisions the Civo Block Storage Volume itself and exports its ID.
func volume(
	ctx *pulumi.Context,
	locals *Locals,
	civoProvider *civo.Provider,
) (*civo.Volume, error) {

	// 1 . Build the resource arguments straight from the proto fields.
	volumeArgs := &civo.VolumeArgs{
		Name:   pulumi.String(locals.CivoVolume.Spec.VolumeName),
		Region: pulumi.String(locals.CivoVolume.Spec.Region.String()),
		SizeGb: pulumi.Int(int(locals.CivoVolume.Spec.SizeGib)),
	}

	// 2 . Create the Volume.
	createdVolume, err := civo.NewVolume(
		ctx,
		"volume",
		volumeArgs,
		pulumi.Provider(civoProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create civo volume")
	}

	// 3 . Export stack output.
	ctx.Export(OpVolumeId, createdVolume.ID())

	return createdVolume, nil
}
