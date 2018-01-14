# click <a> like a button
Not refresh page or redirect to any URL, just call a method like button. 

The way is set `href` to `javascript:void(0)`:
```
<a href="javascript:void(0)" (click)="onClickWorkerCount(server)">{{workerCount(server)}}</a>
```

