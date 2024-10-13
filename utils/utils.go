package utils

import (
	"crypto/tls"
	"encoding/json"
	fmt "fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
	"net/http"
	"os"
	"strings"
	"time"
)

func WriteToFile(body []byte, name string) {
	fileName := fmt.Sprintf("json/%s", name)
	_ = os.Mkdir("json", os.ModePerm)
	_ = os.WriteFile(fileName, body, os.ModePerm)
}
func WriteToJSONByFileName(body []byte, name string) {
	fileName := fmt.Sprintf("json/%s.json", name)
	_ = os.Mkdir("json", os.ModePerm)
	_ = os.WriteFile(fileName, body, os.ModePerm)
}
func WriteJSON(j interface{}, fileName string) {
	f, err := json.MarshalIndent(&j, "", " ")

	if err != nil {
		fmt.Println(err)
	} else {
		os.WriteFile(fmt.Sprintf("%s", fileName), f, 0777)
	}
}
func ReadJSONBytes(fileName string) []byte {
	fileRead, _ := os.ReadFile(fmt.Sprintf("%s", fileName))
	return fileRead
}
func ReadJSON[T interface{}](fileName string) (t T) {
	json.Unmarshal(ReadJSONBytes(fileName), &t)
	return t
}
func GetCommonHeaders() map[string]string {
	return map[string]string{
		"User-Agent": `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36`,
		"Cookie":     "",
	}
}

func GetHttpClient() *resty.Client {

	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(10 * time.Second)

	client.SetHeaders(GetCommonHeaders())
	client.SetContentLength(true)

	//client.SetProxy("socks5://127.0.0.1:7890")

	client.
		SetRetryCount(5).
		SetRetryWaitTime(10 * time.Second).
		SetDebug(false)

	return client
}

func RequestString(url string) (string, error) {
	resp, err := GetHttpClient().R().
		//EnableTrace().
		Get(url)

	if err != nil {
		return "", err
	}
	//fmt.Println("Response Info:")
	//fmt.Println("  Error      :", err)
	//fmt.Println("  Status Code:", resp.StatusCode())

	if resp.StatusCode() == http.StatusOK {
		bodyString := string(resp.Body())
		return bodyString, nil
	} else {
		panic(fmt.Sprintf("错误号码: url is %s ,status code is %d %s", url, resp.StatusCode(), string(resp.Body())))
		return "", err
	}
}
func CheckError(err error) {
	if err != nil {
		fmt.Printf("%d", err)
	}
}
func RequestThanSaveImage(url string, saveImagePath string) error {
	//fmt.Println("start download image ", url)
	client := resty.New()
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	client.SetTimeout(20 * time.Second)
	//client.SetContentLength(true)
	resp, err := client.
		SetRetryCount(4).
		SetRetryWaitTime(20 * time.Second).
		SetDebug(false).
		R().Get(url)

	if err != nil {
		return err
	}
	if resp.StatusCode() == http.StatusOK {
		err = os.WriteFile(saveImagePath, resp.Body(), 0644) // Save image
		if err != nil {
			fmt.Println("Error saving image:", saveImagePath, url, err)
			return err
		}
	} else {
		fmt.Println(fmt.Sprintf("图片返回非200 url is %s ,status code is %d %s", url, resp.StatusCode(), string(resp.Body())))
		//RequestThanSaveImage(url, saveImagePath)
	}
	return nil
}

func RequestGetDocument(url string) (*goquery.Document, error) {
	body, err := RequestString(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(body))
	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}
	return doc, nil
}

type Request struct {
	CategoryID int
	Cursor     string
	Sort       int
	IsVIP      bool
	Limit      int
}

func PostToStructInputStruct[T interface{}](url string, data interface{}, sessionId string) (t T, err error) {
	payload, err := json.Marshal(data)
	return PostToStructInputBytes[T](url, payload, sessionId)
}
func PostToStructInputBytes[T interface{}](url string, payloadBytes []byte, sessionId string) (t T, err error) {
	resp, err := GetHttpClient().R().
		//EnableTrace().
		SetHeader("Content-Type", "application/json").
		SetHeader("Cookie", "sessionid="+sessionId).
		SetBody(payloadBytes).
		Post(url)

	if err != nil {
		return t, err
	}
	//fmt.Println("Response Info:")
	//fmt.Println("  Error      :", err)
	//fmt.Println("  Status Code:", resp.StatusCode())

	if resp.StatusCode() == http.StatusOK {
		if resp.StatusCode() == http.StatusOK {
			//bodyString := string(resp.Body())
			json.Unmarshal(resp.Body(), &t)
			return t, nil
		} else {
			fmt.Println("错误号码:")
			panic(fmt.Sprintf("url is %s ,status code is %d %s", url, resp.StatusCode(), string(resp.Body())))
			return t, err
		}
		json.Unmarshal(resp.Body(), &t)
		return t, nil
	} else {
		fmt.Println("错误号码:")
		panic(fmt.Sprintf("url is %s ,status code is %d %s", url, resp.StatusCode(), string(resp.Body())))
		return t, err
	}
}
