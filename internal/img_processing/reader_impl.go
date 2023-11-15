package img_processing

import (
	"github.com/mymmrac/telego/telegoapi"
	"io"
)

type NamedReader interface {
	Read(p []byte) (n int, err error)
	Name() string
}

type NamedReaderImpl struct {
	reader   io.Reader
	fileName string
}

func (n NamedReaderImpl) Read(p []byte) (number int, err error) {
	return n.reader.Read(p)
}

func (n NamedReaderImpl) Name() string {
	return n.fileName
}

func (n NamedReaderImpl) GetReader() telegoapi.NamedReader {
	return nil
}
