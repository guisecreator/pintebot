package img_processing

import (
	"github.com/disintegration/imaging"
	"github.com/mymmrac/telego"
	"log"
)

func FormatImageForTelegram(
	update telego.Update,
	path string,
) (*telego.InputFile, error) {

	src, err := imaging.Open(path)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	src = imaging.CropAnchor(src, 300, 300, imaging.Center)

	dstImage128 := imaging.Resize(src, 128, 128,
		imaging.ResampleFilter{
			Support: 342,
			Kernel: func(f float64) float64 {
				return f
			},
		})

	err = imaging.Save(dstImage128, "testdata/out_example.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	return &telego.InputFile{}, nil
}
