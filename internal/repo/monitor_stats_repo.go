package repo

import (
	"context"

	"github.com/dushixiang/pika/internal/models"
	"github.com/go-orz/orz"
	"gorm.io/gorm"
)

type MonitorStatsRepo struct {
	orz.Repository[models.MonitorStats, string]
	db *gorm.DB
}

func NewMonitorStatsRepo(db *gorm.DB) *MonitorStatsRepo {
	return &MonitorStatsRepo{
		Repository: orz.NewRepository[models.MonitorStats, string](db),
		db:         db,
	}
}

func (r *MonitorStatsRepo) FindByMonitorId(ctx context.Context, monitorId string) ([]models.MonitorStats, error) {
	var statsList []models.MonitorStats
	err := r.db.WithContext(ctx).
		Where("monitor_id = ?", monitorId).
		Find(&statsList).Error
	return statsList, err
}

func (r *MonitorStatsRepo) FindByMonitorIdIn(ctx context.Context, monitorIds []string) ([]models.MonitorStats, error) {
	var statsList []models.MonitorStats
	err := r.db.WithContext(ctx).
		Where("monitor_id in ?", monitorIds).
		Find(&statsList).Error
	return statsList, err
}

func (r *MonitorStatsRepo) DeleteByMonitorId(ctx context.Context, monitorId string) error {
	return r.db.WithContext(ctx).
		Where("monitor_id = ?", monitorId).
		Delete(&models.MonitorStats{}).Error
}

func (r *MonitorStatsRepo) DeleteByAgentId(ctx context.Context, agentId string) error {
	return r.db.WithContext(ctx).
		Where("agent_id = ?", agentId).
		Delete(&models.MonitorStats{}).Error
}

// FindAll 查找所有统计数据
func (r *MonitorStatsRepo) FindAll(ctx context.Context) ([]models.MonitorStats, error) {
	var statsList []models.MonitorStats
	err := r.db.WithContext(ctx).Find(&statsList).Error
	return statsList, err
}

// DeleteByIDs 批量删除指定ID的统计数据
func (r *MonitorStatsRepo) DeleteByIDs(ctx context.Context, ids []string) error {
	if len(ids) == 0 {
		return nil
	}
	return r.db.WithContext(ctx).
		Where("id IN ?", ids).
		Delete(&models.MonitorStats{}).Error
}
