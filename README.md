# 2017.12.15
## git
git diff // diff between unmanaged change and repository
git diff HEAD // diff between cached change (after git add command) and repository


git difftool [HEAD] // open the diff by 3rd-party tool like bcompare, should set in ~/.gitconfig (if git home enviroment HOME=~)
git mergetool [HEAD] // similar with above but for merge by 3rd-party tool

git stash // stash the current change in cached (staged), and revert the change.
git stash pop // pop the stash out

## yarn
yarn upgrade // upgrade package.json to the latest version in the specified scope.
yarn upgrade --latest // command upgrades packages the same as the upgrade command, but ignores the version range specified in package.json. Instead, the version specified by the latest tag will be used (potentially upgrading the packages across major versions)

# 2017.12.18
## git
### how to put git add commit push in one command?
The way is to write the following code in ~/.bashrc:
```
function git_add_commit_push() {
    git add .
    git commit -a -m "$1"
    git push
}
```
Then in bash, call:
```
source ~/.bashrc  # make the setting available

# call the command
git_add_commit_push "my description ..."
```
### How to handle when git failed with ssh connection failed.
#### One way is the use https instead of ssh, like this:
```
git config --global url."https://".insteadOf git://
```
Then you can see the ~/.gitconfig add the following code automatically:
```
[url "https://"]
        insteadOf = git://
```
If this way still cannot work, try the following way:
#### use another port instead ssh 22
```
$ vi ~/.ssh/config
```
Edit the file with the following content:
```
Host github.com
User liuwenzhe2008@qq.com
Hostname ssh.github.com
PreferredAuthentications publickey
IdentityFile ~/.ssh/id_rsa
Port 443
```

