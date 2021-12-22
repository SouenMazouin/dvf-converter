/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dvf-converter",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		i, _ := cmd.Flags().GetString("in")
		o, _ := cmd.Flags().GetString("out")
		d, _ := cmd.Flags().GetString("delim")

		ConvertFile(i, o, d)
	},
}

func processLine(line, delim string) (res string) {
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

func ConvertFile(in, out, delim string) (err error) {
	srcFile, err := os.Open(in)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(out + ".csv")
	if err != nil {
		return err
	}
	defer dstFile.Close()

	scanner := bufio.NewScanner(srcFile)
	writer := bufio.NewWriter(dstFile)

	defer writer.Flush()

	for scanner.Scan() {
		line := processLine(scanner.Text(), delim)

		_, err = fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.Flags().StringP("in", "i", "", "Input source file")
	rootCmd.Flags().StringP("out", "o", "", "Output dest file")
	rootCmd.Flags().StringP("delim", "d", "", "Current delimiter of input file")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dvf-converter.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".dvf-converter" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".dvf-converter")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
