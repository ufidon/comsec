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
   - 💻 A running instance of Suricata
   
3. **Configure Suricata’s initial settings:**
   - After installation, navigate to the `C:\Program Files\Suricata\etc\suricata.yaml` file.
   - Open `suricata.yaml` in a text editor like Notepad, Notepad++, or Visual Studio Code.
   - **Define network interface:** Find the `af-packet` or `winpcap` section in the config. Change it to monitor the network interface linked to VirtualBox’s NAT network. Run `ipconfig` on the Windows VM to find the interface names.
   - 💻 Setting updates


### **Step 2: Configure Suricata Rules**
1. **Download Rule Sets:**
   - Use Suricata Update to download default rule sets. Open a command prompt as Administrator and run:
     ```
     suricata-update
     ```
     - This will download and install the latest rules, including those from Emerging Threats.
     - 💻 update completed

2. **Customizing Rules:**
   - Open the `C:\Program Files\Suricata\rules` directory. You'll find several rule files here (e.g., `emerging-threats.rules`, `suricata.rules`).
   - You can edit or create custom rules to detect specific traffic. For example, add a rule to detect an ICMP ping:
     ```
     alert icmp any any -> any any (msg:"ICMP Test Rule"; sid:1000001; rev:1;)
     ```
   - Save the changes and reload Suricata by `restarting the service`.
   - 💻 custom rules

👉 To restart the Suricata service on Windows and apply updated rules, follow these steps:

1. **Open Command Prompt as Administrator**
   - Press `Win + S`, type "cmd," and right-click **Command Prompt**.
   - Select **Run as administrator** to open the command prompt with elevated privileges.

2. **Restart the Suricata Service**
   - In the command prompt, enter the following commands to stop and start the Suricata service:

     ```shell
     net stop suricata
     net start suricata
     ```

   - This will stop and then restart the Suricata service, applying any updated rule sets.

- **Alternative: Restarting Through Services Manager**
   1. Open the **Services** manager by pressing `Win + R`, typing `services.msc`, and pressing Enter.
   2. Scroll down to find the **Suricata** service.
   3. Right-click on **Suricata** and select **Restart**.

- **Note**
Some installations may have custom names for the Suricata service, like `SuricataService`. You can list all services with `sc query` to confirm the exact name if needed.

- 💻 Suricata service is running

### **Step 3: Start and Verify Suricata**
1. **Start the Suricata Service:**
   - Open Services (type `services.msc` in the Run window) and locate the Suricata service. Start it if it's not running.
   
2. **Verify Logs:**
   - Suricata writes logs to `C:\Program Files\Suricata\log`. Open the `eve.json` file to view detected network events.
   - Use a log viewer like [jq](https://jqlang.github.io/jq/) or [EveBox](https://evebox.org/) for easier reading of JSON logs.
   - 💻 `eve.json` in `jq`.

---

## **Task 2: Install SSH and Enable RDP on Windows Server 2019/2022**

### **Step 1: Install SSH on Windows Server**
1. **Open Windows PowerShell as Administrator.**
2. **Install the OpenSSH Server feature:**
   ```powershell
   Add-WindowsFeature -Name OpenSSH.Server
   ```
   💻 OpenSSH.Server installed
3. **Start and configure the SSH service:**
   - Start the SSH service:
     ```powershell
     Start-Service sshd
     ```
   - Set the service to automatically start on boot:
     ```powershell
     Set-Service -Name sshd -StartupType 'Automatic'
     ```
   - Check OpenSSH.Server is running
     ```powershell
      Get-Service -Name sshd
     ```
   💻 OpenSSH.Server running

4. **Allow SSH through the Windows Firewall:**
   ```powershell
   New-NetFirewallRule -Name sshd -DisplayName 'OpenSSH Server (sshd)' -Enabled True -Direction Inbound -Protocol TCP -Action Allow -LocalPort 22
   ```
5. **Test SSH: login onto Windows From Parrot Linux**
   - Open a terminal on Parrot Linux
   ```bash
   # 1. find services on Windows server
   nmap -A WindowsIP

   # 2. login onto Windows
   ssh username@windows_ip_address
   ```
   - 💻 Windows open ports
   - 💻 login onto Windows

### **Step 2: Enable RDP on Windows Server**
1. **Enable Remote Desktop:**
   - Open the **Server Manager**, click **Local Server**, and look for the **Remote Desktop** field. Click the field to open the Remote Desktop settings.
   - Check **Allow remote connections to this computer**.

2. **Configure Windows Firewall for RDP:**
   - Ensure that RDP is allowed through the firewall by running this command in PowerShell:
     ```powershell
     Enable-NetFirewallRule -DisplayGroup "Remote Desktop"
     ```
3. **Test Windows RDP:**
   - Ensure RDP is running from inside Windows
      ```powershell
      Get-Service -Name TermService
      ```
      💻 TermService is running
   - Ensure RDP is running from outside Windows and accessible from Parrot Linux
      ```bash
      # 1. find services on Windows server
      nmap -A WindowsIP
      ```
      💻 Windows RDP is ready for connection
4. **Access Windows RDP:**

To access a Windows Remote Desktop from a Linux machine, you can use the **Remote Desktop Protocol (RDP)** client available on Linux. The most common RDP client for Linux is `Remmina`, but there are other tools like `rdesktop`, `xfreerdp`, and `Vinagre` that you can use. Here’s how to set it up:

⚠️ Use **Option 1** only. Other options are for your further interest.

- **Option 1: Using Remmina (GUI Tool)**
  1. **Install Remmina**:
     - On Ubuntu or Debian-based systems:
       ```bash
       sudo apt update
       sudo apt install remmina remmina-plugin-rdp -y
       ```
  2. **Open Remmina**:
     - Launch Remmina from your application menu or by running `remmina` in the terminal.

  3. **Set Up a New Connection**:
     - Click on the **+** icon to create a new connection profile.
     - Set the **Protocol** to **RDP - Remote Desktop Protocol**.
     - Enter the **IP address** or **hostname** of your Windows machine in the **Server** field.
     - Enter your **Windows username** and **password**.
     - Save the connection profile if desired, and click **Connect**.
     - 💻 Windows remote desktop

- **Option 2: Using xfreerdp (Command Line Tool)**
  1. **Install xfreerdp**:
     - On Ubuntu or Debian-based systems:
       ```bash
       sudo apt update
       sudo apt install freerdp2-x11 -y
       ```
  2. **Connect to Windows via RDP**:
     - Use the following command to connect:
       ```bash
       xfreerdp /v:windows_ip_address /u:username /p:password
       ```
     - Replace `windows_ip_address` with your Windows machine’s IP address, `username` with your Windows username, and `password` with your password.

- **Option 3: Using rdesktop (Alternative CLI Tool)**
  1. **Install rdesktop**:
     - On Ubuntu or Debian-based systems:
       ```bash
       sudo apt update
       sudo apt install rdesktop -y
       ```

  2. **Connect to Windows via RDP**:
     - Run the following command:
       ```bash
       rdesktop windows_ip_address -u username -p password
       ```
     - Replace `windows_ip_address`, `username`, and `password` as appropriate.

---

## **Task 3: Testing Suricata with Baseline and Malicious Traffic**

### **Step 1: Generating Baseline Traffic (Normal)**
1. **Ping the Windows Server from Parrot Linux:**
   - From Parrot Linux, open the terminal and run a simple ping to the Windows VM's IP:
     ```
     ping <Windows-VM-IP>
     ```
     💻 ping output
   - Suricata should log this ICMP traffic in the `eve.json` file (assuming the test rule for ICMP is active).
      - 💻 logged ICMP traffic.

2. **Perform File Transfer via SSH:**
   - Use `scp` to transfer a file from Parrot to the Windows VM, assuming SSH is enabled on Windows.
     ```
     scp file.txt user@<Windows-VM-IP>:/path/on/windows/
     ```
      💻 transferred file on both Parrot and Windows
   - Monitor Suricata logs for this traffic. It should not generate an alert as it's benign.
      - 💻 logged scp traffic.

### **Step 2: Generating Malicious Traffic**
1. **Perform a Port Scan Using Nmap:**
   - From Parrot Linux, scan the Windows VM for open ports:
     ```bash
     # TCP SYN scan (also known as a half-open scan)
     # You will learn more in Ethical Hacking
     nmap -sS <Windows-VM-IP>
     ```
     💻 Scan results
   - Suricata should trigger alerts for port scanning, as many default rules detect Nmap scans.
      - 💻 Alerts and logs

2. **Brute Force Attack with Hydra:**
   - Use the `hydra` tool on Parrot Linux to brute-force SSH or RDP login attempts (if RDP is enabled on Windows).
     - SSH:
       ```
       hydra -l admin -P /usr/share/wordlists/rockyou.txt ssh://<Windows-VM-IP>
       ```
       💻 Brute-forcing SSH login attempt result
     - RDP:
       ```
       hydra -l admin -P /usr/share/wordlists/rockyou.txt rdp://<Windows-VM-IP>
       ```
       💻 Brute-forcing RDP login attempt results
   - Suricata will generate alerts for brute force attempts.
      - 💻 Alerts and logs

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
   - 💻 Attack results on Parrot Linux
   - 💻 Alerts and logs
---

## **Task 4: Monitoring and Analyzing Suricata Logs**

#### **Step 1: Viewing Logs**
1. **Eve.json:**
   - Open `C:\Program Files\Suricata\log\eve.json` to see the detected events in JSON format.
   - You can use a tool like `jq` on Windows to format the output for easier readability:
     ```
     jq . < eve.json
     ```
   - 💻 Highlight the logged attacks
2. **Real-time Monitoring with EveBox:**
   - Install EveBox for better visualization of Suricata logs.
   - Download and install EveBox from the [official website](https://evebox.org/).
   - Start EveBox and configure it to read the `eve.json` logs:
     ```
     evebox -r C:\Program Files\Suricata\log\eve.json
     ```
   - 💻 Logs in EveBox
#### **Step 2: Analyzing the Traffic**
- Look for the following in the logs:
  - 💻 ICMP traffic from Parrot to Windows (from the ping).
  - 💻 Port scanning alerts from Nmap.
  - 💻 Brute force detection from Hydra.
  - 💻 Exploit attempts from Metasploit.

- Identify patterns in the traffic, false positives, or any missed attacks. This will help you tune Suricata’s rules.

---

## **Task 5: Enhancing the Lab with Firewall and IPS Features**

#### **Step 1: Integrating Windows Firewall with Suricata**
1. **Create Firewall Rules:**
   - Open Windows Defender Firewall with Advanced Security.
   - Create a rule to block incoming traffic from the Parrot Linux IP if Suricata detects malicious activity.

     - Method 1: Using Windows Defender Firewall GUI

       1. **Open Windows Defender Firewall with Advanced Security**:
          - Press `Win + S`, type "Windows Defender Firewall with Advanced Security," and open it.

       2. **Create a New Inbound Rule**:
          - In the left pane, click on **Inbound Rules**.
          - In the right pane, click on **New Rule...** to open the New Inbound Rule Wizard.

       3. **Select Rule Type**:
          - Choose **Custom** and click **Next**.

       4. **Select Program**:
          - Choose **All programs** and click **Next**.

       5. **Specify Protocol and Ports**:
          - Choose **Any** to apply the rule to all protocols and ports, or configure it specifically if needed.
          - Click **Next**.

       6. **Specify the Scope**:
          - Under **Which remote IP addresses does this rule apply to?**, select **These IP addresses**.
          - Click **Add** and enter the IP address of your Parrot Linux machine.
          - Click **OK**, then **Next**.

       7. **Specify the Action**:
          - Select **Block the connection** and click **Next**.

       8. **Specify the Profile**:
          - Choose the profiles where you want this rule to apply (Domain, Private, Public) based on your network type, then click **Next**.

       9. **Name the Rule**:
          - Enter a name (e.g., "Block Parrot Linux Incoming Traffic") and an optional description.
          - Click **Finish**.

     - Method 2: Using PowerShell

         - Alternatively, you can create the rule using PowerShell with the following command:

            ```powershell
            New-NetFirewallRule -DisplayName "Block Parrot Linux Incoming Traffic" -Direction Inbound -Action Block -RemoteAddress parrot_linux_ip
            ```

         - Replace `parrot_linux_ip` with the actual IP address of your Parrot Linux machine.

     - Verification
       - Once the rule is created, test it by attempting to connect from your Parrot Linux machine to the Windows machine. The firewall should now block incoming traffic from that IP. You can view or modify the rule later in the **Windows Defender Firewall with Advanced Security** interface or by using `Get-NetFirewallRule` in PowerShell.
       - 💻 Show the rule
       - 💻 Show Parrot Linux can't connect to the Windows VM any more.

2. **Optional: Automating Firewall Actions with Scripts:**
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
     💻 configuration
   - Restart Suricata and test again using Parrot Linux.
     - 💻 test results

2. **Testing IPS Functionality:**
   - Re-run the malicious traffic tests from Parrot Linux and verify if Suricata not only detects but also actively blocks the traffic.
   - 💻 detections and automatic blocks

---

## **Conclusion**
In this lab, you have successfully installed and configured Suricata on a Windows Server 2019/2022 VM and tested its ability to detect various types of network traffic using Parrot Linux. By generating both normal and malicious traffic, you evaluated Suricata’s effectiveness in logging and alerting on security events. Additionally, you explored integrating Suricata with Windows Firewall and took initial steps to enable Suricata as an IPS.

This lab provided a solid foundation for understanding how IDS/IPS systems like Suricata work in real-world environments and how they can be used to strengthen network security.