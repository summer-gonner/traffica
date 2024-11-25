package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// ExecuteSearch 查询数据
func (es *Client) ExecuteSearch(index string, field string, value string) ([]*LogInfo, error) {
	// 构建查询
	query := elastic.NewMatchQuery(field, value)

	// 执行查询
	res, err := es.client.Search().
		Index(index).
		Query(query).
		Do(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to execute search: %s", err)
	}
	var ls []*LogInfo
	if res != nil && res.Hits.TotalHits.Value > 0 {
		for _, hit := range res.Hits.Hits {
			var l *LogInfo
			err = json.Unmarshal(hit.Source, &l)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			ls = append(ls, l)
		}

	}
	return ls, nil
}
