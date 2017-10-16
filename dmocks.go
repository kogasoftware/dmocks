package dmocks

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/kogasoftware/dmocks/configuration"
	"github.com/kogasoftware/dmocks/writer"
)

// Dmocks is struct.
type Dmocks struct {
	Port       int
	ConfigPath string
}

// NewDmocks returns Dmocks.
func NewDmocks() *Dmocks {
	return &Dmocks{}
}

// Run executes main process.
func (dmocks *Dmocks) Run() {
	dmocks.setFlag()
	config := configuration.NewConfiguration()
	config.LoadConfigFile(dmocks.ConfigPath)
	go writer.WriteTargetsLoop(config)

	// Listen port 3000
	fs := http.FileServer(http.Dir(config.RootPath))
	http.Handle("/", handleCORS(fs))
	fmt.Println(fmt.Sprintf("Listen port %d.", dmocks.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", dmocks.Port), nil)
}

func (dmocks *Dmocks) setFlag() {
	flag.Usage = defaultHelpMessage

	var port int
	flag.IntVar(&port, "p", 3000, "")
	flag.IntVar(&port, "port", 3000, "")

	var configPath string
	flag.StringVar(&configPath, "c", "config.yml", "")
	flag.StringVar(&configPath, "config", "config.yml", "")

	flag.Parse()

	dmocks.Port = port
	dmocks.ConfigPath = configPath
}

func handleCORS(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

func defaultHelpMessage() {
	message := `Usage: dmocks [options]

options:
  -p, --port
      set port (Default value is 3000)
  -c, --config
      set config file (Default value is 'config.yml')
  -h, --help
      show help message
`
	fmt.Fprintf(os.Stderr, message)
}
