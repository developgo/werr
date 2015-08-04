package werr_test

import (
	"errors"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/txgruppi/werr"
)

func TestWrap(t *testing.T) {
	Convey("it should not wrap nil value", t, func() {
		err := werr.Wrap(nil)
		So(err, ShouldBeNil)
	})
	Convey("it should wrap an error", t, func() {
		original := errors.New("testing")
		err := werr.Wrap(original)
		we, ok := err.(*werr.Wrapper)
		So(ok, ShouldBeTrue)
		So(we.Original, ShouldEqual, original)
	})
	Convey("it should not wrap a Wrapper instance", t, func() {
		original := errors.New("testing")
		err := werr.Wrap(original)
		err = werr.Wrap(err)
		we, ok := err.(*werr.Wrapper)
		So(ok, ShouldBeTrue)
		So(we.Original, ShouldEqual, original)
	})
}

func TestUnwrap(t *testing.T) {
	Convey("it should ignore nil value", t, func() {
		err := werr.Unwrap(nil)
		So(err, ShouldBeNil)
	})
	Convey("it should return the same error if not wrapped", t, func() {
		original := errors.New("testing")
		err := werr.Unwrap(original)
		So(err, ShouldEqual, original)
	})
	Convey("it should return the original error if wrapped", t, func() {
		original := errors.New("testing")
		err := werr.Wrap(original)
		err = werr.Unwrap(err)
		So(err, ShouldEqual, original)
	})
}

func TestWrapper(t *testing.T) {
	Convey("it should return the original error message when calling .Error()", t, func() {
		original := errors.New("testing")
		err := werr.Wrap(original)
		So(err.Error(), ShouldEqual, original.Error())
	})
	Convey("it should return the formatted log when calling .Log()", t, func() {
		original := errors.New("testing")
		err := werr.Wrap(original)
		wrapped, ok := err.(*werr.Wrapper)
		So(ok, ShouldBeTrue)
		log, err := wrapped.Log()
		So(err, ShouldBeNil)
		So(log, ShouldContainSubstring, "werr/funcs_test.go:59 testing\ngoroutine")
	})
}

func TestVars(t *testing.T) {
	Convey("it should be possible to change the stack buffer size", t, func() {
		originalSize := werr.StackArraySize
		werr.StackArraySize = originalSize * 2
		original := errors.New("testing")
		err := werr.Wrap(original)
		wrapped, ok := err.(*werr.Wrapper)
		So(ok, ShouldBeTrue)
		So(len(wrapped.Stack), ShouldEqual, originalSize*2)
	})
	Convey("it should be possible to change the log format string", t, func() {
		werr.LogLine = "{{.File}}@{{.Line}}"
		original := errors.New("testing")
		err := werr.Wrap(original)
		wrapped, ok := err.(*werr.Wrapper)
		So(ok, ShouldBeTrue)
		log, err := wrapped.Log()
		So(err, ShouldBeNil)
		So(log, ShouldEndWith, "werr/funcs_test.go@81")
	})
}
