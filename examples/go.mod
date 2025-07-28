module ephemeral-deployment-example

go 1.21

require github.com/moondev/maas-client-go v0.0.0

require (
	github.com/google/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/spectrocloud/maas-client-go v0.0.2-beta // indirect
)

replace github.com/moondev/maas-client-go => ../
