package gql

import (
	"bytes"
	"embed"
	"errors"
	"io"
	"io/fs"
)

//go:embed schemas/*.gql
var schemas embed.FS

var errSchemaRead = errors.New("could not read schema")

type schemaLoader struct {
	bytes.Buffer
}

func (l *schemaLoader) walkFunc(path string, dirEntry fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

	if dirEntry.IsDir() {
		return nil
	}

	f, err := schemas.Open(path)
	if err != nil {
		return err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	_, _ = l.Write(b)

	return nil
}

// GetSchemaData returns the embedded schema file contents as an array of bytes.
// On failure, it will return an error.
func GetSchemaData() ([]byte, error) {
	var buf bytes.Buffer

	loader := &schemaLoader{buf}

	if err := fs.WalkDir(schemas, ".", loader.walkFunc); err != nil {
		return []byte{}, err
	}

	if loader.Len() == 0 {
		return []byte{}, errSchemaRead
	}

	return loader.Bytes(), nil
}
