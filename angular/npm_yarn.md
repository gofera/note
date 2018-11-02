# node.js
Download node.js from its website, extract, and set to PATH.
```
export PATH="/home/weliu/tools/nodejs/node-v9.3.0-linux-x64/bin:$PATH"
```
Then you can use node and npm, test by checking version:
```
-bash-4.1$ node -v
v9.3.0
-bash-4.1$ npm -v
5.5.1
```

# npm
`npm` is like maven in Java or pip in Python.

## install a package/tool globally
```
npm install -g <package>
``` 
The package will be installed in the ${nodejs_root}/bin and ${nodejs_root}/lib.

## install a package/tool globally with specific version
```
npm install -g <package>@<version>
```
## remove a package/tool globally
```
npm uninstall -g <package>
```

I only suggest you to install global package, not dependency package in a project. For this purpose, use `yarn`.

So I prefer to use `npm` as a installer for tools, and use `yarn` as project dependency management like `maven`.

# yarn
`yarn` is also like maven in Java or pip in Python, but is better than npm to download dependency, the reason is npm will download the same package each time no matter the package does exist locally or not, yarn will download the same package only once for multiple projects.

## install yarn by npm
```
npm install -g yarn
```
check yarn version
```
yarn -v
```
## add a new dependency to the current directory (should be your angular project root directory)
```
yarn add <package...>[@<version>] [--dev/-D]
```
Similar to `npm` command, the version is optional (latest version as default). The command can support download multiple packages if passing more packages in the command. Using `--dev` or `-D` will install one or more packages in your devDependencies.

The package will download to the sub directry `node_modules` (created if not exist).

This will also update your package.json and your yarn.lock so that other developers working on the project will get the same dependencies as you when they run yarn or yarn install.

Most packages will be installed from the npm registry and referred to by simply their package name.

## yarn global
Similar to `npm install package -g`, use: `yarn global add package`. 

## remove dependency
```
yarn remove <package>[@<version>]
```

## update/download all dependencies in the current directory (should be your angular project root directory)
```
yarn install
```

# proxy
If you cannot visit web, you can try to set proxy.

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

# download is very slow in China, please use taobao image

