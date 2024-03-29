package main

import (
	"fmt"
)

// Package represents a software package
type Package struct {
	Name         string
	Version      string
	Description  string
	Dependencies []string
}

// PackageOption represents a function type that modifies Package configuration
type PackageOption func(*Package)

// WithVersion sets the package version
func WithVersion(version string) PackageOption {
	return func(p *Package) {
		p.Version = version
	}
}

// WithDescription sets the package description
func WithDescription(description string) PackageOption {
	return func(p *Package) {
		p.Description = description
	}
}

// WithDependencies sets the package dependencies
func WithDependencies(dependencies ...string) PackageOption {
	return func(p *Package) {
		p.Dependencies = dependencies
	}
}

// NewPackage creates a new Package with the provided options
func NewPackage(name string, options ...PackageOption) *Package {
	pkg := &Package{
		Name:    name,
		Version: "0.0.1",
	}
	for _, option := range options {
		option(pkg)
	}
	return pkg
}

func main1() {
	// Create a new package with custom options
	pkg := NewPackage(
		"example-package",
		// WithVersion("1.0.0"),
		WithDescription("This is an example package."),
		WithDependencies("dependency1", "dependency2"),
	)

	// Print package information
	fmt.Println("Package Name:", pkg.Name)
	fmt.Println("Package Version:", pkg.Version)
	fmt.Println("Package Description:", pkg.Description)
	fmt.Println("Package Dependencies:", pkg.Dependencies)
}
