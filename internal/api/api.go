package api

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"journey/internal/api/spec"
	"journey/internal/pgstore"
	"net/http"
)

type store interface {
	GetParticipant(ctx context.Context, participantID uuid.UUID) (pgstore.Participant, error)
	ConfirmParticipant(ctx context.Context, participantID uuid.UUID) error
}

type API struct {
	store  store
	logger *zap.Logger
}

func NewAPI(pool *pgxpool.Pool, logger *zap.Logger) API {
	return API{
		pgstore.New(pool), logger,
	}
}

func (api API) PatchParticipantsParticipantIDConfirm(w http.ResponseWriter, r *http.Request, participantID string) *spec.Response {
	id, err := uuid.Parse(participantID)
	if err != nil {
		return spec.PatchParticipantsParticipantIDConfirmJSON400Response(spec.Error{Message: "Invalid UUID"})
	}

	participant, err := api.store.GetParticipant(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return spec.PatchParticipantsParticipantIDConfirmJSON400Response(spec.Error{Message: "Participant Not Found"})
		}
		api.logger.Error("failed to get participant", zap.Error(err), zap.String("participant_id", participantID))
		return spec.PatchParticipantsParticipantIDConfirmJSON400Response(spec.Error{Message: "Server Error"})
	}
	if participant.IsConfirmed {
		return spec.PatchParticipantsParticipantIDConfirmJSON400Response(spec.Error{Message: "Participant already confirmed"})
	}

	if err := api.store.ConfirmParticipant(r.Context(), id); err != nil {
		api.logger.Error("failed to confirm participant", zap.Error(err), zap.String("participant_id", participantID))
		return spec.PatchParticipantsParticipantIDConfirmJSON400Response(spec.Error{Message: "Server Error"})
	}

	return spec.PatchParticipantsParticipantIDConfirmJSON204Response(nil)
}

func (A API) PostTrips(w http.ResponseWriter, r *http.Request) *spec.Response {
	//TODO implement me
	panic("implement me")
}

func (A API) GetTripsTripID(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	//TODO implement me
	panic("implement me")
}

func (A API) PutTripsTripID(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	//TODO implement me
	panic("implement me")
}

func (A API) GetTripsTripIDActivities(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	//TODO implement me
	panic("implement me")
}

func (A API) PostTripsTripIDActivities(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	//TODO implement me
	panic("implement me")
}

func (A API) GetTripsTripIDConfirm(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	//TODO implement me
	panic("implement me")
}

func (A API) PostTripsTripIDInvites(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	//TODO implement me
	panic("implement me")
}

func (A API) GetTripsTripIDLinks(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	//TODO implement me
	panic("implement me")
}

func (A API) PostTripsTripIDLinks(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	//TODO implement me
	panic("implement me")
}

func (A API) GetTripsTripIDParticipants(w http.ResponseWriter, r *http.Request, tripID string) *spec.Response {
	//TODO implement me
	panic("implement me")
}
