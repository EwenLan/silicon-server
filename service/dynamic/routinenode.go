package dynamic

import (
	"net/http"

	"github.com/EwenLan/silicon-server/slog"
)

type routineNodeHandleFunc func(http.ResponseWriter, *http.Request)

type routineNode struct {
	routineTable map[string]*routineNode
	handleFunc   routineNodeHandleFunc
}

func (r *routineNode) searchRoutineNode(guider *guiderType) routineNodeHandleFunc {
	curr := guider.getCurrent()
	if curr == "" {
		slog.Debugf("reach end of guider, curr = %s", curr)
		return r.handleFunc
	}
	nextNode, ok := r.routineTable[curr]
	if (!ok) || (nextNode == nil) {
		slog.Debugf("reach leaf of routine tree, curr = %s", curr)
		return r.handleFunc
	}
	slog.Debugf("proceedd one step, curr = %s", curr)
	guider.moveOneStep()
	return nextNode.searchRoutineNode(guider)
}
