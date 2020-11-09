package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/amwolff/be/apps/experimental/pkg/classifier"
	"go.uber.org/zap"
)

func getRootHandler(sugar *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		b, err := httputil.DumpRequest(r.Clone(context.TODO()), true)
		if err != nil {
			sugar.Debugf("DumpRequest: %v", err)
		} else {
			sugar.Debugf("\n%s", strings.TrimSpace(string(b)))
		}

		http.Error(w, "", http.StatusBadRequest)
	}
}

func getFpHandler(sugar *zap.SugaredLogger, c classifier.Classifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		b, err := httputil.DumpRequest(r, false)
		if err != nil {
			sugar.Debugf("DumpRequest: %v", err)
		} else {
			sugar.Debugf("\n%s", strings.TrimSpace(string(b)))
		}

		if r.Method != http.MethodPost {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		var f classifier.Fingerprint

		if err := json.NewDecoder(r.Body).Decode(&f); err != nil {
			sugar.Errorf("NewDecoder: %v", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		sf, ok := c.Do(f)
		if !ok {
			sugar.Infow("Got new", "f", f.VisitorID)
			sugar.Debug(f)

			e := c.Store(f.VisitorID, f)

			sugar.Infof("The set has now %.2f bits of entropy", e)

			http.Error(w, "", http.StatusCreated)
			return
		}

		sugar.Infow("Got similar", "f", f.VisitorID, "similar", sf.VisitorID)
		sugar.Debug(f)
		sugar.Debug(sf)

		e := c.Store(sf.VisitorID, f)

		sugar.Infof("The set has now %.2f bits of entropy", e)
	}
}

func main() {
	l, err := zap.NewDevelopment()
	if err != nil {
		panic(fmt.Sprintf("NewProduction: %v", err))
	}

	defer l.Sync()

	sugar := l.Sugar()

	c := classifier.NewExperimentalInMemory()

	http.Handle("/", getRootHandler(sugar.Named("root handler")))
	http.Handle("/fp", getFpHandler(sugar.Named("fp handler"), c))

	const addr = ":8084"

	go sugar.Infof("Started to listen at %s", addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		sugar.Errorf("ListenAndServe: %v", err)
	}
}
