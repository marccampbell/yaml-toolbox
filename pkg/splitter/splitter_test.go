package splitter

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMarshalIndent(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]string
	}{
		{
			name: "single doc",
			input: `apiVersion: apps/v1
kind: Deployment
metadata:
  name: test
spec: ~`,
			expected: map[string]string{
				"test-deployment.yaml": `apiVersion: apps/v1
kind: Deployment
metadata:
  name: test
spec: ~`,
			},
		},
		{
			name: "two docs",
			input: `apiVersion: apps/v1
kind: Deployment
metadata:
  name: test
spec: ~
---
apiVersion: v1
kind: Service
metadata:
  name: test
spec: ~`,

			expected: map[string]string{
				"test-deployment.yaml": `apiVersion: apps/v1
kind: Deployment
metadata:
  name: test
spec: ~`,
				"test-service.yaml": `apiVersion: v1
kind: Service
metadata:
  name: test
spec: ~`,
			},
		},
		{
			name: "with an empty doc",
			input: `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test
spec: ~
---
apiVersion: v1
kind: Service
metadata:
  name: test
spec: ~`,

			expected: map[string]string{
				"test-deployment.yaml": `apiVersion: apps/v1
kind: Deployment
metadata:
  name: test
spec: ~`,
				"test-service.yaml": `apiVersion: v1
kind: Service
metadata:
  name: test
spec: ~`,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := require.New(t)

			actual, err := SplitYAML([]byte(test.input))
			req.NoError(err)

			// convert to strings to make it easier to view the failures
			actualConverted := map[string]string{}
			for f, v := range actual {
				actualConverted[f] = string(v)
			}
			assert.Equal(t, test.expected, actualConverted)
		})
	}
}
