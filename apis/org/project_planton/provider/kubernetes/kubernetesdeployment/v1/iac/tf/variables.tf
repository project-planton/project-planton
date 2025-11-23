variable "metadata" {
  description = "Metadata for the resource, including name and labels"
  type = object({
    name = string,
    id = optional(string),
    org = optional(string),
    env = optional(string),
    labels = optional(map(string)),
    tags = optional(list(string)),
    version = optional(object({ id = string, message = string }))
  })
}


variable "spec" {
  description = "spec"
  type = object({

    # The Kubernetes cluster to install this microservice on.
    target_cluster = object({
      # Name of the target Kubernetes cluster
      cluster_name = string
    })

    # Kubernetes namespace to install the microservice.
    namespace = string

    # The version of the microservice being deployed.
    # This is usually either "main" (the default git branch name) or "review-<id>" where <id> is the merge request number.
    # It must be between 1 and 20 characters and can only contain lowercase letters, numbers, and hyphens.
    version = string

    # The container specifications for the microservice deployment.
    # This includes configurations for the main application container and any sidecar containers.
    container = object({

      # The main application container specifications.
      app = object({

        # The container image to be used for the application.
        # This value is computed during creation but can be updated.
        # It is derived by combining the Docker repository of the artifact store configured for the environment and the code project path.
        # The `pull_secret_name` is the name of the image pull secret to be configured in the Kubernetes Deployment resource.
        # It is determined by looking up the `container_image_artifact_store_id` from the environment where the microservice is deployed.
        image = object({

          # The repository of the image (e.g., "gcr.io/project/image").
          repo = string

          # The tag of the image (e.g., "latest" or "1.0.0").
          tag = string

          # The name of the image pull secret for private image repositories.
          pull_secret_name = optional(string)
        })

        # The CPU and memory resources allocated to the application container.
        resources = object({

          # The resource limits for the container.
          # Specify the maximum amount of CPU and memory that the container can use.
          limits = object({

            # The amount of CPU allocated (e.g., "500m" for 0.5 CPU cores).
            cpu = string

            # The amount of memory allocated (e.g., "256Mi" for 256 mebibytes).
            memory = string
          })

          # The resource requests for the container.
          # Specify the minimum amount of CPU and memory that the container is guaranteed.
          requests = object({

            # The amount of CPU allocated (e.g., "500m" for 0.5 CPU cores).
            cpu = string

            # The amount of memory allocated (e.g., "256Mi" for 256 mebibytes).
            memory = string
          })
        })

        # The environment variables and secrets for the application container.
        env = object({
          # A map of environment variable names to their values.
          variables = optional(map(string))
          # A map of secret names to their values.
          secrets = optional(map(string))
        })

        # A list of ports to be configured for the application container.
        ports = list(object({

          # The name of the port (e.g., "http", "grpc").
          # The name must only contain lowercase alphanumeric characters and hyphens.
          # Port names must also start and end with an alphanumeric character.
          # For example, "123-abc" and "web" are valid, but "123_abc" and "-web" are not.
          name = string

          # The port number on the container.
          container_port = number

          # The network protocol used by the port (e.g., "TCP", "UDP", "SCTP").
          # Must be one of "TCP", "UDP", or "SCTP".
          network_protocol = string

          # The application protocol for the microservice (e.g., "http").
          # This field is used for setting up the name of the service port in Kubernetes.
          # It is used during microservice deployment and is relevant for deployment and stateful set pod managers.
          # Refer to: https://kubernetes.io/docs/concepts/services-networking/service/#application-protocol
          app_protocol = string

          # The port number on the Kubernetes service.
          service_port = number

          # A flag indicating whether this port should be exposed via ingress.
          is_ingress_port = bool
        }))
      })

      # A list of sidecar containers to be deployed alongside the main application container.
      sidecars = optional(list(object({

        # The name of the container.
        name = string

        # The container image to be used.
        image = string

        # A list of ports exposed by the container.
        ports = list(object({

          # The name of the port.
          name = string

          # The port number on the container.
          # **Note:** The attribute names must use camel case to marshal into the Kubernetes Container spec.
          container_port = number

          # The protocol used by the port (e.g., "TCP" or "UDP").
          protocol = string
        }))

        # Resource specifications for the container, including CPU and memory limits and requests.
        resources = object({

          # The resource limits for the container.
          # Specify the maximum amount of CPU and memory that the container can use.
          limits = object({

            # The amount of CPU allocated (e.g., "500m" for 0.5 CPU cores).
            cpu = string

            # The amount of memory allocated (e.g., "256Mi" for 256 mebibytes).
            memory = string
          })

          # The resource requests for the container.
          # Specify the minimum amount of CPU and memory that the container is guaranteed.
          requests = object({

            # The amount of CPU allocated (e.g., "500m" for 0.5 CPU cores).
            cpu = string

            # The amount of memory allocated (e.g., "256Mi" for 256 mebibytes).
            memory = string
          })
        })

        # A list of environment variables to be set in the container.
        env = list(object({

          # The name of the environment variable.
          name = string

          # The value of the environment variable.
          value = string
        }))
      })))
    })

    # The ingress configuration for the microservice.
    # This defines how the microservice can be accessed externally.
    ingress = object({

      # A flag to enable or disable ingress.
      is_enabled = bool

      # The dns domain.
      dns_domain = string
    })

    # The availability configuration for the microservice.
    # This includes settings for minimum replicas and autoscaling options.
    availability = optional(object({

      # The minimum number of pod replicas to maintain.
      min_replicas = number

      # The configuration for horizontal pod autoscaling.
      horizontal_pod_autoscaling = optional(object({

        # A flag to enable or disable horizontal pod autoscaling.
        is_enabled = bool

        # The target CPU utilization percentage to trigger autoscaling (e.g., 60.0).
        target_cpu_utilization_percent = number

        # The target memory utilization to trigger autoscaling (e.g., "1Gi").
        target_memory_utilization = string
      }))
    }))
  })
}
