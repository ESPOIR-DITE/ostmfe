package collection

import (
	"fmt"
	"ostmfe/controller/misc"
	"ostmfe/domain/collection"
	image3 "ostmfe/domain/image"
	"ostmfe/io/collection_io"
	"ostmfe/io/image_io"
)

type CollectionData struct {
	Collection    collection.CollectionHelper
	ProfileImages image3.ImagesHelper
	Images        []image3.ImagesHelper
}

func GetCollectionData(collectionId string) CollectionData {
	var collectionData CollectionData

	var images []image3.ImagesHelper
	var profileImage image3.ImagesHelper

	collections, err := collection_io.ReadCollection(collectionId)
	if err != nil {
		fmt.Println(err, " error reading collection")
		return collectionData
	}
	collectionImages, err := collection_io.ReadCollectionImgsWithCollectionId(collectionId)
	if err != nil {
		fmt.Println(err, " error reading collectionImages. This collection may not have images yet!")
	} else {
		for _, collectionImage := range collectionImages {
			if collectionImage.Description == "1" || collectionImage.Description == "profile" {
				image, err := image_io.ReadImage(collectionImage.ImageId)
				if err != nil {
					fmt.Println(err, " error reading profileImage")
				}
				profileImage = image3.ImagesHelper{image.Id, misc.ConvertingToString(image.Image), collectionImage.Id}
			}
			image, err := image_io.ReadImage(collectionImage.ImageId)
			if err != nil {
				fmt.Println(err, " error reading Images")
			}
			imageObject := image3.ImagesHelper{image.Id, misc.ConvertingToString(image.Image), collectionImage.Id}
			images = append(images, imageObject)
		}
	}

	collectionHelper := collection.CollectionHelper{collections.Id, collections.Name, collections.ProfileDescription, misc.ConvertingToString(collections.History)}
	collectionData = CollectionData{collectionHelper, profileImage, images}

	return collectionData
}
