# Angular CLI的跨域支持

前后端独立开发测试，启动一个后端服务（比如localhost:8080/api/xxx）提供REST api服务，前端用`ng serve`（默认端口4200），会出现跨域问题，没法访问后端服务，解决的办法不需要后端支持，加上代理即可：

(1) 在项目根目录创建`proxy.conf.json`:
```
{
    "/api": {
        "target": "http://localhost:8080",
        "secure": false
    }
}
```
上面的api为后端REST API服务的前缀，比如get xxx，后端的URL为localhost:8080/api/xxx，angular中可以这样写：
```
this.http.get('/api/xxx')
```
不管是开发独立测试，还是将angular打包到dist并与后端产品一起打包，都不必修改代码。

(2) 前端启动方式：
```
ng serve --proxy-config proxy.conf.json
```
也可以写入在package.json中，
```
  "scripts": {
    "ng": "ng",
    "start": "ng serve --proxy-config proxy.conf.json",
```
然后启动方式如下：
```
npm start
```

这对应前后端独立开发测试非常有用，不影响打包运行。

# How can I convert strings of HTML into HTML in my template (Angular2/TypeScript)?
```
<div [innerHtml]="someProp.htmlText"></div>
``` 

# Reference
[解决angular4本地开发中跨域访问的问题](https://www.bf361.com/wpdesign/angular-ng-serve-proxy-config)

[How can I convert strings of HTML into HTML in my template](https://stackoverflow.com/questions/36921180/how-can-i-convert-strings-of-html-into-html-in-my-template-angular2-typescript)


