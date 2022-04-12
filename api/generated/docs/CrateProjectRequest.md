# CrateProjectRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | Unique project identifier, might be used for idempotency | [default to null]
**Name** | **string** | Project name | [default to null]
**OwnerId** | **string** | Project owner id | [default to null]
**State** | **string** | Project state; Might be created non-delault for creating prioject post-factum | [optional] [default to STATE.PLANNED]
**ParticipantIds** | **[]string** | Ids of the participants | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

