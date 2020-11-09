package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	"github.com/amwolff/be/apps/experimental/pkg/classifier"
	"go.uber.org/zap"
)

func logRequest(sugar *zap.SugaredLogger, r *http.Request) {
	b, err := httputil.DumpRequest(r, true)
	if err != nil {
		sugar.Debugf("DumpRequest: %v", err)
	} else {
		sugar.Debugf("\n%s", strings.TrimSpace(string(b)))
	}
}

func getRootHandler(sugar *zap.SugaredLogger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		logRequest(sugar, r)

		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func getFpHandler(sugar *zap.SugaredLogger, c classifier.Classifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")

		logRequest(sugar, r)

		if r.Method != http.MethodPost {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		d := json.NewDecoder(r.Body)
		d.DisallowUnknownFields()

		var f classifier.Fingerprint

		if err := d.Decode(&f); err != nil {
			sugar.Errorf("NewDecoder: %v", err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		var entropy float64

		sf, ok := c.Do(f)
		if !ok {
			sugar.Infow("Got new", "f", f.VisitorID)
			sugar.Debug(f)

			entropy = c.Store(f.VisitorID, f)

			http.Error(w, http.StatusText(http.StatusCreated), http.StatusCreated)
		} else {
			sugar.Infow("Got similar", "f", f.VisitorID, "similar", sf.VisitorID)
			sugar.Debug(f)
			sugar.Debug(sf)

			entropy = c.Store(sf.VisitorID, f)
		}

		sugar.Infof("The set has now %.2f bits of entropy", entropy)
	}
}

func main() {
	l, err := zap.NewDevelopment(zap.IncreaseLevel(zap.InfoLevel))
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
