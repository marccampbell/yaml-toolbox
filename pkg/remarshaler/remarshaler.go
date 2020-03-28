package remarshaler

import (
	"bytes"

	"github.com/pkg/errors"
	yaml "github.com/replicatedhq/yaml/v3"
)

// RemarshalYAML is a general purpose function
// This ensures that lines aren't wrapped at 80 chars which breaks some functionality
func RemarshalYAML(inputContent []byte) ([]byte, error) {
	yamlObj := map[string]interface{}{}

	err := yaml.Unmarshal(inputContent, &yamlObj)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal yaml")
	}

	inputContent, err = MarshalIndent(2, yamlObj)
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal yaml")
	}

	return inputContent, nil
}

func MarshalIndent(indent int, in interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(indent)
	enc.SetLineLength(-1)
	err := enc.Encode(in)
	if err != nil {
		return nil, errors.Wrapf(err, "marshal with indent %d", indent)
	}

	return buf.Bytes(), nil
}
