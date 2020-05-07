package install

import (
	"fmt"
	"log"
	consulapi "github.com/hashicorp/consul/api"
	"net"
	"os"
)
//增加相关服务
func ConsulRegister(Name string, Port int) {
	config := consulapi.DefaultConfig()
	config.Address = ConsulUrl
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	//创建一个新服务。
	hostname, IP := SystemEnv()
	ID := Name + "-" + IP + "-"+ hostname

	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = ID
	registration.Name = Name + "_exporter"
	registration.Port = Port
	registration.Tags = []string{"prod","prome","node"}
	registration.Address = hostname

	//增加check。
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/metrics")
	//设置超时 5s。
	check.Timeout = "30s"
	//设置间隔 5s。
	check.Interval = "30s"
	//注册check服务。
	registration.Check = check
	log.Println("get check.HTTP:", check)
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatal("register server error : ", err)
	}
}

//删除相关服务
func ConsulDeRegister(Name string)  {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = ConsulUrl
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	hostname, IP := SystemEnv()
	ID := Name + "-" + IP + "-"+ hostname
	client.Agent().ServiceDeregister(ID)
}

func SystemEnv() (Hostname, IP string) {
	var systemIP [5]string
	name, _ := os.Hostname()
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
	}

	for _, address := range addrs {
		i := 0
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				systemIP[i] = ipnet.IP.String()
				break
			}
		}
	}
return name, systemIP[0]
}