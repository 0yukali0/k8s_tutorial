package main

import (
  //"gorm.io/driver/postgres"
  //"gorm.io/gorm"
  "flag"
  "fmt"
)

func main() {
	host := flag.String("host", "127.0.0.1", "Your postgres ip")
	user := flag.String("user", "admin", "Your id")
	pw := flag.String("password", "password", "your password")
	db := flag.String("dbname", "defualt", "dbname")
	port := flag.Int("port", 5432, "port")
	flag.Parse()
	conn_str := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", *host, *user, *pw, *db, *port)
	fmt.Println(conn_str)
}
