// flags defines flags used by primary, backup, and client.
package flags

import "flag"

var Id = flag.Int("id", 0, "ID of the server, backup, or client.")
var ConfigPath = flag.String("config_path", "", "Path to the config file.")
var Listen = flag.Int("listen", 8080, "Port to listen on")
var Backends = flag.String("backend", "", "comma seperated string of backend replicas")

func init() {
	flag.Parse()
}
