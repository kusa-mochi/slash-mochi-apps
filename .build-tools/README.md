# ビルドツール

## 前提環境

- Docker Desktop インストール済

## Linux-AMD64向けビルド手順

1. `.env_linux-amd64`ファイルの中身をすべて`.env`ファイルに上書きペーストし保存する。
1. `.build-tools/`ディレクトリ直下でコマンド`docker compose up`を実行する。
1. `dist/linux/amd64/`ディレクトリ以下にビルド成果物が生成される。

## Windows-AMD64向けビルド手順（デバッグ用途を想定）

1. `.env_windows-amd64`ファイルの中身をすべて`.env`ファイルに上書きペーストし保存する。
1. `.build-tools/`ディレクトリ直下でコマンド`docker compose up`を実行する。
1. `dist/windows/amd64/`ディレクトリ以下にビルド成果物が生成される。

# エディタツール

## 前提環境

- Docker Desktop インストール済
- VSCode インストール済
- VSCode Docker拡張機能 インストール済

## サーバーソースエディター起動手順

1. VSCodeを起動する。
1. VSCodeのターミナルを開く。
1. ターミナルで`.build-tools/`ディレクトリに移動する。
1. ターミナルで`docker compose -f docker-compose.dev.yml up server_editor -d`を実行する。
1. VSCode左サイドバーのDockerアイコンをクリックする。
1. 起動中のコンテナ（CONTAINERS - build-tools - whoatemyapplepie/slash-mochi-go-builder:0.1.0）を右クリックする。
1. "Attach Visual Studio Code" をクリックする。新たにVSCodeのウィンドウが表示される。
1. 左サイドバーに "Open Foler" というボタンが表示されている場合、そのボタンをクリックし、`/project`ディレクトリを開く。
1. 拡張機能「Go」をインストールする。コンテナ内に拡張機能がインストールされる。
1. 以降、コンテナ内でコード編集を行える。

## クライアントソースエディター起動手順

雰囲気はサーバーソースエディターと一緒。

ただし、クライアントはサービス毎に別ディレクトリで開発するところだけ注意。

例：テスト用クライアントは`cmd/clients/test`ディレクトリ以下で開発する。