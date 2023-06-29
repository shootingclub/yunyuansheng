package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleInterceptor(httpRootFunc))
	http.HandleFunc("/health", handleInterceptor(healthFunc))
	err := http.ListenAndServe(":8999", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthFunc(w http.ResponseWriter, r *http.Request) {
	HealthCode := "200"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(HealthCode))
	return
}

func httpRootFunc(w http.ResponseWriter, r *http.Request) {
	//request header写入response header
	//if len(r.Header) > 0 {
	//	for k, v := range r.Header {
	//		log.Printf("%s=%s", k, v[0])
	//		w.Header().Set(k, v[0])
	//	}
	//}
	//
	//// 设置环境值的值
	//os.Setenv("VERSION", "GOLANG YUN YUAN SHENG")
	//// 获取环境变量"VERSION"
	//name := os.Getenv("VERSION")
	//log.Printf("VERSION ENV : %s \n", name)
	//
	////获取Client IP，并且打印出来
	//ip, _, err := net.SplitHostPort(r.RemoteAddr)
	//if err != nil {
	//	fmt.Println("err:", err)
	//}
	//
	//if net.ParseIP(ip) != nil {
	//	log.Printf("ip : %s\n", ip)
	//}
	//log.Printf("http Status Code : %d \n", http.StatusOK)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server Run Success "))

	return
}

func handleInterceptor(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		//request header写入response header
		if len(r.Header) > 0 {
			for k, v := range r.Header {
				log.Printf("%s=%s", k, v[0])
				w.Header().Set(k, v[0])
			}
		}

		// 设置环境值的值
		os.Setenv("VERSION", "GOLANG YUN YUAN SHENG")
		// 获取环境变量"VERSION"
		name := os.Getenv("VERSION")
		log.Printf("VERSION ENV : %s \n", name)

		//获取Client IP，并且打印出来
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			fmt.Println("err:", err)
		}

		if net.ParseIP(ip) != nil {
			log.Printf("ip : %s\n", ip)
		}
		log.Printf("http Status Code : %d \n", http.StatusOK)

		h(w, r)
	}

}
