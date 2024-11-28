package main

import (
	"fmt"
	"github.com/summer-gonner/traffica/pkg/es"
	"log"
	"regexp"
)

func main() {
	// 创建客户端
	client := &es.Client{
		Address:  "http://47.101.198.49:9200", // 设置为你的 Elasticsearch 地址
		Username: "",
		Password: "",
	}

	// 连接到 Elasticsearch
	err := client.Connect()
	if err != nil {
		log.Fatalf("连接 Elasticsearch 失败: %v", err)
	}
	log.Println("成功连接到 Elasticsearch")

	var eqbs []es.EsSearchBody
	a1 := &es.EsSearchBody{
		Field: "logmsg",
		Value: "http://demo-temu.jtjms-sa.com/temutrace/trace/query",
	}
	a2 := &es.EsSearchBody{
		Field: "logmsg",
		Value: "POST",
	}
	eqbs = append(eqbs, *a1)
	eqbs = append(eqbs, *a2)

	// 2. 查询 Elasticsearch 数据
	res, err := client.ExecuteSearch(eqbs)
	if err != nil {
		log.Printf("%v", err)
	}
	for _, item := range res {
		log.Printf("查询成功，返回结果：%v", item.LogMessage)
		value, err := extractFieldValue(item.LogMessage, "Address")
		if err != nil {
			log.Printf("Error extracting 'Content-Type': %v", err)
			continue
		}
		log.Printf("匹配到的值: %v", value)
	}
}

// 从原始文本中提取字段值
func extractFieldValue(input string, fieldName string) (string, error) {
	// 构建正则表达式，动态插入用户输入的字段名
	// 修改后的正则表达式：不使用前瞻断言，直接匹配字段名和值
	re := regexp.MustCompile(fmt.Sprintf(`(?i)%s:\s*(.*?)(?:\n|$)`, regexp.QuoteMeta(fieldName)))

	// 查找匹配项
	match := re.FindStringSubmatch(input)

	// 如果没有找到匹配项，返回错误
	if len(match) < 2 {
		return "", fmt.Errorf("字段 '%s' 未找到", fieldName)
	}

	// 返回字段值
	return match[1], nil
}
