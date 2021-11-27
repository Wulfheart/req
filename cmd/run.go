package cmd

import (
	"fmt"
	"github.com/Wulfheart/req/requester"
	"github.com/Wulfheart/req/str"
	"github.com/fatih/color"
	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/cobra"
	"github.com/theckman/yacspin"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var hideResult bool

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Args:  cobra.ExactValidArgs(1),
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		path := args[0]
		path, err := filepath.Abs(path)
		if err != nil {
			panic(err)
		}

		configPath := filepath.Join(path, requester.ConfigDirectoryName)
		file, _ := os.Stat(configPath)

		for file == nil {
			path = filepath.Dir(path)
			if filepath.Dir(configPath) == path {
				panic("Folder not found")
			}
			configPath = filepath.Join(path, requester.ConfigDirectoryName)
			file, _ = os.Stat(configPath)

		}

		config := requester.Config{
			ConfigPath:  configPath,
			HttpVersion: "HTTP/1.1",
		}

		var session map[string]string = make(map[string]string)
		preflight := requester.BuildRequestsFromFile(filepath.Join(config.ConfigPath, requester.PreRunFileName), config, &session, false)

		requests := requester.BuildRequestsFromFile(args[0], config, &session, true)

		requests = append(preflight, requests...)

		for _, request := range requests {
			err = request.Prepare()
			if err != nil {
				panic(err)
			}

			cfg := yacspin.Config{
				Frequency:       100 * time.Millisecond,
				CharSet:         []string{"   ", "  <", " <<", "<<<", "<< ", "<  ", "   "},
				Suffix:          fmt.Sprintf(" %s %s", getMethodColored(request.OriginalRequest.Method), request.OriginalRequest.RequestURI),
				SuffixAutoColon: true,
				StopCharacter:   "<<<",
			}
			// fmt.Println("<<<", request.OriginalRequest.Method, request.OriginalRequest.RequestURI)

			spinner, _ := yacspin.New(cfg)
			spinner.Start()

			err = request.DoRequest()
			if err != nil {
				panic(err)
			}

			err = request.DoStuffAfterTheRequest()
			if err != nil {
				panic(err)
			}

			spinner.Stop()

			var d string
			if request.TimeNeeded > time.Second {
				d = fmt.Sprintf("%.1fs", request.TimeNeeded.Seconds())
			} else {
				d = fmt.Sprintf("%dms", request.TimeNeeded.Milliseconds())
			}

			fmt.Println(">>>", getResponseCodeColored(request.Response.StatusCode), d, request.Response.Header.Get("Content-Type"))

			if request.ShowResult && !hideResult {
				var b string
				if s := request.Response.Header.Get("Content-Type"); str.StrOf(s).Lower().Upper().Lower().Contains("json") {
					f := prettyjson.NewFormatter()
					f.KeyColor = color.New(color.FgWhite)
					s, err := f.Format([]byte(request.ResponseBody))
					if err != nil {
						panic(err)
					}
					b = string(s)
				} else {
					b = request.ResponseBody
				}

				lines := strings.Split(b, "\n")

				for i, l := range lines {
					lines[i] = str.StrOf(l).Prepend("    ").String()
				}

				fmt.Println(strings.Join(lines, "\n"))
			}
			fmt.Println()

		}
	},
}

func getMethodColored(m string) string {
	switch m {
	case "GET":
		return color.New(color.FgWhite, color.BgBlue, color.Bold).Sprint(m)
	case "POST":
		return color.New(color.FgWhite, color.BgGreen, color.Bold).Sprint(m)
	case "DELETE":
		return color.New(color.BgRed, color.Bold).Sprint(m)
	default:
		return color.New(color.BgWhite, color.FgBlack, color.Bold).Sprint(m)
	}

}

func getResponseCodeColored(c int) string {
	if c < 100 {
		return color.New(color.FgHiBlue).Sprint(c)
	}
	if c < 300 {
		return color.New(color.FgHiGreen).Sprint(c)
	}
	if c < 400 {
		return color.New(color.FgHiWhite).Sprint(c)
	}
	if c < 500 {
		return color.New(color.FgHiCyan).Sprint(c)
	}
	return color.New(color.FgHiRed).Sprint(c)
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	runCmd.Flags().BoolVarP(&hideResult, "hide-response", "r", false, "Hides response body from output")

}
