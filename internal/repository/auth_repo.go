package repository

import (
	"context"
	"time"

	"github.com/TaperoOO5536/special_admin/internal/models"
	"gorm.io/gorm"
)

type AdminAuthRepository interface {
    GetAdmin(ctx context.Context, login string) (*models.Admin, error)
	UpdateRefreshToken(ctx context.Context, tokenHash string, expiresAt time.Time) error
	GetRefreshToken(ctx context.Context, tokenHash string) (*models.Admin, error)
	ClearRefreshToken(ctx context.Context) error
}

type adminAuthRepository struct {
	db *gorm.DB
}

func NewAdminAuthRepository(db *gorm.DB) AdminAuthRepository {
	return &adminAuthRepository{db: db}
}

func (r *adminAuthRepository) GetAdmin(ctx context.Context, login string) (*models.Admin, error) {
    var admin models.Admin
    err := r.db.Where("admin_login = ?", login).First(&admin).Error
    return &admin, err
}

func (r *adminAuthRepository) UpdateRefreshToken(ctx context.Context, tokenHash string, expiresAt time.Time) error {
    return r.db.Model(&models.Admin{}).
        Where("admin_login = ?", "admin").
        Updates(map[string]interface{}{
            "refresh_token_hash": tokenHash,
            "refresh_expires_at": expiresAt,
        }).Error
}

func (r *adminAuthRepository) GetRefreshToken(ctx context.Context, tokenHash string) (*models.Admin, error) {
    var admin models.Admin
    err := r.db.Where("refresh_token_hash = ? AND refresh_expires_at > NOW()", tokenHash).
        First(&admin).Error
    return &admin, err
}

func (r *adminAuthRepository) ClearRefreshToken(ctx context.Context) error {
    return r.db.Model(&models.Admin{Login: "admin"}).
        Update("refresh_token_hash", nil).Error
}
