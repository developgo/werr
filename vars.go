package werr

var (
	// StackArraySize defines the number of bytes which will be allocated for the
	// stack trace.
	// Any value `<= 0` will result in not creating the stack trace.
	StackArraySize = 512

	// LogLine is the format used by `*Wrapper.Log()`.
	// It should be a valid `text/template` string.
	// The data passed to the `Execute` method of the template is the `*Wrapper`
	// instance, which means that you can use all `*Wrapper` methods and
	// attributes inside the template.
	LogLine = "{{.File}}:{{.Line}} {{.Original.Error}}\n{{printf \"%s\" .Stack}}\n\n"
)
