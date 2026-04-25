package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GenerateJwtSecret(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("Erro na geração dos bytes: %w", err)
	}

	// Codifica em Base64 para facilitar o armazenamento em variáveis de ambiente
	return base64.StdEncoding.EncodeToString(bytes), nil
}
