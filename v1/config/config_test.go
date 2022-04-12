package config

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"github.com/stretchr/testify/assert"
	"github.com/trippleflp/go-oci/util"
)

func TestChainError(t *testing.T) {
	data, err := BasicConfig("test", "No File").
		SetAnnotations(map[string]string{
			"123": "123",
		}).
		SetPlatform(v1.Platform{Architecture: "amd64"}).
		SetURLs([]string{"urltest"}).
		Catch(func(err error) {
			fmt.Printf("custom Error: %s", err)
		})

	assert.Error(t, err)
	assert.Nil(t, data)
}

func TestChain(t *testing.T) {
	exampleFilePath := filepath.Join(util.TestPath, "example.json")
	data, err := BasicConfig("test", exampleFilePath).
		SetAnnotations(map[string]string{
			"123": "123",
		}).
		SetPlatform(v1.Platform{Architecture: "amd64"}).
		SetURLs([]string{"urltest"}).
		Unpack()

	assert.NoError(t, err)
	assert.Equal(t, data.MediaType, "test")

	file, _ := os.ReadFile(exampleFilePath)

	assert.Equal(t, data.Size, int64(len(file)))
	assert.Equal(t, data.Annotations, map[string]string{
		"123": "123",
	})
	assert.Equal(t, data.URLs, []string{"urltest"})
	assert.Equal(t, *data.Platform, v1.Platform{Architecture: "amd64"})

}
