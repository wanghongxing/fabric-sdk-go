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
	"encoding/json"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/hyperledger/fabric-sdk-go/fabric-cli/common"
	fabricClient "github.com/hyperledger/fabric-sdk-go/fabric-client"
	"github.com/spf13/pflag"
)

/*
	flags.String(common.PeerFlag, "", "The URL of the peer on which to query the chaincode, e.g. localhost:7051")
	flags.StringVar(&common.ChannelID, common.ChannelIDFlag, common.ChannelID, "The channel ID")
	flags.StringVar(&common.ChaincodeID, common.ChaincodeIDFlag, "", "The chaincode ID")
	flags.StringVar(&common.Args, common.ArgsFlag, common.Args, "The args in JSON format. Example: {\"Args\":\"invoke\",\"arg1\"}")
	flags.IntVar(&common.Iterations, common.IterationsFlag, 1, "The number of times to invoke the chaincode.")
	flags.Int64Var(&common.SleepTime, common.SleepFlag, int64(100), "The number of milliseconds to sleep between invocations of the chaincode.")
*/
type QueryArgs struct {
	ChannelID   string   `json:"channelId"`
	ChaincodeID string   `json:"chaincodeId"`
	Args        []string `json:"args"`
	//Iterations int `json:"iterations"`  //1 is better in production
	//SleepTime int64 `json:"sleepTime"` //only used when iterations is set
}

type queryAction struct {
	common.ActionImpl
	numInvoked uint32
	done       chan bool
}

func NewQueryAction(args *QueryArgs) (*queryAction, error) {
	action, flags := &queryAction{done: make(chan bool)}, &pflag.FlagSet{}

	flags.StringVar(&common.ChannelID, common.ChannelIDFlag, args.ChannelID, "The channel ID")
	flags.StringVar(&common.ChaincodeID, common.ChaincodeIDFlag, args.ChaincodeID, "The chaincode ID")
	flags.StringVar(&common.Args, common.ArgsFlag, common.GetMarshalArgs(common.ArgStruct{args.Args}), "The args in JSON format. Example: {\"Args\":[\"invoke\",\"arg1\",\"arg2\"]}")
	flags.IntVar(&common.Iterations, common.IterationsFlag, 1, "The number of times to invoke the chaincode.")
	flags.Int64Var(&common.SleepTime, common.SleepFlag, int64(100), "The number of milliseconds to sleep between invocations of the chaincode.")

	err := action.Initialize(flags)
	return action, err
}

func (action *queryAction) Execute() (string, error) {
	chain, err := action.NewChain()
	if err != nil {
		return "", fmt.Errorf("Error initializing chain: %v", err)
	}

	argBytes := []byte(common.Args)
	args := &common.ArgStruct{}
	err = json.Unmarshal(argBytes, args)
	if err != nil {
		return "", fmt.Errorf("Error unmarshaling JSON arg string: %v", err)
	}

	if common.Iterations > 1 {
		go action.queryMultiple(chain, args.Args, common.Iterations)

		completed := false
		for !completed {
			select {
			case <-action.done:
				completed = true
			case <-time.After(5 * time.Second):
				fmt.Printf("... completed %d out of %d\n", action.numInvoked, common.Iterations)
			}
		}
	} else {
		response, err := action.doQuery(chain, args.Args)
		if err != nil {
			fmt.Printf("Error invoking chaincode: %v\n", err)
			return "", err
		} else {
			fmt.Printf("***** Response: %s\n", response)
			return string(response), nil
		}
	}
	return "", errors.New("Error invoking chaincode...")
}

func (action *queryAction) queryMultiple(chain fabricClient.Chain, args []string, iterations int) {
	fmt.Printf("Querying CC %d times ...\n", common.Iterations)
	for i := 0; i < common.Iterations; i++ {
		if _, err := action.doQuery(chain, args); err != nil {
			fmt.Printf("Error invoking chaincode: %v\n", err)
		}
		if (i+1) < common.Iterations && common.SleepTime > 0 {
			time.Sleep(time.Duration(common.SleepTime) * time.Millisecond)
		}
		atomic.AddUint32(&action.numInvoked, 1)
	}
	fmt.Printf("Completed %d queries\n", common.Iterations)
	action.done <- true
}

func (action *queryAction) doQuery(chain fabricClient.Chain, args []string) ([]byte, error) {
	common.Logger.Infof("Invoking chaincode: %s or channel: %s, with args: [%v]\n", common.ChaincodeID, common.ChannelID, args)

	response, err := common.QueryChaincode(chain, action.Peers(), common.ChaincodeID, common.ChannelID, args)
	if err != nil {
		return nil, err
	}

	return response, nil
}
