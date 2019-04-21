package loader

import (
	//query "git.sr.ht/~gabe/hod/lang"
	//sparql "git.sr.ht/~gabe/hod/lang/ast"
	//logpb "git.sr.ht/~gabe/hod/proto"
	"git.sr.ht/~gabe/hod/turtle"
	"github.com/dgraph-io/badger"
	//"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/pkg/profile"
	logrus "github.com/sirupsen/logrus"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"
	//"strings"
	"sync"
	"time"
)

// initializing the Hod DB

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, ForceColors: true})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
}

type HodDB struct {
	db *badger.DB
	//versionDB *versionmanager
	cfg *Config

	hashes map[turtle.URI]EntityKey
	uris   map[EntityKey]turtle.URI

	// TODO: serialize/deserialize
	// map graph name to namespaces (map[string]map[string]string)
	namespaces sync.Map
	graphs     map[string]struct{}
}

func MakeHodDB(cfg *Config) (*HodDB, error) {
	// handle profiling
	if cfg.Profile.EnableCpu {
		defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	} else if cfg.Profile.EnableMem {
		defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	} else if cfg.Profile.EnableBlock {
		defer profile.Start(profile.BlockProfile, profile.ProfilePath(".")).Stop()
	}
	// debug performance
	if cfg.Profile.EnableHttp {
		go func() {
			log.Info("Profile at localhost:", cfg.Profile.HttpPort)
			log.Info(http.ListenAndServe("localhost:"+cfg.Profile.HttpPort, nil))
		}()
	}

	/* open view database */
	dbdir := filepath.Join(cfg.Database.Path, "_db_")
	if err := os.MkdirAll(dbdir, 0700); err != nil {
		return nil, err
	}
	opts := badger.DefaultOptions
	opts.Dir = dbdir
	opts.ValueDir = dbdir
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	hod := &HodDB{
		db:     db,
		cfg:    cfg,
		hashes: make(map[turtle.URI]EntityKey),
		uris:   make(map[EntityKey]turtle.URI),
		graphs: make(map[string]struct{}),
	}

	//err = hod.buildVersionManager(cfg)
	//if err != nil {
	//	log.Fatal(err)
	//}

	// start GC on the database
	go func() {
		ticker := time.NewTicker(1 * time.Minute)
		defer ticker.Stop()
		for range ticker.C {
			log.Debug("running gc")
		againDb:
			err := db.RunValueLogGC(0.7)
			if err == nil {
				goto againDb
			}
		}
	}()

	numBuildings := len(cfg.Database.Buildings)
	processed := 0
	for graphname, graphfile := range cfg.Database.Buildings {
		s := time.Now()
		bundle := FileBundle{
			GraphName:     graphname,
			TTLFile:       graphfile,
			OntologyFiles: cfg.Database.Ontologies,
		}
		if err := hod.Load(bundle); err != nil {
			log.Error(errors.Wrapf(err, "Could not load file %s", graphname))
		}
		processtime := time.Since(s)
		processed += 1
		log.Infof("Loaded in %d/%d (%.2f%%) buildings from config file (%s took %s)", processed, numBuildings, 100*float64(processed)/float64(numBuildings), graphname, processtime)
	}

	return hod, nil
}
