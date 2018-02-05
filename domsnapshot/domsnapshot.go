// Package domsnapshot provides the Chrome Debugging Protocol
// commands, types, and events for the DOMSnapshot domain.
//
// This domain facilitates obtaining document snapshots with DOM, layout, and
// style information.
//
// Generated by the chromedp-gen command.
package domsnapshot

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"context"

	"github.com/jbegley1995/cdproto/cdp"
)

// GetSnapshotParams returns a document snapshot, including the full DOM tree
// of the root node (including iframes, template contents, and imported
// documents) in a flattened array, as well as layout and white-listed computed
// style information for the nodes. Shadow DOM in the returned DOM tree is
// flattened.
type GetSnapshotParams struct {
	ComputedStyleWhitelist []string `json:"computedStyleWhitelist"`          // Whitelist of computed styles to return.
	IncludeEventListeners  bool     `json:"includeEventListeners,omitempty"` // Whether or not to retrieve details of DOM listeners (default false).
}

// GetSnapshot returns a document snapshot, including the full DOM tree of
// the root node (including iframes, template contents, and imported documents)
// in a flattened array, as well as layout and white-listed computed style
// information for the nodes. Shadow DOM in the returned DOM tree is flattened.
//
// parameters:
//   computedStyleWhitelist - Whitelist of computed styles to return.
func GetSnapshot(computedStyleWhitelist []string) *GetSnapshotParams {
	return &GetSnapshotParams{
		ComputedStyleWhitelist: computedStyleWhitelist,
	}
}

// WithIncludeEventListeners whether or not to retrieve details of DOM
// listeners (default false).
func (p GetSnapshotParams) WithIncludeEventListeners(includeEventListeners bool) *GetSnapshotParams {
	p.IncludeEventListeners = includeEventListeners
	return &p
}

// GetSnapshotReturns return values.
type GetSnapshotReturns struct {
	DomNodes        []*DOMNode        `json:"domNodes,omitempty"`        // The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
	LayoutTreeNodes []*LayoutTreeNode `json:"layoutTreeNodes,omitempty"` // The nodes in the layout tree.
	ComputedStyles  []*ComputedStyle  `json:"computedStyles,omitempty"`  // Whitelisted ComputedStyle properties for each node in the layout tree.
}

// Do executes DOMSnapshot.getSnapshot against the provided context.
//
// returns:
//   domNodes - The nodes in the DOM tree. The DOMNode at index 0 corresponds to the root document.
//   layoutTreeNodes - The nodes in the layout tree.
//   computedStyles - Whitelisted ComputedStyle properties for each node in the layout tree.
func (p *GetSnapshotParams) Do(ctxt context.Context, h cdp.Executor) (domNodes []*DOMNode, layoutTreeNodes []*LayoutTreeNode, computedStyles []*ComputedStyle, err error) {
	// execute
	var res GetSnapshotReturns
	err = h.Execute(ctxt, CommandGetSnapshot, p, &res)
	if err != nil {
		return nil, nil, nil, err
	}

	return res.DomNodes, res.LayoutTreeNodes, res.ComputedStyles, nil
}

// Command names.
const (
	CommandGetSnapshot = "DOMSnapshot.getSnapshot"
)
