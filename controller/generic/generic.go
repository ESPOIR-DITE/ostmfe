package generic

import (
	"fmt"
	"ostmfe/io/image_io"
)

func GetImageTypeId(imageTypeName string) string {
	imageType, err := image_io.ReadImageTypeWithName(imageTypeName)
	if err != nil {
		fmt.Println(err, " error reading imageType")
		return ""
	}
	return imageType.Id
}
