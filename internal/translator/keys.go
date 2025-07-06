package translator

type Key string

const (
	StartMessage            Key = "START_MESSAGE"
	TrackingList            Key = "TRACKING_LIST_BTN"
	StartTracking           Key = "START_TRACKING_BTN"
	UnknownError            Key = "UNKNOWN_ERROR"
	TrackingWaiting         Key = "TRACKING_WAITING"
	TrackingStarted         Key = "TRACKING_STARTED"
	TrackingCancelled       Key = "TRACKING_CANCELLED"
	TrackingAlreadyAdded    Key = "TRACKING_ALREADY_ADDED"
	TrackingRemovedError    Key = "TRACKING_REMOVED_ERROR"
	TrackingLimitReached    Key = "TRACKING_LIMIT_REACHED"
	BetaOpened              Key = "BETA_OPENED"
	BetaClosed              Key = "BETA_CLOSED"
	JoinBeta                Key = "JOIN_BETA_BTN"
	NewVarErrorUsage        Key = "NEW_VAR_USAGE_ERROR_ADM"
	NewVarAlreadyExists     Key = "NEW_VAR_ALREADY_EXISTS_ADM"
	NewVarSuccess           Key = "NEW_VAR_SUCCESS_ADM"
	RemoveVarErrorUsage     Key = "REMOVE_VAR_USAGE_ERROR_ADM"
	RemoveOrEditVarNotFound Key = "REMOVE_OR_EDIT_VAR_NOT_FOUND_ADM"
	RemoveVarSuccess        Key = "REMOVE_VAR_SUCCESS_ADM"
	EditVarErrorUsage       Key = "EDIT_VAR_USAGE_ERROR_ADM"
	EditVarSuccess          Key = "EDIT_VAR_SUCCESS_ADM"
	SearchVarErrorUsage     Key = "SEARCH_VAR_USAGE_ERROR_ADM"
	SearchVarNotFound       Key = "SEARCH_VAR_NOT_FOUND_ADM"
	SearchVarSuccess        Key = "SEARCH_VAR_SUCCESS_ADM"
	Close                   Key = "CLOSE_BTN"
)
