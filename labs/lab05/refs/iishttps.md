To configure **HTTPS** on **Windows Server 2019/2022** with **IIS (Internet Information Services)**, you’ll need to install IIS, obtain an SSL certificate, and configure the server to serve content over HTTPS.

Here’s a step-by-step guide:

---

### **Step 1: Install IIS on Windows Server**

1. **Open Server Manager**:
   - Click the **Windows icon** and open **Server Manager**.

2. **Add Roles and Features**:
   - In the **Server Manager** Dashboard, click **Add roles and features**.
   - The **Add Roles and Features Wizard** will appear. Click **Next**.

3. **Select Installation Type**:
   - Choose **Role-based or feature-based installation** and click **Next**.

4. **Select the Destination Server**:
   - Select the server from the list (usually the local server) and click **Next**.

5. **Select Server Roles**:
   - Scroll down and select **Web Server (IIS)**. If prompted to add features, click **Add Features**.
   - Click **Next**.

6. **Confirm and Install**:
   - Continue through the wizard, accepting default settings. 
   - Click **Install** and wait for the installation to complete.

---

### **Step 2: Obtain an SSL Certificate**

Before enabling HTTPS, you need an **SSL certificate**. You have two main options:

#### **Option 1: Purchase an SSL Certificate**
- Obtain an SSL certificate from a trusted **Certificate Authority (CA)** such as **DigiCert**, **GoDaddy**, or **Let’s Encrypt**.
- The CA will require you to generate a **Certificate Signing Request (CSR)**, which you can do from IIS.

#### **Option 2: Create a Self-Signed Certificate (for testing)**
For development or testing purposes, you can create a **self-signed certificate** directly in IIS.

---

### **Step 3: Generate a CSR (for purchased SSL certificates)**

If you’re using a **CA-issued certificate**, you first need to generate a **CSR** from IIS.

1. **Open IIS Manager**:
   - Click the **Start** button, type **IIS**, and select **Internet Information Services (IIS) Manager**.

2. **Server Certificates**:
   - In the **Connections** pane on the left, select the root node (your server name).
   - Double-click **Server Certificates** under the **IIS** section in the middle panel.

3. **Create Certificate Request**:
   - In the **Actions** pane on the right, click **Create Certificate Request**.
   - Fill in your company and domain information, including:
     - **Common Name**: This is the fully qualified domain name (e.g., `www.yourdomain.com`).
     - **Organization**: The name of your company.
     - **Organizational Unit**: Department within your organization (e.g., IT).
     - **City/Locality**, **State/Province**, **Country**.

4. **Choose Cryptographic Options**:
   - Select **Microsoft RSA SChannel Cryptographic Provider** and set **2048-bit** key length. Click **Next**.

5. **Save the CSR**:
   - Save the CSR file to a location on your server. You’ll submit this CSR to the **CA** when purchasing your certificate.

6. **Submit CSR to the CA**:
   - Submit the CSR to your chosen **Certificate Authority**.
   - After verification, the CA will issue an SSL certificate.

---

### **Step 4: Install the SSL Certificate**

Once the CA issues your SSL certificate, you can install it on your server.

1. **Open IIS Manager**:
   - Go back to **Server Certificates** in **IIS Manager**.

2. **Complete Certificate Request**:
   - In the **Actions** pane, click **Complete Certificate Request**.
   - Browse for the SSL certificate file provided by the **CA** (usually a `.cer` or `.crt` file).
   - Give the certificate a friendly name for easy identification and click **OK**.

---

### **Step 5: Bind the SSL Certificate to Your Website**

Now that your SSL certificate is installed, you need to bind it to your website to enable HTTPS.

1. **Open IIS Manager**:
   - In the **Connections** pane, expand your server’s name and expand **Sites**.
   - Select the site you want to configure for HTTPS.

2. **Edit Bindings**:
   - In the **Actions** pane on the right, click **Bindings**.
   - In the **Site Bindings** window, click **Add**.

3. **Add HTTPS Binding**:
   - In the **Add Site Binding** window:
     - Set **Type** to **https**.
     - Select the appropriate **IP address** (or leave as `All Unassigned`).
     - In the **SSL Certificate** dropdown, select the SSL certificate you installed.
   - Click **OK**.

4. **Remove HTTP Binding** (Optional):
   - If you want to force users to use HTTPS only, you can remove the HTTP binding by selecting it and clicking **Remove**.

---

### **Step 6: Force HTTPS Redirection (Optional)**

To ensure that all users accessing your site are automatically redirected from HTTP to HTTPS, you can set up redirection in IIS.

1. **Install URL Rewrite Module**:
   - Download and install the **URL Rewrite Module** for IIS from [Microsoft's website](https://www.iis.net/downloads/microsoft/url-rewrite).

2. **Create Redirect Rule**:
   - Open **IIS Manager** and select your website.
   - Double-click **URL Rewrite** in the middle panel.
   - In the **Actions** pane, click **Add Rules**.
   - Select **Blank Rule** and click **OK**.
   
3. **Configure the Rule**:
   - Name the rule something like **Redirect to HTTPS**.
   - In the **Match URL** section, set the **Requested URL** to **Matches the Pattern** and the pattern to `.*`.
   - In the **Conditions** section, click **Add**. Set **Condition Input** to `{HTTPS}` and **Check if input string** to **Matches the pattern**. In the pattern box, enter `^OFF$`.
   - In the **Action** section, select **Redirect** as the action type. In the **Redirect URL** box, enter `https://{HTTP_HOST}/{R:1}`.
   - Set **Redirect Type** to **Permanent (301)** and click **Apply**.

---

### **Step 7: Test the HTTPS Configuration**

1. **Verify HTTPS Binding**:
   - Open a web browser and navigate to your website using HTTPS (`https://yourdomain.com`).
   - You should see the secure lock icon in the browser, confirming the SSL certificate is installed and HTTPS is active.

2. **Check SSL Certificate**:
   - Click on the lock icon in the browser’s address bar to view details about the SSL certificate, including its validity and issuer.

---

### **Step 8: Enable HTTP/2 (Optional)**

Windows Server 2019/2022 supports **HTTP/2**, which can improve website performance over HTTPS. It is enabled by default when you bind a site to HTTPS. You can verify it by inspecting the browser’s network tools.

---

### **Conclusion**

You’ve now set up and configured **HTTPS** on your **Windows Server 2019/2022** using **IIS**. By binding an SSL certificate and redirecting traffic from HTTP to HTTPS, you can ensure secure communication between users and your web server.