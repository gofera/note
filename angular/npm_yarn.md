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

# download is very slow in China, please use taobao image

