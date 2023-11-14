# Domain-Verifier-Tool

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
