# User authentication & authorization

### Lab Objective:
In this lab, you will set up a domain using Active Directory on Windows Server 2019, create users with specific roles and folder permissions for a university structure, and enforce a complex password policy to enhance security.



### Scenario:
Imagine a university with three departments: **Computer Science**, **Mathematics**, and **Engineering**. Each department has one chairperson, three faculty members, and thirty-three students. The entire university has one provost.

1. **Install and configure Active Directory** on Windows Server 2019.
2. **Enforce a complex password policy** to ensure strong passwords for all users in the domain.
3. **Create users** and organize them by department.
4. **Set up folder permissions** for each department according to roles:
    - **Department Announcements Folder**: Faculty and Students have read-only access, and the Chairperson has read/write access.
    - **Department Documents Folder**: Accessible by the Chairperson and Faculty only, with no Student access.
    - **Department Shared Folder**: Accessible by all members of the department with full control.
5. **Create a university-wide folder** accessible to all, with read-only access for all users except the Provost, who has read/write access.



### Prerequisites:
- Windows Server 2019 installed.
- Administrative access to the server.
- Basic understanding of Active Directory.


**Active Directory (AD)** is a directory service developed by Microsoft for Windows domain networks. It manages user data, security, and distributed resources and enables administrators to manage permissions and access to network resources.

Active Directory works with several components:
- **Domain Services (AD DS)**: Central to managing users, groups, and computers.
- **Group Policy**: Manages users and computers' behavior centrally.
- **Organizational Units (OUs)**: Logical units that help organize and manage resources.


⚠️ *All Powershell Methods are optional, just for reference.*

## **Task 1: Installing and Configuring Active Directory**

1. **Install Active Directory Domain Services (AD DS)**

   - Open **Server Manager**, click **Manage**, and select **Add Roles and Features**.
   - Choose **Role-based or feature-based installation**.
   - Select the server from the server pool and click **Next**.
   - Select **Active Directory Domain Services** and click **Next**.
   - Complete the installation and click **Promote this server to a domain controller** in the post-installation window.
   - Choose **Add a new forest**, enter a domain name (e.g., `university.local`), and proceed through the configuration.
   - When prompted, restart the server.
   - 💻 the new forest `university.local`

2. **Configure Active Directory Users and Computers**
   - After the server restarts, open **Active Directory Users and Computers** from the start menu.
   - Create an organizational unit (OU) for the university:
     - Right-click the domain (e.g., `university.local`) > **New** > **Organizational Unit**.
     - Name it **University** and create sub-OUs for **Computer Science**, **Mathematics**, and **Engineering**.
   - 💻 all OUs



## **Task 2: Enable Complex Password Policy for All Users**

**Objective:** Configure Windows Server 2019 to enforce complex password policies for all users in the domain.


1. **Open Group Policy Management**
   - Log in to the **Domain Controller** where Active Directory and Group Policy are managed.
   - Press `Windows Key + R`, type `gpmc.msc`, and hit Enter to open **Group Policy Management**.

2. **Edit the Default Domain Policy**
   - In the **Group Policy Management** window:
     - Navigate to **Forest** > **Domains** > your domain name (e.g., `university.local`).
     - Right-click on **Default Domain Policy**, then click **Edit**.
   
3. **Navigate to Password Policy Settings**
   - In the **Group Policy Management Editor**, go to:
     - **Computer Configuration** > **Policies** > **Windows Settings** > **Security Settings** > **Account Policies** > **Password Policy**.

4. **Enable Complex Password Policy**
   - Find the setting **"Password must meet complexity requirements"** in the **Password Policy** folder.
   - Double-click on it and set the option to **Enabled**.
   
5. **Configure Additional Password Policy Settings (Optional)**
   You can also configure additional policies like:
   - **Minimum password length**: Set a minimum number of characters (e.g., 8 characters).
   - **Enforce password history**: This prevents users from reusing old passwords.
   - **Maximum password age**: Defines how long a password can be used before it needs to be changed.
   - **Minimum password age**: Specifies the minimum amount of time a password must be used before it can be changed.

6. **Apply the Policy**
   - Click **OK** to save the changes, then close the **Group Policy Management Editor**.

- 💻 complex password policy

7. **Force Group Policy Update**
   After configuring the policy, force a Group Policy update to apply the settings immediately:
   - Open **PowerShell** or **Command Prompt** and run the following command:
     ```bash
     gpupdate /force
     ```
   - 💻 policy update confirmation

### *PowerShell Method for Enabling Complex Password Policy:*

1. Open **PowerShell** as Administrator on the Domain Controller.
2. Run the following commands:

```powershell
# Enable complex password requirements
secedit /export /cfg C:\Windows\Temp\secpol.cfg

# Modify the password complexity requirements in the file
(Get-Content C:\Windows\Temp\secpol.cfg) -replace "PasswordComplexity = 0", "PasswordComplexity = 1" | Set-Content C:\Windows\Temp\secpol.cfg

# Import the modified security policy
secedit /import /cfg C:\Windows\Temp\secpol.cfg

# Apply the new security policy
gpupdate /force
```

- This PowerShell method edits the security settings and applies the changes in bulk, ensuring that complex passwords are enforced.

### *Password Complexity Requirements:*

Once enabled, Windows will enforce the following requirements:
- Passwords must contain at least three of the following:
  - Uppercase letters (A-Z)
  - Lowercase letters (a-z)
  - Numbers (0-9)
  - Non-alphabetic characters (e.g., `!`, `@`, `#`, `$`, `%`)
- Passwords cannot contain parts of the user's name or username.
- The minimum password length is typically set to 8 characters (though you can adjust this).

By the end of this task, all users in your domain will be required to create more secure, complex passwords, enhancing the overall security of your Windows Server 2019 environment.



## **Task 3: Creating University Users with Different Privileges**

1. **Create Users**

   - Right-click the OU for each department (e.g., **Computer Science**) > **New** > **User**.
   - Create users for each department:
     - One **Chair** (e.g., `CS_Chair`).
     - Three **Faculty Members** (e.g., `CS_Faculty1`, `CS_Faculty2`, `CS_Faculty3`).
     - Thirty-three **Students** (e.g., `CS_Student1` to `CS_Student33`).
       - ⚠️ Create at least 3 students for each department
   - Repeat for the **Mathematics** and **Engineering** OUs.
   - 💻 all users for each department

2. **Create a Provost User**
   - In the **University** OU, create a user for the **Provost** (e.g., `Provost`).
   - 💻 user `Provost`

3. **Create Security Groups**
   - For each department, create security groups for **Chairs**, **Faculty**, and **Students**.
     - Right-click the department OU > **New** > **Group**.
     - Name the groups (e.g., `CS_Chair`, `CS_Faculty`, `CS_Students`).
   - Add users to the respective groups (e.g., add the `CS_Chair` to the `CS_Chair` group, faculty to the `CS_Faculty` group, etc.).
   - 💻 all users for all groups for each department


### *Using PowerShell to Create Users*
Alternatively, use PowerShell to automate user creation as follows:

```powershell
# Create Provost
New-ADUser -Name "Provost" -GivenName "University" -Surname "Provost" -UserPrincipalName "provost@university.local" -Path "OU=Provost,OU=University,DC=university,DC=local" -AccountPassword (ConvertTo-SecureString "Password123!" -AsPlainText -Force) -Enabled $true

# Department Chairs
$departments = @("ComputerScience", "Mathematics", "Engineering")
foreach ($dept in $departments) {
    New-ADUser -Name "${dept}_Chair" -UserPrincipalName "${dept}_chair@university.local" -Path "OU=$dept,OU=University,DC=university,DC=local" -AccountPassword (ConvertTo-SecureString "Password123!" -AsPlainText -Force) -Enabled $true
}

# Faculty Members
foreach ($dept in $departments) {
    for ($i = 1; $i -le 3; $i++) {
        New-ADUser -Name "${dept}_Faculty$i" -UserPrincipalName "${dept}_faculty$i@university.local" -Path "OU=$dept,OU=University,DC=university,DC=local" -AccountPassword (ConvertTo-SecureString "Password123!" -AsPlainText -Force) -Enabled $true
    }
}

# Students
foreach ($dept in $departments) {
    for ($i = 1; $i -le 33; $i++) {
        New-ADUser -Name "${dept}_Student$i" -UserPrincipalName "${dept}_student$i@university.local" -Path "OU=$dept,OU=University,DC=university,DC=local" -AccountPassword (ConvertTo-SecureString "Password123!" -AsPlainText -Force) -Enabled $true
    }
}
```


## **Task 4: Setting Up Folder Permissions**

1. **Create Departmental Folders**

   - Create a base directory for the university (e.g., `C:\University`).
   - For each department, create three folders:
     - **Announcements**: Faculty and Students have **read-only** access, Chair has **read/write** access.
     - **Documents**: Only Faculty and Chair have access.
     - **Shared**: Full access for all department members.

   - Example structure:
   ```
   - **C:\University\ComputerScience**
     - Announcements
     - DepartmentDocuments
     - SharedResources
   - **C:\University\Mathematics**
     - Announcements
     - DepartmentDocuments
     - SharedResources
   - **C:\University\Engineering**
     - Announcements
     - DepartmentDocuments
     - SharedResources
   ```

   - 💻 all folders with access rights for each department


2. **Assign Folder Permissions Using PowerShell**

   Use the following PowerShell script to automate the permission assignments for each department:

   ```powershell
   $departments = @("ComputerScience", "Mathematics", "Engineering")
   $basePath = "C:\University"

   foreach ($dept in $departments) {
       $deptPath = "$basePath\$dept"

       # Create the department folders
       New-Item -ItemType Directory -Path "$deptPath\Announcements"
       New-Item -ItemType Directory -Path "$deptPath\Documents"
       New-Item -ItemType Directory -Path "$deptPath\Shared"

       # Announcements folder: Read-only for Faculty and Students, Read/Write for Chair
       $announcementPath = "$deptPath\Announcements"
       icacls $announcementPath /grant "${dept}_Chair:(OI)(CI)M"
       icacls $announcementPath /grant "${dept}_Faculty1:(OI)(CI)R" "${dept}_Faculty2:(OI)(CI)R" "${dept}_Faculty3:(OI)(CI)R"
       for ($i = 1; $i -le 33; $i++) {
           icacls $announcementPath /grant "${dept}_Student$i:(OI)(CI)R"
       }

       # Documents folder: Only Faculty and Chair have access
       $documentsPath = "$deptPath\Documents"
       icacls $documentsPath /grant "${dept}_Chair:(OI)(CI)M"
       icacls $documentsPath /grant "${dept}_Faculty1:(OI)(CI)M" "${dept}_Faculty2:(OI)(CI)M" "${dept}_Faculty3:(OI)(CI)M"
       icacls $documentsPath /deny "Students:(OI)(CI)R"

       # Shared folder: Full access for all department members
       $sharedPath = "$deptPath\Shared"
       icacls $sharedPath /grant "${dept}_Chair:(OI)(CI)F" "${dept}_Faculty1:(OI)(CI)F" "${dept}_Faculty2:(OI)(CI)F" "${dept}_Faculty3:(OI)(CI)F"
       for ($i = 1; $i -le 33; $i++) {
           icacls $sharedPath /grant "${dept}_Student$i:(OI)(CI)F"
       }
   }
   ```

3. **Create the University-Level Folder**

   - Create a folder at `C:\University\ProvostAnnouncements`.
   - Set the Provost as the owner with **read/write** access, and all others in the domain as **read-only**:
   - 💻 access rights for the created folder

   ```powershell
   $provostPath = "C:\University\ProvostAnnouncements"
   New-Item -ItemType Directory -Path $provostPath
   icacls $provostPath /grant "Provost:(OI)(CI)M"
   icacls $provostPath /grant "Domain Users:(OI)(CI)R"
   ```

### *NTFS permissions*
In NTFS permissions, access rights are often described using permission flags like `(OI)`, `(CI)`, and specific access types such as `M` (Modify), `F` (Full Control), etc. Here’s a breakdown of what these flags and access types mean:

#### Common NTFS Permission Flags

1. **(OI) - Object Inherit**
   - This flag means that the permissions will apply to files (objects) within the directory (folder). For example, if this flag is set on a folder, any files created in that folder will inherit the same permissions.

2. **(CI) - Container Inherit**
   - This flag means that the permissions will apply to subdirectories (containers) within the directory. Subfolders created under this folder will inherit the same permissions.

3. **(IO) - Inherit Only**
   - This flag specifies that the permission applies only to child objects (files or folders) but not the folder itself. This is used to propagate permissions down the hierarchy without affecting the folder that contains the permissions.

#### Access Rights Codes

1. **F - Full Control**
   - The user has full access to the folder or file, including the ability to read, write, modify, delete, and change permissions or ownership.
   
2. **M - Modify**
   - This gives the user the ability to read, write, and delete files and subfolders. It also includes the ability to modify existing files, but does not include permissions like changing ownership or modifying access rights.
   
3. **RX - Read and Execute**
   - This allows the user to read and run executable files in the folder but does not allow modification.
   
4. **R - Read**
   - The user can view the contents of the folder or file, but cannot make any changes (write, delete, or modify).

5. **W - Write**
   - This allows the user to create new files and subfolders, as well as modify existing files. However, it doesn’t allow deletion of files or changing permissions.

#### Examples:

- **(OI)(CI)F**
   - **(OI)** and **(CI)** mean the permissions apply to both files and subfolders, and **F** grants **Full Control**, allowing the user to perform all actions (read, write, modify, delete, etc.) in both the folder and all its subcontents.
  
- **(OI)(CI)M**
   - **(OI)** and **(CI)** mean the permissions apply to both files and subfolders, and **M** grants **Modify** rights, which allow reading, writing, and deleting files, but do not include changing ownership or permissions.

- **(OI)(CI)R**
   - **(OI)** and **(CI)** mean the permissions apply to both files and subfolders, and **R** grants **Read** access only.

#### Additional Flags:

- **N**: No Access (explicitly denies access).
- **D**: Delete access, allowing the user to delete files or folders.
- **X**: Execute access, allowing the user to execute files.

#### Full Example:

The permission string `(OI)(CI)M` for the **Chair** on the Announcements folder allows the following:
- **Object Inherit (OI)**: Permissions are applied to files inside the folder.
- **Container Inherit (CI)**: Permissions are applied to subfolders.
- **Modify (M)**: The user can read, write, modify, and delete files and subfolders, but cannot change ownership or modify permissions.

For **Faculty and Students**, the permission `(OI)(CI)R` means:
- **Object Inherit (OI)**: Permissions are applied to files inside the folder.
- **Container Inherit (CI)**: Permissions are applied to subfolders.
- **Read (R)**: They can only read the contents but not modify, delete, or execute anything.

These permission flags and access rights allow granular control over who can do what within a directory or file structure in NTFS.


## **Task 5: Verifying all settings**
- 💻 Show that simple passwords won't be accepted
- 💻 Login as `Provost`, `CS_Chair`, `CS_Faculty1`, and `CS_Student1` individually to test all folder permissions.


### Lab Conclusion:
By the end of this lab, you will have:
- Set up an Active Directory domain and created users for a university structure.
- Enforced a complex password policy to improve domain security.
- Configured folder permissions based on user roles and department.

This exercise builds a solid foundation in managing Windows Server 2019, Active Directory, folder permissions, and security policies.