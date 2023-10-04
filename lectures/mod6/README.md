# Malicious Software
ch6


🔭 Explore
---
- [WannaCry ransomware attack](https://en.wikipedia.org/wiki/WannaCry_ransomware_attack)
  - [What Is Wannacry/Wanacrypt0r?](https://www.cisa.gov/sites/default/files/FactSheets/NCCIC%20ICS_FactSheet_WannaCry_Ransomware_S508C.pdf)


## Types of Malware

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
- Propagation mechanisms (replication methods):
  - Infect existing files by viruses 
  - Exploit software vulnerabilities by worms or drive-by-downloads 
  - Social-engineer users to install Trojans or to respond to phishing attacks
- Payloads:
  - Corrupt target system/data files
  - Make the target a zombie as part of a botnet
  - Steal information/service from the target
  - Hide its presence on the victim target


Attack Kits/exploit Kits/crimeware
---
- The development and deployment of malware require considerable technical skills 
- Malware-creation toolkits ease malware development and deployment
  - include a variety of propagation mechanisms and payload modules 
  - generate variants of malware


🔭 Explore
---
- [Old exploit kits](https://www.trendmicro.com/vinfo/us/security/definition/exploit-kit)


💡 Demo
---
- [State-of-the-art expoit kit: Metasploit Framework](https://docs.rapid7.com/metasploit/msf-overview/)


Attack Sources
---
- individuals
- hacker groups
- criminals, organized crime
- organizations that sell their services to companies and nations
  - fostered a large underground economy of attack kits, botnets, and stolen information
- politically motivated attackers
- national government agencies
  - [NSA(National Security Agency)](https://www.nsa.gov/)
    - nickname: No Such Agency


🔭 Explore
---
- [The most powerful hacker groups](https://nordvpn.com/blog/hacker-groups/)

---

Advanced Persistent Threats (APTs)
---
- Well-resourced, persistent malwares with a wide variety of intrusion technologies 
- Typically attributed to state-sponsored organizations and criminal groups
- careful target selection usually business or political
- stealthy intrusion efforts over long periods


🔭 Explore
---
- [List of cyberattacks](https://en.wikipedia.org/wiki/List_of_cyberattacks)


APT Characteristics
---
- Advanced
  - launched by organized, capable, and well-funded attackers
  - use recent intrusion technologies and vulnerability exploits
  - carefully select components which are suitable to the target
- Persistent
  - attack progressively until the target is compromised
  - attack persistently to maximize the loot
- Threats
  - intend to compromise the targets completely
  - the targets greatly raise the threat level to defend

---

## Malware Propagation
- Infect content by viruses
- Exploit vulnerabilities by worms
- Social engineer with spam emails and trojans


[Viruses](https://en.wikipedia.org/wiki/Computer_virus)
---
- Piece of software that infects programs
  - Injects them with a copy of the virus
  - Replicates and goes on to infect other programs
  - Easily spreads through network
- The virus can do anything that the infected program is permitted to do
  - Executes secretly when the host program is running
- Specific to operating system and hardware
  - Takes advantage of their capabilities and weaknesses


Viruses components
---
- Infection mechanism / infection vector
  - How the virus propagates
- Trigger / logic bomb
  - When and how the virus activates or delivers its payload 
- Payload
  - What the virus does (besides spreading)
    - damage or pranks


Virus phases
---
| Phase | Activities |
| --- | --- |
| Dormant | ▶️Virus is idle<br>▶️Will eventually be activated by some event<br>▶️Not all viruses have this stage |
| Trigger | ▶️Virus is activated to deliver its payload<br>▶️Can be caused by a variety of system events |
| Propagation | ▶️Virus inject a copy of itself into other programs or certain disk areas<br>▶️May not be identical to the propagating version<br>▶️Each infected program will enter a propagation phase |
| Execution | ▶️Playload is delivered<br>▶️May be pranking or damaging |


Macro and Scripting Viruses
---
- attaches itself to documents 
- infects scripting code used to support active content in a variety of user document types
  - e.g. word, excel, pdf
- uses the macro programming capabilities of the document’s application to execute and propagate
- characteristics
  - platform independent by using scripting languages
    - easier to write or modify
  - infect documents instead of programs
  - easily to spread
    - file system access controls usually check programs instead of documents


🔭 Explore
---
- [The Melissa Virus: An $80 Million Cyber Crime in 1999 Foreshadowed Modern Threats](https://www.fbi.gov/news/stories/melissa-virus-20th-anniversary-032519)
  - [source code](./resources/mellisa.vbs)
- [What is a Macro virus and how to remove it?](https://cybernews.com/malware/macro-virus-definition/)


Viruses Classification
---
- by target infected
  - Boot sector virus
  - File virus 
  - Macro virus
  - Multipartite virus
- by concealment strategy
  - Encrypted virus
    - exploit and payload are encrypted
  - Stealth virus
    - hides from detection by anti-virus software
  - Polymorphic virus
    - mutates with every infection
  - Metamorphic virus
    - mutates and rewrites itself completely at each iteration 
    - may change behavior and appearance


[How worms reproduce and propagate?](https://en.wikipedia.org/wiki/Computer_worm)
---
- a standalone malware able to replicate itself for propagation
- Exploits software vulnerabilities in client or server programs
- replicates and propagates through 
  - Electronic mail or instant messenger facility
    - e-mails a copy of itself to other systems
    - Sends itself as an attachment via an instant message service
  - File sharing
    - Creates a copy of itself or infects a file as a virus on removable media
  - Remote execution capability
    - executes a copy of itself on another system
  - Remote file transfer capability
    - uses a remote file transfer service to copy itself from one system to the other
  - Remote login capability
    - logs onto a remote system then copies itself from one system to the other
- exhibit similar self-replication and propagation behavior to biological viruses as[epidemic model](https://en.wikipedia.org/wiki/Compartmental_models_in_epidemiology)


How worms find their targets?
---
- scan directly outside the target network
  - thoroughly
  - randomly
  - selectively
- scan indirectly through botnets 
  - outside the target network 
  - inside the target network


🔭 Explore
---
- [Mydoom – $38 billion: the worst computer worm outbreak in history](https://en.wikipedia.org/wiki/Mydoom)
- [Top 10 Worms of all time](https://www.secpoint.com/top-10-worms.html)
- [Timeline of computer viruses and worms](https://en.wikipedia.org/wiki/Timeline_of_computer_viruses_and_worms)



[Mobile Code](https://csrc.nist.gov/pubs/sp/800/28/ver2/final)
---
- not code for mobiles or smartphones
- can be shipped unchanged to a heterogeneous collection of platforms 
  - and executed with identical semantics
  - e.g., script, macro, or other portable instruction such as
    - Java applets, ActiveX, JavaScript, VBScript, etc.
- acts as a mechanism for a virus, worm, or Trojan horse
- common malicious operations
  - Cross-site scripting
  - Interactive and dynamic Web sites
  - E-mail attachments
  - Downloads from untrusted sites or of untrusted software


Mobile Phone Worms
---
- Communicate through Bluetooth, mobile connections or MMS
- Target smartphones
  - disable the phone, 
  - delete data on the phone
  - force the phone to send costly messages


🔭 Explore
---
- [Mobile malware](https://en.wikipedia.org/wiki/Mobile_malware)


Client-Side Vulnerabilities and Drive-by-Downloads
---
- Exploits browser and plugin vulnerabilities 
- downloads and installs malware on the system without the user’s knowledge or consent
  - when the user views a webpage controlled by the attacker
- usually does not actively propagate as a worm does
- Spreads when users visit the malicious Web page


Watering-Hole Attacks
---
- variants of drive-by-downloads used in highly targeted attacks
  - ambush on websites
- compromise the vulnerable websites frequently visited by intended victims
- the intended victims are attacked once they visit one of the compromised websites
- may take no action for other visitors to the website
  - increases the likelihood of remaining undetected


[Malvertising](https://en.wikipedia.org/wiki/Malvertising)
---
- injects malware-laden advertisements into legitimate advertising networks and webpages without compromising the system
- grows rapidly because they are easy to deploy and hard to track
- targets and displays to intended victims only
  - this greatly reduces their visibility
- The malwares reduce the chance of detection and infect specific systems using dynamic generation


🖊️ Practice
---
- [Could you find the malicious ads on this webpage?](https://www.sohu.com/a/652349120_614277)


[Clickjacking](https://en.wikipedia.org/wiki/Clickjacking)
---
- overlays multiple transparent or opaque layers on a normal button or link to trick a user into clicking 
  - also known as user-interface (UI) redress attack
- keystrokes can also be hijacked by similar masquerading textboxes
  - intercepts user ids, passwords, etc.
- intercepted mouse clicks and keystrokes can be used to triger further malicous activities


[Social Engineering](https://en.wikipedia.org/wiki/Social_engineering_(security))
---
- Tricks victims to assist in the compromise of their own systems
- typical methods: 
  - [*pretexting*](https://en.wikipedia.org/wiki/Pretexting) tricks users with a devised scenario
  - [*water holing*](https://en.wikipedia.org/wiki/Watering_hole_attack) exploits the trust users have in familiar websites
  - *baiting* uses physical media and relies on the curiosity or greed of the victim
- typical software tools:
  - spam emails, mails and SMS (short message service)
  - [trojan-horse programs](https://en.wikipedia.org/wiki/Trojan_horse_(computing))
    - smartphone trojans target smartphones


🔭 Explore
---
- [IT threat evolution in Q2 2023. Mobile statistics](https://securelist.com/it-threat-evolution-q2-2023-mobile-statistics/110427/)

---

## Malware Payloads
- corrupt systems
- convert targets into attack agents like zombie and bots
- steal information by keyloggers, phishing and spywares
- hide existence with backdoors and rootkits


Payload—System Corruption
---
- Data Destruction
  - [Chernobyl virus](https://en.wikipedia.org/wiki/CIH_(computer_virus)) overwrites the first megabytes of hard drive with zeroes
  - [Klez worm](https://en.wikipedia.org/wiki/Klez) stops and deletes some antivirus programs, empty files
  - [Ransomware](https://en.wikipedia.org/wiki/Ransomware) such as WannaCry, [PC Cyborg trojan](https://en.wikipedia.org/wiki/AIDS_(Trojan_horse))
    -  enrypts victims' documents and demands ransom
- Real-World (physical equipments) Damage
  - Chernobyl virus rewrites BIOS code
  - [Stuxnet worm](https://en.wikipedia.org/wiki/Stuxnet) caused substantial damage to the nuclear program of Iran
- [Logic Bomb](https://en.wikipedia.org/wiki/Logic_bomb)
  - malicious code execution will be triggered by certain conditions


Payload—Attack Agent—Zombie, Bots
---
- Takes over an Internet attached computer and uses that computer to launch or manage attacks
  - the victim computer is called a zombie or a bot
  - collection of bots, called Botnet, can launch attack coordinately
- Uses of botnet:
  - Distributed denial-of-service (DDoS) attacks
  - Spamming
  - Sniffing traffic
  - Keylogging
  - Spreading new malware
  - Installing advertisement add-ons and browser helper objects (BHOs)
  - Attacking IRC chat networks
  - Manipulating online polls/games


Remote Control Facility
---
- Botnet is controlled by a command-and-control (C&C) center 
  - typically implemented on an [IRC (Internet Relay Chat) server](https://en.wikipedia.org/wiki/Internet_Relay_Chat)
  - fuzzles law enforcement agencies with DNS names generated dynamically
  - or with fast-flux DNS where the address associated with a given server name is frequently changed
- Bots join a specific channel on this server and treat incoming messages as commands
- More recent botnets use covert communication channels via protocols such as HTTP
- Distributed control mechanisms use peer-to-peer protocols to avoid a single point of failure
  - stroke by international coordination between law enforcement agencies 


Payload—Information Theft—Keyloggers, Spyware
---
- [Keylogger](https://en.wikipedia.org/wiki/Keystroke_logging)
  - Captures keystrokes in realtime
  - Typically uses filter to return sensitive information such as user id and password
- [Spyware](https://en.wikipedia.org/wiki/Spyware)
  - monitors browsing history and content
  - Redirects certain Web page requests to fake sites
  - modifies data flow between the browser and Web sites


[Phishing](https://en.wikipedia.org/wiki/Phishing) and Identity Theft
---
- Exploits social engineering to leverage the user’s trust by masquerading as communication from a trusted source
  - Include a URL in a spam e-mail that links to a fake Web site that mimics the login page of a banking, gaming, or similar site
  - Suggests that urgent action is required by the user to authenticate their account
  - Attacker exploits the account using the captured credentials
- Spear phishing
  - Recipients are carefully researched by the attacker
  - E-mail is crafted to specifically suit its recipient, often quoting a range of information to convince them of its authenticity



Payload—Stealthing—[Backdoors](https://en.wikipedia.org/wiki/Backdoor_(computing))
---
- Also known as a trapdoor
- Secret entry point into a program allowing the attacker to gain access and bypass the security access procedures
- Maintenance hook is a backdoor used by Programmers to debug and test programs
- Difficult to implement operating system controls for backdoors in applications


Payload—Stealthing—[Rootkits](https://en.wikipedia.org/wiki/Rootkit)
---
- Set of hidden programs installed on the victim machine to maintain covert access 
- Hide by subverting the monitoring and protection mechanisms 
- Give administrator (or root) privileges to attacker
  - allow to do anything on the victim machine


Rootkit Classification Characteristics
---
- Persistent: 
  - Activates each time the system boots
- Memory based:
  - resides in memory and hart to detect
  - cannot survive a reboot
- User mode:
  - Intercepts calls to APIs (application program interfaces) 
- Kernel mode:
  -  intercepts calls to native APIs in kernel mode
- Virtual machine based:
  - sits between hardware and OS as  a lightweight virtual machine hypervisor
  - such as [Blue Pill](https://en.wikipedia.org/wiki/Blue_Pill_(software))
- External mode:
  - outside the normal operation mode of the targeted system
  - in BIOS or system management mode


Malware Countermeasure Approaches
---
- Ideal solution to the threat of malware is prevention
- Four main elements of prevention:
  - Policy
  - Awareness
  - Vulnerability mitigation
  - Threat mitigation
- If prevention fails, technical mechanisms can be used to support the following threat mitigation options:
  - Detection
  - Identification
  - Removal


[Generations of Anti-Virus Software](https://en.wikipedia.org/wiki/Antivirus_software)
---
- First generation: simple scanners
  - Requires a malware signature to identify the malware
  - Limited to the detection of known malware
- Second generation: heuristic scanners
  - Uses heuristic rules to search for probable malware instances
  - Another approach is integrity checking
- Third generation: activity traps
  - Memory-resident programs that identify malware by its actions rather than its structure in an infected program
- Fourth generation: full-featured protection
  - Packages consisting of a variety of anti-virus techniques used in conjunction
  - Include scanning and activity trap components and access control capability


[Sandbox Analysis](https://en.wikipedia.org/wiki/Sandbox_(computer_security))
---
- Running potentially malicious code in an emulated sandbox or on a virtual machine
- Allows the code to execute in a controlled environment where its behavior can be closely monitored without threatening the security of a real system
- Running potentially malicious software in such environments enables the detection of complex encrypted, polymorphic, or metamorphic malware
- The most difficult design issue with sandbox analysis is to determine how long to run each interpretation


Host-Based Behavior-Blocking Software
---
- Integrates with the operating system of a host computer and monitors program behavior in real time for malicious action 
  - Blocks potentially malicious actions before they have a chance to affect the system
  - Blocks software in real time so it has an advantage over anti-virus detection techniques such as fingerprinting or heuristics
- Limitations
  - Because malicious code must run on the target machine before all its behaviors can be identified, it can cause harm before it has been detected and blocked


Perimeter Scanning Approaches
---
- Anti-virus software typically included in e-mail and Web proxy services running on an organization’s firewall and IDS
- May also be included in the traffic analysis component of an IDS
- May include intrusion prevention measures, blocking the flow of any suspicious traffic
- Approach is limited to scanning malware
- Two types of monitoring software
  - Ingress monitors
    - Located at the border between the enterprise network and the Internet 
    - One technique is to look for incoming traffic to unused local IP addresses
  - Egress monitors
    - Located at the egress point of individual LANs as well as at the border between the enterprise network and the Internet 
    - Monitors outgoing traffic for signs of scanning or other suspicious behavior


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
- [Metaspoit Unleashed](https://www.offsec.com/metasploit-unleashed/)
- [Malware](https://github.com/Samsar4/Ethical-Hacking-Labs/tree/master/6-Malware)
- [Ultimate RAT Collection](https://github.com/yuankong666/Ultimate-RAT-Collection)
  - [NjRat.0.7D](https://github.com/BlackAll9/NjRat.0.7D)