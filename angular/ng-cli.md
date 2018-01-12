# install angular-cli globally
```
npm install -g @angular/cli
```
The latest will be installed by the above command.

For install a specific version, for example, a very old version 1.0.3:
```
npm install -g @angular/cli@1.0.3
```

# uninstall
```
npm uninstall -g @angular/cli@1.0.3
```

# manage angular cli by yarn
The default dependency manager is `npm`. I suggest to use `yarn` to download and manage dependencies, config like this:
```
ng set --global packageManager=yarn  # managed by yarn
```

# create new angular project
```
ng new <project>
```


