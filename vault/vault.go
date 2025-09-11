package vault

import (
	"context"
	"time"

	vault "github.com/hashicorp/vault/api"
	auth "github.com/hashicorp/vault/api/auth/approle"
	"github.com/mitchellh/mapstructure"
)

type DBParams struct {
	DBUser     string `mapstructure:"database-user"`
	DBPassword string `mapstructure:"database-password"`
	DBHost     string `mapstructure:"database-host"`
	DBPort     int    `mapstructure:"database-port"`
	DBName     string `mapstructure:"database-name"`
}

func GetDBParams(ctx context.Context, vaultHost, roleId, secretId, secretsPath, mountPath string) (*DBParams, error) {

	var dbParams DBParams
	config := vault.DefaultConfig()
	config.Timeout = time.Second * 30
	config.Address = vaultHost
	client, err := vault.NewClient(config)
	if err != nil {
		return nil, err
	}
	secretID := &auth.SecretID{FromString: secretId}
	appRoleAuth, err := auth.NewAppRoleAuth(
		roleId,
		secretID,
		auth.WithWrappingToken(),
	)
	if err != nil {
		return nil, err
	}

	authInfo, err := client.Auth().Login(ctx, appRoleAuth)
	if err != nil {
		return nil, err
	}
	if authInfo == nil {
		return nil, err
	}

	secret, err := client.KVv2(mountPath).Get(ctx, secretsPath)
	if err != nil {
		return nil, err
	}
	err = mapstructure.Decode(secret.Data, &dbParams)
	if err != nil {
		return nil, err
	}
	return &dbParams, nil
}
