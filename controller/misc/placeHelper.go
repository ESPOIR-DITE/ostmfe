package misc

import (
	"fmt"
	"ostmfe/domain/image"
	"ostmfe/domain/place"
	"ostmfe/io/place_io"
)

type PlaceGalleryImages struct {
	Gallery      image.GaleryHelper
	PlaceGallery place.PlaceGallery
}

func GetPlaceGallery(placeId string) []PlaceGalleryImages {
	var GalleryImagesList []PlaceGalleryImages

	placeGalleryImages, err := place_io.ReadAllByPlaceGallery(placeId)
	if err != nil {
		fmt.Println(err, "error reading Place Gallery")
		return GalleryImagesList
	}
	for _, placeGalleryImage := range placeGalleryImages {
		GalleryImagesList = append(GalleryImagesList, PlaceGalleryImages{GetGalleryImage(placeGalleryImage.GalleryId), placeGalleryImage})
	}
	return GalleryImagesList
}
