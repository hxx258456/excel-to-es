package esmodel

import (
	"errors"
	"excel-to-es/transfor"
	"fmt"
	"strconv"
)

type MajorScoreline struct {
	ID                int    `json:"id"`                 //
	UniversityId      int    `json:"university_id"`      //
	UniversityName    string `json:"university_name"`    // 大学名称
	Year              int    `json:"year"`               // 年份
	Province          string `json:"province"`           // 省份
	Category          string `json:"category"`           // 类别
	Batch             string `json:"batch"`              // 批次
	MajorName         string `json:"major_name"`         // 专业名称
	AvgScore          int    `json:"avg_score"`          // 平均分
	MinScore          int    `json:"min_score"`          // 最低分
	MaxScore          int    `json:"max_score"`          // 最高分
	MinPosition       int    `json:"min_position"`       // 最低位次
	SelectInformation string `json:"select_information"` // 选科信息
}

func (MajorScoreline) Index() string {
	return "major_scoreline_index"
}

func (MajorScoreline) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "id": {
        "type": "integer"
      },
      "university_id": {
        "type": "integer"
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
        "type": "keyword"
      },
      "avg_score": {
        "type": "integer"
      },
      "min_score": {
        "type": "integer"
      },
      "max_score": {
        "type": "integer"
      },
      "min_position": {
        "type": "integer"
      },
      "select_information": {
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

func (m MajorScoreline) GetId() string {
	return fmt.Sprintf("%d", m.ID)
}

func (MajorScoreline) GenDoc(k int, v []string) (transfor.Indexer, error) {
	if k == 0 {
		return nil, errors.New("invalid key")
	}
	if len(v) < 12 {
		fill := make([]string, 12-(len(v)))
		v = append(v, fill...)
	}
	universityId, _ := strconv.Atoi(v[0])
	year, _ := strconv.Atoi(v[2])

	avgScore, _ := strconv.Atoi(v[7])
	minScore, _ := strconv.Atoi(v[8])
	maxScore, _ := strconv.Atoi(v[9])
	minPosition, _ := strconv.Atoi(v[10])

	return &MajorScoreline{
		ID:                k,
		UniversityId:      universityId,
		UniversityName:    v[1],
		Year:              year,
		Province:          v[3],
		Category:          v[4],
		Batch:             v[5],
		MajorName:         v[6],
		AvgScore:          avgScore,
		MinScore:          minScore,
		MaxScore:          maxScore,
		MinPosition:       minPosition,
		SelectInformation: v[11],
	}, nil
}
