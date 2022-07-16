package dynamic

import (
	"net/http"
	"reflect"
	"testing"
)

func TestRoutineNode_searchRoutineNode(t *testing.T) {
	var targetFunc routineNodeHandleFunc = func(w http.ResponseWriter, r *http.Request) {}
	var notTargetFunc routineNodeHandleFunc = func(w http.ResponseWriter, r *http.Request) {}
	testCases := []struct {
		name        string
		routineNode *routineNode
		guider      *guiderType
		want        routineNodeHandleFunc
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
							http.MethodGet: {
								handleFunc: targetFunc,
							},
							http.MethodPost: {
								handleFunc: notTargetFunc,
							},
						},
						handleFunc: notTargetFunc,
					},
				},
				handleFunc: notTargetFunc,
			},
			want: targetFunc,
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
						handleFunc: targetFunc,
					},
				},
				handleFunc: notTargetFunc,
			},
			want: targetFunc,
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
						handleFunc: targetFunc,
					},
				},
				handleFunc: notTargetFunc,
			},
			want: targetFunc,
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
