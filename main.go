package main

import (
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

func isValidEmail(email string) bool {
	// Regular expression pattern for a valid email address
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return re.MatchString(email)
}

var email string

var cmdVerifyEmail = &cobra.Command{
	Use:   "verify-email",
	Short: "Verify the validity of an email address",
	Long:  "A simple utility to check if a given email address is valid",
	Run: func(cmd *cobra.Command, args []string) {
		if isValidEmail(email) {
			fmt.Println("\u2705 " + email + " is valid")
		} else {
			fmt.Println("\u274C " + email + " is not valid")
		}
	},
}

func main() {

	cmdVerifyEmail.Flags().StringVarP(&email, "email", "e", "", "Email address to verify")
	cmdVerifyEmail.MarkFlagRequired("email")

	if err := cmdVerifyEmail.Execute(); err != nil {
		fmt.Println(err)
	}

	scanner := email

	fmt.Printf("Domain, hasMX, hasSPF, SPF Record, hasDMARC, DMARC Record \n")

	checkdomain(scanner)

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error!! could not resolve from input: %v \n", err)
	}
}

func checkdomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)

	if err != nil {
		log.Printf("Error encountered %v \n", err)
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	txtRecords, err := net.LookupTXT(domain)

	if err != nil {
		log.Printf("Error Encountered:%v \n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error %v \n", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v \n", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
