To install **IIS (Internet Information Services)** on **Windows Server 2019/2022** and enable the default HTTP website, follow this step-by-step guide.

---

### **Step 1: Install IIS on Windows Server**

1. **Open Server Manager**:
   - Click the **Start** button, type **Server Manager**, and open it.

2. **Add Roles and Features**:
   - In **Server Manager**, click **Manage** at the top right and select **Add Roles and Features**.

3. **Start the Wizard**:
   - The **Add Roles and Features Wizard** will appear. Click **Next** to proceed.

4. **Choose Installation Type**:
   - Select **Role-based or feature-based installation** and click **Next**.

5. **Select the Destination Server**:
   - Choose the server from the list (it will usually be the local server).
   - Click **Next**.

6. **Select Server Roles**:
   - Scroll down and check **Web Server (IIS)**.
   - A pop-up will appear asking if you want to add features required for IIS. Click **Add Features**.
   - Click **Next**.

7. **Select Features**:
   - On the **Features** page, you can leave the default selections or add any additional features required for your setup (e.g., **.NET Framework** features, **WebSocket Protocol**).
   - Click **Next**.

8. **Web Server Role (IIS) Overview**:
   - You will see an overview of the IIS role. Just click **Next** to continue.

9. **Select IIS Role Services**:
   - Leave the default options for the IIS role services (HTTP, static content, etc.), or optionally add additional services like **FTP Server**, **Web Management Tools**, or **Application Development** features (e.g., ASP.NET).
   - Click **Next**.

10. **Confirm Installation Selections**:
    - Review your selections and click **Install**.
    - The installation process will begin and may take a few minutes.

11. **Complete Installation**:
    - Once the installation is finished, click **Close**.

---

### **Step 2: Verify IIS Installation**

1. **Open IIS Manager**:
   - Click the **Start** button, type **IIS**, and select **Internet Information Services (IIS) Manager** from the results.
   
2. **Verify Default Website**:
   - In **IIS Manager**, expand the server node in the left-hand pane and click on **Sites**.
   - You should see a **Default Web Site** already created.

---

### **Step 3: Start and Enable the Default Website**

1. **Start the Default Website**:
   - In the **Sites** node of **IIS Manager**, right-click on **Default Web Site** and select **Start** if it’s not already running.

2. **Check Binding Settings**:
   - Right-click on **Default Web Site** and select **Edit Bindings**.
   - Make sure there is a binding for **http** on port **80**:
     - **Type**: HTTP
     - **Host name**: Leave blank (optional, unless you want to bind to a specific domain).
     - **Port**: 80 (default for HTTP)
     - **IP address**: Leave as **All Unassigned** unless you want to bind to a specific IP address.
   - If the binding is missing, click **Add**, set the values, and click **OK**.

---

### **Step 4: Verify the Default Website**

1. **Open a Web Browser**:
   - Open a web browser (either from the server or from another machine on the same network).
   
2. **Access the Default Website**:
   - Type the server’s IP address or hostname in the browser’s address bar:
     ``` 
     http://<server_IP>
     ```
     - For example, if the server’s IP is `192.168.1.100`, enter:
       ```
       http://192.168.1.100
       ```

3. **Default IIS Page**:
   - You should see the default **IIS Welcome Page**, which confirms that IIS is installed and running correctly.
     - The page will typically display a message like **"Welcome to Internet Information Services (IIS)"**.

---

### **Optional: Allow HTTP Traffic Through the Firewall**

If you cannot access the website, it might be because HTTP traffic is blocked by the Windows Firewall. Here’s how to allow it:

1. **Open Windows Defender Firewall**:
   - Open **Control Panel**, go to **System and Security**, then **Windows Defender Firewall**.

2. **Allow HTTP (Port 80)**:
   - Click **Advanced Settings** in the left pane.
   - In the **Windows Defender Firewall with Advanced Security** console, click **Inbound Rules** on the left.
   - On the right-hand side, click **New Rule**.
   - Select **Port** and click **Next**.
   - Select **TCP** and enter **80** in the **Specific local ports** field. Click **Next**.
   - Choose **Allow the connection** and click **Next**.
   - Select when the rule applies (Domain, Private, Public) and click **Next**.
   - Give the rule a name, like **Allow HTTP Traffic**, and click **Finish**.

This will allow inbound traffic on port 80, enabling users to access the website.

---

### **Step 5: Modify or Add Website Content (Optional)**

The default IIS site serves content from a specific folder on the server. You can replace or modify this content as needed:

1. **Default Content Folder**:
   - The default website content is located at:
     ```
     C:\inetpub\wwwroot
     ```

2. **Change Content**:
   - Replace or add new files (like an `index.html` file) in the `wwwroot` folder.
   - This will immediately change the content that users see when they access the website.

---

### **Conclusion**

You have successfully installed **IIS** on **Windows Server 2019/2022** and enabled the default **HTTP** website. Users can now access the site via the server’s IP address or hostname. You can further customize the site by adding more websites, configuring virtual directories, and applying SSL certificates for HTTPS.