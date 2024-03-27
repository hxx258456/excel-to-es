package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
	"strconv"
)

type CareerMajorRequire struct {
	ID                int     `json:"id"`                   //
	CareerCode        string  `json:"career_code"`          // 职业编码
	CareerName        string  `json:"career_name"`          // 职业名称
	MajorCode         string  `json:"major_code"`           // 专业编码
	MajorName         string  `json:"major_name"`           // 专业名称
	EduLevel          string  `json:"edu_level"`            // 学历要求
	PercentOfMajor    float64 `json:"percent_of_major"`     // 专业占比
	PercentOfMajorStr string  `json:"percent_of_major_str"` // 专业占比例
	IsUp              string  `json:"is_up"`                // 是否上涨
}

func (CareerMajorRequire) Index() string {
	return "career_major_require_index"
}

func (CareerMajorRequire) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "id": {
        "type": "integer"
      },
      "career_code": {
        "type": "keyword"
      },
      "career_name": {
        "type": "keyword"
      },
      "major_code": {
        "type": "keyword"
      },
      "edu_level": {
        "type": "keyword"
      },
      "percent_of_major": {
        "type": "double"
      },
      "percent_of_major_str": {
        "type": "keyword"
      },
      "is_up": {
        "type": "keyword"
      }
    }
  }
}
`
}

func (CareerMajorRequire) GenDoc(k int, v []string) (transfor.Indexer, error) {
	percent, _ := strconv.ParseFloat(v[5], 64)
	return &CareerMajorRequire{
		ID:                k,
		CareerCode:        v[0],
		CareerName:        v[1],
		MajorCode:         v[2],
		MajorName:         v[3],
		EduLevel:          v[4],
		PercentOfMajor:    percent,
		PercentOfMajorStr: v[6],
		IsUp:              v[7],
	}, nil
}

func (c CareerMajorRequire) GetId() string {
	return fmt.Sprintf("%d", c.ID)
}
