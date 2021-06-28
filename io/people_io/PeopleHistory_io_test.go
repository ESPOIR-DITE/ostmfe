package people_io

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"ostmfe/domain/people"
	"testing"
)

////generate pdf function https://ichi.pro/fr/generation-de-pdf-dans-go-144610739774987
//func (r *RequestPdf) GeneratePDF(pdfPath string) (bool, error) {
//	t := time.Now().Unix()
//	// write whole the body
//	err1 := ioutil.WriteFile("storage/"+strconv.FormatInt(int64(t), 10)+".html", []byte(r.body), 0644)
//	if err1 != nil {
//		panic(err1)
//	}
//
//	f, err := os.Open("storage/" + strconv.FormatInt(int64(t), 10) + ".html")
//	if f != nil {
//		defer f.Close()
//	}
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	pdfg, err := wkhtmltopdf.NewPDFGenerator()
//	if err != nil {
//		os.Remove("storage/" + strconv.FormatInt(int64(t), 10) + ".html")
//		log.Fatal(err)
//	}
//
//	pdfg.AddPage(wkhtmltopdf.NewPageReader(f))
//
//	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)
//
//	pdfg.Dpi.Set(300)
//
//	err = pdfg.Create()
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = pdfg.WriteFile(pdfPath)
//	if err != nil {
//		log.Fatal(err)
//	}
//	os.Remove("storage/" + strconv.FormatInt(int64(t), 10) + ".html")
//
//	return true, nil
//}

func TestCreatePeopleHistory(t *testing.T) {
	obejct := people.PeopleHistory{"", "003003", "029342"}
	result, err := CreatePeopleHistory(obejct)
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadPeopleHistoryWithPplId(t *testing.T) {
	result, err := ReadPeopleHistoryWithPplId("HF-7c6fc5b1-259e-45f9-8c9e-f2d8d383e383")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadPeopleHistorys(t *testing.T) {
	result, err := ReadPeopleHistorys()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadPeopleImage(t *testing.T) {
	result, err := ReadPeopleImageWithPeopleId("PF-35736764-ef06-46e0-afcb-dc17363cc1e0")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadPeopleImages(t *testing.T) {
	result, err := ReadPeopleImages()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}

func TestReadCategories(t *testing.T) {
	result, err := ReadCategories()
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestDeleteCategory(t *testing.T) {
	result, err := DeleteCategory("CF-38fb1583-5246-45b9-b575-43698e345d12")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}

//Place

func TestCreatePeoplePlace(t *testing.T) {
	peoplePlace := people.PeoplePlace{"", "00001", "99383"}
	result, err := CreatePeoplePlace(peoplePlace)
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestDeletePeoplePlace(t *testing.T) {
	result, err := DeletePeoplePlace("PF-f75dba7a-2942-4c65-96dc-fada02d54c7f")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
func TestReadPeoplePlaceAllByPlaceId(t *testing.T) {
	result, err := ReadPeoplePlaceAllByPlaceId("PF-35736764-ef06-46e0-afcb-dc17363cc1e0")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}

//PeopleCategory

func TestCreatePeopleCategory(t *testing.T) {
	object := people.PeopleCategory{"", "0001", "00001", "sjhdfdjhfkhf"}
	result, err := CreatePeopleCategory(object)
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}

func TestReadPeopleCategoryWithCategoryId(t *testing.T) {
	result, err := ReadPeopleCategoryWithCategoryId("0001")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}

//func TestReadPeopleCategoriesWithPplId(t *testing.T) {
//	result,err := ReadPeopleCategoriesWithPplId("00001")
//	assert.Nil(t, err)
//	fmt.Println("result: ", result)
//}
func TestReadPeopleCategoryWithPplId(t *testing.T) {
	result, err := ReadPeopleCategoryWithPplId("00001")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}

func TestGetAggregatedPeople(t *testing.T) {
	result, err := GetAggregatedPeople("PF-3ad971d3-d032-4c33-86ea-6c72c4f628a7")
	assert.Nil(t, err)
	fmt.Println("result: ", result)
}
