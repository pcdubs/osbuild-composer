package container

import (
	"github.com/containers/image/v5/docker/reference"
	"github.com/opencontainers/go-digest"
)

// A Spec is the specification of how to get a specific
// container from a Source and under what LocalName to
// store it in an image. The container is identified by
// at the Source via Digest and ImageID. The latter one
// should remain the same in the target image as well.
type Spec struct {
	Source    string // does not include the manifest digest
	Digest    string // digest of the manifest at the Source
	TLSVerify *bool  // controls TLS verification
	ImageID   string // container image identifier
	LocalName string // name to use inside the image
}

// NewSpec creates a new Spec from the essential information.
// It also converts is the transition point from container
// specific types (digest.Digest) to generic types (string).
func NewSpec(source reference.Named, digest, imageID digest.Digest) Spec {
	name := source.Name()
	return Spec{
		Source:  name,
		Digest:  digest.String(),
		ImageID: imageID.String(),

		LocalName: name,
	}
}
