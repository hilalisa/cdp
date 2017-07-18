// +build go1.9

// Code generated by cdpgen. DO NOT EDIT.

package dom

// Node DOM interaction is implemented in terms of mirror objects that represent the actual DOM nodes. DOMNode is a base node mirror type.
type Node struct {
	NodeID           NodeID         `json:"nodeId"`                     // Node identifier that is passed into the rest of the DOM messages as the nodeId. Backend will only push node with given id once. It is aware of all requested nodes and will only fire DOM events for nodes known to the client.
	ParentID         *NodeID        `json:"parentId,omitempty"`         // The id of the parent node if any.
	BackendNodeID    BackendNodeID  `json:"backendNodeId"`              // The BackendNodeId for this node.
	NodeType         int            `json:"nodeType"`                   // Node's nodeType.
	NodeName         string         `json:"nodeName"`                   // Node's nodeName.
	LocalName        string         `json:"localName"`                  // Node's localName.
	NodeValue        string         `json:"nodeValue"`                  // Node's nodeValue.
	ChildNodeCount   *int           `json:"childNodeCount,omitempty"`   // Child count for Container nodes.
	Children         []Node         `json:"children,omitempty"`         // Child nodes of this node when requested with children.
	Attributes       []string       `json:"attributes,omitempty"`       // Attributes of the Element node in the form of flat array [name1, value1, name2, value2].
	DocumentURL      *string        `json:"documentURL,omitempty"`      // Document URL that Document or FrameOwner node points to.
	BaseURL          *string        `json:"baseURL,omitempty"`          // Base URL that Document or FrameOwner node uses for URL completion.
	PublicID         *string        `json:"publicId,omitempty"`         // DocumentType's publicId.
	SystemID         *string        `json:"systemId,omitempty"`         // DocumentType's systemId.
	InternalSubset   *string        `json:"internalSubset,omitempty"`   // DocumentType's internalSubset.
	XMLVersion       *string        `json:"xmlVersion,omitempty"`       // Document's XML version in case of XML documents.
	Name             *string        `json:"name,omitempty"`             // Attr's name.
	Value            *string        `json:"value,omitempty"`            // Attr's value.
	PseudoType       PseudoType     `json:"pseudoType,omitempty"`       // Pseudo element type for this node.
	ShadowRootType   ShadowRootType `json:"shadowRootType,omitempty"`   // Shadow root type.
	FrameID          *FrameID       `json:"frameId,omitempty"`          // Frame ID for frame owner elements.
	ContentDocument  *Node          `json:"contentDocument,omitempty"`  // Content document for frame owner elements.
	ShadowRoots      []Node         `json:"shadowRoots,omitempty"`      // Shadow root list for given element host.
	TemplateContent  *Node          `json:"templateContent,omitempty"`  // Content document fragment for template elements.
	PseudoElements   []Node         `json:"pseudoElements,omitempty"`   // Pseudo elements associated with this node.
	ImportedDocument *Node          `json:"importedDocument,omitempty"` // Import document for the HTMLImport links.
	DistributedNodes []BackendNode  `json:"distributedNodes,omitempty"` // Distributed nodes for given insertion point.
	IsSVG            *bool          `json:"isSVG,omitempty"`            // Whether the node is SVG.
}
