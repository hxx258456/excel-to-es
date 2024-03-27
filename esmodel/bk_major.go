package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
)

type BKMajor struct {
	Id         int    `json:"id"`         //
	Subject    string `json:"subject"`    // 学科
	Discipline string `json:"discipline"` // 学类
	Major      string `json:"major"`      // 专业
}

func (BKMajor) Index() string {
	return "bk_major_index"
}

func (BKMajor) Mapping() string {
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

func (BKMajor) GenDoc(k int, v []string) (transfor.Indexer, error) {
	return &BKMajor{
		Id:         k,
		Subject:    v[0],
		Discipline: v[1],
		Major:      v[2],
	}, nil
}

func (bk BKMajor) GetId() string {
	return fmt.Sprintf("%d", bk.Id)
}
