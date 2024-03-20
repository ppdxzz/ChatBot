package main

import (
	"context"
	"fmt"
	"github.com/baidubce/bce-qianfan-sdk/go/qianfan"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func Weather(cityName string, isAll bool) string {
	// 目前仅支持部分测试城市天气查询
	cityCode := getCityCode(cityName)
	if cityCode == "000000" {
		return "目标城市暂未接入天气系统，请联系管理员。"
	}
	var parameters = "key=" + viper.GetString("amap.key") + "&city=" + cityCode
	if isAll {
		parameters += "&extensions=all"
	} else {
		parameters += "&extensions=base"
	}
	respBody := Get("https://restapi.amap.com/v3/weather/weatherInfo?" + parameters)
	if gjson.Get(respBody, "status").Int() == 1 {
		var result string
		if isAll {
			dataJson := gjson.Get(respBody, "forecasts.0").String()
			forecastJson := gjson.Get(dataJson, "casts.1").String()
			result += "省份:" + gjson.Get(dataJson, "province").String()
			result += "\n城市:" + gjson.Get(dataJson, "city").String()
			result += "\n预报日期:" + gjson.Get(forecastJson, "date").String()
			result += "\n白天天气:" + gjson.Get(forecastJson, "dayweather").String()
			result += "\n晚上天气:" + gjson.Get(forecastJson, "nightweather").String()
			result += "\n白天温度:" + gjson.Get(forecastJson, "daytemp").String()
			result += "\n晚上温度:" + gjson.Get(forecastJson, "nighttemp").String()
			result += "\n白天风向:" + gjson.Get(forecastJson, "daywind").String()
			result += "\n晚上风向:" + gjson.Get(forecastJson, "nightwind").String()
			result += "\n白天风力:" + gjson.Get(forecastJson, "daypower").String()
			result += "\n晚上风力:" + gjson.Get(forecastJson, "nightpower").String()
			result += "\n发布时间:" + gjson.Get(dataJson, "reporttime").String()
		} else {
			dataJson := gjson.Get(respBody, "lives.0").String()
			result += "省份:" + gjson.Get(dataJson, "province").String()
			result += "\n城市:" + gjson.Get(dataJson, "city").String()
			result += "\n天气:" + gjson.Get(dataJson, "weather").String()
			result += "\n气温:" + gjson.Get(dataJson, "temperature").String()
			result += "\n风向:" + gjson.Get(dataJson, "winddirection").String()
			result += "\n风力:" + gjson.Get(dataJson, "windpower").String()
			result += "\n空气湿度:" + gjson.Get(dataJson, "humidity").String()
			result += "\n发布时间:" + gjson.Get(dataJson, "reporttime").String()
		}
		return result
	}
	return "天气服务查询异常，请联系开发者"
}

func Chat(question string) string {
	qianfan.GetConfig().AccessKey = viper.GetString("qianfan.AccessKey")
	qianfan.GetConfig().SecretKey = viper.GetString("qianfan.SecretKey")
	chat := qianfan.NewChatCompletion(qianfan.WithModel("ERNIE-Bot"))
	resp, err := chat.Do(
		context.TODO(),
		&qianfan.ChatCompletionRequest{
			Messages: []qianfan.ChatCompletionMessage{
				qianfan.ChatCompletionUserMessage(question),
			},
		},
	)
	if err != nil {
		fmt.Println(err)
		return "出错了哦"
	}
	return resp.Result
}

func ZaoAn(typeName string) string {
	var url string
	key := viper.GetString("tianapi.key")
	if typeName == "ZA" {
		url = "https://apis.tianapi.com/zaoan/index?key=" + key
	} else {
		url = "https://apis.tianapi.com/wanan/index?key=" + key
	}
	respBody := Get(url)
	if gjson.Get(respBody, "code").Int() == 200 {
		return gjson.Get(respBody, "result.content").String()
	}
	return ""
}
