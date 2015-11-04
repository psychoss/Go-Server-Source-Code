
<http://www.mean101.com> 服务器部分源代码

安装和使用

## 依赖项目

		# 安装 LevelDB
		sudo apt-get install libleveldb-dev
		# 安装 Google 的 Snappy 压缩模块
		sudo apt-get install libsnappy-dev
		# 安装 正则表达式
		sudo apt-get install libonig2
		# 安装 pkg-config 
		sudo apt-get install pkg-config



如果出现" Package oniguruma was not found in the pkg-config search path" 这样的错误：

		export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:$GOPATH/src/web/rubex

如果出现" fatal error: oniguruma.h: No such file or directory" 这样的错误：
		
		sudo apt-get -y install libonig-dev


