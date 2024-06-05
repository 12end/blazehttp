<div align="center">
<img src="https://github.com/chaitin/blazehttp/assets/30664688/746026da-6b2f-4f9c-86f1-1e3cb129ca22" width="120"/>
</div>
<h1 align="center">BlazeHTTP</h1>
<h4 align="center"><strong>简体中文</strong> | <a href="https://github.com/chaitin/blazehttp/blob/master/README_EN.md">
English</a></h4>

BlazeHTTP 是一款简单易用的 WAF **防护效果测试**工具。

- 📦 **样本丰富**：目前总样本**33669**条，持续更新中...
- 🚀 **无需配置**：提供**图形化界面**和命令行版本，可直接通过 Release 下载预编译版本，也可以克隆代码本地自行编译
- 📖 **报告导出**：导出所有样本的执行结果，包括样本属性，执行时间，状态码，是否拦截等

## 测试指标

|  指标   | 描述  | 统计方法  |
|  ----  | ----  | ----  |
| 检出率  | 用来反应 WAF 检测能力的全面性，没有检出即为 ”漏报“。 | 攻击样本拦截数量  |
| 误报率  | 用来反应对正常流量的干扰，不靠谱的结果即为 ”误报“。 | 正常样本拦截数量 |
| 准确率  | 准确率是检出率和误报率的综合指标，避免漏报和误报顾此失彼。 |  |
| 检测耗时  | 用来反应 WAF 性能，耗时越大则性能越差。 |  |

## 样本示例

```bash
# 正常样本：testcases/00/02/5ebf56a710da27b73a9ad59219f0.white
GET /rc-virtual-list@3.5.2/lib/hooks/useHeights.js HTTP/1.1
Host: npm.staticblitz.com
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36
Accept: */*
Origin: https://stackblitz.com
Sec-Fetch-Site: cross-site
Sec-Fetch-Mode: cors
Sec-Fetch-Dest: empty
Referer: https://stackblitz.com/
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7

# 黑样本：testcases/8a/36/0bbc7685860c526e33f3cbd83f9c.black
GET /vulnerabilities/sqli_blind/?id=1%27+or+%27%27%3D%27&Submit=Submit HTTP/1.1
Host: 10.10.3.128
Upgrade-Insecure-Requests: 1
User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36
Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7
Referer: http://10.10.3.128/vulnerabilities/sqli_blind/?id=1%27+and+%27%27%3D%27&Submit=Submit
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9,en-US;q=0.8,en;q=0.7
Connection: close
```

## 安装使用

GitHub CI 预编译的产物已上传 Release，可以[直接下载](https://github.com/chaitin/blazehttp/releases)最新的版本使用。

**命令行**

![blazehttp_cmd](https://github.com/chaitin/blazehttp/assets/30664688/7be052e9-2dfb-4f96-a6f2-eb2a0251910e)

**GUI** (MacOS & Windows)

> 如果 MacOS 双击打开报错**不受信任**或者**移到垃圾箱**，执行下面命令后再启动即可：
> ``` bash
> sudo xattr -d com.apple.quarantine blazehttp_1.0.0_darwin_arm64.app
> ```

![gui](https://github.com/chaitin/blazehttp/assets/30664688/dee16f13-8fef-413e-89c8-515b91c52c7a)

## 本地编译

项目只依赖了 Go 语言，首先你的环境上需要有 Go，可以在[这里](https://go.dev/dl/)下载

### 命令行版本

```bash
# 克隆代码
git clone https://github.com/chaitin/blazehttp.git && cd blazehttp
# 本地编译
bash build.sh # 执行后在 build 目录下看到 blazehttp
# 运行
./blazehttp -t https://example.org
```

### GUI 版本

GUI 是基于 [fyne](https://github.com/fyne-io/fyne) 实现。

```bash
# 克隆代码
git clone https://github.com/chaitin/blazehttp.git && cd blazehttp
# 本地运行
go run gui/main.go
```

<img width="810" alt="image" src="https://github.com/chaitin/blazehttp/assets/30664688/3d7f90aa-eb6d-43b0-adea-251114c6ea43">

> 如果需要本地打包，可以参考 fyne 的[打包文档](https://docs.fyne.io/started/packaging)
> 如果需要跨平台打包，也可以参考 [fyne-cross](https://docs.fyne.io/started/cross-compiling)

## 贡献代码

期待大佬们的贡献，添加新样本，新功能，修复 Bug，优化性能等等等等都非常欢迎👏

## Star

用起来还不错的话，帮忙点个 Star ✨
