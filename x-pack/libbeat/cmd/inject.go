// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	"github.com/elastic/beats/v7/libbeat/cmd"

	// register central management
	_ "github.com/elastic/beats/v7/x-pack/libbeat/management"

	// Register fleet
	_ "github.com/elastic/beats/v7/x-pack/libbeat/management/fleet"
	// register processors
	_ "github.com/elastic/beats/v7/x-pack/libbeat/processors/add_cloudfoundry_metadata"
	_ "github.com/elastic/beats/v7/x-pack/libbeat/processors/add_nomad_metadata"

	// register autodiscover providers
	_ "github.com/elastic/beats/v7/x-pack/libbeat/autodiscover/providers/aws/ec2"
	_ "github.com/elastic/beats/v7/x-pack/libbeat/autodiscover/providers/aws/elb"
	_ "github.com/elastic/beats/v7/x-pack/libbeat/autodiscover/providers/nomad"
)

// AddXPack extends the given root folder with XPack features
func AddXPack(root *cmd.BeatsRootCmd, name string) {
	root.AddCommand(genEnrollCmd(name, ""))
}
