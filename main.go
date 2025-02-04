// Initialization of the application.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"time"

	rice "github.com/GeertJohan/go.rice"
	badger "github.com/dgraph-io/badger"
	"github.com/gorilla/mux"
	"github.com/logrusorgru/aurora"
)

func main() {

	// License
	fmt.Println("")
	fmt.Println(aurora.Bold("Permaweb Host :"), "Upload a Git repo to IPNS.")
	fmt.Println("Copyright © 2019 Nato Boram, Permaweb")
	fmt.Println(aurora.Bold("Contact :"), aurora.Blue("https://github.com/Permaweb/Host"))
	fmt.Println("")

	// User
	path, err := initUser()
	if err != nil {
		return
	}
	dirHome = path

	// Forward Compatibility
	err = initCompatibility()
	if err != nil {
		return
	}

	// IPFS
	err = initIPFS()
	if err != nil {
		return
	}

	// Git
	err = initGit()
	if err != nil {
		return
	}

	// Badger
	db, err := initBager()
	if err != nil {
		return
	}
	defer db.Close()

	// Refresh repos
	go func() {
		for {
			onAllRepos(db)
			time.Sleep(24 * time.Hour)
		}
	}()

	// Router
	go initMux(db)

	// Listen to CTRL+C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {

			// Close database
			err = db.Close()
			if err != nil {
				fmt.Println("Couldn't save the database.")
				log.Fatalln(err.Error())
			}

			os.Exit(0)
		}
	}()

	// Wait
	<-make(chan struct{})
}

func initGit() (err error) {

	// Git Directory
	err = os.MkdirAll(dirHome+dirGit, permPrivateDirectory)
	if err != nil {
		fmt.Println("Couldn't create the git directory.")
		fmt.Println(err.Error())
		return
	}

	// Check for Git
	path, err := exec.LookPath("git")
	if err != nil {
		fmt.Println("Git is not installed.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(aurora.Bold("Git :"), aurora.Blue(path))

	fmt.Println("")
	return
}

func initIPFS() (err error) {

	// Check for IPFS
	path, err := exec.LookPath("ipfs")
	if err != nil {
		fmt.Println("IPFS is not installed.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(aurora.Bold("IPFS :"), aurora.Blue(path))

	// Enable sharding
	exec.Command("ipfs", "config", "--json", "Experimental.ShardingEnabled", "true").Run()

	// Check for IPFS Cluster Service
	path, err = exec.LookPath("ipfs-cluster-service")
	if err != nil {
		fmt.Println("IPFS Cluster Service is not installed.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(aurora.Bold("IPFS Cluster Service :"), aurora.Blue(path))

	// Check for IPFS Cluster Control
	path, err = exec.LookPath("ipfs-cluster-ctl")
	if err != nil {
		fmt.Println("IPFS Cluster Control is not installed.")
		fmt.Println(err.Error())
		return
	}
	fmt.Println(aurora.Bold("IPFS Cluster Control :"), aurora.Blue(path))

	// Connect to Swarm
	initSwarm()

	fmt.Println("")
	return
}

func initSwarm() {
	for _, pg := range PublicGateways {
		ipfsSwarmConnect(pg)
	}
}

func initBager() (db *badger.DB, err error) {

	// Badger Directory
	err = os.MkdirAll(dirHome+dirBadger, permPrivateDirectory)
	if err != nil {
		fmt.Println("Couldn't create the badger directory.")
		fmt.Println(err.Error())
		return
	}

	// Options
	options := badger.DefaultOptions(dirHome + dirBadger)

	db, err = badger.Open(options)
	if err != nil {
		fmt.Println("Couldn't open a Badger Database.")
		fmt.Println(err.Error())
	}

	fmt.Println()
	return db, err
}

func initUser() (path string, err error) {

	usr, err := user.Current()
	if err != nil {
		fmt.Println("Couldn't get the current user.")
		fmt.Println(err.Error())
	}

	return usr.HomeDir, err
}

func initMux(db *badger.DB) {
	r := mux.NewRouter()

	// API
	api := r.StrictSlash(true).PathPrefix("/api").Subrouter()

	// Repos
	repos := api.PathPrefix("/repos").Subrouter()
	repos.Methods("GET").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) { reposGetHandler(db, w, r) })
	repos.Methods("POST").Path("/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) { reposPostHandler(db, w, r) })
	repos.Methods("GET").Path("/{link:.*}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) { repoGetHandler(db, w, r) })
	repos.Methods("DELETE").Path("/{link:.*}").HandlerFunc(func(w http.ResponseWriter, r *http.Request) { repoDeleteHandler(db, w, r) })

	// Web Server
	r.PathPrefix("/").Handler(http.FileServer(rice.MustFindBox("web").HTTPBox()))

	fmt.Println("Web server started at", aurora.Blue("http://localhost:62458/"))
	log.Fatal(http.ListenAndServe(":62458", r))
}

func initCompatibility() (err error) {

	// Move config directory
	dirOldConfig := dirHome + "/.config/gi"
	if _, err := os.Stat(dirOldConfig); !os.IsNotExist(err) {
		_, err = mv(dirOldConfig, dirHome+dirConfig)
		if err != nil {
			fmt.Println("Couldn't move old config to new directory")
			fmt.Println(err.Error())
			return err
		}
	}

	// Convert Badger to new version, when applicable.
	// ...

	return
}
