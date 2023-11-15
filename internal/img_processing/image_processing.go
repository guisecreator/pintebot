package img_processing

import (
	"bytes"
	"fmt"
	"github.com/h2non/bimg"
)

func FormatImage(image_path string) (*NamedReaderImpl, error) {
	var (
		processedBuffer bytes.Buffer
		formattingErr   error
	)

	buf, err := bimg.Read(image_path)
	if err != nil {
		return nil, fmt.Errorf("read image error: %v\n", err)
	}

	newImage, err := bimg.NewImage(buf).
		Process(bimg.Options{
			Width:   512,
			Height:  512,
			Crop:    true,
			Quality: 95,
		})
	if err != nil {
		return nil, fmt.Errorf("process image error: %v\n", err)
	}

	_, err = processedBuffer.Write(newImage)
	if err != nil {
		return nil, err
	}

	processed := processedBuffer.Bytes()
	if processed == nil {
		return nil, formattingErr
	}

	newReader := bytes.NewReader(processed)

	return &NamedReaderImpl{
		reader:   newReader,
		fileName: string(processed),
	}, nil
}
