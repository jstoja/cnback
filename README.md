# cnback

*This project is currently not production ready. If you still want to use it, it's to your own riscs and perils.
Originally from https://github.com/stefanprodan/mgob*

CNback is a cloud native backup automation tool built with golang.

#### Features

- [x] schedule backups with cron-style parameter
- [ ] multi source database
  - [ ] mysql
  - [ ] mongodb
  - [ ] postgres
- [ ] multi storage destination
  - [ ] local backups retention
  - [ ] upload to S3 Object Storage (Minio, AWS, Google Cloud)
  - [ ] upload to gcloud storage
  - [ ] upload to SFTP
- [ ] multi notification platforms
  - [ ] emails
  - [ ] slack
- [ ] distributed as an Alpine Docker image

#### Install

...

#### Configure

Define a backup plan (yaml format) for each database you want to backup inside the `config` dir. 
The yaml file name is being used as the backup plan ID, no white spaces or special characters are allowed.

An example can be seen in the `config.example.yml`

#### Restore

...
