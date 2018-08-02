## Usage ##
### For native environment ###
Copy sysconfig.yaml.tpl as sysconfig.yaml and modify it.
```
cp sysconfig.yaml.tpl sysconfig.yaml
vim sysconfig.yaml
```
Sync the 3rd package
```
govendor sync
```

The binary file(eth-agent) will be generated under $GOPATH/bin
```
go install
```

Execute eth-agent
```
eth-agent --config sysconfig.yaml
```

### For docker environment ###
Copy docker/.env.default as docker/.env and modify it.
```
cp docker/.env.default docker/.env
vim docker/.env
```

Execute by docker-compose
```
cd docker 
docker-compose up -d
```

## TODO ##
### Issue 1 ###
When change/add account for mongodb , it has to remove original volume.
The side effect is the all data(e.g. block chain data) will be removed also.

Should try to reserve original data.