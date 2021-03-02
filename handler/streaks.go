package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/denizaygun/streakify/db"
	"github.com/denizaygun/streakify/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

var streakIDKey = "streakID"

func streaks(router chi.Router) {
	router.Get("/", getAllStreaks)
	router.Post("/", createStreak)
	router.Route("/{streakId}", func(router chi.Router) {
		router.Use(StreakContext)
		router.Get("/", getStreak)
		router.Put("/", updateStreak)
		router.Delete("/", deleteStreak)
	})
}

// StreakContext ...
func StreakContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		streakID := chi.URLParam(r, "streakID")
		if streakID == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("streak ID is required")))
			return
		}
		id, err := strconv.Atoi(streakID)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid streak ID")))
		}
		ctx := context.WithValue(r.Context(), streakIDKey, id)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func createStreak(w http.ResponseWriter, r *http.Request) {
	streak := &models.Streak{}
	if err := render.Bind(r, streak); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddStreak(streak); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, streak); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func getAllStreaks(w http.ResponseWriter, r *http.Request) {
	streaks, err := dbInstance.GetAllStreaks()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, streaks); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func getStreak(w http.ResponseWriter, r *http.Request) {
	streakID := r.Context().Value(streakIDKey).(int)
	streak, err := dbInstance.GetStreakByID(streakID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &streak); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func deleteStreak(w http.ResponseWriter, r *http.Request) {
	streakID := r.Context().Value(streakIDKey).(int)
	err := dbInstance.DeleteStreak(streakID)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
}

func updateStreak(w http.ResponseWriter, r *http.Request) {
	streakID := r.Context().Value(streakIDKey).(int)
	streakData := models.Streak{}
	if err := render.Bind(r, &streakData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	streak, err := dbInstance.UpdateStreak(streakID, streakData)
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	if err := render.Render(w, r, &streak); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}
