package werr

import "runtime"

// Wrap creates an `*Wrapper` instance with the file, line and stack trace of
// the moment when it is called.
// If `StackArraySize` is `<= 0` the stack trace will not be generated.
// If `err` is `nil` it will return `nil`.
// If `err` is an `*Wrapper` it will return the same `err` variable.
// Anything else will be put in the `Original` attribute of `*Wrapper`
func Wrap(err error) error {
	if err == nil {
		return nil
	}

	if _, ok := err.(*Wrapper); ok {
		return err
	}

	var stackData []byte
	_, file, line, _ := runtime.Caller(1)
	if StackArraySize > 0 {
		stackData = make([]byte, StackArraySize)
		written := runtime.Stack(stackData, false)
		stackData = stackData[0:written]
	}

	return &Wrapper{
		Original: err,
		File:     file,
		Line:     line,
		Stack:    stackData,
	}
}

// Unwrap returns the original error inside an `*Wrapper` instance.
// If `err` is `nil` it will return `nil`.
// If `err` is an instance of `*Wrapper` it will return the value of the
// `Original` attribute.
// Anything else will be returned without any modification.
func Unwrap(err error) error {
	if err == nil {
		return nil
	}

	if wrapper, ok := err.(*Wrapper); ok {
		return wrapper.Original
	}

	return err
}
