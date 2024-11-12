# **Intrusion Detection**

## **Description**
This lab will simulate a real-world scenario where a server is monitored for potential security breaches. Parrot Linux will act as an attacking machine to generate different types of network traffic. The goal is to test how effectively an `open-source intrusion detection system (IDS)` detects malicious activity, such as port scanning, brute force attacks, and exploitation attempts. 

Both VMs (Ubuntu Server and Parrot Linux) will be connected through a VirtualBox NAT network to simulate network traffic while retaining internet access.

## **Lab Prerequisites**
- VirtualBox installed and configured with a **NAT network** where both virtual machines (VMs) are connected.
- **Ubuntu Linux** and **Parrot Linux** VMs installed and properly connected to the same NAT network.
- A basic understanding of IDS concepts and network monitoring.

## **Lab Objectives**
By the end of this lab, you will:
1. Install and configure IDS on a Ubuntu Server VM.
2. Generate both benign and malicious network traffic using Parrot Linux to test IDS detection capabilities.
3. Analyze and interpret IDS logs for traffic detection.


⚠️ For simplicity, you may choose one of the IDSes:

- WireShark for manual detection
- Native Suricata
- Suricata in docker container


---

## **Task 1: Installing and Configuring IDS on Ubuntu**

### **Step 1: Install Ubuntu Server (SEED)**
1. **Install SEED VM:**
   - Visit the [SEED labs 2.0 website](https://seedsecuritylabs.org/labsetup.html) and choose option `Ubuntu 20.04 VM (for Intel/AMD Machines)`.
   - Choose `Approach 1: Use a pre-built SEED VM`,
     - Follow the VM manual to install the VM.
     - default login: 
       - user: seed
       - pass: dees
   - 💻 A running Ubuntu server

### **Step 2: Install an IDS**
#### Option 1: Install WireShark
- Already installed on SEED Ubuntu
- 💻 A running WireShark

#### Option 2: [Native Suricata](https://docs.suricata.io/en/suricata-7.0.7/quickstart.html)
1. **Install Suricata**
  ```bash
  # add suricata repository
  sudo apt-get install software-properties-common
  sudo add-apt-repository ppa:oisf/suricata-stable
  sudo apt-get update

  # install suricata and jq
  sudo apt-get install suricata jq

  # update suricata rules
  sudo suricata-update

  # check suricata service is running
  sudo systemctl status suricata.service
  ```
   - 💻 Suricata service status
2. **Basic setup**
  ```bash
  # determine the interface(s) and IP address(es) on which Suricata should be inspecting network packets:
  ip a

  # set interface to be monitored, HOME_NET
  sudo vim /etc/suricata/suricata.yaml

  # sudo systemctl restart suricata
  ```
3. **Check Suricata is working**
  ```bash
  # open two terminals, one runs
  sudo tail -f /var/log/suricata/fast.log

  # the other runs
  curl http://testmynids.org/uid/index.html
  ```
  
  You should see something in the first terminal similar to
  ```
  [1:2100498:7] GPL ATTACK_RESPONSE id check returned root [**] [Classification: Potentially Bad Traffic] [Priority: 2] {TCP} 217.160.0.187:80 -> 10.0.2.15:41618
  ```
4. **View specific logs**
  ```bash
  # Alerts:
  sudo tail -f /var/log/suricata/eve.json | jq 'select(.event_type=="alert")'

  # Statistics:
  sudo tail -f /var/log/suricata/eve.json | jq 'select(.event_type=="stats")'
  ```

#### Option 3:  **Docker Suricata**
1. Install Simple-IDS
   - Download then install [Simple-IDS - Suricata & EveBox Simply](https://evebox.org/simple-ids/).
   - **Choose network interface to monitor:**
      ```bash
      # Go the the folder contains simple-ids
      # 1. make it executable
      chmod +x ./simple-ids

      # 2. run simple-ids
      # First time running, you are asked to select the network interface to monitor
      ./simple-ids
      # 3. Under the configure menu select the network interface to monitor
      # 4. select "Start" from the main menu then point your browser at http://127.0.0.1:5636
      ```
   - 💻 simple-ids console interface

2. **Download Rule Sets:**
   - Select "stop" from the main menu to stop simple-ids
   - Select "Update rules" from the main menu to update Suricata rules:
     - This will download and install the latest rules, including those from Emerging Threats.

3. **[Basic setup](https://docs.suricata.io/en/latest/quickstart.html#basic-setup):**
   - In main menu: `Other` → `Suricata Shell`, inside `Suricata Shell`,
     - use command line editor [vi](https://bootlin.com/doc/legacy/command-line/vi_memento.pdf) to edit configuration and rule files
     - `vi /etc/suricata/suricata.yaml`, setup HOME_NET
   - Explore popular [rules](https://docs.suricata.io/en/latest/rules/index.html) in `/var/lib/suricata/rules/suricata.rules`.
   - You can edit or create custom rules to detect specific traffic. For example, add a rule to detect an ICMP ping:
     ```
     alert icmp any any -> any any (msg:"xxxxxxx ICMP Test Rule"; sid:1000001; rev:1;)
     ```
   - Save the changes, exit `vi`
      ```bash
      # check configurations and rules
      suricata -T -c /etc/suricata/suricata.yaml
      ```
   - reload Suricata by start it again.

---

## **Task 2: Test Ubuntu SSH, ftp, and telnet servers from Parrot**

### **Step 1: Test SSH on Ubuntu Server**
1. **Check SSH service is available**
   ```bash
   # check sshd demon installed
   which sshd
   # check sshd is listening
   ss -lt | grep ssh
   ```
   💻 SSH service is available
2. **Start, stop, state the SSH service:**
   - Stop the SSH service:
     ```bash
     sudo systemctl stop ssh
     ```
   - Set the service to automatically start on boot:
     ```bash
     sudo systemctl enable ssh
     ```
   - Start the SSH service:
     ```bash
     sudo systemctl start ssh
     ```
   - Check SSH service is running
     ```bash
      sudo systemctl status ssh
     ```
   💻 SSH service running


3. **Test SSH: login onto Ubuntu From Parrot Linux**
   - Open a terminal on Parrot Linux
   ```bash
   # 1. login onto Ubuntu
   # ssh username@Ubuntu_ip_address, e.g.
   ssh seed@Ubuntu_ip
   ```
   - 💻 login onto Ubuntu

### **Step 2: Test telnet on Ubuntu Server**
1. **Check telnet service is available**
   ```bash
   # check inetd demon installed
   which inetd
   # check inetd is listening
   ss -lt | grep telnet
   ```
   💻 telnet service is available
2. **Start, stop, state the telnet service:**
   - Stop the telnet service:
     ```bash
     sudo systemctl stop inetd
     ```
   - Set the service to automatically start on boot:
     ```bash
     sudo systemctl enable inetd
     ```
   - Start the telnet service:
     ```bash
     sudo systemctl start inetd
     ```
   - Check telnet service is running
     ```bash
      sudo systemctl status inetd
     ```
   💻 telnet service running


3. **Test telnet: login onto Ubuntu From Parrot Linux**
   - Open a terminal on Parrot Linux
   ```bash
   # 1. login onto Ubuntu
   # telnet Ubuntu_ip_address, e.g.
   telnet Ubuntu_ip
   ```
   - 💻 login onto Ubuntu

### **Step 3: Test ftp on Ubuntu Server**
1. **Check ftp service is available**
   ```bash
   # check vsftpd demon installed
   which vsftpd
   # check vsftpd is listening
   ss -lt | grep ftp
   ```
   💻 ftp service is available
2. **Start, stop, state the ftp service:**
   - Stop the ftp service:
     ```bash
     sudo systemctl stop vsftpd
     ```
   - Set the service to automatically start on boot:
     ```bash
     sudo systemctl enable vsftpd
     ```
   - Start the ftp service:
     ```bash
     sudo systemctl start vsftpd
     ```
   - Check ftp service is running
     ```bash
      sudo systemctl status vsftpd
     ```
   💻 ftp service running


3. **Test ftp: Access Ubuntu ftp From Parrot Linux**
   - Open a terminal on Parrot Linux
   ```bash
   # 1. login onto Ubuntu ftp
   # ftp Ubuntu_ip_address, e.g.
   ftp Ubuntu_ip
   ```
   - 💻 login onto Ubuntu ftp


---

## **Task 3: Testing IDS with Baseline and Malicious Traffic**
⚠️ Keep your IDS running and capturing throughout this task.

### **Step 1: Generating Baseline Traffic (Normal)**
1. **Ping the Ubuntu Server from Parrot Linux:**
   - From Parrot Linux, open the terminal and run a simple ping to the Ubuntu VM's IP:
     ```
     ping <Ubuntu-VM-IP>
     ```
     💻 ping output
   - IDS should log this ICMP traffic.
      - 💻 logged ICMP traffic.

2. **Perform File Transfer via SSH:**
   - Use `scp` to transfer a file from Parrot to the Ubuntu VM, assuming SSH is enabled on Ubuntu.
     ```
     echo 'hello ubuntu.' > file.txt
     scp file.txt seed@<Ubuntu-VM-IP>:/home/seed/file.txt
     ```
      💻 transferred file on both Parrot and Ubuntu
   - IDS should log this SSH traffic. 
      - 💻 logged scp traffic.

### **Step 2: Generating Malicious Traffic**
1. **Perform a Port Scan Using Nmap:**
   - From Parrot Linux, scan the Ubuntu VM for open ports:
     ```bash
     # TCP SYN scan (also known as a half-open scan)
     # You will learn more in Ethical Hacking
     nmap -sS <Ubuntu-VM-IP>
     ```
     💻 Scan results
   - IDS should log this TCP SYN scan.
      - 💻 logged TCP SYN can

2. **Brute Force Attack with Hydra:**
   - Use the parallelized login cracker [hydra](https://www.kali.org/tools/hydra/) tool on Parrot Linux to brute-force SSH, telnet or ftp login attempts.
   - Download and extract the collection of popular passwords [rockyou](https://www.kaggle.com/datasets/wjburns/common-password-list-rockyoutxt)
     - SSH:
       ```
       hydra -l admin -P ./rockyou.txt ssh://<Ubuntu-VM-IP>
       ```
       💻 Brute-forcing SSH login attempt result
     - telnet:
       ```
       hydra -l admin -P ./rockyou.txt telnet://<Ubuntu-VM-IP>
       ```
       💻 Brute-forcing telnet login attempt result
     - ftp:
       ```
       hydra -l admin -P ./rockyou.txt ftp://<Ubuntu-VM-IP>
       ```
       💻 Brute-forcing ftp login attempt results
   - IDS should log these brute force attempts.
      - 💻 logged brute force attempts.

3. **Exploit Vulnerabilities with Metasploit:**
   - Use Metasploit on Parrot Linux to exploit a known vulnerability on the Ubuntu VM (even if patched, IDS should log the attempt).
   - Example: Search for Ubuntu samba related exploits in Metasploit:
     ```
     msfconsole
     search name:samba platform:linux
     use exploit/multi/samba/usermap_script
     set RHOSTS <Ubuntu-VM-IP>
     exploit
     ```
   - 💻 Attack results on Parrot Linux
   - 💻 logged exploit traffics.
---

## **Task 4: Analyzing IDS Logs**

- Look for the following in the logs:
  - 💻 ICMP traffic from Parrot to Ubuntu (from the ping).
  - 💻 Port scanning alerts from Nmap.
  - 💻 Brute force detection from Hydra.
  - 💻 Exploit attempts from Metasploit.

- Identify patterns in the traffic. This will help you understand and tune IDS’s rules.

---

## **Conclusion**
In this lab, you have successfully installed and configured IDS on a Ubuntu VM and tested its ability to detect various types of network traffic using Parrot Linux. By generating both normal and malicious traffic, you evaluated IDS’s effectiveness in logging and alerting on security events. 

This lab provided a solid foundation for understanding how IDS systems work in real-world environments and how they can be used to strengthen network security.

# Reference
- [How to Fix VirtualBox UUID already exists Errors](https://www.groovypost.com/howto/virtualbox-error-uuid-hard-disk/)
- [Suricata Docker Image](https://github.com/jasonish/docker-suricata)
- [How to set capabilities with setcap command?](https://unix.stackexchange.com/questions/389879/how-to-set-capabilities-with-setcap-command)
- [Testing ping alert rule](https://forum.suricata.io/t/testing-ping-alert-rule/2666/2)