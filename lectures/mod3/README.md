# User Authentication
ch3


## Digital User Authentication Principles

User authentication
---
- How to prove you are you?
  -  used for access control and user accountability
- The definition in [NIST SP 800-63 Digital Identity Guidelines](https://pages.nist.gov/800-63-4/)
  - The process of determining the validity of one or more authenticators used to claim a digital identity 
  - Authentication establishes that a subject attempting to access a digital service is in control of the technologies used to authenticate
- Includes two functions
  - the user identifies herself to the system by *presenting* one or more authenticators
  - the system verifies these authenticators


[Identification and Authentication Security Requirements](https://nvlpubs.nist.gov/nistpubs/SpecialPublications/NIST.SP.800-171r2.pdf)
---
- Basic Requirements:
  - Identify information system users, processes, or devices
    - Authenticate (or verify) the identities of those users, processes, or devices
- Derived Requirements:
    - Use multifactor authentication for local and network access to privileged accounts and non-privileged accounts
  - Employ replay-resistant authentication mechanisms for network access to privileged and non-privileged accounts
    - Prevent reuse of identifiers for a defined period
    - Disable identifiers after a defined period of inactivity
  - Enforce password policies
    - ensure minimum password complexity and change of characters for new passwords
    - Prohibit password reuse for a specified number of generations
    - Allow temporary password use for system logons with an immediate change to a permanent password
    - Store and transmit only cryptographically-protected passwords
  - Obscure feedback of authentication information


A Model for Digital User Authentication
---
- *Non-federated* [digital identity model](https://pages.nist.gov/800-63-4/sp800-63.html)
- ![Non-federated Digital Identity Model Example](https://pages.nist.gov/800-63-4/sp800-63/media/Non-Federated.png)
- Subject is represented by one of the three roles:
  - Applicant
  - Subscriber
  - Claimant
- Credential Service Provider (CSP)
- Relying Party (RP): 
  - An entity that relies upon the information in the subscriber account, or 
  - An identity provider (IdP) assertion when using federation, typically to process a transaction or grant access to information or a system
- Identity Provider (IdP): an entity in a *federated model* that performs both the CSP and Verifier functions


🖊️ Practice
---
- [Explore the steps in the *Non-federated* digital identity model](https://pages.nist.gov/800-63-4/sp800-63.html)



Authenticator - Means of Authentication
---
- four types of authenticators
  - Something you know (e.g., a password)
  - Something you have (e.g., an ID badge or a cryptographic key)
  - Something you are (e.g., a fingerprint or other biometric characteristic data)
  - How you behave (e.g., voice pattern or handwriting characteristics)
- Multifactor authentication
  - uses at least two types of authenticators


Risk Assessment for User Authentication
---
- Assurance Level is the degree of conifdence on
  - the *credential* presented by the user
  - the *user* presents the credential
- Potential impact
  - three qualitative levels: low, moderate and high
- Areas of risk


Identity Assurance Levels (IAL)
---
- A qualitative demarcation of 4 levels
- Each successive IAL builds on the requirements of lower IALs in order to achieve greater assurance:
- No identity proofing (IAL0)
  - no requirement to link the applicant to a specific, real-life identity
- IAL1
  - The identity proofing process supports the real-world existence of the claimed identity
- IAL2
  - adds additional rigor to the identity proofing process by requiring the collection of stronger types of evidence and a more rigorous process for validating the evidence and verifying the identity
- IAL3
  - adds the requirement for a trained CSP representative to interact directly with the applicant during the entire identity proofing session, either in person or via a supervised remote identity proofing session


🖊️ Practice
---
- [Explore the Table of Maximum Potential Impacts for Each Assurance Level](https://pages.nist.gov/800-63-3/sp800-63-3.html#63Sec6-Table6-1)


## Password-Based Authentication

[Password-Based Authentication](https://en.wikipedia.org/wiki/Password)
---
- Consists of a pair of user name, id or email address and password
- Passwords are saved as hash codes
- The user ID
  - determines that the user is authorized to access the system
  - determines the user’s privileges
  - used in discretionary access control


The Vulnerability of Passwords
---
- dictionary attack
  - Popular password attack
  - Password guessing against single user
- Workstation and session hijacking
- Exploiting user mistakes
  - Exploiting multiple password use
- Electronic monitoring with key logger
- shoulder surfing
- phishing email


The Use of [Hashed Passwords](https://hashcat.net/wiki/doku.php?id=example_hashes)
---
- Linux [/etc/passwd](https://manpages.ubuntu.com/manpages/focal/en/man5/passwd.5.html) file saves user accounts one per line
- each line has 7 fields delimited by colons
  - login name
  - optional encrypted password
    - a lower-case “x” indicates the encrypted password is  stored in the shadow file
  - numerical user ID
  - numerical group ID
  - user name or comment field
  - user home directory
  - optional user command interpreter


Old [Unix password scheme](https://www.oreilly.com/library/view/practical-unix-and/0596003234/ch04s03.html)
---
- Up to eight printable characters in length
- 12-bit salt used to turn DES encryption into a one-way hash function
- Zero value repeatedly encrypted 25 times with user password
- Output translated to 11 character sequence
- Implemented with the POSIX C library function [crypt](https://en.wikipedia.org/wiki/Crypt_(C))
- Now regarded as insecure
  - Still often required for compatibility with existing account management software


Improved Unix password scheme
---
- Recommended hash function is based on MD5
  - Salt of up to 48-bits
  - Password length is unlimited
  - Produces 128-bit hash
  - Uses an inner loop with 1000 iterations to achieve slowdown
- OpenBSD uses Blowfish block cipher based hash algorithm called Bcrypt
  - Most secure version of Unix hash/salt scheme
  - Uses 128-bit salt to create 192-bit hash value


💡 Demo
---
- Explore /etc/passwd and /etc/shadow


🖊️ Practice
---
Explore the password storage schemes on
- [Unix](https://www.oreilly.com/library/view/practical-unix-and/0596003234/ch04s03.html)
- [Linux](https://ma.ttias.be/how-to-generate-a-passwd-password-hash-via-the-command-line-on-linux/)
- [MacOS](https://embracethered.com/blog/posts/2022/grabbing-and-cracking-macos-hashes/)
- [Windows](https://learn.microsoft.com/en-us/windows-server/security/kerberos/passwords-technical-overview)



Password Cracking
---
- Dictionary attacks
  - Develop a large dictionary of possible passwords and try each against the password file
  - Each password must be hashed using each salt value and then compared to stored hash values
- Rainbow table attacks
  - Pre-compute tables of hash values for all salts
  - A mammoth table of hash values 
  - Can be countered by using a sufficiently large salt value and a sufficiently large hash length
- Password crackers exploit the fact that people choose easily guessable passwords
  - Shorter password lengths are also easier to crack
- [John the Ripper](https://www.openwall.com/john/)
  - an Open Source password security auditing and password recovery tool 
  - Uses a combination of brute-force and dictionary techniques


Password strengthening vs. cracking
---
- strengthen by complex password policy
  - Force users to pick stronger passwords
- However password-cracking techniques have also improved
  - processing capacity
  - sophisticated algorithms for generating potential passwords
  - Studying examples and structures of actual passwords
- The percentage of passwords guessed increases with the number of guesses


Password File Access Control
---
- Can block offline guessing attacks by denying access to encrypted passwords
  - Make available only to privileged users
    - Shadow password file
- Vulnerabilities
  - Weakness in the OS that allows access to the file
  - Accident with permissions making it readable
  - Users with same password on other systems
  - Access from backup media
  - Sniff passwords in network traffic


Password Selection Strategies
---
- User education
  - the importance of using hard to guess passwords 
  - guidelines for selecting strong and memorable passwords
- Computer generated passwords
  - secure but hard to remember
- Reactive password checking
  - System periodically runs its own password cracker to find guessable passwords
- Complex password policy
  - accept complex passwords
  - eliminate guessable passwords


Proactive Password Checking
---
- Rule enforcement
  - Specific rules that passwords must satisfy
- Password checker
  - Compile a large dictionary of bad passwords not to use
- [Bloom filter](https://en.wikipedia.org/wiki/Bloom_filter)
  - Used to build a table based on hash values
  - Check desired passwords against this table



## Token-Based Authentication

Token-Based Authentication
---
Tokens are objects that a user possesses for user authentication
- Memory Cards
- Smart Cards
- Electronic Identify Cards


Types of Cards Used as Tokens
---
| Card Type |  Defining Feature | Example |
| --- | --- | --- |
| Embossed  |  Raised characters only, on front | Old credit card |
| Magnetic stripe | Magnetic bar on back, characters on front | Bank card |
| Memory  |  Electronic memory inside | Prepaid phone card |
| Smart<br>∘ Contact<br>∘ Contactless | Electronic memory and processor inside<br>Electrical contacts exposed on surface<br>Radio antenna embedded inside | Biometric ID card |


Memory Cards
---
- Can store but do not process data
- The most common is the magnetic stripe card
- Can include an internal electronic memory
- Can be used alone for physical access
  - Hotel room
  - ATM
- Provides significantly greater security when combined with a password or PIN
- Drawbacks of memory cards include:
  - Requires a special reader
  - Loss of token
  - User dissatisfaction


Smart Tokens
---
- Include an embedded microcontroller
  - looks like bank cards
  - Can look like calculators, keys, small portable objects
- User interface
- Electronic interface
  - communicates with a compatible [reader/writer](https://en.wikipedia.org/wiki/Card_reader)
  - contact and [contactless interfaces](https://en.wikipedia.org/wiki/Contactless_smart_card)
- Three categories of authentication protocol:
  - Static
  - Dynamic password generator
  - Challenge-response


❓ Question
---
- How about a smart phone with Google pay or Apple pay?


[Smart Cards](https://en.wikipedia.org/wiki/Smart_card)
---
- Most important category of smart token
  - Has the appearance of a credit card
  - Has an electronic interface
  - May use any of the smart token protocols
- Contain an entire microcontroller
- Typically include three types of memory: ROM, EEPROM, and RAM



🖊️ Practice
---
- [Explore the eID Card for citizens of the EU and the EEA](https://www.personalausweisportal.de/Webs/PA/EN/citizens/id-card-for-eu-and-eea/eID-card-for-eu-and-eea-node.html)
- [Explore smart card application protocol data unit](https://en.wikipedia.org/wiki/Smart_card_application_protocol_data_unit)



[Electronic Identity Cards (eID)](https://en.wikipedia.org/wiki/Electronic_identification)
---
- a smart card verified by government as valid and authentic
- contains
  - Personal data
  - Document number and card access number (CAN)
  - Machine readable zone (MRZ)
- three functions
  - ePass, electronic passport
  - eID, electronic identity
  - eSign, generating digital signatures 


Password Authenticated Connection Establishment (PACE)
---
- Prevent unauthorized access to the eID card
- For online applications, 
  - access is established by entering the secret 6-digit PIN
- For offline applications, 
  - either the MRZ printed on the back of the card or 
  - the six-digit card access number (CAN) printed on the front is used



## Biometric Authentication

[Biometric Authentication](https://en.wikipedia.org/wiki/Biometrics)
---
- authenticate a user based on unique physical characteristics
  - by pattern recognition
- complex and expensive compared to passwords and tokens
- typical physical characteristics:
  - Facial characteristics, fingerprints, hand geometry
  - Retinal pattern, iris
  - Signature, voice


🖊️ Practice
---
- [Explore biometric Devices: Cost, Types and comparison](https://www.bayometric.com/biometric-devices-cost/)
- [How does face ID work](https://en.wikipedia.org/wiki/Face_ID)
- [How does touch ID work](https://en.wikipedia.org/wiki/Touch_ID)


## Remote User Authentication

Remote User Authentication
---
- complex to authenticate over network
- security threats
  - host attack, client attack
  - eavesdropping 
  - replay attack
  - trojan horse
  - Denial-of-Service
- counter threats with a challenge-response protocol


🖊️ Practice
---
- two challenge-response protocols
  - [Explore the simple example of mutual authentication sequence](https://en.wikipedia.org/wiki/Challenge%E2%80%93response_authentication)
  - [Explore TLS handshake](https://en.wikipedia.org/wiki/Transport_Layer_Security)



Security Issues for User Authentication
---
- [11 Common Authentication Vulnerabilities You Need to Know](https://www.strongdm.com/blog/authentication-vulnerabilities)



# References
- [NIST Special Publication 800-63 Revision 3: Digital Identity Guidelines](https://pages.nist.gov/800-63-3/sp800-63-3.html)
- [FIDO: The Future of Passwords is Passwordless](https://fidoalliance.org/)