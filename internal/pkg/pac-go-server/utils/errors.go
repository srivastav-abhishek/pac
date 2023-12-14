package utils

var (

	// Get Catalogs Error Response Msgs
	// Associated APIs :
	// GET:
	// - /catalogs
	// - /catalogs/:name
	GetCatalogsErrMsg        = "failed to get catalogs Error: <error>"
	CatalogaNameNotSetErrMsg = "catalog name is not set"
	CatalogNotExistErrMsg    = "catalog with name <name> does not exist"
	CatalogGetFailedErrMsg   = "failed to get catalog with name <name> Error: <error>"

	// Create Catalog Error Response Msgs
	// Associated APIs:
	// POST:
	// - /catalogs
	JsonBindErrMsg                      = "failed to bind request, Error: <error>"
	CatalogTypeSetErrMsg                = "catalog type should be set"
	InvalidCatalogTypeErrMsg            = "invalid catalog type <type>, only valid catalog is <type>"
	CatalogNameNotSetErrMsg             = "catalog name should be set"
	CatalogCPUCapacityNotSetErrMsg      = "catalog cpu capacity should be set"
	CatalogMemoryapacityNotSetErrMsg    = "catalog memory capacity should be set"
	CatalogExpiryNotSetErrMsg           = "catalog expiry should be set"
	CatalogVMCrnNotSetErrMsg            = "for catalog type VM crn should be set"
	CatalogVMSystemTypeNotSetErrMsg     = "for catalog type VM system_type should be set"
	CatalogVMProcessorTypeNotSetErrMsg  = "for catalog type VM processor_type should be set"
	CatalogVMImageNotSetErrMsg          = "for catalog type VM image should be set"
	CatalogVMCpuNotSetErrMsg            = "for catalog type VM cpu capacity should be set"
	CatalogVMMemoryCapacityNotSetErrMsg = "for catalog type VM memory capacity should be set"

	// <name> is only for readability purpose, will be replaced with %s in actual implementation (same for all)
	CatalogAlreadyExistErrMsg = "catalog with name <name> already exist"
	CreateCatalogFailedErrMsg = "failed to create catalog Error: <error>"

	// Delete Catalog Error Response Msgs
	// Associated APIs:
	// DELETE:
	// - /catalogs
	DeleteCatalogFailedErrMsg = "failed to delete catalog with name <name> Error: <error>"

	// Retire Catalog Error Response Msgs
	// Associated APIs:
	// PUT:
	// - /catalogs/:name/retire
	RetireCatalogFailedErrMsg = "failed to retire catalog with name <name> Error: <error>"

	// Get Events Error Response Msgs
	// Associated APIs:
	// GET:
	// - /events
	GetEventsFailedErrMsg           = "error getting events: <error>"
	FetchEventsFailedErrMsg         = "error fetching events: <error>"
	GetTotalEventsCountFailedErrMsg = "error getting total count of events: <error>"

	// Get All Groups Error Response Msgs
	// Associated APIs:
	// GET:
	// - /groups
	EmptyGroupsErrMsg       = "groups is empty"
	GetGroupsFailedErrMsg   = "error getting requests: <error>"
	FetchGroupsFailedErrMsg = "error fetching quota: <error>"

	// Get Group Error Response Msgs
	// Associated APIs:
	// GET:
	// - /groups/:id
	GroupNotFoundErrMsg     = "group not found"
	GetGroupFailedErrMsg    = "could not get groups"
	QuotaNotFoundErrMsg     = "quota not found for id: <id>, err: <error>"
	ErrorGettingRequestsMsg = "error getting request: <error>"

	// Get All Keys Error Response Msgs
	// Associated APIs:
	// GET:
	// - /keys
	GetKeysFailedErrMsg   = "error getting keys: <error>"
	FetchKeysFailedErrMsg = "error fetching keys: <error>"

	// Get Key Error Response Msgs
	// Associated APIs:
	// GET:
	// - /keys/:id
	InvalidKeyErrorMsg = "invalid id: <id>"
	KeyNotFoundErrMsg = "key not found with id: <id>"
	ErrorGettingKeyErrMsg = "error getting key: <error>"

	// Create Key Error Response Msgs
	// Associated APIs:
	// POST:
	// - /keys
	EmptyContentErrMsg = "Content cannot be empty."
	InvalidSshKeyErrMsg = "Invalid ssh key"
	InvalidNameErrMsg = "Name must be 32 characters and cannot empty."
	DBInsertFailedErrMsg = "failed to insert the key into the db, err: <error>"

	// Delete Key Error Response Msgs
	// Associated APIs:
	// DELETE:
	// - /keys/:id
	FetchDBRecordFailedErrMsg = "failed to fetch the requested record from the db, err: <error>"
	DeleteNotAllowedErrMsg = "You do not have permission to delete this key."
	DeleteKeyInDBFailedErrMsg = "failed to delete the key from the db, err: <error>"

	// Get Quota Error Response Msgs
	// Associated APIs:
	// GET:
	// - /groups/:id/quota
	GroupIDNotExistErrMsg = "The group ID <group-id> does not exist."
	GetQuotaFailedErrMsg = "An error occured while retriving quota, contact PAC support. Error: <error>"
	QuotaPolicyNotExistErrMsg = "A quota policy does not exist for this group ID. You need to create one first."

	// Create Quota Error Response Msgs
	// Associated APIs:
	// POST:
	// - /groups/:id/quota
	InvalidGroupIDErrMsg = "GroupID must not be set in the request body, or must match the one set in request path."
	InvalidCapacityErrMsg = "minimum supported values for CPU and memory capacity on PowerVS is 0.25C and 2GB respectively"
	InavlidCPUCoresErrMsg = "the CPU cores that can be provisoned on PowerVC is multiples of 0.25"
	CreateQuotaFailedErrMsg = "An error occured while creating quota, contact PAC support."
	QuotaPolicyExistErrMsg = "A quota policy already exists for this group ID. You may delete or update the existing quota."
	InsertQuotaFailedErrMsg = "Failed to insert the quota into the database, Error: <error>"

	// Update Quota Error Response Msgs
	// Associated APIs:
	// PUT:
	// - /groups/:id/quota
	QuotaPolicyNotExistErrMsg = "A quota policy does not exist for this group ID. You need to create one first."

	// Delete Quota Error Response Msgs
	// Associated APIs:
	// DELETE:
	// - /groups/:id/quota
	DeleteQuotaFailedErrMsg = "error": "Error while deleting quota"

	// Delete Quota Error Response Msgs
	// Associated APIs:
	// GET:
	// - /quota
	GetQuotaFailedErrMsg = "failed to get quota, err: <error>"
	GetUsedQuotaFailedErrMsg = "failed to get used quota <error>"

	// Get All Requests Error Response Msgs
	// Associated APIs:
	// GET:
	// - /requests
	InvalidRequestTypeErrMsg = "Invalid request type - <request-type>"
	GetRequestsFailedErrMsg = "error getting requests: <error>"
	FetchRequestsFailedErrMsg = "error fetching requests: <error>"

	// Get Request Error Response Msgs
	// Associated APIs:
	// GET:
	// - /requests/:id
	InvalidIDErrMsg = "invalid id: <ID>"
	RequestNotFoundErrMsg = "request not found with id: <ID>"

	// Update Service Expiry Error Response Msgs
	// Associated APIs:
	// PUT:
	// - /services/:name/expiry
	JustificationNotSetErrMsg = "justification should be set"
	InvalidJustificationErrMsg = "justification must be 500 characters or less"
	ExpiryNotSetErrMsg = "expiry time should be set"
	InvalidRequestTypeErrMsg = "invalid request_type: \"<request-type>\" is set, valid values are <request-type> <request-type> <request-type>"
	ServiceNotExistErrMsg = "service with name <service-name> does not exist"
	GetServiceFailedErrMsg = "failed to get service with name <service-name> Error: <error>"
	ServiceExpiredErrMsg = "Service <service-name> is expired, can't extend the expiry"
	CatalogRetiredErrMsg = "catalog <catalog-name> is retired, can't extend the expiry"
	GetRequestFailedErrMsg = "failed to fetch the request from the db, err: <error>"
	RequestPresentErrMsg = "You have already requested to extend service expiry"
	InsertRequestFailedErrMsg = "failed to insert the request into the db, err: <error>"

	// New Group Request Error Response Msgs
	// Associated APIs:
	// POST:
	// - /groups/:id/request
	AlreadyMemberErrMsg = "You are already a member of this group."
	FetchRecoredFailedErrMsg = "failed to fetch the requested record from the db, err: <error>"
	AccessRequestPresentErrMsg = "You have already requested access to this group."

	// Exit Group Error Response Msgs
	// Associated APIs:
	// POST:
	// - /groups/:id/exit
	MemberNotInGroupErrMsg = "You are already not a member of this group."
	ExitRequestPresentErrMsg = "You have already requested to exit from this group."

	// Approve Request Error Response Msgs
	// Associated APIs:
	// POST:
	// - /requests/:id/approve
	ApproveNotAllowedErrMsg = "You do not have permission to approve requests."
	UserAddFailedErrMsg = "could not add user to group"
	UserDeleteFailedErrMsg = "could not delete user from group"
	ServiceUpdateFailedErrMsg = "failed to update service with name <service-name> Error: <error>"
	StateUpdateFailedErrMsg = "failed to update the state field in the db, err: <error>"

	// Reject Request Error Response Msgs
	// Associated APIs:
	// POST:
	// - /requests/:id/reject
	RejectNotAllowedErrMsg = "You do not have permission to reject requests."
	InvalidJSONBodyErrMsg = "failed to load the body, please feed the proper json body: <json-body>"
	CommentRequiredErrMsg = "Comment is required to reject a request."
	ApprovedStatusErrMsg = "Request is already <'approved'/'rejected'>."

	// Delete Request Error Response Msgs
	// Associated APIs:
	// DELETE:
	// - /requests/:id
	DeleteNotAllowedErrMsg = "You do not have permission to delete this request."
	DeleteDBRecordFailedErrMsg = "failed to delete the record from the db, err: <error>"

	// Get All Services Error Response Msgs
	// Associated APIs:
	// GET:
	// - /services
	GetServicesFailedErrMsg = "failed to get services Error: <error>"

	// Get Service Error Response Msgs
	// Associated APIs:
	// GET:
	// - /services/:name
	ServiceNameNotSetErrMsg = "error": "serviceName name is not set"
	ServiceOwnershipFailedErrMsg = "user id: <user-id> is not owner of service <service-name>"

	// Create Service Error Response Msgs
	// Associated APIs:
	// POST:
	// - /services
	DisplayNameNotSetErrMsg = "display name should be set"
	CatalogRetiredErrMsg = "catalog <catalog-name> is retired, cannot deploy service"
	CatalogNotReadyErrMsg = "catalog <catalog-name> is not in ready state, cannot deploy service"
	SSHKeysNotFoundErrMsg = "no ssh keys found"
	GetCapacityFailedErrMsg = "failed to get needed capacity <catalog-capacity>"
	InsufficientQuotaErrMsg = "user does not have quota to provision resource, Quota: <quota-value> Required: <capacity> Used: <used-quota>"
	ServiceAlreadyExistErrMsg = "service with name <service-name> already exist"
	CreateServiceFailedErrMsg = "failed to create service Error: <error>"

	// Delete Service Error Response Msgs
	// Associated APIs:
	// DELETE:
	// - /services/:name
	ServiceNameNotSetErrMsg = "service name is not set"
	GetServiceFailedErrMsg = "error getting the service with name <service-name> Error: <error>"
	InvalidServiceUserErrMsg = "user id: <user-id> is not the owner of serivce <service-name>"
	DeleteServiceFailedErrMsg = "failed to delete service with name <service-name> Error: <error>"

	// Get TnC Error Response Msgs
	// Associated APIs:
	// GET:
	// - /tnc
	GetTncFailedErrMsg = "failed to get terms and conditions status: <error>"

	// Accept TnC Error Response Msgs
	// Associated APIs:
	// POST:
	// - /tnc
	TncAlreadyAcceptedErrMsg = "terms and conditions already accepted"
	TncAcceptFailedErrMsg = "failed to accept terms and conditions: <error>"

	// Get Users Error Response Msgs
	// Associated APIs:
	// GET:
	// - /users
	GetUserFailedErrMsg = "could not get users"

	// Get User Error Response Msgs
	// Associated APIs:
	// GET:
	// - /users/:id
	// Same as Above

)
