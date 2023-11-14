package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")

	for scanner.Scan() {
		checkDomain(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error: could not read from input: %v\n", err)

	}

}

func checkDomain(domain string) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	handleErr(err)

	if len(mxRecords) > 0 {
		hasMX = true

	}

	txtRecords, err := net.LookupTXT(domain)
	handleErr(err)

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	handleErr(err)

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v, %v, %v, %v, %v, %v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

}

func handleErr(err error) {
	if err != nil {
		log.Printf("Error: %v\n", err)
	}
}

/*
NOTES---

** SENDER POLICY FRAMEWORK (SPF)  **

Sender Policty Framework is an authentication protocol that lists
IP addresses in a DNS TXT

A typical SPF record looks like this:

“v=spf1 ip4:64.34.183.84 ip4:64.34.183.88 include:mmsend.com -all”

** DOMAIN-BASED MESSAGE AUTHENTICATION, REPORTING & CONFORMANCE (DMARC) **

DMARC allows the domain owner to specify how unauthenticated messages should be treated by MBPs. This is accomplished by what is known as a “policy” that is set in the DMARC DNS record. The policy can be set to one of three options: NONE, QUARANTINE, and REJECT.

Policy = (p=none): no action and message delivered as normal
Policy = (p=quarantine): places the message to spam/junk/quarantine folder
Policy = (p=reject): the message rejected/bounced


*/
