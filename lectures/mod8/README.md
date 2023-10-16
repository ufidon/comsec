# Intrusion Detection
CSPP ch8


Objectives
---
- Recognize intrusion patterns
- Describe intrusion detection
  - principles and requirements
  - approaches
- Exploit honeypots and Snort


🔭 Explore
---
- [Attackers: List of computer criminals](https://en.wikipedia.org/wiki/List_of_computer_criminals)
- [Attacks: List of security hacking incidents](https://en.wikipedia.org/wiki/List_of_security_hacking_incidents)
- [Aftermath: List of data breaches](https://en.wikipedia.org/wiki/List_of_data_breaches)
- [Defense: United States Cyber Command (USCYBERCOM)](https://en.wikipedia.org/wiki/United_States_Cyber_Command)


[Intruders](https://en.wikipedia.org/wiki/Security_hacker)
---
- trespass IT systems through
  - unauthorized logon or access to machines
  - malware
- also known as hackers, crackers, or computer criminals
- explore methods for breaching defenses 
- exploit weaknesses in a computer system or network
- motivated by profit, protest, challenge, recreation, etc.
- classified into several categories based on motivations
  - Cyber criminals
  - Hacktivists
  - State-sponsored organizations
  - Miscellaneous

Cyber criminals
---
- pursue financial reward
- typical [cybercrimes](https://en.wikipedia.org/wiki/Cybercrime):
  - Identity theft
  - Theft of financial credentials
  - Corporate espionage
  - Data theft and ransoming
- trade malware, stolen data and identities
- coordinate attacks on the internet
- meet in underground forums, [dark webs](https://en.wikipedia.org/wiki/Dark_web) hosted on [darknet](https://en.wikipedia.org/wiki/Darknet) through [Tor](https://en.wikipedia.org/wiki/Tor_(network))


[Hacktivists](https://en.wikipedia.org/wiki/Hacktivism)
---
- motivated by social or political causes
- promote and publicize their causes typically through:
  - Website defacements and redirects
  - Denial of service attacks
  - Theft and distribution of data that results in negative publicity or compromise of their targets


🔭 Explore
---
- [WikiLeaks](https://en.wikipedia.org/wiki/WikiLeaks)
- [Timeline of events associated with Anonymous](https://en.wikipedia.org/wiki/Timeline_of_events_associated_with_Anonymous)


State-sponsored organizations
---
- conduct espionage or sabotage activities
- sponsored by governments such as
  - Russia, USA, UK, and their intelligence allies
- Also known as Advanced Persistent Threats (APTs) due to 
  - the covert nature and persistence over extended periods


Miscellaneous
---
- motivated by technical challenge, peer-group esteem and reputation, etc.
- many of them discovered new categories of buffer overflow vulnerabilities
- some are hobby hackers using attack toolkits to explore system and network security


Three qualitative skill levels of intruders
---
- Apprentices:
  - Also known as “script-kiddies”
  - comprise the largest number of intruders
  - primarily use existing attack toolkits
  - have rudimentary technical skills
- Journeymen:
  - have sufficient technical skills 
  - able to 
    - modify and extend attack toolkits to 
      - exploit newly discovered or purchased vulnerabilities
    - locate and exploit vulnerabilities similar to the known
  - Adapt tools for use by others
- Masters:
  - have high-level technical skills
  - be able to 
    - discover new vulnerabilities
    - Write new powerful attack toolkits
  - Some are employed by state-sponsored organizations


🔭 Explore
---
- [Cyber attack phases and countermeasures](https://en.wikipedia.org/wiki/Kill_chain)
  1. Reconnaissance
  2. Weaponization
  3. Delivery
  4. Exploitation
  5. Installation
  6. Command and Control
  7. Actions on Objective
- [Types of cyber attack](https://en.wikipedia.org/wiki/Cyberattack)


[Intrusion Detection](https://en.wikipedia.org/wiki/Intrusion_detection_system)
---
- security intrusion: 
  - Unauthorized act of bypassing the security mechanisms of a system
- intrusion detection: 
  - A hardware or software function 
  - gathers and analyzes information from various areas within a computer or a network 
  - identifies ­ possible security intrusions


Intrusion Detection System (IDS) components
---
- Sensors
  - collect evidences of an intrusion such as
  - network packets, log files, and system call traces
- Analyzers
  - aggregate the evidences
  - determine whether an intrusion has occurred
  - provide guidance about reactions to the intrusion
- User interfaces let users
  - view reports
  - configure and control the behavior of the system

IDS types
---
- [Host-based IDS (HIDS)](https://en.wikipedia.org/wiki/Host-based_intrusion_detection_system)
  - Installed on a single host 
  - monitors and analyzes the internals of the host and the network packets on its network interfaces
- Network-based IDS (NIDS)
  - Installed on routers or edge computers
  - Monitors and analyzes network traffic through the routers
- Distributed or hybrid IDS
  - Combines information from a number of sensors, 
    - often both host and network based, 
  - feeds in a central analyzer 
  - able to better identify and respond to intrusion activities


  - Basic Principles
  - The Base-Rate Fallacy
  - Requirements
- Analysis Approaches
  - Anomaly Detection
  - Signature or Heuristic Detection
- Host-Based Intrusion Detection
  - Data Sources and Sensors
  - Anomaly HIDS
  - Signature or Heuristic HIDS
  - Distributed HIDS
- Network-Based Intrusion Detection
  - Types of Network Sensors
  - NIDS Sensor Deployment
  - Intrusion Detection Techniques
  - Logging of Alerts
- Distributed or Hybrid Intrusion Detection
- Intrusion Detection Exchange Format
- Honeypots
- Example System: Snort
  - Snort Architecture
  - Snort Rules