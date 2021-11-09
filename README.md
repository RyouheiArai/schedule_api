# schedule_api

スケジュール管理アプリ（バックエンド部分）

フロントエンド・・・Vue/Vuex  
バックエンド・・・Go/Gin + MySQL  
フロントエンドリポジトリは以下参照。  
https://github.com/RyouheiArai/schedule_front
 
## 簡単な説明
 
スケジュール管理、ユーザ管理に必要なAPIを提供。  
APIは全てrest（json）で提供し、一部のAPIでJWT認証を実施。
 
## 機能
 
ユーザ登録,ログインAPI  
スケジュール登録、取得、変更、削除機能API（JWT認証）　　  
認証機能（JWT認証）
 
## インストール
※MySQLを環境にインストールしてください。  
 
```
$ git clone https://github.com/RyouheiArai/schedule_api.git
$ cd schapi
$ cp .env.example .env // MySQLパスワード等設定
$ go run schapi  // 実行
```
## 今後のアップデート予定
・Docker環境で動作できるようにする  
・googleカレンダーのデータ取り込み  
・一応クライアントがいる為、要望にそった機能を追加予定  

