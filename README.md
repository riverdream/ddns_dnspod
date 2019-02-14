# ddns_dnspod

使用[dnspod](https://www.dnspod.cn/)作为动态域名时，本地需要定时检查ip地址是否变化，并把变化的地址更新到服务器。

代码使用golang语言，若要配置golang开发环境请参考[Golang开发环境配置](https://www.jianshu.com/p/8eb8efbfb4d1)

非开发人员可以直接下载使用：

Linux版本下载：[Linux_x86_ddns-64bit](http://download.pengjx.com/2017/dnspod/linux_x86_ddns-64bit)

Windows版本下载：[Windows_x86_ddns-64bit.exe](http://download.pengjx.com/2017/dnspod/windows_x86_ddns-64bit.exe)

使用的格式如下：

dnspod -login_token "id,token" -sub_domain "www" -domain_id "1234567" -record_id "1234567"

---

dnspod api https://www.dnspod.cn/docs/index.html

0、api 规范说明

https://www.dnspod.cn/docs/info.html

1、获取doman id

https://www.dnspod.cn/docs/domains.html#domain-list

curl --user-agent "pengjx DDNS Client/1.0.0 (yourmail@yeah.net)" -X POST https://dnsapi.cn/Domain.List -d 'login_token=id,token'

2、获取 sub doman id

https://www.dnspod.cn/docs/records.html#record-list

curl --user-agent "pengjx DDNS Client/1.0.0 (yourmail@yeah.net)" -X POST https://dnsapi.cn/Record.List  -d 'login_token=id,token&domain_id=xxxxx'
