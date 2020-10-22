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
	"path/filepath"

	"github.com/labstack/gommon/log"
	"github.com/sclevine/agouti"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Qiita ...
type Qiita struct {
	User     string
	Password string
	URL      string
}

// Setting ...
type Setting struct {
	Browser string
}

// Task ...
type Task struct {
	Xpath string
	URL   string
}

type chromeDriverCmdConfig struct {
	Qiita   Qiita
	Setting Setting
	Tasks   []Task
}

func loginQiita(page *agouti.Page, config *chromeDriverCmdConfig) error {
	if err := page.Navigate(config.Qiita.URL); err != nil {
		log.Printf("Failed to navigate page %v", config.Qiita.URL)
		return err
	}
	log.Info(page.Title())
	page.Screenshot(filepath.Join("outputs", "Screenshot01.png"))

	if err := page.FindByID("identity").Fill(config.Qiita.User); err != nil {
		log.Printf("Failed to fill identity %v", config.Qiita.User)
		return err
	}
	if err := page.FindByID("password").Fill(config.Qiita.Password); err != nil {
		log.Printf("Failed to fill password %v", config.Qiita.Password)
		return err
	}
	if err := page.FindByClass("loginSessionsForm_submit").Submit(); err != nil {
		log.Printf("Failed to submit")
		return err
	}
	return nil
}

func getDriver(browser string) *agouti.WebDriver {
	if browser == "headless" {
		return agouti.ChromeDriver(
			agouti.ChromeOptions("args", []string{
				"--headless",             // headlessモードの指定
				"--window-size=1280,800", // ウィンドウサイズの指定
			}),
			// agouti.Debug,
		)
	}
	return agouti.ChromeDriver()
}

// chromeDriverCmd represents the chromeDriver command
var chromeDriverCmd = &cobra.Command{
	Use:   "chromeDriver",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var config chromeDriverCmdConfig
		if err := viper.Unmarshal(&config); err != nil {
			log.Print("Failed to unmarshal config")
			os.Exit(1)
		}

		// log.Printf("config: %v", config)
		// os.Exit(0)

		driver := getDriver(config.Setting.Browser)
		if err := driver.Start(); err != nil {
			log.Print("Failed to start driver")
			os.Exit(1)
		}
		defer driver.Stop()

		page, err := driver.NewPage()
		if err != nil {
			log.Print("Failed to start page")
			driver.Stop()
			os.Exit(1)
		}

		// if err := loginQiita(page, &config); err != nil {
		// 	driver.Stop()
		// 	os.Exit(1)
		// }

		for _, task := range config.Tasks {
			if err := page.Navigate(task.URL); err != nil {
				log.Printf("Failed to navigate page %v", task.URL)
				driver.Stop()
				os.Exit(1)
			}
			title, err := page.Title()
			if err != nil {
				log.Printf("Failed to get title %v", task.URL)
				driver.Stop()
				os.Exit(1)
			}
			fmt.Printf("[%v](%v)\n", title, task.URL)
			items := page.AllByXPath(task.Xpath)
			itemsCount, err := items.Count()
			if err != nil {
				log.Printf("Failed to get items")
				driver.Stop()
				os.Exit(1)
			}
			for i := 0; i < itemsCount; i++ {
				if text, err := items.At(i).Text(); err == nil {
					fmt.Printf("%d: %v\n", i, text)
				}
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(chromeDriverCmd)
}
