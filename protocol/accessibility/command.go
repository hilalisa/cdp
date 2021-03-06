// Code generated by cdpgen. DO NOT EDIT.

package accessibility

import (
	"github.com/mafredri/cdp/protocol/dom"
	"github.com/mafredri/cdp/protocol/runtime"
)

// GetPartialAXTreeArgs represents the arguments for GetPartialAXTree in the Accessibility domain.
type GetPartialAXTreeArgs struct {
	NodeID         *dom.NodeID             `json:"nodeId,omitempty"`         // Identifier of the node to get the partial accessibility tree for.
	BackendNodeID  *dom.BackendNodeID      `json:"backendNodeId,omitempty"`  // Identifier of the backend node to get the partial accessibility tree for.
	ObjectID       *runtime.RemoteObjectID `json:"objectId,omitempty"`       // JavaScript object id of the node wrapper to get the partial accessibility tree for.
	FetchRelatives *bool                   `json:"fetchRelatives,omitempty"` // Whether to fetch this nodes ancestors, siblings and children. Defaults to true.
}

// NewGetPartialAXTreeArgs initializes GetPartialAXTreeArgs with the required arguments.
func NewGetPartialAXTreeArgs() *GetPartialAXTreeArgs {
	args := new(GetPartialAXTreeArgs)

	return args
}

// SetNodeID sets the NodeID optional argument. Identifier of the node
// to get the partial accessibility tree for.
func (a *GetPartialAXTreeArgs) SetNodeID(nodeID dom.NodeID) *GetPartialAXTreeArgs {
	a.NodeID = &nodeID
	return a
}

// SetBackendNodeID sets the BackendNodeID optional argument.
// Identifier of the backend node to get the partial accessibility tree
// for.
func (a *GetPartialAXTreeArgs) SetBackendNodeID(backendNodeID dom.BackendNodeID) *GetPartialAXTreeArgs {
	a.BackendNodeID = &backendNodeID
	return a
}

// SetObjectID sets the ObjectID optional argument. JavaScript object
// id of the node wrapper to get the partial accessibility tree for.
func (a *GetPartialAXTreeArgs) SetObjectID(objectID runtime.RemoteObjectID) *GetPartialAXTreeArgs {
	a.ObjectID = &objectID
	return a
}

// SetFetchRelatives sets the FetchRelatives optional argument.
// Whether to fetch this nodes ancestors, siblings and children.
// Defaults to true.
func (a *GetPartialAXTreeArgs) SetFetchRelatives(fetchRelatives bool) *GetPartialAXTreeArgs {
	a.FetchRelatives = &fetchRelatives
	return a
}

// GetPartialAXTreeReply represents the return values for GetPartialAXTree in the Accessibility domain.
type GetPartialAXTreeReply struct {
	Nodes []AXNode `json:"nodes"` // The `Accessibility.AXNode` for this DOM node, if it exists, plus its ancestors, siblings and children, if requested.
}
