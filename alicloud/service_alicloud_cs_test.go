package alicloud

import (
	"fmt"
	"github.com/terraform-providers/terraform-provider-alicloud/alicloud/connectivity"
	"testing"
)

func TestCsService_GetPermanentToken(t *testing.T) {
	config := connectivity.Config{
		AccessKey:"",
		SecretKey:"",
	}
	client := connectivity.AliyunClient{
		//RegionId:  "cn-hongkong",
		AccessKey: "",
		SecretKey: "",
		Config:    &config,
	}
	s := CsService{client:&client}
	token, _ := s.GetPermanentToken("")
	fmt.Println(token)
}