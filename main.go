package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	var email string
	var checkFormat, checkHost, checkMX, checkSPF, checkDMARC bool

	rootCmd := &cobra.Command{
		Use:   "email-verify",
		Short: "A command line utility to verify and validate email addresses developed using Go with Cobra and Viper.",
		Run: func(cmd *cobra.Command, args []string) {
			if checkFormat && !isValidFormat(email) {
				color.Red("❗ Invalid email format ❌")
				return
			}
			if checkHost && !isValidHost(email) {
				color.Red("❗ Invalid host ❌")
				color.Yellow("Try using an email address using a proper host 🔰")
				return
			}

			if checkMX && !isValidMX(email) {
				color.Red("❗ Invalid MX record ❌")
				color.Yellow("Try using an email address using a proper host 🔰")
				return
			}

			if checkDMARC && !isValidDMARC(email) {
				color.Red("❗ Invalid DMARC record ❌")
				color.Yellow("Try using an email address using a proper host 🔰")
				return
			}

			if checkSPF && !isValidSPF(email) {
				color.Red("❗ Invalid SPF record ❌")
				color.Yellow("Try using an email address using a proper host 🔰")
				return
			}

			color.Green("Valid email address ✅")
		},
	}

	rootCmd.Flags().StringVarP(&email, "email", "e", "", "Email address to verify")
	rootCmd.MarkFlagRequired("email")
	viper.BindPFlag("email", rootCmd.Flags().Lookup("email"))

	rootCmd.Flags().BoolVar(&checkFormat, "format", true, "Check email format")
	rootCmd.Flags().BoolVar(&checkHost, "host", true, "Check host validity")
	rootCmd.Flags().BoolVar(&checkMX, "mx", true, "Check MX record")
	rootCmd.Flags().BoolVar(&checkDMARC, "dmarc", true, "Check DMARC record")
	rootCmd.Flags().BoolVar(&checkSPF, "spf", true, "Check SPF record")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func isValidFormat(email string) bool {
	at := strings.LastIndex(email, "@")
	if at <= 0 {
		return false
	}
	if at == len(email)-1 {
		return false
	}
	color.Green("Email provided is in valid format ✅")
	return true
}

func isValidHost(email string) bool {
	_, host := split(email)
	hst, err := net.LookupHost(host)
	if err != nil || len(hst) == 0 {
		return false
	}
	color.Green("Email host is valid ✅")
	return true
}

func isValidMX(email string) bool {
	_, host := split(email)
	mxs, err := net.LookupMX(host)
	if err != nil || len(mxs) == 0 {
		return false
	}
	color.Green("Email MX Records are valid ✅")
	return true
}

func isValidDMARC(email string) bool {
	_, host := split(email)
	addrs, err := net.LookupTXT(fmt.Sprintf("%s.%s", "_dmarc", host))
	if err != nil || len(addrs) == 0 {
		return false
	}
	color.Green("Email provided has valid DMARC Records found ✅")
	return true
}

func isValidSPF(email string) bool {
	_, host := split(email)
	addrs, err := net.LookupTXT(fmt.Sprintf("%s.%s", "_spf", host))
	if err != nil || len(addrs) == 0 {
		return false
	}
	color.Green("Email provided has valid SPFv ✅")
	return true
}

func split(email string) (string, string) {
	at := strings.LastIndex(email, "@")
	return email[:at], email[at+1:]
}
