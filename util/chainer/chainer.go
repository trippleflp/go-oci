package chainer

type Chainer[T interface{}] struct {
	element *T
	err     error
}

func BuildChainer[T interface{}](element *T) Chainer[T] {
	return Chainer[T]{
		element: element,
		err:     nil,
	}
}

func BuildChainerWithError[T interface{}](err error) Chainer[T] {
	return Chainer[T]{
		element: nil,
		err:     err,
	}
}

func (c *Chainer[T]) Fail(err error) {
	c.err = err
}

func (c *Chainer[T]) Update(element *T) {
	c.element = element
}

func (c *Chainer[T]) Ok() bool {
	return c.err == nil
}

func (c *Chainer[T]) Data() *T {
	return c.element
}

func (c *Chainer[T]) Unpack() (*T, error) {
	return c.element, c.err
}

func (c *Chainer[T]) Catch(fn func(error)) (*T, error) {
	if !c.Ok() {
		fn(c.err)
	}

	return c.element, c.err
}
