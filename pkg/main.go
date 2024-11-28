package main

import (
	"github.com/summer-gonner/traffica/pkg/es"
	"log"
)

func main() {
	// 创建客户端
	client := &es.Client{
		Address:  "http://47.101.198.49:9200", // 设置为你的 Elasticsearch 地址
		Username: "",
		Password: "",
	}

	//连接到 Elasticsearch
	err := client.Connect()
	if err != nil {
		log.Fatalf("连接 Elasticsearch 失败: %v", err)
	}
	log.Println("成功连接到 Elasticsearch")
	////
	//// 1. 同步数据到 Elasticsearch
	//logInfo := es.LogInfo{
	//	LogMessage: "\t\n\n[aefc3b1e9f3846c4839304c46fe31180]Inbound Message\n----------------------------\nAddress: http://demo-temu.jtjms-sa.com/temutrace/trace/query\nHttpMethod: POST\nQueryString: null\nEncoding: UTF-8\nContent-Type: application/json\nHeaders: {host=[demo-temu.jtjms-sa.com], x-request-id=[84679d1b72aadd50376dd6885df9d013], x-real-ip=[183.178.176.10], x-forwarded-for=[183.178.176.10], x-forwarded-host=[demo-temu.jtjms-sa.com], x-forwarded-port=[443], x-forwarded-proto=[http], x-forwarded-scheme=[http], x-scheme=[http], x-original-forwarded-for=[183.178.176.10], content-length=[30], x-original-url=[/temutrace/trace/query], x-appgw-trace-id=[4423ee7ab5749a86f06133f405c92b33], x-original-host=[demo-temu.jtjms-sa.com], sign=[ugpfiaMS-HDNM-V2LF2GVUqOpoEKmu-74oCvzfWA2Hqkq5CbP_QHNQorQfOOqyFnaI8AMRQnUXW_SyWn3OF5fw], timestamp=[1732721707731], user-agent=[Apache-HttpClient/4.5.3 (Java/17.0.9)], accept-encoding=[gzip,deflate], Content-Type=[application/json;charset=UTF-8]}\nPayload: {\"billCode\":\"UAT100041145050\"}\n----------------------------------------------",
	//	Timestamp:  time.Now().String(),
	//	Level:      "success",
	//}
	//err = client.SynchronizeData("logs", logInfo)
	//if err != nil {
	//	log.Fatalf("同步数据失败: %v", err)
	//}
	//log.Println("数据已成功同步到 Elasticsearch")
	//
	// 2. 查询 Elasticsearch 数据
	res, err := client.ExecuteSearch("logs", "logmsg", "http://demo-temu.jtjms-sa.com/temutrace/trace/query")
	if err != nil {
		log.Printf("%v", err)
	}
	for _, item := range res {
		log.Printf("查询成功，返回结果：%v", item.LogMessage)
	}
}
