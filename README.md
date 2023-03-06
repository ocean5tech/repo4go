# repo4go
# 第1章2022新版go工程师体系课导学以及Go的发展

# 安装
安装在ubuntu然后启动vs code来直接编辑source

        wget https://go.dev/dl/go1.20.1.linux-amd64.tar.gz
        rm -rf /usr/local/go
        sudo tar -C /usr/local -xzf go1.20.1.linux-amd64.tar.gz
        export PATH=$PATH:/usr/local/go/bin
        go version
        code .