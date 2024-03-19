package esmodel

import (
	"excel-to-es/transfor"
	"strconv"
)

type University struct {
	Code           string  `json:"code"`            //编码
	Name           string  `json:"name"`            //名字
	Province       string  `json:"province"`        //省份
	City           string  `json:"city"`            //城市
	Category       string  `json:"category"`        //类别
	Nature         string  `json:"nature"`          //性质
	Belong         string  `json:"belong"`          //隶属
	Feature        string  `json:"feature"`         //特色
	Ranking        int     `json:"ranking"`         //排名
	CompositeIndex float32 `json:"composite_index"` //排名
	Heat           int     `json:"heat"`            // 热度
	Description    string  `json:"description"`     // 简介
}

func (University) Index() string {
	return "university_index"
}

func (University) Mapping() string {
	return `
{
  "mappings": {
    "properties": {
      "code": {
        "type": "keyword"
      },
      "name": {
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
      "province": {
        "type": "keyword"
      },
      "city": {
        "type": "keyword"
      },
      "category": {
        "type": "keyword"
      },
      "nature": {
        "type": "keyword"
      },
      "belong": {
        "type": "keyword"
      },
      "feature": {
        "type": "text",
        "analyzer": "comma",
        "search_analyzer": "comma"
      },
      "ranking": {
        "type": "integer"
      },
      "composite_index": {
        "type": "float"
      },
      "heat": {
        "type": "integer"
      },
      "description": {
        "type": "text"
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

func (University) GenDoc(k int, v []string) (transfor.Indexer, error) {
	ranking, err := strconv.Atoi(v[8])
	if err != nil {
		return nil, err
	}
	compositeIndex, err := strconv.ParseFloat(v[9], 32)
	if err != nil {
		return nil, err
	}
	heat, err := strconv.Atoi(v[10])
	if err != nil {
		return nil, err
	}
	return &University{
		Code:           v[0],
		Name:           v[1],
		Province:       v[2],
		City:           v[3],
		Category:       v[4],
		Nature:         v[5],
		Belong:         v[6],
		Feature:        v[7],
		Ranking:        ranking,
		CompositeIndex: float32(compositeIndex),
		Heat:           heat,
		Description:    v[11],
	}, nil
}

func (u University) GetId() string {
	return u.Code
}
