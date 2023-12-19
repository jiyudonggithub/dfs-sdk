package client

import (
	"dfs-sdk-go/request"
	"dfs-sdk-go/utils"
	"path/filepath"
)

type DfsClient struct {
	client     *CusHttp
	bucketName string
	AccessKey  string
	SecretKey  string
}

/*
*

	request.UploadFileRequest{
			AccessKey:  client.AccessKey,
			SecretKey:  client.SecretKey,
			BucketName: client.bucketName,
			FileName:   filepath.Base(fileName),
		}

		创建 client
*/
func InitClient(accessKey string,
	secretKey string,
	bucketName string) *DfsClient {

	client := NewHttpClient()

	return &DfsClient{client, bucketName, accessKey, secretKey}
}

func (client *DfsClient) UploadFile(fileName string) (request.UploadFileResponse, error) {
	url := utils.HOST

	form_data := map[string]string{
		"accessKey":  client.AccessKey,
		"secretKey":  client.SecretKey,
		"fileName":   filepath.Base(fileName),
		"bucketName": client.bucketName,
	}

	// 上传文件
	multipart, err := client.client.PostMultipart(url+utils.UPLOAD_FILE, fileName, form_data)
	if err != nil {
		return request.UploadFileResponse{}, err
	}
	response := request.UploadFileResponseCode{}
	err = utils.StrToStruct(string(multipart), &response)
	return response.Data, nil
}

func (client *DfsClient) DownloadFile(fileId string) ([]byte, error) {
	url := utils.HOST

	downloadRequest := request.DownloadFileRequest{
		AccessKey:  client.AccessKey,
		SecretKey:  client.SecretKey,
		BucketName: client.bucketName,
		FileId:     fileId,
	}

	multipart, err := client.client.Post(url+utils.DOWNLOAD_FILE, downloadRequest)

	return multipart, err
}

func (client *DfsClient) GetPreSignedUrl(fileId string) (request.UploadFileResponse, error) {
	url := utils.HOST
	preSignedUrl := request.GetPreSignedUrlRequest{
		AccessKey:  client.AccessKey,
		SecretKey:  client.SecretKey,
		BucketName: client.bucketName,
		FileId:     fileId,
	}
	multipart, err := client.client.Post(url+utils.PRE_SIGNED_URL, preSignedUrl)
	response := request.UploadFileResponseCode{}
	err = utils.StrToStruct(string(multipart), &response)
	return response.Data, err
}

func (client *DfsClient) CheckMd5(fileName string) (request.FastUploadResponse, error) {
	url := utils.HOST
	fileMd5, err := utils.GetFileMd5Values(fileName)
	if err != nil {
		return request.FastUploadResponse{}, err
	}
	fastUpload := request.FastUploadRequest{
		AccessKey:  client.AccessKey,
		SecretKey:  client.SecretKey,
		BucketName: client.bucketName,
		Md5:        fileMd5,
	}

	multipart, err := client.client.Post(url+utils.FAST_UPLOAD, fastUpload)
	response := request.FastUploadResponseCode{}
	err = utils.StrToStruct(string(multipart), &response)
	return response.Data, err
}
