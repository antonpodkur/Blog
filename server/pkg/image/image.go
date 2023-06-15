package image

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/h2non/bimg"
)

func ImageProcessing(buffer []byte, quality int, dirname string) (string, error) {
    filename := strings.Replace(uuid.New().String(), "-", "", -1) + ".webp"
    
    converted, err := bimg.NewImage(buffer).Convert(bimg.WEBP)
    if err != nil {
        return "", err
    }

    processed, err := bimg.NewImage(converted).Process(bimg.Options{Quality: quality})
    if err != nil {
        return "", err
    }

    writeError := bimg.Write(fmt.Sprintf("./"+dirname+"/%s", filename), processed)
    if writeError != nil {
        return "", writeError
    }

    return filename, nil
}

