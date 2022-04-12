package manifest

import (
	"github.com/opencontainers/image-spec/specs-go"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/trippleflp/go-oci/util/chainer"
)

type Manifest struct {
	chainer.Chainer[v1.Manifest]
}

func InitManifest() *Manifest {
	manifest := v1.Manifest{
		Versioned: specs.Versioned{SchemaVersion: 1},
		MediaType: "application/vnd.oci.image.manifest.v1+json",
		Layers:    []v1.Descriptor{},
	}

	return &Manifest{chainer.BuildChainer(&manifest)}
}

func (m *Manifest) AddLayer(layer v1.Descriptor) *Manifest {
	m.Data().Layers = append(m.Data().Layers, layer)
	return m
}

func (m *Manifest) SetConfig(config v1.Descriptor) *Manifest {
	m.Data().Config = config
	return m
}

func (m *Manifest) SetAnnotations(annotations map[string]string) *Manifest {
	m.Data().Annotations = annotations
	return m
}
