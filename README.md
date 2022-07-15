# HackChrome

[![Build Status](https://travis-ci.com/godoes/HackChrome.svg?branch=master)](https://travis-ci.com/godoes/HackChrome)

**English** | [中文](https://github.com/godoes/HackChrome/blob/master/README_zh.md)

Get the User:Password from Chrome(include version < 80 and version > 80)

## Chrome version Affect

All version

## Platform

Windows

## Usage

- Install

  ```shell
  set CGO_ENABLED=1 & go install github.com/godoes/HackChrome@latest
  ```

- Open cmd or powershell

  ```shell
  HackChrome -h

  # output:
  Usage of HackChrome:
    -edge
          true or false, default is false and chrome is used.
  ```

- Run

  ```shell
  # Google Chrome:
  HackChrome > netpass.txt

  # Microsoft Edge:
  HackChrome -edge > netpass.txt
  ```

## Demo

![demo](image/result.png)

## Theory

- version < 80

User:Password pairs were stored in the file named "Login Data".

Password was encrypted, But we can use "CryptUnprotectData" Function in "Crypt32.dll" to decrypt them.

Finally, We get the plaintext of the User:Password pairs stored in Chrome

- version > 80

Based on the Algorithm used by "version < 80", It use AES-GCM to encrypt the password via a <master key> and a <nounce>.

The <master key> can be found in the "Local State" file, and can be decrypted by "CryptUnprotectData" mentioned above.

The <nounce> can be found at the beginning of the encrypted_password.

Therefore, we can decrypt all the password.

- Merge the result

If someone updates the Chrome recently, we need to find the two ways of User:Password pairs.

What's more, I use some rules to merge the results into an array.

## LICENSE

The Project follows MIT LICENSE.
