package auth

import (
	"errors"
	"net/http"
	"reflect"
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
		{map[string][]string{"Authorization": []string{"ApiKey abcd"}}, "abcd", nil},
		{map[string][]string{"Authorization": []string{"abcd"}}, "", errors.New("malformed authorization header")},
	}

	for _, tc := range tests {
		got, err := GetAPIKey(tc.input)
		if err != nil {
			if !errors.Is(err, tc.err) {
				t.Fatalf("expected: %v, got: %v", tc.err.Error(), err.Error())
			}

      return
		}

		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
