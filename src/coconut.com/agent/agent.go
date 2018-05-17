package agent

import (
	"github.com/spf13/cobra"
	"runtime"
	"os"
	"github.com/jackc/pgx"
	"log"
	"coconut.com/config/pgconf"
	"github.com/gorilla/mux"
	"net/http"
	h "coconut.com/handlers"
	"coconut.com/db"
	"coconut.com/config"
	"github.com/gorilla/handlers"
	"fmt"
)

var Cmd = &cobra.Command{
	Use:   "deploy",
	Short: "deploy sever",
	Run:   command,
}

func init() {
	Cmd.PersistentFlags().String("port", "4000", "default http port 4000")
}

var faviconHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/favicon.ico")
})

func command(cmd *cobra.Command, args []string) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// database
	pgconf.SetFlags(cmd)
	conn, err := newPgPool(cmd)
	if err != nil {
		log.Fatalf("cannot connect database %v", err)
		os.Exit(1)
	}
	db.Setup(conn)

	// config
	config.ParseFlags(cmd)

	// load build options
	projects, err := db.LoadBuildOptions()
	if err != nil {
		log.Fatalf("cannot connect database %v", err)
		os.Exit(1)
	}
	config.BuildOptions = projects

	// Here we are instantiating the gorilla/mux router
	r := mux.NewRouter()

	// On the default page we will simply serve our static index page.
	r.Handle("/", http.FileServer(http.Dir("./views/")))
	// We will setup our server so we can serve static assest like images, css from the /static/{file} route
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	r.PathPrefix("/payloads/").Handler(http.StripPrefix("/payloads/", http.FileServer(http.Dir("./payloads/"))))

	r.Handle("/favicon.icon", faviconHandler).Methods("GET")

	r.Handle("/list", h.PayloadsHandler).Methods("GET")
	r.Handle("/upload", h.UploadHandler).Methods("POST")
	r.Handle("/event_handler", h.EventHandler).Methods("POST")
	r.Handle("/build_configs/{key}", h.BuildConfigHandler).Methods("GET")
	r.Handle("/build", h.BuildHandler).Methods("POST")
	r.Handle("/build/remove", h.RemoveBuildHandler).Methods("POST")

	// Our application will run on port 8443. Here we declare the port and pass in our router.
	fmt.Printf("Start server listening on port %v\n", config.HttpPort)
	http.ListenAndServe(":" + config.HttpPort, handlers.LoggingHandler(os.Stdout, r))
	//http.ListenAndServeTLS(":" + config.HttpPort, "cert.pem", "key.pem", handlers.LoggingHandler(os.Stdout, r))
}

func newPgPool(cmd *cobra.Command) (pg *pgx.ConnPool, err error) {
	cfg := pgconf.Config(cmd)
	cfg.AfterConnect = func(conn *pgx.Conn) error {
		db.PrepareStmt(conn)
		return nil
	}
	pg, err = pgx.NewConnPool(cfg)
	if err != nil {
		return nil, err
	}
	return
}