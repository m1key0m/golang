[Unit]
Description=cloudsql
[Service]
ExecStart=/root/scripts/01_systemctl_cloudsql.sh
[Install]
WantedBy=multi-user.target

systemctl status cloudsql.service
systemctl start cloudsql.service
systemctl enable cloudsql.service

package main

import "fmt"

var x = 4

func main() {
	fmt.Printf("Test\n")
	f()

}

func f() {
	for i := 0; i < 10; i++ {
		fmt.Printf("hi ")
	}
}
