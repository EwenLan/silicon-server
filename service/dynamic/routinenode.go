package dynamic

import (
	"net/http"

	"github.com/EwenLan/silicon-server/slog"
)

// NodeHandler
type NodeHandler interface {
	HttpHandle(http.ResponseWriter, *http.Request)
	Init()
}

type routineNode struct {
	routineTable map[string]*routineNode
	handler      NodeHandler
}

func (r *routineNode) searchRoutineNode(guider *guiderType) NodeHandler {
	curr := guider.getCurrent()
	if curr == "" {
		slog.Debugf("reach end of guider, curr = %s", curr)
		return r.handler
	}
	nextNode, ok := r.routineTable[curr]
	if (!ok) || (nextNode == nil) {
		slog.Debugf("reach leaf of routine tree, curr = %s", curr)
		return r.handler
	}
	slog.Debugf("proceed one step, curr = %s", curr)
	guider.moveOneStep()
	return nextNode.searchRoutineNode(guider)
}
