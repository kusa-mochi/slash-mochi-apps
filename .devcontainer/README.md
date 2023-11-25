# ビルドツール

## Linux向けビルド手順

1. `.env_linux`ファイルの中身をすべて`.env`ファイルに上書きペーストし保存する。
1. `.build-tools/`ディレクトリ直下でコマンド`docker compose up`を実行する。
1. `dist/linux/amd64/`ディレクトリ以下にビルド成果物が生成される。

## Windows向けビルド手順（デバッグ用途を想定）

1. `.env_windows`ファイルの中身をすべて`.env`ファイルに上書きペーストし保存する。
1. `.build-tools/`ディレクトリ直下でコマンド`docker compose up`を実行する。
1. `dist/windows/amd64/`ディレクトリ以下にビルド成果物が生成される。
