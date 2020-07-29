/*
Copyright © 2020 ks6088ts <ks6088ts@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Package cmd ...
package cmd

import (
	"fmt"
	"os"

	"github.com/ks6088ts/sandbox-go/pkg/scraper"
	"github.com/spf13/cobra"
)

func newHatenaCmd() *cobra.Command {
	type options struct {
		url   string
		debug bool
	}

	opt := options{}
	cmd := &cobra.Command{
		Use:   "hatena",
		Short: "hatena command",
		Long:  `hatena bookmark scrape command`,
		Run: func(cmd *cobra.Command, args []string) {
			scraper := scraper.NewHatenaScraper(opt.url, opt.debug)
			results, err := scraper.Scrape()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			for _, result := range results {
				fmt.Printf("%v  \n", result)
			}
		},
	}

	cmd.PersistentFlags().StringVarP(&opt.url, "url", "u", "https://b.hatena.ne.jp/hotentry/it", "type url")
	cmd.PersistentFlags().BoolVarP(&opt.debug, "debug", "d", false, "true if debug")

	return cmd
}

func init() {
	cmd := newHatenaCmd()
	rootCmd.AddCommand(cmd)
}
