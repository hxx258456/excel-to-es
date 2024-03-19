package main

import (
	"context"
	"excel-to-es/esmodel"
	"excel-to-es/transfor"
	"log"
	"os"

	"github.com/olivere/elastic/v7"
	"github.com/spf13/cobra"
)

var (
	flagEsUrl      string
	flagEsUser     string
	flagEsPassword string
	flagFilePath   string
	flagType       string
	rootCmd        = &cobra.Command{
		Use:   "excel-to-es",
		Short: "excel数据转存elasticsearch工具",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			escli, err := elastic.NewClient(elastic.SetURL(flagEsUrl), elastic.SetBasicAuth(flagEsUser, flagEsPassword), elastic.SetSniff(false))
			if err != nil {
				return err
			}
			switch flagType {
			case "university":
				if err := transfor.ReadExcel(escli, flagFilePath, esmodel.University{}, context.Background()); err != nil {
					return err
				}
			case "position":
				if err := transfor.ReadExcel(escli, flagFilePath, esmodel.Position{}, context.Background()); err != nil {
					return err
				}
			default:
				return nil
			}
			return nil
		},

		// Uncomment the following line if your bare application
		// has an action associated with it:
		// Run: func(cmd *cobra.Command, args []string) { },
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
	rootCmd.MarkFlagRequired("url")

	rootCmd.Flags().StringVarP(&flagEsUser, "user", "u", "elastic", "elasticsearch username")
	rootCmd.MarkFlagRequired("user")

	rootCmd.Flags().StringVarP(&flagEsPassword, "password", "p", "elastic", "elasticsearch user password")
	rootCmd.MarkFlagRequired("password")

	rootCmd.Flags().StringVarP(&flagType, "type", "t", "university", "excel文件解析所用模型")
	rootCmd.MarkFlagRequired("type")

	rootCmd.Flags().StringVarP(&flagFilePath, "filepath", "f", "../../testdata/院校(1).xlsx", "需要导入的excel文件路径")
	rootCmd.MarkFlagRequired("filepath")
}
