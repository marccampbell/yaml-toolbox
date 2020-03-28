package splitter

type OverlySimpleGVK struct {
	APIVersion string               `yaml:"apiVersion"`
	Kind       string               `yaml:"kind"`
	Metadata   OverlySimpleMetadata `yaml:"metadata"`
}

type OverlySimpleMetadata struct {
	Name        string                 `yaml:"name"`
	Annotations map[string]interface{} `yaml:"annotations"`
}
