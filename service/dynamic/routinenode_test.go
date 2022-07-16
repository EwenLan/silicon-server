package dynamic

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/EwenLan/silicon-server/service/dynamic/jsonhandler"
)

func TestRoutineNode_searchRoutineNode(t *testing.T) {
	var targetHandler = &jsonhandler.JsonHandle{}
	var notTargetHandler = &jsonhandler.JsonHandle{}
	testCases := []struct {
		name        string
		routineNode *routineNode
		guider      *guiderType
		want        NodeHandler
	}{
		{
			name: "normal presice success",
			guider: &guiderType{
				steps:    []string{"version", http.MethodGet},
				iterator: 0,
			},
			routineNode: &routineNode{
				routineTable: map[string]*routineNode{
					"version": {
						routineTable: map[string]*routineNode{
							http.MethodGet:  {handler: targetHandler},
							http.MethodPost: {handler: notTargetHandler},
						},
						handler: notTargetHandler,
					},
				},
				handler: notTargetHandler,
			},
			want: targetHandler,
		},
		{
			name: "method wildcard success",
			guider: &guiderType{
				steps:    []string{"version", http.MethodGet},
				iterator: 0,
			},
			routineNode: &routineNode{
				routineTable: map[string]*routineNode{
					"version": {
						handler: targetHandler,
					},
				},
				handler: notTargetHandler,
			},
			want: targetHandler,
		},
		{
			name: "routine tree wildcard success",
			guider: &guiderType{
				steps:    []string{"version", "foo", http.MethodGet},
				iterator: 0,
			},
			routineNode: &routineNode{
				routineTable: map[string]*routineNode{
					"version": {
						handler: targetHandler,
					},
				},
				handler: notTargetHandler,
			},
			want: targetHandler,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.routineNode.searchRoutineNode(tt.guider)
			if reflect.ValueOf(got) != reflect.ValueOf(tt.want) {
				t.Errorf("searchRoutineTree got = %+v, want %+v", got, tt.want)
			}
		})
	}
}
