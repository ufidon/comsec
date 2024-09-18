# DoS Attack, Detection and Defense

## 1. Description
In a denial-of-service attack (DoS attack), the attackers prevent some or all legitimate requests from being fulfilled by flooding the targeted machine or network with superfluous requests. The service interrupt may be temporarily or indefinitely. A successful DoS attack causes a breach of the availability of the CIA triad.

DoS attacks are typically classified into [two categories](https://github.com/Samsar4/Ethical-Hacking-Labs/blob/master/9-Denial-of-Service/0-Introduction.md):
- Buffer overflow attacks
- Flood attacks, mainly in the form of distributed denial-of-service attack (DDoS attack)

In this lab, both the Windows Server and Parrot Linux are required.

*Please don't follow any references blindly*

## 2. Tasks: 
### (50%) Task 1: [SYN Flooding Windows web service](https://github.com/Samsar4/Ethical-Hacking-Labs/blob/master/9-Denial-of-Service/1-SYN-Flooding.md)
- On Windows server
  - [Turn off](https://samsclass.info/123/proj10/123p2win.htm) 
    - (5%) firewall and 
    - (5%) antivirus protection
  - [Install IIS](https://computingforgeeks.com/install-and-configure-iis-web-server-on-windows-server/) and show the default website can be accessed both 
    - (5%) locally and 
    - (5%) from Parrot Linux
  - Install [Wireshark](https://www.wireshark.org/)
    - (5%) run Wireshark and start capturing packets
- On Parrot Linux
  - (5%) Launch SYN flooding attack on Windows web service using hping3
- Go back to Windows server
  - (10%) show the attack results in Wireshark
  - Open Task Manager, check the *performance* tab, show 
    - (5%) Ethernet communication and
    - (5%) CPU usage

### (50%) Task 2: [DDoS attack Parrot web service using HOIC](https://github.com/Samsar4/Ethical-Hacking-Labs/blob/master/9-Denial-of-Service/2-DDoS-using-HOIC.md)
- On Parrot 
  - Launch a simple HTTP server
    ```bash
    # (5%) 1. launch a simple HTTP server
    python3 -m http.server 80
    # (5%) 2. access locally
    http://localhost
    # (5%) 3. access from Windows: http://ParotIP
    ```
  - *In case you cannot access the HTTP server from Windows 2019*, you may try installing other Windows and connect them to the same NAT network as Parrot such as
    - [Windows server 2022](https://www.microsoft.com/en-us/evalcenter/evaluate-windows-server-2022)
    - [Windows 11 development environment](https://developer.microsoft.com/en-us/windows/downloads/virtual-machines/)
    - [Windows 10 Enterprise](https://www.microsoft.com/en-us/evalcenter/download-windows-10-enterprise)    
  - (5%) Run Wireshark and start capturing packets
- (20%) On Windows server, 
  - download and install [High Orbit Ion Cannon](https://en.wikipedia.org/wiki/High_Orbit_Ion_Cannon) 
  - Run HOIC and add the target as Parrot
    - URL: http://ParrotIP
    - Power: high
    - Booster: GenericBoost.hoic
  - THREADS: 30
  - FIRE THE LAZER!
- Go back to Parrot Linux,
  - (10%) show the attack results in Wireshark

## 3. Extra credits (10%): 
[Raven-Storm](https://github.com/Tmpertor/Raven-Storm) is a powerful DDoS toolkit for penetration tests. It can launch DoS attacks on several network layers through modules. In this extra task, attack Windows from the Parrot Linux.
- (2%) Install Raven-Storm on Parrot Linux
- (8% in total, 2% for each) Launch the following modules to attack Windows
  - l3 for ping
  - l4 for udp/tcp services
  - l7 for websites
  - arp for local devices
- For each attack, show at least two screenshots
  - one on Parrot shows the attack
  - the other on Windows shows the result 


## 4. Report
Write a report about the process you complete the tasks above, key screen snapshots are needed as evidences.

# References
- [SYN Flooding](https://github.com/Samsar4/Ethical-Hacking-Labs/blob/master/9-Denial-of-Service/1-SYN-Flooding.md)
- [DDoS attack using HOIC](https://github.com/Samsar4/Ethical-Hacking-Labs/blob/master/9-Denial-of-Service/2-DDoS-using-HOIC.md)
- [Detecting DoS Attack traffic](https://github.com/Samsar4/Ethical-Hacking-Labs/blob/master/9-Denial-of-Service/3-Detecting-DoS-Traffic.md)
- [Anti DDoS Guardian](https://www.anti-ddos.net/)