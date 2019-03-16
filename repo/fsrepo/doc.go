
//<developer>
//    <name>linapex 曹一峰</name>
//    <email>linapex@163.com</email>
//    <wx>superexc</wx>
//    <qqgroup>128148617</qqgroup>
//    <url>https://jsq.ink</url>
//    <role>pku engineer</role>
//    <date>2019-03-16 19:56:44</date>
//</624460183627632640>

//包装FSRPO
//
//要解释程序包路线图…
//
//IPFS/
//——客户/
//——client.lock<------保护client/+信号自己的PID
//珍——ipfs-client.cpuprof
//│——ipfs-client.memprof
//——配置
//——守护进程/
//珍——daemon.lock<------保护daemon/+信号自己的地址
//珍——ipfs-daemon.cpuprof
//│———ipfs-daemon.memprof
//——数据存储/
//——repo.lock<------保护数据存储/和配置
//破译——版本
package fsrepo

//TODO防止多个守护进程运行

