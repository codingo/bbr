# bbr
A command line application to taketemplate information and process it over the command line. Useful for piping reporting from one application to another (such as an automatic submission tool).

# Arguments
| Argument | Description                      |
|----------|----------------------------------|
| -h       | Display help message and exit   |
| -r       | Path to template file to use    |
| -t       | Variable to replace \_target\_ with and to use for `dig` and `whois` commands. |
| -u       | Username to replace \_user\_ with |
| -o       | Output file name. (optional)       |
| -p | Variable to replace \_program\_ (optional) |
| -re | Variable to replace \_researcher\_ (optional) |
| -u | Variable to replace \_username\_ (optional) |

BBR will then process the text file, and make the following replacements (not all fields may be present, some will be present more than once):

| Argument      | Description                                               |
|---------------|-----------------------------------------------------------|
| \_target\_      | Replace with the value of the -t argument                |
| \_username\_    | Replace with the value of the -u argument                |
| \_program\_      | Replace with the value of the -p argument                |
| \_researcher\_ | Replace with the value of the -re argument |
| \_username\_ | Replace with the value of the -u argument |
| \_sha\_         | Replace with the SHA256 encoded value of the -u argument |
| \_nameservers\_ | Replace with the output of "dig NS @8.8.8.8 _target_"     |
| \_dig\_         | Replace with the value of "dig @8.8.8.8 _target_"         |
| \_whois\_       | Replace with the whois output of the target parameter    |
| \_wayback\_ | Replace with an automatic wayback link of the -t argument |
| \_sha\_ | Replace with the SHA256 value of the username parameter |
| \_dig-txt\_ | Replace with the value of DNS TXT records |
| \_curl\_ | Replace with the request response of the -t argument |
| \_joke\_ | Replace with a joke |
| \_punchline\_ | Replace with the punchline for said joke |

# Funcionality
BBR takes a provided template file and makes replacements throughout that file with provided arguments. For example, the following template file (stored in this repository as `template.txt`:

```
 # Summary
The domain _target_ was found to have a CNAME that was pointing to an unregistered domain.

It was possible to register this domain, and to host content on the _target_ website. Given this domain is attributed to _program_(see: attribution) I hosted only a SHA256 string of my researcher account, _researcher).

This can be verified by using the following in the terminal:

\```
echo "_username_" | sha256sum
\```
Which should present the resulting string:
\```
_sha_
\```
Which matches what I placed on _target_ for verification.

This has also been stored on the Wayback engine, in case this is resolved before this submission is able to be triaged: _wayback_

# Attribution
A whois of the domain _target_ shows a direct match to other domains relating to _program_, showing this as beloning to _program_:

\```
_whois_
\```

# Recommendation
Remove the CNAME associated with _target_, or decomission the domain entirely with a redirection to other domains of _program_. If you would like the domain I've claimed to be transferred to you, please don't hestitate to request it within this submission.

# Joke
Triage is a tough gig, here's a joke to lighten the load!

_joke_

... _punchline_
```

When used with the following:

```
➜  ./bbr -t example.com -p Example -u codingo -r ./template.txt | tee  
```
Outputs the following report:
```
 # Summary
The domain example.com was found to have a CNAME that was pointing to an unregistered domain.

It was possible to register this domain, and to host content on the example.com website. Given this domain is attributed to Example(see: attribution) I hosted only a SHA256 string of my researcher account, _researcher).

This can be verified by using the following in the terminal:

\```
echo "codingo" | sha256sum
\```
Which should present the resulting string:
\```
10c989bbd4963c465e0941acd70833d5579ca846f5a68eadc8bcf63801b3993b
\```
Which matches what I placed on example.com for verification.

This has also been stored on the Wayback engine, in case this is resolved before this submission is able to be triaged: example.com

# Attribution
A whois of the domain example.com shows a direct match to other domains relating to Example, showing this as beloning to Example:

\```
   Domain Name: EXAMPLE.COM
   Registry Domain ID: 2336799_DOMAIN_COM-VRSN
   Registrar WHOIS Server: whois.iana.org
   Registrar URL: http://res-dom.iana.org
   Updated Date: 2020-08-14T07:02:37Z
   Creation Date: 1995-08-14T04:00:00Z
   Registry Expiry Date: 2021-08-13T04:00:00Z
   Registrar: RESERVED-Internet Assigned Numbers Authority
   Registrar IANA ID: 376
   Registrar Abuse Contact Email:
   Registrar Abuse Contact Phone:
   Domain Status: clientDeleteProhibited https://icann.org/epp#clientDeleteProhibited
   Domain Status: clientTransferProhibited https://icann.org/epp#clientTransferProhibited
   Domain Status: clientUpdateProhibited https://icann.org/epp#clientUpdateProhibited
   Name Server: A.IANA-SERVERS.NET
   Name Server: B.IANA-SERVERS.NET
   DNSSEC: signedDelegation
   DNSSEC DS Data: 31589 8 1 3490A6806D47F17A34C29E2CE80E8A999FFBE4BE
   DNSSEC DS Data: 31589 8 2 CDE0D742D6998AA554A92D890F8184C698CFAC8A26FA59875A990C03E576343C
   DNSSEC DS Data: 43547 8 1 B6225AB2CC613E0DCA7962BDC2342EA4F1B56083
   DNSSEC DS Data: 43547 8 2 615A64233543F66F44D68933625B17497C89A70E858ED76A2145997EDF96A918
   DNSSEC DS Data: 31406 8 1 189968811E6EBA862DD6C209F75623D8D9ED9142
   DNSSEC DS Data: 31406 8 2 F78CF3344F72137235098ECBBD08947C2C9001C7F6A085A17F518B5D8F6B916D
   URL of the ICANN Whois Inaccuracy Complaint Form: https://www.icann.org/wicf/
>>> Last update of whois database: 2020-08-22T03:11:57Z <<<

For more information on Whois status codes, please visit https://icann.org/epp

NOTICE: The expiration date displayed in this record is the date the
registrar's sponsorship of the domain name registration in the registry is
currently set to expire. This date does not necessarily reflect the expiration
date of the domain name registrant's agreement with the sponsoring
registrar.  Users may consult the sponsoring registrar's Whois database to
view the registrar's reported date of expiration for this registration.

TERMS OF USE: You are not authorized to access or query our Whois
database through the use of electronic processes that are high-volume and
automated except as reasonably necessary to register domain names or
modify existing registrations; the Data in VeriSign Global Registry
Services' ("VeriSign") Whois database is provided by VeriSign for
information purposes only, and to assist persons in obtaining information
about or related to a domain name registration record. VeriSign does not
guarantee its accuracy. By submitting a Whois query, you agree to abide
by the following terms of use: You agree that you may use this Data only
for lawful purposes and that under no circumstances will you use this Data
to: (1) allow, enable, or otherwise support the transmission of mass
unsolicited, commercial advertising or solicitations via e-mail, telephone,
or facsimile; or (2) enable high volume, automated, electronic processes
that apply to VeriSign (or its computer systems). The compilation,
repackaging, dissemination or other use of this Data is expressly
prohibited without the prior written consent of VeriSign. You agree not to
use electronic processes that are automated and high-volume to access or
query the Whois database except as reasonably necessary to register
domain names or modify existing registrations. VeriSign reserves the right
to restrict your access to the Whois database in its sole discretion to ensure
operational stability.  VeriSign may restrict or terminate your access to the
Whois database for failure to abide by these terms of use. VeriSign
reserves the right to modify these terms at any time.

The Registry database contains ONLY .COM, .NET, .EDU domains and
Registrars.
% IANA WHOIS server
% for more information on IANA, visit http://www.iana.org
% This query returned 1 object

domain:       EXAMPLE.COM

organisation: Internet Assigned Numbers Authority

created:      1992-01-01
source:       IANA


\```

# Recommendation
Remove the CNAME associated with example.com, or decomission the domain entirely with a redirection to other domains of Example. If you would like the domain I've claimed to be transferred to you, please don't hestitate to request it within this submission.

# Joke
Triage is a tough gig, here's a joke to lighten the load!

What was the pumpkin’s favorite sport?

... Squash.
```

This can then be submitted to your platform of choice, and is a repeatable template as you find similar vulnerablities of the same type.
