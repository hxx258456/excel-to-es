package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
	"strconv"
)

type Enrollment struct {
	ID             int    `json:"id"`              //
	UniversityId   string `json:"university_id"`   //
	UniversityName string `json:"university_name"` // 大学名称
	Year           int    `json:"year"`            // 年份
	Province       string `json:"province"`        // 省份
	Category       string `json:"category"`        // 类别
	Batch          string `json:"batch"`           // 批次
	MajorName      string `json:"major_name"`      // 专业名称
	PlanNum        int    `json:"plan_num"`        // 计划招生
	Education      string `json:"education"`       // 学制
	Fees           int    `json:"fees"`            // 学费
	Course         string `json:"course"`          // 学科
	SelectInfo     string `json:"select_info"`     // 选科信息
}

func (e Enrollment) GetId() string {
	return fmt.Sprintf("%d", e.ID)
}

func (Enrollment) Index() string {
	return "enrollment_index"
}

func (Enrollment) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "id": {
        "type": "integer"
      },
      "university_id": {
        "type": "keyword"
      },
      "university_name": {
        "type": "keyword"
      },
      "year": {
        "type": "integer"
      },
      "province": {
        "type": "keyword"
      },
      "category": {
        "type": "keyword"
      },
      "batch": {
        "type": "keyword"
      },
      "major_name": {
        "type": "text",
        "analyzer": "ik_max_word",
        "search_analyzer": "ik_max_word",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      },
      "plan_num": {
        "type": "integer"
      },
      "education": {
        "type": "keyword"
      },
      "fees": {
        "type": "integer"
      },
      "course": {
        "type": "keyword"
      },
      "select_info": {
        "type": "text",
        "analyzer": "ik_max_word",
        "search_analyzer": "ik_max_word",
        "fields": {
          "keyword": {
            "type": "keyword",
            "ignore_above": 256
          }
        }
      }
    }
  }
}
`
}

func (Enrollment) GenDoc(k int, v []string) (transfor.Indexer, error) {
	if len(v) < 12 {
		fill := make([]string, 12-len(v))
		v = append(v, fill...)
	}
	year, _ := strconv.Atoi(v[2])
	planNum, _ := strconv.Atoi(v[7])
	fees, _ := strconv.Atoi(v[9])
	return &Enrollment{
		ID:             k,
		UniversityId:   v[0],
		UniversityName: v[1],
		Year:           year,
		Province:       v[3],
		Category:       v[4],
		Batch:          v[5],
		MajorName:      v[6],
		PlanNum:        planNum,
		Education:      v[8],
		Fees:           fees,
		Course:         v[10],
		SelectInfo:     v[11],
	}, nil
}
