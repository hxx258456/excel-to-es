package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
)

type CareerEduRequire struct {
	ID         int    `json:"id"`          //
	CareerCode string `json:"career_code"` // 职业编码
	CareerName string `json:"career_name"` // 职业名称
	TrendData  string `json:"trend_data"`  // 趋势数据
}

func (CareerEduRequire) Index() string {
	return "career_edu_require_index"
}

func (CareerEduRequire) Mapping() string {
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
      }
    }
  }
}
`
}

func (CareerEduRequire) GenDoc(k int, v []string) (transfor.Indexer, error) {
	return &CareerNeeds{
		ID:         k,
		CareerCode: v[0],
		CareerName: v[1],
		TrendData:  v[2],
	}, nil
}

func (c CareerEduRequire) GetId() string {
	return fmt.Sprintf("%d", c.ID)
}
