package scheduler

import (
	"github.com/G-Research/armada/internal/scheduler/schedulerobjects"
	"github.com/G-Research/armada/pkg/api"
	v1 "k8s.io/api/core/v1"
)

type taintsFilterFunc func(*v1.Taint) bool
type labelsFilterFunc func(key, value string) bool

// NodeType represents a particular combination of taints and labels.
// The scheduler groups nodes by node type. When assigning pods to nodes,
// the scheduler only considers nodes with a NodeType for which the taints and labels match.
type NodeType struct {
	// Unique identifier of this node type.
	// Used for map lookup.
	id     string
	Taints []v1.Taint
	Labels map[string]string
}

func NewNodeTypeFromNode(node *v1.Node, taintsFilter taintsFilterFunc, labelsFilter labelsFilterFunc) *NodeType {
	return NewNodeType(node.Spec.Taints, node.GetLabels(), taintsFilter, labelsFilter)
}

func NewNodeTypeFromNodeInfo(nodeInfo *api.NodeInfo, taintsFilter taintsFilterFunc, labelsFilter labelsFilterFunc) *NodeType {
	return NewNodeType(nodeInfo.GetTaints(), nodeInfo.GetLabels(), taintsFilter, labelsFilter)
}

func NewNodeType(taints []v1.Taint, labels map[string]string, taintsFilter taintsFilterFunc, labelsFilter labelsFilterFunc) *NodeType {
	return &NodeType{
		Taints: getFilteredTaints(taints, taintsFilter),
		Labels: getFilteredLabels(labels, labelsFilter),
	}
}

// getFilteredTaints returns a list of taints satisfying the filter predicate.
func getFilteredTaints(taints []v1.Taint, inclusionFilter taintsFilterFunc) []v1.Taint {
	if inclusionFilter == nil {
		return taints
	}
	filteredTaints := []v1.Taint{}
	for _, taint := range taints {
		if !inclusionFilter(&taint) {
			continue
		}
		filteredTaints = append(filteredTaints, taint)
	}
	return filteredTaints
}

// getFilteredLabels returns a map of labels satisfying the filter predicate.
func getFilteredLabels(labels map[string]string, inclusionFilter labelsFilterFunc) map[string]string {
	if inclusionFilter == nil {
		return labels
	}
	filteredLabels := make(map[string]string)
	for key, value := range labels {
		if !inclusionFilter(key, value) {
			continue
		}
		filteredLabels[key] = value
	}
	return filteredLabels
}

// canSchedulePod determines whether a pod can be scheduled on nodes of this NodeType.
// If the pod can't be scheduled, the returned error indicates why.
// If no error is returned, the pod can be scheduled on nodes of this NodeType.
//
// TODO: Remove in favour of PodRequirementsMet.
func (nodeType *NodeType) canSchedulePod(req *PodSchedulingRequirements) error {
	return req.toleratesTaintsAndMatchesSelector(nodeType.Taints, nodeType.Labels)
}

func (nodeType *NodeType) PodRequirementsMet(req *schedulerobjects.PodRequirements) (bool, PodRequirementsNotMetReason, error) {
	return PodRequirementsMet(nodeType.Taints, nodeType.Labels, req)
}
