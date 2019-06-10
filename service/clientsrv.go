package service

import (
  "github.com/liu578101804/kun_auth2.0/models"
  "github.com/liu578101804/kun_auth2.0/utils"
  "github.com/liu578101804/de/database"
)

func CreateClient() *models.OauthClient {
  var(
    client models.OauthClient
    clientKey string
    clientSecret string
  )
  clientKey = utils.CreateClientKey()
  clientSecret = utils.CreateClientSecret()
  client = models.OauthClient{
    ClientKey: clientKey,
    ClientSecret: clientSecret,
  }
  return &client
}

func ClientAddService(clientM *models.OauthClient) (int64, error) {
  return database.G_engine.Insert(clientM)
}

