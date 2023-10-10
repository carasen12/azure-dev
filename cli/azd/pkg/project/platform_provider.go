package project

import "github.com/azure/azure-dev/cli/azd/pkg/ioc"

// PlatformProvider is an interface for a platform provider
type PlatformProvider interface {
	// Name returns the name of the platform
	Name() string

	// IsEnabled returns true if the platform is enabled
	IsEnabled() bool

	// ConfigureContainer configures the IoC container for the platform
	ConfigureContainer(container *ioc.NestedContainer) error
}
