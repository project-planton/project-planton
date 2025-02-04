variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id = optional(string),
    org = optional(string),
    env = optional(object({
      name = optional(string),
      id = optional(string),
    })),
    labels = optional(object({
      key = string, value = string
    })),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
  })
}


variable "spec" {
  description = "spec"
  type = object({

    # aws lambda function spec
    function = object({

      # Instruction set architecture for your Lambda function. Valid values are `["x8664"]` and `["arm64"]`.
      # Default is `["x8664"]`. Removing this attribute, function's architecture stay the same.
      architectures = list(string)

      # Description of what your Lambda Function does.
      description = string

      # Configuration block. Detailed below.
      file_system_config = object({

        # Amazon Resource Name (ARN) of the Amazon EFS Access Point that provides access to the file system.
        arn = string

        # Path where the function can access the file system, starting with /mnt/.
        local_mount_path = string
      })

      # Function [entrypoint](https://docs.aws.amazon.com/lambda/latest/dg/walkthrough-custom-events-create-test-function.html) in your code.
      handler = string

      # ECR image URI containing the function's deployment package. Exactly one of `filename`, `imageUri`,  or `s3Bucket` must be specified.
      image_uri = string

      # Amazon Resource Name (ARN) of the AWS Key Management Service (KMS) key that is used to encrypt environment variables.
      # If this configuration is not provided when environment variables are in use, AWS Lambda uses a default service key.
      # If this configuration is provided when environment variables are not in use,
      # the AWS Lambda API does not save this configuration and the provider will show a perpetual difference of adding
      # the key. To fix the perpetual difference, remove this configuration.
      kms_key_arn = string

      # List of Lambda Layer Version ARNs (maximum of 5) to attach to your Lambda Function.
      # See [Lambda Layers](https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
      layers = list(string)

      # Amount of memory in MB your Lambda Function can use at runtime. Defaults to `128`.
      # See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html)
      memory_size = number

      # Lambda deployment package type. Valid values are `Zip` and `Image`. Defaults to `Zip`.
      package_type = string

      # Whether to publish creation/change as new Lambda Function Version. Defaults to `false`.
      publish = bool

      # Amount of reserved concurrent executions for this lambda function. A value of `0` disables lambda from
      # being triggered and `-1` removes any concurrency limitations. Defaults to Unreserved Concurrency Limits `-1`.
      # See [Managing Concurrency](https://docs.aws.amazon.com/lambda/latest/dg/concurrent-executions.html)
      reserved_concurrent_executions = number

      # Identifier of the function's runtime.
      # See [Runtimes](https://docs.aws.amazon.com/lambda/latest/dg/API_CreateFunction.html#SSS-CreateFunction-request-Runtime) for valid values.
      runtime = string

      # S3 bucket location containing the function's deployment package. This bucket must reside in the same AWS region
      # where you are creating the Lambda function. Exactly one of `filename`, `imageUri`, or `s3Bucket` must be specified.
      # When `s3Bucket` is set, `s3Key` is required.
      s3_bucket = string

      # S3 key of an object containing the function's deployment package. When `s3Bucket` is set, `s3Key` is required.
      s3_key = string

      # Object version containing the function's deployment package. Conflicts with `filename` and `imageUri`.
      s3_object_version = string

      # Used to trigger updates. Must be set to a base64-encoded SHA256 hash of the package file specified with either
      # filename or s3_key. The usual way to set this is filebase64sha256('file.zip') where 'file.zip' is the local filename
      # of the lambda function source archive.
      source_code_hash = string

      # Amount of time your Lambda Function has to run in seconds. Defaults to `3`.
      # See [Limits](https://docs.aws.amazon.com/lambda/latest/dg/limits.html).
      timeout = number

      # Map of environment variables that are accessible from the function code during execution. If provided at least one key must be present.
      variables = object({

        # Description for key
        key = string

        # Description for value
        value = string
      })

      # ARN of an SNS topic or SQS queue to notify when an invocation fails. If this option is used, the function's IAM
      # role must be granted suitable access to write to the target object, which means allowing either
      # the `sns:Publish` or `sqs:SendMessage` action on this ARN, depending on which service is targeted.
      dead_letter_config_target_arn = string

      # The Lambda OCI [image configurations](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_function#image_config)
      image_config = object({

        # Parameters that you want to pass in with `entryPoint`.
        commands = list(string)

        # Entry point to your application, which is typically the location of the runtime executable.
        entry_points = list(string)

        # Working directory.
        working_directory = string
      })

      # Whether to sample and trace a subset of incoming requests with AWS X-Ray. Valid values are `PassThrough` and `Active`.
      # If `PassThrough`, Lambda will only trace the request from an upstream service if it contains a tracing header
      # with "sampled=1". If `Active`, Lambda will respect any tracing header it receives from an upstream service.
      # If no tracing header is received, Lambda will call X-Ray for a tracing decision.
      tracing_config_mode = string

      # VPC configuration
      vpc_config = object({

        # List of security group IDs associated with the Lambda function.
        security_group_ids = list(string)

        # List of subnet IDs associated with the Lambda function.
        subnet_ids = list(string)

        # ID of the VPC.
        vpc_id = string
      })

      # The size of the Lambda function Ephemeral storage(`/tmp`) represented in MB.
      # The minimum supported `ephemeralStorage` value defaults to `512`MB and the maximum supported value is `10240`MB.
      ephemeral_storage_size = number
    })

    # aws lambda function iam spec
    iam_role = object({

      # ARN of the policy that is used to set the permissions boundary for the role
      permissions_boundary = string

      # Enable Lambda@Edge for your Node.js or Python functions. The required trust relationship and publishing of
      # function versions will be configured in this module.
      lambda_at_edge_enabled = bool

      # Enable CloudWatch Lambda Insights for the Lambda Function.
      cloudwatch_lambda_insights_enabled = bool

      # List of AWS Systems Manager Parameter Store parameter names. The IAM role of this Lambda function will be enhanced
      # with read permissions for those parameters. Parameters must start with a forward slash and can be encrypted with the
      # default KMS key.
      ssm_parameter_names = list(string)

      # ARNs of custom policies to be attached to the lambda role
      custom_iam_policy_arns = list(string)

      # Inline policy document (JSON) to attach to the lambda role
      inline_iam_policy = string
    })

    # aws lambda cloud watch log group
    cloudwatch_log_group = object({

      # The ARN of the KMS Key to use when encrypting log data.
      # Please note, after the AWS KMS CMK is disassociated from the log group, AWS CloudWatch Logs stops encrypting newly
      # ingested data for the log group.
      # All previously ingested data remains encrypted, and AWS CloudWatch Logs requires permissions for the CMK whenever
      # the encrypted data is requested.
      kms_key_arn = string

      # Number of days you want to retain log events in the log group
      retention_in_days = number
    })

    # Defines which external source(s) can invoke this function (action 'lambda:InvokeFunction'). Attributes map to
    # those of https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/lambda_permission.
    # NOTE: to keep things simple, we only expose a subset of said attributes. If a more complex configuration is
    # needed, declare the necessary lambda permissions outside of this module
    invoke_function_permissions = list(object({

      # The principal who is getting this permission e.g., `s3.amazonaws.com`,
      # an AWS account ID, or AWS IAM principal, or AWS service principal such as `events.amazonaws.com` or `sns.amazonaws.com`.
      principal = string

      # When the principal is an AWS service, the ARN of the specific resource within that service to grant permission to.
      # Without this, any resource from `principal` will be granted permission – even if that resource is from another account.
      # For S3, this should be the ARN of the S3 Bucket.
      # For EventBridge events, this should be the ARN of the EventBridge Rule.
      # For API Gateway, this should be the ARN of the API, as described
      # [here](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-control-access-using-iam-policies-to-invoke-api.html).
      source_arn = string
    }))
  })
}