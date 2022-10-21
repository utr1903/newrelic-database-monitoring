#!/bin/bash

sudo systemctl stop mysql

echo Y | sudo apt-get purge mysql-server mysql-client mysql-common mysql-server-core-* mysql-client-core-*

sudo rm -rf /etc/mysql /var/lib/mysql

echo Y | sudo apt autoremove

echo Y | sudo apt autoclean
