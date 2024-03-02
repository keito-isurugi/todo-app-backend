#!/bin/bash

# goenvのインストール
if [ ! -d ~/.goenv ]; then
    git clone https://github.com/syndbg/goenv.git ~/.goenv

    # goenvを利用可能にするために、現在のシェルセッションを更新
    echo 'export GOENV_ROOT="$HOME/.goenv"' >> ~/.zshrc
    echo 'export PATH="$GOENV_ROOT/bin:$PATH"' >> ~/.zshrc
    echo 'eval "$(goenv init -)"' >> ~/.zshrc
    source ~/.zshrc
fi

GO_VERSION="1.20.5"

# Goの指定のバージョンがインストールされていなければインストール
if ! goenv versions | grep -q "$GO_VERSION"; then
    goenv install $GO_VERSION
fi

# 指定のバージョンをデフォルトとして設定
goenv global $GO_VERSION