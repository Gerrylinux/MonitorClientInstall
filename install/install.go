package install

import (
	"fmt"
	"os/exec"
)

func NodeInstall() error {
	node_exporter := "node_exporter-0.17.0.linux-amd64.tar.gz"
	node_ini := "node_exporter.ini"

	if !PathExists("/usr/bin/", "supervisorctl") {
		err := SupervisorInstall()
		if err != nil {
			return  err
		}
	}

	if !PathExists("/usr/local/src/", node_exporter) {
		Download(node_exporter, "/usr/local/src/")
	}
	if !PathExists("/usr/local/supervisord/conf.d/", node_ini) {
		Download(node_ini, "/usr/local/supervisord/conf.d/")
	}
	FileTar("/usr/local/src/node_exporter-0.17.0.linux-amd64","/usr/local/src/", "node_exporter")
	cmd := exec.Command("/bin/bash", "-c", `supervisorctl update`)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err,": supervisorctl start err",node_ini)
		return err
	}
	ConsulRegister("node",9100)
	return nil
}

func RedisInstall() error {
	node_exporter := "redis_exporter-v1.5.3.linux-amd64.tar.gz"
	node_ini := "redis_exporter.ini"

	if !PathExists("/usr/bin/", "supervisorctl") {
		err := SupervisorInstall()
		if err != nil {
			return  err
		}
	}

	if !PathExists("/usr/local/src/",  node_exporter) {
		Download(node_exporter, "/usr/local/src/")
	}
	if !PathExists("/usr/local/supervisord/conf.d/", node_ini) {
		Download(node_ini, "/usr/local/supervisord/conf.d/")
	}
	FileTar("/usr/local/src/redis_exporter-v1.5.3.linux-amd64","/usr/local/src/", "redis_exporter")
	_, IP := SystemEnv()
	ConfSed := "sed -i 's/127.0.0.1/" + IP +"/g' /usr/local/supervisord/conf.d/redis_exporter.ini && supervisorctl update"
	cmd := exec.Command("/bin/bash", "-c", ConfSed)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err,": supervisorctl start err",node_ini)
		return err
	}
	ConsulRegister("redis",9121)
	return nil
}

func MysqlInstall() error {
	node_exporter := "mysqld_exporter-0.12.1.linux-amd64.tar.gz"
	node_ini := "mysqld_exporter.ini"

	if !PathExists("/usr/bin/", "supervisorctl") {
		err := SupervisorInstall()
		if err != nil {
			return  err
		}
	}

	if !PathExists("/usr/local/src/",  node_exporter) {
		Download(node_exporter, "/usr/local/src/")
	}
	if !PathExists("/usr/local/supervisord/conf.d/", node_ini) {
		Download(node_ini, "/usr/local/supervisord/conf.d/")
	}
	FileTar("/usr/local/src/mysqld_exporter-0.12.1.linux-amd64","/usr/local/src/", "mysqld_exporter")
	cmd := exec.Command("/bin/bash", "-c", `supervisorctl update`)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err,": supervisorctl start err",node_ini)
		return err
	}
	ConsulRegister("mysqld",9104)
	return nil
}

func KafkaInstall() error {
	node_exporter := "kafka_exporter-1.2.0.linux-amd64.tar.gz"
	node_ini := "kafka_exporter.ini"

	if !PathExists("/usr/bin/", "supervisorctl") {
		err := SupervisorInstall()
		if err != nil {
			return  err
		}
	}

	if !PathExists("/usr/local/src/",  node_exporter) {
		Download(node_exporter, "/usr/local/src/")
	}
	if !PathExists("/usr/local/supervisord/conf.d/", node_ini) {
		Download(node_ini, "/usr/local/supervisord/conf.d/")
	}
	FileTar("/usr/local/src/kafka_exporter-1.2.0.linux-amd64","/usr/local/src/", "kafka_exporter")
	_, IP := SystemEnv()
	ConfSed := "sed -i 's/127.0.0.1/" + IP +"/g' /usr/local/supervisord/conf.d/kafka_exporter.ini && supervisorctl update"
	cmd := exec.Command("/bin/bash", "-c", ConfSed)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err,": supervisorctl start err",node_ini)
		return err
	}
	ConsulRegister("kafka",9308)
	return nil
}

func SupervisorInstall() error {
	cmd := exec.Command("/bin/bash", "-c", `yum -y install epel* && yum -y install supervisor&& mkdir -p /usr/local/supervisord/conf.d/ /data/wwwlogs/supervisord && rm -f /usr/lib/systemd/system/supervisord.service && wget -P /usr/local/supervisord/ http://172.17.126.214/supervisor/supervisord.ini && wget -P /usr/lib/systemd/system/ http://172.17.126.214/supervisor/supervisord.service && systemctl enable supervisord && systemctl restart supervisord `)
	_, err := cmd.Output()
	if err != nil {
		fmt.Println(err, ":install supervisor server defeated")
		return err
	}
	return nil
}