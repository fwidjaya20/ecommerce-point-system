package user_point

import (
	"context"
	"fmt"
	models2 "github.com/fwidjaya20/ecommerce-point-system/internal/databases/models"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/models"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/repositories"
	"github.com/fwidjaya20/ecommerce-point-system/internal/domains/user_point/values"
	libError "github.com/fwidjaya20/ecommerce-point-system/lib/error"
	"github.com/fwidjaya20/ecommerce-point-system/lib/nats"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"
	"net/http"
)

type service struct {
	actor string
	logger log.Logger
	repository repositories.Interface
}

func (s *service) PublishAddPoint(ctx context.Context, payload models.PublishAddOrDeductPointRequest) (*models.UserPointInquiryResponse, error) {
	logger := log.With(s.logger, "METHOD", "PublishAddPoint()")

	var err error

	eventId, err := uuid.NewRandom()
	if nil != err {
		_ = level.Error(logger).Log(values.CREATE_UUID_ERROR, err)
		return nil, libError.NewError(err, http.StatusInternalServerError, values.CREATE_UUID_ERROR, "Gagal Generate UUID")
	}

	err = nats.GetGlobalPublisher().Store(nats.StoreDetail{
		Channel:     "user_point",
		Domain:      "user_point",
		Subject:     "user_point.add",
		EventSource: "user_point.store_event",
		Data: map[string]interface{}{
			"data": models.AddOrDeductPointRequest{
				Id:        eventId.String(),
				UserId:    payload.UserId,
				Point:     payload.Point,
				PointType: values.ADD_POINT_TYPE,
				Notes:     fmt.Sprintf(`{"message": "%s", "amount": %f}`, values.ADD_POINT_TYPE, payload.Point),
			},
			"metadata": map[string]interface{}{
				"version": 1,
			},
		},
	})

	if nil != err {
		_ = level.Error(logger).Log(values.PUBLISH_ADD_POINT_ERROR, err)
		return nil, libError.NewError(err, http.StatusInternalServerError, values.PUBLISH_ADD_POINT_ERROR,"Gagal publish event")
	}

	return &models.UserPointInquiryResponse{
		RequestId: eventId.String(),
	}, nil
}

func (s *service) PublishDeductPoint(ctx context.Context, payload models.PublishAddOrDeductPointRequest) (*models.UserPointInquiryResponse, error) {
	logger := log.With(s.logger, "METHOD", "PublishDeductPoint()")

	var err error

	eventId, err := uuid.NewRandom()
	if nil != err {
		_ = level.Error(logger).Log(values.CREATE_UUID_ERROR, err)
		return nil, libError.NewError(err, http.StatusInternalServerError, values.CREATE_UUID_ERROR, "Gagal Generate UUID")
	}

	err = nats.GetGlobalPublisher().Store(nats.StoreDetail{
		Channel:     "user_point",
		Domain:      "user_point",
		Subject:     "user_point.deduct",
		EventSource: "user_point.store_event",
		Data: map[string]interface{}{
			"data": models.AddOrDeductPointRequest{
				Id:        eventId.String(),
				UserId:    payload.UserId,
				Point:     payload.Point,
				PointType: values.DEDUCT_POINT_TYPE,
				Notes:     fmt.Sprintf(`{"message": "%s", "amount": %f}`, values.DEDUCT_POINT_TYPE, payload.Point),
			},
			"metadata": map[string]interface{}{
				"version": 1,
			},
		},
	})

	if nil != err {
		_ = level.Error(logger).Log(values.PUBLISH_ADD_POINT_ERROR, err)
		return nil, libError.NewError(err, http.StatusInternalServerError, values.PUBLISH_ADD_POINT_ERROR,"Gagal publish event")
	}

	return &models.UserPointInquiryResponse{
		RequestId: eventId.String(),
	}, nil
}

func (s *service) StoreEvent(ctx context.Context, payload models.AddOrDeductPointRequest) error {
	logger := log.With(s.logger, "METHOD", "StoreEvent()")

	var err error

	err = s.repository.StoreEvent(ctx, &models2.UserPointEvent{
		Id:        payload.Id,
		UserId:    payload.UserId,
		Point:     payload.Point,
		PointType: payload.PointType,
		Notes:     payload.Notes,
	})

	if nil != err {
		_ = level.Error(logger).Log("Error", err)
		return libError.NewError(err, http.StatusInternalServerError, values.STORE_POINT_ERROR,"Gagal memperbaharui point user")
	}

	return nil
}

func (s *service) GetPoint(ctx context.Context, payload models.GetPointRequest) (*models.UserPointResponse, error) {
	logger := log.With(s.logger, "METHOD", "GetPoint()")

	var result *models2.UserPointSnapshot
	var err error

	result, err = s.repository.GetPoint(ctx, payload.UserId)

	if nil != err {
		_ = level.Error(logger).Log("Error", err)
		return nil, libError.NewError(err, http.StatusInternalServerError, values.GET_POINT_ERROR, "Gagal memperoleh data poin")
	}

	return &models.UserPointResponse{
		TotalPoint: result.Point,
	}, nil
}

func NewUserPointService(
	logger log.Logger,
	repository repositories.Interface,
) UseCase {
	service := service{
		actor:      "USER_POINT",
		logger:     nil,
		repository: repository,
	}

	service.logger = log.With(logger, "ACTOR", service.actor)

	return &service
}