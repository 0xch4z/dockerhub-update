package update

import "testing"

func TestMustParseURI(t *testing.T) {
	for _, tc := range []struct {
		uri       string
		repo      string
		namespace string
		invalid   bool
	}{
		{
			uri:     "",
			invalid: true,
		},
		{
			uri:       "someuser/somerepo",
			namespace: "someuser",
			repo:      "somerepo",
		},
		{
			uri:     "some/user/some/repo",
			invalid: true,
		},
		{
			uri:       "namespace/repo",
			namespace: "namespace",
			repo:      "repo",
		},
	} {
		namespace, repo, err := mustParseURI(tc.uri)
		if tc.invalid && err == nil {
			t.Errorf("Expected error to be thrown for '%s'\n", tc.uri)
		}

		if err != nil && !tc.invalid {
			t.Errorf("Expected error not to be thrown for '%s', but got %v\n", tc.uri, err)
		}

		if !tc.invalid {
			if namespace != tc.namespace {
				t.Errorf("Expected namespace to be '%s' but got '%s'\n", tc.namespace, namespace)
			}

			if repo != tc.repo {
				t.Errorf("Expected repo to be '%s' but got '%s'\n", tc.repo, repo)
			}
		}
	}
}
