package main

import (
	"context"
	"excel-to-es/excel"
	"flag"
	"fmt"
	"github.com/olivere/elastic/v7"
)

var (
	flagEsUrl      string
	flagEsUser     string
	flagEsPassword string
	flagFilePath   string
	ctx            = context.Background()
)

func main() {
	flag.Parse()
	// es init
	es, err := elastic.NewClient(elastic.SetURL(flagEsUrl), elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetSniff(false))
	if err != nil {
		panic(fmt.Sprintf("open es client error: %v", err))
	}

	excel.ReadExcel(es, flagFilePath, ctx)
}

func init() {

	flag.StringVar(&flagEsUrl, "es_url", "http://0.0.0.0:9200", "elasticsearch url, eg: -es_url http://0.0.0.0:9200")
	flag.StringVar(&flagEsUser, "es_user", "elastic", "elasticsearch user, eg: -es_user elastic")
	flag.StringVar(&flagEsPassword, "es_password", "elastic", "elasticsearch password, eg: -es_password elastic")

	flag.StringVar(&flagFilePath, "filepath", "../../testdata/院校(1).xlsx", "需要导入的excel路径, eg: -filepath ../../院校(1).xlsx")

}
