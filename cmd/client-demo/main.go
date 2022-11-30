/*
Copyright AppsCode Inc. and Contributors.

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

package main

import (
	"fmt"
	"os"

	"go.bytebuilders.dev/license-verifier/client"
	_ "go.bytebuilders.dev/license-verifier/client"
)

func main() {
	cc, err := client.NewClient("https://api.appscode.ninja", os.Getenv("APPSCODE_NINJA_TOKEN"), "4a650d06-8dd3-4cfd-951b-32331b3b2fc7")
	if err != nil {
		panic(err)
	}
	lic, contract, err := cc.AcquireLicense([]string{"kubedb-enterprise"})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(lic))
	fmt.Println(contract.ID)
	fmt.Println(contract.StartTimestamp)
	fmt.Println(contract.ExpiryTimestamp)
}
