/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package didclient provides REST operations.
package didclient

import (
	"fmt"
	"net/http"

	"github.com/trustbloc/agent-sdk/pkg/controller/command/didclient"
	"github.com/trustbloc/agent-sdk/pkg/controller/internal/cmdutil"
	"github.com/trustbloc/agent-sdk/pkg/controller/rest"
)

// constants for endpoints of DIDClient.
const (
	OperationID       = "/didclient"
	CreateBlocDIDPath = OperationID + "/create-trustbloc-did"
	CreatePeerDIDPath = OperationID + "/create-peer-did"
)

// Operation is controller REST service controller for DID Client.
type Operation struct {
	command  *didclient.Command
	handlers []rest.Handler
}

// New returns new DID client rest instance.
func New(ctx didclient.Provider, domain string) (*Operation, error) {
	client, err := didclient.New(domain, ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize did-client command: %w", err)
	}

	o := &Operation{command: client}
	o.registerHandler()

	return o, nil
}

// GetRESTHandlers get all controller API handler available for this protocol service.
func (c *Operation) GetRESTHandlers() []rest.Handler {
	return c.handlers
}

// registerHandler register handlers to be exposed from this protocol service as REST API endpoints.
func (c *Operation) registerHandler() {
	// Add more protocol endpoints here to expose them as controller API endpoints
	c.handlers = []rest.Handler{
		cmdutil.NewHTTPHandler(CreateBlocDIDPath, http.MethodPost, c.CreateTrustBlocDID),
		cmdutil.NewHTTPHandler(CreatePeerDIDPath, http.MethodPost, c.CreatePeerDID),
	}
}

// CreateTrustBlocDID swagger:route POST /didclient/create-trustbloc-did didclient createTrustBlocDID
//
// Creates a new trust bloc DID.
//
// Responses:
//    default: genericError
//    200: createDIDResp
func (c *Operation) CreateTrustBlocDID(rw http.ResponseWriter, req *http.Request) {
	rest.Execute(c.command.CreateTrustBlocDID, rw, req.Body)
}

// CreatePeerDID swagger:route POST /didclient/create-peer-did didclient createPeerDID
//
// Creates a new peer DID.
//
// Responses:
//    default: genericError
//    200: createDIDResp
func (c *Operation) CreatePeerDID(rw http.ResponseWriter, req *http.Request) {
	rest.Execute(c.command.CreatePeerDID, rw, req.Body)
}
