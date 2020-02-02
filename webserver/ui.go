package webserver

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/edznux/wonderxss/api"
	"github.com/edznux/wonderxss/api/local"
	"github.com/edznux/wonderxss/config"
	"github.com/edznux/wonderxss/storage/models"
)

//UI is the interface/struct for the web server (excluding the APIs)
type UI struct {
	indexPath  string
	staticPath string
	api        api.API
}

// New return a new UI struct.
// It's already set up and have create a new *local* api.
func New() *UI {
	ui := UI{}
	ui.indexPath = "/index.html"
	ui.staticPath = "webserver/wonderxss/build"
	log.Println("Connecting to local API from ui")
	ui.api = local.New()
	return &ui
}

// LoggingMiddleware print all the http requests done by the application on a standard format
func (ui *UI) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Infof("%s - %s - \"%s %s %s\"-\"%s\"",
			r.RemoteAddr,
			strings.Split(r.Host, ":")[0],
			r.Method,
			r.RequestURI,
			r.Proto,
			r.UserAgent(),
		)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// HandleIndex differenciate request to be able to send the correct response
// If there is no subdomain, send the user interface (React app)
// Else, return the payload (if it exist, otherwise 404)
func (ui *UI) HandleIndex(w http.ResponseWriter, req *http.Request) {
	hostname := req.Host
	subdomain := strings.TrimSuffix(hostname, "."+config.Current.Domain)
	log.Debugln("req.URL.Path:", req.URL.Path)
	log.Debugln("hostname:", hostname)
	log.Debugln("Subdomain:", subdomain)
	content, err := ui.api.ServePayload(subdomain)

	// Index page, should return the UI
	if subdomain == hostname {
		log.Debug("Index page called, redirecting to UI")
		ui.ServeUI(w, req)
		return
	}
	if err == models.NoSuchItem {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No such payload"))
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Encountered an error :/"))
		return
	}
	w.Write([]byte(content))
}

// ServeUI render the React APP
// TODO: use bindata to embed it in the golang binary
func (ui *UI) ServeUI(w http.ResponseWriter, r *http.Request) {
	log.Debugln("Serving UI")
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// prepend the path with the path to the static directory
	path = filepath.Join(ui.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		fullIndexPath := filepath.Join(ui.staticPath, ui.indexPath)
		log.Info("Non-existing path, returning indexPath", fullIndexPath)
		http.ServeFile(w, r, fullIndexPath)
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(ui.staticPath)).ServeHTTP(w, r)
}
