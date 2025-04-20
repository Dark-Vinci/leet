package dynamicProgramming

import (
	"errors"
	"fmt"
	"reflect"
)

type Result[T any] struct {
	err error
	ok  T
}

func (r Result[T]) And(other Result[T]) Result[T] {

	if r.err != nil {
		return Result[T]{err: r.err}
	}

	return other
}

func (r Result[T]) Expect(s string) T {
	if r.err != nil {
		panic(s)
	}

	return r.ok
}

func (r Result[T]) ExpectErr(s string) error {
	if r.err == nil {
		panic(s)
	}

	return r.err
}

func (r Result[T]) Map(f func(T) Result[T]) Result[T] {
	if r.err != nil {
		return Result[T]{err: r.err}
	}

	return f(r.ok)
}

func (r Result[T]) MapErr(f func(error) Result[T]) Result[T] {
	if r.err == nil {
		return Result[T]{ok: r.ok}
	}

	return f(r.err)
}

func (r Result[T]) AndThen(s string) error {
	if r.err == nil {
		panic(s)
	}

	return r.err
}

func (r Result[T]) MapOr(s string) error {
	if r.err == nil {
		panic(s)
	}

	return r.err
}

func (r Result[T]) Or(rr Result[T]) Result[T] {
	if r.err == nil {
		return rr
	}

	return r
}

func (r Result[T]) OrElse(f func(Result[T]) Result[T]) Result[T] {
	if r.err != nil {
		return f(r)
	}

	return r
}

func NewOk[T any](value T) Result[T] {
	return Result[T]{ok: value}
}

func NewErr[T any](err error) Result[T] {
	return Result[T]{err: err}
}

type F[T any] func() (T, error)

func (r Result[T]) Flatten() Result[T] {
	if r.err != nil {
		return Result[T]{err: r.err}
	}

	okType := reflect.TypeOf(r.ok)
	resultType := reflect.TypeOf(Result[T]{})

	if okType.Kind() == reflect.Struct && okType.Name() == resultType.Name() && okType.PkgPath() == resultType.PkgPath() {
		okValue := reflect.ValueOf(r.ok)

		innerOk := okValue.FieldByName("ok").Interface()
		innerErr := okValue.FieldByName("err").Interface()

		if innerErr != nil {
			return Result[T]{err: innerErr.(error)}
		}

		return Result[T]{ok: innerOk.(T)}
	}

	return r
}

func (r Result[T]) IsErr() bool {
	if r.err != nil {
		return true
	}

	return false
}

func (r Result[T]) Is(a error) bool {
	if r.err == nil {
		return false
	}

	return errors.Is(r.err, a)
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(r.err)
	}

	return r.ok
}

func (r Result[T]) UnwrapErr() error {
	if r.err == nil {
		panic(r.err)
	}

	return r.err
}

func (r Result[T]) IsOk() bool {
	if r.err == nil {
		return true
	}

	return false
}

func abc(a int, b string) (int, error) {
	fmt.Print(a, b)
	return 0, errors.New(b)
}

func With[T any](f func() (T, error)) Result[T] {
	a, b := f()

	if b == nil {
		return NewOk(a)
	}

	return NewErr[T](b)
}

func From[T any](a T, b error) Result[T] {
	if b == nil {
		return NewOk(a)
	}

	return NewErr[T](b)
}

func main() {
	a := From(abc(2, ""))

	if a.IsErr() {
		fmt.Println("SOMETHING WENT WRONG")
	}

	if a.Is(errors.New("melon")) {

	}

	b := a.Unwrap()
	fmt.Println(a, b)
}
