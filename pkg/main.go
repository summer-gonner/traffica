package main

import (
	"github.com/summer-gonner/traffica/pkg/es"
	"log"
	"time"
)

func main() {
	// 创建客户端
	client := &es.Client{
		Address: "http://47.101.198.49:9200", // 设置为你的 Elasticsearch 地址
	}

	// 连接到 Elasticsearch
	err := client.Connect()
	if err != nil {
		log.Fatalf("连接 Elasticsearch 失败: %v", err)
	}
	log.Println("成功连接到 Elasticsearch")

	// 1. 同步数据到 Elasticsearch
	logInfo := es.LogInfo{
		LogMessage: "[09066b070d8b45eabd880e44a462fed6]清关提单日志信息,消息监听原始报文={\"headers\":{\"traceId\":\"a813894e13b4454f8f68a36f05241cc3\",\"amqp_receivedDeliveryMode\":\"PERSISTENT\",\"amqp_receivedExchange\":\"customs-clearance-mawb-log-save\",\"amqp_deliveryTag\":1,\"deliveryAttempt\":1,\"amqp_consumerQueue\":\"customs-clearance-mawb-log-save.ordermq\",\"messageId\":728868126040555523,\"amqp_redelivered\":false,\"amqp_receivedRoutingKey\":\"customs-clearance-mawb-log-save\",\"id\":\"e77ab795-f842-0d98-d956-649f926e9679\",\"amqp_consumerTag\":\"amq.ctag-QImmDlBr8EYrkYf00IXZOw\",\"contentType\":{\"concrete\":true,\"parameters\":{},\"subtype\":\"json\",\"type\":\"application\",\"wildcardSubtype\":false,\"wildcardType\":false},\"timestamp\":1732514442964},\"payload\":{\"createTime\":\"2024-11-25 09:00:42\",\"masterWaybillNo\":\"MASTER1732514442094\",\"mawbId\":728868125331718168,\"operationType\":\"保存\",\"operationTypeCode\":\"SAVE\",\"operationUser\":\"system\",\"operationUserCode\":\"system\",\"payload\":\"{\\\"arrivePortCode\\\":\\\"110\\\",\\\"bigBagList\\\":[{\\\"bigBagNo\\\":\\\"BIG1732514439731\\\",\\\"masterWaybillNo\\\":\\\"MASTER1732514442094\\\",\\\"mawbId\\\":728868125331718168,\\\"packageList\\\":[{\\\"bigBagId\\\":728868125444964369,\\\"bigBagNo\\\":\\\"BIG1732514439731\\\",\\\"masterWaybillNo\\\":\\\"MASTER1732514442094\\\",\\\"mawbId\\\":728868125331718168,\\\"packageId\\\":688649114367307819,\\\"providerOrderId\\\":\\\"LPCNG11634884396544\\\",\\\"sourceCode\\\":\\\"D445\\\",\\\"sourceName\\\":\\\"CAINIAO\\\",\\\"trackingNo\\\":\\\"UTE720000000059\\\"}],\\\"sourceCode\\\":\\\"D445\\\",\\\"sourceName\\\":\\\"CAINIAO\\\"}],\\\"bigBagQuantity\\\":10,\\\"buyerRegion\\\":\\\"KSA\\\",\\\"customerCode\\\":\\\"J0086024511\\\",\\\"customsWaybillId\\\":\\\"CB202417634339\\\",\\\"departPortCode\\\":\\\"5314\\\",\\\"masterWaybillNo\\\":\\\"MASTER1732514442094\\\",\\\"realWeight\\\":\\\"100000.0\\\",\\\"sourceCode\\\":\\\"D445\\\",\\\"sourceName\\\":\\\"CAINIAO\\\",\\\"transportCode\\\":\\\"AIR_TRANSPORT\\\",\\\"transportName\\\":\\\"航空运输\\\"}\"}}",
		Timestamp:  time.Now().String(),
		Level:      "success",
	}
	err = client.SynchronizeData("logs", logInfo)
	if err != nil {
		log.Fatalf("同步数据失败: %v", err)
	}
	log.Println("数据已成功同步到 Elasticsearch")

	// 2. 查询 Elasticsearch 数据
	res, err := client.ExecuteSearch("logs", "logmsg", "09066b070d8b45eabd880e44a462fed6")

	for _, item := range res {
		log.Printf("查询成功，返回结果：%v", item.Level)
		log.Printf("查询成功，返回结果：%v", item.LogMessage)
		log.Printf("查询成功，返回结果：%v", item.Timestamp)
	}
}
