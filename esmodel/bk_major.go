package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
)

type BKMajor struct {
	Id             int    `json:"id"`              //
	Subject        string `json:"subject"`         // 学科
	SubjectCode    string `json:"subject_code"`    // 学科编码
	Discipline     string `json:"discipline"`      // 学类
	DisciplineCode string `json:"discipline_code"` // 学类编码
	Major          string `json:"major"`           // 专业
	MajorCode      string `json:"major_code"`      // 专业编码
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
	  "subject_code": {
		"type": "keyword"
	  },
	  "discipline": {
		"type": "keyword"
	  },
	  "discipline_code": {
		"type": "keyword"
	  },
	  "major": {
		"type": "keyword"
	  },
	  "major_code": {
		"type": "keyword"
	  }
	}
  }
}
`
}

func (BKMajor) GenDoc(k int, v []string) (transfor.Indexer, error) {
	return &BKMajor{
		Id:             k,
		Subject:        v[0],
		SubjectCode:    v[1],
		Discipline:     v[2],
		DisciplineCode: v[3],
		Major:          v[4],
		MajorCode:      v[5],
	}, nil
}

func (bk BKMajor) GetId() string {
	return fmt.Sprintf("%d", bk.Id)
}
