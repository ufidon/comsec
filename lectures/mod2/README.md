# Cryptographic Tools
ch2

The biggest US intelligence leaks
---
- In 2023, Jack Teixeira revealed highly sensitive Pentagon documents about the Ukrainian war
- In 2013, Edward Snowden revealed the U.S. government’s actions in covertly surveilling millions of Americans
- In 2010, Julian Assange, on WikiLeaks, revealed the actions of U.S. and coalition forces in the Iraq war from 2004 to 2009


How to keep confidentiality
---
- set up policies and rules on accessing, storing and transmitting confidential information
- place restrictions on confidential information 
- encrypt confidential information with **cryptographic tools**



## Confidentiality with symmetric ciphers

💡 Demo: Caesar cipher 
---
- Demo encryption & decryption with [Caesar cipher](https://www.xarg.org/tools/caesar-cipher/)


Encryption & decryption
---
- provide confidentiality for transmitted and stored data
- *plaintext*: human readable message, text, image, audio or video
- *ciphertext*: human unreadable message scrambled with encryption algorithm
- *encryption algorithm*: performs various substitution and transformation on plaintext with an *encryption key*
- *decryption algorithm*: reverse the operations of encryption to get the original message with a *decryption key*



Symmetric cipher
---
- the encryption key = the decryption key
- also called single-key cipher
- Two requirements for secure use
  - a strong encryption algorithm
  - Sender and receiver must 
    - have obtained copies of the secret key in a secure fashion
    - must keep the key secure
- the key revealed is calls compromised


Attack symmetric cipher by **brute-force**
---
- Try all possible keys on some ciphertexts
- averagely half of all possible keys must be tried
  - best case? worst case? average case?


Attack symmetric cipher by **cryptanalysis**
---
- rely on:
  - Nature of the algorithm
  - Some knowledge of the general characteristics of the plaintext
    - such as [letter frequencies](https://en.wikipedia.org/wiki/Letter_frequency)
  - Some sample pairs of plaintext-ciphertext


🖊️ Practice: attack Caesar cipher by 
---
-  [brute-force attack (exhaustive key search)](https://www.dcode.fr/caesar-cipher)
-  cryptanalysis attack


Symmetric [block cipher](https://en.wikipedia.org/wiki/Block_cipher)
---
- operates on blocks of bits or bytes
- 3 popular symmetric block ciphers
  - AES: [Advanced Encryption Standard](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)
  - DES: [Data Encryption Standard](https://en.wikipedia.org/wiki/Data_Encryption_Standard)
  - Triple DES: [Triple DES](https://en.wikipedia.org/wiki/Triple_DES)


Comparison of the [3 popular symmetric block ciphers](https://en.wikipedia.org/wiki/Cipher_security_summary)
---
| block & key size in *bits*\cipher | AES | DES | 3DES |
| --- | --- | --- | --- |
| plaintext block size | 128 | 64 | 64 |
| ciphertext block size | 128 | 64 | 64 |
| key size | 128, 192, or 256 | 56 | 112 or 168 |


[Data Encryption Standard (DES)](https://en.wikipedia.org/wiki/Data_Encryption_Standard)
---
- widely used
- published on Federal Information Processing Standard (FIPS) PUB 46
- uses 64 bit plaintext block and 56 bit key to produce a 64 bit ciphertext block
- strength concerns
  - DES is studied thoroughly
  - 56-bit key length is vulnerable to brute-force attack


🖊️ Practice: Find the average time required for exhaustive key search
---
| key size in bits | cipher | number of alternative keys | time required at $10^{12}$ decryptions / $\mu s$ |
| --- | --- | --- | --- |
| 56 | DES | $2^{56}$ |  |
| 112 | 3DES | $2^{112}$ |  |
| 128 | AES | $2^{128}$ | $5.395 \times 10^{12}$ years |
| 168 | 3DES | $2^{168}$ |  |
| 192 | AES | $2^{192}$ |  |
| 256 | AES | $2^{256}$ |  |

Calculation steps
---
- average case tries half of the possible alternative keys: $2^{keyLength-1}$
- time needed = $2^{keyLength-1}/exhaustiveSpeed$
- do the unit conversion
- years needed = $2^{keyLength-1}/exhaustiveSpeed/10^6/60/60/24/365$


Triple DES: [Triple DES](https://en.wikipedia.org/wiki/Triple_DES)
---
- published in several standard documents
- widely used but start deprecation in 2017
  - complete deprecation by the end of 2023
- repeats DES algorithm three times using either two or three unique keys
  - two unique keys: $56×2=112$
  - three unique keys: $56×3=168$
- strong key length against brute-force attack, but
- data block size is still 64bits


AES: [Advanced Encryption Standard](https://en.wikipedia.org/wiki/Advanced_Encryption_Standard)
---
- a replacement for 3DES
- established by the U.S. National Institute of Standards and Technology (NIST) in 2001
  - U.S. Federal Information Processing Standards (FIPS) PUB 197
- stronger and more efficient than t3DES
- block size: 128bits; key sizes: 128, 192, 256bits


How to apply block cipher on data much big than one block?
---
- divide data into multiple blocks of input size: 64 or 128 bits
  - then apply cipher on each block in multiple ways
  - called [block cipher mode of operation](https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation)
- Electronic CodeBook (ECB) mode is the simplest approach
  - Each block of plaintext is encrypted using the same key independently
  - Regularities in the plaintext could occur in the cyphertext
    - can be overcomed with chained mode or counter mode


Popular [block cipher modes of operation](https://en.wikipedia.org/wiki/Block_cipher_mode_of_operation)
---
- Electronic codebook (ECB)
- Cipher block chaining (CBC), Propagating cipher block chaining (PCBC)
- Cipher feedback (CFB), Output feedback (OFB), Counter (CTR)


Block cipher vs. stream cipher
---
- [Block cipher](https://en.wikipedia.org/wiki/Block_cipher)
  - input is processed block-wise
  - an output block is produced for each input block
  - can reuse keys
- [Stream cipher](https://en.wikipedia.org/wiki/Stream_cipher)
  - input is processed element-wise
    - bit by bit, or byte by byte
  - an output element is produced for each input element
  - almost always faster and use far less code than block cipher
  - a pseudorandom stream (key stream) is generated with the input key
    - unpredictable without knowledge of the input key
    - bit-wise xor-ed with plaintext to get ciphertext


💡 Demo
---
- Encrypt & decrypt a file in Parrot Linux
  ```bash
  # see the list under the 'Cipher commands' heading
  openssl -h

  # or get a long list, one cipher per line
  openssl list-cipher-commands

  # encrypt file.txt to file.enc using 256-bit AES in CBC mode
  openssl enc -aes-256-cbc -salt -in file.txt -out file.enc

  # the same, only the output is base64 encoded for, e.g., e-mail
  openssl enc -aes-256-cbc -a -salt -in file.txt -out file.enc

  # decrypt binary file.enc
  openssl enc -d -aes-256-cbc -in file.enc

  # decrypt base64-encoded version
  openssl enc -d -aes-256-cbc -a -in file.enc

  # provide password on command line
  openssl enc -aes-256-cbc -salt -in file.txt \
  -out file.enc -pass pass:mySillyPassword

  # provide password in a file
  openssl enc -aes-256-cbc -salt -in file.txt \
  -out file.enc -pass file:/path/to/secret/password.txt
  ```

## Message authentication and hash function

Message authentication
---
- assure data integrity against data tampering
- verify received message is authentic
  - contents have not been altered
  - from authentic source
  - timely and in correct sequence
- can use conventional encryption
  - sender and receiver need to share the key
  - sender and receiver must trust mutually


Message Authentication Without Confidentiality
---
- Message encryption only does not provide a secure form of authentication
- authentication and confidentiality combined in a single algorithm is desired
  - message authentication is provided as a separate function from message encryption
- applicable situations
  - message in plaintext is desired, such as broadcast
  - it is a burden to decrypt incoming message
  - authentication in plaintext is desired


Message Authentication Using a Message Authentication Code (MAC)
---
- for the sender
  - generate $MAC$ from the message $M$ and a key $K$ with a *MAC algorithm*
  - send out $MAC + M$
- for the receiver
  - recalculate the $MAC'$ from the received message $M'$ and the key $K$
  - compare the recalculated $MAC'$ with the received $MAC''$
    - if $MAC'=MAC''$, message is authentic
- the MAC algorithm is usually a *cryptographic hash function + cipher*


Message Authentication using a [hash function](https://en.wikipedia.org/wiki/Hash_function) and a cipher
---
- for the sender
  - get the hash tag $H$ of the message $M: Hash(M)$
  - encrypt $H$ with a key $K$ to get $MAC$
    - a shared key for symmetric encryption
    - a public key for asymmetric encryption
  - send out: $M + MAC$
- for the receiver
  - recalculate the hash tag $H'$ of the received message $M': Hash(M')$
  - decrypt the received $MAC'$ to get $H''$ with key $K$
  - compare $H'$ with $H''$
    - if $H'=H''$, message is authentic


Message Authentication using a hash function and a shared key
---
- for the sender
  - get the hash tag $H$ of the combination of the message $M$ and the shared key $K: Hash(K+M+K)$
  - send out: $M + H$
- for the receiver
  - recalculate the hash tag $H'$ of the received message $M': Hash(K+M'+K)$
  - compare $H'$ with the received $H''$
    - if $H'=H''$, message is authentic


[Hash function properties](https://en.wikipedia.org/wiki/Cryptographic_hash_function)
---
- universality:
  - accept input of any sizes
  - generate a fixed-length output
- efficiency: 
  - $∀x: H(x)$ can be computed efficiently
- One-way or pre-image resistant:
  - Computationally infeasible to find $x$ such that $H(x)=h$
- Second pre-image resistance:
  - Computationally infeasible to find a pair $x≠y$ such that $H(x)=H(y)$
- Collision resistant or strong collision resistance:
  - Computationally infeasible to find any pair $x≠y$ such that $H(x)=H(y)$


💡 Demo
---
- Explore [popular hash functions](https://en.wikipedia.org/wiki/List_of_hash_functions)
  - MD5: message digest
  - SHA: secure hash algorithm
- Generate hash code for any file
- On Linux,
  ```bash
  # get MD5 of a string
  echo "hello" | md5sum

  # get SHA1 of typed input, end by CTRL+D
  sha1sum

  # get SHA256 of program sha256sum
  sha256sum `which sha256sum`
  ```
- On Windows, use PowerShell commandlet [Get-FileHash](https://learn.microsoft.com/en-us/powershell/module/microsoft.powershell.utility/get-filehash)
  ```powershell
  Get-FileHash myFile
  ```


Security of [Hash Functions](https://hashcat.net/wiki/doku.php?id=example_hashes)
---
- two ways attacking secure hash functions
  - cryptanalysis attack
    - exploit logical weakness in the algorithm
  - brute-force attack
    - strength of hash function depends solely on the length of the hash code
- applications
  - save the hashes of passwords instead in plaintext
  - intrusion detection by saving the hash of each file securely


## Authentication and non-repudiation with asymmetric ciphers

❓ What is the problem with the previous message authentication
---
- Both authentication and non-repudiation can be broken
  - if the sender and receiver lack trust
  - why?
- can be solved with asymmetric ciphers


Asymmetric cipher
---
- also called public-key cipher
- based on mathematical functions
- asymmetric
  - use two separate but related keys
  - a public key and a private key
  - private key is never released
  - public key is made public for others to use
    - it is a challenge to distribute public keys


Public-key cryptography
---
- public key is used for encryption
  - the generated ciphertext can only be decrypted with the related private key
- the private key can also be used to encrypt message
  - but this process is called signing the message instead of encryption
  - the ciphertext can only be decrypted with the related public key
    - but it is called signature verification instead of decryption


Popular public-key cryptosystems
---
- [RSA (Rivest–Shamir–Adleman)](https://en.wikipedia.org/wiki/RSA_(cryptosystem))
  - Its security of RSA relies on the "factoring problem":
    - the practical difficulty of factoring the product of two large prime numbers 
  - widely accepted and implemented
  - 
- [Diffie-Hellman](https://en.wikipedia.org/wiki/Diffie%E2%80%93Hellman_key_exchange)
  -  a mathematical method of securely exchanging cryptographic keys over a public channel
- [DSS (Digital Signature Standard)](https://en.wikipedia.org/wiki/Digital_Signature_Standard)
  - a FIPS specifying a suite of algorithms generating signatures
- [ECC (Elliptic-curve cryptography)](https://en.wikipedia.org/wiki/Elliptic-curve_cryptography)
  - based on the algebraic structure of elliptic curves over finite fields
  - smaller keys compared to non-EC cryptography to provide equivalent security


Applications for public-key cryptosystems
---
| Algorithm | Digital signature | Public key exchange | Encryption |
| --- | --- | --- | --- |
| RSA | Y | Y | Y |
| Diffie-Hellman | N | Y | N |
| DSS | Y | N | N |
| ECC | Y | Y | Y |


Requirements for Public-Key Cryptosystems
---
- Efficiency
  - Efficient to create key pairs
  - Efficient for sender knowing public key to encrypt messages
  - Efficient for receiver knowing private key to decrypt ciphertext
- Security
  - Computationally infeasible for opponent to derive private key from public key
  - Computationally infeasible for opponent to recover original message


## Digital signatures and key management

Digital signatures
---
- defined in NIST FIPS PUB 186-4
  - the result of a cryptographic transformation of data
  - with properly implementation, provides a mechanism for verifying 
    - origin authentication, 
    - data integrity 
    - signatory non-repudiation
- FIPS 186-4 specified three digital signature algorithms:
  - Digital Signature Algorithm (DSA)
  - RSA Digital Signature Algorithm
  - Elliptic Curve Digital Signature Algorithm (ECDSA)


Sign a message and verify the signature
---
- Sign a message $M$
  - the signature $S$ on $M$ is the ciphertext of a hash code $h$ of $M$
  - encrypted with signer's private key $K_{pri}: S=E_{K_{pri}}(h)$
  - $M$ and $S$ are sent out
- Verify the signature $S'$ for $M'$ by the receiver
  - generate the hash code $h'$ of $M'$
  - decrypt $S'$ with signer's public key $P_{pub}$ to get $h''=D_{K_{pub}}(S')$
  - if $h' = h''$, the signature is valid


Public key certificate
---
- public key is vulnerable to spoofing (impersonation)
  - can be overcomed with public key certificate
- a public key certificate binds 
  - the unsigned certificate which contains
    - the public key, 
    - its owner's identity  
    - the certificate authority (CA) 's information
  - and the CA's signature for the unsigned certificate
    - can be verified with CA's public key
  - most saved in the format specified in the the [X.509 standard](https://en.wikipedia.org/wiki/X.509)


💡 Demo
---
- Show a sample public key certificate


Symmetric key exchange using public-key encryption
---
- symmetric cipher is efficient for large data
- public-key encryption is suitable for small chunk of data
  - e.g. the *secret key* used in symmetric cipher
- two popular ways of symmetric key exchange
  - Diffie–Hellman key exchange
    -  no authentication of the two communicating partners
   - public-key encryption with public key certificate
     - referred to as a *digital envelope*


Digital envelope creation and opening
---
- for the sender:
  - generate a random symmetric key $K_s$
  - encrypt a message $M$ with $K_s$ using a symmetric cipher to get the encrypted message $C_M$
  - encrypt $K_s$ with receiver's public key $K_{pub}$ to get $C_{K_s}$
  - the digital envelope = $C_M + C_{K_s}$ is sent to the receiver
- for the receiver:
  - decrypt $C_{K_s}$ with her private key $K_{pri}$ to get $K_s$
  - decrypt $C_M$ with $K_s$ to get $M$



## Random and pseudorandom numbers


Applications of random numbers
---
Generation of 
- key for asymmetric cipher
- stream key for stream cipher
- temporary session key
- secret key in digital envelope
- handshaking to prevent replay attacks


Random number requirements
---
- Randomness criteria
  - uniform distribution
    - occurrence frequency of each number should be approximately the same
  - independence
    - no one value in the sequence can be inferred from the others
- Unpredictability
  - statistically independence between numbers in the sequence
  - future elements of the sequence are not predictable


Random vs. Pseudorandom
---
- True random number generator (TRNG)
  - use nondeterministic sources
  - Most operate by measuring unpredictable natural processes
    - e.g. capacitor leakage, gas discharge, radiation
    - increasingly provided on modern processors
- Pseudorandom number generator (PRNG or DRNG)
  - satisfy statistical randomness tests
  - likely to be predictable
- random numbers generated from cryptographic algorithms are not statistically random
  - due to the deterministic algorithms


Practical application: encryption of stored data
---
- Common to encrypt transmitted data
  - HTTPS
- Much less common for stored data
  - little protection beyond domain authentication and access controls
  - data are archived for indefinite periods
  - data are recoverable even though erased until disk sectors are reused
- Approaches to encrypt stored data
  - available encryption packages
    -  Pretty Good Privacy (PGP), VeraCrypt
   - back-end appliance
     -  hardware device that sits between servers and storage systems and encrypts/decrypts all passing data
  - library-based tape encryption
    - a cryptographic co-processor embedded in the tape drive and tape library hardware
  - background laptop and PC data encryption
    - Windows BitLocker, MacOS FileVault


💡 Demo:
---
- [Gpg4win](https://www.gpg4win.org/)


# References
- [The 11 Largest National Security Leaks in American History](https://www.saturdayeveningpost.com/2022/09/the-11-largest-national-security-leaks-in-american-history/)
- [OpenSSL Command-Line HOWTO](https://www.madboa.com/geek/openssl/)
- [OpenSSL Cookbook](https://www.feistyduck.com/library/openssl-cookbook/online/)