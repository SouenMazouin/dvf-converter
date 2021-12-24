/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var TxtToCsvCmd = &cobra.Command{
	Use:   "TxtToCsv",
	Short: "Convert DVF TXT files to CSV",
	Run: func(cmd *cobra.Command, args []string) {
		i, _ := cmd.Flags().GetString("in")
		o, _ := cmd.Flags().GetString("out")
		d, _ := cmd.Flags().GetString("delim")

		ConvertTxtToCsv(i, o, d)
	},
}

func processTxtLine(line, delim string) (res string) {
	res = line
	//TODO:
	// - replace "," notation at numbers by "." notation, there is probably a better way to do it
	if strings.Contains(line, ",") {
		res = strings.Replace(line, ",", ".", -1)
	}

	// TODO:
	// - leave the choice of delimiter with default value at ","
	if strings.Contains(line, delim) || strings.Contains(line, ",") {
		res = strings.Replace(res, delim, ",", -1)
	}
	return res
}

func ConvertTxtToCsv(in, out, delim string) (err error) {
	srcFile, err := os.Open(in)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if out[len(out)-4:] != ".csv" {
		out += ".csv"
	}

	dstFile, err := os.Create(out)
	if err != nil {
		return err
	}

	defer dstFile.Close()

	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)

	defer writer.Flush()

	for scanner.Scan() {
		line := processTxtLine(scanner.Text(), delim)

		_, err = fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}

	return nil
}

func init() {
	rootCmd.AddCommand(TxtToCsvCmd)
	TxtToCsvCmd.Flags().StringP("in", "i", "", "Input source file")
	TxtToCsvCmd.Flags().StringP("out", "o", "", "Output dest file")
	TxtToCsvCmd.Flags().StringP("delim", "d", "", "Current delimiter of input file")
}
