## Intrusion Prevention

In this lab, we learn configuring Windows firewall rules, monitoring firewall activity, and troubleshooting firewall-related issues.

### **Windows Firewall Features and Capabilities**

**Windows Defender Firewall with Advanced Security** is a powerful tool built into Windows operating systems, including Windows Server 2019/2022, to manage and control inbound and outbound network traffic. 

Windows Firewall offers flexibility, allowing you to finely control traffic flows both to and from your server, thus improving the server’s security posture.

Its core functions are to protect the server from unauthorized access and help secure network communications. Key features include:

#### **1. Three Profiles for Different Network Environments**
- **Domain**: Applied when the server is connected to a corporate domain network.
- **Private**: Used for trusted networks like home or small office networks.
- **Public**: Used for public networks like Wi-Fi at a coffee shop. It is the most restrictive profile.

Each profile allows different rules to be applied based on the network's security requirements.

#### **2. Inbound and Outbound Rules**
- **Inbound Rules**: Control incoming traffic to the server. These rules can allow or block traffic from specific ports, programs, or IP addresses.
- **Outbound Rules**: Control outgoing traffic from the server to external networks or other devices.

#### **3. Predefined Rules**
Windows Firewall has a set of predefined rules for common services such as **Remote Desktop**, **File Sharing**, **ICMP** (ping), and more, which can be easily enabled or customized.

#### **4. Granular Control**
- **Port-Based Rules**: Allow or block traffic based on port numbers (e.g., HTTP on port 80, HTTPS on port 443).
- **Program-Based Rules**: Control network traffic for specific programs or executables.
- **Protocol-Based Rules**: Rules based on protocols like TCP, UDP, ICMP, etc.
- **IP-Based Rules**: Filter traffic by specific IP addresses or IP ranges.

#### **5. Connection Security Rules (IPsec)**
Windows Firewall supports **IPsec (Internet Protocol Security)**, which can be used to secure communications between hosts on a network by encrypting traffic and enforcing authentication for network connections.

#### **6. Logging and Monitoring**
Windows Firewall provides logging capabilities to monitor successful and dropped connections. Logs can help track what traffic is allowed or blocked, which is crucial for troubleshooting.

#### **7. Network Isolation**
Firewall can enforce **network isolation** rules that control access to the server based on membership in certain network subnets, IP ranges, or based on the computer's domain status.

#### **8. Group Policy Integration**
In domain environments, firewall settings can be configured and enforced centrally through **Group Policy**, ensuring that all machines in the domain adhere to the same security policies.

### **Lab Overview**

- **Environment**: Two VMs connected in a NAT network in VirtualBox:
  - Windows Server 2019/2022 with IIS installed and a website hosted **done in Lab05**
  - Parrot Linux
- **Goals**: 
  1. Configure Windows Firewall.
  2. Test configurations using Parrot Linux.
  3. Monitor and troubleshoot firewall settings.

---

### **Task 1: Configure Basic Windows Firewall Settings**

1. **Launch Windows Defender Firewall** on the Windows Server VM.
   - Open **Windows Defender Firewall with Advanced Security**.

2. **Configure Inbound Rules**:
   - **Allow HTTP (Port 80)**: Create a rule to allow incoming web traffic.
   - **Block ICMP (Ping)**: Create a rule to block ICMP traffic.

#### **Step-by-Step Guide to Configuring Inbound Rules**

##### **1. Open Windows Defender Firewall with Advanced Security**
- Press **Windows + R**, type `wf.msc`, and press **Enter**.
- This opens the **Windows Defender Firewall with Advanced Security** console.

##### **2. Navigate to Inbound Rules**
- On the left pane, click **Inbound Rules** to view a list of all active inbound rules.

##### **3. Create a New Inbound Rule**
- On the right-hand side, click **New Rule**. This launches the **New Inbound Rule Wizard**.
  
##### **4. Choose Rule Type**
In the wizard, choose the type of rule you want to create:
- **Program**: Controls traffic for a specific application.
- **Port**: Controls traffic to a specific TCP/UDP port.
- **Predefined**: Choose from predefined rules for services like **Remote Desktop** or **File Sharing**.
- **Custom**: Create a custom rule based on specific conditions (protocol, program, port, IP address).

For example, to allow HTTP traffic:
- Select **Port** and click **Next**.

##### **5. Configure Port**
- Choose **TCP** or **UDP** based on your requirement (e.g., HTTP uses TCP).
- Specify the port number (e.g., **80** for HTTP).
- Click **Next**.

##### **6. Choose Action**
- **Allow the connection**: Permit the traffic.
- **Block the connection**: Deny the traffic.

Choose **Allow the connection** and click **Next**.

##### **7. Choose the Profile**
Select the profiles (Domain, Private, Public) to which this rule should apply, depending on the network environment. Typically, if you're running a web server, you might allow traffic on **Private** and **Public**.

##### **8. Name the Rule**
- Give the rule a descriptive name (e.g., **Allow HTTP Traffic**).
- Click **Finish** to create the rule.

##### **9. Test the Rule**
After creating the rule, test it by using a client machine (like your **Parrot Linux VM**) to access the service. For example, if you created an HTTP rule, you could access it using `curl` or a web browser.

3. **Configure Outbound Rules**:
   - **Block All Traffic Except HTTP**: Create a rule that blocks all outbound traffic except on port 80 (HTTP).

#### **Step-by-Step Guide to Configuring Outbound Rules**

##### **1. Open Outbound Rules**
- On the left pane, click **Outbound Rules** to see the current outbound rules.

##### **2. Create a New Outbound Rule**
- Click **New Rule** from the right-hand side to launch the **New Outbound Rule Wizard**.

##### **3. Choose Rule Type**
Select the appropriate type (e.g., **Port** or **Program**).

For example, to block all traffic except HTTP:
- Select **Port** and click **Next**.

##### **4. Configure Port**
- Choose **TCP** and specify port **80** (for HTTP).
- Click **Next**.

##### **5. Choose Action**
- Select **Allow the connection** for traffic on port 443.
- Click **Next**.

##### **6. Choose Profile**
Select the profiles (Domain, Private, Public) this rule should apply to.

##### **7. Name the Rule**
Give the rule a name (e.g., **Allow HTTP Outbound**).
- Click **Finish**.

##### **8. Block All Other Outbound Traffic**
- To block all other outbound traffic, create a separate **block rule**:
  - Repeat the process, but instead of selecting **Allow**, choose **Block** for **All Ports** except **443**.

##### **9. Test the Rule**
After creating the outbound rule, try accessing external websites using HTTP (port 80)from a client machine. Only HTTP connections should succeed.

4. **Switch Between Firewall Profiles**:
   - Test different profiles (Domain, Private, Public) to observe how firewall settings differ across network environments.

---

### **Task 2: Test Windows Firewall Settings Using Parrot Linux**

1. **Basic Connectivity Tests**:
   - **Ping the Windows Server** from Parrot Linux:
     ```bash
     ping <Windows_Server_IP>
     ```
     Ensure ICMP is blocked.

   - **Access HTTP Services**:
     - Use curl from Parrot Linux to access the Windows Server:
       ```bash
       curl http://<Windows_Server_IP>
       ```
     - Confirm that port 80 (HTTP) traffic is allowed.

2. **Nmap Scanning**:
   - Run an **nmap** scan from Parrot Linux:
     ```bash
     nmap -Pn <Windows_Server_IP>
     ```
     Ensure only allowed services (e.g., HTTP, HTTPS) are visible.

---

### **Task 3: Monitor Firewall Activity on Windows Server**

1. **Enable Firewall Logging**:
   - Go to **Monitoring** in the **Windows Defender Firewall** console and configure logging for both dropped packets and successful connections.

2. **Review Firewall Logs**:
   - Logs are located at: `C:\Windows\System32\LogFiles\Firewall\pfirewall.log`.
   - Use PowerShell to view logs in real time:
     ```powershell
     Get-Content -Path "C:\Windows\System32\LogFiles\Firewall\pfirewall.log" -Wait
     ```

3. **Monitor in Real-Time**:
   - Use **Resource Monitor** to view real-time network activity and observe blocked or allowed connections.

---

### **Task 4: Troubleshoot Firewall-Related Issues**

1. **Simulate Blocked RDP Connection**:
   - **Block RDP Traffic**: Create a rule to block inbound RDP on port 3389.
   - Attempt to connect via RDP from Parrot Linux, then troubleshoot and unblock the service.

2. **Troubleshoot HTTP Service Unavailability**:
   - **Block HTTP Traffic**: Create a rule to block port 80.
   - Try accessing the web service from Parrot Linux. Investigate using logs and remove the block to restore service.

3. **Resolve ICMP Block**:
   - Disable the ICMP blocking rule temporarily, test with ping, and then reapply the rule to allow ICMP from a specific IP range.

---

### **Task 5: Automate Firewall Configuration and Export Rules**

1. **Automate Rule Creation with PowerShell**:
   - Use PowerShell to create and apply firewall rules:
     ```powershell
     New-NetFirewallRule -DisplayName "Allow HTTP" -Direction Inbound -Action Allow -Protocol TCP -LocalPort 80
     ```

2. **Export and Import Firewall Configurations**:
   - Export firewall rules:
     ```powershell
     netsh advfirewall export "C:\path\to\firewall_config.wfw"
     ```
   - Import firewall rules on another server:
     ```powershell
     netsh advfirewall import "C:\path\to\firewall_config.wfw"
     ```

---

### **Lab Wrap-Up**

By completing these tasks, you will:
- Create, configure, and test both inbound and outbound firewall rules.
- Use penetration testing tools to evaluate the firewall’s strength.
- Monitor firewall activity using logs and built-in tools.
- Troubleshoot and resolve firewall-related issues efficiently.
