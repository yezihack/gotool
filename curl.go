package too

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/cast"
)

// curl 是网络请求工具, 提供 get, post.json, post.form

// curl 发起 get请求
func Get(uri string, params map[string]interface{}, headers map[string]string, timeout ...time.Duration) (result []byte, err error) {
	// 判断
	if uri == "" {
		err = errors.New("url is empty")
		return
	}
	// 创建一个 http 客户端
	cli := &http.Client{}
	buf := strings.NewReader("")
	if len(params) > 0 {
		val := url.Values{}
		for k, v := range params {
			val.Add(k, cast.ToString(v))
		}
		buf.Reset(val.Encode())
	}
	// 写入 uri 请求信息
	req, err := http.NewRequest(http.MethodGet, uri, buf)
	if err != nil {
		return
	}
	// 设置超时
	_timeout := time.Second * 5
	if len(timeout) > 0 {
		_timeout = timeout[0]
	}
	ctx, cancel := context.WithTimeout(context.Background(), _timeout)
	defer cancel()
	req.WithContext(ctx)
	// 设置 header
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	// 发起请求
	resp, err := cli.Do(req)
	if err != nil {
		return
	}
	// 关闭连接
	defer resp.Body.Close()
	// 读取 body
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// curl 支持POST form表单形式
// timeout default 5second
func PostForm(uri string, params map[string]interface{}, headers map[string]string, timeout ...time.Duration) (result []byte, err error) {
	// 判断
	if uri == "" {
		err = errors.New("url is empty")
		return
	}
	// 创建一个 http 客户端
	cli := &http.Client{}
	values := url.Values{}
	for k, v := range params {
		if v != nil {
			values.Set(k, cast.ToString(v))
		}
	}
	// 写入 post 请求数据
	req, err := http.NewRequest(http.MethodPost, uri, strings.NewReader(values.Encode()))
	if err != nil {
		return
	}
	// 设置超时
	_timeout := time.Second * 5
	if len(timeout) > 0 {
		_timeout = timeout[0]
	}
	ctx, cancel := context.WithTimeout(context.Background(), _timeout)
	defer cancel()
	req.WithContext(ctx)
	// 设置 header
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := cli.Do(req)
	if err != nil {
		return
	}
	// 必须关闭
	defer resp.Body.Close()
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// curl 支持POST json
// timeout default 5second
func PostJson(uri string, params map[string]interface{}, headers map[string]string, timeout ...time.Duration) (result []byte, err error) {
	// 判断
	if uri == "" {
		err = errors.New("url is empty")
		return
	}
	// 创建一个 http 客户端
	cli := &http.Client{}
	// 数据打包
	data, err := json.Marshal(params)
	if err != nil {
		return
	}
	// 写入 post 请求数据
	req, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(data))
	if err != nil {
		return
	}
	// 设置超时
	_timeout := time.Second * 5
	if len(timeout) > 0 {
		_timeout = timeout[0]
	}
	ctx, cancel := context.WithTimeout(context.Background(), _timeout)
	defer cancel()
	req = req.WithContext(ctx)
	// 设置 header
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	req.Header.Set("Content-Type", "application/json")
	// 发起 http 请求
	resp, err := cli.Do(req)
	if err != nil {
		return
	}
	// 必须关闭
	defer resp.Body.Close()
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}
