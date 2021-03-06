
# Firehose-to-Logdna

This is a Nozzle to drain log from firehose in CloudFoundry to LogDNA.

# Usage

Currently the Nozzle only run as an app in Cloud Foundry.

1 - Download the latest release of Firehose-to-Logdna.
```
    $git clone https://github.com/LYHuang/firehose-to-syslog.git
    $cd firehose-to-syslog
```

2 - Utilize the CF cli to authenticate with your PCF instance.
```
    cf login -a https://api.[your cf system domain] -u [your id] --skip-ssl-validation
```

3 - Push firehose-to-logdna.
```
    cf push firehose-to-logdna --no-start
```

4 - Set environment variables with cf cli.
```
    cf set-env firehose-to-logdna API_ENDPOINT https://api.[your cf system domain]
    cf set-env firehose-to-logdna DOPPLER_ENDPOINT wss://doppler.[your cf system domain]:443
    cf set-env firehose-to-logdna INGESTION_KEY [your ingestion key in LogDNA]
    cf set-env firehose-to-logdna FIREHOSE_USER  [your doppler.firehose enabled user]
    cf set-env firehose-to-logdna FIREHOSE_PASSWORD  [your doppler.firehose enabled user password]
    cf set-env firehose-to-logdna FIREHOSE_CLIENT_ID  [your doppler.firehose enabled client id]
    cf set-env firehose-to-logdna FIREHOSE_CLIENT_SECRET  [your doppler.firehose enabled client secret]
```

5 - Push and run.
```
    cf push firehose-to-logdna
```
