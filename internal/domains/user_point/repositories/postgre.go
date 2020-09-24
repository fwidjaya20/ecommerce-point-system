package repositories

import (
"context"
"database/sql"
"github.com/fwidjaya20/ecommerce-point-system/internal/databases/models"
"github.com/fwidjaya20/ecommerce-point-system/internal/globals"
"strings"
)

type postgre struct {}

func (p postgre) StoreEvent(ctx context.Context, model *models.UserPointEvent) error {
	var err error
	query, args := p.buildStoreEventQuery(model)
	_, err = globals.GetQuery(ctx).NamedExecContext(ctx, query, args)

	if nil != err {
		return err
	}

	err = p.createSnapshot(ctx, model)

	return err
}

func (p postgre) GetPoint(ctx context.Context, userId string) (*models.UserPointSnapshot, error) {
	snapshot, err := p.getLastSnapshot(ctx, userId)

	if nil != err {
		return nil, err
	}

	return snapshot, nil
}

func NewUserPointRepository() Interface {
	return &postgre{}
}

func (p *postgre) createSnapshot(ctx context.Context, model *models.UserPointEvent) error {
	lastSnapshot, err := p.getLastSnapshot(ctx, model.UserId)
	if nil != err && err != sql.ErrNoRows {
		return err
	}

	var lastPoint float64 = 0

	if lastSnapshot != nil {
		lastPoint = lastSnapshot.Point
	}

	snapshotModel := models.UserPointSnapshot{
		Id:          model.Id,
		UserId:      model.UserId,
		Point:       lastPoint + model.Point,
		LastEventId: model.Id,
	}

	query, args := p.buildCreateSnapshotQuery(snapshotModel)
	_, err = globals.GetQuery(ctx).NamedExecContext(ctx, query, args)

	return err
}

func (p *postgre) getLastSnapshot(ctx context.Context, userId string) (*models.UserPointSnapshot, error) {
	var result models.UserPointSnapshot
	var err error

	query, arg := p.buildGetLastSnapshotQuery(userId)
	row, err := globals.GetQuery(ctx).NamedQueryRowxContext(ctx, query, arg)

	if nil != err {
		return nil, err
	}

	err = row.StructScan(&result)

	return &result, err
}


func (p *postgre) buildGetLastSnapshotQuery(userId string) (string, interface{}) {
	var query strings.Builder
	var arg map[string]interface{} = make(map[string]interface{})

	query.WriteString(`SELECT "id", "user_id", "point", "last_event_id" `)
	query.WriteString(`FROM "user_point_snapshots" `)
	query.WriteString(`WHERE "user_id"=:userId `)
	query.WriteString(`ORDER BY "created_at" DESC `)
	query.WriteString(`LIMIT 1`)

	arg["userId"] = userId

	return query.String(), arg
}

func (p *postgre) buildStoreEventQuery(model *models.UserPointEvent) (string, interface{}) {
	var query strings.Builder
	var arg map[string]interface{} = make(map[string]interface{})

	query.WriteString(`INSERT INTO "user_point_events" `)
	query.WriteString(`("id", "user_id", "point", "point_type", "notes", "created_by") `)
	query.WriteString(`VALUES `)
	query.WriteString(`(:id, :userId, :point, :pointType, :notes, :createdBy)`)

	arg["id"] = model.Id
	arg["userId"] = model.UserId
	arg["point"] = model.Point
	arg["pointType"] = model.PointType
	arg["notes"] = model.Notes
	arg["createdBy"] = "SYSTEM"

	return query.String(), arg
}

func (p *postgre) buildCreateSnapshotQuery(model models.UserPointSnapshot) (string, interface{}) {
	var query strings.Builder
	var arg map[string]interface{} = make(map[string]interface{})

	query.WriteString(`INSERT INTO "user_point_snapshots" `)
	query.WriteString(`("id", "user_id", "point", "last_event_id", "created_by") `)
	query.WriteString(`VALUES `)
	query.WriteString(`(:id, :userId, :point, :lastEventId, :createdBy)`)

	arg["id"] = model.Id
	arg["userId"] = model.UserId
	arg["point"] = model.Point
	arg["lastEventId"] = model.LastEventId
	arg["createdBy"] = "SYSTEM"

	return query.String(), arg
}