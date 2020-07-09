package file

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

// Definition represents serverless.yaml file structure
type Definition struct {
	Service     string              `yaml:"service,omitempty"`
	Description string              `yaml:"description,omitempty"`
	Provider    TriggermeshProvider `yaml:"provider,omitempty"`
	Repository  string              `yaml:"repository,omitempty"`
	Functions   map[string]Function `yaml:"functions,omitempty"`
	Include     []string            `yaml:"include,omitempty"`
}

// TriggermeshProvider structure contains serverless provider parameters specific to triggermesh
type TriggermeshProvider struct {
	Name         string            `yaml:"name,omitempty"`
	PullPolicy   string            `yaml:"pull-policy,omitempty"`
	Namespace    string            `yaml:"namespace,omitempty"`
	Runtime      string            `yaml:"runtime,omitempty"`
	Buildtimeout string            `yaml:"buildtimeout,omitempty"`
	Environment  map[string]string `yaml:"environment,omitempty"`
	EnvSecrets   []string          `yaml:"env-secrets,omitempty"`
	Annotations  map[string]string `yaml:"annotations,omitempty"`

	// registry configs moved to client Configset
	// these variables kept for backward compatibility
	Registry       string `yaml:"registry,omitempty"`
	RegistrySecret string `yaml:"registry-secret,omitempty"`
}

// Function describes function definition in serverless format
type Function struct {
	Handler     string            `yaml:"handler,omitempty"`
	Source      string            `yaml:"source,omitempty"`
	Revision    string            `yaml:"revision,omitempty"`
	Runtime     string            `yaml:"runtime,omitempty"`
	Concurrency int               `yaml:"concurrency,omitempty"`
	Buildargs   []string          `yaml:"buildargs,omitempty"`
	Description string            `yaml:"description,omitempty"`
	Labels      []string          `yaml:"labels,omitempty"`
	Environment map[string]string `yaml:"environment,omitempty"`
	EnvSecrets  []string          `yaml:"env-secrets,omitempty"`
	Annotations map[string]string `yaml:"annotations,omitempty"`
	Schedule    []Schedule        `yaml:"schedule,omitempty"`
}

// Schedule struct contains a data in JSON format and a cron
// that defines how often events should be sent to a function.
// Description string may be used to explain events purpose.
type Schedule struct {
	Cron        string
	Data        string
	Description string
}

// Aos returns filesystem object with standard set of os methods implemented by afero package
var Aos = afero.NewOsFs()

// ParseManifest accepts serverless yaml file path and returns decoded structure
func ParseManifest(path string) (Definition, error) {
	var definition Definition

	exists, err := afero.Exists(Aos, path)

	if !exists || err != nil {
		return definition, errors.New("could not find manifest file")
	}

	data, err := afero.ReadFile(Aos, path)

	if err != nil {
		return definition, err
	}

	definition.Repository = filepath.Base(filepath.Dir(path))
	err = yaml.UnmarshalStrict(data, &definition)

	return definition, err
}

// Validate function verifies that provided service Definition object contains required set of keys and values
func (definition Definition) Validate() error {
	if definition.Provider.Name != "" && definition.Provider.Name != "triggermesh" {
		return fmt.Errorf("%s provider is not supported", definition.Provider.Name)
	}

	if len(definition.Service) == 0 {
		return errors.New("Service name can't be empty")
	}

	return nil
}
