package test

import (
	"github.com/go-chi/render"
	"golang.org/x/exp/slog"
	"net/http"
	"time"
)

func New(log *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		const op = "handlers.url.test.New"

		go func() {
			log.Info("test req started")

			time.Sleep(time.Second * 5)

			log.Info("test req finished")

			render.PlainText(w, r, "Request Finished")
			return
		}()
		//log := log.With(
		//	slog.String("op", op),
		//	slog.String("request_id", middleware.GetReqID(r.Context())),
		//)
		//
		//alias := chi.URLParam(r, "alias")
		//if alias == "" {
		//	log.Info("alias is empty")
		//
		//	render.JSON(w, r, resp.Error("invalid request"))
		//
		//	return
		//}
		//
		//resURL, err := urlGetter.GetURL(alias)
		//if errors.Is(err, storage.ErrURLNotFound) {
		//	log.Info("url not found", "alias", alias)
		//
		//	render.JSON(w, r, resp.Error("not found"))
		//
		//	return
		//}
		//if err != nil {
		//	log.Error("failed to get url", sl.Err(err))
		//
		//	render.JSON(w, r, resp.Error("internal error"))
		//
		//	return
		//}
		//
		//log.Info("got url", slog.String("url", resURL))
		//
		//// redirect to found url
		//http.Redirect(w, r, resURL, http.StatusFound)
	}
}
