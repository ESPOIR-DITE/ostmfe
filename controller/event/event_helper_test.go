package event

import (
	"fmt"
	"testing"
)

func TestGetGroupsData(t *testing.T) {
	groupData := GetGroupsData("EF-4470cdf3-88d6-4f1f-a3ac-05151efedbb5")
	fmt.Println(groupData)
}
