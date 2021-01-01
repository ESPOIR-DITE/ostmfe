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

func GetPageBannerData() []PageBannerData {
	var pageBannerList []PageBannerData

	pageBanners, err := pageData_io.ReadPageBanners()
	if err != nil {
		fmt.Println(err, " error reading pageBanners")
		return pageBannerList
	}
	for _, pageBanner := range pageBanners {
		pageBannerList = append(pageBannerList, PageBannerData{pageBanner, GetBannerImage(pageBanner.BannerId)})
	}
	return pageBannerList
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
