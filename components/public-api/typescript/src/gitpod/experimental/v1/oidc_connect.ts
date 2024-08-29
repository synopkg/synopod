/**
 * Copyright (c) 2024 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

// @generated by protoc-gen-connect-es v1.1.2 with parameter "target=ts"
// @generated from file gitpod/experimental/v1/oidc.proto (package gitpod.experimental.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateClientConfigRequest, CreateClientConfigResponse, DeleteClientConfigRequest, DeleteClientConfigResponse, GetClientConfigRequest, GetClientConfigResponse, ListClientConfigsRequest, ListClientConfigsResponse, SetClientConfigActivationRequest, SetClientConfigActivationResponse, UpdateClientConfigRequest, UpdateClientConfigResponse } from "./oidc_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service gitpod.experimental.v1.OIDCService
 */
export const OIDCService = {
  typeName: "gitpod.experimental.v1.OIDCService",
  methods: {
    /**
     * Creates a new OIDC client configuration.
     *
     * @generated from rpc gitpod.experimental.v1.OIDCService.CreateClientConfig
     */
    createClientConfig: {
      name: "CreateClientConfig",
      I: CreateClientConfigRequest,
      O: CreateClientConfigResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Retrieves an OIDC client configuration by ID.
     *
     * @generated from rpc gitpod.experimental.v1.OIDCService.GetClientConfig
     */
    getClientConfig: {
      name: "GetClientConfig",
      I: GetClientConfigRequest,
      O: GetClientConfigResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Lists OIDC client configurations.
     *
     * @generated from rpc gitpod.experimental.v1.OIDCService.ListClientConfigs
     */
    listClientConfigs: {
      name: "ListClientConfigs",
      I: ListClientConfigsRequest,
      O: ListClientConfigsResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Updates modifiable properties of an existing OIDC client configuration.
     *
     * @generated from rpc gitpod.experimental.v1.OIDCService.UpdateClientConfig
     */
    updateClientConfig: {
      name: "UpdateClientConfig",
      I: UpdateClientConfigRequest,
      O: UpdateClientConfigResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Removes an OIDC client configuration by ID.
     *
     * @generated from rpc gitpod.experimental.v1.OIDCService.DeleteClientConfig
     */
    deleteClientConfig: {
      name: "DeleteClientConfig",
      I: DeleteClientConfigRequest,
      O: DeleteClientConfigResponse,
      kind: MethodKind.Unary,
    },
    /**
     * Activates an OIDC client configuration by ID.
     *
     * @generated from rpc gitpod.experimental.v1.OIDCService.SetClientConfigActivation
     */
    setClientConfigActivation: {
      name: "SetClientConfigActivation",
      I: SetClientConfigActivationRequest,
      O: SetClientConfigActivationResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;
