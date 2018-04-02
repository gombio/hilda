// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/gombio/hilda/config"
	ht "github.com/gombio/hilda/test"
	cmpt "github.com/gombio/hilda/test/component"
	"github.com/spf13/cobra"
	"net/http"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Verify provided hosts",
	Long: `Visits each host and run verification process.

Provide list of full /healthz URLs to visit separated with space.
Ex. http://example.com/healthz http://example2.com/healthz http://example3.com/healthz`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide a list of full /healthz URLs to visit separated with space")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, url := range args {
			ctx := ht.NewContext(url)
			rpt := ht.NewReport(url)

			//tests
			cmpt.Request(
				&http.Client{Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				}},
				&http.Client{Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
				}},
			)(ctx, rpt)
			cmpt.HTTP()(ctx, rpt)
			cmpt.Services()(ctx, rpt)

			fmt.Println(rpt.URL + " => " + rpt.Status)
			for c, f := range rpt.Components {
				if !showReport(c, cmd) {
					continue
				}
				status := "OK"
				if len(f) > 0 {
					status = "Error"
				}
				fmt.Println("=> " + c + ": " + status)
				for k, v := range f {
					fmt.Println("--> " + k + ": " + v)
				}
			}
		}
	},
}

func init() {
	config.Init(testCmd)
	RootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("hosts", "", "Comma separated list of full /healthz URLs to visit. Ex. http://example.com/healthz")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showReport(name string, cmd *cobra.Command) bool {
	flagValue := cmd.Flag(config.LongName(name))

	if flagValue == nil {
		return true
	}

	return flagValue.Value.String() == "false"
}
