[Unit]
Description=cloudsql
[Service]
ExecStart=/root/scripts/01_systemctl_cloudsql.sh
[Install]
WantedBy=multi-user.target

systemctl status cloudsql.service 
systemctl start cloudsql.service
systemctl enable cloudsql.service

