package module

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ec2"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/ecs"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lb"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	awsecsservicev1 "github.com/project-planton/project-planton/apis/project/planton/provider/aws/awsecsservice/v1"
)

// service creates and wires up the ECS Task Definition and AWS ECS Service resources.
func ecsService(ctx *pulumi.Context, locals *Locals, provider *aws.Provider) error {
	spec := locals.AwsEcsService.Spec
	serviceName := locals.AwsEcsService.Metadata.Name

	containerDefs, err := buildContainerDefinitions(
		serviceName,
		spec.Container.Image.Repo,
		spec.Container.Image.Tag,
		spec.Container.Port,
		spec.Container.Env, // new param
	)
	if err != nil {
		return errors.Wrap(err, "failed to build container definitions JSON")
	}

	taskDefinitionArgs := &ecs.TaskDefinitionArgs{
		Family:                  pulumi.String(serviceName),
		RequiresCompatibilities: pulumi.StringArray{pulumi.String("FARGATE")},
		Cpu:                     pulumi.String(fmt.Sprintf("%d", spec.Container.Cpu)),
		Memory:                  pulumi.String(fmt.Sprintf("%d", spec.Container.Memory)),
		NetworkMode:             pulumi.String("awsvpc"),
		ContainerDefinitions:    pulumi.String(containerDefs),
		Tags:                    pulumi.ToStringMap(locals.AwsTags),
	}

	if spec.Iam != nil {
		if spec.Iam.TaskExecutionRoleArn != "" {
			taskDefinitionArgs.ExecutionRoleArn = pulumi.String(spec.Iam.TaskExecutionRoleArn)
		}
		if spec.Iam.TaskRoleArn != "" {
			taskDefinitionArgs.TaskRoleArn = pulumi.String(spec.Iam.TaskRoleArn)
		}
	}

	taskDef, err := ecs.NewTaskDefinition(ctx,
		serviceName+"-taskdef",
		taskDefinitionArgs,
		pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "unable to create ECS task definition")
	}

	serviceArgs := &ecs.ServiceArgs{
		Name:           pulumi.String(serviceName),
		Cluster:        pulumi.String(spec.ClusterArn),
		LaunchType:     pulumi.String("FARGATE"),
		DesiredCount:   pulumi.Int(int(spec.Container.Replicas)),
		TaskDefinition: taskDef.Arn,
		NetworkConfiguration: &ecs.ServiceNetworkConfigurationArgs{
			Subnets:        toPulumiStrings(spec.Network.Subnets),
			SecurityGroups: toPulumiStrings(spec.Network.SecurityGroups),
		},
		Tags: pulumi.ToStringMap(locals.AwsTags),
	}

	var loadBalancerDNS pulumi.StringInput = pulumi.String("")

	if spec.Alb.Enabled && spec.Container.Port != 0 {
		if len(spec.Network.Subnets) == 0 {
			return errors.New("at least one subnet is required for ALB usage")
		}
		if spec.Alb.Arn == "" {
			return errors.New("alb.arn is required when alb.enabled = true")
		}

		firstSubnetID := spec.Network.Subnets[0]
		subnetLookup, err := ec2.LookupSubnet(ctx, &ec2.LookupSubnetArgs{
			Id: &firstSubnetID,
		}, pulumi.Provider(provider))
		if err != nil {
			return errors.Wrap(err, "failed to lookup subnet for ALB target group")
		}

		targetGroup, err := lb.NewTargetGroup(ctx, serviceName+"-tg", &lb.TargetGroupArgs{
			Port:        pulumi.Int(int(spec.Container.Port)),
			Protocol:    pulumi.String("HTTP"),
			TargetType:  pulumi.String("ip"),
			VpcId:       pulumi.String(subnetLookup.VpcId),
			HealthCheck: &lb.TargetGroupHealthCheckArgs{Path: pulumi.String("/")},
			Tags:        pulumi.ToStringMap(locals.AwsTags),
		}, pulumi.Provider(provider))
		if err != nil {
			return errors.Wrap(err, "failed to create ALB target group")
		}

		serviceArgs.LoadBalancers = ecs.ServiceLoadBalancerArray{
			&ecs.ServiceLoadBalancerArgs{
				TargetGroupArn: targetGroup.Arn,
				ContainerName:  pulumi.String(serviceName),
				ContainerPort:  pulumi.Int(int(spec.Container.Port)),
			},
		}

		foundAlb, err := lb.LookupLoadBalancer(ctx, &lb.LookupLoadBalancerArgs{
			Arn: &spec.Alb.Arn,
		}, pulumi.Provider(provider))
		if err != nil {
			return errors.Wrap(err, "failed to find ALB by ARN")
		}
		loadBalancerDNS = pulumi.String(foundAlb.DnsName)

		// Convert user-supplied int32 -> int to match the lookup function's signature
		listenerPort := int(spec.Alb.ListenerPort)

		foundListener, err := lb.LookupListener(ctx, &lb.LookupListenerArgs{
			LoadBalancerArn: &foundAlb.Arn,
			Port:            &listenerPort,
		}, pulumi.Provider(provider))
		if err != nil {
			return errors.Wrap(err, "failed to find ALB listener on the given port")
		}

		conditions := lb.ListenerRuleConditionArray{}

		if spec.Alb.RoutingType == "path" {
			if spec.Alb.Path == "" {
				return errors.New("alb.path must be set if routingType is 'path'")
			}
			conditions = lb.ListenerRuleConditionArray{
				&lb.ListenerRuleConditionArgs{
					PathPattern: &lb.ListenerRuleConditionPathPatternArgs{
						Values: pulumi.StringArray{
							pulumi.String(spec.Alb.Path),
						},
					},
				},
			}
		}

		if spec.Alb.RoutingType == "hostname" {
			if spec.Alb.Hostname == "" {
				return errors.New("alb.hostname must be set if routingType is 'hostname'")
			}
			conditions = lb.ListenerRuleConditionArray{
				&lb.ListenerRuleConditionArgs{
					HostHeader: &lb.ListenerRuleConditionHostHeaderArgs{
						Values: pulumi.StringArray{
							pulumi.String(spec.Alb.Hostname),
						},
					},
				},
			}
		}

		if len(conditions) > 0 {
			_, err := lb.NewListenerRule(ctx, serviceName+"-rule", &lb.ListenerRuleArgs{
				ListenerArn: pulumi.String(foundListener.Arn),
				Actions: lb.ListenerRuleActionArray{
					&lb.ListenerRuleActionArgs{
						Type:           pulumi.String("forward"),
						TargetGroupArn: targetGroup.Arn,
					},
				},
				Conditions: conditions,
				Priority:   pulumi.Int(100),
				Tags:       pulumi.ToStringMap(locals.AwsTags),
			}, pulumi.Provider(provider))
			if err != nil {
				return errors.Wrap(err, "failed to create listener rule for path/hostname-based routing")
			}
		}
	}

	awsEcsService, err := ecs.NewService(ctx, serviceName+"-service", serviceArgs, pulumi.Provider(provider))
	if err != nil {
		return errors.Wrap(err, "unable to create ECS service")
	}

	ctx.Export(OpAwsEcsServiceName, awsEcsService.Name)
	ctx.Export(OpEcsClusterName, pulumi.String(spec.ClusterArn))
	ctx.Export(OpLoadBalancerDnsName, loadBalancerDNS)

	var serviceUrl pulumi.StringInput = pulumi.String("")
	if spec.Alb.RoutingType == "HOSTNAME" &&
		spec.Alb.Enabled && spec.Alb.Hostname != "" {
		serviceUrl = pulumi.String(fmt.Sprintf("http://%s", spec.Alb.Hostname))
	}
	ctx.Export(OpServiceUrl, serviceUrl)
	ctx.Export(OpServiceDiscoveryName, pulumi.String(""))

	return nil
}

// buildContainerDefinitions constructs a JSON array of container definitions.
// It now honours env.variables, env.secrets, and env.files (S3 env-file objects).
func buildContainerDefinitions(
	serviceName, repo, tag string,
	port int32,
	env *awsecsservicev1.AwsEcsServiceContainerEnv,
) (string, error) {

	// -------- environment (key-value) ----------
	envVars := []map[string]string{}
	if env != nil {
		// deterministic order for easier diffs
		keys := []string{}
		for k := range env.Variables {
			keys = append(keys, k)
		}
		for k := range env.Secrets {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			val := ""
			if v, ok := env.Variables[k]; ok {
				val = v
			}
			if v, ok := env.Secrets[k]; ok {
				val = v
			}
			envVars = append(envVars, map[string]string{
				"name":  k,
				"value": val,
			})
		}
	}

	// -------- environmentFiles (S3) ------------
	envFiles := []map[string]string{}
	if env != nil {
		for _, f := range env.Files {
			if f.S3Uri != "" {
				envFiles = append(envFiles, map[string]string{
					"type":  "s3",
					"value": f.S3Uri,
				})
			}
		}
	}

	// -------- container base -------------------
	container := map[string]interface{}{
		"name":      serviceName,
		"image":     fmt.Sprintf("%s:%s", repo, tag),
		"essential": true,
	}

	if len(envVars) > 0 {
		container["environment"] = envVars
	}
	if len(envFiles) > 0 {
		container["environmentFiles"] = envFiles
	}

	if port != 0 {
		container["portMappings"] = []map[string]int32{
			{
				"containerPort": port,
			},
		}
	}

	containerDefinitions := []interface{}{container}
	encoded, err := json.Marshal(containerDefinitions)
	if err != nil {
		return "", err
	}
	return string(encoded), nil
}

// toPulumiStrings is a helper that converts a native string slice to a pulumi.StringArray.
func toPulumiStrings(input []string) pulumi.StringArray {
	output := make(pulumi.StringArray, len(input))
	for i, s := range input {
		output[i] = pulumi.String(s)
	}
	return output
}
