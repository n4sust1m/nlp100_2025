# nlp100_2025

https://nlp100.github.io/2025/ja/index.html を解きます。AI は使いません。

前回: https://github.com/n4sust1m/nlp100

## Setup

[mise](https://github.com/jdx/mise) を利用して各言語のランタイムをインストールします

macOS & zsh の場合

```bash
# install
brew install mise
```

`~/.zshrc` に以下の設定を追加

```bash
# mise
if type mise &>/dev/null; then
  eval "$(mise activate zsh)"
  eval "$(mise activate --shims)"
fi
```