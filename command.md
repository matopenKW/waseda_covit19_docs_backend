# Test
dev_appserver.py .

# gcloud deploy
+ 参考URL
    + https://blog.a-know.me/entry/2018/10/28/215508
+ export GO111MODULE=on
+ gcloud app deploy app.yaml --project waseda-covite19-docs

# docker
docker-compose exec -it app sh

# go module 
GO111MODULE=on go get -u 

# mySQL
+ localでコネクション接続 Userディレクトリで実行
    + $ cd ~
    + $ ./cloud_sql_proxy -instances=waseda-covite19-docs:us-central1:wasephil=tcp:13306 -credential_file=key_json/cloud_SQL/waseda-covite19-docs.json
