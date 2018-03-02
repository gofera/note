# update what will be committed
```
git add <file>...
```
Usually use the following command to update current folder:
```
git add .
```

# discard changes in the working directory
```
git checkout -- <file>...
```
# file change history
```
gitk <file>
```
or:
```
gitk --follow <file>
```
For example:
```
gitk cwm-ng/src/app/app.module.ts
```

# disable converting to windows line end
```
$ git config --global core.autocrlf false //禁用自动转换
```

# git find history
```
$ git log
```

# git list branch
```
git branch
```

# git create new branch
```
git branch <new_branch_name>
```

# git switch branch
```
git checkout <other_branch_name>
```

# git create and switch branch
```
git checkout -b <new_branch_name>
```
