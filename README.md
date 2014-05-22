Properties
========
[![Build Status](https://drone.io/github.com/widuu/goini/status.png)](https://drone.io/github.com/nosix-me/Properties/1)

##描述

使用Properties读取properties文件。

##安装方法
	
	go get github.com/nosix-me/Properties

##使用方法

>properties配置文件样式样例
	
	#test
	string=hello

>初始化

	p, err := NewProperties("./settings.properties")
>修改配置文件路径

	p.SetFilepath("./settings.properties")
>读取配置信息

	p.ReadProperties()
>获取配置信息

	p.Get("string")
>获得所有配置信息

	p.GetProperties()
---


	
