# 2017.12.15
## git
git diff // diff between unmanaged change and repository

git diff HEAD // diff between cached change (after git add command) and repository


git difftool [HEAD] // open the diff by 3rd-party tool like bcompare, should set in ~/.gitconfig (if git home enviroment HOME=~)
git mergetool [HEAD] // similar with above but for merge by 3rd-party tool

git stash // stash the current change in cached (staged), and revert the change.
git stash pop // pop the stash out

## yarn
```
npm install --global yarn
yarn upgrade // upgrade package.json to the latest version in the specified scope.
yarn upgrade --latest // command upgrades packages the same as the upgrade command, but ignores the version range specified in package.json. Instead, the version specified by the latest tag will be used (potentially upgrading the packages across major versions)
```

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

# 2017.12.27
## java run windows command
```
cmd /c dir
```
# 2018.1.8
## extract tar.xz
```
tar -Jxf node-v9.3.0-linux-x64.tar.xz
```
## maven build test but not run unit test
use `-DskipTests`

## VS Code Typescript plugin for checking style
Search `TSLink` in VS Code

## install npm
Just download node.js and extract, set to PATH env, then we can use `node` and `npm`.
```
node -v
npm -v
```

## set proxy for npm and yarn
Simply paste the following code at the bottom of your ~/.bashrc file:
```
######################
# User Variables (Edit These!)
######################
username="myusername"
password="mypassword"
proxy="mycompany:8080"

######################
# Environement Variables
# (npm does use these variables, and they are vital to lots of applications)
######################
export HTTPS_PROXY="http://$username:$password@$proxy"
export HTTP_PROXY="http://$username:$password@$proxy"
export http_proxy="http://$username:$password@$proxy"
export https_proxy="http://$username:$password@$proxy"
export all_proxy="http://$username:$password@$proxy"
export ftp_proxy="http://$username:$password@$proxy"
export dns_proxy="http://$username:$password@$proxy"
export rsync_proxy="http://$username:$password@$proxy"
export no_proxy="127.0.0.10/8, localhost, 10.0.0.0/8, 172.16.0.0/12, 192.168.0.0/16"

######################
# npm Settings
######################
npm config set registry http://registry.npmjs.org/
npm config set proxy "http://$username:$password@$proxy"
npm config set https-proxy "http://$username:$password@$proxy"
npm config set strict-ssl false
echo "registry=http://registry.npmjs.org/" > ~/.npmrc
echo "proxy=http://$username:$password@$proxy" >> ~/.npmrc
echo "strict-ssl=false" >> ~/.npmrc
echo "http-proxy=http://$username:$password@$proxy" >> ~/.npmrc
echo "http_proxy=http://$username:$password@$proxy" >> ~/.npmrc
echo "https_proxy=http://$username:$password@$proxy" >> ~/.npmrc
echo "https-proxy=http://$username:$password@$proxy" >> ~/.npmrc

######################
# yarn Settings
######################
yarn config set registry https://registry.yarnpkg.com/
yarn config set proxy "http://$username:$password@$proxy"
yarn config set https-proxy "http://$username:$password@$proxy"
yarn config set strict-ssl false

######################
# WGET SETTINGS
# (Bonus Settings! Not required for npm to work, but needed for lots of other programs)
######################
echo "https_proxy = http://$username:$password@$proxy/" > ~/.wgetrc
echo "http_proxy = http://$username:$password@$proxy/" >> ~/.wgetrc
echo "ftp_proxy = http://$username:$password@$proxy/" >> ~/.wgetrc
echo "use_proxy = on" >> ~/.wgetrc

######################
# CURL SETTINGS
# (Bonus Settings! Not required for npm to work, but needed for lots of other programs)
######################
echo "proxy=http://$username:$password@$proxy" > ~/.curlrc
```
Then edit the "username", "password", and "proxy" fields in the code you pasted.

Then call: `source ~/.bashrc`.

Check your settings by running `npm config list` and `cat ~/.npmrc` (or `yarn config list` and `cat ~/.yarnrc`).

Then you can use `npm install (-g)` (or `yarn install`) to install package from npm.


## how to fix git bash install in windows if very slow because home set to a remote location
set the env `HOME` to your local location, git bash will use it as home, I suggest to set `HOME` to your user home.

## git proxy setting
See my ~/.gitconfig file:
```
[user]
        email = liuwenzhe2008@qq.com
        name = WenzheLiu
[http]
        proxy = http://$username:$password@$host:$port
[https]
        proxy = https://$username:$password@$host:$port
[diff]
        tool = bc3
[difftool]
        prompt = false
[difftool "bc3"]
        path = /c/Users/weliu/AppData/Local/Beyond Compare 4/BComp.exe
[merge]
        tool = bc3
[mergetool]
        prompt = false
        keepBackup = false
[mergetool "bc3"]
        path = /c/Users/weliu/AppData/Local/Beyond Compare 4/BComp.exe
        trustExitCode = true
[alias]
        dt = difftool
        mt = mergetool

[url "https://"]
        insteadOf = git://
```
Replace the `$username`, `$password`, `$host` and `$port` with the real one. If the password includes special charactors such as `#`, use `%23` instead of `#`. For others special charactors, google yourself.

## angular cli
```
npm uninstall -g angular-cli  # this is old version ng cli, unintall
npm install -g @angular/cli   # new version is this
ng set --global packageManager=yarn  # managed by yarn
```

## RPM install to my directory without root
```
rpm2cpio ../google-chrome-stable_current_x86_64.rpm | cpio -idv
```
## portable app without install
https://portableapps.com/download

# old version firefox cannot load angular app
fix solution: find an older version angular cli, create a ng app for testing. For example:
```
npm uninstall -g @angular/cli
npm install -g @angular/cli@1.0.3
ng new xxx
ng serve
ng build
```

# 2018.01.16
## tree command ignore folder/pattern
```
tree -I 'test*|docs|bin|lib'
```

