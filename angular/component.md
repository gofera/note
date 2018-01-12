# apply CSS style
## directly binding with [style.background-color]
```
<div [style.background-color]="getStyle()">
...
getStyle() { return "yellow"; }
```
## adding a class [class.my-class]
```
<div [class.my-class]="isClassVisible">
...
isClassVisible = false;
```
## using NgClass [ngClass]
```
[ngClass]="{ 'myclass1': true, 'myclass2': false }"
```
