package dynamic

import (
	"strings"

	"github.com/EwenLan/silicon-server/globaldefine"
	"github.com/EwenLan/silicon-server/service/dynamic/about"
	"github.com/EwenLan/silicon-server/service/dynamic/jsonhandler"
)

type guiderType struct {
	steps    []string
	iterator int
}

var rootRoutineNode routineNode

func (g *guiderType) getCurrent() string {
	if g.iterator < len(g.steps) {
		return g.steps[g.iterator]
	}
	return ""
}

func (g *guiderType) moveOneStep() {
	if g.iterator < len(g.steps) {
		g.iterator++
	}
}

func (g *guiderType) init(method string, url string) {
	parts := strings.Split(url, pathSeprator)
	if len(parts) >= guiderStartIndex {
		g.steps = parts[guiderStartIndex:]
	}
	g.steps = append(g.steps, method)
	g.iterator = 0
}

func InitRootRoutineNode() {
	rootRoutineNode.routineTable = map[string]*routineNode{
		"version": {
			handler: &jsonhandler.JsonHandle{
				ResponseContent: &globaldefine.VersionInfoPrototype{},
				JsonHandleFunc:  about.ServeAbout,
			},
		},
	}
}
