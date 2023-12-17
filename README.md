# embulk-hands-on-local-postgresql
Hands-on with Embulk to extract data from a database and convert to csv.


## ハンズオン手順
- docker環境の構築  
`docker-compose build`
`docker-compose up -d`

- `make init`  
postgresにteasテーブルを作成
teasテーブルにデータを登録

- dockerコンテナの中に入る  
`docker exec -it embulk bash`

- embulkの処理を実行。postgresの中身が./output_files/teas.csvとして出力される  
`embulk run config.yml`
