# Upstream's text blocks (cmd/ate-setup/internal/steps/overlay.go), byte for
# byte the Pulumi module's splice.go: the additional ext_proc filter and its
# cluster for Envoy's egress configuration, and the credential-provider block
# for agentgateway's egress policy. documents.tf fills each in at upstream's
# marker, indented like the marker line.

locals {
  extproc_filter = trimsuffix(<<-EOT
    - name: envoy.filters.http.ext_proc
      typed_config:
        "@type": type.googleapis.com/envoy.extensions.filters.http.ext_proc.v3.ExternalProcessor
        grpc_service:
          envoy_grpc:
            cluster_name: additional_egress_ext_proc
          timeout: 2s
        failure_mode_allow: false
        message_timeout: 2s
        request_attributes:
        - filter_state['dev.ate.actor.identity']
        processing_mode:
          request_header_mode: SEND
          response_header_mode: SKIP
          request_body_mode: NONE
          response_body_mode: NONE
          request_trailer_mode: SKIP
          response_trailer_mode: SKIP
        mutation_rules:
          disallow_system: true
          disallow_is_error: true
    EOT
  , "\n")

  # sni, SAN, address, port.
  extproc_cluster = trimsuffix(<<-EOT
    - name: additional_egress_ext_proc
      type: STRICT_DNS
      lb_policy: ROUND_ROBIN
      connect_timeout: 1s
      typed_extension_protocol_options:
        envoy.extensions.upstreams.http.v3.HttpProtocolOptions:
          "@type": type.googleapis.com/envoy.extensions.upstreams.http.v3.HttpProtocolOptions
          explicit_http_config:
            http2_protocol_options: {}
      transport_socket:
        name: envoy.transport_sockets.tls
        typed_config:
          "@type": type.googleapis.com/envoy.extensions.transport_sockets.tls.v3.UpstreamTlsContext
          sni: %s
          common_tls_context:
            tls_params:
              tls_minimum_protocol_version: TLSv1_3
              tls_maximum_protocol_version: TLSv1_3
            tls_certificate_sds_secret_configs:
            - name: podidentity_client_cert
              sds_config:
                resource_api_version: V3
                path_config_source:
                  path: /etc/envoy/sds-podidentity-cert.yaml
            combined_validation_context:
              default_validation_context:
                match_typed_subject_alt_names:
                - san_type: DNS
                  matcher:
                    exact: %s
              validation_context_sds_secret_config:
                name: servicedns_validation_context
                sds_config:
                  resource_api_version: V3
                  path_config_source:
                    path: /etc/envoy/sds-servicedns-validation.yaml
      load_assignment:
        cluster_name: additional_egress_ext_proc
        endpoints:
        - lb_endpoints:
          - endpoint:
              address:
                socket_address:
                  address: %s
                  port_value: %d
    EOT
  , "\n")

  # name, address, server name, each quoted.
  agentgateway_providers = trimsuffix(<<-EOT
    credentialProviders:
    - uriAuthority: %s
      target:
        host: %s
        policies:
          backendTLS:
            hostname: %s
            cert: /run/podidentity.podcert.ate.dev/credential-bundle.pem
            key: /run/podidentity.podcert.ate.dev/credential-bundle.pem
            root: /run/servicedns-ca/trust-bundle.pem
    EOT
  , "\n")
}
