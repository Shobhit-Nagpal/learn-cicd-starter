package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type test struct {
		input http.Header
		want  string
		err   error
	}

	tests := []test{
		{map[string][]string{}, "", ErrNoAuthHeaderIncluded},
		{map[string][]string{"Authorization": []string{"ApiKey abcd"}}, "abcd", nil},
		{map[string][]string{"Authorization": []string{"ApiKey abcd"}}, "lol", nil},
		{map[string][]string{"Authorization": []string{"abcd"}}, "", errors.New("malformed authorization header")},
		{map[string][]string{"Authorization": []string{"abcd"}}, "", errors.New("malformed authorization header")},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if err != nil {
			if tc.err == nil {
				t.Fatalf("expected: nil, got: %v", err)
			}

			if err.Error() != tc.err.Error() {
				t.Fatalf("expected error: %v, got: %v", tc.err.Error(), err.Error())
			}

			continue

		} else if tc.err != nil {
			t.Fatalf("expected error: %v, got: nil", tc.err.Error())
		}

		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
