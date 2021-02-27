# waseda_covit19_docs_backend

## Makefile
### init
**Please docker installed**
```
$ make init
```

### run app 
```
$ make server-run
```

### stop app 
```
$ make server-run
```

### run web server
```
$ make run
```

### run db
```
$ make db-run
```

### execute sql
**db is running**
```
$ make create-table sqlname=<sql file name>.sql"
```

## install google drive api token
1. Open Go Quickstart of Google Drive API v3
    + [Google Drive API](https://developers.google.com/drive/api/v3/quickstart/go)
2. Click Enable the Drive API button
    + Config
    + QuickStart → Desktop app → Create!
3. credentials.json will be downloaded
4. Exeute quickstart.go
    + Changed drive scope (L79)
```golang：quickstart.go
	
    // Change from drive.DriveMetadataReadonlyScope to drive.DriveScope
	config, err := google.ConfigFromJSON(b, drive.DriveMetadataReadonlyScope)　// ← here!
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

```
5. OAuth2 URL is output to the console log
6. Access the target OAuth URL in your browser
7. Login to the target Google Accunt
    + If an error is presented
    + Forced access!
8. Paste the access code displayed in the browser into the log
9. Get google drive api token!
10. Please set base64 encode token to env
```
$ echo GOOGLE_DRIVE_API_TOKEN=$(cat ./app/cmd/token.json | base64) >> .env
$ echo GOOGLE_DRIVE_API_CLIENT=$(cat ./app/cmd/client.json | base64) >> .env 
```