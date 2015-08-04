package werr

import (
	"bytes"
	"text/template"
)

// Wrapper is the struct which holds the `Original` error and the data related
// to this error.
// You should not create an instance of `*Wrapper` by yourself, you should call
// the `Wrap` function, which will set all the values for you automatically.
type Wrapper struct {
	Original error
	File     string
	Line     int
	Stack    []byte
}

func (w *Wrapper) Error() string {
	return w.Original.Error()
}

// Log returns an string which is created by `LogLine` execution.
// If can return an error if any template-related function returns an error.
func (w *Wrapper) Log() (string, error) {
	tmpl, err := template.New("").Parse(LogLine)
	if err != nil {
		return "", err
	}
	var b bytes.Buffer
	err = tmpl.Execute(&b, w)
	if err != nil {
		return "", err
	}
	return b.String(), nil
}
