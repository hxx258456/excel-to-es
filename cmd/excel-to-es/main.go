package main

import (
	"context"
	"excel-to-es/esmodel"
	"excel-to-es/old_transfor"
	"excel-to-es/transfor"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
	"log"
	"os"
)

var (
	flagEsUrl      string
	flagEsUser     string
	flagEsPassword string
	flagFilePath   string
	flagType       string
	flagChunkSize  int

	rootCmd = &cobra.Command{
		Use:   "excel-to-es",
		Short: "excel数据转存elasticsearch工具",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			switch flagType {
			case "university":
				//if err := transfor.ReadExcel(flagEsUrl, flagEsUser, flagEsPassword, flagFilePath, esmodel.University{}, flagChunkSize, context.Background()); err != nil {
				//	return err
				//}

				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.University{}, flagChunkSize, context.Background()); err != nil {
					return err
				}
			case "position":
				if err := transfor.ReadExcel(flagEsUrl, flagEsUser, flagEsPassword, flagFilePath, esmodel.Position{}, flagChunkSize, context.Background()); err != nil {
					return err
				}
			case "university_score_line":
				if err := transfor.ReadExcel(flagEsUrl, flagEsUser, flagEsPassword, flagFilePath, esmodel.UniversityScoreLine{}, flagChunkSize, context.Background()); err != nil {
					return err
				}
			case "batch_line":
				if err := transfor.ReadExcel(flagEsUrl, flagEsUser, flagEsPassword, flagFilePath, esmodel.BatchLine{}, flagChunkSize, context.Background()); err != nil {
					return err
				}
			case "early_batch":
				if err := transfor.ReadExcel(flagEsUrl, flagEsUser, flagEsPassword, flagFilePath, esmodel.EarlyBatch{}, flagChunkSize, context.Background()); err != nil {
					return err
				}
			case "zk_major":
				if err := transfor.ReadExcel(flagEsUrl, flagEsUser, flagEsPassword, flagFilePath, esmodel.ZKMajor{}, flagChunkSize, context.Background()); err != nil {
					return err
				}
			case "bk_major":
				if err := transfor.ReadExcel(flagEsUrl, flagEsUser, flagEsPassword, flagFilePath, esmodel.BKMajor{}, flagChunkSize, context.Background()); err != nil {
					return err
				}
			default:
				return nil
			}
			return nil
		},
	}
)

func execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func main() {
	execute()
}

func init() {
	rootCmd.Flags().StringVarP(&flagEsUrl, "addr", "a", "http://0.0.0.0:9200", "elasticsearch url")
	if err := rootCmd.MarkFlagRequired("addr"); err != nil {
		panic(err)
	}

	rootCmd.Flags().StringVarP(&flagEsUser, "user", "u", "elastic", "elasticsearch username")
	if err := rootCmd.MarkFlagRequired("user"); err != nil {
		panic(err)
	}

	rootCmd.Flags().StringVarP(&flagEsPassword, "password", "p", "elastic", "elasticsearch user password")
	if err := rootCmd.MarkFlagRequired("password"); err != nil {
		panic(err)
	}

	rootCmd.Flags().StringVarP(&flagType, "type", "t", "university", "excel文件解析所用模型")
	if err := rootCmd.MarkFlagRequired("type"); err != nil {
		panic(err)
	}

	rootCmd.Flags().StringVarP(&flagFilePath, "filepath", "f", "../../testdata/院校(1).xlsx", "需要导入的excel文件路径")
	if err := rootCmd.MarkFlagRequired("filepath"); err != nil {
		panic(err)
	}

	rootCmd.Flags().IntVarP(&flagChunkSize, "chunk", "c", 500, "文件切片大小")
}
