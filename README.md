## 二维码模块
### 新增二维码
新增二维码需要填写植物的学名，上传一张图片，填写植物的介绍信息。

提交之后，系统会根据提交的内容生成一个访问链接，点击该链接可以预览该植物页面内容。

数据结构：

id、name、desc是存储在数据库里的静态信息；

link、pic、code存储相对路径信息；

其中link与id有关，数据完整性能够保证。pic与code需要考虑文件不统一的问题。为了处理文件统一性质的问题，每次使用get方法检查的同时，检查文件夹里是否存在相应的pic，执行下载的时候也检查是否存在qrcode。

read每次访问时+1，数据完整性也较好。

### 修改二维码
修改二维码可以修改植物的学名，图片和植物的描述信息。

### 删除二维码
在预览页面点击删除可以将该二维码的信息删除

## 栏目模块
### 预览栏目&新增栏目
点击栏目按钮，右侧有一列当前栏目状态的预览栏。左侧有一个表单用来新增加栏目

## 数据结构
- 支持最多2级目录，多的无法添加；
- 已经有挂载文章、页面、栏目的无法执行删除操作
- 选择父目录只能选择类型为目录类型的目录
- 需要设定几个目录无法修改
- 需要显示目录的类型和旗下挂载的内容
- 在不能删除的栏目上不显示删除按钮
- 给添加部分增加一个二维码验证功能
- 需要针对不同的类型展示不同的数量
