package repos

import (
	"fmt"
	"io/ioutil"

	"github.com/jenkins-x/go-scm/scm/factory"
	"sigs.k8s.io/yaml"
)

// Setup configures the DefaultIdentifier by parsing the mappings from the
// provided filename.
func Setup(filename string) error {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read %q: %w", filename, err)
	}

	data := map[string]string{}
	err = yaml.Unmarshal(b, &data)
	if err != nil {
		return fmt.Errorf("failed to parse %q: %w", filename, err)
	}

	matches := []factory.MappingFunc{}
	for k, v := range data {
		matches = append(matches, factory.Mapping(k, v))
	}
	identifier := factory.NewDriverIdentifier(matches...)
	factory.DefaultIdentifier = identifier
	return nil
}
