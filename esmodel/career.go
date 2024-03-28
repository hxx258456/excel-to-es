package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
)

type Career struct {
	ID                   int    `json:"id"`                      // id
	OneLevelCareerCode   string `json:"one_level_career_code"`   // 一级职业编号
	OneLevelCareerName   string `json:"one_level_career_name"`   // 一级职业名称
	TwoLevelCareerCode   string `json:"two_level_career_code"`   // 二级职业编号
	TwoLevelCareerName   string `json:"two_level_career_name"`   // 二级职业名称
	ThreeLevelCareerCode string `json:"three_level_career_code"` // 三级职业编号
	ThreeLevelCareerName string `json:"three_level_career_name"` // 三级职业名称
	TwoLevelCareerDesc   string `json:"two_level_career_desc"`   // 二级职业描述
	Ranking              int    `json:"ranking"`                 // 热度
}

func (c Career) GetId() string {
	return fmt.Sprintf("%d", c.ID)
}

func (Career) Index() string {
	return "career_index"
}

func (Career) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "id": {
        "type": "integer"
      },
      "one_level_career_code": {
        "type": "keyword"
      },
      "one_level_career_name": {
        "type": "keyword"
      },
      "two_level_career_code": {
        "type": "keyword"
      },
      "two_level_career_name": {
        "type": "keyword"
      },
      "three_level_career_code": {
        "type": "keyword"
      },
      "three_level_career_name": {
        "type": "keyword"
      },
      "ranking": {
        "type": "integer"
      }
    }
  }
}
`
}

func (Career) GenDoc(k int, v []string) (transfor.Indexer, error) {
	if len(v) < 7 {
		fill := make([]string, 7-(len(v)))
		v = append(v, fill...)
	}
	return &Career{
		ID:                   k,
		OneLevelCareerCode:   v[0],
		OneLevelCareerName:   v[1],
		TwoLevelCareerCode:   v[2],
		TwoLevelCareerName:   v[3],
		ThreeLevelCareerCode: v[4],
		ThreeLevelCareerName: v[5],
		TwoLevelCareerDesc:   v[6],
	}, nil
}
