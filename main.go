package bsdiff

import (
	"io"

	"github.com/jeffpc1993/go-bsdiff/diff"
	"github.com/jeffpc1993/go-bsdiff/patch"
)

func Diff(oldReader, newReader io.Reader, patchWriter io.Writer) (err error) {
	return diff.Diff(oldReader, newReader, patchWriter)
}

func Patch(oldReader io.Reader, newWriter io.Writer, patchReader io.Reader) (err error) {
	return patch.Patch(oldReader, newWriter, patchReader)
}
