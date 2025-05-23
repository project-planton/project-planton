go_build_dir=pkg

#https://github.com/kubernetes-sigs/gateway-api/tree/main/config/crd/standard
#https://github.com/kubernetes-sigs/gateway-api/tree/v1.1.0/config/crd/standard
gateway_apis=v1.1.0

#https://github.com/istio/istio/tree/master/manifests/charts/base/crds
#https://github.com/istio/istio/tree/release-1.22
istio_release=release-1.22

#https://github.com/cert-manager/cert-manager/tree/master/deploy/crds
#https://github.com/cert-manager/cert-manager/tree/release-1.15
cert_manager_release=release-1.15

#https://github.com/apache/solr-operator/tree/v0.8.1/helm/solr-operator/crds
#https://github.com/apache/solr-operator/releases/tag/v0.8.1
solr_operator_release=v0.8.1

#https://github.com/strimzi/strimzi-kafka-operator/tree/main/install/cluster-operator
#https://github.com/strimzi/strimzi-kafka-operator/tree/release-0.42.x
strimzi_operator_release=release-0.42.x

#https://github.com/zalando/postgres-operator/tree/v1.12.2/manifests
#https://github.com/zalando/postgres-operator/tree/v1.12.2
zalando_operator_release=v1.12.2

#https://github.com/external-secrets/external-secrets/blob/v0.9.20/deploy/crds/bundle.yaml
#https://github.com/external-secrets/external-secrets/tree/v0.9.20
external_secrets_release=v0.9.20

#https://raw.githubusercontent.com/elastic/cloud-on-k8s/v2.14.0/deploy/eck-operator/charts/eck-operator-crds/templates/all-crds.yaml
elastic_operator_release=v2.14.0

#https://github.com/keycloak/keycloak-k8s-resources/tree/25.0.4/kubernetes
keycloak_release=25.0.4

.PHONY: clean
clean:
	rm -rf ${go_build_dir}

.PHONY: gen-gateway-apis
gen-gateway-apis:
	crd2pulumi --force --goPath=gatewayapis https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/${gateway_apis}/config/crd/standard/gateway.networking.k8s.io_gatewayclasses.yaml \
	https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/${gateway_apis}/config/crd/standard/gateway.networking.k8s.io_gateways.yaml \
	https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/${gateway_apis}/config/crd/standard/gateway.networking.k8s.io_grpcroutes.yaml \
	https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/${gateway_apis}/config/crd/standard/gateway.networking.k8s.io_httproutes.yaml \
	https://raw.githubusercontent.com/kubernetes-sigs/gateway-api/${gateway_apis}/config/crd/standard/gateway.networking.k8s.io_referencegrants.yaml

.PHONY: gen-istio
gen-istio:
	crd2pulumi --force --goPath=istio https://raw.githubusercontent.com/istio/istio/${istio_release}/manifests/charts/base/crds/crd-all.gen.yaml

.PHONY: gen-cert-manager
gen-cert-manager:
	crd2pulumi --force --goPath=certmanager https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-certificates.yaml \
		https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-certificaterequests.yaml \
		https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-challenges.yaml \
		https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-clusterissuers.yaml \
		https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-issuers.yaml \
		https://raw.githubusercontent.com/cert-manager/cert-manager/${cert_manager_release}/deploy/crds/crd-orders.yaml

.PHONY: gen-solr-operator
gen-solr-operator:
	crd2pulumi --force --goPath=solroperator https://raw.githubusercontent.com/apache/solr-operator/${solr_operator_release}/helm/solr-operator/crds/crds.yaml

.PHONY: gen-strimzi-operator
gen-strimzi-operator:
	crd2pulumi --force --goPath=strimzioperator https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/main/install/cluster-operator/040-Crd-kafka.yaml \
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/041-Crd-kafkaconnect.yaml \
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/042-Crd-strimzipodset.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/043-Crd-kafkatopic.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/044-Crd-kafkauser.yaml \
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/045-Crd-kafkamirrormaker.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/046-Crd-kafkabridge.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/047-Crd-kafkaconnector.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/048-Crd-kafkamirrormaker2.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/049-Crd-kafkarebalance.yaml\
		https://raw.githubusercontent.com/strimzi/strimzi-kafka-operator/${strimzi_operator_release}/install/cluster-operator/04A-Crd-kafkanodepool.yaml

.PHONY: gen-zalando-operator
gen-zalando-operator:
	crd2pulumi --force --goPath=zalandooperator https://raw.githubusercontent.com/zalando/postgres-operator/${zalando_operator_release}/manifests/operatorconfiguration.crd.yaml \
		https://raw.githubusercontent.com/zalando/postgres-operator/${zalando_operator_release}/manifests/postgresql.crd.yaml \
		https://raw.githubusercontent.com/zalando/postgres-operator/${zalando_operator_release}/manifests/postgresteam.crd.yaml

.PHONY: gen-external-secrets
gen-external-secrets:
	crd2pulumi --force --goPath=externalsecrets https://raw.githubusercontent.com/external-secrets/external-secrets/${external_secrets_release}/deploy/crds/bundle.yaml

.PHONY: gen-elastic-operator
gen-elastic-operator:
	crd2pulumi --force --goPath=elasticsearch https://raw.githubusercontent.com/elastic/cloud-on-k8s/${elastic_operator_release}/deploy/eck-operator/charts/eck-operator-crds/templates/all-crds.yaml

.PHONY: gen-keycloak
gen-keycloak:
	crd2pulumi --force --goPath=keycloak https://raw.githubusercontent.com/keycloak/keycloak-k8s-resources/${keycloak_release}/kubernetes/keycloakrealmimports.k8s.keycloak.org-v1.yml \
		https://raw.githubusercontent.com/keycloak/keycloak-k8s-resources/${keycloak_release}/kubernetes/keycloaks.k8s.keycloak.org-v1.yml

.PHONY: gen
gen: clean gen-gateway-apis gen-istio gen-cert-manager gen-solr-operator gen-strimzi-operator gen-zalando-operator gen-external-secrets gen-elastic-operator gen-keycloak

.PHONY: build
build: clean gen
