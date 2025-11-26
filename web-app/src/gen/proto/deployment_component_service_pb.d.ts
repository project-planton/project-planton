// TypeScript declarations for protobuf messages
// Generated to make Closure-style .pb.js files compatible with TypeScript

import type { DescMessage, DescFile } from "@bufbuild/protobuf";

// Create file descriptor
// Edition.EDITION_PROTO3 = 999 (for proto3 syntax)
export const mockFile: DescFile = {
  kind: "file",
  edition: 999, // EDITION_PROTO3
  name: "proto/deployment_component_service.proto",
  dependencies: [],
  enums: [],
  messages: [],
  extensions: [],
  services: [],
  members: [], // Make members iterable
  deprecated: false,
  proto: {} as any,
  toString: () => "proto/deployment_component_service.proto",
} as DescFile;

export interface ListDeploymentComponentsRequest {
  provider?: string;
  kind?: string;
}

export interface DeploymentComponent {
  id: string;
  kind: string;
  provider: string;
  name: string;
  version: string;
  idPrefix: string;
  isServiceKind: boolean;
  createdAt?: { seconds: number; nanos?: number };
  updatedAt?: { seconds: number; nanos?: number };
}

export interface ListDeploymentComponentsResponse {
  components: DeploymentComponent[];
}

// Export as message schemas compatible with @bufbuild/protobuf
export const ListDeploymentComponentsRequestSchema: DescMessage = {
  kind: "message",
  typeName: "backend.v1.ListDeploymentComponentsRequest",
  name: "ListDeploymentComponentsRequest",
  file: mockFile,
  fields: [],
  field: {},
  oneofs: [],
  oneof: {},
  reservedNames: [],
  reservedRanges: [],
  extensionRanges: [],
  members: [], // Make members iterable for createZeroMessage
  proto: {} as any,
  toString: () => "backend.v1.ListDeploymentComponentsRequest",
} as DescMessage;

export const ListDeploymentComponentsResponseSchema: DescMessage = {
  kind: "message",
  typeName: "backend.v1.ListDeploymentComponentsResponse",
  name: "ListDeploymentComponentsResponse",
  file: mockFile,
  fields: [],
  field: {},
  oneofs: [],
  oneof: {},
  reservedNames: [],
  reservedRanges: [],
  extensionRanges: [],
  members: [], // Make members iterable for createZeroMessage
  proto: {} as any,
  toString: () => "backend.v1.ListDeploymentComponentsResponse",
} as DescMessage;

// Also export the types
export type { ListDeploymentComponentsRequest, ListDeploymentComponentsResponse, DeploymentComponent };
