package utils

import "testing"

func TestGetFilenameFromPath(t *testing.T) {
	testCases := []struct {
		name string
		path string
		want string
	}{
		{
			name: "normal case",
			path: "silicon-server/utils/utils.go",
			want: "utils.go",
		},
		{
			name: "empty path",
			path: "",
			want: "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := GetFilenameFromPath(tt.path)
			if got != tt.want {
				t.Errorf("GetFilenameFromPath got = %s, want = %s", got, tt.want)
			}
		})
	}
}

func TestGetDirectoryFromPath(t *testing.T) {
	testCases := []struct {
		name string
		path string
		want string
	}{
		{
			name: "normal relative path",
			path: "silicon-server/utils/utils.go",
			want: "silicon-server/utils",
		},
		{
			name: "normal absolute path",
			path: "C:/silicon-server/utils/utils.go",
			want: "C:/silicon-server/utils",
		},
		{
			name: "only filename",
			path: "utils.go",
			want: "",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := GetDirectoryFromPath(tt.path)
			if got != tt.want {
				t.Errorf("GetDirectoryFromPath got = %s, want = %s", got, tt.want)
			}
		})
	}
}
