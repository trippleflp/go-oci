package index

import (
	"github.com/opencontainers/image-spec/specs-go"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/trippleflp/go-oci/util/chainer"
)

type Index struct {
	chainer.Chainer[v1.Index]
}

func InitIndex() *Index {
	index := v1.Index{
		Versioned: specs.Versioned{SchemaVersion: 1},
		MediaType: "application/vnd.oci.image.index.v1+json",
		Manifests: []v1.Descriptor{},
	}

	return &Index{chainer.BuildChainer(&index)}
}

func (i *Index) AddManifestAsDescriptor(manifest v1.Descriptor) *Index {
	if !i.Ok() {
		return i
	}
	i.Data().Manifests = append(i.Data().Manifests, manifest)
	return i
}

// Frist marshal
// func (m *Index) AddManifest(manifest v1.Manifest) *Index{
// 	if !m.Ok(){
// 		return m
// 	}
// 	m.Data().Manifests = append(m.Data().Manifests, manifest)
// 	return m
// }

func (i *Index) SetAnnotations(annotaions map[string]string) *Index {
	if !i.Ok() {
		return i
	}

	i.Data().Annotations = annotaions
	return i
}
