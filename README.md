# goimg

#### 介绍
一个轻量型的图片服务器


#### 软件架构
上传接口：
http://127.0.0.1:8080/upload

参数：Files 类型：文件


返回结果：
[{"success":true,"message":"OK","version":"v0.1.1","data":{"size":49160,"mime":"image/jpeg","fileId":"5781339b809d5f18132f5c4fbe9df2fe","fileName":"gss0.baidu.jpg"}}]

使用说明：
http://127.0.0.1:8080/5781339b809d5f18132f5c4fbe9df2fe  默认：压缩质量为75%

访问原图：
http://127.0.0.1:8080/5781339b809d5f18132f5c4fbe9df2fe?p=1   //p=1 查看原始图片

下载图片
http://127.0.0.1:8080/5781339b809d5f18132f5c4fbe9df2fe?d=1  //d=1 下载图片，浏览器不再展示图片

灰阶图
http://127.0.0.1:8080/5781339b809d5f18132f5c4fbe9df2fe?g=1  //g=1 灰阶图

缩放
http://127.0.0.1:8080/5781339b809d5f18132f5c4fbe9df2fe?w=100&h=100  //w宽度 h高度

压缩
http://127.0.0.1:8080/5781339b809d5f18132f5c4fbe9df2fe?q=75     //q 压缩质量 

转换格式
http://127.0.0.1:8080/5781339b809d5f18132f5c4fbe9df2fe?f=png    //f 转换格式 

旋转
http://127.0.0.1:8080/5781339b809d5f18132f5c4fbe9df2fe?r=90   //r 旋转图像



#### 安装教程

1.  xxxx
2.  xxxx
3.  xxxx

#### 使用说明

1.  xxxx
2.  xxxx
3.  xxxx

#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request


#### 码云特技

