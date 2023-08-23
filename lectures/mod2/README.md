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
-  brute-force attack (exhaustive key search)
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


## Message authentication and hash function


## Authentication and non-repudiation with asymmetric ciphers


## Digital signatures and key management

## Random and pseudorandom numbers


# References
- [The 11 Largest National Security Leaks in American History](https://www.saturdayeveningpost.com/2022/09/the-11-largest-national-security-leaks-in-american-history/)