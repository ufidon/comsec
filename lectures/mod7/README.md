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


Source Address Spoofing in Flooding attacks
---
- Forge large volumes of IP packets with random source addresses different from the attacker $A$
  - Usually via the raw socket interface on operating systems
  - Makes $A$ harder to be identified
  - The destination address is set to be the target system $T$
- The ICMP echo responses from $T$ will not be reflected back to $A$
  - For the forged addresses, 
    - some correspond to real systems, they respond with error packets to $T$
    - some not used or unreachable, then ICMP destination unreachable packets might be sent back to $T$ or discarded
    - These returned packets form a second wave of attack on $T$
    - Congestion could occur in the router connected to $T$


Countermeasures against flooding attacks
---
- require flow information along the path from $A$ to $T$
  - it takes time and effort to organize
- this vulnerability born with TCP/IP 
  - which assumes a cooperative, trusting environment
  - no assurance about the source address in a packet really corresponds with the sender $S$
    - impose filtering on routers to ensure this, needs to be imposed
       - as close to $S$ as possible
       - at the borders of the ISP’s providing this connection 
         - unfortunately, many ISPs do not implement such filtering
- [Honeypot](https://en.wikipedia.org/wiki/Honeypot_(computing)) exploits a useful side effect of this scattering of response packets
  - monitors attack traffic for unused IP addresses in backscatter traffic
  - gives valuable information on the type and scale of attacks 

---

## Flooding Attacks

[SYN Flood](https://en.wikipedia.org/wiki/SYN_flood)
---


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