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
## yarn
yarn global install a package, like angular cli:
```
yarn global add @angular/cli
```
Also need to add yarn path to the PATH. So it is better to use npm to install global, because it is no need to cache for global, so use npm is better.
```
npm install -g @angular/cli
```

# 2017.12.22
## maven
### create a maven project by command line
```
mvn archetype:generate -DgroupId=com.wenzhe -DartifactId=cwm -DarchetypeArtifactId=maven-archetype-quickstart -DinteractiveMode=false
```
The generated maven java project has a root folder named the same with artifactId.

### create eclipse project from maven command line
```
cd $project_root
mvn eclipse:eclipse
```
For idea, no action need to do, just open directly.

## git bash
### install linux command to git bash in windows
For example, to install `tree` command, we can download the `tree.exe` from `http://gnuwin32.sourceforge.net/packages/tree.htm` and unzip, put the tree.exe to `$git_install_location/usr/bin` (the `$git_install_location` is root path in linux, i.e. `/`.) Then we can use it in Linux.
```
$ which tree
/usr/bin/tree
$ tree
.
|-- README.md
`-- cwm
    |-- pom.xml
    `-- src
        |-- main
        |   `-- java
        |       `-- com
        |           `-- wenzhe
        |               `-- App.java
        `-- test
            `-- java
                `-- com
                    `-- wenzhe
                        `-- AppTest.java
```

