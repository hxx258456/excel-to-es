package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
)

type ZKMajor struct {
	Id         int    `json:"id"`         //
	Subject    string `json:"subject"`    // 学科
	Discipline string `json:"discipline"` // 学类
	Major      string `json:"major"`      // 专业
}

func (ZKMajor) Index() string {
	return "zk_major_index"
}

func (ZKMajor) Mapping() string {
	return `
{
  "mappings": {
	"properties": {
	  "id": {
		"type": "integer"
	  },
	  "subject": {
		"type": "keyword"
	  },
	  "discipline": {
		"type": "keyword"
	  },
	  "major": {
		"type": "keyword"
	  }
	}
  }
}
`
}

func (ZKMajor) GenDoc(k int, v []string) (transfor.Indexer, error) {
	return &ZKMajor{
		Id:         k,
		Subject:    v[0],
		Discipline: v[1],
		Major:      v[2],
	}, nil
}

func (zk ZKMajor) GetId() string {
	return fmt.Sprintf("%d", zk.Id)
}
