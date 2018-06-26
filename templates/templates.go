package templates

import (
	"fmt"
	"html/template"
	"log"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/blastbeatsandcode/blastbeatsandcode-website/utils"
	"github.com/fsnotify/fsnotify"
)

// Reloader holds the required templates, a watcher to watch for changes, and a mutex
// Mutex will lock our map concurrently
type Reloader struct {
	Templates map[string]*template.Template

	*fsnotify.Watcher
	*sync.RWMutex
}

// Define strings for templates
var (
	templatePath = filepath.Join("templates/")
	templateExt  = ".gohtml"
)

/* GetTemplates returns a map of templates to their respective parts */
func GetTemplates() Reloader {
	str := "templates/"
	r := New(str)

	// Map templates to parent templates
	r.Templates = map[string]*template.Template{
		"index":    parsedTemplates("index"),
		"projects": parsedTemplates("projects"),
		"blog":     parsedTemplates("blog"),
		"contact":  parsedTemplates("contact"),
		"login":    parsedTemplates("login"),
	}

	return *r
}

/* Get retrieves a template with the given name from the internal map */
func (r *Reloader) Get(name string) *template.Template {
	r.RLock()         // Begin mutex lock
	defer r.RUnlock() // Defer mutex lock until we are finished

	// See if we have name in templates map
	if t, ok := r.Templates[name]; ok {
		return t
	}

	return nil
}

/* Watch calls a goroutine that waits for fsnotify events and will hot-swap the modified template */
func (r *Reloader) Watch() {
	go func() {
		for {
			select {
			case evt := <-r.Watcher.Events: // Detects if there are any file events
				fmt.Println("WATCHER EVENT DETECTED")
				fmt.Printf("File: %s Event: %s. Hot reloading.\n", evt.Name, evt.String())

				// Do reload
				if err := r.reload(evt.Name); err != nil {
					fmt.Print(err)
				}
			case err := <-r.Watcher.Errors:
				fmt.Println("Watcher error")
				fmt.Println(err)
			default:
				// Sleep the routines because they execute approximately a bajillion times a second if you dont!
				time.Sleep(250 * time.Millisecond)
			}
		}
	}()
}

/* New creates a Reloader so files can be watched for changes concurrently. */
func New(dirs ...string) *Reloader {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}

	for _, path := range dirs {
		watcher.Add(path)
		fmt.Println("Add path: ")
		fmt.Println(path)
	}

	return &Reloader{
		Watcher: watcher,
		RWMutex: &sync.RWMutex{},
	}
}

/* doReload gets called by reload() and reparses given template. */
func (r *Reloader) doReload(name string) error {
	if len(name) >= len(templateExt) && name[len(name)-len(templateExt):] == templateExt {
		tmpl := parsedTemplates(name)

		// Gather what would be the key in our template map.
		// 'name' is in the format: "path/indentifier.extension",
		// so we trim the 'path/' and the '.extension' to get
		// name (minus new extension) used inside of our map.
		key := name[:len(name)-len(templateExt)]

		r.Lock()
		r.Templates[key] = tmpl
		r.Unlock()

		return nil
	}
	return fmt.Errorf("unable to reload file %s", name)
}

/*	reload takes the initial reloader and does some checking to see if the file is a base template file
	then parses the files in doReload() */
func (r *Reloader) reload(name string) error {
	if strings.HasPrefix(name, "templates/") {
		name = strings.Replace(name, "templates/", "", -1)
	}

	// Check for base templates to be reloaded so all template paths are reparsed to update across the board
	if name == "base.gohtml" || name == "nav.gohtml" || name == "social.gohtml" {
		var err error

		// Reload all with the updated "base".
		for names := range r.Templates {
			names += ".gohtml"

			if err = r.doReload(names); err != nil {
				log.Println(err)
			}
		}
		return err
	}
	return r.doReload(name)
}

/* parsedTemplates takes in a fileName and parses all other base files, returns parsed template.
This is how we will serve templates in a modular fashion. */
func parsedTemplates(fileName string) *template.Template {
	// See if extension was passed, if not add it
	if !strings.Contains(fileName, ".gohtml") {
		fileName += ".gohtml"
	}

	pref := "templates/"

	// In this return statement, we also assign any functions we want to be able to pass to the templates
	// This will help us to unescape HTML given from sources other than the templates that we want parsed
	return template.Must(template.New("").Funcs(template.FuncMap{"noescape": utils.Noescape}).
		ParseFiles(pref+fileName, pref+"base.gohtml", pref+"nav.gohtml", pref+"social.gohtml"))
}
