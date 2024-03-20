package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func getCityCode(cityName string) string {
	switch cityName {
	case "北京":
		return "110105"
	case "安阳":
		return "410502"
	case "鹤壁":
		return "410611"
	default:
		return "000000"
	}
}

// Get 请求
func Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	reader := bufio.NewReader(resp.Body)
	bytes, _ := reader.ReadBytes('\n')
	respBody := string(bytes)
	return respBody
}
