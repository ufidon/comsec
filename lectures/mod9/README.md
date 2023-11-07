# Antivirus, Firewalls and Intrusion Prevention Systems
CSPP ch9


🔭 Explore
---
- [What is an antivirus?](https://en.wikipedia.org/wiki/Antivirus_software)
  - [Comparison of antivirus software](https://en.wikipedia.org/wiki/Comparison_of_antivirus_software)


🖊️ Practice
---
- [Explore Windows security](https://www.microsoft.com/en-us/windows/comprehensive-security)
- [Explore, download and install ClamAM on Windows](https://www.clamav.net/)


🦮 Guide
---
- [NIST SP 800-83: Guide to Malware Incident Prevention and Handling for Desktops and Laptops](https://csrc.nist.gov/pubs/sp/800/83/r1/final)


🔭 Explore
---
- [What is a firewall?](https://en.wikipedia.org/wiki/Firewall_(computing))
  - [Comparison of firewalls](https://en.wikipedia.org/wiki/Comparison_of_firewalls)


Defense in depth
---
- Most organizations have internet connectivity today
  - threats are introduced
- Effective *means* of protecting LANs are needed
  - the means could be IDS, IPS, IDPS, firewall, or combined
  - Inserted between the premises network and the Internet to establish a controlled link
    - Used as a perimeter defense
    - Single choke point to impose security and auditing
    - Insulates the internal systems from external networks
  - Can be a single computer system or a set of two or more systems working together


Firewall design goals
---
- All traffic between inside and outside must pass through the firewall
  - Only traffic authorized by security policy will be allowed to pass
- The firewall is immune to penetration


Firewall Access Policy
---
- A critical component in firewall planning and implementation
  - lists the types of traffic authorized to pass through the firewall
    - based on address ranges, protocols, applications and content
- developed from 
  - the organization’s information security risk assessment and policy
  - a broad specification of allowed traffic types
- refined to detail the network traffic filters


Network traffic characteristics used by firewall filters
---
- IP address and protocol values
  - used by packet filter and stateful inspection firewalls
  - Typically used to limit access to specific services
- Application protocol
  - used by an application-level gateway that relays and monitors the exchange of information for specific application protocols
- User identity
  - Typically for inside users who identify themselves using some form of secure authentication technology
- Network activity
  - such as request timestamp, request rate, or other activity patterns


Firewall Capabilities And Limits
---
- Capabilities:
  - Defines a single choke point
    - Provides a location for monitoring security events
    - Integrates several Internet functions that are not security related
  - Can serve as the platform for IPSec and VPN
- Limitations:
  - Cannot protect against attacks bypassing firewall
  - May not protect fully against internal threats
  - Improperly secured wireless LAN can be accessed from outside the organization
  - Infected laptop, PDA, or portable storage device may be used internally


Types of Firewalls
---
- packet filtering firewall
- stateful inspection firewall
- application proxy firewall
- circuit-level proxy firewall


[Packet Filtering Firewall](https://en.wikipedia.org/wiki/Berkeley_Packet_Filter)
---
- Applies rules to each incoming and outgoing IP packet
  - Typically a list of rules based on matches in the IP or TCP header
  - Forwards or discards the packet based on rules match
- Filtering rules are based on information contained in a network packet
  - Source and destination IP address
  - Source and destination port number
  - IP protocol fields
  - mac address specified interface
- Two default policies:
  - Discard - prohibit unless expressly permitted
    - More conservative, controlled, visible to users
  - Forward - permit unless expressly prohibited
    - Easier to manage and use but less secure


🖊️ Practice
---
- [How to Use UFW (Uncomplicated Firewall)?](https://www.baeldung.com/linux/uncomplicated-firewall)
- [UFW](https://help.ubuntu.com/community/UFW)


Packet Filter Advantages And Weaknesses
---
- Advantages
  - Simplicity
  - Typically transparent to users and are very fast
- Weaknesses
  - Cannot prevent attacks that employ application specific vulnerabilities or functions
  - Limited logging functionality
  - Do not support advanced user authentication
  - Vulnerable to attacks on TCP/IP protocol bugs
  - Improper configuration can lead to breaches


🦮 Guide
---
- [NIST SP 800-41: Guidelines on Firewalls and Firewall Policy](https://csrc.nist.gov/pubs/sp/800/41/r1/final)
- [NIST SP 800-94: Guide to Intrusion Detection and Prevention Systems (IDPS)](https://csrc.nist.gov/pubs/sp/800/94/final)


🔭 Explore
---
- [What is an IDPS?](https://en.wikipedia.org/wiki/Intrusion_detection_system)