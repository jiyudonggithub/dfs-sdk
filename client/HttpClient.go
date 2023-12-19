package client

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// 初始化连接池配置
var httpTrans = &http.Transport{
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).DialContext,
	ForceAttemptHTTP2:     true,
	MaxIdleConns:          100,                                     //所有Host的连接池最大空闲连接数
	MaxIdleConnsPerHost:   2,                                       //每个Host的连接池最大空闲连接数
	MaxConnsPerHost:       4,                                       //每个Host的连接池最大连接数
	IdleConnTimeout:       time.Duration(90000) * time.Millisecond, //空闲连接保留时间
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}

// 初始化连接池
var httpClient = &http.Client{
	Timeout:   30 * time.Second,
	Transport: httpTrans,
}

type CusHttp struct {
	Header map[string]string
}

func NewHttpClient() *CusHttp {

	return &CusHttp{Header: map[string]string{}}
}

// 发送GET请求
// url：         请求地址
func (cus *CusHttp) Get(url string) []byte {

	// resp, err := httpClient.Get(url)

	//初始化请求
	reqest, err := http.NewRequest("GET", url, nil)

	//增加Header选项
	if cus.Header != nil && len(cus.Header) > 0 {
		for k, vla := range cus.Header {
			reqest.Header.Add(k, vla)
		}
	}

	if err != nil {
		panic(err)
	}

	resp, err := httpClient.Do(reqest)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return result
}

// 发送POST请求 默认使用
// url：         请求地址
// data：        POST请求提交的数据
func (cus *CusHttp) Post(url string, data interface{}) ([]byte, error) {
	cus.Header["Content-Type"] = "application/json"
	return cus.post(url, data)
}

// 发送POST 表单
// url：         请求地址
// data：        POST请求提交的数据
func (cus *CusHttp) PostForm(url string, data interface{}) ([]byte, error) {
	cus.Header["Content-Type"] = "application/x-www-form-urlencoded"
	return cus.post(url, data)
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
func (cus *CusHttp) post(url string, data interface{}) ([]byte, error) {

	jsonStr, _ := json.Marshal(data)
	//初始化请求
	reqest, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	//增加Header选项
	if cus.Header != nil && len(cus.Header) > 0 {
		for k, vla := range cus.Header {
			reqest.Header.Add(k, vla)
		}
	}

	if err != nil {
		panic(err)
	}

	resp, err := httpClient.Do(reqest)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)
	return result, err
}

func (cus *CusHttp) PostMultipartValue(url string, data map[string]string) ([]byte, error) {

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	// 参数2 fileType (普通参数)
	for k, vla := range data {
		fileWriter2, err := bodyWriter.CreateFormField(k)
		if err != nil {
			return nil, err
		}
		_, errs2 := fileWriter2.Write([]byte(vla))
		if errs2 != nil {
			return nil, err
		}
	}

	// 一定要记着关闭
	err := bodyWriter.Close()
	if err != nil {
		return nil, err
	}
	//发送post请求
	req, err := http.NewRequest("POST", url, bodyBuf)
	if err != nil {
		return nil, err
	}
	//添加头文件
	if cus.Header != nil && len(cus.Header) > 0 {
		for k, vla := range cus.Header {
			req.Header.Add(k, vla)
		}
	}

	//req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	//获取返回值
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (cus *CusHttp) PostMultipart(url string, filePath string, data map[string]string) ([]byte, error) {

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	fileWriter1, err := bodyWriter.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return nil, err
	}
	//把文件流写入到缓冲区里去
	_, err1 := io.Copy(fileWriter1, f)
	if err1 != nil {
		return nil, err
	}
	// 参数2 fileType (普通参数)
	for k, vla := range data {
		fileWriter2, err := bodyWriter.CreateFormField(k)
		if err != nil {
			return nil, err
		}
		_, errs2 := fileWriter2.Write([]byte(vla))
		if errs2 != nil {
			return nil, err
		}
	}

	// 一定要记着关闭
	err = bodyWriter.Close()
	if err != nil {
		return nil, err
	}
	//发送post请求
	req, err := http.NewRequest("POST", url, bodyBuf)
	if err != nil {
		return nil, err
	}
	//添加头文件
	if cus.Header != nil && len(cus.Header) > 0 {
		for k, vla := range cus.Header {
			req.Header.Add(k, vla)
		}
	}
	//添加头文件
	//req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	//获取返回值
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (cus *CusHttp) PostMultipartChunk(url string, chunkBytes []byte, data map[string]string) ([]byte, error) {

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter1, err := bodyWriter.CreateFormFile("file", filepath.Base("chunk"))
	if err != nil {
		return nil, err
	}
	//把文件流写入到缓冲区里去
	_, err1 := io.Copy(fileWriter1, bytes.NewReader(chunkBytes))
	if err1 != nil {
		return nil, err
	}
	// 参数2 fileType (普通参数)
	for k, vla := range data {
		fileWriter2, err := bodyWriter.CreateFormField(k)
		if err != nil {
			return nil, err
		}
		_, errs2 := fileWriter2.Write([]byte(vla))
		if errs2 != nil {
			return nil, err
		}
	}

	// 一定要记着关闭
	err = bodyWriter.Close()
	if err != nil {
		return nil, err
	}
	//发送post请求
	req, err := http.NewRequest("POST", url, bodyBuf)
	if err != nil {
		return nil, err
	}
	//添加头文件
	if cus.Header != nil && len(cus.Header) > 0 {
		for k, vla := range cus.Header {
			req.Header.Add(k, vla)
		}
	}
	//添加头文件
	//req.Header.Set("Content-Type", "multipart/form-data")
	req.Header.Set("Content-Type", bodyWriter.FormDataContentType())
	//获取返回值
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
