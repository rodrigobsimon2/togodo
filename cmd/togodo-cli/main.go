package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/rodrigobsimon/togodo/internal/infra/db"
	"github.com/rodrigobsimon/togodo/internal/service"
	"github.com/rodrigobsimon/togodo/internal/usecase"
)

func main() {
	godotenv.Load()
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	pi := db.PersistentImplementationMySQL{
		Host: os.Getenv("DB_HOST"),
		Port: port,
		User: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB_DATABASE"),
	}
	
	list := flag.Bool("list", false, "on/off flag, determines if tasks will be listed")
	create := flag.Bool("create", false, "on/off flag, determines if a task will be created")
	description := flag.String("description", "default", "task description")
	priority := flag.Int("priority", 0, "task priority (3 - IMPORTANT URGENT, 2 - IMPORTANT NOT URGENT, 1 - NOT IMPORTANT URGENT, 0 - NOT IMPORTANT NO URGENT)")
	
	flag.Parse()

	if *create {
		t := usecase.CreateTask(*description, *priority)
		service.SaveTask(t, pi)
	}

	if *list {
		ts := service.ListTask(pi)
		fmt.Println(ts)
	}
}