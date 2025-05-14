package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	// header := http.Header{}
	// header.Set("Authorization", "ApiKey someRandomApiKey")
	// result, _ := GetAPIKey(header)
	// if result != "someRandomApiKey" {
	// 	t.Errorf("Result was incorrect, got: %s, want: %s.", result, "someRandomApiKey")
	// }
	var tests = []struct {
		name    string
		headers http.Header
		want    string
	}{
		{
			"Valid api key",
			http.Header{
				"Authorization": []string{"ApiKey someRandomApiKey"},
			},
			"someRandomApiKey",
		},
		{
			"Empty api key",
			http.Header{
				"Authorization": []string{"ApiKey"},
			},
			"",
		},
		{
			"Empty header value",
			http.Header{
				"Authorization": []string{""},
			},
			"",
		},
		{
			"No header",
			http.Header{},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, _ := GetAPIKey(tt.headers)
			if result != tt.want {
				t.Errorf("got %s, want %s", result, tt.want)
			}
		})
	}
}
