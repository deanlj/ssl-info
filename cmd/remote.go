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
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Show certificate information, about a remote service or website.",
	Long:  `remote will poll an endpoint and return interesting cert information.`,
	Run: func(cmd *cobra.Command, args []string) {
		address += ":" + port

		// I think this is the best way to determine allowed versions...
		sslVersions := map[string]uint16{
			"SSL30": 0x0300,
			"TLS10": 0x0301,
			"TLS11": 0x0302,
			"TLS12": 0x0303,
		}

		var sslVersionsSupported []string

		for k, v := range sslVersions {
			conf := &tls.Config{
				InsecureSkipVerify: skipVerify,
				MinVersion:         v,
				MaxVersion:         v,
			}

			conn, err := tls.Dial("tcp", address, conf)

			if err == nil {
				sslVersionsSupported = append(sslVersionsSupported, k)
				conn.Close()
			}
		}

		conf := &tls.Config{
			InsecureSkipVerify: skipVerify,
		}

		conn, err := tls.Dial("tcp", address, conf)
		if err != nil {
			log.Println(err)
			return
		}

		defer conn.Close()

		interesting := map[string]interface{}{
			"DNS Names":              conn.ConnectionState().PeerCertificates[0].DNSNames,
			"NotAfter":               conn.ConnectionState().PeerCertificates[0].NotAfter,
			"NotBefore":              conn.ConnectionState().PeerCertificates[0].NotBefore,
			"Issuer Organization":    conn.ConnectionState().PeerCertificates[0].Issuer.Organization,
			"Signature Algorithm":    conn.ConnectionState().PeerCertificates[0].SignatureAlgorithm.String(),
			"Supported SSL Versions": sslVersionsSupported,
		}

		jsonResponse := json.NewEncoder(os.Stdout)
		jsonResponse.SetIndent("", "  ")
		jsonResponse.Encode(interesting)
	}}

var (
	address    string
	port       string
	skipVerify bool
)

func init() {
	rootCmd.AddCommand(remoteCmd)

	cobra.EnableCommandSorting = true
	remoteCmd.Flags().StringVarP(&address, "address", "a", "", "address of endpoint")
	remoteCmd.Flags().StringVarP(&port, "port", "p", "443", "port of endpoint")
	remoteCmd.Flags().BoolVarP(&skipVerify, "skipVerify", "s", false, "skip SSL chain verification")
}
