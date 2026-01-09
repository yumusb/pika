package repo

import (
	"context"

	"github.com/dushixiang/pika/internal/models"
	"github.com/go-orz/orz"
	"gorm.io/gorm"
)

// SSHLoginEventRepo SSH登录事件数据访问层
type SSHLoginEventRepo struct {
	orz.Repository[models.SSHLoginEvent, string]
}

// NewSSHLoginEventRepo 创建仓库
func NewSSHLoginEventRepo(db *gorm.DB) *SSHLoginEventRepo {
	return &SSHLoginEventRepo{
		Repository: orz.NewRepository[models.SSHLoginEvent, string](db),
	}
}

// DeleteEventsByAgentID 删除探针的所有登录事件
func (r *SSHLoginEventRepo) DeleteEventsByAgentID(ctx context.Context, agentID string) error {
	return r.GetDB(ctx).Where("agent_id = ?", agentID).Delete(&models.SSHLoginEvent{}).Error
}

// DeleteOldEvents 删除旧的事件记录（保留最近N天）
func (r *SSHLoginEventRepo) DeleteOldEvents(ctx context.Context, beforeTimestamp int64) error {
	return r.GetDB(ctx).Where("timestamp < ?", beforeTimestamp).Delete(&models.SSHLoginEvent{}).Error
}

// FindEventByTimestamp 查找指定时间范围内的事件（用于去重）
func (r *SSHLoginEventRepo) FindEventByTimestamp(ctx context.Context, agentID string, timestamp, tolerance int64) (*models.SSHLoginEvent, error) {
	var event models.SSHLoginEvent
	err := r.GetDB(ctx).Where("agent_id = ? AND timestamp >= ? AND timestamp <= ?",
		agentID, timestamp-tolerance, timestamp+tolerance).
		First(&event).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &event, err
}
