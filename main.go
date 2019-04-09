package main

import (
	//	"context"
	//	"fmt"
	"git.sr.ht/~gabe/hod/hod"
	//	"github.com/chzyer/readline"
	"github.com/pkg/errors"
	//	"io"
	"log"
	//	"strings"
	//	"time"
)

var debug = false

func main() {

	cfg, err := hod.ReadConfig("hodconfig.yml")
	if err != nil {
		log.Fatal(errors.Wrap(err, "Could not load config file"))
	}
	_ = cfg

	L, err := hod.NewLog(cfg)
	if err != nil {
		log.Fatal(errors.Wrap(err, "open log"))
	}
	defer L.Close()
	//if err := L.ServeGRPC(); err != nil {
	//	log.Fatal(errors.Wrap(err, "grpc"))
	//}
	//version, err := L.LoadFile("test", "BrickFrame.ttl", "bf")
	//log.Println("V>", version)
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "load brickframe"))
	//}
	//version, err = L.LoadFile("test", "Brick.ttl", "b")
	//log.Println("V>", version)
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "load brick"))
	//}
	//version, err = L.LoadFile("test", "example.ttl", "ex")
	//log.Println("V>", version)
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "load example"))
	//}
	//version, err = L.LoadFile("test", "Brick.ttl", "b")
	//log.Println("V>", version)
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "load brick"))
	//}
	//version, err = L.LoadFile("test", "example.ttl", "ex")
	//log.Println("V>", version)
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "load example"))
	//}

	//q := "SELECT ?vav ?room FROM test WHERE { ?vav rdf:type brick:VAV . ?room rdf:type brick:Room . ?zone rdf:type brick:HVAC_Zone . ?vav bf:feeds+ ?zone . ?room bf:isPartOf ?zone };"

	//q := "SELECT ?x ?y FROM * WHERE { ?r rdf:type brick:Room . ?x ?y ?r };"
	//version, err := L.LoadFile("soda", "BrickFrame.ttl", "brickframe")
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "load brickframe"))
	//}
	//version, err = L.LoadFile("soda", "Brick.ttl", "brick")
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "load brick"))
	//}
	//version, err = L.LoadFile("soda", "berkeley.ttl", "berkeley")
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "load berkeley"))
	//}
	//_ = version

	//q := "SELECT ?x FROM soda WHERE { ?x rdf:type brick:Room };"

	//cur, err := L.createCursor("test", 0, version)
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "create cursor"))
	//}
	//_ = cur
	////}
	//key := cur.ContextualizeURI(&logpb.URI{
	//	Namespace: "https://brickschema.org/schema/1.0.3/Brick",
	//	Value:     "Zone_Temperature_Sensor",
	//})
	//log.Println(key)
	//log.Println(S(key))
	//entity, err := cur.getEntity(key)
	//if err != nil {
	//	log.Fatal(errors.Wrap(err, "get ent"))
	//}
	//L.Dump(entity)

	//go func() {
	log.Fatal(L.ServeGRPC())
	//}()

	//	if x, e := L.ParseQuery("SELECT ?x FROM movie WHERE { ?x rdf:type ?y };", time.Now().UnixNano()); e != nil {
	//		log.Fatal(e)
	//	} else {
	//		log.Println(x)
	//	}
	//	//if res, err := L.Select(context.Background(), x); err != nil {
	//	//	log.Fatal(e)
	//	//} else {
	//	//	log.Println(x)
	//	//}
	//
	//	prompt, err := readline.NewEx(&readline.Config{
	//		Prompt:            "\033[32m»\033[0m ",
	//		HistoryFile:       "/tmp/readline.tmp",
	//		EOFPrompt:         "exit",
	//		HistorySearchFold: true,
	//	})
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer prompt.Close()
	//	//	prompt.HistoryDisable()
	//
	//	var buildup = ""
	//	ts := time.Now().UnixNano()
	//	for {
	//		line, err := prompt.Readline()
	//		if err == readline.ErrInterrupt {
	//			if len(line) == 0 {
	//				break
	//			} else {
	//				continue
	//			}
	//		} else if err == io.EOF {
	//			break
	//		}
	//
	//		buildup += strings.TrimSpace(line)
	//		if strings.HasSuffix(buildup, ";") {
	//			prompt.SetPrompt("\033[32m»\033[0m ")
	//			query := buildup
	//			buildup = ""
	//			if err = prompt.SaveHistory(query); err != nil {
	//				log.Fatal(err)
	//			}
	//
	//			selectquery, err := L.ParseQuery(query, ts)
	//			if err != nil {
	//				log.Println(err)
	//				continue
	//			}
	//			res, err := L.Select(context.Background(), selectquery)
	//			if err != nil {
	//				log.Println(err)
	//				continue
	//			}
	//			fmt.Fprintf(prompt, "rows %d\n", len(res.Rows))
	//
	//			prompt.SetPrompt("\033[32m»\033[0m ")
	//		} else if len(buildup) > 0 {
	//			prompt.SetPrompt("...  ")
	//		}
	//	}
	//
	//	selectquery, err := L.ParseQuery(q, ts)
	//	log.Println(selectquery)
	//
	//	_, err = L.Select(context.Background(), selectquery)
	//	if err != nil {
	//		log.Println(errors.Wrap(err, "select q1"))
	//	} else {
	//		log.Println("successful q1")
	//	}
	//	//cur.dumpResponse(resp)
	//
	//	_, err = L.Select(context.Background(), selectquery)
	//	if err != nil {
	//		log.Println(errors.Wrap(err, "select q2"))
	//	} else {
	//		log.Println("successful q2")
	//	}
}
