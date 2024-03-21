package esmodel

import (
	"excel-to-es/transfor"
	"fmt"
	"strconv"
)

type UniversityScoreLine struct {
	ID                  int    `json:"id"`                   // 唯一标识
	Year                int    `json:"year"`                 // 年份
	ProvinceCode        int    `json:"province_code"`        // 省份编码
	Province            string `json:"province"`             // 省份
	UniversityCode      string `json:"university_code"`      // 学校编码
	University          string `json:"university"`           // 学校
	EnrollmentDirection string `json:"enrollment_direction"` // 招生方向
	City                string `json:"city"`                 // 城市
	Subjects            string `json:"subjects"`             // 科类
	Batch               string `json:"batch"`                // 批次
	EnrollmentNums      int    `json:"enrollment_nums"`      // 录取人数
	MaxScore            int    `json:"max_score"`            // 最高分
	MaxPosition         int    `json:"max_position"`         // 最高名次
	MinScore            int    `json:"min_score"`            // 最低分
	MinPosition         int    `json:"min_position"`         // 最低名次
	AvgScore            int    `json:"avg_score"`            // 平均分
}

func (UniversityScoreLine) Index() string {
	return "university_score_line_index"
}

func (UniversityScoreLine) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "id": {
        "type": "integer"
      },
      "year": {
        "type": "integer"
      },
      "province_code": {
        "type": "integer"
      },
      "province": {
        "type": "keyword"
      },
      "university_code": {
        "type": "keyword"
      },
      "university": {
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
      "enrollment_direction": {
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
      "city": {
        "type": "keyword"
      },
      "subjects": {
        "type": "keyword"
      },
      "batch": {
        "type": "keyword"
      },
      "enrollment_nums": {
        "type": "integer"
      },
      "max_score": {
        "type": "integer"
      },
      "max_position": {
        "type": "integer"
      },
      "min_score": {
        "type": "integer"
      },
      "min_position": {
        "type": "integer"
      },
      "avg_score": {
        "type": "integer"
      }
    }
  }
}
`
}

func (UniversityScoreLine) GenDoc(k int, v []string) (transfor.Indexer, error) {
	if k == 0 {
		return nil, fmt.Errorf("invalid k")
	}
	year, err := strconv.Atoi(v[0])
	if err != nil {
		return nil, err
	}
	provinceCode, _ := strconv.Atoi(v[1])

	enrollmentNums, _ := strconv.Atoi(v[9])
	maxScore, _ := strconv.Atoi(v[10])
	maxPosition, _ := strconv.Atoi(v[11])
	minScore, _ := strconv.Atoi(v[12])
	minPosition, _ := strconv.Atoi(v[13])
	avgScore, _ := strconv.Atoi(v[14])
	return UniversityScoreLine{
		ID:                  k,
		Year:                year,
		ProvinceCode:        provinceCode,
		Province:            v[2],
		UniversityCode:      v[3],
		University:          v[4],
		EnrollmentDirection: v[5],
		City:                v[6],
		Subjects:            v[7],
		Batch:               v[8],
		EnrollmentNums:      enrollmentNums,
		MaxScore:            maxScore,
		MaxPosition:         maxPosition,
		MinScore:            minScore,
		MinPosition:         minPosition,
		AvgScore:            avgScore,
	}, nil
}

func (u UniversityScoreLine) GetId() string {
	return fmt.Sprintf("%d", u.ID)
}
