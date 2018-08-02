#!/bin/bash

# According to Dockerfile
source .env
cd $WORKDIR_CT

CONFIG=sysconfig.yaml

update_config() {
    local config_tpl=sysconfig.yaml.tpl
    [ ! -f "$config_tpl" ] && echo "$config_tpl is not existed." && exit 256
    
    cp $config_tpl $CONFIG

    sed -i "s%<ETH_AGENT_PROTOCOL>%$ETH_AGENT_PROTOCOL%g" $CONFIG
    sed -i "s%<ETH_AGENT_DOMAIN>%$ETH_AGENT_DOMAIN%g" $CONFIG
    sed -i "s%<ETH_AGENT_PORT>%$ETH_AGENT_PORT%g" $CONFIG
    sed -i "s%<ETH_AGENT_LOG_CT>%$ETH_AGENT_LOG_CT%g" $CONFIG
    sed -i "s%<ETH_AGENT_LOG_LEVEL>%$ETH_AGENT_LOG_LEVEL%g" $CONFIG
    sed -i "s%<ETH_AGENT_BIN>%$ETH_AGENT_BIN%g" $CONFIG
    sed -i "s%<ETH_DOMAIN>%$ETH_DOMAIN%g" $CONFIG
    sed -i "s%<ETH_PORT>%$ETH_PORT%g" $CONFIG
    sed -i "s%<SSO_DOMAIN>%$SSO_DOMAIN%g" $CONFIG
    sed -i "s%<SSO_PORT>%$SSO_PORT%g" $CONFIG
    sed -i "s%<REDIS_DOMAIN>%$REDIS_DOMAIN%g" $CONFIG
    sed -i "s%<REDIS_PORT>%$REDIS_PORT%g" $CONFIG
    sed -i "s%<REDIS_PASSWORD>%$REDIS_PASSWORD%g" $CONFIG
}

main() {
    update_config
    eth-agent --config $CONFIG
}

main "$@"