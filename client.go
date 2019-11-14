package mob_smssdk

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const verifyURL = "https://webapi.sms.mob.com/sms/verify"
const contentType = "application/x-www-form-urlencoded;charset=UTF-8"
const status = "status"

type MobSmsSdkClient struct {
	HttpClient *http.Client
	VerifyURL  string
	AppKey     string
}

func NewMobSmsSdk(appKey string) *MobSmsSdkClient {
	result := &MobSmsSdkClient{AppKey: appKey}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	result.HttpClient = &http.Client{Transport: tr}
	result.VerifyURL = verifyURL
	return result
}

func (sdk *MobSmsSdkClient) Verify(phoneNumber, zone, code string) (err error) {
	reqBody := fmt.Sprintf("appkey=%s&amp;phone=%s&amp;zone=%s&amp;code=%s", sdk.AppKey, phoneNumber, zone, code)
	resp, err := sdk.HttpClient.Post(sdk.VerifyURL, contentType, strings.NewReader(reqBody))
	if nil != err {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return
	}
	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)
	if nil != err {
		return
	}
	//fmt.Printf("%+v\n", result)
	iStatus, ok := result[status]
	if !ok {
		return errors.New("SmsSdk Server错误")
	}
	status, ok := iStatus.(float64)
	if !ok {
		return errors.New("SmsSdk Server回复参数错误")
	}
	if 200 != int(status) {
		return ErrorCode(status)
	}
	return
}
