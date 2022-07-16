package dynamic

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGuiderType_init(t *testing.T) {
	testCases := []struct {
		name   string
		method string
		url    string
		want1  []string
		want2  int
	}{
		{
			name:   "normal success",
			method: http.MethodGet,
			url:    "/api/version",
			want1:  []string{"version", http.MethodGet},
			want2:  0,
		},
		{
			name:   "empty path",
			method: http.MethodGet,
			url:    "",
			want1:  []string{http.MethodGet},
			want2:  0,
		},
		{
			name:   "bear api",
			method: http.MethodGet,
			url:    "/api",
			want1:  []string{http.MethodGet},
			want2:  0,
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, (func(t *testing.T) {
			gotInst := &guiderType{}
			gotInst.init(tt.method, tt.url)
			if !reflect.DeepEqual(gotInst.steps, tt.want1) {
				t.Errorf("init got steps = %+v, want %+v", gotInst.steps, tt.want1)
			}
			if gotInst.iterator != tt.want2 {
				t.Errorf("init got iterator = %d, want %d", gotInst.iterator, tt.want2)
			}
		}))
	}
}
