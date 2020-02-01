package webserver

import (
	"fmt"
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

type UI struct {
	indexPath  string
	staticPath string
	api        api.API
}

func New() *UI {
	ui := UI{}
	ui.indexPath = "/index.html"
	ui.staticPath = "webserver/wonderxss/build"
	log.Println("Connecting to local API from ui")
	ui.api = local.New()
	return &ui
}

func (ui *UI) HandleIndex(w http.ResponseWriter, req *http.Request) {
	hostname := req.Host
	subdomain := strings.TrimSuffix(hostname, "."+config.Current.Domain)
	log.Println("req.URL.Path:", req.URL.Path)
	log.Println("hostname:", hostname)
	log.Println("Subdomain:", subdomain)
	content, err := ui.api.ServePayload(subdomain)

	// Index page, should return the UI
	if subdomain == hostname {
		fmt.Println("Index page called, redirecting to UI")
		ui.ServeUI(w, req)
		return
	}
	if err == models.NoSuchItem {
		w.Write([]byte("No such payload"))
		return
	}
	if err != nil {
		w.Write([]byte("Encountered an error :/"))
		return
	}
	w.Write([]byte(content))
}

func (ui *UI) ServeUI(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Serving UI")
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
	fmt.Println("Path:", path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		fullIndexPath := filepath.Join(ui.staticPath, ui.indexPath)
		fmt.Println("Non-existing path, returning indexPath", fullIndexPath)
		http.ServeFile(w, r, fullIndexPath)
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(ui.staticPath)).ServeHTTP(w, r)
}
