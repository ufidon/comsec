# Application of Cryptography

## 1. Description
Modern symmetric cryptography, asymmetric cryptography, and hash functions form the foundation of secure communication on the Internet, Intranet, secure systems and services, as well as Blockchain technology. These technologies ensure confidentiality, integrity, authenticity, and non-repudiation, all of which contribute to overall availability.

This lab is designed to be completed on a Windows server virtual machine.

_Note: You may also complete this lab on your host system if it runs Windows._

- **Collaboration Requirement**: You will need a friend or classmate to complete this lab.
- 💻 : screenshot


## 2. Tasks: Secure Emails and Files Using Gpg4win
In this lab, you will use Gpg4win to manage your private/public keys, certificates, and your friend's certificates, enabling you to secure emails, files, and folders.

### Task 1: Generate, Import, and Manage OpenPGP Certificates
1. Download and install [Gpg4win](https://www.gpg4win.org/).
   1. 💻 - a running Gpg4win
2. Generate an OpenPGP certificate using your school email.
   - **Ensure you protect your private key with a passphrase.**
   - 💻 - all components of your certificate
3. Publish your OpenPGP certificate on a public keyserver.
   - If you encounter issues publishing your certificate, manually upload your exported OpenPGP certificate to a keyserver like [keys.openpgp.org](https://keys.openpgp.org/) and follow the web prompts to verify your email.
   - 💻 - search result of your certificate
4. Backup your private (secret) key.
   1. 💻 - all components of your private key
5. Export your public key from your OpenPGP certificate.
   1. 💻 - all components of your public key
6. Search the OpenPGP certificate of a friend or classmate, download then import it.
   1. 💻 - all components of the certificate

👉 Show all public/private key components

```bash
# open a powershell, command windows, or terminal from the folder contains the key file
gpg --list-packets key_file
```

### Task 2: Data/Program Integrity Assurance
1. Take a photo with your smartphone and transfer it to your computer.
   1. 💻 - the photo
2. Create a checksum for this photo.
   1. 💻 - the checksum
3. Verify that the checksum is valid.
   1. 💻 - verification result
4. Modify the photo slightly and confirm that the old checksum is now invalid.
   1. 💻 - modified photo
   2. 💻 - verification result
5. Answer the following questions:
   - What checksum algorithm did you use?
   - What is the length of this checksum in bits, bytes, and hexadecimal digits?
   - Open the checksum file with Notepad and review its content.


### Task 3: Privacy Assurance
GnuPG can be used to protect data. Create a folder named "secrets" containing at least three image files, and use your OpenPGP key pair along with your friend's public key to complete the following subtasks:

**Subtask 3.1: Send Secrets to a Friend**
- Sign the "secrets" folder with your private key and encrypt it with your friend's public key.
- Send the signed and encrypted folder to your friend via email, USB drive, or a shared drive.
- 💻 transferred result

**Subtask 3.2: Read Secrets from a Friend**
- Receive the signed and encrypted file from your friend via email, USB drive, or a shared drive.
- Decrypt the file using your private key, verify its contents, and confirm your friend's signature is valid.
- 💻 signature verification
- 💻 decrypted content

**Subtask 3.3: Keep Secrets for Yourself**
- Encrypt the "secrets" folder using only your OpenPGP certificate.
  - 💻 encrypted result
- Decrypt the encrypted folder and verify its contents.
  - 💻 signature verification
  - 💻 decrypted content

### Task 4: Secure Emails with Mailvelope
Watch the video [Encrypt Your Gmail/Yahoo/Outlook/iCloud and Other Webmail](https://youtu.be/-Hz40_P6bVE) and use [Mailvelope](https://www.mailvelope.com/) to:
1. Import your private OpenPGP key exported from Kleopatra.
   1. 💻 imported result
2. Import your partner's OpenPGP certificate exported from Kleopatra or a key server.
   1. 💻  imported result
3. Send an encrypted and signed email to your partner.
   1. 💻  the encrypted and signed email
4. Decrypt and verify your partner's encrypted and signed email.
  - 💻 signature verification
  - 💻 decrypted content

## 3. Report
Write a report detailing the steps you took to complete the tasks above. Include key screenshots as evidence.

## References
- [Pretty Good Privacy](https://en.wikipedia.org/wiki/Pretty_Good_Privacy)
- [GNU Privacy Guard](https://en.wikipedia.org/wiki/GNU_Privacy_Guard)
- [Gpg4win](https://en.wikipedia.org/wiki/Gpg4win)
- [Gpg4win Official Site](https://www.gpg4win.org)
- [The Gpg4win Compendium](https://files.gpg4win.org/doc/gpg4win-compendium-en.pdf)
- [Mailvelope Wiki](https://en.wikipedia.org/wiki/Mailvelope)
- [Mailvelope Official Site](https://www.mailvelope.com)
- [Mailvelope Source Code](https://github.com/mailvelope/)
- [Encrypt Your Gmail/Yahoo/Outlook/iCloud and Other Webmail](https://www.youtube.com/watch?v=-Hz40_P6bVE)