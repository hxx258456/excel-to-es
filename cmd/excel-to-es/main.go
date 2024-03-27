package main

import (
	"context"
	"excel-to-es/esmodel"
	"excel-to-es/old_transfor"
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
	flagReverse    bool
	flagOffset     int

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
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.University{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "position":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.Position{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "university_score_line":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.UniversityScoreLine{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "batch_line":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.BatchLine{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "early_batch":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.EarlyBatch{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "zk_major":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.ZKMajor{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "bk_major":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.BKMajor{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "major_score_line":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.MajorScoreline{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "career_major_require":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.CareerMajorRequire{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "career_needs":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.CareerNeeds{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "c_c_r":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.CareerCompanyRecruitment{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "c_pc_r":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.CareerPCRecruitment{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "c_edu_r":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.CareerEduRequire{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "c_ctype_c":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.CareerCTypeCount{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}

			case "c_company_count":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.CareerCompanyCount{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
					return err
				}
			case "career":
				esCli, err := elastic.NewClient(elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetURL(flagEsUrl), elastic.SetSniff(false))
				if err != nil {
					return err
				}
				if err := old_transfor.ReadExcel(esCli, flagFilePath, esmodel.Career{}, flagChunkSize, flagReverse, flagOffset, context.Background()); err != nil {
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
	rootCmd.Flags().BoolVarP(&flagReverse, "reverse", "r", false, "是否转义")
	rootCmd.Flags().IntVarP(&flagOffset, "offset", "o", 0, "id偏移量")
}
