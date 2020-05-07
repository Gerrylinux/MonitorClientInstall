package install

import (
	"fmt"
	"os/exec"
)

func Noderemove() {

	cmd := exec.Command("/bin/bash", "-c", `rm -rf /usr/local/src/node_exporter* && rm -rf /usr/local/node_exporter && rm -rf /usr/local/supervisord/conf.d/node_exporter.ini  && supervisorctl update`)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	ConsulDeRegister("node")
}

func Redisremove() {
	cmd := exec.Command("/bin/bash", "-c", `rm -rf /usr/local/src/redis_exporter* && rm -rf /usr/local/redis_exporter && rm -rf /usr/local/supervisord/conf.d/redis_exporter.ini && supervisorctl update`)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	ConsulDeRegister("redis")
}

func Kafkaremove() {
	cmd := exec.Command("/bin/bash", "-c", `rm -rf /usr/local/src/kafka_exporter* && rm -rf /usr/local/kafka_exporter && rm -rf  /usr/local/supervisord/conf.d/kafka_exporter.ini && supervisorctl update`)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	ConsulDeRegister("kafka")
}

func Mysqlremove() {
	cmd := exec.Command("/bin/bash", "-c", `rm -rf /usr/local/src/mysqld_exporter* && rm -rf /usr/local/mysqld_exporter && rm -rf /usr/local/supervisord/conf.d/mysqld_exporter.ini && supervisorctl update`)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	ConsulDeRegister("mysqld")
}