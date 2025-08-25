# Setup lab environment

## 1.Description
**In this lab, we will build our own virtual computer security environment managed with VirtualBox, which consists of Parrot Security Linux and Windows server 2025**

* The physical computer you use is called a "Host". The virtual machines that run on it are called "Guests". 

**Prerequisite:**

Your PC must support virtualization which is popular today, following the section 'Check virtualization availability' in the reference to find its availability and enable it.

* Your own PC or laptop needs more than 100GB free disk space for these VMs
  * no less than 8GB RAM
* If you don't have sufficient free space, you may buy an external USB **3.0+** Flash drive or SSD, for example
  * _USB 3.0+ Flash drive_
    * [Samsung BAR Plus 256GB - 300MB/s USB 3.1 Flash Drive Titan Gray (MUF-256BE4/AM)](https://www.amazon.com/Samsung-BAR-Plus-32GB-MUF-32BE4/dp/B07BPKL2D2?ref\_=fsclp\_pl\_dp\_2&th=1)
    * [SanDisk 256GB Extreme PRO USB 3.1 Solid State Flash Drive - SDCZ880-256G-G46](https://www.amazon.com/dp/B01N7QDO7M/ref=emc\_b\_5\_t)
  * _USB 3.0+  External SSD_
    * [SanDisk 500GB Extreme Portable External SSD - Up to 550MB/s - USB-C, USB 3.1 - SDSSDE60-500G-G25](https://www.amazon.com/SanDisk-500GB-Extreme-Portable-External/dp/B078SWJ3CF/ref=sr\_1\_1?dchild=1&keywords=SanDisk\+500GB\+Extreme\+Portable\+External\+SSD\+-\+Up\+to\+550MB%2Fs\+-\+USB-C%2C\+USB\+3\.1\+-\+SDSSDE60-500G-G25&qid=1588950864&s=electronics&sr=1-1)
    * [Samsung (MU-PA500B/AM)T5 Portable SSD - 500GB - USB 3.1 External SSD , Blue ](https://www.amazon.com/Samsung-T5-Portable-SSD-MU-PA500B/dp/B073GZBT36?ref\_=fsclp\_pl\_dp\_3&th=1)

## 2.Steps
To set up the lab using a `VirtualBox image for Parrot Linux` and a `VHD image for Windows Server 2025`, follow the steps below. This approach skips the installation process by using pre-configured virtual disk images.

### **1. Install VirtualBox**

If you haven't already installed VirtualBox, download and install it from the [VirtualBox website](https://www.virtualbox.org/).

### **2. Download the VirtualBox and VHD Images**

- **Parrot Linux**: Download the Parrot Linux VirtualBox image from the [Parrot OS official website](https://www.parrotsec.org). Ensure you choose the `.ova` format, which is a pre-configured VirtualBox appliance.
- **Windows Server 2025**: Download the VHD image from the [Microsoft Evaluation Center](https://www.microsoft.com/en-us/evalcenter/evaluate-windows-server-2025).

### **3. Import the Parrot Linux VirtualBox Image**

1. **Open VirtualBox**.
2. Go to **File** > **Import Appliance**.
3. Browse and select the Parrot Linux `.ova` file you downloaded.
4. Follow the prompts to import the appliance. You can adjust the settings during the import if necessary 
   1. e.g., allocate more RAM or change the virtual disk size
   2. A suggested allocation: 2GB RAM, 60GB disk
5. After the import completes, you should see the Parrot Linux VM listed in VirtualBox.

### **4. Create the Windows Server 2025 VM Using the VHD Image**

1. **Create a New VM**:
   - In VirtualBox, click on **"New"**.
   - Name the VM (e.g., `WindowsServer2025`), select **Type**: Microsoft Windows, and **Version**: Windows 2025 (64-bit).
   - Allocate memory (e.g., 4096 MB).

2. **Attach the VHD Image**:
   - In the **Hard disk** section of the VM creation wizard, choose **Use an existing virtual hard disk file**.
   - Click on the folder icon and browse to select the downloaded Windows Server 2025 `.vhd` file.
   - Finish creating the VM.

### **5. Create and Configure the NAT Network**

1. **Create NAT Network**:
   - For VirtualBox v6 (obsolete since 2024), go to **File** > **Preferences** > **Network** > **NAT Networks**.
     - For VirtualBox v7, this is on the toolbar
   - Click the **"+"** button to create a new NAT network.
   - Edit the NAT Network to ensure it is enabled, and note the IP range (e.g., `10.0.2.0/24`).

2. **Attach VMs to the NAT Network**:
   - For each VM (Parrot Linux and Windows Server 2025), go to **Settings** > **Network**.
   - Set **Adapter 1** to **Attached to**: NAT Network.
   - Select the NAT Network you just created.
   - Enable **Promiscuous** mode: `All all` or `Allow VMs`.

### **6. Start the VMs**

1. **Start the Parrot Linux VM** by selecting it in VirtualBox and clicking **Start**.
2. **Start the Windows Server 2025 VM** similarly.

### **7. Verify Network Configuration**

Once both VMs are running:

#### **Find the IP Addresses**

- **Windows Server 2025**: Open a command prompt and run:
  ```bash
  ipconfig
  ```
  This will display the IP address assigned to the VM.

- **Parrot Linux**: Open a terminal and run:
  ```bash
  ip addr
  ```
  This will display the IP address assigned to the VM.

#### **Test Communication Between VMs**

- From **Parrot Linux**, ping **indows Server 2025**:
  ```bash
  # To enable ping without sudo
  sudo setcap cap_net_raw+p $(which ping)

  ping <Windows Server 2025 IP>
  ```
  - By default, Windows server 2025 disabled its ping echo service. Following this [link](https://kb.iu.edu/d/aopy) to turn on its echo services
    - File and Printer Sharing (Echo Request - ICMPv4-In), and
    - File and Printer Sharing (Echo Request - ICMPv6-In)
- From **Windows Server 2025**, ping **Parrot Linux**:
  ```bash
  ping <Parrot Linux IP>
  ```


### **8. Install VirtualBox Guest Additions**

Guest Additions need to be installed on both VMs to enable bidirectional copy-paste and shared folders. The steps differ slightly for Parrot Linux and Windows Server 2025.

#### **For Parrot Linux**

1. **Start the Parrot Linux VM**.
   - Guest additions may be already installed.
     ```bash
     # check whether the Guest additions are installed
     lsmod | grep vboxguest
     # skip the following steps if you see output like
     # vboxguest 219348  6 vboxsf

     # or check with
     sudo /usr/sbin/VBoxService -V
     ```
2. **Insert the Guest Additions CD**:
   - In VirtualBox, with the VM running, go to **Devices** > **Insert Guest Additions CD Image**.
3. **Install Required Packages**:
   - Open a terminal in Parrot Linux and run:
     ```bash
     sudo apt update
     sudo apt install build-essential dkms linux-headers-$(uname -r)
     ```
4. **Mount and Install Guest Additions**:
   - In the terminal, run:
     ```bash
     sudo mkdir /mnt/cdrom
     sudo mount /dev/cdrom /mnt/cdrom
     sudo /mnt/cdrom/VBoxLinuxAdditions.run
     ```
5. **Reboot the VM**:
   - After installation completes, reboot the VM:
     ```bash
     sudo reboot
     ```

#### **For Windows Server 2025**

1. **Start the Windows Server 2025 VM**.
2. **Insert the Guest Additions CD**:
   - In VirtualBox, with the VM running, go to **Devices** > **Insert Guest Additions CD Image**.
3. **Install Guest Additions**:
   - When the CD is detected, run the installer (`VBoxWindowsAdditions.exe`) from the CD drive.
   - Follow the prompts to install Guest Additions.
   - Reboot the VM after installation.

### **9. Enable Bidirectional Copy-Paste**

1. **With the VM Powered On**:
   - Go to **Devices** > **Shared Clipboard**.
   - Select **Bidirectional**.
   - Go to **Devices** > **Drag and Drop**.
   - Select **Bidirectional**.

Repeat these steps for both the Parrot Linux and Windows Server 2025 VMs.

### **10. Create Shared Folders**

Shared folders allow you to share files between your host system and the guest VMs.

#### **Set Up Shared Folders in VirtualBox**

1. **With the VM Powered Off**:
   - Go to **Settings** > **Shared Folders**.
   - Click the **"+"** icon on the right to add a new shared folder.
   - Choose a folder path on your host machine.
   - Set the **Folder Name** (this is how it will appear in the VM).
   - Check **Auto-mount** and optionally **Make Permanent**.

   Repeat this for both the Parrot Linux and Windows Server 2025 VMs.

#### **Access Shared Folders in Parrot Linux**

1. **After Booting the VM**:
   - The shared folder should automatically appear in the `/media/sf_<folder_name>` directory.
   - If it doesn't appear, add the user to the `vboxsf` group:
     ```bash
     sudo usermod -aG vboxsf $USER
     ```
   - Log out and log back in, or reboot the VM to apply the changes.

#### **Access Shared Folders in Windows Server 2025**

1. **After Booting the VM**:
   - Open **File Explorer**.
   - The shared folder should be listed as a network drive under **This PC**.
   - You can now access the shared folder directly from Windows Explorer.

### **11. Final Verification**

After setting up the shared folders and enabling bidirectional copy-paste:

- **Copy-Paste Test**: Try copying text or files between the host and the guest VMs to verify that bidirectional copy-paste is working.
- **Shared Folders Test**: Place files in the shared folder on the host and verify that they appear in both VMs.

### **Conclusion**

By using a pre-configured VirtualBox image for Parrot Linux and a VHD for Windows Server 2025, you save time on OS installation and quickly set up a virtual lab. After importing the images and configuring the NAT network, the VMs should be able to communicate with each other, which you can verify using ping commands. You've enabled bidirectional copy-paste and configured shared folders for easy file transfer between the host and the guest VMs, improving the usability of your virtual environment.

---

## Alternative environment on [Google Cloud Platform (GCP)](https://cloud.google.com/)

### Create 2 VMs on GCP
- Create a Windows server 2025 VM
- Create a Ubuntu 24.04 LTS VM

### Setup ubuntu user
- Install [Google Cloud SDK (which includes Google Cloud CLI)](https://cloud.google.com/sdk/docs/install)
  - On windows, run the installation file with all default options
    - refer to this [video](https://youtu.be/rpmOM5jJJfY?si=Jm8RDhbWXsc2zjRZ)
  - On Linux
  ```bash
  ## A. on ubuntu/debian
  # 0.  Install required network tools
  sudo apt-get install apt-transport-https ca-certificates gnupg curl

  # 1. Add Google Cloud SDK repository
  echo "deb [signed-by=/usr/share/keyrings/cloud.google.gpg] http://packages.cloud.google.com/apt cloud-sdk main" | sudo tee -a /etc/apt/sources.list.d/google-cloud-sdk.list

  # 2. Import public key
  curl https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -

  # 3. update then install
  sudo apt update && sudo apt install google-cloud-sdk -y

  ## B. other linux distribution
  sudo snap install google-cloud-sdk --classic
  ```
- Initialize GC CLI
  ```bash
  gcloud init
  ```
- Logon ubuntu vm through SSH
  ```bash
  # with gcloud
  gcloud compute ssh USERNAME@VM_NAME --zone ZONE
  # or on webpage
  ```
- Setup a password for a user in SSH
  ```bash
    sudo passwd USERNAME
  ```

### Install and enable rdp on ubuntu
```bash
#!/bin/bash

# 1. update system
sudo apt update && sudo apt upgrade -y

# 2. install XFCE desktop and necessary plugins
sudo apt install xfce4 xfce4-goodies -y

# 3. install xrdp
sudo apt install xrdp -y

# 4. enable then start xrdp service
sudo systemctl enable xrdp
sudo systemctl start xrdp

# 5. configure xrdp to use XFCE session
echo "startxfce4" > ~/.xsession
sudo sed -i.bak '/^new_cursors=true/ s/true/false/' /etc/xrdp/xrdp.ini

# 6. restart xrdp service to apply configuration
sudo systemctl restart xrdp

# 7. allow xrdp sudo
# How to use vim? refer to [Give Me 20 Minutes and I’ll Make You a Vim Motions Expert](https://youtu.be/z4eA2eC28qg)
sudo vim /etc/polkit-1/localauthority.conf.d/02-allow-sudoers.conf
# add the following 4 lines to 02-allow-sudoers.conf
[Allow Sudoers]
Identity=unix-group:sudo
Action=*
ResultActive=yes

# 8. install browsers
sudo apt install chromium-browser firefox
# create firewall rule (not recommended)
#gcloud compute firewall-rules create allow-rdp \
#    --allow tcp:3389 \
#    --network default \
#    --source-ranges 0.0.0.0/0 \
#    --description "Allow RDP"

echo "Installation done！"
echo "GCP IAP Desktop or SSH tunnelling for RDP is recommended，avoid exposing port 3389."

# 9. reboot the system
sudo reboot
```

### Access RDP by GCP IAP
```bash
# 1. login gcloud CLI
gcloud auth login

# 2. enable IAP
gcloud services enable iap.googleapis.com

# 3. Create tunnel via IAP to access RDP
# Change vm name my-vm and zone to yours
gcloud compute start-iap-tunnel my-vm 3389 --local-host-port=localhost:4000 --zone=us-central1-a

# 4. Tunnel to the VM: localhost:4000 → VM:3389
# With RDP client, username/pass: vm logon account
```

⚠️ Remember to POWER OFF all your VMs every time.

---

- **Wireshark Filters**
  - [Capture filters](https://wiki.wireshark.org/CaptureFilters)
  - [Display filters](https://wiki.wireshark.org/DisplayFilters)

## References
* Demo video: [Create virtual machines for Ubuntu & Windows Server 2019 in VirtualBox 6](https://youtu.be/3PbnBVNWXpk)
* [Install the guest additions using the provided ISO from VirtualBox](https://www.parrotsec.org/docs/virtualization/virtualbox-guest-additions)
* [Linux Capabilities and Ping](https://unixetc.co.uk/2016/05/30/linux-capabilities-and-ping/)
* *Projects from samsclass*
  * [Project 1: Kali Virtual Machine](https://samsclass.info/152/proj/123p1kali.htm)
    * [katoolin3](https://github.com/s-h-3-l-l/katoolin3)
  * [Project 2: Windows 2016 Server Virtual Machine](https://samsclass.info/123/proj10/123p2win.htm)
* *Check virtualization availability*
  * [Linux Find Out If CPU Support Intel VT/AMD-V Virtualization For KVM](https://www.cyberciti.biz/faq/linux-xen-vmware-kvm-intel-vt-amd-v-support/)
  * [Easy Ways to Check If Your Processor Supports Virtualization](https://www.technorms.com/8208/check-if-processor-supports-virtualization)
  * [How to find out if Intel VT-x or AMD-V virtualization technology is supported?](https://www.auslogics.com/en/articles/how-to-find-out-if-intel-vt-x-or-amd-v-virtualization-technology-is-supported/)
  * [How to check if Intel Virtualization is enabled without going to BIOS in Windows 10](https://stackoverflow.com/questions/49005791/how-to-check-if-intel-virtualization-is-enabled-without-going-to-bios-in-windows)
  * [How to Enable Intel VT-x in Your Computer’s BIOS or UEFI Firmware](https://www.howtogeek.com/213795/how-to-enable-intel-vt-x-in-your-computers-bios-or-uefi-firmware/)
  * [Virtualization (VT-x/AMD-V) - Enabling virtualization on your computer for running 2N® Access Commander](https://2nwiki.2n.cz/pages/viewpage.action?pageId=75202968)
* *For Mac computers with ARM processor such as m1/m1*
  * [Developer preview for macOS / Arm64 (M1/M2) hosts](https://www.virtualbox.org/wiki/Download_Old_Builds_7_0)
  * [How To Install Parrot OS on Mac M1?](https://minkcoregame.medium.com/how-to-install-parrot-os-on-mac-m1-c1f20438631)
    * [Securely run ANY operating systems on your Mac m1](https://mac.getutm.app/)
* *VM manager*
  * [VMware Workstation Player](https://www.vmware.com/products/workstation-player.html)
    * [Promiscuous Mode by Default?](https://communities.vmware.com/t5/VMware-Workstation-Pro/Promiscuous-Mode-by-Default/td-p/2717392)
    * [Using Virtual Ethernet Adapters in Promiscuous Mode on a Linux Host](https://kb.vmware.com/s/article/287)
  * [VirtualBox](https://www.virtualbox.org/)
  * [Hyper-V](https://docs.microsoft.com/en-us/virtualization/hyper-v-on-windows/quick-start/enable-hyper-v)
* [Top 21 Operating Systems For Ethical Hacking And Pen Testing](https://techlog360.com/top-ethical-hacking-operating-systems/)
  * [Kali Linux](https://www.kali.org/)
* [Linux distribution](https://en.wikipedia.org/wiki/Linux_distribution)
  * [List of Linux distributions](https://en.wikipedia.org/wiki/List_of_Linux_distributions)
  * [Comparison of Linux distributions](https://en.wikipedia.org/wiki/Comparison_of_Linux_distributions)
  * [Light-weight Linux distribution](https://en.wikipedia.org/wiki/Light-weight_Linux_distribution)
* [Windows Server evaluation center](https://www.microsoft.com/en-us/evalcenter/evaluate-windows-server)
* *command line basics*
  * [Linux journey](https://linuxjourney.com/)
  * [Windows commands cheatsheet](https://github.com/ufidon/script/tree/master/lectures/module01)
  * [Linux commands cheatsheet](https://bootlin.com/doc/legacy/command-line/)
* _old os_
  * [WinWorld](https://winworldpc.com/library/operating-systems)
  * [Old versions of Linux](https://soft.lafibre.info/)
  * [archiveos](https://archiveos.org/)
* [Install Kali on VMware](https://samsclass.info/152/proj/M108.htm)
* [Kali Linux and Windows server 2012 VMs](https://drive.google.com/drive/folders/1fT7DlwAQjaDjCRsDDSDtaYZU2sCSLa_v)
  * Windows user/pass: Lab250/toor, Administrator/Admin123
  * Kali user/pass: root/toor
* [Immersive learning games](https://drive.google.com/drive/folders/1lrMrlt7txA1VviePt4koUjyxB6nedtLg)
* _Mac m1/m2/m3/m4_
  * [How to Install and use UTM on Mac: A Guide to Virtual Machines on Mac M1/M2 Apple Silicon Chips](https://youtu.be/i8yUvohhAXw)
    * [How to Install Windows 11 VM on a Mac (M1 / M2 / M3 / M4) with UTM?](https://youtu.be/uis3pvt4wBU)
  * [Install Windows Server 2025 on macOS ARM 64 bit](https://youtu.be/K3QfTnYePxM)
  * [How to Install Parrot Security OS in UTM on MacOS, 2025?](https://youtu.be/_mBuvZIbHl4)