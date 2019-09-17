# SCSP
# 基于区块链的服务交易平台
该平台在缺乏信任且无中心化的环境中为用户和服务提供方提供了一个自主协商的服务交易平台，保证交易的安全、合理。
# 上手指南
以下指南将帮助你在本地机器上安装和运行该项目，进行开发和测试。
## 安装要求
     curl:              v7.47.0（根据系统版本下载对应版本curl）
       docker：          v17.06或以上版本
     docker-compose：   v1.14或以上版本
     golang：           v1.10或以上版本
     node.js：           v8.9或以上版本
## 安装步骤
### 安装curl：
	官网找到相应的版本，下载压缩包；
	在终端解压压缩包tar -zxvf curl-7.47.0.tar.gz；
	进入到文件夹下cd curl-7.47.0
	终端命令make
	终端命令make install
### 安装docker：
	如果安装过旧版本的docker，先将其移除：sudo apt-get remove docker-engine docker.io；
	安装包允许apt通过HTTPS使用仓库：sudo apt-get install apt-transport-https ca-certificates curl software-properties-common；
	添加docker官方GPG key：curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -；
	sudo apt-key fingerprint 0EBFCD88；
	设置docker稳定版仓库：sudo add-apt-repository “deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable”；
	更新apt索引源：sudo apt-get update；
	安装docker CE（社区版）：sudo apt-get install docker-ce；
	如果上一步出现错误，检查一下版本：apt-cache madison docker-ce；
	安装指定版本：sudo apt-get install docker-ce=<version>；
	测试docker是否安装成功：sudo docker run hello world，出现如下信息表示安装成功：
 
	查看docker版本：docker -v
### 安装docker-compose：
	使用命令行下载docker-compose：sudo curl -L https://github.com/docker/compose/releases/download/1.17.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose；
	授权：sudo chmod +x /usr/local/bin/docker-compose；
	查看版本信息：docker-compose --version
### 安装golang：
	安装：sudo apt-get install golang；
	设置环境变量：打开配置文件gedit /etc/profile，末尾添加如下信息：
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
	配置文件生效：source /etc/profile；
	查看golang版本：go version；
### 安装node.js：
	安装node.js：sudo apt-get install nodejs；
	安装npm：sudo apt-get install npm；
### 下载镜像：
我们使用fabric1.1版本镜像，使用官网链接下载即可：curl -sSL https://goo.gl/6wtTN5 | bash -s 1.1.0-rc1；
## 网络要求
包括orderer节点、peer节点、kafka节点和zookeeper节点，每种节点的个数可以根据自己的需要在配置文件中进行修改即可，在我们的项目中，使用了3个orderer节点、4个peer节点（两个组织，每个组织两个peer节点）、4个kafka节点、3个zookeeper节点（个数必须是奇数）
##部署
SCSP主要是区块链部分，提供了关于服务交易的一些功能，部署位置没有特别要求，为了操作方便，我们放在根目录下即可；
ctrip主要提供展示界面，需要将该文件夹部署在/tomcat/apache-tomcat-8.5.28/webapps/ROOT路径下；
## 启动
### Terminal 1：启动网络
	进入SCSP文件夹：cd SCSP；
	运行脚本：./runApp.sh；
### Terminal 2
	进入scripts文件夹：cd scripts；
	运行install脚本：./install.sh；
注意：运行该脚本后，可能会提示未安装`jq`，按照提示要求安装即可；该脚本的作用是创建channel、把节点加入到channel中、在每个组织中安装链码（chaincode）并在节点上实例化；
	运行regist脚本：./regist.sh；
该脚本的作用是用户注册，为进行注册的用户生成身份凭证token；
	运行task脚本：./task.sh；
该脚本的作用是用户将所需服务的性能要求写入智能合约，并发布任务；
	运行listener脚本：./listener.sh；
该脚本的作用是将服务提供方加入到候选队列中；
	运行request脚本：./request.sh；
该脚本的作用是在众多的服务提供方提供的服务中选择一个性能最优的作为最后所使用的服务；
至此，我们完成了通过脚本运行测试整个项目，可以通过界面查看最终平台展示效果。
