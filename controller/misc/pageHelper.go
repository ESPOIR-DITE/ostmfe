package misc

import (
	"fmt"
	"ostmfe/domain/pageData"
	"ostmfe/io/pageData_io"
)

type PageBannerData struct {
	PageBanner pageData.PageBanner
	Image      string
}

func GetPageBannerData(data pageData.PageData) pageData.BannerImageHelper {
	var pageBannerList pageData.BannerImageHelper
	if data.BannerId == "" {
		return pageBannerList
	}
	banner, err := pageData_io.ReadBannerN(data.BannerId)
	if err != nil {
		fmt.Println(err, " error reading pageBanners")
		return pageBannerList
	}
	return banner
}

//This method will take a bannerId and returns a string containing an image.
func GetBannerImage(bannerId string) string {
	bannerObject, err := pageData_io.ReadBanner(bannerId)
	if err != nil {
		fmt.Println(err, " error reading Banner")
		return ""
	}
	return ConvertingToString(bannerObject.Image)
}
