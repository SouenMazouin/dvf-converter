/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// CsvToJsonCmd represents the CsvToJson command
var CsvToJsonCmd = &cobra.Command{
	Use:   "CsvToJson",
	Short: "Convert DVF CSV files to JSON",
	Run: func(cmd *cobra.Command, args []string) {
		i, _ := cmd.Flags().GetString("in")
		o, _ := cmd.Flags().GetString("out")

		ConvertCsvToJson(i, o)
	},
}

// func processCsvLine(c [][]string) (res string) {

// }

func ConvertCsvToJson(in, out string) (err error) {
	srcFile, err := os.Open(in)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if out[len(out)-5:] != ".json" {
		out += ".json"
	}

	dstFile, err := os.Create(out)
	if err != nil {
		return err
	}

	defer dstFile.Close()

	reader := csv.NewReader(srcFile)

	content, err := reader.ReadAll()
	if err != nil {
		return err
	}

	if len(content) < 1 {
		fmt.Println("Something wrong, the file maybe empty")
	}

	// TODO: refacto
	headers := append(make([]string, 0), content[0]...)

	content = content[1:]

	var buffer bytes.Buffer
	buffer.WriteString("[")
	for i, d := range content {
		buffer.WriteString("{")
		for j, y := range d {
			buffer.WriteString(`"` + headers[j] + `":`)
			_, fErr := strconv.ParseFloat(y, 32)
			_, bErr := strconv.ParseBool(y)
			if fErr == nil {
				buffer.WriteString(y)
			} else if bErr == nil {
				buffer.WriteString(strings.ToLower(y))
			} else {
				buffer.WriteString((`"` + y + `"`))
			}
			if j < len(d)-1 {
				buffer.WriteString(",")
			}

		}
		buffer.WriteString("}")
		if i < len(content)-1 {
			buffer.WriteString(",")
		}
	}
	buffer.WriteString(`]`)

	rawMessage := buffer.String()

	jsonData, err := json.Marshal(rawMessage)
	if err != nil {
		return err
	}
	fmt.Println(jsonData)

	return nil
}

func init() {
	rootCmd.AddCommand(CsvToJsonCmd)
	CsvToJsonCmd.Flags().StringP("in", "i", "", "Input source file")
	CsvToJsonCmd.Flags().StringP("out", "o", "", "Output dest file")
}
