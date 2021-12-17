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

package login

import (
	"encoding/json"
	"github.com/mimiro-io/datahub-cli/internal/config"
	"github.com/mimiro-io/datahub-cli/internal/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var CopyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copy a login profile",
	Long: `Copies an existing login profile to a new one, keeping its values. For example:
mim login copy my-alias --to new-alias
`,
	Run: func(cmd *cobra.Command, args []string) {
		pterm.EnableDebugMessages()

		alias, err := cmd.Flags().GetString("alias")
		utils.HandleError(err)

		if alias == "" && len(args) > 0 {
			alias = args[0]
		}

		to, err := cmd.Flags().GetString("to")
		utils.HandleError(err)

		if alias == "" || to == "" {
			pterm.Error.Println("You need both an existing alias and a to alias")
		}

		server, _ := cmd.Flags().GetString("server")
		audience, _ := cmd.Flags().GetString("audience")
		err = copy(alias, to, server, audience)
		utils.HandleError(err)

		pterm.Println()
	},
}

func init() {
	CopyCmd.Flags().StringP("alias", "a", "", "An alias value to copy from")
	CopyCmd.Flags().String("to", "", "An alias value to copy to")
	CopyCmd.Flags().StringP("server", "s", "", "Server to replace existing with")
	CopyCmd.Flags().String("audience", "", "Audience to replace existing with")
}

func copy(from string, to string, server string, audience string) error {
	data := &config.Config{}
	err := config.Load(from, data)
	if err != nil {
		return err
	}

	/*home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	db, err := bolt.Open(home+"/.mim/conf.db", 0666, &bolt.Options{Timeout: 1 * time.Second})
	defer db.Close()

	data := &config.Config{}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("logins"))
		res := b.Get([]byte(from))
		if res == nil {
			return errors.New("alias not found")
		}
		err := json.Unmarshal(res, data)
		if err != nil {
			data = nil
			return err
		}
		return nil
	})*/

	if server != "" {
		data.Server = server
	}
	if audience != "" {
		data.Audience = audience
	}

	p, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return config.WriteValue(to, p)

	/*err = db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("logins"))
		if err != nil {
			return err
		}
		return b.Put([]byte(to), p)
	})

	return nil*/

}
