# cnback

Originally from https://github.com/stefanprodan/mgob

CNback is a cloud native backup automation tool built with golang.

#### Features

* schedule backups with cron-style parameter
* multi source database
** mysql
** mongodb
** postgres
* multi storage destination
** local backups retention
** upload to S3 Object Storage (Minio, AWS, Google Cloud)
** upload to gcloud storage
** upload to SFTP
* multi notification platforms
** emails
** slack
* distributed as an Alpine Docker image

#### Install

#### Configure

Define a backup plan (yaml format) for each database you want to backup inside the `config` dir. 
The yaml file name is being used as the backup plan ID, no white spaces or special characters are allowed. 

_Backup plan_

```yaml
scheduler:
  # run every day at 6:00 and 18:00 UTC
  cron: "0 6,18 */1 * *"
  # number of backups to keep locally
  retention: 14
  # backup operation timeout in minutes
  timeout: 60
target:
  # mongod IP or host name
  host: "172.18.7.21"
  # mongodb port
  port: 27017
  # mongodb database name, leave blank to backup all databases
  database: "test"
  # leave blank if auth is not enabled
  username: "admin"
  password: "secret"
  # add custom params to mongodump (eg. Auth or SSL support), leave blank if not needed
  params: "--ssl --authenticationDatabase admin"
# S3 upload (optional)
s3:
  url: "https://play.minio.io:9000"
  bucket: "backup"
  accessKey: "Q3AM3UQ867SPQQA43P2F"
  secretKey: "zuf+tfteSlswRu7BJ86wekitnifILbZam1KYY3TG"
  api: "S3v4"
# GCloud upload (optional)
gcloud:
  bucket: "backup"
  keyFilePath: /path/to/service-account.json
# SFTP upload (optional)
sftp:
  host: sftp.company.com
  port: 2022
  username: user
  password: secret
  # dir must exist on the SFTP server
  dir: backup
# Email notifications (optional)
smtp:
  server: smtp.company.com
  port: 465
  username: user
  password: secret
  from: cnback@company.com
  to:
    - devops@company.com
    - alerts@company.com
# Slack notifications (optional)
slack:
  url: https://hooks.slack.com/services/xxxx/xxx/xx
  channel: devops-alerts
  username: cnback
  # 'true' to notify only on failures 
  warnOnly: false
```

ReplicaSet example:

```yaml
target:
  host: "mongo-0.mongo.db,mongo-1.mongo.db,mongo-2.mongo.db"
  port: 27017
  database: "test"
```

Sharded cluster with authentication and SSL example:

```yaml
target:
  host: "mongos-0.db,mongos-1.db"
  port: 27017
  database: "test"
  username: "admin"
  password: "secret"
  params: "--ssl --authenticationDatabase admin"
```

#### Restore

In order to restore from a local backup you have two options:

Browse `cnback-host:8090/storage` to identify the backup you want to restore. 
Login to your MongoDB server and download the archive using `curl` and restore the backup with `mongorestore` command line.

```bash
curl -o /tmp/mongo-test-1494056760.gz http://cnback-host:8090/storage/mongo-test/mongo-test-1494056760.gz
mongorestore --gzip --archive=/tmp/mongo-test-1494056760.gz --drop
```

You can also restore a backup from within cnback container. 
Exec into cnback, identify the backup you want to restore and use `mongorestore` to connect to your MongoDB server.

```bash
docker exec -it cnback sh
ls /storage/mongo-test
mongorestore --gzip --archive=/storage/mongo-test/mongo-test-1494056760.gz --host mongohost:27017 --drop
```
# cnback
