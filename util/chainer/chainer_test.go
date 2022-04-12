package chainer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildChainerWithError(t *testing.T) {
	expectedErr := errors.New("error")
	chainer := BuildChainerWithError[string](expectedErr)
	assert.False(t, chainer.Ok())
	assert.Nil(t, chainer.Data())

	data, err := chainer.Catch(func(err error) {
		assert.Equal(t, err, expectedErr)
	})

	assert.Nil(t, data)
	assert.Equal(t, expectedErr, err)
}

func TestChainerFail(t *testing.T) {
	str := ""
	expectedErr := errors.New("error")
	chainer := BuildChainer(&str)
	chainer.Fail(expectedErr).Catch(func(err error) {
		assert.Equal(t, err, expectedErr)
	})

}

func TestChainer(t *testing.T) {
	str := "123"
	chainer := BuildChainer(&str)
	data, err := chainer.Unpack()
	assert.Nil(t, err)
	assert.Equal(t, *data, str)

	newStr := "456"
	data, err = chainer.Update(&newStr).Unpack()
	assert.Nil(t, err)
	assert.Equal(t, *data, newStr)
}
