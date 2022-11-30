package main

import (
	"go_patterns/creational/builder"
	"go_patterns/creational/builder/builder_ex_2"
	"go_patterns/creational/factory"
)

func main() {
	// Builder Pattern
	builder.MainBuilder()
	builder_ex_2.BuildSiteMap()
	// Prototype factory
	factory.ExampleService()
}
