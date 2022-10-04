package main

import (
	"bufio"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"local-p/internal/crud"
	"local-p/internal/database"
	"local-p/internal/instance"
	"local-p/internal/project"
	"os"
	"time"
)

// Get time now
func GetTimeNow() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
func main() {

	fmt.Println("Start: ", GetTimeNow())

	// Lista os projetos
	project.List()

	project := "totemic-chalice-364501"
	instanceName := "my-instance"
	databaseName := "mysql-db"

	// Cria a instancia
	fmt.Println("Create instance: ", GetTimeNow())
	instance.Create(project)
	for {
		i := instance.Get(project, instanceName)
		if i.State == "RUNNABLE" {
			break
		}
		time.Sleep(5 * time.Second)
	}

	// Libera conexão
	time.Sleep(20 * time.Second)
	fmt.Println("Allow Network: ", GetTimeNow())
	instance.Update(project, instanceName)

	// Obtem as informações da instancia
	i := instance.Get(project, instanceName)

	// Cria um banco de dados

	time.Sleep(20 * time.Second)
	fmt.Println("Create database: ", GetTimeNow())
	database.Create(project, instanceName)

	// Obtem as informações do banco de dados
	database.Get(project, instanceName, databaseName)

	// Add um usuário
	database.AddUser(project, instanceName)

	//Alimenta o banco de dados
	time.Sleep(20 * time.Second)
	fmt.Println("Insert data: ", GetTimeNow())
	fmt.Println("IP:", i.IpAddresses[0].IpAddress)
	crud.Create(i.IpAddresses[0].IpAddress)

	// Deleta a instancia
	fmt.Print("Press 'Enter' to continue...")
	fmt.Println("Delete instance: ", GetTimeNow())
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	instance.Delete(project, instanceName)

	fmt.Println("End: ", GetTimeNow())
}
