module github.com/hashicorp/packer-plugin-chef

go 1.16

require (
	github.com/hashicorp/hcl/v2 v2.10.1
	github.com/hashicorp/packer-plugin-sdk v0.2.3
	github.com/zclconf/go-cty v1.8.4
)

// Incorrect plugin registration for chef provisioners; see hashicorp/packer-plugin-chef/pull/25
retract v0.0.1
