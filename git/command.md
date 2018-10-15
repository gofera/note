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

# Submit existing git project to remote git repository (new)
First, create a new repository in GitHub/Bitbucket. Then:
```
$ cd ${your_code_root}
$ git remote add origin ${remote_git_repository}
```
If origin has already existed, the above command will fail, we can remove it first:
```
$ git remote -v
origin  ${original_git_repository} (fetch)
origin  ${original_git_repository} (push)
$ git remote remove origin
```
Or change instead remove/add:
```
$ git remote set-url origin ${remote_git_repository}
```
And then add as the above command. Then push the code to remove repository:
```
git push -u origin master
```

# Tip for git completion in Linux
```
$ vi ~/.bashrc
```
Then add the following:
```
for file in /etc/bash_completion.d/* ; do
  source "$file"
done
```
Then `source ~/.bashrc`. Then your git will have tip for completion by press tab (or double press):
```
-bash-4.1$ git push origin 
comment_for_activator          FETCH_HEAD                     master                         origin/comment_for_activator 
control_access                 HEAD                           ORIG_HEAD                      origin/master
```

# git无法pull仓库refusing to merge unrelated histories
```
git pull --allow-unrelated-histories
```

# git issue: SSL certificate problem: Unable to get local issuer certificate
refer: [SSL certificate problem: Unable to get local issuer certificate](https://confluence.atlassian.com/bitbucketserverkb/ssl-certificate-problem-unable-to-get-local-issuer-certificate-816521128.html). 

Work around: Tell git to not perform the validation of the certificate using the global option:

```
git config --global http.sslVerify false
```
