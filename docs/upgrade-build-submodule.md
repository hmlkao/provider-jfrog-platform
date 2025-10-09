# How to upgrade `build` submodule

1. Check current version of git submodules

    ```shell
    git submodule
    ```

2. (*optional*) You can find Git commit hash in history to see how old is your current version

    ```shell
    cd build
    git log --pretty=oneline
    ```

3. Upgrade all submodules

    ```shell
    git submodule update --recursive --remote
    ```

4. Commit changes

    ```shell
    git add .
    git commit -m "Upgraded build git submodule"
    ```
