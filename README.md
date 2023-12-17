# embulk-hands-on-local-postgresql
Hands-on with Embulk to extract data from a database and convert to csv.


docker-compose build
docker-compose up -d
・ docker環境の構築

make initで初期設定。
・　postgresにteasテーブルを作成
・ teasテーブルにデータを登録

docker exec -it embulk bash
・ dockerコンテナの中に入る

embulk run config.yml
・ embulkの処理を実行。postgresの中身が./output_files/teas.csvとして出力される。
