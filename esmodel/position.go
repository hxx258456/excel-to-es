package esmodel

import (
	"errors"
	"excel-to-es/transfor"
	"fmt"
	"strconv"
)

var _ transfor.Indexer = (*Position)(nil)

type Position struct {
	Id                int    `json:"id"`
	ProvinceCode      int    `json:"province_code"`       // 省份编码
	Province          string `json:"province"`            //省份
	Year              int    `json:"year"`                // 年份
	Subjects          string `json:"subjects"`            // 科目
	Batch             string `json:"batch"`               // 批次
	ScoreRangeLow     int    `json:"score_range_low"`     // 分数区间低
	ScoreRangeHigh    int    `json:"score_range_high"`    // 分数区间高
	PositionRangeHigh int    `json:"position_range_high"` // 位次区间高
	PositionRangeLow  int    `json:"position_range_low"`  // 位次区间低
	SameNumber        int    `json:"same_number"`         // 同分人数
}

func (Position) Index() string {
	return "position_index"
}

func (Position) Mapping() string {
	return `
	{
		"mappings": {
		  "properties": {
			"id": {
			  "type": "keyword"
			},
			"province_code": {
			  "type": "integer"
			},
			"province": {
			  "type": "keyword"
			},
			"year": {
			  "type": "integer"
			},
			"subjects": {
			  "type": "keyword"
			},
			"batch": {
			  "type": "keyword"
			},
			"score_range_low": {
			  "type": "long"
			},
			"score_range_high": {
				"type": "long"
			},
			"position_range_high": {
			  "type": "long"
			},
			"position_range_low": {
				"type": "long"
			},
			"same_number": {
			  "type": "long"
			}
		  }
		}
	  }
`
}

func (Position) GenDoc(k int, v []string) (transfor.Indexer, error) {
	if k == 0 {
		return nil, errors.New("invalid k")
	}
	provinceCode, err := strconv.Atoi(v[0])
	if err != nil {
		return nil, err
	}
	year, err := strconv.Atoi(v[2])
	if err != nil {
		return nil, err
	}

	// 分数区间
	scoreRangeLow, err := strconv.Atoi(v[5])
	if err != nil {
		return nil, err
	}
	scoreRangeHigh, err := strconv.Atoi(v[6])
	if err != nil {
		return nil, err
	}

	// 位次区间
	positionRangeHigh, err := strconv.Atoi(v[7])
	if err != nil {
		return nil, err
	}
	positionRangeLow, err := strconv.Atoi(v[8])
	if err != nil {
		return nil, err
	}

	sameCount, err := strconv.Atoi(v[9])
	if err != nil {
		return nil, err
	}
	return &Position{
		Id:                k,
		ProvinceCode:      provinceCode,
		Province:          v[1],
		Year:              year,
		Subjects:          v[3],
		Batch:             v[4],
		ScoreRangeLow:     scoreRangeLow,
		ScoreRangeHigh:    scoreRangeHigh,
		PositionRangeHigh: positionRangeHigh,
		PositionRangeLow:  positionRangeLow,
		SameNumber:        sameCount,
	}, nil
}

func (p Position) GetId() string {
	return fmt.Sprintf("%d", p.Id)
}
