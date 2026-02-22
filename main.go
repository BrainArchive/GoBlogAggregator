package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/brainarchive/goblogaggregator/internal/config"
	"github.com/brainarchive/goblogaggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.ReadConfig()
	sessionState := &state{
		ConfigPtr: &cfg,
	}
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("list", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", middlewareLoggedIn(handlerCreateFeed))
	cmds.register("feeds", handlerListFeed)
	cmds.register("follow", middlewareLoggedIn(handlerFollow))
	cmds.register("following", middlewareLoggedIn(handlerFollowing))
	cmds.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	cmds.register("browse", handlerBrowse)

	// creates a connection to the database.
	dbURL := cfg.DBURL
	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	sessionState.db = dbQueries

	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	cmdName := os.Args[1]
	cmdArgs := os.Args[2:]
	err = cmds.run(sessionState, command{name: cmdName, args: cmdArgs})
	if err != nil {
		log.Fatalf("Command error: %v\n", err)
	}

}

type state struct {
	db        *database.Queries
	ConfigPtr *config.Config
}
