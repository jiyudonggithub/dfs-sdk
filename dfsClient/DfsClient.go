package dfsClient

import (
	"github.com/jiyudonggithub/dfs-sdk/request"
	"github.com/jiyudonggithub/dfs-sdk/utils"
	"mime/multipart"
	"path/filepath"
)

type DfsClient struct {
	client     *CusHttp
	bucketName string
	AccessKey  string
	SecretKey  string
	Host       string
}

/*
*
创建 dfsClient
*/
func InitClient(accessKey string,
	secretKey string,
	bucketName string,
	host string) *DfsClient {

	client := NewHttpClient()

	return &DfsClient{client, bucketName, accessKey, secretKey, host}
}

func (client *DfsClient) UploadFile(fileName string) (request.UploadFileResponse, error) {
	url := client.Host

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

func (client *DfsClient) UploadMultipartFile(file multipart.File, fileName string) (request.UploadFileResponse, error) {
	url := client.Host

	form_data := map[string]string{
		"accessKey":  client.AccessKey,
		"secretKey":  client.SecretKey,
		"fileName":   fileName,
		"bucketName": client.bucketName,
	}

	// 上传文件
	multipart, err := client.client.PostMultipartFile(url+utils.UPLOAD_FILE, file, fileName, form_data)
	if err != nil {
		return request.UploadFileResponse{}, err
	}
	response := request.UploadFileResponseCode{}
	err = utils.StrToStruct(string(multipart), &response)
	return response.Data, nil
}

func (client *DfsClient) DownloadFile(fileId string) ([]byte, error) {
	url := client.Host

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
	url := client.Host
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
	url := client.Host
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
