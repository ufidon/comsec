# Denial-of-Service (DoS) Attack

🔭 Explore
---
- [Top 10 DDoS Attacks](https://socradar.io/top-10-ddos-attacks/)


[What is Denial-of-Service Attack](https://en.wikipedia.org/wiki/Denial-of-service_attack)
---
- The [NIST Computer Security Incident Handling Guide](https://csrc.nist.gov/pubs/sp/800/61/r2/final) defines a DoS  attack as
- An action that 
  - prevents or impairs the authorized use of networks, systems, or applications 
  - by exhausting resources such as central processing units (CPU), memory, bandwidth, and disk space
- Attacks on the availability of services and resources
- Popular form
  - Distributed denial-of-service (DDoS) attack


DoS attack targets
---
- Network bandwidth
  - The capacity of the network links connecting a server to the Internet
  - For most organizations this is their connection to their Internet Service Provider (ISP)
- System resources
  - Overload or crash the network handling software
  - by exhausting the client connections a server can provide
- Application resources
  -  Consume significant resources by valid requests 
  -  Exhaust the response capability of the server 


Flooding attack with ping command
---
- A classic denial-of-service attack
- Overwhelm the capacity of the network connection to the target organization
- Traffic can be handled by higher capacity links on the path
  - but packets are discarded as capacity decreases
- Source of the attack is clearly identified 
  - unless a spoofed address is used
- Network performance is noticeably affected


Source Address Spoofing
---
- Forge random and different source addresses in flooding IP packets
  - Usually via the raw socket interface on operating systems
  - Makes attacking systems harder to be identified
- set the destination address to be the target system  
- Congestion would result in the router connected to the final, lower capacity link
- Requires network engineers to specifically query flow information from their routers
- Backscatter traffic
  - Advertise routes to unused IP addresses to monitor attack traffic


SYN Spoofing
Flooding Attacks
ICMP Flood
UDP Flood
TCP SYN Flood
Distributed Denial-of-Service Attacks
Application-Based Bandwidth Attacks
SIP Flood
HTTP-Based Attacks
Reflector and Amplifier Attacks
Reflection Attacks
Amplification Attacks
DNS Amplification Attacks
Defenses Against Denial-of-Service Attacks
Responding to a Denial-of-Service Attack