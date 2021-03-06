/*
Copyright SecureKey Technologies Inc. All Rights Reserved.


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

package chaincode

import (
	"github.com/hyperledger/fabric-sdk-go/fabric-cli/common"
	"github.com/spf13/cobra"
)

var chaincodeCmd = &cobra.Command{
	Use:   "chaincode",
	Short: "Chaincode commands",
	Long:  "Chaincode commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

// Cmd returns the chaincode command
func Cmd() *cobra.Command {
	flags := chaincodeCmd.Flags()
	flags.StringVar(&common.ChannelID, common.ChannelIDFlag, common.ChannelID, "The channel ID")

	chaincodeCmd.AddCommand(getInstallCmd())
	chaincodeCmd.AddCommand(getInstantiateCmd())
	chaincodeCmd.AddCommand(getInvokeCmd())
	chaincodeCmd.AddCommand(getQueryCmd())
	chaincodeCmd.AddCommand(getGetInfoCmd())

	return chaincodeCmd
}
