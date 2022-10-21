#!/bin/bash

# Variables
password="newrelic"

# Create user
sudo mysql -e "CREATE USER 'newrelic'@'localhost' IDENTIFIED BY '$password' WITH MAX_USER_CONNECTIONS 5;"

# Grant replication privileges
sudo mysql -e "GRANT REPLICATION CLIENT ON *.* TO 'newrelic'@'localhost' WITH MAX_USER_CONNECTIONS 5;"

# Grant select privileges
sudo mysql -e "GRANT SELECT ON *.* TO 'newrelic'@'localhost' WITH MAX_USER_CONNECTIONS 5;"

# Install the agents
curl -Ls https://download.newrelic.com/install/newrelic-cli/scripts/install.sh | \
  bash && sudo NEW_RELIC_API_KEY=$NEWRELIC_API_KEY NEW_RELIC_ACCOUNT_ID=$NEWRELIC_ACCOUNT_ID NEW_RELIC_REGION=EU /usr/local/bin/newrelic install

# Add MySQL logging directory to config
echo '
  - name: mysql.log
    file: /var/lib/mysql/general_logs.log
    attributes:
      logtype: mysql_log
' | \
  sudo tee -a /etc/newrelic-infra/logging.d/logging.yml
