package golang_generics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}

func TestData(t *testing.T)  {
	data := Data[string]{
		First: "Fahril",
		Second: "Hadi",
	}

	fmt.Println(data)
}

func TestGenericMethod(t *testing.T)  {
	data := Data[string]{
		First: "Fahril",
		Second: "Hadi",
	}

	assert.Equal(t, "Abu", data.ChangeFirst("Abu"))
	assert.Equal(t, "Hello Fahril", data.SayHello("Fahril"))
}