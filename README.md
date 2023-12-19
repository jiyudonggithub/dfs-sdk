# dfs-sdk-go

#### 介绍
这 dfs 服务的go 的sdk 应用





#### 使用说明

1.  向 dfs 服务管理员申请 accesskey 
2.  下载 dfs 响应的sdk
   先初始化一个dfs client 。 
  ```
        var accessKey = "*****"
	var sercertKey = "******"
	dfsClient := client.InitClient(accessKey, sercertKey, "bucketName")
 ```
示例：

```
        #  获取文件id 的url
        var accessKey = "lM6Maad1lapad1666611"
	var sercertKey = "T2dPjPpTDSejeP2peT2fziUD5dF22F2eTdDTTjF2"
	dfsClient := client.InitClient(accessKey, sercertKey, "idata")
	url, err := dfsClient.GetPreSignedUrl("756c8331d351447e924920c5790ba532")
	fmt.Println(url)
	fmt.Println(err)
```
