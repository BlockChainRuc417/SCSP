# 基于区块链的服务协商框架

该平台在缺乏信任且无中心化的环境中为用户和服务提供方提供了一个协商平台。

# 上手指南

以下指南将帮助你在本地机器上安装和运行该项目，进行开发和测试。

## 安装要求
Hyperledger Fabric v1.1 (https://hyperledger-fabric.readthedocs.io/en/release-1.1/) 目前只支持1.1版本

[curl]( https://curl.haxx.se/)-v7.47.0（根据系统版本下载对应版本curl）

[docker](https://docs.docker.com/compose/overview/)-v17.06或以上版本

[docker-compose](https://docs.docker.com/compose/overview/)-v1.14或以上版本

[golang]( https://golang.google.cn/)-v1.10或以上版本

node.js：  v8.9或以上版本

## 安装步骤

### 安装curl：

1、	官网找到相应的版本，下载压缩包；

2、	在终端解压压缩包：

```
tar -zxvf curl-7.47.0.tar.gz
```

3、	进入到文件夹下：

```
cd curl-7.47.0
```

4、	终端命令

```
make

make install
```

### 安装docker：

1、	如果安装过旧版本的docker，先将其移除：

```
sudo apt-get remove docker-engine docker.io
```

2、	安装包允许apt通过HTTPS使用仓库：

```
sudo apt-get install apt-transport-https ca-certificates curl software-properties-common
```

3、	添加docker官方GPG key：
```
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

sudo apt-key fingerprint 0EBFCD88
```

4、	设置docker稳定版仓库：

```
sudo add-apt-repository “deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable”
```

5、	更新apt索引源：

```
sudo apt-get update
```

6、	安装docker CE（社区版）：

```
sudo apt-get install docker-ce
```

7、	如果上一步出现错误，检查一下版本：

```
apt-cache madison docker-ce
```

8、	安装指定版本：

```
sudo apt-get install docker-ce=<version>
```

9、	测试docker是否安装成功：

```
sudo docker run hello world
```

10、	查看docker版本：

```
docker -v
```

### 安装docker-compose：
1、	使用命令行下载docker-compose：

```
sudo curl -L https://github.com/docker/compose/releases/download/1.17.0/docker-compose-`uname -s`-`uname -m` -o /usr/local/bin/docker-compose
```

2、	授权：

```
sudo chmod +x /usr/local/bin/docker-compose
```

3、	查看版本信息：

```
docker-compose –version
```

### 安装golang：
1、	安装：

```
sudo apt-get install golang
```
2、	设置环境变量：打开配置文件

```
gedit /etc/profile
```

末尾添加如下信息：

```
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

3、	配置文件生效：

```
source /etc/profile
```

4、	查看golang版本：

```
go version
```

### 安装node.js：
1、	安装node.js：

```
sudo apt-get install nodejs
```

2、	安装npm：

```
sudo apt-get install npm
```

### 下载镜像：

我们使用fabric1.1版本镜像，使用官网链接下载即可：

```
curl -sSL https://goo.gl/6wtTN5 | bash -s 1.1.0-rc1
```

## 网络要求

包括orderer节点、peer节点、kafka节点和zookeeper节点，每种节点的个数可以根据自己的需要在配置文件中进行修改即可，在我们的项目中，使用了3个orderer节点、4个peer节点（两个组织，每个组织两个peer节点）、4个kafka节点、3个zookeeper节点（个数必须是奇数）

## 部署

SCSP主要是区块链部分，提供了关于服务交易的一些功能，部署位置没有特别要求，为了操作方便，我们放在根目录下即可；  

## 启动

### Terminal 1：启动网络

进入SCSP文件夹：

```
cd SCSP
```

运行脚本：

```
./runApp.sh
```

### Terminal 2

进入scripts文件夹：

```
cd scripts
```

运行install脚本：

```
./install.sh
```

注意：运行该脚本后，可能会提示未安装`jq`，按照提示要求安装即可；该脚本的作用是创建channel、把节点加入到channel中、在每个组织中安装链码（chaincode）并在节点上实例化；

运行regist脚本：

```
./regist.sh
```

该脚本的作用是用户注册，为进行注册的用户生成身份凭证token；

运行task脚本：

```
./task.sh
```

该脚本用于任务发布测试；

运行listener脚本：

```
./listener.sh
```

该脚本的作用是将服务提供方加入到候选队列中；

运行request脚本：

```
./request.sh
```

+ 该脚本的作用是在众多的服务提供方提供的服务中选择一个性能最优的作为最后所使用的服务；

至此，我们完成了通过脚本运行测试整个项目，可以通过界面查看最终平台展示效果。

