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


Network traffic ingredients used by firewall filters
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


🦮 Guide
---
- [NIST SP 800-41: Guidelines on Firewalls and Firewall Policy](https://csrc.nist.gov/pubs/sp/800/41/r1/final)
- [NIST SP 800-94: Guide to Intrusion Detection and Prevention Systems (IDPS)](https://csrc.nist.gov/pubs/sp/800/94/final)


🔭 Explore
---
- [What is an IDPS?](https://en.wikipedia.org/wiki/Intrusion_detection_system)