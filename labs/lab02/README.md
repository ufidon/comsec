# Application of Cryptography

## 1. Description
The combination of modern symmetric cryptography, asymmetric cryptography and hash functions form the basis of secure Internet, secure Intranet, secure system and service as well as Blockchain, etc. by assuring their confidentiality, integrity, authenticity and nonrepudiation, which leads to availability.

This lab is supposed to be done on the Windows server virtual machine.

_Note: you may do this lab on your host system if it's Windows_

**You need a friend or classmate to complete this lab**

## 2. Tasks: Secure e-mails and files using Gpg4win
In this task, you will use Gpg4win to manage your private/public keys, certificates and your friends' certificates. By which, you can secure e-mails, files and folders.

## Task 1 (30%) Generate, import and manage OpenPGP certificates
- (5%) Download and install [Gpg4win](https://www.gpg4win.org/)
- (5%) Generate OpenPGP certificate using your school email
- (5%) Publish your OpenPGP certificate on a public keyserver
- (5%) Backup your private key
- (5%) Export your public key
- (5%) Find the OpenPGP certificate of one of your fiends or classmates, then import it

## Task 2 (20%) Data/program integrity assurance 
Take a photo with your smartphone and send it to your computer.
- (5%) Create a checksum for this photo
- (5%) Verify the checksum is valid
- (5%) modify the photo a little bit then verify the old checksum becomes invalid
- (5%) What is the checksum algorithms? What is the length of this checksum in bits? in bytes? in hex digits? Open the checksum file with Notepad.


## Task 3 Privacy assurance (50%)
GnuPG can be used to _protect data_. Create a folder 'secrets' containing at least three image files, use your OpenPGP key pair and your friend's public key (download from the GPG keyserver on which your friend published her/his public key) to complete the subtasks:
	
* Subtask 3.1 (20%): Send secrets to a friend
  * Sign this folder with your private key and encrypt it with your friends public key. 
  * Then send it to your friend through email, USB drive or sharing drive.
	
* Subtask 3.2 (20%): Read secrets from a friend
  * Receive the signed and encrypted file from your friend through email, USB drive or sharing drive. 
  * Decrypt this file with your private key and check it contents, and check your friend's signature is valid.

* Subtask 3.3 (10%): Keep secrets for yourself
  * Encrypt the folder only with your OpenPGP certificate
  * Decrypt the encrypted file and check its contents

## 3. Extra credits (10%): Secure emails with Mailvelope
Watch this video [Encrypt Your Gmail/Yahoo/Outlook/iCloud and Other Webmail](https://youtu.be/\-Hz40\_P6bVE), using [Mailvelope](https://www.mailvelope.com/) to 
* (2%) import your private OpenPGP key exported from Kleopatra
* (2%) import your partner's OpenPGP certificate exported from Kleopatra or a key server
* (3%) send an encrypted and signed email to your partner, 
* (3%) decrypt and verify his/her encrypted and signed email that sent to you 
 

## 4. Report
Write a report about the process you complete the tasks above, key screen snapshots are needed as evidences.

# References
* [Pretty Good Privacy.](https://en.wikipedia.org/wiki/Pretty\_Good\_Privacy)
  * [GNU Privacy Guard.](https://en.wikipedia.org/wiki/GNU\_Privacy\_Guard)
  * [Gpg4win.](https://en.wikipedia.org/wiki/Gpg4win)
  * [Gpg4win.](https://www.gpg4win.org)
  * [The Gpg4win Compendium.](https://files.gpg4win.org/doc/gpg4win-compendium-en.pdf)
* [Mailvelope Wiki.](https://en.wikipedia.org/wiki/Mailvelope)
  * [Mailvelope official website.](https://www.mailvelope.com)
  * [Mailvelope source code.](https://github.com/mailvelope/)
  * [Encrypt Your Gmail/Yahoo/Outlook/iCloud and Other Webmail.](https://www.youtube.com/watch?v=\-Hz40\_P6bVE)