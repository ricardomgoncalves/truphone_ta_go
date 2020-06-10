package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("should return error with message", func(t *testing.T) {
		input := []string{
			"error",
			"test1",
		}

		for i, message := range input {
			err := New(message)

			require.NotNil(t, err, i)
			assert.Equal(t, message, err.Error(), i)
		}
	})
}

func TestErrorf(t *testing.T) {
	t.Run("should return error with formatted message", func(t *testing.T) {
		input := []string{
			"error",
			"test1",
		}

		for i, message := range input {
			err := Errorf("%v", message)

			require.NotNil(t, err, i)
			assert.Equal(t, message, err.Error(), i)
		}
	})
}

func TestIterate(t *testing.T) {
	t.Run("should return false on nil error", func(t *testing.T) {
		value := Iterate(nil, func(err error) bool {
			return true
		})

		assert.False(t, value)
	})
	t.Run("should return first iteration", func(t *testing.T) {
		value := Iterate(New("test"), func(err error) bool {
			return true
		})

		assert.True(t, value)

		value = Iterate(New("test"), func(err error) bool {
			return false
		})

		assert.False(t, value)
	})
	t.Run("should return on iteration", func(t *testing.T) {
		inner := New("inner")

		value := Iterate(Wrap(inner, New("outer")), func(err error) bool {
			if err == inner {
				return true
			}
			return true
		})

		assert.True(t, value)

		value = Iterate(Wrap(inner, New("outer")), func(err error) bool {
			if err == inner {
				return false
			}
			return true
		})

		assert.False(t, value)
	})
}

func TestCode(t *testing.T) {
	t.Run("should return 0 on nil error", func(t *testing.T) {
		assert.Equal(t, 0, Code(nil))
	})
	t.Run("should return code on first iteration", func(t *testing.T) {
		err := WithCode(New("error"), 1)
		require.NotNil(t, err)
		assert.Equal(t, 1, Code(err))
	})
	t.Run("should return 0 with wrapped error without code", func(t *testing.T) {
		err := Wrap(New("inner"), New("outer"))
		require.NotNil(t, err)
		assert.Equal(t, 0, Code(err))
	})
	t.Run("should return code with wrapped code", func(t *testing.T) {
		err := Wrap(WithCode(New("inner"), 1), New("outer"))
		require.NotNil(t, err)
		assert.Equal(t, 1, Code(err))
	})
}

func TestIsTemporary(t *testing.T) {
	t.Run("should return false on nil err", func(t *testing.T) {
		assert.False(t, IsTemporary(nil))
	})
	t.Run("should return false", func(t *testing.T) {
		assert.False(t, IsTemporary(New("error")))
		assert.False(t, IsTemporary(Wrap(New("inner"), New("outer"))))
	})
	t.Run("should return true", func(t *testing.T) {
		assert.True(t, IsTemporary(AsTemporary(New("error"))))
		assert.True(t, IsTemporary(Wrap(AsTemporary(New("inner")), New("outer"))))
	})
}

func TestIs(t *testing.T) {
	t.Run("should return false on nil error", func(t *testing.T) {
		assert.False(t, Is(nil, New("test")))
	})
	t.Run("should return false", func(t *testing.T) {
		assert.False(t, Is(New("one"), New("two")))
	})
	t.Run("should return true", func(t *testing.T) {
		assert.False(t, Is(New("one"), New("one")))
	})
}

func TestAnnotate(t *testing.T) {
	t.Run("should return nil on nil error", func(t *testing.T) {
		assert.Nil(t, Annotate(nil, "message"))
	})
	t.Run("should correctly format error string", func(t *testing.T) {
		assert.Equal(t, "error: something", Annotate(New("error"), "something").Error())
	})
}

func TestAnnotatef(t *testing.T) {
	t.Run("should return nil on nil error", func(t *testing.T) {
		assert.Nil(t, Annotatef(nil, "%v", "message"))
	})
	t.Run("should correctly format error string", func(t *testing.T) {
		assert.Equal(t, "error: something more", Annotatef(New("error"), "%v %v", "something", "more").Error())
	})
}

func TestAnnotated_Unwrap(t *testing.T) {
	t.Run("should unwrap properly the original error", func(t *testing.T) {
		originalErr := New("original")
		annotatedErr := Annotate(originalErr, "extra message")
		assert.NotEqual(t, originalErr, annotatedErr)
		assert.Equal(t, originalErr, errors.Unwrap(annotatedErr))
	})
}

func TestWrap(t *testing.T) {
	t.Run("should correctly wrap errors", func(t *testing.T) {
		assert.Equal(t, "outer: inner", Wrap(New("inner"), New("outer")).Error())
	})
}

func TestWrap_Unwrap(t *testing.T) {
	t.Run("should unwrap properly the inner error", func(t *testing.T) {
		innerErr := New("inner")
		outerErr := New("outer")
		wrappedErr := Wrap(innerErr, outerErr)
		assert.NotEqual(t, innerErr, wrappedErr)
		assert.Equal(t, innerErr, errors.Unwrap(wrappedErr))
	})
}

func TestWithCode(t *testing.T) {
	t.Run("should return with code", func(t *testing.T) {
		assert.Equal(t, "123: error", WithCode(New("error"), 123).Error())
	})
}

func TestCoded_Unwrap(t *testing.T) {
	t.Run("should unwrap properly the inner error", func(t *testing.T) {
		innerErr := New("inner")
		codeErr := WithCode(innerErr, 123)
		assert.NotEqual(t, innerErr, codeErr)
		assert.Equal(t, innerErr, errors.Unwrap(codeErr))
	})
}

func TestCoded_Code(t *testing.T) {
	t.Run("should retrieve error code", func(t *testing.T) {
		innerErr := New("inner")
		codeErr := WithCode(innerErr, 123)
		assert.Equal(t, 123, Code(codeErr))
	})
}

func TestAsTemporary(t *testing.T) {
	t.Run("should return temporary error", func(t *testing.T) {
		assert.Equal(t, "error[temporary]", AsTemporary(New("error")).Error())
	})
}

func TestTemporary_Unwrap(t *testing.T) {
	t.Run("should unwrap properly the inner error", func(t *testing.T) {
		innerErr := New("inner")
		temporaryErr := AsTemporary(innerErr)
		assert.NotEqual(t, innerErr, temporaryErr)
		assert.Equal(t, innerErr, errors.Unwrap(temporaryErr))
	})
}
