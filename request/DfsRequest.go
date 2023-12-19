package request

type UploadFileRequest struct {
	AccessKey  string `json:"accessKey"`
	SecretKey  string `json:"secretKey"`
	BucketName string `json:"bucketName"` //存储桶
	FileName   string `json:"fileName"`   //文件名字
}

type DownloadFileRequest struct {
	AccessKey  string `json:"accessKey"`
	SecretKey  string `json:"secretKey"`
	BucketName string `json:"bucketName"` //存储桶
	FileId     string `json:"fileId"`     //文件主键
}

type FastUploadRequest struct {
	AccessKey  string `json:"accessKey"`
	SecretKey  string `json:"secretKey"`
	BucketName string `json:"bucketName"` //存储桶
	Md5        string `json:"md5"`        //文件md5
}

type GetPreSignedUrlRequest struct {
	AccessKey  string `json:"accessKey"`
	SecretKey  string `json:"secretKey"`
	BucketName string `json:"bucketName"` //存储桶
	FileId     string `json:"fileId"`     //文件主键
}

type CreateUploadChunkRequest struct {
	AccessKey  string `json:"accessKey"`
	SecretKey  string `json:"secretKey"`
	BucketName string `json:"bucketName"` //存储桶
	Identifier string `json:"identifier"` //文件md5
	FileName   string `json:"fileName"`   //文件名字
	ChunkSize  uint64 `json:"chunkSize"`  //分片大小
	TotalSize  uint64 `json:"totalSize"`  //分片大小
}

type UploadChunkRequest struct {
	AccessKey  string `json:"accessKey"`
	SecretKey  string `json:"secretKey"`
	BucketName string `json:"bucketName"` //存储桶
	Identifier string `json:"identifier"` //文件md5
	PartNumber int8   `json:"partNumber"` //文件名字
}

type UploadChunkInfoRequest struct {
	AccessKey  string `json:"accessKey"`
	SecretKey  string `json:"secretKey"`
	BucketName string `json:"bucketName"` //存储桶
	Identifier string `json:"identifier"` //文件md5
}

type MergeUploadChunkRequest struct {
	AccessKey  string `json:"accessKey"`
	SecretKey  string `json:"secretKey"`
	BucketName string `json:"bucketName"` //存储桶
	Identifier string `json:"identifier"` //文件md5
}
