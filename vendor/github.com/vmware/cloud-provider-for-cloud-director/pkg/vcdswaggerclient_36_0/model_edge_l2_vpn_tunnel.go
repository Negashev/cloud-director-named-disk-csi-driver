/*
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Apache-2.0
*/

/*
 * VMware Cloud Director OpenAPI
 *
 * VMware Cloud Director OpenAPI is a new API that is defined using the OpenAPI standards.<br/> This ReSTful API borrows some elements of the legacy VMware Cloud Director API and establishes new patterns for use as described below. <h4>Authentication</h4> Authentication and Authorization schemes are the same as those for the legacy APIs. You can authenticate using the JWT token via the <code>Authorization</code> header or specifying a session using <code>x-vcloud-authorization</code> (The latter form is deprecated). <h4>Operation Patterns</h4> This API follows the following general guidelines to establish a consistent CRUD pattern: <table> <tr>   <th>Operation</th><th>Description</th><th>Response Code</th><th>Response Content</th> </tr><tr>   <td>GET /items<td>Returns a paginated list of items<td>200<td>Response will include Navigational links to the items in the list. </tr><tr>   <td>POST /items<td>Returns newly created item<td>201<td>Content-Location header links to the newly created item </tr><tr>   <td>GET /items/urn<td>Returns an individual item<td>200<td>A single item using same data type as that included in list above </tr><tr>   <td>PUT /items/urn<td>Updates an individual item<td>200<td>Updated view of the item is returned </tr><tr>   <td>DELETE /items/urn<td>Deletes the item<td>204<td>No content is returned. </tr> </table> <h5>Asynchronous operations</h5> Asynchronous operations are determined by the server. In those cases, instead of responding as described above, the server responds with an HTTP Response code 202 and an empty body. The tracking task (which is the same task as all legacy API operations use) is linked via the URI provided in the <code>Location</code> header.<br/> All API calls can choose to service a request asynchronously or synchronously as determined by the server upon interpreting the request. Operations that choose to exhibit this dual behavior will have both options documented by specifying both response code(s) below. The caller must be prepared to handle responses to such API calls by inspecting the HTTP Response code. <h5>Error Conditions</h5> <b>All</b> operations report errors using the following error reporting rules: <ul>   <li>400: Bad Request - In event of bad request due to incorrect data or other user error</li>   <li>401: Bad Request - If user is unauthenticated or their session has expired</li>   <li>403: Forbidden - If the user is not authorized or the entity does not exist</li> </ul> <h4>OpenAPI Design Concepts and Principles</h4> <ul>   <li>IDs are full Uniform Resource Names (URNs).</li>   <li>OpenAPI's <code>Content-Type</code> is always <code>application/json</code></li>   <li>REST links are in the Link header.</li>   <ul>     <li>Multiple relationships for any link are represented by multiple values in a space-separated list.</li>     <li>Links have a custom VMware Cloud Director-specific &quot;model&quot; attribute that hints at the applicable data         type for the links.</li>     <li>title + rel + model attributes evaluates to a unique link.</li>     <li>Links follow Hypermedia as the Engine of Application State (HATEOAS) principles. Links are present if         certain operations are present and permitted for the user&quot;s current role and the state of the         referred entities.</li>   </ul>   <li>APIs follow a flat structure relying on cross-referencing other entities instead of the navigational style       used by the legacy VMware Cloud Director APIs.</li>   <li>Most endpoints that return a list support filtering and sorting similar to the query service in the legacy       VMware Cloud Director APIs.</li>   <li>Accept header must be included to specify the API version for the request similar to calls to existing legacy       VMware Cloud Director APIs.</li>   <li>Each feature has a version in the path element present in its URL.<br/>       <b>Note</b> API URL's without a version in their paths must be considered experimental.</li> </ul> 
 *
 * API version: 36.0
 * Contact: https://code.vmware.com/support
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

// Specifies the L2 VPN tunnel configuration.
type EdgeL2VpnTunnel struct {
	// The unique id of this L2 VPN tunnel. On updates, the id is required for the tunnel, while for create a new id will be generated. This id is not a VCD URN. 
	Id string `json:"id,omitempty"`
	// Name for the tunnel.
	Name string `json:"name"`
	// Description for the tunnel.
	Description string `json:"description,omitempty"`
	// The IP address of the local endpoint, which corresponds to the Edge Gateway the tunnel is being configured on. 
	LocalEndpointIP string `json:"localEndpointIP"`
	// The IP address of the remote endpoint, which corresponds to the device on the remote site terminating the VPN tunnel. 
	RemoteEndpointIP string `json:"remoteEndpointIP"`
	// The network CIDR block over which the session interfaces.
	TunnelInterface string `json:"tunnelInterface,omitempty"`
	// This is the Pre-shared key used for authentication, no specific format is required.
	PreSharedKey string `json:"preSharedKey,omitempty"`
	// This is the mode used by the local endpoint to establish an IKE Connection with the remote site. The default is INITIATOR. <ul>   <li>INITIATOR - Local endpoint initiates tunnel setup and will also respond to incoming tunnel setup requests from the peer gateway.</li>   <li>RESPOND_ONLY - Local endpoint shall only respond to incoming tunnel setup requests, it shall not initiate the tunnel setup.</li>   <li>ON_DEMAND - In this mode local endpoint will initiate tunnel creation once first packet matching the policy rule is received, and will also respond to   incoming initiation requests.</li> </ul> 
	ConnectorInitiationMode string `json:"connectorInitiationMode,omitempty"`
	// Described whether the tunnel is enabled or not. The default is true.
	Enabled bool `json:"enabled,omitempty"`
	// The current session mode, one of either SERVER or CLIENT. <ul>   <li>SERVER - In which the edge gateway acts as the server side of the L2 VPN tunnel and generates peer codes to distribute to client sessions.</li>   <li>CLIENT - In which the edge gateway receives peer codes from the server side of the L2 VPN tunnel to establish a connection.</li> </ul> 
	SessionMode string `json:"sessionMode"`
	// This property is a base64 encoded string of the full configuration for the tunnel, generated by the server-side L2 VPN session. An L2 VPN client session must receive and validate this string in order to successfully establish a tunnel, but be careful sharing or storing this code since it does contain the encoded PSK. Leave this property blank if this call is being used to establish a server-side session. 
	PeerCode string `json:"peerCode,omitempty"`
	// The list of OrgVDC Network entity references which are currently attached to this L2VPN tunnel. 
	AttachedNetworks []EntityReference `json:"attachedNetworks,omitempty"`
	// The list of OrgVDC Network entity references which are currently attached to this L2VPN tunnel. 
	StretchedNetworks []EdgeL2VpnStretchedNetwork `json:"stretchedNetworks,omitempty"`
	// Whether logging for the tunnel is enabled or not.
	Logging bool           `json:"logging,omitempty"`
	Version *ObjectVersion `json:"version,omitempty"`
}
