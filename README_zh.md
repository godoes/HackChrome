# HackChrome

[![Build Status](https://travis-ci.com/godoes/HackChrome.svg?branch=master)](https://travis-ci.com/godoes/HackChrome)

[English](https://github.com/godoes/HackChrome/blob/master/README.md) | **中文**

从 Chrome 中获取自动保存的用户名密码

## 影响 Chrome 版本

全版本

## 平台

Windows

## 使用

- 安装

  ```shell
  set CGO_ENABLED=1 & go install github.com/godoes/HackChrome@latest
  ```

- 打开 cmd 或者 powershell

  ```shell
  HackChrome -h

  # output:
  Usage of HackChrome:
    -edge
          true or false, default is false and chrome is used.
  ```

- 运行

  ```shell
  # Google Chrome:
  HackChrome > netpass.txt

  # Microsoft Edge:
  HackChrome -edge > netpass.txt
  ```

## 效果截图

![demo](image/result.png)

## 原理

- version < 80

Chrome的用户名密码存储在 "Login Data" 文件中。

密码是加密的，但我们可以使用 "Crypt32.dll" 中的 "CryptUnprotectData" 函数来解密密码。

最终，我们获取到了明文的用户名密码。

- version > 80

算法整体基于 "version < 80" 的算法，然而现在利用一个 master key 和 nounce，使用 AES-GCM 加密密码。

master key 在 "Local State" 文件中，可以使用上面提到的 "CryptUnprotectData" 函数解密。

nounce 则在密文的头部。

因此，我们可以解密出明文密码。

- 结果合并操作

如果刚刚升级 Chrome 到 v80，之前保存的密码和之后保存的密码存储加密方式将会不同，所以 HackChrome 将两组结果
以某种算法聚合起来，得到最终的用户名密码数组。

## 协议

本项目遵循 MIT 协议
