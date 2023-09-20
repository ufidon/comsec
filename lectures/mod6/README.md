# Malicious Software
ch6


🔭 Explore
---
- [WannaCry ransomware attack](https://en.wikipedia.org/wiki/WannaCry_ransomware_attack)
- [What Is Wannacry/Wanacrypt0r?](https://www.cisa.gov/sites/default/files/FactSheets/NCCIC%20ICS_FactSheet_WannaCry_Ransomware_S508C.pdf)


What is [malware](https://en.wikipedia.org/wiki/Malware)?
---
- [NIST 800-83](https://csrc.nist.gov/pubs/sp/800/83/r1/final) defines malware as:
- a program that is inserted into a system, 
  - usually covertly, with the intent of compromising 
    - the confidentiality, integrity, or availability 
      - of the victim’s data, applications, or operating system 
    - or otherwise annoying or disrupting the victim


Classification of Malware
---
- Classified into two broad categories based on
  - how it spreads or propagates to the targets
  - the actions or payloads it performs on the targets
- classified by existence form
  - Those that need a host program (parasitic code such as viruses)
  - Those that are independent, self-contained programs (worms, trojans, and bots)
- classified by replication capability
  - Malware that does not replicate (trojans and spam e-mail)
  - Malware that does replicate (viruses and worms)


Types of Malware
---
- Propagation mechanisms:
  - Infection of existing content by viruses that is subsequently spread to other systems
  - Exploit of software vulnerabilities by worms or drive-by-downloads to allow the malware to replicate
  - Social engineering attacks that convince users to bypass security mechanisms to install Trojans or to respond to phishing attacks
- Payloads:
  - Corruption of system or data files
  - Theft of service/make the system a zombie agent of attack as part of a botnet
  - Theft of information from the system
  - Hiding its presence on the system


Attack/exploit Kits/crimeware
---
- The development and deployment of malware require considerable technical skill 
- Malware-creation toolkits help novices develop and deploy malware
  - includes a variety of propagation mechanisms and payload modules 
  - generate variants of malware


🔭 Explore
---
- [exploit kit](https://www.trendmicro.com/vinfo/us/security/definition/exploit-kit)


Attack Sources
Advanced Persistent Threat

Propagation—Infected Content—Viruses
The Nature of Viruses
Macro and Scripting Viruses
Viruses Classification
Propagation—Vulnerability Exploit—Worms
Target Discovery

Worm Propagation Model
The Morris Worm
A Brief History of Worm Attacks
State of Worm Technology
Mobile Code
Mobile Phone Worms

Client-Side Vulnerabilities and Drive-by-Downloads
Clickjacking
Propagation—Social Engineering—Spam E-Mail, Trojans
Spam (Unsolicited Bulk) E-Mail

Trojan Horses
Mobile Phone Trojans
Payload—System Corruption
Data Destruction
Real-World Damage

Logic Bomb
Payload—Attack Agent—Zombie, Bots
Uses of Bots
Remote Control Facility
Payload—Information Theft—Keyloggers, Phishing, Spyware
Credential Theft, Keyloggers, and Spyware
Phishing and Identity Theft
Reconnaissance, Espionage, and Data Exfiltration
Payload—Stealthing—Backdoors, Rootkits

Backdoor
Rootkit
Kernel Mode Rootkits
Virtual Machine and Other External Rootkits
Countermeasures
Malware Countermeasure Approaches
Host-Based Scanners and Signature-Based Anti-Virus
Perimeter Scanning Approaches
Distributed Intelligence Gathering Approaches


Terminology for malware
---
| Name | Description |
| --- | --- |
| Advanced Persistent Threat (APT) | Cybercrime directed at business and political targets, using a wide variety of intrusion technologies and malware, applied persistently and effectively to specific targets over an extended period, often attributed to state-sponsored organizations. |
| Adware | Advertising that is integrated into software. It can result in pop-up ads or redirection of a browser to a commercial site. |
| Attack kit | Set of tools for generating new malware automatically using a variety of supplied propagation and payload mechanisms. |
| Auto-rooter | Malicious hacker tools used to break into new machines remotely. |
| Backdoor (trapdoor) | Any mechanism that bypasses a normal security check; it may allow unauthorized access to functionality in a program, or onto a compromised system. |
| Downloaders | Code that installs other items on a machine that is under attack. It is normally included in the malware code first inserted on to a compromised system to then import a larger malware package. |
| Drive-by-download | An attack using code on a compromised website that exploits a browser Vulnerability to attack a client system when the site is viewed. |
| Exploits | Code specific to a single vulnerability or set of vulnerabilities. |
| Flooders (DoS client) | Used to generate a large volume of data to attack networked computer systems, by carrying out some form of denial-of-service (DoS) attack. |
| Keyloggers | Captures keystrokes on a compromised system. | 
| Logic bomb | Code inserted into malware by an intruder. A logic bomb lies dormant until a Predefined condition is met; the code then triggers some payload. |
| Macro virus | A type of virus that uses macro or scripting code, typically embedded in a Document or document template, and triggered when the document is viewed or edited, to run and replicate itself into other such documents. |
| Mobile code | Software (e.g., script and macro) that can be shipped unchanged to a heterogeneous collection of platforms and execute with identical semantics. |
| Rootkit | Set of hacker tools used after attacker has broken into a computer system and gained root-level access. |
| Spammer programs | Used to send large volumes of unwanted e-mail. |
| Spyware | Software that collects information from a computer and transmits it to another system by monitoring keystrokes, screen data, and/or network traffic; or by scanning files on the system for sensitive information. |
| Trojan horse | A computer program that appears to have a useful function, but also has a hidden and potentially malicious function that evades security mechanisms, sometimes by exploiting legitimate authorizations of a system entity that invokes it. |
| Virus | Malware that, when executed, tries to replicate itself into other executable machine or script code; when it succeeds, the code is said to be infected. When the infected code is executed, the virus also executes. |
| Worm | A computer program that can run independently and can propagate a complete working version of itself onto other hosts on a network, by exploiting software vulnerabilities in the target system, or using captured authorization credentials. |
| Zombie, bot | Program installed on an infected machine that is activated to launch attacks on other machines. |

# References
- [Live cyber threat map](https://threatmap.checkpoint.com/)