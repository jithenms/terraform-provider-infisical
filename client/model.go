package infisicalclient

import "time"

type GetEncryptedSecretsV3Request struct {
	Environment string `json:"environment"`
	WorkspaceId string `json:"workspaceId"`
	SecretPath  string `json:"secretPath"`
}

type EncryptedSecretV3 struct {
	ID        string `json:"_id"`
	Version   int    `json:"version"`
	Workspace string `json:"workspace"`
	Type      string `json:"type"`
	Tags      []struct {
		ID        string `json:"_id"`
		Name      string `json:"name"`
		Slug      string `json:"slug"`
		Workspace string `json:"workspace"`
	} `json:"tags"`
	Environment             string    `json:"environment"`
	SecretKeyCiphertext     string    `json:"secretKeyCiphertext"`
	SecretKeyIV             string    `json:"secretKeyIV"`
	SecretKeyTag            string    `json:"secretKeyTag"`
	SecretValueCiphertext   string    `json:"secretValueCiphertext"`
	SecretValueIV           string    `json:"secretValueIV"`
	SecretValueTag          string    `json:"secretValueTag"`
	SecretCommentCiphertext string    `json:"secretCommentCiphertext"`
	SecretCommentIV         string    `json:"secretCommentIV"`
	SecretCommentTag        string    `json:"secretCommentTag"`
	Algorithm               string    `json:"algorithm"`
	KeyEncoding             string    `json:"keyEncoding"`
	Folder                  string    `json:"folder"`
	V                       int       `json:"__v"`
	CreatedAt               time.Time `json:"createdAt"`
	UpdatedAt               time.Time `json:"updatedAt"`
}

type GetEncryptedSecretsV3Response struct {
	Secrets []EncryptedSecretV3 `json:"secrets"`
}

type GetServiceTokenDetailsResponse struct {
	ID           string    `json:"_id"`
	Name         string    `json:"name"`
	Workspace    string    `json:"workspace"`
	Environment  string    `json:"environment"`
	ExpiresAt    time.Time `json:"expiresAt"`
	EncryptedKey string    `json:"encryptedKey"`
	Iv           string    `json:"iv"`
	Tag          string    `json:"tag"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	V            int       `json:"__v"`
}

type UniversalMachineIdentityAuthResponse struct {
	AccessToken       string `json:"accessToken"`
	ExpiresIn         int    `json:"expiresIn"`
	AccessTokenMaxTTL int    `json:"accessTokenMaxTTL"`
	TokenType         string `json:"tokenType"`
}

type SingleEnvironmentVariable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
	ID    string `json:"_id"`
	Tags  []struct {
		ID        string `json:"_id"`
		Name      string `json:"name"`
		Slug      string `json:"slug"`
		Workspace string `json:"workspace"`
	} `json:"tags"`
	Comment string `json:"comment"`
}

type SymmetricEncryptionResult struct {
	CipherText []byte `json:"CipherText"`
	Nonce      []byte `json:"Nonce"`
	AuthTag    []byte `json:"AuthTag"`
}

// Workspace key request
type GetEncryptedWorkspaceKeyRequest struct {
	WorkspaceId string `json:"workspaceId"`
}

// Workspace key response
type GetEncryptedWorkspaceKeyResponse struct {
	ID           string `json:"_id"`
	EncryptedKey string `json:"encryptedKey"`
	Nonce        string `json:"nonce"`
	Sender       struct {
		ID             string    `json:"_id"`
		Email          string    `json:"email"`
		RefreshVersion int       `json:"refreshVersion"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
		V              int       `json:"__v"`
		FirstName      string    `json:"firstName"`
		LastName       string    `json:"lastName"`
		PublicKey      string    `json:"publicKey"`
	} `json:"sender"`
	Receiver  string    `json:"receiver"`
	Workspace string    `json:"workspace"`
	V         int       `json:"__v"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// encrypted secret
type EncryptedSecret struct {
	SecretName              string `json:"secretName"`
	WorkspaceID             string `json:"workspaceId"`
	Type                    string `json:"type"`
	Environment             string `json:"environment"`
	SecretKeyCiphertext     string `json:"secretKeyCiphertext"`
	SecretKeyIV             string `json:"secretKeyIV"`
	SecretKeyTag            string `json:"secretKeyTag"`
	SecretValueCiphertext   string `json:"secretValueCiphertext"`
	SecretValueIV           string `json:"secretValueIV"`
	SecretValueTag          string `json:"secretValueTag"`
	SecretCommentCiphertext string `json:"secretCommentCiphertext"`
	SecretCommentIV         string `json:"secretCommentIV"`
	SecretCommentTag        string `json:"secretCommentTag"`
	SecretPath              string `json:"secretPath"`
}

// create secrets
type CreateSecretV3Request struct {
	SecretName              string `json:"secretName"`
	WorkspaceID             string `json:"workspaceId"`
	Type                    string `json:"type"`
	Environment             string `json:"environment"`
	SecretKeyCiphertext     string `json:"secretKeyCiphertext"`
	SecretKeyIV             string `json:"secretKeyIV"`
	SecretKeyTag            string `json:"secretKeyTag"`
	SecretValueCiphertext   string `json:"secretValueCiphertext"`
	SecretValueIV           string `json:"secretValueIV"`
	SecretValueTag          string `json:"secretValueTag"`
	SecretCommentCiphertext string `json:"secretCommentCiphertext"`
	SecretCommentIV         string `json:"secretCommentIV"`
	SecretCommentTag        string `json:"secretCommentTag"`
	SecretPath              string `json:"secretPath"`
}

// delete secret by name api
type DeleteSecretV3Request struct {
	SecretName  string `json:"secretName"`
	WorkspaceId string `json:"workspaceId"`
	Environment string `json:"environment"`
	Type        string `json:"type"`
	SecretPath  string `json:"secretPath"`
}

// update secret by name api
type UpdateSecretByNameV3Request struct {
	SecretName            string `json:"secretName"`
	WorkspaceID           string `json:"workspaceId"`
	Environment           string `json:"environment"`
	Type                  string `json:"type"`
	SecretPath            string `json:"secretPath"`
	SecretValueCiphertext string `json:"secretValueCiphertext"`
	SecretValueIV         string `json:"secretValueIV"`
	SecretValueTag        string `json:"secretValueTag"`
}

// get secret by name api
type GetSingleSecretByNameV3Request struct {
	SecretName  string `json:"secretName"`
	WorkspaceId string `json:"workspaceId"`
	Environment string `json:"environment"`
	Type        string `json:"type"`
	SecretPath  string `json:"secretPath"`
}

type GetSingleSecretByNameSecretResponse struct {
	Secret EncryptedSecret `json:"secret"`
}

type GetRawSecretsV3Request struct {
	Environment string `json:"environment"`
	WorkspaceId string `json:"workspaceId"`
	SecretPath  string `json:"secretPath"`
}

type RawV3Secret struct {
	Version       int    `json:"version"`
	Workspace     string `json:"workspace"`
	Type          string `json:"type"`
	Environment   string `json:"environment"`
	SecretKey     string `json:"secretKey"`
	SecretValue   string `json:"secretValue"`
	SecretComment string `json:"secretComment"`
}

type GetRawSecretsV3Response struct {
	Secrets []RawV3Secret `json:"secrets"`
}

type GetSingleRawSecretByNameSecretResponse struct {
	Secret RawV3Secret `json:"secret"`
}

// create secrets
type CreateRawSecretV3Request struct {
	WorkspaceID   string `json:"workspaceId"`
	Type          string `json:"type"`
	Environment   string `json:"environment"`
	SecretKey     string `json:"secretKey"`
	SecretValue   string `json:"secretValue"`
	SecretComment string `json:"secretComment"`
	SecretPath    string `json:"secretPath"`
}

type DeleteRawSecretV3Request struct {
	SecretName  string `json:"secretName"`
	WorkspaceId string `json:"workspaceId"`
	Environment string `json:"environment"`
	Type        string `json:"type"`
	SecretPath  string `json:"secretPath"`
}

// update secret by name api
type UpdateRawSecretByNameV3Request struct {
	SecretName  string `json:"secretName"`
	WorkspaceID string `json:"workspaceId"`
	Environment string `json:"environment"`
	Type        string `json:"type"`
	SecretPath  string `json:"secretPath"`
	SecretValue string `json:"secretValue"`
}
