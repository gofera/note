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

# git delete branch
```
git branch -d <branch_name>
```

# 引用其它项目作为子模块
创建.gitmodules文件，参考：[Git 工具 - 子模块](https://git-scm.com/book/zh/v1/Git-%E5%B7%A5%E5%85%B7-%E5%AD%90%E6%A8%A1%E5%9D%97)

# 忽略文件（但不是ignore，不能提交的忽略，在git status也不提示）
```
git update-index --assume-unchanged <file_path>    # 忽略：git认为文件没变化
git update-index --no-assume-unchanged <file_path> # 撤销忽略
git ls-files -v | grep '^h'                        # 列出忽略的文件
```

# show the file content before modified
```
git show HEAD:<file path>
```

# git pull local repository
For github, we can use `git pull`, for local repo, we can use:
```
git pull <local_repo_dir>
```
