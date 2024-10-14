# **Intrusion Detection**

## **Description**
In this lab, we will configure Suricata, an open-source intrusion detection system (IDS), on a Windows Server 2019/2022 virtual machine. The lab will simulate a real-world scenario where a server is monitored for potential security breaches. Parrot Linux will act as an attacking machine to generate different types of network traffic. The goal is to test how effectively Suricata detects malicious activity, such as port scanning, brute force attacks, and exploitation attempts.

Both VMs (Windows Server and Parrot Linux) will be connected through a VirtualBox NAT network to simulate network traffic while retaining internet access. This lab also explores integrating Suricata with Windows Firewall and introduces the possibility of running Suricata in Intrusion Prevention System (IPS) mode.

## **Lab Prerequisites**
- VirtualBox installed and configured with a **NAT network** where both virtual machines (VMs) are connected.
- **Windows Server 2019 or 2022** and **Parrot Linux** VMs installed and properly connected to the same NAT network.
- A basic understanding of IDS/IPS concepts and network monitoring.

## **Lab Objectives**
By the end of this lab, you will:
1. Install and configure Suricata on a Windows Server VM.
2. Install and configure SSH and enable RDP on the Windows Server.
3. Generate both benign and malicious network traffic using Parrot Linux to test Suricata’s detection capabilities.
4. Analyze and interpret Suricata logs for traffic detection.
5. Optionally integrate Windows Firewall with Suricata to block detected threats.
6. Explore Suricata’s potential as an Intrusion Prevention System (IPS).

---

## **Task 1: Installing and Configuring Suricata on Windows Server 2019/2022**

### **Step 1: Install Suricata on Windows Server**
1. **Download Suricata:**
   - Visit the [Suricata website](https://suricata.io/download/) and download the latest version for Windows.
   
2. **Install Suricata:**
   - Run the installer. During the installation, you’ll be asked to select the components to install. Choose:
     - Suricata executable
     - Suricata Update (for downloading rules)
   
3. **Configure Suricata’s initial settings:**
   - After installation, navigate to the `C:\Program Files\Suricata\etc\suricata.yaml` file.
   - Open `suricata.yaml` in a text editor like Notepad++.
   - **Define network interface:** Find the `af-packet` or `winpcap` section in the config. Change it to monitor the network interface linked to VirtualBox’s NAT network. Run `ipconfig` on the Windows VM to check the interface names.


### **Step 2: Configure Suricata Rules**
1. **Download Rule Sets:**
   - Use Suricata Update to download default rule sets. Open a command prompt as Administrator and run:
     ```
     suricata-update
     ```
   - This will download and install the latest rules, including those from Emerging Threats.

2. **Customizing Rules:**
   - Open the `C:\Program Files\Suricata\rules` directory. You'll find several rule files here (e.g., `emerging-threats.rules`, `suricata.rules`).
   - You can edit or create custom rules to detect specific traffic. For example, add a rule to detect an ICMP ping:
     ```
     alert icmp any any -> any any (msg:"ICMP Test Rule"; sid:1000001; rev:1;)
     ```
   - Save the changes and reload Suricata by restarting the service.


### **Step 3: Start and Verify Suricata**
1. **Start the Suricata Service:**
   - Open Services (type `services.msc` in the Run window) and locate the Suricata service. Start it if it's not running.
   
2. **Verify Logs:**
   - Suricata writes logs to `C:\Program Files\Suricata\log`. Open the `eve.json` file to view detected network events.
   - Use a log viewer like `jq` or EveBox for easier reading of JSON logs.

---

## **Task 2: Install SSH and Enable RDP on Windows Server 2019/2022**

### **Step 1: Install SSH on Windows Server**
1. **Open Windows PowerShell as Administrator.**
2. **Install the OpenSSH Server feature:**
   ```powershell
   Add-WindowsFeature -Name OpenSSH.Server
   ```
3. **Start and configure the SSH service:**
   - Start the SSH service:
     ```powershell
     Start-Service sshd
     ```
   - Set the service to automatically start on boot:
     ```powershell
     Set-Service -Name sshd -StartupType 'Automatic'
     ```

4. **Allow SSH through the Windows Firewall:**
   ```powershell
   New-NetFirewallRule -Name sshd -DisplayName 'OpenSSH Server (sshd)' -Enabled True -Direction Inbound -Protocol TCP -Action Allow -LocalPort 22
   ```

### **Step 2: Enable RDP on Windows Server**
1. **Enable Remote Desktop:**
   - Open the **Server Manager**, click **Local Server**, and look for the **Remote Desktop** field. Click the field to open the Remote Desktop settings.
   - Check **Allow remote connections to this computer**.

2. **Configure Windows Firewall for RDP:**
   - Ensure that RDP is allowed through the firewall by running this command in PowerShell:
     ```powershell
     Enable-NetFirewallRule -DisplayGroup "Remote Desktop"
     ```

---

## **Task 3: Testing Suricata with Baseline and Malicious Traffic**

### **Step 1: Generating Baseline Traffic (Normal)**
1. **Ping the Windows Server from Parrot Linux:**
   - From Parrot Linux, open the terminal and run a simple ping to the Windows VM's IP:
     ```
     ping <Windows-VM-IP>
     ```
   - Suricata should log this ICMP traffic in the `eve.json` file (assuming the test rule for ICMP is active).

2. **Perform File Transfer via SSH:**
   - Use `scp` to transfer a file from Parrot to the Windows VM, assuming SSH is enabled on Windows.
     ```
     scp file.txt user@<Windows-VM-IP>:/path/on/windows/
     ```

   - Monitor Suricata logs for this traffic. It should not generate an alert as it's benign.

### **Step 2: Generating Malicious Traffic**
1. **Perform a Port Scan Using Nmap:**
   - From Parrot Linux, scan the Windows VM for open ports:
     ```bash
     nmap -sS <Windows-VM-IP>
     ```
   - Suricata should trigger alerts for port scanning, as many default rules detect Nmap scans.

2. **Brute Force Attack with Hydra:**
   - Use the `hydra` tool on Parrot Linux to brute-force SSH or RDP login attempts (if RDP is enabled on Windows).
     - SSH:
       ```
       hydra -l admin -P /usr/share/wordlists/rockyou.txt ssh://<Windows-VM-IP>
       ```
     - RDP:
       ```
       hydra -l admin -P /usr/share/wordlists/rockyou.txt rdp://<Windows-VM-IP>
       ```
   - Suricata will generate alerts for brute force attempts.

3. **Exploit Vulnerabilities with Metasploit:**
   - Use Metasploit on Parrot Linux to exploit a known vulnerability on the Windows VM (even if patched, Suricata should log the attempt).
   - Example: Search for Windows SMB-related exploits in Metasploit:
     ```
     msfconsole
     search smb
     use exploit/windows/smb/ms17_010_eternalblue
     set RHOSTS <Windows-VM-IP>
     exploit
     ```

---

## **Task 4: Monitoring and Analyzing Suricata Logs**

#### **Step 1: Viewing Logs**
1. **Eve.json:**
   - Open `C:\Program Files\Suricata\log\eve.json` to see the detected events in JSON format.
   - You can use a tool like `jq` on Windows to format the output for easier readability:
     ```
     jq . < eve.json
     ```

2. **Real-time Monitoring with EveBox:**
   - Install EveBox for better visualization of Suricata logs.
   - Download and install EveBox from the [official website](https://evebox.org/).
   - Start EveBox and configure it to read the `eve.json` logs:
     ```
     evebox -r C:\Program Files\Suricata\log\eve.json
     ```

#### **Step 2: Analyzing the Traffic**
- Look for the following in the logs:
  - ICMP traffic from Parrot to Windows (from the ping).
  - Port scanning alerts from Nmap.
  - Brute force detection from Hydra.
  - Exploit attempts from Metasploit.

- Identify patterns in the traffic, false positives, or any missed attacks. This will help you tune Suricata’s rules.

---

## **Task 5: Enhancing the Lab with Firewall and IPS Features**

#### **Step 1: Integrating Windows Firewall with Suricata**
1. **Create Firewall Rules:**
   - Open Windows Defender Firewall with Advanced Security.
   - Create a rule to block incoming traffic from the Parrot Linux IP if Suricata detects malicious activity.

2. **Automating Firewall Actions with Scripts:**
   - Write a PowerShell script to monitor Suricata logs and automatically add the attacker’s IP to the Windows Firewall block list:
     ```powershell
     $log = "C:\Program Files\Suricata\log\eve.json"
     $attackerIP = Get-Content $log | Select-String "src_ip" | ForEach-Object { $_.Matches.Groups[1].Value }
     New-NetFirewallRule -DisplayName "Block Attacker" -Direction Inbound -RemoteAddress $attackerIP -Action Block
     ```

### **Step 2: Switching Suricata to IPS Mode**
1. **Switch to IPS Mode:**
   - Modify the Suricata configuration (`suricata.yaml`) to enable IPS functionality (if supported on Windows):
     ```yaml
     - drop:
         src_ip: any
         action: drop
     ```
   - Restart Suricata and test again using Parrot Linux.

2. **Testing IPS Functionality:**
   - Re-run the malicious traffic tests from Parrot Linux and verify if Suricata not only detects but also actively blocks the traffic.

---

## **Conclusion**
In this lab, you have successfully installed and configured Suricata on a Windows Server 2019/2022 VM and tested its ability to detect various types of network traffic using Parrot Linux. By generating both normal and malicious traffic, you evaluated Suricata’s effectiveness in logging and alerting on security events. Additionally, you explored integrating Suricata with Windows Firewall and took initial steps to enable Suricata as an IPS.

This lab provided a solid foundation for understanding how IDS/IPS systems like Suricata work in real-world environments and how they can be used to strengthen network security.