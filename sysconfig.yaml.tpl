# use YAML format
# change these to your prefer setting
# use YAML format
# change these to your prefer setting
# log_level, represent the log level:
# PanicLevel Level = 0
# FatalLevel = 1
# ErrorLevel = 2
# WarnLevel = 3
# InfoLevel = 4
# DebugLevel = 5

eth_proxy:
    protocol: <ETH_AGENT_PROTOCOL>
    domain: <ETH_AGENT_DOMAIN>
    port: <ETH_AGENT_PORT>
    log_file: <ETH_AGENT_LOG_CT>
    log_level: <ETH_AGENT_LOG_LEVEL>
#    product_bin: <ETH_AGENT_BIN>

eth:
    domain: <ETH_DOMAIN>
    port: <ETH_PORT>

sso:
    domain: <SSO_DOMAIN>
    port: <SSO_PORT>

redis:
    domain: <REDIS_DOMAIN>
    port: <REDIS_PORT>
    password: <REDIS_PASSWORD>

mongo:
    domain: <MONGODB_DOMAIN>
    port: <MONGODB_PORT>
    #root_username: <MONGO_INITDB_ROOT_USERNAME>
    #root_password: <MONGO_INITDB_ROOT_PASSWORD> 
    agent_db: <MONGO_AGENT_DB>
    agent_user: <MONGO_USERNAME>
    agent_pwd: <MONGO_PASSWORD>
    agent_readonly_user: <MONGO_AGENT_READONLY_USERNAME>
    agent_readonly_pwd: <MONGO_AGENT_READONLY_PASSWORD>