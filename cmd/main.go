package main

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"poker/model"
	"strconv"
	"time"
)

const (
	urlDomain = "https://1024.hellogroup.com"
	UserToken = "7f4afd39-727e-4d4a-9e60-e76d79ac4d78"
)

var client = &http.Client{
	Timeout: time.Second * 60,
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 300 * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          150,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConnsPerHost:   100,
	},
}

func main() {

}

// applyRouter 报名接口
func applyRouter(ctx context.Context) (string, error) {
	url := urlDomain + "/card/ready?userToken=" + UserToken
	request, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		respBody, _ := ioutil.ReadAll(resp.Body) //true/flase
		return string(respBody), nil
	}
	return "", errors.New("err apply router")
}

// cardInfo 获取整体⽐赛信息接口
func cardInfo(ctx context.Context, roundNum int) (model.CardInfoResp, error) {
	cardInfoResp := model.CardInfoResp{}
	url := urlDomain + "/card/info?userToken=" + UserToken + "&startTime=" +
		strconv.FormatInt(time.Now().Add(-time.Hour).Unix() * 1000, 10) + "&endTime=" + strconv.FormatInt(time.Now().Unix() * 1000, 10)
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := client.Do(request)
	if err != nil {
		return cardInfoResp, err
	}
	if resp.StatusCode == 200 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(respBody, &cardInfoResp)
		if err != nil {
			return cardInfoResp, err
		}
		return cardInfoResp, nil
	}
	return cardInfoResp, errors.New("err cardInfo")
}

// roundInfo 获取某场比赛信息接口
func roundInfo(ctx context.Context, roundNum int) (model.RoundInfoResp, error) {
	roundResp := model.RoundInfoResp{}
	url := urlDomain + "/card/roundInfo?userToken=" + UserToken + "&roundNum=" + strconv.Itoa(roundNum)
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	resp, err := client.Do(request)
	if err != nil {
		return roundResp, err
	}
	if resp.StatusCode == 200 {
		respBody, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(respBody, &roundResp)
		if err != nil {
			return roundResp, err
		}
		return roundResp, nil
	}
	return roundResp, errors.New("err roundInfo")
}

// cardOperate 获取某场比赛信息接口
func cardOperate(ctx context.Context, roundNum int, cardsInfo string) (string, error) {
	url := urlDomain + "/card/operate?userToken=" + UserToken + "&roundNum=" +
		strconv.Itoa(roundNum) + "&cardsInfo=" + cardsInfo
	request, err := http.NewRequestWithContext(ctx, "POST", url, nil)
	resp, err := client.Do(request)
	if err != nil {
		return "", err
	}
	if resp.StatusCode == 200 {
		respBody, _ := ioutil.ReadAll(resp.Body) //ture/false
		return string(respBody), nil
	}
	return "", errors.New("err cardOperate")
}
