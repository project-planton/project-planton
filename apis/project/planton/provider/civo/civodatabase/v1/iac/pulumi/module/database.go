package module

import (
	"github.com/pkg/errors"
	civodatabasev1 "github.com/project-planton/project-planton/apis/project/planton/provider/civo/civodatabase/v1"
	"github.com/pulumi/pulumi-civo/sdk/v2/go/civo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// database provisions the managed database instance and exports outputs.
func database(
	ctx *pulumi.Context,
	locals *Locals,
	civoProvider *civo.Provider,
) (*civo.Database, error) {

	// 1. Translate proto enum → engine slug accepted by Civo.
	var engineSlug string
	switch locals.CivoDatabase.Spec.Engine {
	case civodatabasev1.CivoDatabaseEngine_mysql:
		engineSlug = "mysql"
	case civodatabasev1.CivoDatabaseEngine_postgres:
		engineSlug = "postgres"
	default:
		return nil, errors.Errorf("unsupported database engine: %v", locals.CivoDatabase.Spec.Engine)
	}

	// 2. Build resource arguments from proto fields.
	databaseArgs := &civo.DatabaseArgs{
		Name:    pulumi.String(locals.CivoDatabase.Spec.DbInstanceName),
		Engine:  pulumi.String(engineSlug),
		Version: pulumi.String(locals.CivoDatabase.Spec.EngineVersion),
		Region:  pulumi.String(locals.CivoDatabase.Spec.Region.String()),
		Size:    pulumi.String(locals.CivoDatabase.Spec.SizeSlug),
		// Replica count in Civo = total nodes; primary + replicas.
		Nodes: pulumi.Int(int(locals.CivoDatabase.Spec.Replicas) + 1),
	}

	// Private network attachment
	if locals.CivoDatabase.Spec.NetworkId != nil && locals.CivoDatabase.Spec.NetworkId.GetValue() != "" {
		databaseArgs.NetworkId = pulumi.StringPtr(locals.CivoDatabase.Spec.NetworkId.GetValue())
	}

	// Optional: firewalls (first one only—Civo currently allows one).
	if len(locals.CivoDatabase.Spec.FirewallIds) > 0 &&
		locals.CivoDatabase.Spec.FirewallIds[0].GetValue() != "" {
		databaseArgs.FirewallId = pulumi.String(
			locals.CivoDatabase.Spec.FirewallIds[0].GetValue())
	}

	// 3. Provision the database.
	createdDatabase, err := civo.NewDatabase(
		ctx,
		"database",
		databaseArgs,
		pulumi.Provider(civoProvider),
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create Civo database")
	}

	// 4. Export stack outputs so ProjectPlanton can wire dependencies.
	ctx.Export(OpDatabaseId, createdDatabase.ID())
	ctx.Export(OpHost, createdDatabase.DnsEndpoint)
	ctx.Export(OpPort, createdDatabase.Port)
	ctx.Export(OpUsername, createdDatabase.Username)
	ctx.Export(OpPasswordSecretRef, createdDatabase.Password) // secret‑ref in future

	return createdDatabase, nil
}
