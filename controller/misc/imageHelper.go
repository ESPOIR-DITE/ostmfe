package misc

import (
	"fmt"
	image2 "ostmfe/domain/image"
	"ostmfe/io/image_io"
)

func CreateImageHelper(content []byte, description string) (image2.Images, error) {
	//image
	image := image2.Images{"", content, description}
	imagePlace, err := image_io.CreateImage(image)
	if err != nil {
		fmt.Println(err, " error creating a new image")
		return image2.Images{}, err
	}
	return imagePlace, nil
}
