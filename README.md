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



