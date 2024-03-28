package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
)

type UniversityMajor struct {
	ID             int    `json:"id"`              // 1
	UniversityCode string `json:"university_code"` //编码唯一
	UniversityName string `json:"university_name"` // 学校名
	FacultyId      string `json:"faculty_id"`      // 院系id
	FacultyName    string `json:"faculty_name"`    // 院系名称
	Website        string `json:"website"`         // 网站
	MajorCode      string `json:"major_code"`      // 专业编码
	MajorName      string `json:"major_name"`      // 专业名称
	FeatureTags    string `json:"feature_tags"`    // 特色标签
	Batch          string `json:"batch"`           // 批次
}

func (m UniversityMajor) GetId() string {
	return fmt.Sprintf("%d", m.ID)
}

func (UniversityMajor) Index() string {
	return "university_major_index"
}

func (UniversityMajor) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "id": {
        "type": "integer"
      },
      "university_code": {
        "type": "keyword"
      },
      "university_name": {
        "type": "keyword"
      },
      "faculty_id": {
        "type": "keyword"
      },
      "faculty_name": {
        "type": "keyword"
      },
      "website": {
        "type": "keyword"
      },
      "major_code": {
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
      "feature_tags": {
        "type": "text",
        "analyzer": "comma",
        "search_analyzer": "comma"
      },
      "batch": {
        "type": "keyword"
      }
    }
  },
  "settings": {
    "analysis": {
      "analyzer": {
        "comma": {
          "type": "pattern",
          "pattern": ","
        }
      }
    }
  }
}
`
}

func (UniversityMajor) GenDoc(k int, v []string) (transfor.Indexer, error) {
	if len(v) < 9 {
		fill := make([]string, 9-len(v))
		v = append(v, fill...)
	}
	return &UniversityMajor{
		ID:             k,
		UniversityCode: v[0],
		UniversityName: v[1],
		FacultyId:      v[2],
		FacultyName:    v[3],
		Website:        v[4],
		MajorCode:      v[5],
		MajorName:      v[6],
		FeatureTags:    v[7],
		Batch:          v[8],
	}, nil
}
