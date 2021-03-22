// Copyright 2021 MIMIRO AS
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"fmt"
	"os"

	"github.com/mimiro-io/datahub-cli/internal/login"
	"github.com/mimiro-io/datahub-cli/internal/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var Login2Cmd = &cobra.Command{
	Use:   "login",
	Short: "Log in to the datahub",
	Long: `Log in to the datahub, or add and use login profiles.

Example:
	mim login "<profile>"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
			os.Exit(0)
		}
		out, _ := cmd.Flags().GetBool("out")
		if out {
			pterm.DisableOutput = true
		}

		token := login.UseLogin(args[0])

		if out {
			fmt.Printf("%s\n", token)
		} else {
			login.UpdateConfig(args[0])
		}

	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// jobsCmd.PersistentFlags().String("foo", "", "A help for foo")
	Login2Cmd.Flags().Bool("out", false, "Export the token to stdout instead of login in")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	Login2Cmd.AddCommand(login.AddCmd)
	Login2Cmd.AddCommand(login.ListCmd)
	Login2Cmd.AddCommand(login.UseCmd)
	Login2Cmd.AddCommand(login.DeleteCmd)
	Login2Cmd.AddCommand(login.CopyCmd)

	Login2Cmd.SetHelpFunc(func(command *cobra.Command, strings []string) {

		pterm.Println()
		result := utils.RenderMarkdown(command, "resources/doc-login.md")
		pterm.Println(result)
	})

}
