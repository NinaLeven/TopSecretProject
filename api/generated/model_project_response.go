/*
 * Project Manager API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type ProjectResponse struct {
	// Unique project identifier
	Uid string `json:"uid"`
	// Project name
	Name string `json:"name"`
	// Project owner id
	OwnerId string `json:"owner_id"`
	// Project state
	State string `json:"state"`
	// Project progress in %
	Progress int32 `json:"progress,omitempty"`
	// Ids of the participants
	ParticipantIds []string `json:"participant_ids,omitempty"`
	// Created at timestamp rfc 3339
	CreatedAt string `json:"created_at,omitempty"`
	// Updated at timestamp rfc 3339
	UpdatedAt string `json:"updated_at,omitempty"`
}
