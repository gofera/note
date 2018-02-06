# 自动导入指定类的方法
在`IDE`中输入方法名，能够补全提示，能够自动地静态导入该方法，设置方法如下：在`Preferences`中，`Java/Editor/Content Assist/Favorites`，点击`New Type...`按钮，`Browser`指定的类，这样，在该类中的所有公有静态方法和变量都能够自动导入。

# 指定Java版本
在eclipse.ini文件中，加入：
```
-vm
${your_java_home}/bin/javaw.exe
```
