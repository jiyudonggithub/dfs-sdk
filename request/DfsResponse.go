package request

type UploadFileResponse struct {
	FileUrl  string `json:"fileUrl"`  //GET请求方式的访问 URL
	FileName string `json:"fileName"` //文件名字
	Md5      string `json:"md5"`      //文件md5
	FileId   string `json:"fileId"`   //文件主键
}

type FastUploadResponse struct {
	Exist    bool   `json:"exist"`    //是否存在
	FileUrl  string `json:"fileUrl"`  //GET请求方式的访问 URL
	FileName string `json:"fileName"` //文件名字
	Md5      string `json:"md5"`      //文件md5
	FileId   string `json:"fileId"`   //文件主键
}

type UploadChunkInfoResponse struct {
	Finished       bool   `json:"finished"`       //是否完成上传（是否已经合并分片）
	FileIdentifier string `json:"fileIdentifier"` //文件md5
	UploadId       string `json:"uploadId"`       //分片上传的uploadId
	ChunkSize      uint64 `json:"chunkSize"`      //分片大小
	TotalSize      uint64 `json:"totalSize"`      //文件大小
	ChunkNum       int8   `json:"chunkNum"`       //分片数量
	ExitPartList   []int8 `json:"exitPartList"`   //已上传完的分片
}

type UploadChunkInfoResponseCode struct {
	Code int16                   `json:"code"`
	Msg  string                  `json:"msg"`
	Data UploadChunkInfoResponse `json:"data"`
}

type UploadFileResponseCode struct {
	Code int16              `json:"code"`
	Msg  string             `json:"msg"`
	Data UploadFileResponse `json:"data"`
}

type FastUploadResponseCode struct {
	Code int16              `json:"code"`
	Msg  string             `json:"msg"`
	Data FastUploadResponse `json:"data"`
}
