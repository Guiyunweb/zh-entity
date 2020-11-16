package utils

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
)

const secretId = ""
const secretKey = ""

type Res struct {
	Response Response `json:"Response"`
}

type Response struct {
	RecognizeStatus int64  `json:"RecognizeStatus"`
	RequestId       string `json:"RequestId"`
	Seq             int64  `json:"Seq"`
	SessionUuid     string `json:"SessionUuid"`
	Source          string `json:"Source"`
	SourceText      string `json:"SourceText"`
	Target          string `json:"Target"`
	TargetText      string `json:"TargetText"`
}

func Translate(sourceText string) (string, error) {
	credential := common.NewCredential(secretId, secretKey)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tmt.tencentcloudapi.com"
	client, _ := tmt.NewClient(credential, "ap-shanghai", cpf)

	request := tmt.NewTextTranslateRequest()

	request.SourceText = common.StringPtr(sourceText)
	request.Source = common.StringPtr("zh")
	request.Target = common.StringPtr("en")
	request.ProjectId = common.Int64Ptr(0)

	response, err := client.TextTranslate(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return "", nil
	}
	if err != nil {
		panic(err)
		return "", err
	}
	fmt.Println("%s", response.ToJsonString())
	var res Res
	if err := json.Unmarshal([]byte(response.ToJsonString()), &res); err != nil {
		return "", err
	}
	return res.Response.TargetText, err
}
