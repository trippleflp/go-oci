package config

import (
	"os"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"

	"github.com/trippleflp/go-oci/util"
	"github.com/trippleflp/go-oci/util/chainer"
)

type Config struct {
	chainer.Chainer[v1.Descriptor]
}

func buildChainer(spec v1.Descriptor) chainer.Chainer[v1.Descriptor] {
	return chainer.BuildChainer(&spec)
}

func buildChainerErr(err error) chainer.Chainer[v1.Descriptor] {
	return chainer.BuildChainerWithError[v1.Descriptor](err)
}

func BasicConfigFromBytes(mediaType string, data []byte) *Config {
	configSpec := v1.Descriptor{
		MediaType: mediaType,
		Size:      util.GetSize(data),
		Digest:    util.GetDigest(data),
	}
	return &Config{buildChainer(configSpec)}
}

func BasicConfig(mediaType string, filePath string) *Config {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return &Config{buildChainerErr(err)}
	}
	return BasicConfigFromBytes(mediaType, data)
}

func (c *Config) SetAnnotations(annotations map[string]string) *Config {
	if !c.Ok() {
		return c
	}
	c.Data().Annotations = annotations
	return c
}

func (c *Config) SetURLs(urls []string) *Config {
	if !c.Ok() {
		return c
	}
	c.Data().URLs = urls
	return c
}

func (c *Config) SetPlatform(platform v1.Platform) *Config {
	if !c.Ok() {
		return c
	}
	c.Data().Platform = &platform
	return c
}
