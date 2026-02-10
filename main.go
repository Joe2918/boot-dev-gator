package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/bootdotdev/boot-dev-gator/internal/config"
	"github.com/bootdotdev/boot-dev-gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	programState := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerListUsers)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	commandName := args[1]
	commandArgs := args[2:]

	err = cmds.run(programState, command{Name: commandName, Args: commandArgs})
	if err != nil {
		log.Fatal(err)
	}
}
