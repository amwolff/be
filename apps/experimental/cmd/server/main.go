package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"

	"go.uber.org/zap"

	"github.com/amwolff/be/apps/experimental/pkg/classifier"
)

func getFpHandler(sugar *zap.SugaredLogger, c classifier.Classifier) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var f classifier.Fingerprint

		b, err := httputil.DumpRequest(r.Clone(context.TODO()), true)
		if err != nil {
			sugar.Warnf("DumpRequest: %v", err)
		} else {
			sugar.Debug(string(b))
		}

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
	l, err := zap.NewProduction()
	if err != nil {
		panic(fmt.Sprintf("NewProduction: %v", err))
	}

	defer l.Sync()

	sugar := l.Sugar()
}
