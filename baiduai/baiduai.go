package baiduai

import (
	"encoding/json"
	"errors"

	"github.com/mlogclub/simple"
	"github.com/tidwall/gjson"
	"gopkg.in/resty.v1"
)

func GetTags(title, content string) *AiTags {
	if title == "" || content == "" {
		return nil
	}
	data := make(map[string]interface{})
	data["title"] = title
	data["content"] = simple.Substr(content, 0, 10000)

	bytesData, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	url := "https://aip.baidubce.com/rpc/2.0/nlp/v1/keyword?charset=UTF-8&access_token=" + GetToken()
	response, err := resty.R().SetBody(string(bytesData)).Post(url)
	if err != nil {
		return nil
	}

	tags := &AiTags{}
	err = json.Unmarshal(response.Body(), tags)
	if err != nil {
		return nil
	}
	return tags
}

func GetCategories(title, content string) *AiCategories {
	if title == "" || content == "" {
		return nil
	}

	data := make(map[string]interface{})
	data["title"] = title
	data["content"] = simple.Substr(content, 0, 10000)

	bytesData, err := json.Marshal(data)
	if err != nil {
		return nil
	}

	url := "https://aip.baidubce.com/rpc/2.0/nlp/v1/topic?charset=UTF-8&access_token=" + GetToken()
	response, err := resty.R().SetBody(string(bytesData)).Post(url)
	if err != nil {
		return nil
	}

	categories := &AiCategories{}
	err = json.Unmarshal(response.Body(), categories)
	if err != nil {
		return nil
	}
	return categories
}

func GetNewsSummary(title, content string, maxSummaryLen int) (string, error) {
	if title == "" || content == "" {
		return "", errors.New("标题或内容为空")
	}
	if maxSummaryLen <= 0 {
		maxSummaryLen = 256
	}

	data := make(map[string]interface{})
	data["title"] = title
	data["content"] = simple.Substr(content, 0, 3000)
	data["max_summary_len"] = maxSummaryLen

	bytesData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	url := "https://aip.baidubce.com/rpc/2.0/nlp/v1/news_summary?charset=UTF-8&access_token=" + GetToken()
	response, err := resty.R().SetBody(string(bytesData)).Post(url)
	if err != nil {
		return "", err
	}
	ret := gjson.Get(string(response.Body()), "summary")
	return ret.String(), nil
}
