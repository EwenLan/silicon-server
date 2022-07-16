package static

import "testing"

func TestGetContentTypeByFilename(t *testing.T) {
	testCases := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "normal html",
			filename: "www/index.html",
			want:     "text/html",
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := getContentTypeByFilename(tt.filename)
			if got != tt.want {
				t.Errorf("getContentTypeByFilename got = %s, want = %s", got, tt.want)
			}
		})
	}
}
