package manifest

import (
	"testing"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/stretchr/testify/assert"
)

func TestManifest(t *testing.T) {
	manifest := InitManifest()
	data, err := manifest.AddLayer(v1.Descriptor{}).
		SetConfig(v1.Descriptor{}).
		SetAnnotations(map[string]string{"123": "456"}).
		Unpack()

	assert.NoError(t, err)
	assert.NotNil(t, data)

}
