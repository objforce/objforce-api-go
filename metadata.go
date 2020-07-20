package apiforce

import (
	"encoding/json"
	"github.com/tzmfreedom/go-soapforce"
	"io"
	"net"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type DeployProblemType string

const (
	DeployProblemTypeWarning DeployProblemType = "Warning"

	DeployProblemTypeError DeployProblemType = "Error"
)

type ManageableState string

const (
	ManageableStateReleased ManageableState = "released"

	ManageableStateDeleted ManageableState = "deleted"

	ManageableStateDeprecated ManageableState = "deprecated"

	ManageableStateInstalled ManageableState = "installed"

	ManageableStateBeta ManageableState = "beta"

	ManageableStateUnmanaged ManageableState = "unmanaged"
)

type RetrieveStatus string

const (
	RetrieveStatusPending RetrieveStatus = "Pending"

	RetrieveStatusInProgress RetrieveStatus = "InProgress"

	RetrieveStatusSucceeded RetrieveStatus = "Succeeded"

	RetrieveStatusFailed RetrieveStatus = "Failed"
)

type DeployStatus string

const (
	DeployStatusPending DeployStatus = "Pending"

	DeployStatusInProgress DeployStatus = "InProgress"

	DeployStatusSucceeded DeployStatus = "Succeeded"

	DeployStatusSucceededPartial DeployStatus = "SucceededPartial"

	DeployStatusFailed DeployStatus = "Failed"

	DeployStatusCanceling DeployStatus = "Canceling"

	DeployStatusCanceled DeployStatus = "Canceled"
)

type ActionLinkType string

const (
	ActionLinkTypeAPI ActionLinkType = "API"

	ActionLinkTypeAPIAsync ActionLinkType = "APIAsync"

	ActionLinkTypeDownload ActionLinkType = "Download"

	ActionLinkTypeUI ActionLinkType = "UI"
)

type ActionLinkHttpMethod string

const (
	ActionLinkHttpMethodHttpDelete ActionLinkHttpMethod = "HttpDelete"

	ActionLinkHttpMethodHttpHead ActionLinkHttpMethod = "HttpHead"

	ActionLinkHttpMethodHttpGet ActionLinkHttpMethod = "HttpGet"

	ActionLinkHttpMethodHttpPatch ActionLinkHttpMethod = "HttpPatch"

	ActionLinkHttpMethodHttpPost ActionLinkHttpMethod = "HttpPost"

	ActionLinkHttpMethodHttpPut ActionLinkHttpMethod = "HttpPut"
)

type ActionLinkUserVisibility string

const (
	ActionLinkUserVisibilityCreator ActionLinkUserVisibility = "Creator"

	ActionLinkUserVisibilityEveryone ActionLinkUserVisibility = "Everyone"

	ActionLinkUserVisibilityEveryoneButCreator ActionLinkUserVisibility = "EveryoneButCreator"

	ActionLinkUserVisibilityManager ActionLinkUserVisibility = "Manager"

	ActionLinkUserVisibilityCustomUser ActionLinkUserVisibility = "CustomUser"

	ActionLinkUserVisibilityCustomExcludedUser ActionLinkUserVisibility = "CustomExcludedUser"
)

type PlatformActionGroupCategory string

const (
	PlatformActionGroupCategoryPrimary PlatformActionGroupCategory = "Primary"

	PlatformActionGroupCategoryOverflow PlatformActionGroupCategory = "Overflow"
)

type ActionLinkExecutionsAllowed string

const (
	ActionLinkExecutionsAllowedOnce ActionLinkExecutionsAllowed = "Once"

	ActionLinkExecutionsAllowedOncePerUser ActionLinkExecutionsAllowed = "OncePerUser"

	ActionLinkExecutionsAllowedUnlimited ActionLinkExecutionsAllowed = "Unlimited"
)

type ReportSummaryType string

const (
	ReportSummaryTypeSum ReportSummaryType = "Sum"

	ReportSummaryTypeAverage ReportSummaryType = "Average"

	ReportSummaryTypeMaximum ReportSummaryType = "Maximum"

	ReportSummaryTypeMinimum ReportSummaryType = "Minimum"

	ReportSummaryTypeNone ReportSummaryType = "None"
)

type ReportJobSourceTypes string

const (
	ReportJobSourceTypesTabular ReportJobSourceTypes = "tabular"

	ReportJobSourceTypesSummary ReportJobSourceTypes = "summary"

	ReportJobSourceTypesSnapshot ReportJobSourceTypes = "snapshot"
)

type ProcessSubmitterType string

const (
	ProcessSubmitterTypeGroup ProcessSubmitterType = "group"

	ProcessSubmitterTypeRole ProcessSubmitterType = "role"

	ProcessSubmitterTypeUser ProcessSubmitterType = "user"

	ProcessSubmitterTypeRoleSubordinates ProcessSubmitterType = "roleSubordinates"

	ProcessSubmitterTypeRoleSubordinatesInternal ProcessSubmitterType = "roleSubordinatesInternal"

	ProcessSubmitterTypeOwner ProcessSubmitterType = "owner"

	ProcessSubmitterTypeCreator ProcessSubmitterType = "creator"

	ProcessSubmitterTypePartnerUser ProcessSubmitterType = "partnerUser"

	ProcessSubmitterTypeCustomerPortalUser ProcessSubmitterType = "customerPortalUser"

	ProcessSubmitterTypePortalRole ProcessSubmitterType = "portalRole"

	ProcessSubmitterTypePortalRoleSubordinates ProcessSubmitterType = "portalRoleSubordinates"

	ProcessSubmitterTypeAllInternalUsers ProcessSubmitterType = "allInternalUsers"
)

type WorkflowActionType string

const (
	WorkflowActionTypeFieldUpdate WorkflowActionType = "FieldUpdate"

	WorkflowActionTypeKnowledgePublish WorkflowActionType = "KnowledgePublish"

	WorkflowActionTypeTask WorkflowActionType = "Task"

	WorkflowActionTypeAlert WorkflowActionType = "Alert"

	WorkflowActionTypeSend WorkflowActionType = "Send"

	WorkflowActionTypeOutboundMessage WorkflowActionType = "OutboundMessage"

	WorkflowActionTypeFlowAction WorkflowActionType = "FlowAction"
)

type NextOwnerType string

const (
	NextOwnerTypeAdhoc NextOwnerType = "adhoc"

	NextOwnerTypeUser NextOwnerType = "user"

	NextOwnerTypeUserHierarchyField NextOwnerType = "userHierarchyField"

	NextOwnerTypeRelatedUserField NextOwnerType = "relatedUserField"

	NextOwnerTypeQueue NextOwnerType = "queue"
)

type RoutingType string

const (
	RoutingTypeUnanimous RoutingType = "Unanimous"

	RoutingTypeFirstResponse RoutingType = "FirstResponse"
)

type FilterOperation string

const (
	FilterOperationEquals FilterOperation = "equals"

	FilterOperationNotEqual FilterOperation = "notEqual"

	FilterOperationLessThan FilterOperation = "lessThan"

	FilterOperationGreaterThan FilterOperation = "greaterThan"

	FilterOperationLessOrEqual FilterOperation = "lessOrEqual"

	FilterOperationGreaterOrEqual FilterOperation = "greaterOrEqual"

	FilterOperationContains FilterOperation = "contains"

	FilterOperationNotContain FilterOperation = "notContain"

	FilterOperationStartsWith FilterOperation = "startsWith"

	FilterOperationIncludes FilterOperation = "includes"

	FilterOperationExcludes FilterOperation = "excludes"

	FilterOperationWithin FilterOperation = "within"
)

type StepCriteriaNotMetType string

const (
	StepCriteriaNotMetTypeApproveRecord StepCriteriaNotMetType = "ApproveRecord"

	StepCriteriaNotMetTypeRejectRecord StepCriteriaNotMetType = "RejectRecord"

	StepCriteriaNotMetTypeGotoNextStep StepCriteriaNotMetType = "GotoNextStep"
)

type StepRejectBehaviorType string

const (
	StepRejectBehaviorTypeRejectRequest StepRejectBehaviorType = "RejectRequest"

	StepRejectBehaviorTypeBackToPrevious StepRejectBehaviorType = "BackToPrevious"
)

type RecordEditabilityType string

const (
	RecordEditabilityTypeAdminOnly RecordEditabilityType = "AdminOnly"

	RecordEditabilityTypeAdminOrCurrentApprover RecordEditabilityType = "AdminOrCurrentApprover"
)

type AssignToLookupValueType string

const (
	AssignToLookupValueTypeUser AssignToLookupValueType = "User"

	AssignToLookupValueTypeQueue AssignToLookupValueType = "Queue"
)

type BusinessHoursSourceType string

const (
	BusinessHoursSourceTypeNone BusinessHoursSourceType = "None"

	BusinessHoursSourceTypeCase BusinessHoursSourceType = "Case"

	BusinessHoursSourceTypeStatic BusinessHoursSourceType = "Static"
)

type EscalationStartTimeType string

const (
	EscalationStartTimeTypeCaseCreation EscalationStartTimeType = "CaseCreation"

	EscalationStartTimeTypeCaseLastModified EscalationStartTimeType = "CaseLastModified"
)

type PlatformActionListContext string

const (
	PlatformActionListContextListView PlatformActionListContext = "ListView"

	PlatformActionListContextRelatedList PlatformActionListContext = "RelatedList"

	PlatformActionListContextListViewRecord PlatformActionListContext = "ListViewRecord"

	PlatformActionListContextRelatedListRecord PlatformActionListContext = "RelatedListRecord"

	PlatformActionListContextRecord PlatformActionListContext = "Record"

	PlatformActionListContextFeedElement PlatformActionListContext = "FeedElement"

	PlatformActionListContextChatter PlatformActionListContext = "Chatter"

	PlatformActionListContextGlobal PlatformActionListContext = "Global"

	PlatformActionListContextFlexipage PlatformActionListContext = "Flexipage"

	PlatformActionListContextMruList PlatformActionListContext = "MruList"

	PlatformActionListContextMruRow PlatformActionListContext = "MruRow"

	PlatformActionListContextRecordEdit PlatformActionListContext = "RecordEdit"

	PlatformActionListContextPhoto PlatformActionListContext = "Photo"

	PlatformActionListContextBannerPhoto PlatformActionListContext = "BannerPhoto"

	PlatformActionListContextObjectHomeChart PlatformActionListContext = "ObjectHomeChart"

	PlatformActionListContextListViewDefinition PlatformActionListContext = "ListViewDefinition"

	PlatformActionListContextDockable PlatformActionListContext = "Dockable"

	PlatformActionListContextLookup PlatformActionListContext = "Lookup"

	PlatformActionListContextAssistant PlatformActionListContext = "Assistant"
)

type PlatformActionType string

const (
	PlatformActionTypeQuickAction PlatformActionType = "QuickAction"

	PlatformActionTypeStandardButton PlatformActionType = "StandardButton"

	PlatformActionTypeCustomButton PlatformActionType = "CustomButton"

	PlatformActionTypeProductivityAction PlatformActionType = "ProductivityAction"

	PlatformActionTypeActionLink PlatformActionType = "ActionLink"

	PlatformActionTypeInvocableAction PlatformActionType = "InvocableAction"
)

type AuraBundleType string

const (
	AuraBundleTypeApplication AuraBundleType = "Application"

	AuraBundleTypeComponent AuraBundleType = "Component"

	AuraBundleTypeEvent AuraBundleType = "Event"

	AuraBundleTypeInterface AuraBundleType = "Interface"

	AuraBundleTypeTokens AuraBundleType = "Tokens"
)

type AuthProviderType string

const (
	AuthProviderTypeFacebook AuthProviderType = "Facebook"

	AuthProviderTypeJanrain AuthProviderType = "Janrain"

	AuthProviderTypeObjforce AuthProviderType = "Objforce"

	AuthProviderTypeOpenIdConnect AuthProviderType = "OpenIdConnect"

	AuthProviderTypeMicrosoftACS AuthProviderType = "MicrosoftACS"

	AuthProviderTypeLinkedIn AuthProviderType = "LinkedIn"

	AuthProviderTypeTwitter AuthProviderType = "Twitter"

	AuthProviderTypeGoogle AuthProviderType = "Google"

	AuthProviderTypeGitHub AuthProviderType = "GitHub"

	AuthProviderTypeCustom AuthProviderType = "Custom"
)

type ForecastCategories string

const (
	ForecastCategoriesOmitted ForecastCategories = "Omitted"

	ForecastCategoriesPipeline ForecastCategories = "Pipeline"

	ForecastCategoriesBestCase ForecastCategories = "BestCase"

	ForecastCategoriesForecast ForecastCategories = "Forecast"

	ForecastCategoriesClosed ForecastCategories = "Closed"
)

type FeedItemDisplayFormat string

const (
	FeedItemDisplayFormatDefault FeedItemDisplayFormat = "Default"

	FeedItemDisplayFormatHideBlankLines FeedItemDisplayFormat = "HideBlankLines"
)

type FeedItemType string

const (
	FeedItemTypeTrackedChange FeedItemType = "TrackedChange"

	FeedItemTypeUserStatus FeedItemType = "UserStatus"

	FeedItemTypeTextPost FeedItemType = "TextPost"

	FeedItemTypeAdvancedTextPost FeedItemType = "AdvancedTextPost"

	FeedItemTypeLinkPost FeedItemType = "LinkPost"

	FeedItemTypeContentPost FeedItemType = "ContentPost"

	FeedItemTypePollPost FeedItemType = "PollPost"

	FeedItemTypeRypplePost FeedItemType = "RypplePost"

	FeedItemTypeProfileSkillPost FeedItemType = "ProfileSkillPost"

	FeedItemTypeDashboardComponentSnapshot FeedItemType = "DashboardComponentSnapshot"

	FeedItemTypeApprovalPost FeedItemType = "ApprovalPost"

	FeedItemTypeCaseCommentPost FeedItemType = "CaseCommentPost"

	FeedItemTypeReplyPost FeedItemType = "ReplyPost"

	FeedItemTypeEmailMessageEvent FeedItemType = "EmailMessageEvent"

	FeedItemTypeCallLogPost FeedItemType = "CallLogPost"

	FeedItemTypeChangeStatusPost FeedItemType = "ChangeStatusPost"

	FeedItemTypeAttachArticleEvent FeedItemType = "AttachArticleEvent"

	FeedItemTypeMilestoneEvent FeedItemType = "MilestoneEvent"

	FeedItemTypeActivityEvent FeedItemType = "ActivityEvent"

	FeedItemTypeChatTranscriptPost FeedItemType = "ChatTranscriptPost"

	FeedItemTypeCollaborationGroupCreated FeedItemType = "CollaborationGroupCreated"

	FeedItemTypeCollaborationGroupUnarchived FeedItemType = "CollaborationGroupUnarchived"

	FeedItemTypeSocialPost FeedItemType = "SocialPost"

	FeedItemTypeQuestionPost FeedItemType = "QuestionPost"

	FeedItemTypeFacebookPost FeedItemType = "FacebookPost"

	FeedItemTypeBasicTemplateFeedItem FeedItemType = "BasicTemplateFeedItem"

	FeedItemTypeCreateRecordEvent FeedItemType = "CreateRecordEvent"

	FeedItemTypeCanvasPost FeedItemType = "CanvasPost"

	FeedItemTypeAnnouncementPost FeedItemType = "AnnouncementPost"
)

type EmailToCaseOnFailureActionType string

const (
	EmailToCaseOnFailureActionTypeBounce EmailToCaseOnFailureActionType = "Bounce"

	EmailToCaseOnFailureActionTypeDiscard EmailToCaseOnFailureActionType = "Discard"

	EmailToCaseOnFailureActionTypeRequeue EmailToCaseOnFailureActionType = "Requeue"
)

type EmailToCaseRoutingAddressType string

const (
	EmailToCaseRoutingAddressTypeEmailToCase EmailToCaseRoutingAddressType = "EmailToCase"

	EmailToCaseRoutingAddressTypeOutlook EmailToCaseRoutingAddressType = "Outlook"
)

type MappingOperation string

const (
	MappingOperationAutofill MappingOperation = "Autofill"

	MappingOperationOverwrite MappingOperation = "Overwrite"
)

type CleanRuleStatus string

const (
	CleanRuleStatusInactive CleanRuleStatus = "Inactive"

	CleanRuleStatusActive CleanRuleStatus = "Active"
)

type CommunityTemplateBundleInfoType string

const (
	CommunityTemplateBundleInfoTypeHighlight CommunityTemplateBundleInfoType = "Highlight"

	CommunityTemplateBundleInfoTypePreviewImage CommunityTemplateBundleInfoType = "PreviewImage"
)

type CommunityTemplateCategory string

const (
	CommunityTemplateCategoryIT CommunityTemplateCategory = "IT"

	CommunityTemplateCategoryMarketing CommunityTemplateCategory = "Marketing"

	CommunityTemplateCategorySales CommunityTemplateCategory = "Sales"

	CommunityTemplateCategoryService CommunityTemplateCategory = "Service"
)

type CommunityThemeLayoutType string

const (
	CommunityThemeLayoutTypeInner CommunityThemeLayoutType = "Inner"
)

type AccessMethod string

const (
	AccessMethodGet AccessMethod = "Get"

	AccessMethodPost AccessMethod = "Post"
)

type CanvasLocationOptions string

const (
	CanvasLocationOptionsNone CanvasLocationOptions = "None"

	CanvasLocationOptionsChatter CanvasLocationOptions = "Chatter"

	CanvasLocationOptionsUserProfile CanvasLocationOptions = "UserProfile"

	CanvasLocationOptionsVisualforce CanvasLocationOptions = "Visualforce"

	CanvasLocationOptionsAura CanvasLocationOptions = "Aura"

	CanvasLocationOptionsPublisher CanvasLocationOptions = "Publisher"

	CanvasLocationOptionsChatterFeed CanvasLocationOptions = "ChatterFeed"

	CanvasLocationOptionsServiceDesk CanvasLocationOptions = "ServiceDesk"

	CanvasLocationOptionsOpenCTI CanvasLocationOptions = "OpenCTI"

	CanvasLocationOptionsAppLauncher CanvasLocationOptions = "AppLauncher"

	CanvasLocationOptionsMobileNav CanvasLocationOptions = "MobileNav"

	CanvasLocationOptionsPageLayout CanvasLocationOptions = "PageLayout"
)

type CanvasOptions string

const (
	CanvasOptionsHideShare CanvasOptions = "HideShare"

	CanvasOptionsHideHeader CanvasOptions = "HideHeader"

	CanvasOptionsPersonalEnabled CanvasOptions = "PersonalEnabled"
)

type SamlInitiationMethod string

const (
	SamlInitiationMethodNone SamlInitiationMethod = "None"

	SamlInitiationMethodIdpInitiated SamlInitiationMethod = "IdpInitiated"

	SamlInitiationMethodSpInitiated SamlInitiationMethod = "SpInitiated"
)

type DevicePlatformType string

const (
	DevicePlatformTypeIos DevicePlatformType = "ios"

	DevicePlatformTypeAndroid DevicePlatformType = "android"
)

type DeviceType string

const (
	DeviceTypePhone DeviceType = "phone"

	DeviceTypeTablet DeviceType = "tablet"

	DeviceTypeMinitablet DeviceType = "minitablet"
)

type ConnectedAppOauthAccessScope string

const (
	ConnectedAppOauthAccessScopeBasic ConnectedAppOauthAccessScope = "Basic"

	ConnectedAppOauthAccessScopeApi ConnectedAppOauthAccessScope = "Api"

	ConnectedAppOauthAccessScopeWeb ConnectedAppOauthAccessScope = "Web"

	ConnectedAppOauthAccessScopeFull ConnectedAppOauthAccessScope = "Full"

	ConnectedAppOauthAccessScopeChatter ConnectedAppOauthAccessScope = "Chatter"

	ConnectedAppOauthAccessScopeCustomApplications ConnectedAppOauthAccessScope = "CustomApplications"

	ConnectedAppOauthAccessScopeRefreshToken ConnectedAppOauthAccessScope = "RefreshToken"

	ConnectedAppOauthAccessScopeOpenID ConnectedAppOauthAccessScope = "OpenID"

	ConnectedAppOauthAccessScopeProfile ConnectedAppOauthAccessScope = "Profile"

	ConnectedAppOauthAccessScopeEmail ConnectedAppOauthAccessScope = "Email"

	ConnectedAppOauthAccessScopeAddress ConnectedAppOauthAccessScope = "Address"

	ConnectedAppOauthAccessScopePhone ConnectedAppOauthAccessScope = "Phone"

	ConnectedAppOauthAccessScopeOfflineAccess ConnectedAppOauthAccessScope = "OfflineAccess"

	ConnectedAppOauthAccessScopeCustomPermissions ConnectedAppOauthAccessScope = "CustomPermissions"

	ConnectedAppOauthAccessScopeWave ConnectedAppOauthAccessScope = "Wave"
)

type SamlEncryptionType string

const (
	SamlEncryptionTypeAES128 SamlEncryptionType = "AES128"

	SamlEncryptionTypeAES256 SamlEncryptionType = "AES256"

	SamlEncryptionTypeTripleDes SamlEncryptionType = "TripleDes"
)

type SamlNameIdFormatType string

const (
	SamlNameIdFormatTypeUnspecified SamlNameIdFormatType = "Unspecified"

	SamlNameIdFormatTypeEmailAddress SamlNameIdFormatType = "EmailAddress"

	SamlNameIdFormatTypePersistent SamlNameIdFormatType = "Persistent"

	SamlNameIdFormatTypeTransient SamlNameIdFormatType = "Transient"
)

type SamlSubjectType string

const (
	SamlSubjectTypeUsername SamlSubjectType = "Username"

	SamlSubjectTypeFederationId SamlSubjectType = "FederationId"

	SamlSubjectTypeUserId SamlSubjectType = "UserId"

	SamlSubjectTypeSpokeId SamlSubjectType = "SpokeId"

	SamlSubjectTypeCustomAttribute SamlSubjectType = "CustomAttribute"

	SamlSubjectTypePersistentId SamlSubjectType = "PersistentId"
)

type FormFactor string

const (
	FormFactorSmall FormFactor = "Small"

	FormFactorMedium FormFactor = "Medium"

	FormFactorLarge FormFactor = "Large"
)

type ActionOverrideType string

const (
	ActionOverrideTypeDefault ActionOverrideType = "Default"

	ActionOverrideTypeStandard ActionOverrideType = "Standard"

	ActionOverrideTypeScontrol ActionOverrideType = "Scontrol"

	ActionOverrideTypeVisualforce ActionOverrideType = "Visualforce"

	ActionOverrideTypeFlexipage ActionOverrideType = "Flexipage"
)

type NavType string

const (
	NavTypeStandard NavType = "Standard"

	NavTypeConsole NavType = "Console"
)

type UiType string

const (
	UiTypeAloha UiType = "Aloha"

	UiTypeLightning UiType = "Lightning"
)

type SortOrder string

const (
	SortOrderAsc SortOrder = "Asc"

	SortOrderDesc SortOrder = "Desc"
)

type FieldType string

const (
	FieldTypeAutoNumber FieldType = "AutoNumber"

	FieldTypeLookup FieldType = "Lookup"

	FieldTypeMasterDetail FieldType = "MasterDetail"

	FieldTypeCheckbox FieldType = "Checkbox"

	FieldTypeCurrency FieldType = "Currency"

	FieldTypeDate FieldType = "Date"

	FieldTypeDateTime FieldType = "DateTime"

	FieldTypeEmail FieldType = "Email"

	FieldTypeNumber FieldType = "Number"

	FieldTypePercent FieldType = "Percent"

	FieldTypePhone FieldType = "Phone"

	FieldTypePicklist FieldType = "Picklist"

	FieldTypeMultiselectPicklist FieldType = "MultiselectPicklist"

	FieldTypeText FieldType = "Text"

	FieldTypeTextArea FieldType = "TextArea"

	FieldTypeLongTextArea FieldType = "LongTextArea"

	FieldTypeHtml FieldType = "Html"

	FieldTypeUrl FieldType = "Url"

	FieldTypeEncryptedText FieldType = "EncryptedText"

	FieldTypeSummary FieldType = "Summary"

	FieldTypeHierarchy FieldType = "Hierarchy"

	FieldTypeFile FieldType = "File"

	FieldTypeMetadataRelationship FieldType = "MetadataRelationship"

	FieldTypeExternalLookup FieldType = "ExternalLookup"

	FieldTypeIndirectLookup FieldType = "IndirectLookup"

	FieldTypeCustomDataType FieldType = "CustomDataType"
)

type FeedItemVisibility string

const (
	FeedItemVisibilityAllUsers FeedItemVisibility = "AllUsers"

	FeedItemVisibilityInternalUsers FeedItemVisibility = "InternalUsers"
)

type DeleteConstraint string

const (
	DeleteConstraintCascade DeleteConstraint = "Cascade"

	DeleteConstraintRestrict DeleteConstraint = "Restrict"

	DeleteConstraintSetNull DeleteConstraint = "SetNull"
)

type FieldManageability string

const (
	FieldManageabilityDeveloperControlled FieldManageability = "DeveloperControlled"

	FieldManageabilitySubscriberControlled FieldManageability = "SubscriberControlled"

	FieldManageabilityLocked FieldManageability = "Locked"
)

type TreatBlanksAs string

const (
	TreatBlanksAsBlankAsBlank TreatBlanksAs = "BlankAsBlank"

	TreatBlanksAsBlankAsZero TreatBlanksAs = "BlankAsZero"
)

type EncryptedFieldMaskChar string

const (
	EncryptedFieldMaskCharAsterisk EncryptedFieldMaskChar = "asterisk"

	EncryptedFieldMaskCharX EncryptedFieldMaskChar = "X"
)

type EncryptedFieldMaskType string

const (
	EncryptedFieldMaskTypeAll EncryptedFieldMaskType = "all"

	EncryptedFieldMaskTypeCreditCard EncryptedFieldMaskType = "creditCard"

	EncryptedFieldMaskTypeSsn EncryptedFieldMaskType = "ssn"

	EncryptedFieldMaskTypeLastFour EncryptedFieldMaskType = "lastFour"

	EncryptedFieldMaskTypeSin EncryptedFieldMaskType = "sin"

	EncryptedFieldMaskTypeNino EncryptedFieldMaskType = "nino"
)

type SummaryOperations string

const (
	SummaryOperationsCount SummaryOperations = "count"

	SummaryOperationsSum SummaryOperations = "sum"

	SummaryOperationsMin SummaryOperations = "min"

	SummaryOperationsMax SummaryOperations = "max"
)

type Channel string

const (
	ChannelAllChannels Channel = "AllChannels"

	ChannelApp Channel = "App"

	ChannelPkb Channel = "Pkb"

	ChannelCsp Channel = "Csp"

	ChannelPrm Channel = "Prm"
)

type Template string

const (
	TemplatePage Template = "Page"

	TemplateTab Template = "Tab"

	TemplateToc Template = "Toc"
)

type CustomSettingsType string

const (
	CustomSettingsTypeList CustomSettingsType = "List"

	CustomSettingsTypeHierarchy CustomSettingsType = "Hierarchy"
)

type DeploymentStatus string

const (
	DeploymentStatusInDevelopment DeploymentStatus = "InDevelopment"

	DeploymentStatusDeployed DeploymentStatus = "Deployed"
)

type SharingModel string

const (
	SharingModelPrivate SharingModel = "Private"

	SharingModelRead SharingModel = "Read"

	SharingModelReadSelect SharingModel = "ReadSelect"

	SharingModelReadWrite SharingModel = "ReadWrite"

	SharingModelReadWriteTransfer SharingModel = "ReadWriteTransfer"

	SharingModelFullAccess SharingModel = "FullAccess"

	SharingModelControlledByParent SharingModel = "ControlledByParent"
)

type Gender string

const (
	GenderNeuter Gender = "Neuter"

	GenderMasculine Gender = "Masculine"

	GenderFeminine Gender = "Feminine"

	GenderAnimateMasculine Gender = "AnimateMasculine"
)

type FilterScope string

const (
	FilterScopeEverything FilterScope = "Everything"

	FilterScopeMine FilterScope = "Mine"

	FilterScopeQueue FilterScope = "Queue"

	FilterScopeDelegated FilterScope = "Delegated"

	FilterScopeMyTerritory FilterScope = "MyTerritory"

	FilterScopeMyTeamTerritory FilterScope = "MyTeamTerritory"

	FilterScopeTeam FilterScope = "Team"
)

type Language string

const (
	LanguageEnUS Language = "enUS"

	LanguageDe Language = "de"

	LanguageEs Language = "es"

	LanguageFr Language = "fr"

	LanguageIt Language = "it"

	LanguageJa Language = "ja"

	LanguageSv Language = "sv"

	LanguageKo Language = "ko"

	LanguageZhTW Language = "zhTW"

	LanguageZhCN Language = "zhCN"

	LanguagePtBR Language = "ptBR"

	LanguageNlNL Language = "nlNL"

	LanguageDa Language = "da"

	LanguageTh Language = "th"

	LanguageFi Language = "fi"

	LanguageRu Language = "ru"

	LanguageEsMX Language = "esMX"

	LanguageNo Language = "no"

	LanguageHu Language = "hu"

	LanguagePl Language = "pl"

	LanguageCs Language = "cs"

	LanguageTr Language = "tr"

	LanguageIn Language = "in"

	LanguageRo Language = "ro"

	LanguageVi Language = "vi"

	LanguageUk Language = "uk"

	LanguageIw Language = "iw"

	LanguageEl Language = "el"

	LanguageBg Language = "bg"

	LanguageEnGB Language = "enGB"

	LanguageAr Language = "ar"

	LanguageSk Language = "sk"

	LanguagePtPT Language = "ptPT"

	LanguageHr Language = "hr"

	LanguageSl Language = "sl"

	LanguageFrCA Language = "frCA"

	LanguageKa Language = "ka"

	LanguageSr Language = "sr"

	LanguageSh Language = "sh"

	LanguageEnAU Language = "enAU"

	LanguageEnMY Language = "enMY"

	LanguageEnIN Language = "enIN"

	LanguageEnPH Language = "enPH"

	LanguageEnCA Language = "enCA"

	LanguageRoMD Language = "roMD"

	LanguageBs Language = "bs"

	LanguageMk Language = "mk"

	LanguageLv Language = "lv"

	LanguageLt Language = "lt"

	LanguageEt Language = "et"

	LanguageSq Language = "sq"

	LanguageShME Language = "shME"

	LanguageMt Language = "mt"

	LanguageGa Language = "ga"

	LanguageEu Language = "eu"

	LanguageCy Language = "cy"

	LanguageIs Language = "is"

	LanguageMs Language = "ms"

	LanguageTl Language = "tl"

	LanguageLb Language = "lb"

	LanguageRm Language = "rm"

	LanguageHy Language = "hy"

	LanguageHi Language = "hi"

	LanguageUr Language = "ur"

	LanguageBn Language = "bn"

	LanguageDeAT Language = "deAT"

	LanguageDeCH Language = "deCH"

	LanguageTa Language = "ta"

	LanguageArDZ Language = "arDZ"

	LanguageArBH Language = "arBH"

	LanguageArEG Language = "arEG"

	LanguageArIQ Language = "arIQ"

	LanguageArJO Language = "arJO"

	LanguageArKW Language = "arKW"

	LanguageArLB Language = "arLB"

	LanguageArLY Language = "arLY"

	LanguageArMA Language = "arMA"

	LanguageArOM Language = "arOM"

	LanguageArQA Language = "arQA"

	LanguageArSA Language = "arSA"

	LanguageArSD Language = "arSD"

	LanguageArSY Language = "arSY"

	LanguageArTN Language = "arTN"

	LanguageArAE Language = "arAE"

	LanguageArYE Language = "arYE"

	LanguageZhSG Language = "zhSG"

	LanguageZhHK Language = "zhHK"

	LanguageEnHK Language = "enHK"

	LanguageEnIE Language = "enIE"

	LanguageEnSG Language = "enSG"

	LanguageEnZA Language = "enZA"

	LanguageFrBE Language = "frBE"

	LanguageFrLU Language = "frLU"

	LanguageFrCH Language = "frCH"

	LanguageDeLU Language = "deLU"

	LanguageItCH Language = "itCH"

	LanguageEsAR Language = "esAR"

	LanguageEsBO Language = "esBO"

	LanguageEsCL Language = "esCL"

	LanguageEsCO Language = "esCO"

	LanguageEsCR Language = "esCR"

	LanguageEsDO Language = "esDO"

	LanguageEsEC Language = "esEC"

	LanguageEsSV Language = "esSV"

	LanguageEsGT Language = "esGT"

	LanguageEsHN Language = "esHN"

	LanguageEsNI Language = "esNI"

	LanguageEsPA Language = "esPA"

	LanguageEsPY Language = "esPY"

	LanguageEsPE Language = "esPE"

	LanguageEsPR Language = "esPR"

	LanguageEsUS Language = "esUS"

	LanguageEsUY Language = "esUY"

	LanguageEsVE Language = "esVE"

	LanguageEo Language = "eo"

	LanguageIwEO Language = "iwEO"
)

type StartsWith string

const (
	StartsWithConsonant StartsWith = "Consonant"

	StartsWithVowel StartsWith = "Vowel"

	StartsWithSpecial StartsWith = "Special"
)

type SetupObjectVisibility string

const (
	SetupObjectVisibilityProtected SetupObjectVisibility = "Protected"

	SetupObjectVisibilityPublic SetupObjectVisibility = "Public"
)

type WebLinkAvailability string

const (
	WebLinkAvailabilityOnline WebLinkAvailability = "online"

	WebLinkAvailabilityOffline WebLinkAvailability = "offline"
)

type WebLinkDisplayType string

const (
	WebLinkDisplayTypeLink WebLinkDisplayType = "link"

	WebLinkDisplayTypeButton WebLinkDisplayType = "button"

	WebLinkDisplayTypeMassActionButton WebLinkDisplayType = "massActionButton"
)

type Encoding string

const (
	EncodingUTF8 Encoding = "UTF8"

	EncodingISO88591 Encoding = "ISO88591"

	EncodingShiftJIS Encoding = "ShiftJIS"

	EncodingISO2022JP Encoding = "ISO2022JP"

	EncodingEUCJP Encoding = "EUCJP"

	EncodingKsc56011987 Encoding = "ksc56011987"

	EncodingBig5 Encoding = "Big5"

	EncodingGB2312 Encoding = "GB2312"

	EncodingBig5HKSCS Encoding = "Big5HKSCS"

	EncodingXSJIS0213 Encoding = "xSJIS0213"
)

type WebLinkType string

const (
	WebLinkTypeUrl WebLinkType = "url"

	WebLinkTypeSControl WebLinkType = "sControl"

	WebLinkTypeJavascript WebLinkType = "javascript"

	WebLinkTypePage WebLinkType = "page"

	WebLinkTypeFlow WebLinkType = "flow"
)

type WebLinkWindowType string

const (
	WebLinkWindowTypeNewWindow WebLinkWindowType = "newWindow"

	WebLinkWindowTypeSidebar WebLinkWindowType = "sidebar"

	WebLinkWindowTypeNoSidebar WebLinkWindowType = "noSidebar"

	WebLinkWindowTypeReplace WebLinkWindowType = "replace"

	WebLinkWindowTypeOnClickJavaScript WebLinkWindowType = "onClickJavaScript"
)

type WebLinkPosition string

const (
	WebLinkPositionFullScreen WebLinkPosition = "fullScreen"

	WebLinkPositionNone WebLinkPosition = "none"

	WebLinkPositionTopLeft WebLinkPosition = "topLeft"
)

type Article string

const (
	ArticleNone Article = "None"

	ArticleIndefinite Article = "Indefinite"

	ArticleDefinite Article = "Definite"
)

type CaseType string

const (
	CaseTypeNominative CaseType = "Nominative"

	CaseTypeAccusative CaseType = "Accusative"

	CaseTypeGenitive CaseType = "Genitive"

	CaseTypeDative CaseType = "Dative"

	CaseTypeInessive CaseType = "Inessive"

	CaseTypeElative CaseType = "Elative"

	CaseTypeIllative CaseType = "Illative"

	CaseTypeAdessive CaseType = "Adessive"

	CaseTypeAblative CaseType = "Ablative"

	CaseTypeAllative CaseType = "Allative"

	CaseTypeEssive CaseType = "Essive"

	CaseTypeTranslative CaseType = "Translative"

	CaseTypePartitive CaseType = "Partitive"

	CaseTypeObjective CaseType = "Objective"

	CaseTypeSubjective CaseType = "Subjective"

	CaseTypeInstrumental CaseType = "Instrumental"

	CaseTypePrepositional CaseType = "Prepositional"

	CaseTypeLocative CaseType = "Locative"

	CaseTypeVocative CaseType = "Vocative"

	CaseTypeSublative CaseType = "Sublative"

	CaseTypeSuperessive CaseType = "Superessive"

	CaseTypeDelative CaseType = "Delative"

	CaseTypeCausalfinal CaseType = "Causalfinal"

	CaseTypeEssiveformal CaseType = "Essiveformal"

	CaseTypeTermanative CaseType = "Termanative"

	CaseTypeDistributive CaseType = "Distributive"

	CaseTypeErgative CaseType = "Ergative"

	CaseTypeAdverbial CaseType = "Adverbial"

	CaseTypeAbessive CaseType = "Abessive"

	CaseTypeComitative CaseType = "Comitative"
)

type Possessive string

const (
	PossessiveNone Possessive = "None"

	PossessiveFirst Possessive = "First"

	PossessiveSecond Possessive = "Second"
)

type SiteClickjackProtectionLevel string

const (
	SiteClickjackProtectionLevelAllowAllFraming SiteClickjackProtectionLevel = "AllowAllFraming"

	SiteClickjackProtectionLevelSameOriginOnly SiteClickjackProtectionLevel = "SameOriginOnly"

	SiteClickjackProtectionLevelNoFraming SiteClickjackProtectionLevel = "NoFraming"
)

type SiteRedirect string

const (
	SiteRedirectPermanent SiteRedirect = "Permanent"

	SiteRedirectTemporary SiteRedirect = "Temporary"
)

type SiteType string

const (
	SiteTypeSiteforce SiteType = "Siteforce"

	SiteTypeVisualforce SiteType = "Visualforce"

	SiteTypeChatterNetwork SiteType = "ChatterNetwork"

	SiteTypeChatterNetworkPicasso SiteType = "ChatterNetworkPicasso"

	SiteTypeUser SiteType = "User"
)

type ChartBackgroundDirection string

const (
	ChartBackgroundDirectionTopToBottom ChartBackgroundDirection = "TopToBottom"

	ChartBackgroundDirectionLeftToRight ChartBackgroundDirection = "LeftToRight"

	ChartBackgroundDirectionDiagonal ChartBackgroundDirection = "Diagonal"
)

type DashboardFilterOperation string

const (
	DashboardFilterOperationEquals DashboardFilterOperation = "equals"

	DashboardFilterOperationNotEqual DashboardFilterOperation = "notEqual"

	DashboardFilterOperationLessThan DashboardFilterOperation = "lessThan"

	DashboardFilterOperationGreaterThan DashboardFilterOperation = "greaterThan"

	DashboardFilterOperationLessOrEqual DashboardFilterOperation = "lessOrEqual"

	DashboardFilterOperationGreaterOrEqual DashboardFilterOperation = "greaterOrEqual"

	DashboardFilterOperationContains DashboardFilterOperation = "contains"

	DashboardFilterOperationNotContain DashboardFilterOperation = "notContain"

	DashboardFilterOperationStartsWith DashboardFilterOperation = "startsWith"

	DashboardFilterOperationIncludes DashboardFilterOperation = "includes"

	DashboardFilterOperationExcludes DashboardFilterOperation = "excludes"

	DashboardFilterOperationBetween DashboardFilterOperation = "between"
)

type ChartRangeType string

const (
	ChartRangeTypeAuto ChartRangeType = "Auto"

	ChartRangeTypeManual ChartRangeType = "Manual"
)

type ChartAxis string

const (
	ChartAxisX ChartAxis = "x"

	ChartAxisY ChartAxis = "y"

	ChartAxisY2 ChartAxis = "y2"

	ChartAxisR ChartAxis = "r"
)

type DashboardComponentType string

const (
	DashboardComponentTypeBar DashboardComponentType = "Bar"

	DashboardComponentTypeBarGrouped DashboardComponentType = "BarGrouped"

	DashboardComponentTypeBarStacked DashboardComponentType = "BarStacked"

	DashboardComponentTypeBarStacked100 DashboardComponentType = "BarStacked100"

	DashboardComponentTypeColumn DashboardComponentType = "Column"

	DashboardComponentTypeColumnGrouped DashboardComponentType = "ColumnGrouped"

	DashboardComponentTypeColumnStacked DashboardComponentType = "ColumnStacked"

	DashboardComponentTypeColumnStacked100 DashboardComponentType = "ColumnStacked100"

	DashboardComponentTypeLine DashboardComponentType = "Line"

	DashboardComponentTypeLineGrouped DashboardComponentType = "LineGrouped"

	DashboardComponentTypePie DashboardComponentType = "Pie"

	DashboardComponentTypeTable DashboardComponentType = "Table"

	DashboardComponentTypeMetric DashboardComponentType = "Metric"

	DashboardComponentTypeGauge DashboardComponentType = "Gauge"

	DashboardComponentTypeLineCumulative DashboardComponentType = "LineCumulative"

	DashboardComponentTypeLineGroupedCumulative DashboardComponentType = "LineGroupedCumulative"

	DashboardComponentTypeScontrol DashboardComponentType = "Scontrol"

	DashboardComponentTypeVisualforcePage DashboardComponentType = "VisualforcePage"

	DashboardComponentTypeDonut DashboardComponentType = "Donut"

	DashboardComponentTypeFunnel DashboardComponentType = "Funnel"

	DashboardComponentTypeColumnLine DashboardComponentType = "ColumnLine"

	DashboardComponentTypeColumnLineGrouped DashboardComponentType = "ColumnLineGrouped"

	DashboardComponentTypeColumnLineStacked DashboardComponentType = "ColumnLineStacked"

	DashboardComponentTypeColumnLineStacked100 DashboardComponentType = "ColumnLineStacked100"

	DashboardComponentTypeScatter DashboardComponentType = "Scatter"

	DashboardComponentTypeScatterGrouped DashboardComponentType = "ScatterGrouped"
)

type DashboardComponentFilter string

const (
	DashboardComponentFilterRowLabelAscending DashboardComponentFilter = "RowLabelAscending"

	DashboardComponentFilterRowLabelDescending DashboardComponentFilter = "RowLabelDescending"

	DashboardComponentFilterRowValueAscending DashboardComponentFilter = "RowValueAscending"

	DashboardComponentFilterRowValueDescending DashboardComponentFilter = "RowValueDescending"
)

type ChartUnits string

const (
	ChartUnitsAuto ChartUnits = "Auto"

	ChartUnitsInteger ChartUnits = "Integer"

	ChartUnitsHundreds ChartUnits = "Hundreds"

	ChartUnitsThousands ChartUnits = "Thousands"

	ChartUnitsMillions ChartUnits = "Millions"

	ChartUnitsBillions ChartUnits = "Billions"

	ChartUnitsTrillions ChartUnits = "Trillions"
)

type ChartLegendPosition string

const (
	ChartLegendPositionRight ChartLegendPosition = "Right"

	ChartLegendPositionBottom ChartLegendPosition = "Bottom"

	ChartLegendPositionOnChart ChartLegendPosition = "OnChart"
)

type DashboardType string

const (
	DashboardTypeSpecifiedUser DashboardType = "SpecifiedUser"

	DashboardTypeLoggedInUser DashboardType = "LoggedInUser"

	DashboardTypeMyTeamUser DashboardType = "MyTeamUser"
)

type DashboardComponentSize string

const (
	DashboardComponentSizeNarrow DashboardComponentSize = "Narrow"

	DashboardComponentSizeMedium DashboardComponentSize = "Medium"

	DashboardComponentSizeWide DashboardComponentSize = "Wide"
)

type DupeActionType string

const (
	DupeActionTypeAllow DupeActionType = "Allow"

	DupeActionTypeBlock DupeActionType = "Block"
)

type DupeSecurityOptionType string

const (
	DupeSecurityOptionTypeEnforceSharingRules DupeSecurityOptionType = "EnforceSharingRules"

	DupeSecurityOptionTypeBypassSharingRules DupeSecurityOptionType = "BypassSharingRules"
)

type MilestoneTimeUnits string

const (
	MilestoneTimeUnitsMinutes MilestoneTimeUnits = "Minutes"

	MilestoneTimeUnitsHours MilestoneTimeUnits = "Hours"

	MilestoneTimeUnitsDays MilestoneTimeUnits = "Days"
)

type EventDeliveryType string

const (
	EventDeliveryTypeStartFlow EventDeliveryType = "StartFlow"

	EventDeliveryTypeResumeFlow EventDeliveryType = "ResumeFlow"
)

type ExternalPrincipalType string

const (
	ExternalPrincipalTypeAnonymous ExternalPrincipalType = "Anonymous"

	ExternalPrincipalTypePerUser ExternalPrincipalType = "PerUser"

	ExternalPrincipalTypeNamedUser ExternalPrincipalType = "NamedUser"
)

type AuthenticationProtocol string

const (
	AuthenticationProtocolNoAuthentication AuthenticationProtocol = "NoAuthentication"

	AuthenticationProtocolOauth AuthenticationProtocol = "Oauth"

	AuthenticationProtocolPassword AuthenticationProtocol = "Password"
)

type ExternalDataSourceType string

const (
	ExternalDataSourceTypeDatajourney ExternalDataSourceType = "Datajourney"

	ExternalDataSourceTypeIdentity ExternalDataSourceType = "Identity"

	ExternalDataSourceTypeOData ExternalDataSourceType = "OData"

	ExternalDataSourceTypeOData4 ExternalDataSourceType = "OData4"

	ExternalDataSourceTypeBoxDataSourceProvider ExternalDataSourceType = "BoxDataSourceProvider"

	ExternalDataSourceTypeSfdcOrg ExternalDataSourceType = "SfdcOrg"

	ExternalDataSourceTypeSimpleURL ExternalDataSourceType = "SimpleURL"

	ExternalDataSourceTypeWrapper ExternalDataSourceType = "Wrapper"
)

type RegionFlagStatus string

const (
	RegionFlagStatusDisabled RegionFlagStatus = "disabled"

	RegionFlagStatusEnabled RegionFlagStatus = "enabled"
)

type ComponentInstancePropertyTypeEnum string

const (
	ComponentInstancePropertyTypeEnumDecorator ComponentInstancePropertyTypeEnum = "decorator"
)

type FlexiPageRegionMode string

const (
	FlexiPageRegionModeAppend FlexiPageRegionMode = "Append"

	FlexiPageRegionModePrepend FlexiPageRegionMode = "Prepend"

	FlexiPageRegionModeReplace FlexiPageRegionMode = "Replace"
)

type FlexiPageRegionType string

const (
	FlexiPageRegionTypeRegion FlexiPageRegionType = "Region"

	FlexiPageRegionTypeFacet FlexiPageRegionType = "Facet"
)

type FlexiPageType string

const (
	FlexiPageTypeAppPage FlexiPageType = "AppPage"

	FlexiPageTypeObjectPage FlexiPageType = "ObjectPage"

	FlexiPageTypeRecordPage FlexiPageType = "RecordPage"

	FlexiPageTypeHomePage FlexiPageType = "HomePage"

	FlexiPageTypeMailAppAppPage FlexiPageType = "MailAppAppPage"

	FlexiPageTypeCommAppPage FlexiPageType = "CommAppPage"

	FlexiPageTypeCommObjectPage FlexiPageType = "CommObjectPage"

	FlexiPageTypeCommQuickActionCreatePage FlexiPageType = "CommQuickActionCreatePage"

	FlexiPageTypeCommRecordPage FlexiPageType = "CommRecordPage"

	FlexiPageTypeCommRelatedListPage FlexiPageType = "CommRelatedListPage"

	FlexiPageTypeCommSearchResultPage FlexiPageType = "CommSearchResultPage"

	FlexiPageTypeCommThemeLayoutPage FlexiPageType = "CommThemeLayoutPage"

	FlexiPageTypeUtilityBar FlexiPageType = "UtilityBar"
)

type FlowAssignmentOperator string

const (
	FlowAssignmentOperatorAssign FlowAssignmentOperator = "Assign"

	FlowAssignmentOperatorAdd FlowAssignmentOperator = "Add"

	FlowAssignmentOperatorSubtract FlowAssignmentOperator = "Subtract"

	FlowAssignmentOperatorAddItem FlowAssignmentOperator = "AddItem"
)

type FlowComparisonOperator string

const (
	FlowComparisonOperatorEqualTo FlowComparisonOperator = "EqualTo"

	FlowComparisonOperatorNotEqualTo FlowComparisonOperator = "NotEqualTo"

	FlowComparisonOperatorGreaterThan FlowComparisonOperator = "GreaterThan"

	FlowComparisonOperatorLessThan FlowComparisonOperator = "LessThan"

	FlowComparisonOperatorGreaterThanOrEqualTo FlowComparisonOperator = "GreaterThanOrEqualTo"

	FlowComparisonOperatorLessThanOrEqualTo FlowComparisonOperator = "LessThanOrEqualTo"

	FlowComparisonOperatorStartsWith FlowComparisonOperator = "StartsWith"

	FlowComparisonOperatorEndsWith FlowComparisonOperator = "EndsWith"

	FlowComparisonOperatorContains FlowComparisonOperator = "Contains"

	FlowComparisonOperatorIsNull FlowComparisonOperator = "IsNull"

	FlowComparisonOperatorWasSet FlowComparisonOperator = "WasSet"

	FlowComparisonOperatorWasSelected FlowComparisonOperator = "WasSelected"

	FlowComparisonOperatorWasVisited FlowComparisonOperator = "WasVisited"
)

type FlowRecordFilterOperator string

const (
	FlowRecordFilterOperatorEqualTo FlowRecordFilterOperator = "EqualTo"

	FlowRecordFilterOperatorNotEqualTo FlowRecordFilterOperator = "NotEqualTo"

	FlowRecordFilterOperatorGreaterThan FlowRecordFilterOperator = "GreaterThan"

	FlowRecordFilterOperatorLessThan FlowRecordFilterOperator = "LessThan"

	FlowRecordFilterOperatorGreaterThanOrEqualTo FlowRecordFilterOperator = "GreaterThanOrEqualTo"

	FlowRecordFilterOperatorLessThanOrEqualTo FlowRecordFilterOperator = "LessThanOrEqualTo"

	FlowRecordFilterOperatorStartsWith FlowRecordFilterOperator = "StartsWith"

	FlowRecordFilterOperatorEndsWith FlowRecordFilterOperator = "EndsWith"

	FlowRecordFilterOperatorContains FlowRecordFilterOperator = "Contains"

	FlowRecordFilterOperatorIsNull FlowRecordFilterOperator = "IsNull"
)

type FlowDataType string

const (
	FlowDataTypeCurrency FlowDataType = "Currency"

	FlowDataTypeDate FlowDataType = "Date"

	FlowDataTypeNumber FlowDataType = "Number"

	FlowDataTypeString FlowDataType = "String"

	FlowDataTypeBoolean FlowDataType = "Boolean"

	FlowDataTypeSObject FlowDataType = "SObject"

	FlowDataTypeDateTime FlowDataType = "DateTime"

	FlowDataTypePicklist FlowDataType = "Picklist"

	FlowDataTypeMultipicklist FlowDataType = "Multipicklist"
)

type FlowScreenFieldType string

const (
	FlowScreenFieldTypeDisplayText FlowScreenFieldType = "DisplayText"

	FlowScreenFieldTypeInputField FlowScreenFieldType = "InputField"

	FlowScreenFieldTypeLargeTextArea FlowScreenFieldType = "LargeTextArea"

	FlowScreenFieldTypePasswordField FlowScreenFieldType = "PasswordField"

	FlowScreenFieldTypeRadioButtons FlowScreenFieldType = "RadioButtons"

	FlowScreenFieldTypeDropdownBox FlowScreenFieldType = "DropdownBox"

	FlowScreenFieldTypeMultiSelectCheckboxes FlowScreenFieldType = "MultiSelectCheckboxes"

	FlowScreenFieldTypeMultiSelectPicklist FlowScreenFieldType = "MultiSelectPicklist"
)

type IterationOrder string

const (
	IterationOrderAsc IterationOrder = "Asc"

	IterationOrderDesc IterationOrder = "Desc"
)

type InvocableActionType string

const (
	InvocableActionTypeApex InvocableActionType = "apex"

	InvocableActionTypeChatterPost InvocableActionType = "chatterPost"

	InvocableActionTypeContentWorkspaceEnableFolders InvocableActionType = "contentWorkspaceEnableFolders"

	InvocableActionTypeEmailAlert InvocableActionType = "emailAlert"

	InvocableActionTypeEmailSimple InvocableActionType = "emailSimple"

	InvocableActionTypeFlow InvocableActionType = "flow"

	InvocableActionTypeMetricRefresh InvocableActionType = "metricRefresh"

	InvocableActionTypeQuickAction InvocableActionType = "quickAction"

	InvocableActionTypeSubmit InvocableActionType = "submit"

	InvocableActionTypeThanks InvocableActionType = "thanks"

	InvocableActionTypeThunderResponse InvocableActionType = "thunderResponse"
)

type FlowProcessType string

const (
	FlowProcessTypeAutoLaunchedFlow FlowProcessType = "AutoLaunchedFlow"

	FlowProcessTypeFlow FlowProcessType = "Flow"

	FlowProcessTypeWorkflow FlowProcessType = "Workflow"

	FlowProcessTypeCustomEvent FlowProcessType = "CustomEvent"

	FlowProcessTypeInvocableProcess FlowProcessType = "InvocableProcess"

	FlowProcessTypeLoginFlow FlowProcessType = "LoginFlow"

	FlowProcessTypeActionPlan FlowProcessType = "ActionPlan"

	FlowProcessTypeJourneyBuilderIntegration FlowProcessType = "JourneyBuilderIntegration"

	FlowProcessTypeUserProvisioningFlow FlowProcessType = "UserProvisioningFlow"
)

type FolderAccessTypes string

const (
	FolderAccessTypesShared FolderAccessTypes = "Shared"

	FolderAccessTypesPublic FolderAccessTypes = "Public"

	FolderAccessTypesHidden FolderAccessTypes = "Hidden"

	FolderAccessTypesPublicInternal FolderAccessTypes = "PublicInternal"
)

type FolderShareAccessLevel string

const (
	FolderShareAccessLevelView FolderShareAccessLevel = "View"

	FolderShareAccessLevelEditAllContents FolderShareAccessLevel = "EditAllContents"

	FolderShareAccessLevelManage FolderShareAccessLevel = "Manage"
)

type FolderSharedToType string

const (
	FolderSharedToTypeGroup FolderSharedToType = "Group"

	FolderSharedToTypeRole FolderSharedToType = "Role"

	FolderSharedToTypeRoleAndSubordinates FolderSharedToType = "RoleAndSubordinates"

	FolderSharedToTypeRoleAndSubordinatesInternal FolderSharedToType = "RoleAndSubordinatesInternal"

	FolderSharedToTypeManager FolderSharedToType = "Manager"

	FolderSharedToTypeManagerAndSubordinatesInternal FolderSharedToType = "ManagerAndSubordinatesInternal"

	FolderSharedToTypeOrganization FolderSharedToType = "Organization"

	FolderSharedToTypeTerritory FolderSharedToType = "Territory"

	FolderSharedToTypeTerritoryAndSubordinates FolderSharedToType = "TerritoryAndSubordinates"

	FolderSharedToTypeAllPrmUsers FolderSharedToType = "AllPrmUsers"

	FolderSharedToTypeUser FolderSharedToType = "User"

	FolderSharedToTypePartnerUser FolderSharedToType = "PartnerUser"

	FolderSharedToTypeAllCspUsers FolderSharedToType = "AllCspUsers"

	FolderSharedToTypeCustomerPortalUser FolderSharedToType = "CustomerPortalUser"

	FolderSharedToTypePortalRole FolderSharedToType = "PortalRole"

	FolderSharedToTypePortalRoleAndSubordinates FolderSharedToType = "PortalRoleAndSubordinates"
)

type PublicFolderAccess string

const (
	PublicFolderAccessReadOnly PublicFolderAccess = "ReadOnly"

	PublicFolderAccessReadWrite PublicFolderAccess = "ReadWrite"
)

type DisplayCurrency string

const (
	DisplayCurrencyCORPORATE DisplayCurrency = "CORPORATE"

	DisplayCurrencyPERSONAL DisplayCurrency = "PERSONAL"
)

type PeriodTypes string

const (
	PeriodTypesMonth PeriodTypes = "Month"

	PeriodTypesQuarter PeriodTypes = "Quarter"

	PeriodTypesWeek PeriodTypes = "Week"

	PeriodTypesYear PeriodTypes = "Year"
)

type PageComponentType string

const (
	PageComponentTypeLinks PageComponentType = "links"

	PageComponentTypeHtmlArea PageComponentType = "htmlArea"

	PageComponentTypeImageOrNote PageComponentType = "imageOrNote"

	PageComponentTypeVisualforcePage PageComponentType = "visualforcePage"
)

type PageComponentWidth string

const (
	PageComponentWidthNarrow PageComponentWidth = "narrow"

	PageComponentWidthWide PageComponentWidth = "wide"
)

type KnowledgeCaseEditor string

const (
	KnowledgeCaseEditorSimple KnowledgeCaseEditor = "simple"

	KnowledgeCaseEditorStandard KnowledgeCaseEditor = "standard"
)

type KnowledgeLanguageLookupValueType string

const (
	KnowledgeLanguageLookupValueTypeUser KnowledgeLanguageLookupValueType = "User"

	KnowledgeLanguageLookupValueTypeQueue KnowledgeLanguageLookupValueType = "Queue"
)

type FeedLayoutFilterPosition string

const (
	FeedLayoutFilterPositionCenterDropDown FeedLayoutFilterPosition = "CenterDropDown"

	FeedLayoutFilterPositionLeftFixed FeedLayoutFilterPosition = "LeftFixed"

	FeedLayoutFilterPositionLeftFloat FeedLayoutFilterPosition = "LeftFloat"
)

type FeedLayoutFilterType string

const (
	FeedLayoutFilterTypeAllUpdates FeedLayoutFilterType = "AllUpdates"

	FeedLayoutFilterTypeFeedItemType FeedLayoutFilterType = "FeedItemType"

	FeedLayoutFilterTypeCustom FeedLayoutFilterType = "Custom"
)

type FeedLayoutComponentType string

const (
	FeedLayoutComponentTypeHelpAndToolLinks FeedLayoutComponentType = "HelpAndToolLinks"

	FeedLayoutComponentTypeCustomButtons FeedLayoutComponentType = "CustomButtons"

	FeedLayoutComponentTypeFollowing FeedLayoutComponentType = "Following"

	FeedLayoutComponentTypeFollowers FeedLayoutComponentType = "Followers"

	FeedLayoutComponentTypeCustomLinks FeedLayoutComponentType = "CustomLinks"

	FeedLayoutComponentTypeMilestones FeedLayoutComponentType = "Milestones"

	FeedLayoutComponentTypeTopics FeedLayoutComponentType = "Topics"

	FeedLayoutComponentTypeCaseUnifiedFiles FeedLayoutComponentType = "CaseUnifiedFiles"

	FeedLayoutComponentTypeVisualforce FeedLayoutComponentType = "Visualforce"
)

type LayoutHeader string

const (
	LayoutHeaderPersonalTagging LayoutHeader = "PersonalTagging"

	LayoutHeaderPublicTagging LayoutHeader = "PublicTagging"
)

type UiBehavior string

const (
	UiBehaviorEdit UiBehavior = "Edit"

	UiBehaviorRequired UiBehavior = "Required"

	UiBehaviorReadonly UiBehavior = "Readonly"
)

type ReportChartComponentSize string

const (
	ReportChartComponentSizeSMALL ReportChartComponentSize = "SMALL"

	ReportChartComponentSizeMEDIUM ReportChartComponentSize = "MEDIUM"

	ReportChartComponentSizeLARGE ReportChartComponentSize = "LARGE"
)

type LayoutSectionStyle string

const (
	LayoutSectionStyleTwoColumnsTopToBottom LayoutSectionStyle = "TwoColumnsTopToBottom"

	LayoutSectionStyleTwoColumnsLeftToRight LayoutSectionStyle = "TwoColumnsLeftToRight"

	LayoutSectionStyleOneColumn LayoutSectionStyle = "OneColumn"

	LayoutSectionStyleCustomLinks LayoutSectionStyle = "CustomLinks"
)

type SummaryLayoutStyle string

const (
	SummaryLayoutStyleDefault SummaryLayoutStyle = "Default"

	SummaryLayoutStyleQuoteTemplate SummaryLayoutStyle = "QuoteTemplate"

	SummaryLayoutStyleDefaultQuoteTemplate SummaryLayoutStyle = "DefaultQuoteTemplate"

	SummaryLayoutStyleCaseInteraction SummaryLayoutStyle = "CaseInteraction"

	SummaryLayoutStyleQuickActionLayoutLeftRight SummaryLayoutStyle = "QuickActionLayoutLeftRight"

	SummaryLayoutStyleQuickActionLayoutTopDown SummaryLayoutStyle = "QuickActionLayoutTopDown"

	SummaryLayoutStylePathAssistant SummaryLayoutStyle = "PathAssistant"
)

type VisibleOrRequired string

const (
	VisibleOrRequiredVisibleOptional VisibleOrRequired = "VisibleOptional"

	VisibleOrRequiredVisibleRequired VisibleOrRequired = "VisibleRequired"

	VisibleOrRequiredNotVisible VisibleOrRequired = "NotVisible"
)

type LetterheadHorizontalAlignment string

const (
	LetterheadHorizontalAlignmentNone LetterheadHorizontalAlignment = "None"

	LetterheadHorizontalAlignmentLeft LetterheadHorizontalAlignment = "Left"

	LetterheadHorizontalAlignmentCenter LetterheadHorizontalAlignment = "Center"

	LetterheadHorizontalAlignmentRight LetterheadHorizontalAlignment = "Right"
)

type LetterheadVerticalAlignment string

const (
	LetterheadVerticalAlignmentNone LetterheadVerticalAlignment = "None"

	LetterheadVerticalAlignmentTop LetterheadVerticalAlignment = "Top"

	LetterheadVerticalAlignmentMiddle LetterheadVerticalAlignment = "Middle"

	LetterheadVerticalAlignmentBottom LetterheadVerticalAlignment = "Bottom"
)

type SupervisorAgentStatusFilter string

const (
	SupervisorAgentStatusFilterOnline SupervisorAgentStatusFilter = "Online"

	SupervisorAgentStatusFilterAway SupervisorAgentStatusFilter = "Away"

	SupervisorAgentStatusFilterOffline SupervisorAgentStatusFilter = "Offline"
)

type LiveChatButtonPresentation string

const (
	LiveChatButtonPresentationSlide LiveChatButtonPresentation = "Slide"

	LiveChatButtonPresentationFade LiveChatButtonPresentation = "Fade"

	LiveChatButtonPresentationAppear LiveChatButtonPresentation = "Appear"

	LiveChatButtonPresentationCustom LiveChatButtonPresentation = "Custom"
)

type LiveChatButtonInviteEndPosition string

const (
	LiveChatButtonInviteEndPositionTopLeft LiveChatButtonInviteEndPosition = "TopLeft"

	LiveChatButtonInviteEndPositionTop LiveChatButtonInviteEndPosition = "Top"

	LiveChatButtonInviteEndPositionTopRight LiveChatButtonInviteEndPosition = "TopRight"

	LiveChatButtonInviteEndPositionLeft LiveChatButtonInviteEndPosition = "Left"

	LiveChatButtonInviteEndPositionCenter LiveChatButtonInviteEndPosition = "Center"

	LiveChatButtonInviteEndPositionRight LiveChatButtonInviteEndPosition = "Right"

	LiveChatButtonInviteEndPositionBottomLeft LiveChatButtonInviteEndPosition = "BottomLeft"

	LiveChatButtonInviteEndPositionBottom LiveChatButtonInviteEndPosition = "Bottom"

	LiveChatButtonInviteEndPositionBottomRight LiveChatButtonInviteEndPosition = "BottomRight"
)

type LiveChatButtonInviteStartPosition string

const (
	LiveChatButtonInviteStartPositionTopLeft LiveChatButtonInviteStartPosition = "TopLeft"

	LiveChatButtonInviteStartPositionTopLeftTop LiveChatButtonInviteStartPosition = "TopLeftTop"

	LiveChatButtonInviteStartPositionTop LiveChatButtonInviteStartPosition = "Top"

	LiveChatButtonInviteStartPositionTopRightTop LiveChatButtonInviteStartPosition = "TopRightTop"

	LiveChatButtonInviteStartPositionTopRight LiveChatButtonInviteStartPosition = "TopRight"

	LiveChatButtonInviteStartPositionTopRightRight LiveChatButtonInviteStartPosition = "TopRightRight"

	LiveChatButtonInviteStartPositionRight LiveChatButtonInviteStartPosition = "Right"

	LiveChatButtonInviteStartPositionBottomRightRight LiveChatButtonInviteStartPosition = "BottomRightRight"

	LiveChatButtonInviteStartPositionBottomRight LiveChatButtonInviteStartPosition = "BottomRight"

	LiveChatButtonInviteStartPositionBottomRightBottom LiveChatButtonInviteStartPosition = "BottomRightBottom"

	LiveChatButtonInviteStartPositionBottom LiveChatButtonInviteStartPosition = "Bottom"

	LiveChatButtonInviteStartPositionBottomLeftBottom LiveChatButtonInviteStartPosition = "BottomLeftBottom"

	LiveChatButtonInviteStartPositionBottomLeft LiveChatButtonInviteStartPosition = "BottomLeft"

	LiveChatButtonInviteStartPositionBottomLeftLeft LiveChatButtonInviteStartPosition = "BottomLeftLeft"

	LiveChatButtonInviteStartPositionLeft LiveChatButtonInviteStartPosition = "Left"

	LiveChatButtonInviteStartPositionTopLeftLeft LiveChatButtonInviteStartPosition = "TopLeftLeft"
)

type LiveChatButtonRoutingType string

const (
	LiveChatButtonRoutingTypeChoice LiveChatButtonRoutingType = "Choice"

	LiveChatButtonRoutingTypeLeastActive LiveChatButtonRoutingType = "LeastActive"

	LiveChatButtonRoutingTypeMostAvailable LiveChatButtonRoutingType = "MostAvailable"
)

type LiveChatButtonType string

const (
	LiveChatButtonTypeStandard LiveChatButtonType = "Standard"

	LiveChatButtonTypeInvite LiveChatButtonType = "Invite"
)

type SensitiveDataActionType string

const (
	SensitiveDataActionTypeRemove SensitiveDataActionType = "Remove"

	SensitiveDataActionTypeReplace SensitiveDataActionType = "Replace"
)

type BlankValueBehavior string

const (
	BlankValueBehaviorMatchBlanks BlankValueBehavior = "MatchBlanks"

	BlankValueBehaviorNullNotAllowed BlankValueBehavior = "NullNotAllowed"
)

type MatchingMethod string

const (
	MatchingMethodExact MatchingMethod = "Exact"

	MatchingMethodFirstName MatchingMethod = "FirstName"

	MatchingMethodLastName MatchingMethod = "LastName"

	MatchingMethodCompanyName MatchingMethod = "CompanyName"

	MatchingMethodPhone MatchingMethod = "Phone"

	MatchingMethodCity MatchingMethod = "City"

	MatchingMethodStreet MatchingMethod = "Street"

	MatchingMethodZip MatchingMethod = "Zip"

	MatchingMethodTitle MatchingMethod = "Title"
)

type MatchingRuleStatus string

const (
	MatchingRuleStatusInactive MatchingRuleStatus = "Inactive"

	MatchingRuleStatusDeactivationFailed MatchingRuleStatus = "DeactivationFailed"

	MatchingRuleStatusActivating MatchingRuleStatus = "Activating"

	MatchingRuleStatusDeactivating MatchingRuleStatus = "Deactivating"

	MatchingRuleStatusActive MatchingRuleStatus = "Active"

	MatchingRuleStatusActivationFailed MatchingRuleStatus = "ActivationFailed"
)

type ApexCodeUnitStatus string

const (
	ApexCodeUnitStatusInactive ApexCodeUnitStatus = "Inactive"

	ApexCodeUnitStatusActive ApexCodeUnitStatus = "Active"

	ApexCodeUnitStatusDeleted ApexCodeUnitStatus = "Deleted"
)

type ContentAssetFormat string

const (
	ContentAssetFormatOriginal ContentAssetFormat = "Original"

	ContentAssetFormatZippedVersions ContentAssetFormat = "ZippedVersions"
)

type ContentAssetAccess string

const (
	ContentAssetAccessVIEWER ContentAssetAccess = "VIEWER"

	ContentAssetAccessCOLLABORATOR ContentAssetAccess = "COLLABORATOR"

	ContentAssetAccessINFERRED ContentAssetAccess = "INFERRED"
)

type DataPipelineType string

const (
	DataPipelineTypePig DataPipelineType = "Pig"
)

type EmailTemplateStyle string

const (
	EmailTemplateStyleNone EmailTemplateStyle = "none"

	EmailTemplateStyleFreeForm EmailTemplateStyle = "freeForm"

	EmailTemplateStyleFormalLetter EmailTemplateStyle = "formalLetter"

	EmailTemplateStylePromotionRight EmailTemplateStyle = "promotionRight"

	EmailTemplateStylePromotionLeft EmailTemplateStyle = "promotionLeft"

	EmailTemplateStyleNewsletter EmailTemplateStyle = "newsletter"

	EmailTemplateStyleProducts EmailTemplateStyle = "products"
)

type EmailTemplateType string

const (
	EmailTemplateTypeText EmailTemplateType = "text"

	EmailTemplateTypeHtml EmailTemplateType = "html"

	EmailTemplateTypeCustom EmailTemplateType = "custom"

	EmailTemplateTypeVisualforce EmailTemplateType = "visualforce"
)

type SControlContentSource string

const (
	SControlContentSourceHTML SControlContentSource = "HTML"

	SControlContentSourceURL SControlContentSource = "URL"

	SControlContentSourceSnippet SControlContentSource = "Snippet"
)

type StaticResourceCacheControl string

const (
	StaticResourceCacheControlPrivate StaticResourceCacheControl = "Private"

	StaticResourceCacheControlPublic StaticResourceCacheControl = "Public"
)

type MilestoneTypeRecurrenceType string

const (
	MilestoneTypeRecurrenceTypeNone MilestoneTypeRecurrenceType = "none"

	MilestoneTypeRecurrenceTypeRecursIndependently MilestoneTypeRecurrenceType = "recursIndependently"

	MilestoneTypeRecurrenceTypeRecursChained MilestoneTypeRecurrenceType = "recursChained"
)

type ModerationRuleAction string

const (
	ModerationRuleActionBlock ModerationRuleAction = "Block"

	ModerationRuleActionFreezeAndNotify ModerationRuleAction = "FreezeAndNotify"

	ModerationRuleActionReview ModerationRuleAction = "Review"

	ModerationRuleActionReplace ModerationRuleAction = "Replace"

	ModerationRuleActionFlag ModerationRuleAction = "Flag"
)

type NetworkStatus string

const (
	NetworkStatusUnderConstruction NetworkStatus = "UnderConstruction"

	NetworkStatusLive NetworkStatus = "Live"

	NetworkStatusDownForMaintenance NetworkStatus = "DownForMaintenance"
)

type APIAccessLevel string

const (
	APIAccessLevelUnrestricted APIAccessLevel = "Unrestricted"

	APIAccessLevelRestricted APIAccessLevel = "Restricted"
)

type PermissionSetTabVisibility string

const (
	PermissionSetTabVisibilityNone PermissionSetTabVisibility = "None"

	PermissionSetTabVisibilityAvailable PermissionSetTabVisibility = "Available"

	PermissionSetTabVisibilityVisible PermissionSetTabVisibility = "Visible"
)

type PlatformCacheType string

const (
	PlatformCacheTypeSession PlatformCacheType = "Session"

	PlatformCacheTypeOrganization PlatformCacheType = "Organization"
)

type PortalRoles string

const (
	PortalRolesExecutive PortalRoles = "Executive"

	PortalRolesManager PortalRoles = "Manager"

	PortalRolesWorker PortalRoles = "Worker"

	PortalRolesPersonAccount PortalRoles = "PersonAccount"
)

type PortalType string

const (
	PortalTypeCustomerSuccess PortalType = "CustomerSuccess"

	PortalTypePartner PortalType = "Partner"

	PortalTypeNetwork PortalType = "Network"
)

type TabVisibility string

const (
	TabVisibilityHidden TabVisibility = "Hidden"

	TabVisibilityDefaultOff TabVisibility = "DefaultOff"

	TabVisibilityDefaultOn TabVisibility = "DefaultOn"
)

type QuickActionLabel string

const (
	QuickActionLabelLogACall QuickActionLabel = "LogACall"

	QuickActionLabelLogANote QuickActionLabel = "LogANote"

	QuickActionLabelNew QuickActionLabel = "New"

	QuickActionLabelNewRecordType QuickActionLabel = "NewRecordType"

	QuickActionLabelUpdate QuickActionLabel = "Update"

	QuickActionLabelNewChild QuickActionLabel = "NewChild"

	QuickActionLabelNewChildRecordType QuickActionLabel = "NewChildRecordType"

	QuickActionLabelCreateNew QuickActionLabel = "CreateNew"

	QuickActionLabelCreateNewRecordType QuickActionLabel = "CreateNewRecordType"

	QuickActionLabelSendEmail QuickActionLabel = "SendEmail"

	QuickActionLabelQuickRecordType QuickActionLabel = "QuickRecordType"

	QuickActionLabelQuick QuickActionLabel = "Quick"

	QuickActionLabelEditDescription QuickActionLabel = "EditDescription"

	QuickActionLabelDefer QuickActionLabel = "Defer"

	QuickActionLabelChangeDueDate QuickActionLabel = "ChangeDueDate"

	QuickActionLabelChangePriority QuickActionLabel = "ChangePriority"

	QuickActionLabelChangeStatus QuickActionLabel = "ChangeStatus"

	QuickActionLabelSocialPost QuickActionLabel = "SocialPost"

	QuickActionLabelEscalate QuickActionLabel = "Escalate"

	QuickActionLabelEscalateToRecord QuickActionLabel = "EscalateToRecord"

	QuickActionLabelOfferFeedback QuickActionLabel = "OfferFeedback"

	QuickActionLabelRequestFeedback QuickActionLabel = "RequestFeedback"

	QuickActionLabelAddRecord QuickActionLabel = "AddRecord"

	QuickActionLabelAddMember QuickActionLabel = "AddMember"
)

type QuickActionType string

const (
	QuickActionTypeCreate QuickActionType = "Create"

	QuickActionTypeVisualforcePage QuickActionType = "VisualforcePage"

	QuickActionTypePost QuickActionType = "Post"

	QuickActionTypeSendEmail QuickActionType = "SendEmail"

	QuickActionTypeLogACall QuickActionType = "LogACall"

	QuickActionTypeSocialPost QuickActionType = "SocialPost"

	QuickActionTypeCanvas QuickActionType = "Canvas"

	QuickActionTypeUpdate QuickActionType = "Update"

	QuickActionTypeLightningComponent QuickActionType = "LightningComponent"
)

type ReportAggregateDatatype string

const (
	ReportAggregateDatatypeCurrency ReportAggregateDatatype = "currency"

	ReportAggregateDatatypePercent ReportAggregateDatatype = "percent"

	ReportAggregateDatatypeNumber ReportAggregateDatatype = "number"
)

type ReportBucketFieldType string

const (
	ReportBucketFieldTypeText ReportBucketFieldType = "text"

	ReportBucketFieldTypeNumber ReportBucketFieldType = "number"

	ReportBucketFieldTypePicklist ReportBucketFieldType = "picklist"
)

type ReportFormulaNullTreatment string

const (
	ReportFormulaNullTreatmentN ReportFormulaNullTreatment = "n"

	ReportFormulaNullTreatmentZ ReportFormulaNullTreatment = "z"
)

type ChartType string

const (
	ChartTypeNone ChartType = "None"

	ChartTypeScatter ChartType = "Scatter"

	ChartTypeScatterGrouped ChartType = "ScatterGrouped"

	ChartTypeBubble ChartType = "Bubble"

	ChartTypeBubbleGrouped ChartType = "BubbleGrouped"

	ChartTypeHorizontalBar ChartType = "HorizontalBar"

	ChartTypeHorizontalBarGrouped ChartType = "HorizontalBarGrouped"

	ChartTypeHorizontalBarStacked ChartType = "HorizontalBarStacked"

	ChartTypeHorizontalBarStackedTo100 ChartType = "HorizontalBarStackedTo100"

	ChartTypeVerticalColumn ChartType = "VerticalColumn"

	ChartTypeVerticalColumnGrouped ChartType = "VerticalColumnGrouped"

	ChartTypeVerticalColumnStacked ChartType = "VerticalColumnStacked"

	ChartTypeVerticalColumnStackedTo100 ChartType = "VerticalColumnStackedTo100"

	ChartTypeLine ChartType = "Line"

	ChartTypeLineGrouped ChartType = "LineGrouped"

	ChartTypeLineCumulative ChartType = "LineCumulative"

	ChartTypeLineCumulativeGrouped ChartType = "LineCumulativeGrouped"

	ChartTypePie ChartType = "Pie"

	ChartTypeDonut ChartType = "Donut"

	ChartTypeFunnel ChartType = "Funnel"

	ChartTypeVerticalColumnLine ChartType = "VerticalColumnLine"

	ChartTypeVerticalColumnGroupedLine ChartType = "VerticalColumnGroupedLine"

	ChartTypeVerticalColumnStackedLine ChartType = "VerticalColumnStackedLine"

	ChartTypePlugin ChartType = "Plugin"
)

type ChartPosition string

const (
	ChartPositionCHARTTOP ChartPosition = "CHARTTOP"

	ChartPositionCHARTBOTTOM ChartPosition = "CHARTBOTTOM"
)

type ReportChartSize string

const (
	ReportChartSizeTiny ReportChartSize = "Tiny"

	ReportChartSizeSmall ReportChartSize = "Small"

	ReportChartSizeMedium ReportChartSize = "Medium"

	ReportChartSizeLarge ReportChartSize = "Large"

	ReportChartSizeHuge ReportChartSize = "Huge"
)

type ObjectFilterOperator string

const (
	ObjectFilterOperatorWith ObjectFilterOperator = "with"

	ObjectFilterOperatorWithout ObjectFilterOperator = "without"
)

type CurrencyIsoCode string

const (
	CurrencyIsoCodeADP CurrencyIsoCode = "ADP"

	CurrencyIsoCodeAED CurrencyIsoCode = "AED"

	CurrencyIsoCodeAFA CurrencyIsoCode = "AFA"

	CurrencyIsoCodeAFN CurrencyIsoCode = "AFN"

	CurrencyIsoCodeALL CurrencyIsoCode = "ALL"

	CurrencyIsoCodeAMD CurrencyIsoCode = "AMD"

	CurrencyIsoCodeANG CurrencyIsoCode = "ANG"

	CurrencyIsoCodeAOA CurrencyIsoCode = "AOA"

	CurrencyIsoCodeARS CurrencyIsoCode = "ARS"

	CurrencyIsoCodeATS CurrencyIsoCode = "ATS"

	CurrencyIsoCodeAUD CurrencyIsoCode = "AUD"

	CurrencyIsoCodeAWG CurrencyIsoCode = "AWG"

	CurrencyIsoCodeAZM CurrencyIsoCode = "AZM"

	CurrencyIsoCodeAZN CurrencyIsoCode = "AZN"

	CurrencyIsoCodeBAM CurrencyIsoCode = "BAM"

	CurrencyIsoCodeBBD CurrencyIsoCode = "BBD"

	CurrencyIsoCodeBDT CurrencyIsoCode = "BDT"

	CurrencyIsoCodeBEF CurrencyIsoCode = "BEF"

	CurrencyIsoCodeBGL CurrencyIsoCode = "BGL"

	CurrencyIsoCodeBGN CurrencyIsoCode = "BGN"

	CurrencyIsoCodeBHD CurrencyIsoCode = "BHD"

	CurrencyIsoCodeBIF CurrencyIsoCode = "BIF"

	CurrencyIsoCodeBMD CurrencyIsoCode = "BMD"

	CurrencyIsoCodeBND CurrencyIsoCode = "BND"

	CurrencyIsoCodeBOB CurrencyIsoCode = "BOB"

	CurrencyIsoCodeBOV CurrencyIsoCode = "BOV"

	CurrencyIsoCodeBRB CurrencyIsoCode = "BRB"

	CurrencyIsoCodeBRL CurrencyIsoCode = "BRL"

	CurrencyIsoCodeBSD CurrencyIsoCode = "BSD"

	CurrencyIsoCodeBTN CurrencyIsoCode = "BTN"

	CurrencyIsoCodeBWP CurrencyIsoCode = "BWP"

	CurrencyIsoCodeBYB CurrencyIsoCode = "BYB"

	CurrencyIsoCodeBYR CurrencyIsoCode = "BYR"

	CurrencyIsoCodeBZD CurrencyIsoCode = "BZD"

	CurrencyIsoCodeCAD CurrencyIsoCode = "CAD"

	CurrencyIsoCodeCDF CurrencyIsoCode = "CDF"

	CurrencyIsoCodeCHF CurrencyIsoCode = "CHF"

	CurrencyIsoCodeCLF CurrencyIsoCode = "CLF"

	CurrencyIsoCodeCLP CurrencyIsoCode = "CLP"

	CurrencyIsoCodeCNY CurrencyIsoCode = "CNY"

	CurrencyIsoCodeCOP CurrencyIsoCode = "COP"

	CurrencyIsoCodeCRC CurrencyIsoCode = "CRC"

	CurrencyIsoCodeCSD CurrencyIsoCode = "CSD"

	CurrencyIsoCodeCUC CurrencyIsoCode = "CUC"

	CurrencyIsoCodeCUP CurrencyIsoCode = "CUP"

	CurrencyIsoCodeCVE CurrencyIsoCode = "CVE"

	CurrencyIsoCodeCYP CurrencyIsoCode = "CYP"

	CurrencyIsoCodeCZK CurrencyIsoCode = "CZK"

	CurrencyIsoCodeDEM CurrencyIsoCode = "DEM"

	CurrencyIsoCodeDJF CurrencyIsoCode = "DJF"

	CurrencyIsoCodeDKK CurrencyIsoCode = "DKK"

	CurrencyIsoCodeDOP CurrencyIsoCode = "DOP"

	CurrencyIsoCodeDZD CurrencyIsoCode = "DZD"

	CurrencyIsoCodeECS CurrencyIsoCode = "ECS"

	CurrencyIsoCodeEEK CurrencyIsoCode = "EEK"

	CurrencyIsoCodeEGP CurrencyIsoCode = "EGP"

	CurrencyIsoCodeERN CurrencyIsoCode = "ERN"

	CurrencyIsoCodeESP CurrencyIsoCode = "ESP"

	CurrencyIsoCodeETB CurrencyIsoCode = "ETB"

	CurrencyIsoCodeEUR CurrencyIsoCode = "EUR"

	CurrencyIsoCodeFIM CurrencyIsoCode = "FIM"

	CurrencyIsoCodeFJD CurrencyIsoCode = "FJD"

	CurrencyIsoCodeFKP CurrencyIsoCode = "FKP"

	CurrencyIsoCodeFRF CurrencyIsoCode = "FRF"

	CurrencyIsoCodeGBP CurrencyIsoCode = "GBP"

	CurrencyIsoCodeGEL CurrencyIsoCode = "GEL"

	CurrencyIsoCodeGHC CurrencyIsoCode = "GHC"

	CurrencyIsoCodeGHS CurrencyIsoCode = "GHS"

	CurrencyIsoCodeGIP CurrencyIsoCode = "GIP"

	CurrencyIsoCodeGMD CurrencyIsoCode = "GMD"

	CurrencyIsoCodeGNF CurrencyIsoCode = "GNF"

	CurrencyIsoCodeGRD CurrencyIsoCode = "GRD"

	CurrencyIsoCodeGTQ CurrencyIsoCode = "GTQ"

	CurrencyIsoCodeGWP CurrencyIsoCode = "GWP"

	CurrencyIsoCodeGYD CurrencyIsoCode = "GYD"

	CurrencyIsoCodeHKD CurrencyIsoCode = "HKD"

	CurrencyIsoCodeHNL CurrencyIsoCode = "HNL"

	CurrencyIsoCodeHRD CurrencyIsoCode = "HRD"

	CurrencyIsoCodeHRK CurrencyIsoCode = "HRK"

	CurrencyIsoCodeHTG CurrencyIsoCode = "HTG"

	CurrencyIsoCodeHUF CurrencyIsoCode = "HUF"

	CurrencyIsoCodeIDR CurrencyIsoCode = "IDR"

	CurrencyIsoCodeIEP CurrencyIsoCode = "IEP"

	CurrencyIsoCodeILS CurrencyIsoCode = "ILS"

	CurrencyIsoCodeINR CurrencyIsoCode = "INR"

	CurrencyIsoCodeIQD CurrencyIsoCode = "IQD"

	CurrencyIsoCodeIRR CurrencyIsoCode = "IRR"

	CurrencyIsoCodeISK CurrencyIsoCode = "ISK"

	CurrencyIsoCodeITL CurrencyIsoCode = "ITL"

	CurrencyIsoCodeJMD CurrencyIsoCode = "JMD"

	CurrencyIsoCodeJOD CurrencyIsoCode = "JOD"

	CurrencyIsoCodeJPY CurrencyIsoCode = "JPY"

	CurrencyIsoCodeKES CurrencyIsoCode = "KES"

	CurrencyIsoCodeKGS CurrencyIsoCode = "KGS"

	CurrencyIsoCodeKHR CurrencyIsoCode = "KHR"

	CurrencyIsoCodeKMF CurrencyIsoCode = "KMF"

	CurrencyIsoCodeKPW CurrencyIsoCode = "KPW"

	CurrencyIsoCodeKRW CurrencyIsoCode = "KRW"

	CurrencyIsoCodeKWD CurrencyIsoCode = "KWD"

	CurrencyIsoCodeKYD CurrencyIsoCode = "KYD"

	CurrencyIsoCodeKZT CurrencyIsoCode = "KZT"

	CurrencyIsoCodeLAK CurrencyIsoCode = "LAK"

	CurrencyIsoCodeLBP CurrencyIsoCode = "LBP"

	CurrencyIsoCodeLKR CurrencyIsoCode = "LKR"

	CurrencyIsoCodeLRD CurrencyIsoCode = "LRD"

	CurrencyIsoCodeLSL CurrencyIsoCode = "LSL"

	CurrencyIsoCodeLTL CurrencyIsoCode = "LTL"

	CurrencyIsoCodeLUF CurrencyIsoCode = "LUF"

	CurrencyIsoCodeLVL CurrencyIsoCode = "LVL"

	CurrencyIsoCodeLYD CurrencyIsoCode = "LYD"

	CurrencyIsoCodeMAD CurrencyIsoCode = "MAD"

	CurrencyIsoCodeMDL CurrencyIsoCode = "MDL"

	CurrencyIsoCodeMGA CurrencyIsoCode = "MGA"

	CurrencyIsoCodeMGF CurrencyIsoCode = "MGF"

	CurrencyIsoCodeMKD CurrencyIsoCode = "MKD"

	CurrencyIsoCodeMMK CurrencyIsoCode = "MMK"

	CurrencyIsoCodeMNT CurrencyIsoCode = "MNT"

	CurrencyIsoCodeMOP CurrencyIsoCode = "MOP"

	CurrencyIsoCodeMRO CurrencyIsoCode = "MRO"

	CurrencyIsoCodeMTL CurrencyIsoCode = "MTL"

	CurrencyIsoCodeMUR CurrencyIsoCode = "MUR"

	CurrencyIsoCodeMVR CurrencyIsoCode = "MVR"

	CurrencyIsoCodeMWK CurrencyIsoCode = "MWK"

	CurrencyIsoCodeMXN CurrencyIsoCode = "MXN"

	CurrencyIsoCodeMXV CurrencyIsoCode = "MXV"

	CurrencyIsoCodeMYR CurrencyIsoCode = "MYR"

	CurrencyIsoCodeMZM CurrencyIsoCode = "MZM"

	CurrencyIsoCodeMZN CurrencyIsoCode = "MZN"

	CurrencyIsoCodeNAD CurrencyIsoCode = "NAD"

	CurrencyIsoCodeNGN CurrencyIsoCode = "NGN"

	CurrencyIsoCodeNIO CurrencyIsoCode = "NIO"

	CurrencyIsoCodeNLG CurrencyIsoCode = "NLG"

	CurrencyIsoCodeNOK CurrencyIsoCode = "NOK"

	CurrencyIsoCodeNPR CurrencyIsoCode = "NPR"

	CurrencyIsoCodeNZD CurrencyIsoCode = "NZD"

	CurrencyIsoCodeOMR CurrencyIsoCode = "OMR"

	CurrencyIsoCodePAB CurrencyIsoCode = "PAB"

	CurrencyIsoCodePEN CurrencyIsoCode = "PEN"

	CurrencyIsoCodePGK CurrencyIsoCode = "PGK"

	CurrencyIsoCodePHP CurrencyIsoCode = "PHP"

	CurrencyIsoCodePKR CurrencyIsoCode = "PKR"

	CurrencyIsoCodePLN CurrencyIsoCode = "PLN"

	CurrencyIsoCodePTE CurrencyIsoCode = "PTE"

	CurrencyIsoCodePYG CurrencyIsoCode = "PYG"

	CurrencyIsoCodeQAR CurrencyIsoCode = "QAR"

	CurrencyIsoCodeRMB CurrencyIsoCode = "RMB"

	CurrencyIsoCodeROL CurrencyIsoCode = "ROL"

	CurrencyIsoCodeRON CurrencyIsoCode = "RON"

	CurrencyIsoCodeRSD CurrencyIsoCode = "RSD"

	CurrencyIsoCodeRUB CurrencyIsoCode = "RUB"

	CurrencyIsoCodeRUR CurrencyIsoCode = "RUR"

	CurrencyIsoCodeRWF CurrencyIsoCode = "RWF"

	CurrencyIsoCodeSAR CurrencyIsoCode = "SAR"

	CurrencyIsoCodeSBD CurrencyIsoCode = "SBD"

	CurrencyIsoCodeSCR CurrencyIsoCode = "SCR"

	CurrencyIsoCodeSDD CurrencyIsoCode = "SDD"

	CurrencyIsoCodeSDG CurrencyIsoCode = "SDG"

	CurrencyIsoCodeSEK CurrencyIsoCode = "SEK"

	CurrencyIsoCodeSGD CurrencyIsoCode = "SGD"

	CurrencyIsoCodeSHP CurrencyIsoCode = "SHP"

	CurrencyIsoCodeSIT CurrencyIsoCode = "SIT"

	CurrencyIsoCodeSKK CurrencyIsoCode = "SKK"

	CurrencyIsoCodeSLL CurrencyIsoCode = "SLL"

	CurrencyIsoCodeSOS CurrencyIsoCode = "SOS"

	CurrencyIsoCodeSRD CurrencyIsoCode = "SRD"

	CurrencyIsoCodeSRG CurrencyIsoCode = "SRG"

	CurrencyIsoCodeSSP CurrencyIsoCode = "SSP"

	CurrencyIsoCodeSTD CurrencyIsoCode = "STD"

	CurrencyIsoCodeSUR CurrencyIsoCode = "SUR"

	CurrencyIsoCodeSVC CurrencyIsoCode = "SVC"

	CurrencyIsoCodeSYP CurrencyIsoCode = "SYP"

	CurrencyIsoCodeSZL CurrencyIsoCode = "SZL"

	CurrencyIsoCodeTHB CurrencyIsoCode = "THB"

	CurrencyIsoCodeTJR CurrencyIsoCode = "TJR"

	CurrencyIsoCodeTJS CurrencyIsoCode = "TJS"

	CurrencyIsoCodeTMM CurrencyIsoCode = "TMM"

	CurrencyIsoCodeTMT CurrencyIsoCode = "TMT"

	CurrencyIsoCodeTND CurrencyIsoCode = "TND"

	CurrencyIsoCodeTOP CurrencyIsoCode = "TOP"

	CurrencyIsoCodeTPE CurrencyIsoCode = "TPE"

	CurrencyIsoCodeTRL CurrencyIsoCode = "TRL"

	CurrencyIsoCodeTRY CurrencyIsoCode = "TRY"

	CurrencyIsoCodeTTD CurrencyIsoCode = "TTD"

	CurrencyIsoCodeTWD CurrencyIsoCode = "TWD"

	CurrencyIsoCodeTZS CurrencyIsoCode = "TZS"

	CurrencyIsoCodeUAH CurrencyIsoCode = "UAH"

	CurrencyIsoCodeUGX CurrencyIsoCode = "UGX"

	CurrencyIsoCodeUSD CurrencyIsoCode = "USD"

	CurrencyIsoCodeUYU CurrencyIsoCode = "UYU"

	CurrencyIsoCodeUZS CurrencyIsoCode = "UZS"

	CurrencyIsoCodeVEB CurrencyIsoCode = "VEB"

	CurrencyIsoCodeVEF CurrencyIsoCode = "VEF"

	CurrencyIsoCodeVND CurrencyIsoCode = "VND"

	CurrencyIsoCodeVUV CurrencyIsoCode = "VUV"

	CurrencyIsoCodeWST CurrencyIsoCode = "WST"

	CurrencyIsoCodeXAF CurrencyIsoCode = "XAF"

	CurrencyIsoCodeXCD CurrencyIsoCode = "XCD"

	CurrencyIsoCodeXOF CurrencyIsoCode = "XOF"

	CurrencyIsoCodeXPF CurrencyIsoCode = "XPF"

	CurrencyIsoCodeYER CurrencyIsoCode = "YER"

	CurrencyIsoCodeYUM CurrencyIsoCode = "YUM"

	CurrencyIsoCodeZAR CurrencyIsoCode = "ZAR"

	CurrencyIsoCodeZMK CurrencyIsoCode = "ZMK"

	CurrencyIsoCodeZMW CurrencyIsoCode = "ZMW"

	CurrencyIsoCodeZWD CurrencyIsoCode = "ZWD"

	CurrencyIsoCodeZWL CurrencyIsoCode = "ZWL"
)

type DataCategoryFilterOperation string

const (
	DataCategoryFilterOperationAbove DataCategoryFilterOperation = "above"

	DataCategoryFilterOperationBelow DataCategoryFilterOperation = "below"

	DataCategoryFilterOperationAt DataCategoryFilterOperation = "at"

	DataCategoryFilterOperationAboveOrBelow DataCategoryFilterOperation = "aboveOrBelow"
)

type ReportFormat string

const (
	ReportFormatMultiBlock ReportFormat = "MultiBlock"

	ReportFormatMatrix ReportFormat = "Matrix"

	ReportFormatSummary ReportFormat = "Summary"

	ReportFormatTabular ReportFormat = "Tabular"
)

type ReportAggrType string

const (
	ReportAggrTypeSum ReportAggrType = "Sum"

	ReportAggrTypeAverage ReportAggrType = "Average"

	ReportAggrTypeMaximum ReportAggrType = "Maximum"

	ReportAggrTypeMinimum ReportAggrType = "Minimum"

	ReportAggrTypeRowCount ReportAggrType = "RowCount"
)

type UserDateGranularity string

const (
	UserDateGranularityNone UserDateGranularity = "None"

	UserDateGranularityDay UserDateGranularity = "Day"

	UserDateGranularityWeek UserDateGranularity = "Week"

	UserDateGranularityMonth UserDateGranularity = "Month"

	UserDateGranularityQuarter UserDateGranularity = "Quarter"

	UserDateGranularityYear UserDateGranularity = "Year"

	UserDateGranularityFiscalQuarter UserDateGranularity = "FiscalQuarter"

	UserDateGranularityFiscalYear UserDateGranularity = "FiscalYear"

	UserDateGranularityMonthInYear UserDateGranularity = "MonthInYear"

	UserDateGranularityDayInMonth UserDateGranularity = "DayInMonth"

	UserDateGranularityFiscalPeriod UserDateGranularity = "FiscalPeriod"

	UserDateGranularityFiscalWeek UserDateGranularity = "FiscalWeek"
)

type ReportSortType string

const (
	ReportSortTypeColumn ReportSortType = "Column"

	ReportSortTypeAggregate ReportSortType = "Aggregate"

	ReportSortTypeCustomSummaryFormula ReportSortType = "CustomSummaryFormula"
)

type UserDateInterval string

const (
	UserDateIntervalINTERVALCURRENT UserDateInterval = "INTERVALCURRENT"

	UserDateIntervalINTERVALCURNEXT1 UserDateInterval = "INTERVALCURNEXT1"

	UserDateIntervalINTERVALCURPREV1 UserDateInterval = "INTERVALCURPREV1"

	UserDateIntervalINTERVALNEXT1 UserDateInterval = "INTERVALNEXT1"

	UserDateIntervalINTERVALPREV1 UserDateInterval = "INTERVALPREV1"

	UserDateIntervalINTERVALCURNEXT3 UserDateInterval = "INTERVALCURNEXT3"

	UserDateIntervalINTERVALCURFY UserDateInterval = "INTERVALCURFY"

	UserDateIntervalINTERVALPREVFY UserDateInterval = "INTERVALPREVFY"

	UserDateIntervalINTERVALPREV2FY UserDateInterval = "INTERVALPREV2FY"

	UserDateIntervalINTERVALAGO2FY UserDateInterval = "INTERVALAGO2FY"

	UserDateIntervalINTERVALNEXTFY UserDateInterval = "INTERVALNEXTFY"

	UserDateIntervalINTERVALPREVCURFY UserDateInterval = "INTERVALPREVCURFY"

	UserDateIntervalINTERVALPREVCUR2FY UserDateInterval = "INTERVALPREVCUR2FY"

	UserDateIntervalINTERVALCURNEXTFY UserDateInterval = "INTERVALCURNEXTFY"

	UserDateIntervalINTERVALCUSTOM UserDateInterval = "INTERVALCUSTOM"

	UserDateIntervalINTERVALYESTERDAY UserDateInterval = "INTERVALYESTERDAY"

	UserDateIntervalINTERVALTODAY UserDateInterval = "INTERVALTODAY"

	UserDateIntervalINTERVALTOMORROW UserDateInterval = "INTERVALTOMORROW"

	UserDateIntervalINTERVALLASTWEEK UserDateInterval = "INTERVALLASTWEEK"

	UserDateIntervalINTERVALTHISWEEK UserDateInterval = "INTERVALTHISWEEK"

	UserDateIntervalINTERVALNEXTWEEK UserDateInterval = "INTERVALNEXTWEEK"

	UserDateIntervalINTERVALLASTMONTH UserDateInterval = "INTERVALLASTMONTH"

	UserDateIntervalINTERVALTHISMONTH UserDateInterval = "INTERVALTHISMONTH"

	UserDateIntervalINTERVALNEXTMONTH UserDateInterval = "INTERVALNEXTMONTH"

	UserDateIntervalINTERVALLASTTHISMONTH UserDateInterval = "INTERVALLASTTHISMONTH"

	UserDateIntervalINTERVALTHISNEXTMONTH UserDateInterval = "INTERVALTHISNEXTMONTH"

	UserDateIntervalINTERVALCURRENTQ UserDateInterval = "INTERVALCURRENTQ"

	UserDateIntervalINTERVALCURNEXTQ UserDateInterval = "INTERVALCURNEXTQ"

	UserDateIntervalINTERVALCURPREVQ UserDateInterval = "INTERVALCURPREVQ"

	UserDateIntervalINTERVALNEXTQ UserDateInterval = "INTERVALNEXTQ"

	UserDateIntervalINTERVALPREVQ UserDateInterval = "INTERVALPREVQ"

	UserDateIntervalINTERVALCURNEXT3Q UserDateInterval = "INTERVALCURNEXT3Q"

	UserDateIntervalINTERVALCURY UserDateInterval = "INTERVALCURY"

	UserDateIntervalINTERVALPREVY UserDateInterval = "INTERVALPREVY"

	UserDateIntervalINTERVALPREV2Y UserDateInterval = "INTERVALPREV2Y"

	UserDateIntervalINTERVALAGO2Y UserDateInterval = "INTERVALAGO2Y"

	UserDateIntervalINTERVALNEXTY UserDateInterval = "INTERVALNEXTY"

	UserDateIntervalINTERVALPREVCURY UserDateInterval = "INTERVALPREVCURY"

	UserDateIntervalINTERVALPREVCUR2Y UserDateInterval = "INTERVALPREVCUR2Y"

	UserDateIntervalINTERVALCURNEXTY UserDateInterval = "INTERVALCURNEXTY"

	UserDateIntervalINTERVALLAST7 UserDateInterval = "INTERVALLAST7"

	UserDateIntervalINTERVALLAST30 UserDateInterval = "INTERVALLAST30"

	UserDateIntervalINTERVALLAST60 UserDateInterval = "INTERVALLAST60"

	UserDateIntervalINTERVALLAST90 UserDateInterval = "INTERVALLAST90"

	UserDateIntervalINTERVALLAST120 UserDateInterval = "INTERVALLAST120"

	UserDateIntervalINTERVALNEXT7 UserDateInterval = "INTERVALNEXT7"

	UserDateIntervalINTERVALNEXT30 UserDateInterval = "INTERVALNEXT30"

	UserDateIntervalINTERVALNEXT60 UserDateInterval = "INTERVALNEXT60"

	UserDateIntervalINTERVALNEXT90 UserDateInterval = "INTERVALNEXT90"

	UserDateIntervalINTERVALNEXT120 UserDateInterval = "INTERVALNEXT120"

	UserDateIntervalLASTFISCALWEEK UserDateInterval = "LASTFISCALWEEK"

	UserDateIntervalTHISFISCALWEEK UserDateInterval = "THISFISCALWEEK"

	UserDateIntervalNEXTFISCALWEEK UserDateInterval = "NEXTFISCALWEEK"

	UserDateIntervalLASTFISCALPERIOD UserDateInterval = "LASTFISCALPERIOD"

	UserDateIntervalTHISFISCALPERIOD UserDateInterval = "THISFISCALPERIOD"

	UserDateIntervalNEXTFISCALPERIOD UserDateInterval = "NEXTFISCALPERIOD"

	UserDateIntervalLASTTHISFISCALPERIOD UserDateInterval = "LASTTHISFISCALPERIOD"

	UserDateIntervalTHISNEXTFISCALPERIOD UserDateInterval = "THISNEXTFISCALPERIOD"

	UserDateIntervalCURRENTENTITLEMENTPERIOD UserDateInterval = "CURRENTENTITLEMENTPERIOD"

	UserDateIntervalPREVIOUSENTITLEMENTPERIOD UserDateInterval = "PREVIOUSENTITLEMENTPERIOD"

	UserDateIntervalPREVIOUSTWOENTITLEMENTPERIODS UserDateInterval = "PREVIOUSTWOENTITLEMENTPERIODS"

	UserDateIntervalTWOENTITLEMENTPERIODSAGO UserDateInterval = "TWOENTITLEMENTPERIODSAGO"

	UserDateIntervalCURRENTANDPREVIOUSENTITLEMENTPERIOD UserDateInterval = "CURRENTANDPREVIOUSENTITLEMENTPERIOD"

	UserDateIntervalCURRENTANDPREVIOUSTWOENTITLEMENTPERIODS UserDateInterval = "CURRENTANDPREVIOUSTWOENTITLEMENTPERIODS"
)

type ReportTypeCategory string

const (
	ReportTypeCategoryAccounts ReportTypeCategory = "accounts"

	ReportTypeCategoryOpportunities ReportTypeCategory = "opportunities"

	ReportTypeCategoryForecasts ReportTypeCategory = "forecasts"

	ReportTypeCategoryCases ReportTypeCategory = "cases"

	ReportTypeCategoryLeads ReportTypeCategory = "leads"

	ReportTypeCategoryCampaigns ReportTypeCategory = "campaigns"

	ReportTypeCategoryActivities ReportTypeCategory = "activities"

	ReportTypeCategoryBusop ReportTypeCategory = "busop"

	ReportTypeCategoryProducts ReportTypeCategory = "products"

	ReportTypeCategoryAdmin ReportTypeCategory = "admin"

	ReportTypeCategoryTerritory ReportTypeCategory = "territory"

	ReportTypeCategoryOther ReportTypeCategory = "other"

	ReportTypeCategoryContent ReportTypeCategory = "content"

	ReportTypeCategoryUsageentitlement ReportTypeCategory = "usageentitlement"

	ReportTypeCategoryWdc ReportTypeCategory = "wdc"

	ReportTypeCategoryCalibration ReportTypeCategory = "calibration"

	ReportTypeCategoryTerritory2 ReportTypeCategory = "territory2"
)

type SamlIdentityLocationType string

const (
	SamlIdentityLocationTypeSubjectNameId SamlIdentityLocationType = "SubjectNameId"

	SamlIdentityLocationTypeAttribute SamlIdentityLocationType = "Attribute"
)

type SamlIdentityType string

const (
	SamlIdentityTypeUsername SamlIdentityType = "Username"

	SamlIdentityTypeFederationId SamlIdentityType = "FederationId"

	SamlIdentityTypeUserId SamlIdentityType = "UserId"
)

type SamlType string

const (
	SamlTypeSAML11 SamlType = "SAML11"

	SamlTypeSAML20 SamlType = "SAML20"
)

type Complexity string

const (
	ComplexityNoRestriction Complexity = "NoRestriction"

	ComplexityAlphaNumeric Complexity = "AlphaNumeric"

	ComplexitySpecialCharacters Complexity = "SpecialCharacters"

	ComplexityUpperLowerCaseNumeric Complexity = "UpperLowerCaseNumeric"

	ComplexityUpperLowerCaseNumericSpecialCharacters Complexity = "UpperLowerCaseNumericSpecialCharacters"
)

type Expiration string

const (
	ExpirationThirtyDays Expiration = "ThirtyDays"

	ExpirationSixtyDays Expiration = "SixtyDays"

	ExpirationNinetyDays Expiration = "NinetyDays"

	ExpirationSixMonths Expiration = "SixMonths"

	ExpirationOneYear Expiration = "OneYear"

	ExpirationNever Expiration = "Never"
)

type LockoutInterval string

const (
	LockoutIntervalFifteenMinutes LockoutInterval = "FifteenMinutes"

	LockoutIntervalThirtyMinutes LockoutInterval = "ThirtyMinutes"

	LockoutIntervalSixtyMinutes LockoutInterval = "SixtyMinutes"

	LockoutIntervalForever LockoutInterval = "Forever"
)

type MaxLoginAttempts string

const (
	MaxLoginAttemptsThreeAttempts MaxLoginAttempts = "ThreeAttempts"

	MaxLoginAttemptsFiveAttempts MaxLoginAttempts = "FiveAttempts"

	MaxLoginAttemptsTenAttempts MaxLoginAttempts = "TenAttempts"

	MaxLoginAttemptsNoLimit MaxLoginAttempts = "NoLimit"
)

type QuestionRestriction string

const (
	QuestionRestrictionNone QuestionRestriction = "None"

	QuestionRestrictionDoesNotContainPassword QuestionRestriction = "DoesNotContainPassword"
)

type SessionTimeout string

const (
	SessionTimeoutTwentyFourHours SessionTimeout = "TwentyFourHours"

	SessionTimeoutTwelveHours SessionTimeout = "TwelveHours"

	SessionTimeoutEightHours SessionTimeout = "EightHours"

	SessionTimeoutFourHours SessionTimeout = "FourHours"

	SessionTimeoutTwoHours SessionTimeout = "TwoHours"

	SessionTimeoutSixtyMinutes SessionTimeout = "SixtyMinutes"

	SessionTimeoutThirtyMinutes SessionTimeout = "ThirtyMinutes"

	SessionTimeoutFifteenMinutes SessionTimeout = "FifteenMinutes"
)

type MonitoredEvents string

const (
	MonitoredEventsAuditTrail MonitoredEvents = "AuditTrail"

	MonitoredEventsLogin MonitoredEvents = "Login"

	MonitoredEventsEntity MonitoredEvents = "Entity"

	MonitoredEventsDataExport MonitoredEvents = "DataExport"

	MonitoredEventsAccessResource MonitoredEvents = "AccessResource"
)

type VisualizationResourceType string

const (
	VisualizationResourceTypeJs VisualizationResourceType = "js"

	VisualizationResourceTypeCss VisualizationResourceType = "css"
)

type LookupValueType string

const (
	LookupValueTypeUser LookupValueType = "User"

	LookupValueTypeQueue LookupValueType = "Queue"

	LookupValueTypeRecordType LookupValueType = "RecordType"
)

type FieldUpdateOperation string

const (
	FieldUpdateOperationFormula FieldUpdateOperation = "Formula"

	FieldUpdateOperationLiteral FieldUpdateOperation = "Literal"

	FieldUpdateOperationNull FieldUpdateOperation = "Null"

	FieldUpdateOperationNextValue FieldUpdateOperation = "NextValue"

	FieldUpdateOperationPreviousValue FieldUpdateOperation = "PreviousValue"

	FieldUpdateOperationLookupValue FieldUpdateOperation = "LookupValue"
)

type KnowledgeWorkflowAction string

const (
	KnowledgeWorkflowActionPublishAsNew KnowledgeWorkflowAction = "PublishAsNew"

	KnowledgeWorkflowActionPublish KnowledgeWorkflowAction = "Publish"
)

type SendAction string

const (
	SendActionSend SendAction = "Send"
)

type ActionTaskAssignedToTypes string

const (
	ActionTaskAssignedToTypesUser ActionTaskAssignedToTypes = "user"

	ActionTaskAssignedToTypesRole ActionTaskAssignedToTypes = "role"

	ActionTaskAssignedToTypesOpportunityTeam ActionTaskAssignedToTypes = "opportunityTeam"

	ActionTaskAssignedToTypesAccountTeam ActionTaskAssignedToTypes = "accountTeam"

	ActionTaskAssignedToTypesOwner ActionTaskAssignedToTypes = "owner"

	ActionTaskAssignedToTypesAccountOwner ActionTaskAssignedToTypes = "accountOwner"

	ActionTaskAssignedToTypesCreator ActionTaskAssignedToTypes = "creator"

	ActionTaskAssignedToTypesAccountCreator ActionTaskAssignedToTypes = "accountCreator"

	ActionTaskAssignedToTypesPartnerUser ActionTaskAssignedToTypes = "partnerUser"

	ActionTaskAssignedToTypesPortalRole ActionTaskAssignedToTypes = "portalRole"
)

type ActionEmailRecipientTypes string

const (
	ActionEmailRecipientTypesGroup ActionEmailRecipientTypes = "group"

	ActionEmailRecipientTypesRole ActionEmailRecipientTypes = "role"

	ActionEmailRecipientTypesUser ActionEmailRecipientTypes = "user"

	ActionEmailRecipientTypesOpportunityTeam ActionEmailRecipientTypes = "opportunityTeam"

	ActionEmailRecipientTypesAccountTeam ActionEmailRecipientTypes = "accountTeam"

	ActionEmailRecipientTypesRoleSubordinates ActionEmailRecipientTypes = "roleSubordinates"

	ActionEmailRecipientTypesOwner ActionEmailRecipientTypes = "owner"

	ActionEmailRecipientTypesCreator ActionEmailRecipientTypes = "creator"

	ActionEmailRecipientTypesPartnerUser ActionEmailRecipientTypes = "partnerUser"

	ActionEmailRecipientTypesAccountOwner ActionEmailRecipientTypes = "accountOwner"

	ActionEmailRecipientTypesCustomerPortalUser ActionEmailRecipientTypes = "customerPortalUser"

	ActionEmailRecipientTypesPortalRole ActionEmailRecipientTypes = "portalRole"

	ActionEmailRecipientTypesPortalRoleSubordinates ActionEmailRecipientTypes = "portalRoleSubordinates"

	ActionEmailRecipientTypesContactLookup ActionEmailRecipientTypes = "contactLookup"

	ActionEmailRecipientTypesUserLookup ActionEmailRecipientTypes = "userLookup"

	ActionEmailRecipientTypesRoleSubordinatesInternal ActionEmailRecipientTypes = "roleSubordinatesInternal"

	ActionEmailRecipientTypesEmail ActionEmailRecipientTypes = "email"

	ActionEmailRecipientTypesCaseTeam ActionEmailRecipientTypes = "caseTeam"

	ActionEmailRecipientTypesCampaignMemberDerivedOwner ActionEmailRecipientTypes = "campaignMemberDerivedOwner"
)

type ActionEmailSenderType string

const (
	ActionEmailSenderTypeCurrentUser ActionEmailSenderType = "CurrentUser"

	ActionEmailSenderTypeOrgWideEmailAddress ActionEmailSenderType = "OrgWideEmailAddress"

	ActionEmailSenderTypeDefaultWorkflowUser ActionEmailSenderType = "DefaultWorkflowUser"
)

type WorkflowTriggerTypes string

const (
	WorkflowTriggerTypesOnCreateOnly WorkflowTriggerTypes = "onCreateOnly"

	WorkflowTriggerTypesOnCreateOrTriggeringUpdate WorkflowTriggerTypes = "onCreateOrTriggeringUpdate"

	WorkflowTriggerTypesOnAllChanges WorkflowTriggerTypes = "onAllChanges"

	WorkflowTriggerTypesOnRecursiveUpdate WorkflowTriggerTypes = "OnRecursiveUpdate"
)

type WorkflowTimeUnits string

const (
	WorkflowTimeUnitsHours WorkflowTimeUnits = "Hours"

	WorkflowTimeUnitsDays WorkflowTimeUnits = "Days"
)

type ExtendedErrorCode string

const ()

type TestLevel string

const (
	TestLevelNoTestRun TestLevel = "NoTestRun"

	TestLevelRunSpecifiedTests TestLevel = "RunSpecifiedTests"

	TestLevelRunLocalTests TestLevel = "RunLocalTests"

	TestLevelRunAllTestsInOrg TestLevel = "RunAllTestsInOrg"
)

type AsyncRequestState string

const (
	AsyncRequestStateQueued AsyncRequestState = "Queued"

	AsyncRequestStateInProgress AsyncRequestState = "InProgress"

	AsyncRequestStateCompleted AsyncRequestState = "Completed"

	AsyncRequestStateError AsyncRequestState = "Error"
)

type LogCategory string

const (
	LogCategoryDb LogCategory = "Db"

	LogCategoryWorkflow LogCategory = "Workflow"

	LogCategoryValidation LogCategory = "Validation"

	LogCategoryCallout LogCategory = "Callout"

	LogCategoryApexcode LogCategory = "Apexcode"

	LogCategoryApexprofiling LogCategory = "Apexprofiling"

	LogCategoryVisualforce LogCategory = "Visualforce"

	LogCategorySystem LogCategory = "System"

	LogCategoryAll LogCategory = "All"
)

type LogCategoryLevel string

const (
	LogCategoryLevelNone LogCategoryLevel = "None"

	LogCategoryLevelFinest LogCategoryLevel = "Finest"

	LogCategoryLevelFiner LogCategoryLevel = "Finer"

	LogCategoryLevelFine LogCategoryLevel = "Fine"

	LogCategoryLevelDebug LogCategoryLevel = "Debug"

	LogCategoryLevelInfo LogCategoryLevel = "Info"

	LogCategoryLevelWarn LogCategoryLevel = "Warn"

	LogCategoryLevelError LogCategoryLevel = "Error"
)

type LogType string

const (
	LogTypeNone LogType = "None"

	LogTypeDebugonly LogType = "Debugonly"

	LogTypeDb LogType = "Db"

	LogTypeProfiling LogType = "Profiling"

	LogTypeCallout LogType = "Callout"

	LogTypeDetail LogType = "Detail"
)

type ID string

const ()

type StatusCode string

const (
	StatusCodeALLORNONEOPERATIONROLLEDBACK StatusCode = "ALLORNONEOPERATIONROLLEDBACK"

	StatusCodeALREADYINPROCESS StatusCode = "ALREADYINPROCESS"

	StatusCodeAPEXDATAACCESSRESTRICTION StatusCode = "APEXDATAACCESSRESTRICTION"

	StatusCodeASSIGNEETYPEREQUIRED StatusCode = "ASSIGNEETYPEREQUIRED"

	StatusCodeAURACOMPILEERROR StatusCode = "AURACOMPILEERROR"

	StatusCodeBADCUSTOMENTITYPARENTDOMAIN StatusCode = "BADCUSTOMENTITYPARENTDOMAIN"

	StatusCodeBCCNOTALLOWEDIFBCCCOMPLIANCEENABLED StatusCode = "BCCNOTALLOWEDIFBCCCOMPLIANCEENABLED"

	StatusCodeCANNOTCASCADEPRODUCTACTIVE StatusCode = "CANNOTCASCADEPRODUCTACTIVE"

	StatusCodeCANNOTCHANGEFIELDTYPEOFAPEXREFERENCEDFIELD StatusCode = "CANNOTCHANGEFIELDTYPEOFAPEXREFERENCEDFIELD"

	StatusCodeCANNOTCHANGEFIELDTYPEOFREFERENCEDFIELD StatusCode = "CANNOTCHANGEFIELDTYPEOFREFERENCEDFIELD"

	StatusCodeCANNOTCREATEANOTHERMANAGEDPACKAGE StatusCode = "CANNOTCREATEANOTHERMANAGEDPACKAGE"

	StatusCodeCANNOTDEACTIVATEDIVISION StatusCode = "CANNOTDEACTIVATEDIVISION"

	StatusCodeCANNOTDELETEGLOBALACTIONLIST StatusCode = "CANNOTDELETEGLOBALACTIONLIST"

	StatusCodeCANNOTDELETELASTDATEDCONVERSIONRATE StatusCode = "CANNOTDELETELASTDATEDCONVERSIONRATE"

	StatusCodeCANNOTDELETEMANAGEDOBJECT StatusCode = "CANNOTDELETEMANAGEDOBJECT"

	StatusCodeCANNOTDISABLELASTADMIN StatusCode = "CANNOTDISABLELASTADMIN"

	StatusCodeCANNOTENABLEIPRESTRICTREQUESTS StatusCode = "CANNOTENABLEIPRESTRICTREQUESTS"

	StatusCodeCANNOTEXECUTEFLOWTRIGGER StatusCode = "CANNOTEXECUTEFLOWTRIGGER"

	StatusCodeCANNOTFREEZESELF StatusCode = "CANNOTFREEZESELF"

	StatusCodeCANNOTINSERTUPDATEACTIVATEENTITY StatusCode = "CANNOTINSERTUPDATEACTIVATEENTITY"

	StatusCodeCANNOTMODIFYMANAGEDOBJECT StatusCode = "CANNOTMODIFYMANAGEDOBJECT"

	StatusCodeCANNOTPASSWORDLOCKOUT StatusCode = "CANNOTPASSWORDLOCKOUT"

	StatusCodeCANNOTPOSTTOARCHIVEDGROUP StatusCode = "CANNOTPOSTTOARCHIVEDGROUP"

	StatusCodeCANNOTRENAMEAPEXREFERENCEDFIELD StatusCode = "CANNOTRENAMEAPEXREFERENCEDFIELD"

	StatusCodeCANNOTRENAMEAPEXREFERENCEDOBJECT StatusCode = "CANNOTRENAMEAPEXREFERENCEDOBJECT"

	StatusCodeCANNOTRENAMEREFERENCEDFIELD StatusCode = "CANNOTRENAMEREFERENCEDFIELD"

	StatusCodeCANNOTRENAMEREFERENCEDOBJECT StatusCode = "CANNOTRENAMEREFERENCEDOBJECT"

	StatusCodeCANNOTREPARENTRECORD StatusCode = "CANNOTREPARENTRECORD"

	StatusCodeCANNOTUPDATECONVERTEDLEAD StatusCode = "CANNOTUPDATECONVERTEDLEAD"

	StatusCodeCANTDISABLECORPCURRENCY StatusCode = "CANTDISABLECORPCURRENCY"

	StatusCodeCANTUNSETCORPCURRENCY StatusCode = "CANTUNSETCORPCURRENCY"

	StatusCodeCHILDSHAREFAILSPARENT StatusCode = "CHILDSHAREFAILSPARENT"

	StatusCodeCIRCULARDEPENDENCY StatusCode = "CIRCULARDEPENDENCY"

	StatusCodeCLEANSERVICEERROR StatusCode = "CLEANSERVICEERROR"

	StatusCodeCOLLISIONDETECTED StatusCode = "COLLISIONDETECTED"

	StatusCodeCOMMUNITYNOTACCESSIBLE StatusCode = "COMMUNITYNOTACCESSIBLE"

	StatusCodeCONFLICTINGENVIRONMENTHUBMEMBER StatusCode = "CONFLICTINGENVIRONMENTHUBMEMBER"

	StatusCodeCONFLICTINGSSOUSERMAPPING StatusCode = "CONFLICTINGSSOUSERMAPPING"

	StatusCodeCUSTOMAPEXERROR StatusCode = "CUSTOMAPEXERROR"

	StatusCodeCUSTOMCLOBFIELDLIMITEXCEEDED StatusCode = "CUSTOMCLOBFIELDLIMITEXCEEDED"

	StatusCodeCUSTOMENTITYORFIELDLIMIT StatusCode = "CUSTOMENTITYORFIELDLIMIT"

	StatusCodeCUSTOMFIELDINDEXLIMITEXCEEDED StatusCode = "CUSTOMFIELDINDEXLIMITEXCEEDED"

	StatusCodeCUSTOMINDEXEXISTS StatusCode = "CUSTOMINDEXEXISTS"

	StatusCodeCUSTOMLINKLIMITEXCEEDED StatusCode = "CUSTOMLINKLIMITEXCEEDED"

	StatusCodeCUSTOMMETADATALIMITEXCEEDED StatusCode = "CUSTOMMETADATALIMITEXCEEDED"

	StatusCodeCUSTOMSETTINGSLIMITEXCEEDED StatusCode = "CUSTOMSETTINGSLIMITEXCEEDED"

	StatusCodeCUSTOMTABLIMITEXCEEDED StatusCode = "CUSTOMTABLIMITEXCEEDED"

	StatusCodeDATACLOUDADDRESSNORECORDSFOUND StatusCode = "DATACLOUDADDRESSNORECORDSFOUND"

	StatusCodeDATACLOUDADDRESSPROCESSINGERROR StatusCode = "DATACLOUDADDRESSPROCESSINGERROR"

	StatusCodeDATACLOUDADDRESSSERVERERROR StatusCode = "DATACLOUDADDRESSSERVERERROR"

	StatusCodeDELETEFAILED StatusCode = "DELETEFAILED"

	StatusCodeDELETEOPERATIONTOOLARGE StatusCode = "DELETEOPERATIONTOOLARGE"

	StatusCodeDELETEREQUIREDONCASCADE StatusCode = "DELETEREQUIREDONCASCADE"

	StatusCodeDEPENDENCYEXISTS StatusCode = "DEPENDENCYEXISTS"

	StatusCodeDUPLICATESDETECTED StatusCode = "DUPLICATESDETECTED"

	StatusCodeDUPLICATECASESOLUTION StatusCode = "DUPLICATECASESOLUTION"

	StatusCodeDUPLICATECOMMNICKNAME StatusCode = "DUPLICATECOMMNICKNAME"

	StatusCodeDUPLICATECUSTOMENTITYDEFINITION StatusCode = "DUPLICATECUSTOMENTITYDEFINITION"

	StatusCodeDUPLICATECUSTOMTABMOTIF StatusCode = "DUPLICATECUSTOMTABMOTIF"

	StatusCodeDUPLICATEDEVELOPERNAME StatusCode = "DUPLICATEDEVELOPERNAME"

	StatusCodeDUPLICATEEXTERNALID StatusCode = "DUPLICATEEXTERNALID"

	StatusCodeDUPLICATEMASTERLABEL StatusCode = "DUPLICATEMASTERLABEL"

	StatusCodeDUPLICATESENDERDISPLAYNAME StatusCode = "DUPLICATESENDERDISPLAYNAME"

	StatusCodeDUPLICATEUSERNAME StatusCode = "DUPLICATEUSERNAME"

	StatusCodeDUPLICATEVALUE StatusCode = "DUPLICATEVALUE"

	StatusCodeEMAILADDRESSBOUNCED StatusCode = "EMAILADDRESSBOUNCED"

	StatusCodeEMAILEXTERNALTRANSPORTCONNECTIONERROR StatusCode = "EMAILEXTERNALTRANSPORTCONNECTIONERROR"

	StatusCodeEMAILEXTERNALTRANSPORTTOKENERROR StatusCode = "EMAILEXTERNALTRANSPORTTOKENERROR"

	StatusCodeEMAILEXTERNALTRANSPORTTOOMANYREQUESTSERROR StatusCode = "EMAILEXTERNALTRANSPORTTOOMANYREQUESTSERROR"

	StatusCodeEMAILEXTERNALTRANSPORTUNKNOWNERROR StatusCode = "EMAILEXTERNALTRANSPORTUNKNOWNERROR"

	StatusCodeEMAILNOTPROCESSEDDUETOPRIORERROR StatusCode = "EMAILNOTPROCESSEDDUETOPRIORERROR"

	StatusCodeEMAILOPTEDOUT StatusCode = "EMAILOPTEDOUT"

	StatusCodeEMAILTEMPLATEFORMULAERROR StatusCode = "EMAILTEMPLATEFORMULAERROR"

	StatusCodeEMAILTEMPLATEMERGEFIELDACCESSERROR StatusCode = "EMAILTEMPLATEMERGEFIELDACCESSERROR"

	StatusCodeEMAILTEMPLATEMERGEFIELDERROR StatusCode = "EMAILTEMPLATEMERGEFIELDERROR"

	StatusCodeEMAILTEMPLATEMERGEFIELDVALUEERROR StatusCode = "EMAILTEMPLATEMERGEFIELDVALUEERROR"

	StatusCodeEMAILTEMPLATEPROCESSINGERROR StatusCode = "EMAILTEMPLATEPROCESSINGERROR"

	StatusCodeEMPTYSCONTROLFILENAME StatusCode = "EMPTYSCONTROLFILENAME"

	StatusCodeENTITYFAILEDIFLASTMODIFIEDONUPDATE StatusCode = "ENTITYFAILEDIFLASTMODIFIEDONUPDATE"

	StatusCodeENTITYISARCHIVED StatusCode = "ENTITYISARCHIVED"

	StatusCodeENTITYISDELETED StatusCode = "ENTITYISDELETED"

	StatusCodeENTITYISLOCKED StatusCode = "ENTITYISLOCKED"

	StatusCodeENTITYSAVEERROR StatusCode = "ENTITYSAVEERROR"

	StatusCodeENTITYSAVEVALIDATIONERROR StatusCode = "ENTITYSAVEVALIDATIONERROR"

	StatusCodeENVIRONMENTHUBMEMBERSHIPCONFLICT StatusCode = "ENVIRONMENTHUBMEMBERSHIPCONFLICT"

	StatusCodeENVIRONMENTHUBMEMBERSHIPERRORJOININGHUB StatusCode = "ENVIRONMENTHUBMEMBERSHIPERRORJOININGHUB"

	StatusCodeENVIRONMENTHUBMEMBERSHIPUSERALREADYINHUB StatusCode = "ENVIRONMENTHUBMEMBERSHIPUSERALREADYINHUB"

	StatusCodeENVIRONMENTHUBMEMBERSHIPUSERNOTORGADMIN StatusCode = "ENVIRONMENTHUBMEMBERSHIPUSERNOTORGADMIN"

	StatusCodeERRORINMAILER StatusCode = "ERRORINMAILER"

	StatusCodeEXCHANGEWEBSERVICESURLINVALID StatusCode = "EXCHANGEWEBSERVICESURLINVALID"

	StatusCodeFAILEDACTIVATION StatusCode = "FAILEDACTIVATION"

	StatusCodeFIELDCUSTOMVALIDATIONEXCEPTION StatusCode = "FIELDCUSTOMVALIDATIONEXCEPTION"

	StatusCodeFIELDFILTERVALIDATIONEXCEPTION StatusCode = "FIELDFILTERVALIDATIONEXCEPTION"

	StatusCodeFIELDINTEGRITYEXCEPTION StatusCode = "FIELDINTEGRITYEXCEPTION"

	StatusCodeFIELDKEYWORDLISTMATCHLIMIT StatusCode = "FIELDKEYWORDLISTMATCHLIMIT"

	StatusCodeFIELDMAPPINGERROR StatusCode = "FIELDMAPPINGERROR"

	StatusCodeFIELDMODERATIONRULEBLOCK StatusCode = "FIELDMODERATIONRULEBLOCK"

	StatusCodeFIELDNOTUPDATABLE StatusCode = "FIELDNOTUPDATABLE"

	StatusCodeFILEEXTENSIONNOTALLOWED StatusCode = "FILEEXTENSIONNOTALLOWED"

	StatusCodeFILESIZELIMITEXCEEDED StatusCode = "FILESIZELIMITEXCEEDED"

	StatusCodeFILTEREDLOOKUPLIMITEXCEEDED StatusCode = "FILTEREDLOOKUPLIMITEXCEEDED"

	StatusCodeFINDDUPLICATESERROR StatusCode = "FINDDUPLICATESERROR"

	StatusCodeFUNCTIONALITYNOTENABLED StatusCode = "FUNCTIONALITYNOTENABLED"

	StatusCodeHASPUBLICREFERENCES StatusCode = "HASPUBLICREFERENCES"

	StatusCodeHTMLFILEUPLOADNOTALLOWED StatusCode = "HTMLFILEUPLOADNOTALLOWED"

	StatusCodeIMAGETOOLARGE StatusCode = "IMAGETOOLARGE"

	StatusCodeINACTIVEOWNERORUSER StatusCode = "INACTIVEOWNERORUSER"

	StatusCodeINACTIVERULEERROR StatusCode = "INACTIVERULEERROR"

	StatusCodeINSERTUPDATEDELETENOTALLOWEDDURINGMAINTENANCE StatusCode = "INSERTUPDATEDELETENOTALLOWEDDURINGMAINTENANCE"

	StatusCodeINSUFFICIENTACCESSONCROSSREFERENCEENTITY StatusCode = "INSUFFICIENTACCESSONCROSSREFERENCEENTITY"

	StatusCodeINSUFFICIENTACCESSORREADONLY StatusCode = "INSUFFICIENTACCESSORREADONLY"

	StatusCodeINSUFFICIENTACCESSTOINSIGHTSEXTERNALDATA StatusCode = "INSUFFICIENTACCESSTOINSIGHTSEXTERNALDATA"

	StatusCodeINSUFFICIENTCREDITS StatusCode = "INSUFFICIENTCREDITS"

	StatusCodeINVALIDACCESSLEVEL StatusCode = "INVALIDACCESSLEVEL"

	StatusCodeINVALIDARGUMENTTYPE StatusCode = "INVALIDARGUMENTTYPE"

	StatusCodeINVALIDASSIGNEETYPE StatusCode = "INVALIDASSIGNEETYPE"

	StatusCodeINVALIDASSIGNMENTRULE StatusCode = "INVALIDASSIGNMENTRULE"

	StatusCodeINVALIDBATCHOPERATION StatusCode = "INVALIDBATCHOPERATION"

	StatusCodeINVALIDCONTENTTYPE StatusCode = "INVALIDCONTENTTYPE"

	StatusCodeINVALIDCREDITCARDINFO StatusCode = "INVALIDCREDITCARDINFO"

	StatusCodeINVALIDCROSSREFERENCEKEY StatusCode = "INVALIDCROSSREFERENCEKEY"

	StatusCodeINVALIDCROSSREFERENCETYPEFORFIELD StatusCode = "INVALIDCROSSREFERENCETYPEFORFIELD"

	StatusCodeINVALIDCURRENCYCONVRATE StatusCode = "INVALIDCURRENCYCONVRATE"

	StatusCodeINVALIDCURRENCYCORPRATE StatusCode = "INVALIDCURRENCYCORPRATE"

	StatusCodeINVALIDCURRENCYISO StatusCode = "INVALIDCURRENCYISO"

	StatusCodeINVALIDDATACATEGORYGROUPREFERENCE StatusCode = "INVALIDDATACATEGORYGROUPREFERENCE"

	StatusCodeINVALIDDATAURI StatusCode = "INVALIDDATAURI"

	StatusCodeINVALIDEMAILADDRESS StatusCode = "INVALIDEMAILADDRESS"

	StatusCodeINVALIDEMPTYKEYOWNER StatusCode = "INVALIDEMPTYKEYOWNER"

	StatusCodeINVALIDENTITYFORMATCHENGINEERROR StatusCode = "INVALIDENTITYFORMATCHENGINEERROR"

	StatusCodeINVALIDENTITYFORMATCHOPERATIONERROR StatusCode = "INVALIDENTITYFORMATCHOPERATIONERROR"

	StatusCodeINVALIDENTITYFORUPSERT StatusCode = "INVALIDENTITYFORUPSERT"

	StatusCodeINVALIDENVIRONMENTHUBMEMBER StatusCode = "INVALIDENVIRONMENTHUBMEMBER"

	StatusCodeINVALIDEVENTDELIVERY StatusCode = "INVALIDEVENTDELIVERY"

	StatusCodeINVALIDEVENTSUBSCRIPTION StatusCode = "INVALIDEVENTSUBSCRIPTION"

	StatusCodeINVALIDFIELD StatusCode = "INVALIDFIELD"

	StatusCodeINVALIDFIELDFORINSERTUPDATE StatusCode = "INVALIDFIELDFORINSERTUPDATE"

	StatusCodeINVALIDFIELDWHENUSINGTEMPLATE StatusCode = "INVALIDFIELDWHENUSINGTEMPLATE"

	StatusCodeINVALIDFILTERACTION StatusCode = "INVALIDFILTERACTION"

	StatusCodeINVALIDGOOGLEDOCSURL StatusCode = "INVALIDGOOGLEDOCSURL"

	StatusCodeINVALIDIDFIELD StatusCode = "INVALIDIDFIELD"

	StatusCodeINVALIDINETADDRESS StatusCode = "INVALIDINETADDRESS"

	StatusCodeINVALIDINPUT StatusCode = "INVALIDINPUT"

	StatusCodeINVALIDLINEITEMCLONESTATE StatusCode = "INVALIDLINEITEMCLONESTATE"

	StatusCodeINVALIDMARKUP StatusCode = "INVALIDMARKUP"

	StatusCodeINVALIDMASTERORTRANSLATEDSOLUTION StatusCode = "INVALIDMASTERORTRANSLATEDSOLUTION"

	StatusCodeINVALIDMESSAGEIDREFERENCE StatusCode = "INVALIDMESSAGEIDREFERENCE"

	StatusCodeINVALIDNAMESPACEPREFIX StatusCode = "INVALIDNAMESPACEPREFIX"

	StatusCodeINVALIDOAUTHURL StatusCode = "INVALIDOAUTHURL"

	StatusCodeINVALIDOPERATION StatusCode = "INVALIDOPERATION"

	StatusCodeINVALIDOPERATOR StatusCode = "INVALIDOPERATOR"

	StatusCodeINVALIDORNULLFORRESTRICTEDPICKLIST StatusCode = "INVALIDORNULLFORRESTRICTEDPICKLIST"

	StatusCodeINVALIDOWNER StatusCode = "INVALIDOWNER"

	StatusCodeINVALIDPACKAGELICENSE StatusCode = "INVALIDPACKAGELICENSE"

	StatusCodeINVALIDPACKAGEVERSION StatusCode = "INVALIDPACKAGEVERSION"

	StatusCodeINVALIDPARTNERNETWORKSTATUS StatusCode = "INVALIDPARTNERNETWORKSTATUS"

	StatusCodeINVALIDPERSONACCOUNTOPERATION StatusCode = "INVALIDPERSONACCOUNTOPERATION"

	StatusCodeINVALIDQUERYLOCATOR StatusCode = "INVALIDQUERYLOCATOR"

	StatusCodeINVALIDREADONLYUSERDML StatusCode = "INVALIDREADONLYUSERDML"

	StatusCodeINVALIDRUNTIMEVALUE StatusCode = "INVALIDRUNTIMEVALUE"

	StatusCodeINVALIDSAVEASACTIVITYFLAG StatusCode = "INVALIDSAVEASACTIVITYFLAG"

	StatusCodeINVALIDSESSIONID StatusCode = "INVALIDSESSIONID"

	StatusCodeINVALIDSETUPOWNER StatusCode = "INVALIDSETUPOWNER"

	StatusCodeINVALIDSIGNUPCOUNTRY StatusCode = "INVALIDSIGNUPCOUNTRY"

	StatusCodeINVALIDSIGNUPOPTION StatusCode = "INVALIDSIGNUPOPTION"

	StatusCodeINVALIDSITEDELETEEXCEPTION StatusCode = "INVALIDSITEDELETEEXCEPTION"

	StatusCodeINVALIDSITEFILEIMPORTEDEXCEPTION StatusCode = "INVALIDSITEFILEIMPORTEDEXCEPTION"

	StatusCodeINVALIDSITEFILETYPEEXCEPTION StatusCode = "INVALIDSITEFILETYPEEXCEPTION"

	StatusCodeINVALIDSTATUS StatusCode = "INVALIDSTATUS"

	StatusCodeINVALIDSUBDOMAIN StatusCode = "INVALIDSUBDOMAIN"

	StatusCodeINVALIDTYPE StatusCode = "INVALIDTYPE"

	StatusCodeINVALIDTYPEFOROPERATION StatusCode = "INVALIDTYPEFOROPERATION"

	StatusCodeINVALIDTYPEONFIELDINRECORD StatusCode = "INVALIDTYPEONFIELDINRECORD"

	StatusCodeINVALIDUSERID StatusCode = "INVALIDUSERID"

	StatusCodeIPRANGELIMITEXCEEDED StatusCode = "IPRANGELIMITEXCEEDED"

	StatusCodeJIGSAWIMPORTLIMITEXCEEDED StatusCode = "JIGSAWIMPORTLIMITEXCEEDED"

	StatusCodeLICENSELIMITEXCEEDED StatusCode = "LICENSELIMITEXCEEDED"

	StatusCodeLIGHTPORTALUSEREXCEPTION StatusCode = "LIGHTPORTALUSEREXCEPTION"

	StatusCodeLIMITEXCEEDED StatusCode = "LIMITEXCEEDED"

	StatusCodeMALFORMEDID StatusCode = "MALFORMEDID"

	StatusCodeMANAGERNOTDEFINED StatusCode = "MANAGERNOTDEFINED"

	StatusCodeMASSMAILRETRYLIMITEXCEEDED StatusCode = "MASSMAILRETRYLIMITEXCEEDED"

	StatusCodeMASSMAILLIMITEXCEEDED StatusCode = "MASSMAILLIMITEXCEEDED"

	StatusCodeMATCHDEFINITIONERROR StatusCode = "MATCHDEFINITIONERROR"

	StatusCodeMATCHOPERATIONERROR StatusCode = "MATCHOPERATIONERROR"

	StatusCodeMATCHOPERATIONINVALIDENGINEERROR StatusCode = "MATCHOPERATIONINVALIDENGINEERROR"

	StatusCodeMATCHOPERATIONINVALIDRULEERROR StatusCode = "MATCHOPERATIONINVALIDRULEERROR"

	StatusCodeMATCHOPERATIONMISSINGENGINEERROR StatusCode = "MATCHOPERATIONMISSINGENGINEERROR"

	StatusCodeMATCHOPERATIONMISSINGOBJECTTYPEERROR StatusCode = "MATCHOPERATIONMISSINGOBJECTTYPEERROR"

	StatusCodeMATCHOPERATIONMISSINGOPTIONSERROR StatusCode = "MATCHOPERATIONMISSINGOPTIONSERROR"

	StatusCodeMATCHOPERATIONMISSINGRULEERROR StatusCode = "MATCHOPERATIONMISSINGRULEERROR"

	StatusCodeMATCHOPERATIONUNKNOWNRULEERROR StatusCode = "MATCHOPERATIONUNKNOWNRULEERROR"

	StatusCodeMATCHOPERATIONUNSUPPORTEDVERSIONERROR StatusCode = "MATCHOPERATIONUNSUPPORTEDVERSIONERROR"

	StatusCodeMATCHPRECONDITIONFAILED StatusCode = "MATCHPRECONDITIONFAILED"

	StatusCodeMATCHRUNTIMEERROR StatusCode = "MATCHRUNTIMEERROR"

	StatusCodeMATCHSERVICEERROR StatusCode = "MATCHSERVICEERROR"

	StatusCodeMATCHSERVICETIMEDOUT StatusCode = "MATCHSERVICETIMEDOUT"

	StatusCodeMATCHSERVICEUNAVAILABLEERROR StatusCode = "MATCHSERVICEUNAVAILABLEERROR"

	StatusCodeMAXIMUMCCEMAILSEXCEEDED StatusCode = "MAXIMUMCCEMAILSEXCEEDED"

	StatusCodeMAXIMUMDASHBOARDCOMPONENTSEXCEEDED StatusCode = "MAXIMUMDASHBOARDCOMPONENTSEXCEEDED"

	StatusCodeMAXIMUMHIERARCHYCHILDRENREACHED StatusCode = "MAXIMUMHIERARCHYCHILDRENREACHED"

	StatusCodeMAXIMUMHIERARCHYLEVELSREACHED StatusCode = "MAXIMUMHIERARCHYLEVELSREACHED"

	StatusCodeMAXIMUMHIERARCHYTREESIZEREACHED StatusCode = "MAXIMUMHIERARCHYTREESIZEREACHED"

	StatusCodeMAXIMUMSIZEOFATTACHMENT StatusCode = "MAXIMUMSIZEOFATTACHMENT"

	StatusCodeMAXIMUMSIZEOFDOCUMENT StatusCode = "MAXIMUMSIZEOFDOCUMENT"

	StatusCodeMAXACTIONSPERRULEEXCEEDED StatusCode = "MAXACTIONSPERRULEEXCEEDED"

	StatusCodeMAXACTIVERULESEXCEEDED StatusCode = "MAXACTIVERULESEXCEEDED"

	StatusCodeMAXAPPROVALSTEPSEXCEEDED StatusCode = "MAXAPPROVALSTEPSEXCEEDED"

	StatusCodeMAXDEPTHINFLOWEXECUTION StatusCode = "MAXDEPTHINFLOWEXECUTION"

	StatusCodeMAXFORMULASPERRULEEXCEEDED StatusCode = "MAXFORMULASPERRULEEXCEEDED"

	StatusCodeMAXRULESEXCEEDED StatusCode = "MAXRULESEXCEEDED"

	StatusCodeMAXRULEENTRIESEXCEEDED StatusCode = "MAXRULEENTRIESEXCEEDED"

	StatusCodeMAXTASKDESCRIPTIONEXCEEEDED StatusCode = "MAXTASKDESCRIPTIONEXCEEEDED"

	StatusCodeMAXTMRULESEXCEEDED StatusCode = "MAXTMRULESEXCEEDED"

	StatusCodeMAXTMRULEITEMSEXCEEDED StatusCode = "MAXTMRULEITEMSEXCEEDED"

	StatusCodeMERGEFAILED StatusCode = "MERGEFAILED"

	StatusCodeMETADATAFIELDUPDATEERROR StatusCode = "METADATAFIELDUPDATEERROR"

	StatusCodeMISSINGARGUMENT StatusCode = "MISSINGARGUMENT"

	StatusCodeMISSINGRECORD StatusCode = "MISSINGRECORD"

	StatusCodeMIXEDDMLOPERATION StatusCode = "MIXEDDMLOPERATION"

	StatusCodeNONUNIQUESHIPPINGADDRESS StatusCode = "NONUNIQUESHIPPINGADDRESS"

	StatusCodeNOAPPLICABLEPROCESS StatusCode = "NOAPPLICABLEPROCESS"

	StatusCodeNOATTACHMENTPERMISSION StatusCode = "NOATTACHMENTPERMISSION"

	StatusCodeNOINACTIVEDIVISIONMEMBERS StatusCode = "NOINACTIVEDIVISIONMEMBERS"

	StatusCodeNOMASSMAILPERMISSION StatusCode = "NOMASSMAILPERMISSION"

	StatusCodeNOPARTNERPERMISSION StatusCode = "NOPARTNERPERMISSION"

	StatusCodeNOSUCHUSEREXISTS StatusCode = "NOSUCHUSEREXISTS"

	StatusCodeNUMBEROUTSIDEVALIDRANGE StatusCode = "NUMBEROUTSIDEVALIDRANGE"

	StatusCodeNUMHISTORYFIELDSBYSOBJECTEXCEEDED StatusCode = "NUMHISTORYFIELDSBYSOBJECTEXCEEDED"

	StatusCodeOPTEDOUTOFMASSMAIL StatusCode = "OPTEDOUTOFMASSMAIL"

	StatusCodeOPWITHINVALIDUSERTYPEEXCEPTION StatusCode = "OPWITHINVALIDUSERTYPEEXCEPTION"

	StatusCodePACKAGELICENSEREQUIRED StatusCode = "PACKAGELICENSEREQUIRED"

	StatusCodePACKAGINGAPIINSTALLFAILED StatusCode = "PACKAGINGAPIINSTALLFAILED"

	StatusCodePACKAGINGAPIUNINSTALLFAILED StatusCode = "PACKAGINGAPIUNINSTALLFAILED"

	StatusCodePALIINVALIDACTIONID StatusCode = "PALIINVALIDACTIONID"

	StatusCodePALIINVALIDACTIONNAME StatusCode = "PALIINVALIDACTIONNAME"

	StatusCodePALIINVALIDACTIONTYPE StatusCode = "PALIINVALIDACTIONTYPE"

	StatusCodePALINVALIDASSISTANTRECOMMENDATIONTYPEID StatusCode = "PALINVALIDASSISTANTRECOMMENDATIONTYPEID"

	StatusCodePALINVALIDENTITYID StatusCode = "PALINVALIDENTITYID"

	StatusCodePALINVALIDFLEXIPAGEID StatusCode = "PALINVALIDFLEXIPAGEID"

	StatusCodePALINVALIDLAYOUTID StatusCode = "PALINVALIDLAYOUTID"

	StatusCodePALINVALIDPARAMETERS StatusCode = "PALINVALIDPARAMETERS"

	StatusCodePAAPIEXCEPTION StatusCode = "PAAPIEXCEPTION"

	StatusCodePAAXISFAULT StatusCode = "PAAXISFAULT"

	StatusCodePAINVALIDIDEXCEPTION StatusCode = "PAINVALIDIDEXCEPTION"

	StatusCodePANOACCESSEXCEPTION StatusCode = "PANOACCESSEXCEPTION"

	StatusCodePANODATAFOUNDEXCEPTION StatusCode = "PANODATAFOUNDEXCEPTION"

	StatusCodePAURISYNTAXEXCEPTION StatusCode = "PAURISYNTAXEXCEPTION"

	StatusCodePAVISIBLEACTIONSFILTERORDERINGEXCEPTION StatusCode = "PAVISIBLEACTIONSFILTERORDERINGEXCEPTION"

	StatusCodePORTALNOACCESS StatusCode = "PORTALNOACCESS"

	StatusCodePORTALUSERALREADYEXISTSFORCONTACT StatusCode = "PORTALUSERALREADYEXISTSFORCONTACT"

	StatusCodePORTALUSERCREATIONRESTRICTEDWITHENCRYPTION StatusCode = "PORTALUSERCREATIONRESTRICTEDWITHENCRYPTION"

	StatusCodePRIVATECONTACTONASSET StatusCode = "PRIVATECONTACTONASSET"

	StatusCodePROCESSINGHALTED StatusCode = "PROCESSINGHALTED"

	StatusCodeQAINVALIDCREATEFEEDITEM StatusCode = "QAINVALIDCREATEFEEDITEM"

	StatusCodeQAINVALIDSUCCESSMESSAGE StatusCode = "QAINVALIDSUCCESSMESSAGE"

	StatusCodeQUERYTIMEOUT StatusCode = "QUERYTIMEOUT"

	StatusCodeQUICKACTIONLISTITEMNOTALLOWED StatusCode = "QUICKACTIONLISTITEMNOTALLOWED"

	StatusCodeQUICKACTIONLISTNOTALLOWED StatusCode = "QUICKACTIONLISTNOTALLOWED"

	StatusCodeRECORDINUSEBYWORKFLOW StatusCode = "RECORDINUSEBYWORKFLOW"

	StatusCodeRELFIELDBADACCESSIBILITY StatusCode = "RELFIELDBADACCESSIBILITY"

	StatusCodeREPUTATIONMINIMUMNUMBERNOTREACHED StatusCode = "REPUTATIONMINIMUMNUMBERNOTREACHED"

	StatusCodeREQUESTRUNNINGTOOLONG StatusCode = "REQUESTRUNNINGTOOLONG"

	StatusCodeREQUIREDFEATUREMISSING StatusCode = "REQUIREDFEATUREMISSING"

	StatusCodeREQUIREDFIELDMISSING StatusCode = "REQUIREDFIELDMISSING"

	StatusCodeRETRIEVEEXCHANGEATTACHMENTFAILED StatusCode = "RETRIEVEEXCHANGEATTACHMENTFAILED"

	StatusCodeRETRIEVEEXCHANGEEMAILFAILED StatusCode = "RETRIEVEEXCHANGEEMAILFAILED"

	StatusCodeRETRIEVEEXCHANGEEVENTFAILED StatusCode = "RETRIEVEEXCHANGEEVENTFAILED"

	StatusCodeOBJFORCEINBOXTRANSPORTCONNECTIONERROR StatusCode = "OBJFORCEINBOXTRANSPORTCONNECTIONERROR"

	StatusCodeOBJFORCEINBOXTRANSPORTTOKENERROR StatusCode = "OBJFORCEINBOXTRANSPORTTOKENERROR"

	StatusCodeOBJFORCEINBOXTRANSPORTUNKNOWNERROR StatusCode = "OBJFORCEINBOXTRANSPORTUNKNOWNERROR"

	StatusCodeSELFREFERENCEFROMFLOW StatusCode = "SELFREFERENCEFROMFLOW"

	StatusCodeSELFREFERENCEFROMTRIGGER StatusCode = "SELFREFERENCEFROMTRIGGER"

	StatusCodeSHARENEEDEDFORCHILDOWNER StatusCode = "SHARENEEDEDFORCHILDOWNER"

	StatusCodeSINGLEEMAILLIMITEXCEEDED StatusCode = "SINGLEEMAILLIMITEXCEEDED"

	StatusCodeSOCIALACCOUNTNOTFOUND StatusCode = "SOCIALACCOUNTNOTFOUND"

	StatusCodeSOCIALACTIONINVALID StatusCode = "SOCIALACTIONINVALID"

	StatusCodeSOCIALPOSTINVALID StatusCode = "SOCIALPOSTINVALID"

	StatusCodeSOCIALPOSTNOTFOUND StatusCode = "SOCIALPOSTNOTFOUND"

	StatusCodeSTANDARDPRICENOTDEFINED StatusCode = "STANDARDPRICENOTDEFINED"

	StatusCodeSTORAGELIMITEXCEEDED StatusCode = "STORAGELIMITEXCEEDED"

	StatusCodeSTRINGTOOLONG StatusCode = "STRINGTOOLONG"

	StatusCodeSUBDOMAININUSE StatusCode = "SUBDOMAININUSE"

	StatusCodeTABSETLIMITEXCEEDED StatusCode = "TABSETLIMITEXCEEDED"

	StatusCodeTEMPLATENOTACTIVE StatusCode = "TEMPLATENOTACTIVE"

	StatusCodeTEMPLATENOTFOUND StatusCode = "TEMPLATENOTFOUND"

	StatusCodeTERRITORYREALIGNINPROGRESS StatusCode = "TERRITORYREALIGNINPROGRESS"

	StatusCodeTEXTDATAOUTSIDESUPPORTEDCHARSET StatusCode = "TEXTDATAOUTSIDESUPPORTEDCHARSET"

	StatusCodeTOOMANYAPEXREQUESTS StatusCode = "TOOMANYAPEXREQUESTS"

	StatusCodeTOOMANYENUMVALUE StatusCode = "TOOMANYENUMVALUE"

	StatusCodeTOOMANYPOSSIBLEUSERSEXIST StatusCode = "TOOMANYPOSSIBLEUSERSEXIST"

	StatusCodeTRANSFERREQUIRESREAD StatusCode = "TRANSFERREQUIRESREAD"

	StatusCodeUNABLETOLOCKROW StatusCode = "UNABLETOLOCKROW"

	StatusCodeUNAVAILABLERECORDTYPEEXCEPTION StatusCode = "UNAVAILABLERECORDTYPEEXCEPTION"

	StatusCodeUNAVAILABLEREF StatusCode = "UNAVAILABLEREF"

	StatusCodeUNDELETEFAILED StatusCode = "UNDELETEFAILED"

	StatusCodeUNKNOWNEXCEPTION StatusCode = "UNKNOWNEXCEPTION"

	StatusCodeUNSAFEHTMLCONTENT StatusCode = "UNSAFEHTMLCONTENT"

	StatusCodeUNSPECIFIEDEMAILADDRESS StatusCode = "UNSPECIFIEDEMAILADDRESS"

	StatusCodeUNSUPPORTEDAPEXTRIGGEROPERATON StatusCode = "UNSUPPORTEDAPEXTRIGGEROPERATON"

	StatusCodeUNVERIFIEDSENDERADDRESS StatusCode = "UNVERIFIEDSENDERADDRESS"

	StatusCodeUSEROWNSPORTALACCOUNTEXCEPTION StatusCode = "USEROWNSPORTALACCOUNTEXCEPTION"

	StatusCodeUSERWITHAPEXSHARESEXCEPTION StatusCode = "USERWITHAPEXSHARESEXCEPTION"

	StatusCodeVFCOMPILEERROR StatusCode = "VFCOMPILEERROR"

	StatusCodeWEBLINKSIZELIMITEXCEEDED StatusCode = "WEBLINKSIZELIMITEXCEEDED"

	StatusCodeWEBLINKURLINVALID StatusCode = "WEBLINKURLINVALID"

	StatusCodeWRONGCONTROLLERTYPE StatusCode = "WRONGCONTROLLERTYPE"

	StatusCodeXCLEANUNEXPECTEDERROR StatusCode = "XCLEANUNEXPECTEDERROR"
)

type AllOrNoneHeader struct {
	Name string `json:"http://soap.sforce.com/2006/04/metadata AllOrNoneHeader"`

	AllOrNone bool `json:"allOrNone,omitempty"`
}

type CallOptions struct {
	Name string `json:"http://soap.sforce.com/2006/04/metadata CallOptions"`
	Client string `json:"client,omitempty"`
}

type DebuggingHeader struct {
	Name string `json:"http://soap.sforce.com/2006/04/metadata DebuggingHeader"`

	Categories []*LogInfo `json:"categories,omitempty"`

	DebugLevel *LogType `json:"debugLevel,omitempty"`
}

type DebuggingInfo struct {
	Name string `json:"http://soap.sforce.com/2006/04/metadata DebuggingInfo"`

	DebugLog string `json:"debugLog,omitempty"`
}

type SessionHeader struct {
	Name string `json:"http://soap.sforce.com/2006/04/metadata SessionHeader"`

	SessionId string `json:"sessionId,omitempty"`
}

type CancelDeploy struct {
	Name string `json:"http://soap.sforce.com/2006/04/metadata cancelDeploy"`

	AsyncProcessId ID `json:"String,omitempty"`
}

type CancelDeployResponse struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata cancelDeployResponse"`

	Result *CancelDeployResult `json:"result,omitempty"`
}

type CheckDeployStatus struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata checkDeployStatus"`

	AsyncProcessId ID `json:"asyncProcessId,omitempty"`

	IncludeDetails bool `json:"includeDetails,omitempty"`
}

type CheckDeployStatusResponse struct {
	Name string `json:"http://soap.sforce.com/2006/04/metadata checkDeployStatusResponse"`

	Result *DeployResult `json:"result,omitempty"`
}

type CheckRetrieveStatus struct {
	Name string `json:"http://soap.sforce.com/2006/04/metadata checkRetrieveStatus"`

	AsyncProcessId ID `json:"asyncProcessId,omitempty"`

	IncludeZip bool `json:"includeZip,omitempty"`
}

type CheckRetrieveStatusResponse struct {
	Name string `json:"http://soap.sforce.com/2006/04/metadata checkRetrieveStatusResponse"`

	Result *RetrieveResult `json:"result,omitempty"`
}

type CreateMetadata struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata createMetadata"`

	Metadata []MetadataInterface `json:"metadata,omitempty"`
}

type CreateMetadataResponse struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata createMetadataResponse"`

	Result []*SaveResult `json:"result,omitempty"`
}

type DeleteMetadata struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata deleteMetadata"`

	Type string `json:"type,omitempty"`

	FullNames []string `json:"fullNames,omitempty"`
}

type DeleteMetadataResponse struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata deleteMetadataResponse"`

	Result []*DeleteResult `json:"result,omitempty"`
}

type Deploy struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata deploy"`

	ZipFile string `json:"ZipFile,omitempty"`

	DeployOptions *DeployOptions `json:"DeployOptions,omitempty"`
}

type DeployResponse struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata deployResponse"`

	Result *AsyncResult `json:"result,omitempty"`
}

type DeployRecentValidation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata deployRecentValidation"`

	ValidationId ID `json:"validationId,omitempty"`
}

type DeployRecentValidationResponse struct {
	Name string `json:"http://soap.sforce.com/2006/04/metadata deployRecentValidationResponse"`

	Result string `json:"result,omitempty"`
}

type DescribeMetadata struct {
	Name string `json:"http://soap.sforce.com/2006/04/metadata describeMetadata"`

	AsOfVersion float64 `json:"asOfVersion,omitempty"`
}

type DescribeMetadataResponse struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata describeMetadataResponse"`

	Result *DescribeMetadataResult `json:"result,omitempty"`
}

type DescribeValueType struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata describeValueType"`

	Type string `json:"type,omitempty"`
}

type DescribeValueTypeResponse struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata describeValueTypeResponse"`

	Result *DescribeValueTypeResult `json:"result,omitempty"`
}

type ListMetadata struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata listMetadata"`

	Queries []*ListMetadataQuery `json:"queries,omitempty"`

	AsOfVersion float64 `json:"asOfVersion,omitempty"`
}

type ListMetadataResponse struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata listMetadataResponse"`

	Result []*FileProperties `json:"result,omitempty"`
}

type ReadMetadata struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata readMetadata"`

	Type string `json:"type,omitempty"`

	FullNames []string `json:"fullNames,omitempty"`
}

type ReadMetadataResponse struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata readMetadataResponse"`

	Result *ReadResult `json:"result,omitempty"`
}

type RenameMetadata struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata renameMetadata"`

	Type string `json:"type,omitempty"`

	OldFullName string `json:"oldFullName,omitempty"`

	NewFullName string `json:"newFullName,omitempty"`
}

type RenameMetadataResponse struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata renameMetadataResponse"`

	Result *SaveResult `json:"result,omitempty"`
}

type Retrieve struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata retrieve"`

	RetrieveRequest *RetrieveRequest `json:"retrieveRequest,omitempty"`
}

type RetrieveResponse struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata retrieveResponse"`

	Result *AsyncResult `json:"result,omitempty"`
}

type UpdateMetadata struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata updateMetadata"`

	Metadata []MetadataInterface `json:"metadata,omitempty"`
}

type UpdateMetadataResponse struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata updateMetadataResponse"`

	Result []*SaveResult `json:"result,omitempty"`
}

type UpsertMetadata struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata upsertMetadata"`

	Metadata []MetadataInterface `json:"metadata,omitempty"`
}

type UpsertMetadataResponse struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata upsertMetadataResponse"`

	Result []*UpsertResult `json:"result,omitempty"`
}

type CancelDeployResult struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CancelDeployResult"`

	Done bool `json:"done,omitempty"`

	Id ID `json:"id,omitempty"`
}

type DeployResult struct {
	//XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DeployResult"`

	CanceledBy string `json:"canceledBy,omitempty"`

	CanceledByName string `json:"canceledByName,omitempty"`

	CheckOnly bool `json:"checkOnly,omitempty"`

	CompletedDate time.Time `json:"completedDate,omitempty"`

	CreatedBy string `json:"createdBy,omitempty"`

	CreatedByName string `json:"createdByName,omitempty"`

	CreatedDate time.Time `json:"createdDate,omitempty"`

	Details *DeployDetails `json:"details,omitempty"`

	Done bool `json:"done,omitempty"`

	ErrorMessage string `json:"errorMessage,omitempty"`

	ErrorStatusCode StatusCode `json:"errorStatusCode,omitempty"`

	Id ID `json:"id,omitempty"`

	IgnoreWarnings bool `json:"ignoreWarnings,omitempty"`

	LastModifiedDate time.Time `json:"lastModifiedDate,omitempty"`

	NumberComponentErrors int32 `json:"numberComponentErrors,omitempty"`

	NumberComponentsDeployed int32 `json:"numberComponentsDeployed,omitempty"`

	NumberComponentsTotal int32 `json:"numberComponentsTotal,omitempty"`

	NumberTestErrors int32 `json:"numberTestErrors,omitempty"`

	NumberTestsCompleted int32 `json:"numberTestsCompleted,omitempty"`

	NumberTestsTotal int32 `json:"numberTestsTotal,omitempty"`

	RollbackOnError bool `json:"rollbackOnError,omitempty"`

	RunTestsEnabled bool `json:"runTestsEnabled,omitempty"`

	StartDate time.Time `json:"startDate,omitempty"`

	StateDetail string `json:"stateDetail,omitempty"`

	Status *DeployStatus `json:"status,omitempty"`

	Success bool `json:"success,omitempty"`
}

type DeployDetails struct {
	//XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DeployDetails"`

	ComponentFailures []*DeployMessage `json:"componentFailures,omitempty"`

	ComponentSuccesses []*DeployMessage `json:"componentSuccesses,omitempty"`

	RetrieveResult *RetrieveResult `json:"retrieveResult,omitempty"`

	RunTestResult *RunTestsResult `json:"runTestResult,omitempty"`
}

type DeployMessage struct {
	//XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DeployMessage"`

	Changed bool `json:"changed,omitempty"`

	ColumnNumber int32 `json:"columnNumber,omitempty"`

	ComponentType string `json:"componentType,omitempty"`

	Created bool `json:"created,omitempty"`

	CreatedDate time.Time `json:"createdDate,omitempty"`

	Deleted bool `json:"deleted,omitempty"`

	FileName string `json:"fileName,omitempty"`

	FullName string `json:"fullName,omitempty"`

	Id string `json:"id,omitempty"`

	LineNumber int32 `json:"lineNumber,omitempty"`

	Problem string `json:"problem,omitempty"`

	ProblemType *DeployProblemType `json:"problemType,omitempty"`

	Success bool `json:"success,omitempty"`
}

type RetrieveResult struct {
	//XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RetrieveResult"`

	Done bool `json:"done,omitempty"`

	ErrorMessage string `json:"errorMessage,omitempty"`

	ErrorStatusCode StatusCode `json:"errorStatusCode,omitempty"`

	FileProperties []*FileProperties `json:"fileProperties,omitempty"`

	Id string `json:"id,omitempty"`

	Messages []*RetrieveMessage `json:"messages,omitempty"`

	Status *RetrieveStatus `json:"status,omitempty"`

	Success bool `json:"success,omitempty"`

	ZipFile []byte `json:"zipFile,omitempty"`
}

type FileProperties struct {
	CreatedById string `json:"createdById,omitempty"`

	CreatedByName string `json:"createdByName,omitempty"`

	CreatedDate time.Time `json:"createdDate,omitempty"`

	FileName string `json:"fileName,omitempty"`

	FullName string `json:"fullName,omitempty"`

	Id string `json:"id,omitempty"`

	LastModifiedById string `json:"lastModifiedById,omitempty"`

	LastModifiedByName string `json:"lastModifiedByName,omitempty"`

	LastModifiedDate time.Time `json:"lastModifiedDate,omitempty"`

	ManageableState ManageableState `json:"manageableState,omitempty"`

	NamespacePrefix string `json:"namespacePrefix,omitempty"`

	Type_ string `json:"type,omitempty"`
}

type RetrieveMessage struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RetrieveMessage"`

	FileName string `json:"fileName,omitempty"`

	Problem string `json:"problem,omitempty"`
}

type RunTestsResult struct {
	//XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RunTestsResult"`

	ApexLogId string `json:"apexLogId,omitempty"`

	CodeCoverage []*CodeCoverageResult `json:"codeCoverage,omitempty"`

	CodeCoverageWarnings []*CodeCoverageWarning `json:"codeCoverageWarnings,omitempty"`

	Failures []*RunTestFailure `json:"failures,omitempty"`

	NumFailures int32 `json:"numFailures,omitempty"`

	NumTestsRun int32 `json:"numTestsRun,omitempty"`

	Successes []*RunTestSuccess `json:"successes,omitempty"`

	TotalTime float64 `json:"totalTime,omitempty"`
}

type CodeCoverageResult struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CodeCoverageResult"`

	DmlInfo []*CodeLocation `json:"dmlInfo,omitempty"`

	Id ID `json:"id,omitempty"`

	LocationsNotCovered []*CodeLocation `json:"locationsNotCovered,omitempty"`

	MethodInfo []*CodeLocation `json:"methodInfo,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	NumLocations int32 `json:"numLocations,omitempty"`

	NumLocationsNotCovered int32 `json:"numLocationsNotCovered,omitempty"`

	SoqlInfo []*CodeLocation `json:"soqlInfo,omitempty"`

	SoslInfo []*CodeLocation `json:"soslInfo,omitempty"`

	Type_ string `json:"type,omitempty"`
}

type CodeLocation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CodeLocation"`

	Column int32 `json:"column,omitempty"`

	Line int32 `json:"line,omitempty"`

	NumExecutions int32 `json:"numExecutions,omitempty"`

	Time float64 `json:"time,omitempty"`
}

type CodeCoverageWarning struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CodeCoverageWarning"`

	Id ID `json:"id,omitempty"`

	Message string `json:"message,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type RunTestFailure struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RunTestFailure"`

	Id ID `json:"id,omitempty"`

	Message string `json:"message,omitempty"`

	MethodName string `json:"methodName,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	PackageName string `json:"packageName,omitempty"`

	SeeAllData bool `json:"seeAllData,omitempty"`

	StackTrace string `json:"stackTrace,omitempty"`

	Time float64 `json:"time,omitempty"`

	Type_ string `json:"type,omitempty"`
}

type RunTestSuccess struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RunTestSuccess"`

	Id ID `json:"id,omitempty"`

	MethodName string `json:"methodName,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace,omitempty"`

	SeeAllData bool `json:"seeAllData,omitempty"`

	Time float64 `json:"time,omitempty"`
}

type MetadataInterface interface {}

type Metadata struct {
	FullName string `json:"fullName,omitempty"`
}

type AccountSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AccountSettings"`

	*Metadata

	EnableAccountOwnerReport bool `json:"enableAccountOwnerReport,omitempty"`

	EnableAccountTeams bool `json:"enableAccountTeams,omitempty"`

	ShowViewHierarchyLink bool `json:"showViewHierarchyLink,omitempty"`
}

type ActionLinkGroupTemplate struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ActionLinkGroupTemplate"`

	*Metadata

	ActionLinkTemplates []*ActionLinkTemplate `json:"actionLinkTemplates,omitempty"`

	Category *PlatformActionGroupCategory `json:"category,omitempty"`

	ExecutionsAllowed *ActionLinkExecutionsAllowed `json:"executionsAllowed,omitempty"`

	HoursUntilExpiration int32 `json:"hoursUntilExpiration,omitempty"`

	IsPublished bool `json:"isPublished,omitempty"`

	Name string `json:"name,omitempty"`
}

type ActionLinkTemplate struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ActionLinkTemplate"`

	ActionUrl string `json:"actionUrl,omitempty"`

	Headers string `json:"headers,omitempty"`

	IsConfirmationRequired bool `json:"isConfirmationRequired,omitempty"`

	IsGroupDefault bool `json:"isGroupDefault,omitempty"`

	Label string `json:"label,omitempty"`

	LabelKey string `json:"labelKey,omitempty"`

	LinkType *ActionLinkType `json:"linkType,omitempty"`

	Method *ActionLinkHttpMethod `json:"method,omitempty"`

	Position int32 `json:"position,omitempty"`

	RequestBody string `json:"requestBody,omitempty"`

	UserAlias string `json:"userAlias,omitempty"`

	UserVisibility *ActionLinkUserVisibility `json:"userVisibility,omitempty"`
}

type ActivitiesSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ActivitiesSettings"`

	*Metadata

	AllowUsersToRelateMultipleContactsToTasksAndEvents bool `json:"allowUsersToRelateMultipleContactsToTasksAndEvents,omitempty"`

	EnableActivityReminders bool `json:"enableActivityReminders,omitempty"`

	EnableClickCreateEvents bool `json:"enableClickCreateEvents,omitempty"`

	EnableDragAndDropScheduling bool `json:"enableDragAndDropScheduling,omitempty"`

	EnableEmailTracking bool `json:"enableEmailTracking,omitempty"`

	EnableGroupTasks bool `json:"enableGroupTasks,omitempty"`

	EnableListViewScheduling bool `json:"enableListViewScheduling,omitempty"`

	EnableLogNote bool `json:"enableLogNote,omitempty"`

	EnableMultidayEvents bool `json:"enableMultidayEvents,omitempty"`

	EnableRecurringEvents bool `json:"enableRecurringEvents,omitempty"`

	EnableRecurringTasks bool `json:"enableRecurringTasks,omitempty"`

	EnableSidebarCalendarShortcut bool `json:"enableSidebarCalendarShortcut,omitempty"`

	EnableSimpleTaskCreateUI bool `json:"enableSimpleTaskCreateUI,omitempty"`

	EnableUNSTaskDelegatedToNotifications bool `json:"enableUNSTaskDelegatedToNotifications,omitempty"`

	MeetingRequestsLogo string `json:"meetingRequestsLogo,omitempty"`

	ShowCustomLogoMeetingRequests bool `json:"showCustomLogoMeetingRequests,omitempty"`

	ShowEventDetailsMultiUserCalendar bool `json:"showEventDetailsMultiUserCalendar,omitempty"`

	ShowHomePageHoverLinksForEvents bool `json:"showHomePageHoverLinksForEvents,omitempty"`

	ShowMyTasksHoverLinks bool `json:"showMyTasksHoverLinks,omitempty"`

	ShowRequestedMeetingsOnHomePage bool `json:"showRequestedMeetingsOnHomePage,omitempty"`
}

type AddressSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AddressSettings"`

	*Metadata

	CountriesAndStates *CountriesAndStates `json:"countriesAndStates,omitempty"`
}

type CountriesAndStates struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CountriesAndStates"`

	Countries []*Country `json:"countries,omitempty"`
}

type Country struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Country"`

	Active bool `json:"active,omitempty"`

	IntegrationValue string `json:"integrationValue,omitempty"`

	IsoCode string `json:"isoCode,omitempty"`

	Label string `json:"label,omitempty"`

	OrgDefault bool `json:"orgDefault,omitempty"`

	Standard bool `json:"standard,omitempty"`

	States []*State `json:"states,omitempty"`

	Visible bool `json:"visible,omitempty"`
}

type State struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata State"`

	Active bool `json:"active,omitempty"`

	IntegrationValue string `json:"integrationValue,omitempty"`

	IsoCode string `json:"isoCode,omitempty"`

	Label string `json:"label,omitempty"`

	Standard bool `json:"standard,omitempty"`

	Visible bool `json:"visible,omitempty"`
}

type AnalyticSnapshot struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AnalyticSnapshot"`

	*Metadata

	Description string `json:"description,omitempty"`

	GroupColumn string `json:"groupColumn,omitempty"`

	Mappings []*AnalyticSnapshotMapping `json:"mappings,omitempty"`

	Name string `json:"name,omitempty"`

	RunningUser string `json:"runningUser,omitempty"`

	SourceReport string `json:"sourceReport,omitempty"`

	TargetObject string `json:"targetObject,omitempty"`
}

type AnalyticSnapshotMapping struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AnalyticSnapshotMapping"`

	AggregateType *ReportSummaryType `json:"aggregateType,omitempty"`

	SourceField string `json:"sourceField,omitempty"`

	SourceType *ReportJobSourceTypes `json:"sourceType,omitempty"`

	TargetField string `json:"targetField,omitempty"`
}

type ApexTestSuite struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ApexTestSuite"`

	*Metadata

	TestClassName []string `json:"testClassName,omitempty"`
}

type AppMenu struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AppMenu"`

	*Metadata

	AppMenuItems []*AppMenuItem `json:"appMenuItems,omitempty"`
}

type AppMenuItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AppMenuItem"`

	Name string `json:"name,omitempty"`

	Type_ string `json:"type,omitempty"`
}

type ApprovalProcess struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ApprovalProcess"`

	*Metadata

	Active bool `json:"active,omitempty"`

	AllowRecall bool `json:"allowRecall,omitempty"`

	AllowedSubmitters []*ApprovalSubmitter `json:"allowedSubmitters,omitempty"`

	ApprovalPageFields *ApprovalPageField `json:"approvalPageFields,omitempty"`

	ApprovalStep []*ApprovalStep `json:"approvalStep,omitempty"`

	Description string `json:"description,omitempty"`

	EmailTemplate string `json:"emailTemplate,omitempty"`

	EnableMobileDeviceAccess bool `json:"enableMobileDeviceAccess,omitempty"`

	EntryCriteria *ApprovalEntryCriteria `json:"entryCriteria,omitempty"`

	FinalApprovalActions *ApprovalAction `json:"finalApprovalActions,omitempty"`

	FinalApprovalRecordLock bool `json:"finalApprovalRecordLock,omitempty"`

	FinalRejectionActions *ApprovalAction `json:"finalRejectionActions,omitempty"`

	FinalRejectionRecordLock bool `json:"finalRejectionRecordLock,omitempty"`

	InitialSubmissionActions *ApprovalAction `json:"initialSubmissionActions,omitempty"`

	Label string `json:"label,omitempty"`

	NextAutomatedApprover *NextAutomatedApprover `json:"nextAutomatedApprover,omitempty"`

	PostTemplate string `json:"postTemplate,omitempty"`

	RecallActions *ApprovalAction `json:"recallActions,omitempty"`

	RecordEditability *RecordEditabilityType `json:"recordEditability,omitempty"`

	ShowApprovalHistory bool `json:"showApprovalHistory,omitempty"`
}

type ApprovalSubmitter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ApprovalSubmitter"`

	Submitter string `json:"submitter,omitempty"`

	Type_ *ProcessSubmitterType `json:"type,omitempty"`
}

type ApprovalPageField struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ApprovalPageField"`

	Field []string `json:"field,omitempty"`
}

type ApprovalStep struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ApprovalStep"`

	AllowDelegate bool `json:"allowDelegate,omitempty"`

	ApprovalActions *ApprovalAction `json:"approvalActions,omitempty"`

	AssignedApprover *ApprovalStepApprover `json:"assignedApprover,omitempty"`

	Description string `json:"description,omitempty"`

	EntryCriteria *ApprovalEntryCriteria `json:"entryCriteria,omitempty"`

	IfCriteriaNotMet *StepCriteriaNotMetType `json:"ifCriteriaNotMet,omitempty"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`

	RejectBehavior *ApprovalStepRejectBehavior `json:"rejectBehavior,omitempty"`

	RejectionActions *ApprovalAction `json:"rejectionActions,omitempty"`
}

type ApprovalAction struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ApprovalAction"`

	Action []*WorkflowActionReference `json:"action,omitempty"`
}

type WorkflowActionReference struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowActionReference"`

	Name string `json:"name,omitempty"`

	Type_ *WorkflowActionType `json:"type,omitempty"`
}

type ApprovalStepApprover struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ApprovalStepApprover"`

	Approver []*Approver `json:"approver,omitempty"`

	WhenMultipleApprovers *RoutingType `json:"whenMultipleApprovers,omitempty"`
}

type Approver struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Approver"`

	Name string `json:"name,omitempty"`

	Type_ *NextOwnerType `json:"type,omitempty"`
}

type ApprovalEntryCriteria struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ApprovalEntryCriteria"`

	BooleanFilter string `json:"booleanFilter,omitempty"`

	CriteriaItems []*FilterItem `json:"criteriaItems,omitempty"`

	Formula string `json:"formula,omitempty"`
}

type FilterItem struct {
	Field string `json:"field,omitempty"`

	Operation *FilterOperation `json:"operation,omitempty"`

	Value string `json:"value,omitempty"`

	ValueField string `json:"valueField,omitempty"`
}

type DuplicateRuleFilterItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DuplicateRuleFilterItem"`

	*FilterItem

	SortOrder int32 `json:"sortOrder,omitempty"`

	Table string `json:"table,omitempty"`
}

type ApprovalStepRejectBehavior struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ApprovalStepRejectBehavior"`

	Type_ *StepRejectBehaviorType `json:"type,omitempty"`
}

type NextAutomatedApprover struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata NextAutomatedApprover"`

	UseApproverFieldOfRecordOwner bool `json:"useApproverFieldOfRecordOwner,omitempty"`

	UserHierarchyField string `json:"userHierarchyField,omitempty"`
}

type AssignmentRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AssignmentRule"`

	*Metadata

	Active bool `json:"active,omitempty"`

	RuleEntry []*RuleEntry `json:"ruleEntry,omitempty"`
}

type RuleEntry struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RuleEntry"`

	AssignedTo string `json:"assignedTo,omitempty"`

	AssignedToType *AssignToLookupValueType `json:"assignedToType,omitempty"`

	BooleanFilter string `json:"booleanFilter,omitempty"`

	BusinessHours string `json:"businessHours,omitempty"`

	BusinessHoursSource *BusinessHoursSourceType `json:"businessHoursSource,omitempty"`

	CriteriaItems []*FilterItem `json:"criteriaItems,omitempty"`

	DisableEscalationWhenModified bool `json:"disableEscalationWhenModified,omitempty"`

	EscalationAction []*EscalationAction `json:"escalationAction,omitempty"`

	EscalationStartTime *EscalationStartTimeType `json:"escalationStartTime,omitempty"`

	Formula string `json:"formula,omitempty"`

	NotifyCcRecipients bool `json:"notifyCcRecipients,omitempty"`

	OverrideExistingTeams bool `json:"overrideExistingTeams,omitempty"`

	ReplyToEmail string `json:"replyToEmail,omitempty"`

	SenderEmail string `json:"senderEmail,omitempty"`

	SenderName string `json:"senderName,omitempty"`

	Team []string `json:"team,omitempty"`

	Template string `json:"template,omitempty"`
}

type EscalationAction struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EscalationAction"`

	AssignedTo string `json:"assignedTo,omitempty"`

	AssignedToTemplate string `json:"assignedToTemplate,omitempty"`

	AssignedToType *AssignToLookupValueType `json:"assignedToType,omitempty"`

	MinutesToEscalation int32 `json:"minutesToEscalation,omitempty"`

	NotifyCaseOwner bool `json:"notifyCaseOwner,omitempty"`

	NotifyEmail []string `json:"notifyEmail,omitempty"`

	NotifyTo string `json:"notifyTo,omitempty"`

	NotifyToTemplate string `json:"notifyToTemplate,omitempty"`
}

type AssignmentRules struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AssignmentRules"`

	*Metadata

	AssignmentRule []*AssignmentRule `json:"assignmentRule,omitempty"`
}

type AssistantRecommendationType struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AssistantRecommendationType"`

	*Metadata

	Description string `json:"description,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	PlatformActionlist *PlatformActionList `json:"platformActionlist,omitempty"`

	SobjectType string `json:"sobjectType,omitempty"`

	Title string `json:"title,omitempty"`
}

type PlatformActionList struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PlatformActionList"`

	*Metadata

	ActionListContext *PlatformActionListContext `json:"actionListContext,omitempty"`

	PlatformActionListItems []*PlatformActionListItem `json:"platformActionListItems,omitempty"`

	RelatedSourceEntity string `json:"relatedSourceEntity,omitempty"`
}

type PlatformActionListItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PlatformActionListItem"`

	ActionName string `json:"actionName,omitempty"`

	ActionType *PlatformActionType `json:"actionType,omitempty"`

	SortOrder int32 `json:"sortOrder,omitempty"`

	Subtype string `json:"subtype,omitempty"`
}

type AuraDefinitionBundle struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AuraDefinitionBundle"`

	*Metadata

	SVGContent []byte `json:"SVGContent,omitempty"`

	ApiVersion float64 `json:"apiVersion,omitempty"`

	ControllerContent []byte `json:"controllerContent,omitempty"`

	Description string `json:"description,omitempty"`

	DesignContent []byte `json:"designContent,omitempty"`

	DocumentationContent []byte `json:"documentationContent,omitempty"`

	HelperContent []byte `json:"helperContent,omitempty"`

	Markup []byte `json:"markup,omitempty"`

	ModelContent []byte `json:"modelContent,omitempty"`

	PackageVersions []*PackageVersion `json:"packageVersions,omitempty"`

	RendererContent []byte `json:"rendererContent,omitempty"`

	StyleContent []byte `json:"styleContent,omitempty"`

	TestsuiteContent []byte `json:"testsuiteContent,omitempty"`

	Type_ *AuraBundleType `json:"type,omitempty"`
}

type PackageVersion struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PackageVersion"`

	MajorNumber int32 `json:"majorNumber,omitempty"`

	MinorNumber int32 `json:"minorNumber,omitempty"`

	Namespace string `json:"namespace,omitempty"`
}

type AuthProvider struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AuthProvider"`

	*Metadata

	AuthorizeUrl string `json:"authorizeUrl,omitempty"`

	ConsumerKey string `json:"consumerKey,omitempty"`

	ConsumerSecret string `json:"consumerSecret,omitempty"`

	CustomMetadataTypeRecord string `json:"customMetadataTypeRecord,omitempty"`

	DefaultScopes string `json:"defaultScopes,omitempty"`

	ErrorUrl string `json:"errorUrl,omitempty"`

	ExecutionUser string `json:"executionUser,omitempty"`

	FriendlyName string `json:"friendlyName,omitempty"`

	IconUrl string `json:"iconUrl,omitempty"`

	IdTokenIssuer string `json:"idTokenIssuer,omitempty"`

	IncludeOrgIdInIdentifier bool `json:"includeOrgIdInIdentifier,omitempty"`

	LogoutUrl string `json:"logoutUrl,omitempty"`

	Plugin string `json:"plugin,omitempty"`

	Portal string `json:"portal,omitempty"`

	ProviderType *AuthProviderType `json:"providerType,omitempty"`

	RegistrationHandler string `json:"registrationHandler,omitempty"`

	SendAccessTokenInHeader bool `json:"sendAccessTokenInHeader,omitempty"`

	SendClientCredentialsInHeader bool `json:"sendClientCredentialsInHeader,omitempty"`

	TokenUrl string `json:"tokenUrl,omitempty"`

	UserInfoUrl string `json:"userInfoUrl,omitempty"`
}

type AutoResponseRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AutoResponseRule"`

	*Metadata

	Active bool `json:"active,omitempty"`

	RuleEntry []*RuleEntry `json:"ruleEntry,omitempty"`
}

type AutoResponseRules struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AutoResponseRules"`

	*Metadata

	AutoResponseRule []*AutoResponseRule `json:"autoResponseRule,omitempty"`
}

type BusinessHoursEntry struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata BusinessHoursEntry"`

	*Metadata

	Active bool `json:"active,omitempty"`

	Default_ bool `json:"default,omitempty"`

	FridayEndTime time.Time `json:"fridayEndTime,omitempty"`

	FridayStartTime time.Time `json:"fridayStartTime,omitempty"`

	MondayEndTime time.Time `json:"mondayEndTime,omitempty"`

	MondayStartTime time.Time `json:"mondayStartTime,omitempty"`

	Name string `json:"name,omitempty"`

	SaturdayEndTime time.Time `json:"saturdayEndTime,omitempty"`

	SaturdayStartTime time.Time `json:"saturdayStartTime,omitempty"`

	SundayEndTime time.Time `json:"sundayEndTime,omitempty"`

	SundayStartTime time.Time `json:"sundayStartTime,omitempty"`

	ThursdayEndTime time.Time `json:"thursdayEndTime,omitempty"`

	ThursdayStartTime time.Time `json:"thursdayStartTime,omitempty"`

	TimeZoneId string `json:"timeZoneId,omitempty"`

	TuesdayEndTime time.Time `json:"tuesdayEndTime,omitempty"`

	TuesdayStartTime time.Time `json:"tuesdayStartTime,omitempty"`

	WednesdayEndTime time.Time `json:"wednesdayEndTime,omitempty"`

	WednesdayStartTime time.Time `json:"wednesdayStartTime,omitempty"`
}

type BusinessHoursSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata BusinessHoursSettings"`

	*Metadata

	BusinessHours []*BusinessHoursEntry `json:"businessHours,omitempty"`

	Holidays []*Holiday `json:"holidays,omitempty"`
}

type Holiday struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Holiday"`

	ActivityDate time.Time `json:"activityDate,omitempty"`

	BusinessHours []string `json:"businessHours,omitempty"`

	Description string `json:"description,omitempty"`

	EndTime time.Time `json:"endTime,omitempty"`

	IsRecurring bool `json:"isRecurring,omitempty"`

	Name string `json:"name,omitempty"`

	RecurrenceDayOfMonth int32 `json:"recurrenceDayOfMonth,omitempty"`

	RecurrenceDayOfWeek []string `json:"recurrenceDayOfWeek,omitempty"`

	RecurrenceDayOfWeekMask int32 `json:"recurrenceDayOfWeekMask,omitempty"`

	RecurrenceEndDate time.Time `json:"recurrenceEndDate,omitempty"`

	RecurrenceInstance string `json:"recurrenceInstance,omitempty"`

	RecurrenceInterval int32 `json:"recurrenceInterval,omitempty"`

	RecurrenceMonthOfYear string `json:"recurrenceMonthOfYear,omitempty"`

	RecurrenceStartDate time.Time `json:"recurrenceStartDate,omitempty"`

	RecurrenceType string `json:"recurrenceType,omitempty"`

	StartTime time.Time `json:"startTime,omitempty"`
}

type BusinessProcess struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata BusinessProcess"`

	*Metadata

	Description string `json:"description,omitempty"`

	IsActive bool `json:"isActive,omitempty"`

	Values []*PicklistValue `json:"values,omitempty"`
}

type PicklistValue struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PicklistValue"`

	*GlobalPicklistValue

	AllowEmail bool `json:"allowEmail,omitempty"`

	Closed bool `json:"closed,omitempty"`

	ControllingFieldValues []string `json:"controllingFieldValues,omitempty"`

	Converted bool `json:"converted,omitempty"`

	CssExposed bool `json:"cssExposed,omitempty"`

	ForecastCategory *ForecastCategories `json:"forecastCategory,omitempty"`

	HighPriority bool `json:"highPriority,omitempty"`

	Probability int32 `json:"probability,omitempty"`

	ReverseRole string `json:"reverseRole,omitempty"`

	Reviewed bool `json:"reviewed,omitempty"`

	Won bool `json:"won,omitempty"`
}

type GlobalPicklistValue struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata GlobalPicklistValue"`

	*Metadata

	Color string `json:"color,omitempty"`

	Default_ bool `json:"default,omitempty"`

	Description string `json:"description,omitempty"`

	IsActive bool `json:"isActive,omitempty"`
}

type CallCenter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CallCenter"`

	*Metadata

	AdapterUrl string `json:"adapterUrl,omitempty"`

	CustomSettings string `json:"customSettings,omitempty"`

	DisplayName string `json:"displayName,omitempty"`

	DisplayNameLabel string `json:"displayNameLabel,omitempty"`

	InternalNameLabel string `json:"internalNameLabel,omitempty"`

	Sections []*CallCenterSection `json:"sections,omitempty"`

	Version string `json:"version,omitempty"`
}

type CallCenterSection struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CallCenterSection"`

	Items []*CallCenterItem `json:"items,omitempty"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type CallCenterItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CallCenterItem"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`

	Value string `json:"value,omitempty"`
}

type CampaignInfluenceModel struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CampaignInfluenceModel"`

	*Metadata

	IsDefaultModel bool `json:"isDefaultModel,omitempty"`

	IsModelLocked bool `json:"isModelLocked,omitempty"`

	ModelDescription string `json:"modelDescription,omitempty"`

	Name string `json:"name,omitempty"`
}

type CaseSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CaseSettings"`

	*Metadata

	CaseAssignNotificationTemplate string `json:"caseAssignNotificationTemplate,omitempty"`

	CaseCloseNotificationTemplate string `json:"caseCloseNotificationTemplate,omitempty"`

	CaseCommentNotificationTemplate string `json:"caseCommentNotificationTemplate,omitempty"`

	CaseCreateNotificationTemplate string `json:"caseCreateNotificationTemplate,omitempty"`

	CaseFeedItemSettings []*FeedItemSettings `json:"caseFeedItemSettings,omitempty"`

	CloseCaseThroughStatusChange bool `json:"closeCaseThroughStatusChange,omitempty"`

	DefaultCaseOwner string `json:"defaultCaseOwner,omitempty"`

	DefaultCaseOwnerType string `json:"defaultCaseOwnerType,omitempty"`

	DefaultCaseUser string `json:"defaultCaseUser,omitempty"`

	EmailActionDefaultsHandlerClass string `json:"emailActionDefaultsHandlerClass,omitempty"`

	EmailToCase *EmailToCaseSettings `json:"emailToCase,omitempty"`

	EnableCaseFeed bool `json:"enableCaseFeed,omitempty"`

	EnableDraftEmails bool `json:"enableDraftEmails,omitempty"`

	EnableEarlyEscalationRuleTriggers bool `json:"enableEarlyEscalationRuleTriggers,omitempty"`

	EnableEmailActionDefaultsHandler bool `json:"enableEmailActionDefaultsHandler,omitempty"`

	EnableSuggestedArticlesApplication bool `json:"enableSuggestedArticlesApplication,omitempty"`

	EnableSuggestedArticlesCustomerPortal bool `json:"enableSuggestedArticlesCustomerPortal,omitempty"`

	EnableSuggestedArticlesPartnerPortal bool `json:"enableSuggestedArticlesPartnerPortal,omitempty"`

	EnableSuggestedSolutions bool `json:"enableSuggestedSolutions,omitempty"`

	KeepRecordTypeOnAssignmentRule bool `json:"keepRecordTypeOnAssignmentRule,omitempty"`

	NotifyContactOnCaseComment bool `json:"notifyContactOnCaseComment,omitempty"`

	NotifyDefaultCaseOwner bool `json:"notifyDefaultCaseOwner,omitempty"`

	NotifyOwnerOnCaseComment bool `json:"notifyOwnerOnCaseComment,omitempty"`

	NotifyOwnerOnCaseOwnerChange bool `json:"notifyOwnerOnCaseOwnerChange,omitempty"`

	ShowFewerCloseActions bool `json:"showFewerCloseActions,omitempty"`

	UseSystemEmailAddress bool `json:"useSystemEmailAddress,omitempty"`

	WebToCase *WebToCaseSettings `json:"webToCase,omitempty"`
}

type FeedItemSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FeedItemSettings"`

	CharacterLimit int32 `json:"characterLimit,omitempty"`

	CollapseThread bool `json:"collapseThread,omitempty"`

	DisplayFormat *FeedItemDisplayFormat `json:"displayFormat,omitempty"`

	FeedItemType *FeedItemType `json:"feedItemType,omitempty"`
}

type EmailToCaseSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EmailToCaseSettings"`

	EnableEmailToCase bool `json:"enableEmailToCase,omitempty"`

	EnableHtmlEmail bool `json:"enableHtmlEmail,omitempty"`

	EnableOnDemandEmailToCase bool `json:"enableOnDemandEmailToCase,omitempty"`

	EnableThreadIDInBody bool `json:"enableThreadIDInBody,omitempty"`

	EnableThreadIDInSubject bool `json:"enableThreadIDInSubject,omitempty"`

	NotifyOwnerOnNewCaseEmail bool `json:"notifyOwnerOnNewCaseEmail,omitempty"`

	OverEmailLimitAction *EmailToCaseOnFailureActionType `json:"overEmailLimitAction,omitempty"`

	PreQuoteSignature bool `json:"preQuoteSignature,omitempty"`

	RoutingAddresses []*EmailToCaseRoutingAddress `json:"routingAddresses,omitempty"`

	UnauthorizedSenderAction *EmailToCaseOnFailureActionType `json:"unauthorizedSenderAction,omitempty"`
}

type EmailToCaseRoutingAddress struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EmailToCaseRoutingAddress"`

	AddressType *EmailToCaseRoutingAddressType `json:"addressType,omitempty"`

	AuthorizedSenders string `json:"authorizedSenders,omitempty"`

	CaseOrigin string `json:"caseOrigin,omitempty"`

	CaseOwner string `json:"caseOwner,omitempty"`

	CaseOwnerType string `json:"caseOwnerType,omitempty"`

	CasePriority string `json:"casePriority,omitempty"`

	CreateTask bool `json:"createTask,omitempty"`

	EmailAddress string `json:"emailAddress,omitempty"`

	EmailServicesAddress string `json:"emailServicesAddress,omitempty"`

	IsVerified bool `json:"isVerified,omitempty"`

	RoutingName string `json:"routingName,omitempty"`

	SaveEmailHeaders bool `json:"saveEmailHeaders,omitempty"`

	TaskStatus string `json:"taskStatus,omitempty"`
}

type WebToCaseSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WebToCaseSettings"`

	CaseOrigin string `json:"caseOrigin,omitempty"`

	DefaultResponseTemplate string `json:"defaultResponseTemplate,omitempty"`

	EnableWebToCase bool `json:"enableWebToCase,omitempty"`
}

type ChannelLayout struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ChannelLayout"`

	*Metadata

	EnabledChannels []string `json:"enabledChannels,omitempty"`

	Label string `json:"label,omitempty"`

	LayoutItems []*ChannelLayoutItem `json:"layoutItems,omitempty"`
}

type ChannelLayoutItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ChannelLayoutItem"`

	Field string `json:"field,omitempty"`
}

type ChatterAnswersSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ChatterAnswersSettings"`

	*Metadata

	EmailFollowersOnBestAnswer bool `json:"emailFollowersOnBestAnswer,omitempty"`

	EmailFollowersOnReply bool `json:"emailFollowersOnReply,omitempty"`

	EmailOwnerOnPrivateReply bool `json:"emailOwnerOnPrivateReply,omitempty"`

	EmailOwnerOnReply bool `json:"emailOwnerOnReply,omitempty"`

	EnableAnswerViaEmail bool `json:"enableAnswerViaEmail,omitempty"`

	EnableChatterAnswers bool `json:"enableChatterAnswers,omitempty"`

	EnableFacebookSSO bool `json:"enableFacebookSSO,omitempty"`

	EnableInlinePublisher bool `json:"enableInlinePublisher,omitempty"`

	EnableReputation bool `json:"enableReputation,omitempty"`

	EnableRichTextEditor bool `json:"enableRichTextEditor,omitempty"`

	FacebookAuthProvider string `json:"facebookAuthProvider,omitempty"`

	ShowInPortals bool `json:"showInPortals,omitempty"`
}

type CleanDataService struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CleanDataService"`

	*Metadata

	CleanRules []*CleanRule `json:"cleanRules,omitempty"`

	Description string `json:"description,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	MatchEngine string `json:"matchEngine,omitempty"`
}

type CleanRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CleanRule"`

	BulkEnabled bool `json:"bulkEnabled,omitempty"`

	BypassTriggers bool `json:"bypassTriggers,omitempty"`

	BypassWorkflow bool `json:"bypassWorkflow,omitempty"`

	Description string `json:"description,omitempty"`

	DeveloperName string `json:"developerName,omitempty"`

	FieldMappings []*FieldMapping `json:"fieldMappings,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	MatchRule string `json:"matchRule,omitempty"`

	SourceSobjectType string `json:"sourceSobjectType,omitempty"`

	Status *CleanRuleStatus `json:"status,omitempty"`

	TargetSobjectType string `json:"targetSobjectType,omitempty"`
}

type FieldMapping struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FieldMapping"`

	SObjectType string `json:"SObjectType,omitempty"`

	DeveloperName string `json:"developerName,omitempty"`

	FieldMappingRows []*FieldMappingRow `json:"fieldMappingRows,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`
}

type FieldMappingRow struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FieldMappingRow"`

	SObjectType string `json:"SObjectType,omitempty"`

	FieldMappingFields []*FieldMappingField `json:"fieldMappingFields,omitempty"`

	FieldName string `json:"fieldName,omitempty"`

	MappingOperation *MappingOperation `json:"mappingOperation,omitempty"`
}

type FieldMappingField struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FieldMappingField"`

	DataServiceField string `json:"dataServiceField,omitempty"`

	DataServiceObjectName string `json:"dataServiceObjectName,omitempty"`

	Priority int32 `json:"priority,omitempty"`
}

type Community struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Community"`

	*Metadata

	Active bool `json:"active,omitempty"`

	ChatterAnswersFacebookSsoUrl string `json:"chatterAnswersFacebookSsoUrl,omitempty"`

	CommunityFeedPage string `json:"communityFeedPage,omitempty"`

	DataCategoryName string `json:"dataCategoryName,omitempty"`

	Description string `json:"description,omitempty"`

	EmailFooterDocument string `json:"emailFooterDocument,omitempty"`

	EmailHeaderDocument string `json:"emailHeaderDocument,omitempty"`

	EmailNotificationUrl string `json:"emailNotificationUrl,omitempty"`

	EnableChatterAnswers bool `json:"enableChatterAnswers,omitempty"`

	EnablePrivateQuestions bool `json:"enablePrivateQuestions,omitempty"`

	ExpertsGroup string `json:"expertsGroup,omitempty"`

	Portal string `json:"portal,omitempty"`

	ReputationLevels *ReputationLevels `json:"reputationLevels,omitempty"`

	ShowInPortal bool `json:"showInPortal,omitempty"`

	Site string `json:"site,omitempty"`
}

type ReputationLevels struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReputationLevels"`

	ChatterAnswersReputationLevels []*ChatterAnswersReputationLevel `json:"chatterAnswersReputationLevels,omitempty"`

	IdeaReputationLevels []*IdeaReputationLevel `json:"ideaReputationLevels,omitempty"`
}

type ChatterAnswersReputationLevel struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ChatterAnswersReputationLevel"`

	Name string `json:"name,omitempty"`

	Value int32 `json:"value,omitempty"`
}

type IdeaReputationLevel struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata IdeaReputationLevel"`

	Name string `json:"name,omitempty"`

	Value int32 `json:"value,omitempty"`
}

type CommunityTemplateDefinition struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CommunityTemplateDefinition"`

	*Metadata

	BundlesInfo []*CommunityTemplateBundleInfo `json:"bundlesInfo,omitempty"`

	Category *CommunityTemplateCategory `json:"category,omitempty"`

	DefaultThemeDefinition string `json:"defaultThemeDefinition,omitempty"`

	Description string `json:"description,omitempty"`

	EnableExtendedCleanUpOnDelete bool `json:"enableExtendedCleanUpOnDelete,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	PageSetting []*CommunityTemplatePageSetting `json:"pageSetting,omitempty"`
}

type CommunityTemplateBundleInfo struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CommunityTemplateBundleInfo"`

	Description string `json:"description,omitempty"`

	Image string `json:"image,omitempty"`

	Order int32 `json:"order,omitempty"`

	Title string `json:"title,omitempty"`

	Type_ *CommunityTemplateBundleInfoType `json:"type,omitempty"`
}

type CommunityTemplatePageSetting struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CommunityTemplatePageSetting"`

	Page string `json:"page,omitempty"`
}

type CommunityThemeDefinition struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CommunityThemeDefinition"`

	*Metadata

	Description string `json:"description,omitempty"`

	EnableExtendedCleanUpOnDelete bool `json:"enableExtendedCleanUpOnDelete,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	ThemeSetting []*CommunityThemeSetting `json:"themeSetting,omitempty"`
}

type CommunityThemeSetting struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CommunityThemeSetting"`

	ThemeLayout string `json:"themeLayout,omitempty"`

	ThemeLayoutType *CommunityThemeLayoutType `json:"themeLayoutType,omitempty"`
}

type CompactLayout struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CompactLayout"`

	*Metadata

	Fields []string `json:"fields,omitempty"`

	Label string `json:"label,omitempty"`
}

type CompanySettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CompanySettings"`

	*Metadata

	FiscalYear *FiscalYearSettings `json:"fiscalYear,omitempty"`
}

type FiscalYearSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FiscalYearSettings"`

	FiscalYearNameBasedOn string `json:"fiscalYearNameBasedOn,omitempty"`

	StartMonth string `json:"startMonth,omitempty"`
}

type ConnectedApp struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ConnectedApp"`

	*Metadata

	Attributes []*ConnectedAppAttribute `json:"attributes,omitempty"`

	CanvasConfig *ConnectedAppCanvasConfig `json:"canvasConfig,omitempty"`

	ContactEmail string `json:"contactEmail,omitempty"`

	ContactPhone string `json:"contactPhone,omitempty"`

	Description string `json:"description,omitempty"`

	IconUrl string `json:"iconUrl,omitempty"`

	InfoUrl string `json:"infoUrl,omitempty"`

	IpRanges []*ConnectedAppIpRange `json:"ipRanges,omitempty"`

	Label string `json:"label,omitempty"`

	LogoUrl string `json:"logoUrl,omitempty"`

	MobileAppConfig *ConnectedAppMobileDetailConfig `json:"mobileAppConfig,omitempty"`

	MobileStartUrl string `json:"mobileStartUrl,omitempty"`

	OauthConfig *ConnectedAppOauthConfig `json:"oauthConfig,omitempty"`

	Plugin string `json:"plugin,omitempty"`

	SamlConfig *ConnectedAppSamlConfig `json:"samlConfig,omitempty"`

	StartUrl string `json:"startUrl,omitempty"`
}

type ConnectedAppAttribute struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ConnectedAppAttribute"`

	Formula string `json:"formula,omitempty"`

	Key string `json:"key,omitempty"`
}

type ConnectedAppCanvasConfig struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ConnectedAppCanvasConfig"`

	AccessMethod *AccessMethod `json:"accessMethod,omitempty"`

	CanvasUrl string `json:"canvasUrl,omitempty"`

	LifecycleClass string `json:"lifecycleClass,omitempty"`

	Locations []*CanvasLocationOptions `json:"locations,omitempty"`

	Options []*CanvasOptions `json:"options,omitempty"`

	SamlInitiationMethod *SamlInitiationMethod `json:"samlInitiationMethod,omitempty"`
}

type ConnectedAppIpRange struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ConnectedAppIpRange"`

	Description string `json:"description,omitempty"`

	End string `json:"end,omitempty"`

	Start string `json:"start,omitempty"`
}

type ConnectedAppMobileDetailConfig struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ConnectedAppMobileDetailConfig"`

	ApplicationBinaryFile []byte `json:"applicationBinaryFile,omitempty"`

	ApplicationBinaryFileName string `json:"applicationBinaryFileName,omitempty"`

	ApplicationBundleIdentifier string `json:"applicationBundleIdentifier,omitempty"`

	ApplicationFileLength int32 `json:"applicationFileLength,omitempty"`

	ApplicationIconFile string `json:"applicationIconFile,omitempty"`

	ApplicationIconFileName string `json:"applicationIconFileName,omitempty"`

	ApplicationInstallUrl string `json:"applicationInstallUrl,omitempty"`

	DevicePlatform *DevicePlatformType `json:"devicePlatform,omitempty"`

	DeviceType *DeviceType `json:"deviceType,omitempty"`

	MinimumOsVersion string `json:"minimumOsVersion,omitempty"`

	PrivateApp bool `json:"privateApp,omitempty"`

	Version string `json:"version,omitempty"`
}

type ConnectedAppOauthConfig struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ConnectedAppOauthConfig"`

	CallbackUrl string `json:"callbackUrl,omitempty"`

	Certificate string `json:"certificate,omitempty"`

	ConsumerKey string `json:"consumerKey,omitempty"`

	ConsumerSecret string `json:"consumerSecret,omitempty"`

	Scopes []*ConnectedAppOauthAccessScope `json:"scopes,omitempty"`
}

type ConnectedAppSamlConfig struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ConnectedAppSamlConfig"`

	AcsUrl string `json:"acsUrl,omitempty"`

	Certificate string `json:"certificate,omitempty"`

	EncryptionCertificate string `json:"encryptionCertificate,omitempty"`

	EncryptionType *SamlEncryptionType `json:"encryptionType,omitempty"`

	EntityUrl string `json:"entityUrl,omitempty"`

	Issuer string `json:"issuer,omitempty"`

	SamlNameIdFormat *SamlNameIdFormatType `json:"samlNameIdFormat,omitempty"`

	SamlSubjectCustomAttr string `json:"samlSubjectCustomAttr,omitempty"`

	SamlSubjectType *SamlSubjectType `json:"samlSubjectType,omitempty"`
}

type ContractSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ContractSettings"`

	*Metadata

	AutoCalculateEndDate bool `json:"autoCalculateEndDate,omitempty"`

	AutoExpirationDelay string `json:"autoExpirationDelay,omitempty"`

	AutoExpirationRecipient string `json:"autoExpirationRecipient,omitempty"`

	AutoExpireContracts bool `json:"autoExpireContracts,omitempty"`

	EnableContractHistoryTracking bool `json:"enableContractHistoryTracking,omitempty"`

	NotifyOwnersOnContractExpiration bool `json:"notifyOwnersOnContractExpiration,omitempty"`
}

type CorsWhitelistOrigin struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CorsWhitelistOrigin"`

	*Metadata

	UrlPattern string `json:"urlPattern,omitempty"`
}

type CustomApplication struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomApplication"`

	*Metadata

	ActionOverrides []*AppActionOverride `json:"actionOverrides,omitempty"`

	Brand *AppBrand `json:"brand,omitempty"`

	CustomApplicationComponents *CustomApplicationComponents `json:"customApplicationComponents,omitempty"`

	DefaultLandingTab string `json:"defaultLandingTab,omitempty"`

	Description string `json:"description,omitempty"`

	DetailPageRefreshMethod string `json:"detailPageRefreshMethod,omitempty"`

	DomainWhitelist *DomainWhitelist `json:"domainWhitelist,omitempty"`

	EnableCustomizeMyTabs bool `json:"enableCustomizeMyTabs,omitempty"`

	EnableKeyboardShortcuts bool `json:"enableKeyboardShortcuts,omitempty"`

	EnableListViewHover bool `json:"enableListViewHover,omitempty"`

	EnableListViewReskin bool `json:"enableListViewReskin,omitempty"`

	EnableMultiMonitorComponents bool `json:"enableMultiMonitorComponents,omitempty"`

	EnablePinTabs bool `json:"enablePinTabs,omitempty"`

	EnableTabHover bool `json:"enableTabHover,omitempty"`

	EnableTabLimits bool `json:"enableTabLimits,omitempty"`

	FooterColor string `json:"footerColor,omitempty"`

	FormFactors []*FormFactor `json:"formFactors,omitempty"`

	HeaderColor string `json:"headerColor,omitempty"`

	IsServiceCloudConsole bool `json:"isServiceCloudConsole,omitempty"`

	KeyboardShortcuts *KeyboardShortcuts `json:"keyboardShortcuts,omitempty"`

	Label string `json:"label,omitempty"`

	ListPlacement *ListPlacement `json:"listPlacement,omitempty"`

	ListRefreshMethod string `json:"listRefreshMethod,omitempty"`

	LiveAgentConfig *LiveAgentConfig `json:"liveAgentConfig,omitempty"`

	Logo string `json:"logo,omitempty"`

	NavType *NavType `json:"navType,omitempty"`

	PrimaryTabColor string `json:"primaryTabColor,omitempty"`

	PushNotifications *PushNotifications `json:"pushNotifications,omitempty"`

	SaveUserSessions bool `json:"saveUserSessions,omitempty"`

	Tab []string `json:"tab,omitempty"`

	TabLimitConfig *TabLimitConfig `json:"tabLimitConfig,omitempty"`

	UiType *UiType `json:"uiType,omitempty"`

	UtilityBar string `json:"utilityBar,omitempty"`

	WorkspaceMappings *WorkspaceMappings `json:"workspaceMappings,omitempty"`
}

type AppActionOverride struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AppActionOverride"`

	*ActionOverride

	PageOrSobjectType string `json:"pageOrSobjectType,omitempty"`
}

type ActionOverride struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ActionOverride"`

	ActionName string `json:"actionName,omitempty"`

	Comment string `json:"comment,omitempty"`

	Content string `json:"content,omitempty"`

	FormFactor *FormFactor `json:"formFactor,omitempty"`

	SkipRecordTypeSelect bool `json:"skipRecordTypeSelect,omitempty"`

	Type_ *ActionOverrideType `json:"type,omitempty"`
}

type AppBrand struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AppBrand"`

	FooterColor string `json:"footerColor,omitempty"`

	HeaderColor string `json:"headerColor,omitempty"`

	Logo string `json:"logo,omitempty"`

	LogoVersion int32 `json:"logoVersion,omitempty"`
}

type CustomApplicationComponents struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomApplicationComponents"`

	Alignment string `json:"alignment,omitempty"`

	CustomApplicationComponent []string `json:"customApplicationComponent,omitempty"`
}

type DomainWhitelist struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DomainWhitelist"`

	Domain []string `json:"domain,omitempty"`
}

type KeyboardShortcuts struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata KeyboardShortcuts"`

	CustomShortcut []*CustomShortcut `json:"customShortcut,omitempty"`

	DefaultShortcut []*DefaultShortcut `json:"defaultShortcut,omitempty"`
}

type CustomShortcut struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomShortcut"`

	*DefaultShortcut

	Description string `json:"description,omitempty"`

	EventName string `json:"eventName,omitempty"`
}

type DefaultShortcut struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DefaultShortcut"`

	Action string `json:"action,omitempty"`

	Active bool `json:"active,omitempty"`

	KeyCommand string `json:"keyCommand,omitempty"`
}

type ListPlacement struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ListPlacement"`

	Height int32 `json:"height,omitempty"`

	Location string `json:"location,omitempty"`

	Units string `json:"units,omitempty"`

	Width int32 `json:"width,omitempty"`
}

type LiveAgentConfig struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LiveAgentConfig"`

	EnableLiveChat bool `json:"enableLiveChat,omitempty"`

	OpenNewAccountSubtab bool `json:"openNewAccountSubtab,omitempty"`

	OpenNewCaseSubtab bool `json:"openNewCaseSubtab,omitempty"`

	OpenNewContactSubtab bool `json:"openNewContactSubtab,omitempty"`

	OpenNewLeadSubtab bool `json:"openNewLeadSubtab,omitempty"`

	OpenNewVFPageSubtab bool `json:"openNewVFPageSubtab,omitempty"`

	PagesToOpen *PagesToOpen `json:"pagesToOpen,omitempty"`

	ShowKnowledgeArticles bool `json:"showKnowledgeArticles,omitempty"`
}

type PagesToOpen struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PagesToOpen"`

	PageToOpen []string `json:"pageToOpen,omitempty"`
}

type PushNotifications struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PushNotifications"`

	PushNotification []*PushNotification `json:"pushNotification,omitempty"`
}

type PushNotification struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PushNotification"`

	FieldNames []string `json:"fieldNames,omitempty"`

	ObjectName string `json:"objectName,omitempty"`
}

type TabLimitConfig struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata TabLimitConfig"`

	MaxNumberOfPrimaryTabs string `json:"maxNumberOfPrimaryTabs,omitempty"`

	MaxNumberOfSubTabs string `json:"maxNumberOfSubTabs,omitempty"`
}

type WorkspaceMappings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkspaceMappings"`

	Mapping []*WorkspaceMapping `json:"mapping,omitempty"`
}

type WorkspaceMapping struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkspaceMapping"`

	FieldName string `json:"fieldName,omitempty"`

	Tab string `json:"tab,omitempty"`
}

type CustomApplicationComponent struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomApplicationComponent"`

	*Metadata

	ButtonIconUrl string `json:"buttonIconUrl,omitempty"`

	ButtonStyle string `json:"buttonStyle,omitempty"`

	ButtonText string `json:"buttonText,omitempty"`

	ButtonWidth int32 `json:"buttonWidth,omitempty"`

	Height int32 `json:"height,omitempty"`

	IsHeightFixed bool `json:"isHeightFixed,omitempty"`

	IsHidden bool `json:"isHidden,omitempty"`

	IsWidthFixed bool `json:"isWidthFixed,omitempty"`

	VisualforcePage string `json:"visualforcePage,omitempty"`

	Width int32 `json:"width,omitempty"`
}

type CustomDataType struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomDataType"`

	*Metadata

	CustomDataTypeComponents []*CustomDataTypeComponent `json:"customDataTypeComponents,omitempty"`

	Description string `json:"description,omitempty"`

	DisplayFormula string `json:"displayFormula,omitempty"`

	EditComponentsOnSeparateLines bool `json:"editComponentsOnSeparateLines,omitempty"`

	Label string `json:"label,omitempty"`

	RightAligned bool `json:"rightAligned,omitempty"`

	SupportComponentsInReports bool `json:"supportComponentsInReports,omitempty"`
}

type CustomDataTypeComponent struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomDataTypeComponent"`

	DeveloperSuffix string `json:"developerSuffix,omitempty"`

	EnforceFieldRequiredness bool `json:"enforceFieldRequiredness,omitempty"`

	Label string `json:"label,omitempty"`

	Length int32 `json:"length,omitempty"`

	Precision int32 `json:"precision,omitempty"`

	Scale int32 `json:"scale,omitempty"`

	SortOrder *SortOrder `json:"sortOrder,omitempty"`

	SortPriority int32 `json:"sortPriority,omitempty"`

	Type_ *FieldType `json:"type,omitempty"`
}

type CustomExperience struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomExperience"`

	*Metadata

	AllowInternalUserLogin bool `json:"allowInternalUserLogin,omitempty"`

	Branding *CustomExperienceBranding `json:"branding,omitempty"`

	ChangePasswordEmailTemplate string `json:"changePasswordEmailTemplate,omitempty"`

	EmailFooterLogo string `json:"emailFooterLogo,omitempty"`

	EmailFooterText string `json:"emailFooterText,omitempty"`

	EmailSenderAddress string `json:"emailSenderAddress,omitempty"`

	EmailSenderName string `json:"emailSenderName,omitempty"`

	EnableErrorPageOverridesForVisualforce bool `json:"enableErrorPageOverridesForVisualforce,omitempty"`

	ForgotPasswordEmailTemplate string `json:"forgotPasswordEmailTemplate,omitempty"`

	PicassoSite string `json:"picassoSite,omitempty"`

	SObjectType string `json:"sObjectType,omitempty"`

	SendWelcomeEmail bool `json:"sendWelcomeEmail,omitempty"`

	Site string `json:"site,omitempty"`

	SiteAsContainerEnabled bool `json:"siteAsContainerEnabled,omitempty"`

	Tabs *CustomExperienceTabSet `json:"tabs,omitempty"`

	UrlPathPrefix string `json:"urlPathPrefix,omitempty"`

	WelcomeEmailTemplate string `json:"welcomeEmailTemplate,omitempty"`
}

type CustomExperienceBranding struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomExperienceBranding"`

	LoginFooterText string `json:"loginFooterText,omitempty"`

	LoginLogo string `json:"loginLogo,omitempty"`

	PageFooter string `json:"pageFooter,omitempty"`

	PageHeader string `json:"pageHeader,omitempty"`

	PrimaryColor string `json:"primaryColor,omitempty"`

	PrimaryComplementColor string `json:"primaryComplementColor,omitempty"`

	QuaternaryColor string `json:"quaternaryColor,omitempty"`

	QuaternaryComplementColor string `json:"quaternaryComplementColor,omitempty"`

	SecondaryColor string `json:"secondaryColor,omitempty"`

	TertiaryColor string `json:"tertiaryColor,omitempty"`

	TertiaryComplementColor string `json:"tertiaryComplementColor,omitempty"`

	ZeronaryColor string `json:"zeronaryColor,omitempty"`

	ZeronaryComplementColor string `json:"zeronaryComplementColor,omitempty"`
}

type CustomExperienceTabSet struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomExperienceTabSet"`

	CustomTab []string `json:"customTab,omitempty"`

	DefaultTab string `json:"defaultTab,omitempty"`

	StandardTab []string `json:"standardTab,omitempty"`
}

type CustomFeedFilter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomFeedFilter"`

	*Metadata

	Criteria []*FeedFilterCriterion `json:"criteria,omitempty"`

	Description string `json:"description,omitempty"`

	IsProtected bool `json:"isProtected,omitempty"`

	Label string `json:"label,omitempty"`
}

type FeedFilterCriterion struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FeedFilterCriterion"`

	FeedItemType *FeedItemType `json:"feedItemType,omitempty"`

	FeedItemVisibility *FeedItemVisibility `json:"feedItemVisibility,omitempty"`

	RelatedSObjectType string `json:"relatedSObjectType,omitempty"`
}

type CustomField struct {
	CaseSensitive bool `json:"caseSensitive,omitempty"`

	CustomDataType string `json:"customDataType,omitempty"`

	DefaultValue string `json:"defaultValue,omitempty"`

	DeleteConstraint *DeleteConstraint `json:"deleteConstraint,omitempty"`

	Deprecated bool `json:"deprecated,omitempty"`

	Description string `json:"description,omitempty"`

	DisplayFormat string `json:"displayFormat,omitempty"`

	Encrypted bool `json:"encrypted,omitempty"`

	EscapeMarkup bool `json:"escapeMarkup,omitempty"`

	ExternalDeveloperName string `json:"externalDeveloperName,omitempty"`

	ExternalId bool `json:"externalId,omitempty"`

	FieldManageability *FieldManageability `json:"fieldManageability,omitempty"`

	Formula string `json:"formula,omitempty"`

	FormulaTreatBlanksAs *TreatBlanksAs `json:"formulaTreatBlanksAs,omitempty"`

	FullName string `json:"fullName"`

	InlineHelpText string `json:"inlineHelpText,omitempty"`

	IsConvertLeadDisabled bool `json:"isConvertLeadDisabled,omitempty"`

	IsFilteringDisabled bool `json:"isFilteringDisabled,omitempty"`

	IsNameField bool `json:"isNameField,omitempty"`

	IsSortingDisabled bool `json:"isSortingDisabled,omitempty"`

	Label string `json:"label,omitempty"`

	Length int32 `json:"length,omitempty"`

	LookupFilter *LookupFilter `json:"lookupFilter,omitempty"`

	MaskChar *EncryptedFieldMaskChar `json:"maskChar,omitempty"`

	MaskType *EncryptedFieldMaskType `json:"maskType,omitempty"`

	Picklist *Picklist `json:"picklist,omitempty"`

	PopulateExistingRows bool `json:"populateExistingRows,omitempty"`

	Precision int32 `json:"precision,omitempty"`

	ReferenceTargetField string `json:"referenceTargetField,omitempty"`

	ReferenceTo string `json:"referenceTo,omitempty"`

	RelationshipLabel string `json:"relationshipLabel,omitempty"`

	RelationshipName string `json:"relationshipName,omitempty"`

	RelationshipOrder int32 `json:"relationshipOrder,omitempty"`

	ReparentableMasterDetail bool `json:"reparentableMasterDetail,omitempty"`

	Required bool `json:"required,omitempty"`

	RestrictedAdminField bool `json:"restrictedAdminField,omitempty"`

	Scale int32 `json:"scale,omitempty"`

	StartingNumber int32 `json:"startingNumber,omitempty"`

	StripMarkup bool `json:"stripMarkup,omitempty"`

	SummarizedField string `json:"summarizedField,omitempty"`

	SummaryFilterItems []*FilterItem `json:"summaryFilterItems,omitempty"`

	SummaryForeignKey string `json:"summaryForeignKey,omitempty"`

	SummaryOperation *SummaryOperations `json:"summaryOperation,omitempty"`

	TrackFeedHistory bool `json:"trackFeedHistory,omitempty"`

	TrackHistory bool `json:"trackHistory,omitempty"`

	TrackTrending bool `json:"trackTrending,omitempty"`

	Type FieldType `json:"type,omitempty"`

	Unique bool `json:"unique,omitempty"`

	ValueSet *ValueSet `json:"valueSet,omitempty"`

	VisibleLines int32 `json:"visibleLines,omitempty"`

	WriteRequiresMasterRead bool `json:"writeRequiresMasterRead,omitempty"`
}

type LookupFilter struct {
	Active bool `json:"active,omitempty"`

	BooleanFilter string `json:"booleanFilter,omitempty"`

	Description string `json:"description,omitempty"`

	ErrorMessage string `json:"errorMessage,omitempty"`

	FilterItems []*FilterItem `json:"filterItems,omitempty"`

	InfoMessage string `json:"infoMessage,omitempty"`

	IsOptional bool `json:"isOptional,omitempty"`
}

type Picklist struct {
	ControllingField string `json:"controllingField,omitempty"`

	PicklistValues []*PicklistValue `json:"picklistValues,omitempty"`

	RestrictedPicklist bool `json:"restrictedPicklist,omitempty"`

	Sorted bool `json:"sorted,omitempty"`
}

type ValueSet struct {
	ControllingField string `json:"controllingField,omitempty"`

	Restricted bool `json:"restricted,omitempty"`

	ValueSetDefinition *ValueSetValuesDefinition `json:"valueSetDefinition,omitempty"`

	ValueSetName string `json:"valueSetName,omitempty"`

	ValueSettings []*ValueSettings `json:"valueSettings,omitempty"`
}

type ValueSetValuesDefinition struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ValueSetValuesDefinition"`

	Sorted bool `json:"sorted,omitempty"`

	Value []*CustomValue `json:"value,omitempty"`
}

type CustomValue struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomValue"`

	*Metadata

	Color string `json:"color,omitempty"`

	Default_ bool `json:"default,omitempty"`

	Description string `json:"description,omitempty"`

	IsActive bool `json:"isActive,omitempty"`
}

type StandardValue struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata StandardValue"`

	*CustomValue

	AllowEmail bool `json:"allowEmail,omitempty"`

	Closed bool `json:"closed,omitempty"`

	Converted bool `json:"converted,omitempty"`

	CssExposed bool `json:"cssExposed,omitempty"`

	ForecastCategory *ForecastCategories `json:"forecastCategory,omitempty"`

	HighPriority bool `json:"highPriority,omitempty"`

	Probability int32 `json:"probability,omitempty"`

	ReverseRole string `json:"reverseRole,omitempty"`

	Reviewed bool `json:"reviewed,omitempty"`

	Won bool `json:"won,omitempty"`
}

type ValueSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ValueSettings"`

	ControllingFieldValue []string `json:"controllingFieldValue,omitempty"`

	ValueName string `json:"valueName,omitempty"`
}

type CustomLabel struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomLabel"`

	*Metadata

	Categories string `json:"categories,omitempty"`

	Language string `json:"language,omitempty"`

	Protected bool `json:"protected,omitempty"`

	ShortDescription string `json:"shortDescription,omitempty"`

	Value string `json:"value,omitempty"`
}

type CustomLabels struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomLabels"`

	*Metadata

	Labels []*CustomLabel `json:"labels,omitempty"`
}

type CustomMetadata struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomMetadata"`

	*Metadata

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	Protected bool `json:"protected,omitempty"`

	Values []*CustomMetadataValue `json:"values,omitempty"`
}

type CustomMetadataValue struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomMetadataValue"`

	Field string `json:"field,omitempty"`

	Value interface{} `json:"value,omitempty"`
}

type CustomObject struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata metadata"`

	Type string `json:"http://www.w3.org/2001/XMLSchema-instance type,attr"`

	FullName string `json:"fullName"`

	ActionOverrides []*ActionOverride `json:"actionOverrides,omitempty"`

	AllowInChatterGroups bool `json:"allowInChatterGroups,omitempty"`

	ArticleTypeChannelDisplay *ArticleTypeChannelDisplay `json:"articleTypeChannelDisplay,omitempty"`

	BusinessProcesses []*BusinessProcess `json:"businessProcesses,omitempty"`

	CompactLayoutAssignment string `json:"compactLayoutAssignment,omitempty"`

	CompactLayouts []*CompactLayout `json:"compactLayouts,omitempty"`

	CustomHelp string `json:"customHelp,omitempty"`

	CustomHelpPage string `json:"customHelpPage,omitempty"`

	CustomSettingsType *CustomSettingsType `json:"customSettingsType,omitempty"`

	DeploymentStatus DeploymentStatus `json:"deploymentStatus,omitempty"`

	Deprecated bool `json:"deprecated,omitempty"`

	Description string `json:"description,omitempty"`

	EnableActivities bool `json:"enableActivities,omitempty"`

	EnableBulkApi bool `json:"enableBulkApi,omitempty"`

	EnableDivisions bool `json:"enableDivisions,omitempty"`

	EnableEnhancedLookup bool `json:"enableEnhancedLookup,omitempty"`

	EnableFeeds bool `json:"enableFeeds,omitempty"`

	EnableHistory bool `json:"enableHistory,omitempty"`

	EnableReports bool `json:"enableReports,omitempty"`

	EnableSearch bool `json:"enableSearch,omitempty"`

	EnableSharing bool `json:"enableSharing,omitempty"`

	EnableStreamingApi bool `json:"enableStreamingApi,omitempty"`

	ExternalDataSource string `json:"externalDataSource,omitempty"`

	ExternalName string `json:"externalName,omitempty"`

	ExternalRepository string `json:"externalRepository,omitempty"`

	ExternalSharingModel SharingModel `json:"externalSharingModel,omitempty"`

	FieldSets []*FieldSet `json:"fieldSets,omitempty"`

	Fields []*CustomField `json:"fields,omitempty"`

	Gender *Gender `json:"gender,omitempty"`

	HistoryRetentionPolicy *HistoryRetentionPolicy `json:"historyRetentionPolicy,omitempty"`

	Household bool `json:"household,omitempty"`

	Label string `json:"label,omitempty"`

	ListViews []*ListView `json:"listViews,omitempty"`

	NameField *CustomField `json:"nameField,omitempty"`

	PluralLabel string `json:"pluralLabel,omitempty"`

	RecordTypeTrackFeedHistory bool `json:"recordTypeTrackFeedHistory,omitempty"`

	RecordTypeTrackHistory bool `json:"recordTypeTrackHistory,omitempty"`

	RecordTypes []*RecordType `json:"recordTypes,omitempty"`

	SearchLayouts *SearchLayouts `json:"searchLayouts,omitempty"`

	SharingModel SharingModel `json:"sharingModel,omitempty"`

	SharingReasons []*SharingReason `json:"sharingReasons,omitempty"`

	SharingRecalculations []*SharingRecalculation `json:"sharingRecalculations,omitempty"`

	StartsWith *StartsWith `json:"startsWith,omitempty"`

	ValidationRules []*ValidationRule `json:"validationRules,omitempty"`

	Visibility *SetupObjectVisibility `json:"visibility,omitempty"`

	WebLinks []*WebLink `json:"webLinks,omitempty"`
}

type ArticleTypeChannelDisplay struct {
	ArticleTypeTemplates []*ArticleTypeTemplate `json:"articleTypeTemplates,omitempty"`
}

type ArticleTypeTemplate struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ArticleTypeTemplate"`

	Channel *Channel `json:"channel,omitempty"`

	Page string `json:"page,omitempty"`

	Template *Template `json:"template,omitempty"`
}

type FieldSet struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FieldSet"`

	*Metadata

	AvailableFields []*FieldSetItem `json:"availableFields,omitempty"`

	Description string `json:"description,omitempty"`

	DisplayedFields []*FieldSetItem `json:"displayedFields,omitempty"`

	Label string `json:"label,omitempty"`
}

type FieldSetItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FieldSetItem"`

	Field string `json:"field,omitempty"`

	IsFieldManaged bool `json:"isFieldManaged,omitempty"`

	IsRequired bool `json:"isRequired,omitempty"`
}

type HistoryRetentionPolicy struct {
	ArchiveAfterMonths int32 `json:"archiveAfterMonths,omitempty"`

	ArchiveRetentionYears int32 `json:"archiveRetentionYears,omitempty"`

	Description string `json:"description,omitempty"`
}

type ListView struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ListView"`

	*Metadata

	BooleanFilter string `json:"booleanFilter,omitempty"`

	Columns []string `json:"columns,omitempty"`

	Division string `json:"division,omitempty"`

	FilterScope *FilterScope `json:"filterScope,omitempty"`

	Filters []*ListViewFilter `json:"filters,omitempty"`

	Label string `json:"label,omitempty"`

	Language *Language `json:"language,omitempty"`

	Queue string `json:"queue,omitempty"`

	SharedTo *SharedTo `json:"sharedTo,omitempty"`
}

type ListViewFilter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ListViewFilter"`

	Field string `json:"field,omitempty"`

	Operation *FilterOperation `json:"operation,omitempty"`

	Value string `json:"value,omitempty"`
}

type SharedTo struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SharedTo"`

	AllCustomerPortalUsers string `json:"allCustomerPortalUsers,omitempty"`

	AllInternalUsers string `json:"allInternalUsers,omitempty"`

	AllPartnerUsers string `json:"allPartnerUsers,omitempty"`

	Group []string `json:"group,omitempty"`

	Groups []string `json:"groups,omitempty"`

	ManagerSubordinates []string `json:"managerSubordinates,omitempty"`

	Managers []string `json:"managers,omitempty"`

	PortalRole []string `json:"portalRole,omitempty"`

	PortalRoleAndSubordinates []string `json:"portalRoleAndSubordinates,omitempty"`

	Queue []string `json:"queue,omitempty"`

	Role []string `json:"role,omitempty"`

	RoleAndSubordinates []string `json:"roleAndSubordinates,omitempty"`

	RoleAndSubordinatesInternal []string `json:"roleAndSubordinatesInternal,omitempty"`

	Roles []string `json:"roles,omitempty"`

	RolesAndSubordinates []string `json:"rolesAndSubordinates,omitempty"`

	Territories []string `json:"territories,omitempty"`

	TerritoriesAndSubordinates []string `json:"territoriesAndSubordinates,omitempty"`

	Territory []string `json:"territory,omitempty"`

	TerritoryAndSubordinates []string `json:"territoryAndSubordinates,omitempty"`
}

type RecordType struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RecordType"`

	*Metadata

	Active bool `json:"active,omitempty"`

	BusinessProcess string `json:"businessProcess,omitempty"`

	CompactLayoutAssignment string `json:"compactLayoutAssignment,omitempty"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	PicklistValues []*RecordTypePicklistValue `json:"picklistValues,omitempty"`
}

type RecordTypePicklistValue struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RecordTypePicklistValue"`

	Picklist string `json:"picklist,omitempty"`

	Values []*PicklistValue `json:"values,omitempty"`
}

type SearchLayouts struct {
	CustomTabListAdditionalFields []string `json:"customTabListAdditionalFields,omitempty"`

	ExcludedStandardButtons []string `json:"excludedStandardButtons,omitempty"`

	ListViewButtons []string `json:"listViewButtons,omitempty"`

	LookupDialogsAdditionalFields []string `json:"lookupDialogsAdditionalFields,omitempty"`

	LookupFilterFields []string `json:"lookupFilterFields,omitempty"`

	LookupPhoneDialogsAdditionalFields []string `json:"lookupPhoneDialogsAdditionalFields,omitempty"`

	SearchFilterFields []string `json:"searchFilterFields,omitempty"`

	SearchResultsAdditionalFields []string `json:"searchResultsAdditionalFields,omitempty"`

	SearchResultsCustomButtons []string `json:"searchResultsCustomButtons,omitempty"`
}

type SharingReason struct {
	*Metadata

	Label string `json:"label,omitempty"`
}

type SharingRecalculation struct {
	ClassName string `json:"className,omitempty"`
}

type ValidationRule struct {
	*Metadata

	Active bool `json:"active,omitempty"`

	Description string `json:"description,omitempty"`

	ErrorConditionFormula string `json:"errorConditionFormula,omitempty"`

	ErrorDisplayField string `json:"errorDisplayField,omitempty"`

	ErrorMessage string `json:"errorMessage,omitempty"`
}

type WebLink struct {
	*Metadata

	Availability *WebLinkAvailability `json:"availability,omitempty"`

	Description string `json:"description,omitempty"`

	DisplayType *WebLinkDisplayType `json:"displayType,omitempty"`

	EncodingKey *Encoding `json:"encodingKey,omitempty"`

	HasMenubar bool `json:"hasMenubar,omitempty"`

	HasScrollbars bool `json:"hasScrollbars,omitempty"`

	HasToolbar bool `json:"hasToolbar,omitempty"`

	Height int32 `json:"height,omitempty"`

	IsResizable bool `json:"isResizable,omitempty"`

	LinkType *WebLinkType `json:"linkType,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	OpenType *WebLinkWindowType `json:"openType,omitempty"`

	Page string `json:"page,omitempty"`

	Position *WebLinkPosition `json:"position,omitempty"`

	Protected bool `json:"protected,omitempty"`

	RequireRowSelection bool `json:"requireRowSelection,omitempty"`

	Scontrol string `json:"scontrol,omitempty"`

	ShowsLocation bool `json:"showsLocation,omitempty"`

	ShowsStatus bool `json:"showsStatus,omitempty"`

	Url string `json:"url,omitempty"`

	Width int32 `json:"width,omitempty"`
}

type CustomObjectTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomObjectTranslation"`

	*Metadata

	CaseValues []*ObjectNameCaseValue `json:"caseValues,omitempty"`

	Fields []*CustomFieldTranslation `json:"fields,omitempty"`

	Gender *Gender `json:"gender,omitempty"`

	Layouts []*LayoutTranslation `json:"layouts,omitempty"`

	NameFieldLabel string `json:"nameFieldLabel,omitempty"`

	QuickActions []*QuickActionTranslation `json:"quickActions,omitempty"`

	RecordTypes []*RecordTypeTranslation `json:"recordTypes,omitempty"`

	SharingReasons []*SharingReasonTranslation `json:"sharingReasons,omitempty"`

	StandardFields []*StandardFieldTranslation `json:"standardFields,omitempty"`

	StartsWith *StartsWith `json:"startsWith,omitempty"`

	ValidationRules []*ValidationRuleTranslation `json:"validationRules,omitempty"`

	WebLinks []*WebLinkTranslation `json:"webLinks,omitempty"`

	WorkflowTasks []*WorkflowTaskTranslation `json:"workflowTasks,omitempty"`
}

type ObjectNameCaseValue struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ObjectNameCaseValue"`

	Article *Article `json:"article,omitempty"`

	CaseType *CaseType `json:"caseType,omitempty"`

	Plural bool `json:"plural,omitempty"`

	Possessive *Possessive `json:"possessive,omitempty"`

	Value string `json:"value,omitempty"`
}

type CustomFieldTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomFieldTranslation"`

	CaseValues []*ObjectNameCaseValue `json:"caseValues,omitempty"`

	Gender *Gender `json:"gender,omitempty"`

	Help string `json:"help,omitempty"`

	Label string `json:"label,omitempty"`

	LookupFilter *LookupFilterTranslation `json:"lookupFilter,omitempty"`

	Name string `json:"name,omitempty"`

	PicklistValues []*PicklistValueTranslation `json:"picklistValues,omitempty"`

	RelationshipLabel string `json:"relationshipLabel,omitempty"`

	StartsWith *StartsWith `json:"startsWith,omitempty"`
}

type LookupFilterTranslation struct {
	ErrorMessage string `json:"errorMessage,omitempty"`

	InformationalMessage string `json:"informationalMessage,omitempty"`
}

type PicklistValueTranslation struct {
	MasterLabel string `json:"masterLabel,omitempty"`

	Translation string `json:"translation,omitempty"`
}

type LayoutTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LayoutTranslation"`

	Layout string `json:"layout,omitempty"`

	LayoutType string `json:"layoutType,omitempty"`

	Sections []*LayoutSectionTranslation `json:"sections,omitempty"`
}

type LayoutSectionTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LayoutSectionTranslation"`

	Label string `json:"label,omitempty"`

	Section string `json:"section,omitempty"`
}

type QuickActionTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata QuickActionTranslation"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type RecordTypeTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RecordTypeTranslation"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type SharingReasonTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SharingReasonTranslation"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type StandardFieldTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata StandardFieldTranslation"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type ValidationRuleTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ValidationRuleTranslation"`

	ErrorMessage string `json:"errorMessage,omitempty"`

	Name string `json:"name,omitempty"`
}

type WebLinkTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WebLinkTranslation"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type WorkflowTaskTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowTaskTranslation"`

	Description string `json:"description,omitempty"`

	Name string `json:"name,omitempty"`

	Subject string `json:"subject,omitempty"`
}

type CustomPageWebLink struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomPageWebLink"`

	*Metadata

	Availability *WebLinkAvailability `json:"availability,omitempty"`

	Description string `json:"description,omitempty"`

	DisplayType *WebLinkDisplayType `json:"displayType,omitempty"`

	EncodingKey *Encoding `json:"encodingKey,omitempty"`

	HasMenubar bool `json:"hasMenubar,omitempty"`

	HasScrollbars bool `json:"hasScrollbars,omitempty"`

	HasToolbar bool `json:"hasToolbar,omitempty"`

	Height int32 `json:"height,omitempty"`

	IsResizable bool `json:"isResizable,omitempty"`

	LinkType *WebLinkType `json:"linkType,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	OpenType *WebLinkWindowType `json:"openType,omitempty"`

	Page string `json:"page,omitempty"`

	Position *WebLinkPosition `json:"position,omitempty"`

	Protected bool `json:"protected,omitempty"`

	RequireRowSelection bool `json:"requireRowSelection,omitempty"`

	Scontrol string `json:"scontrol,omitempty"`

	ShowsLocation bool `json:"showsLocation,omitempty"`

	ShowsStatus bool `json:"showsStatus,omitempty"`

	Url string `json:"url,omitempty"`

	Width int32 `json:"width,omitempty"`
}

type CustomPermission struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomPermission"`

	*Metadata

	ConnectedApp string `json:"connectedApp,omitempty"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	RequiredPermission []*CustomPermissionDependencyRequired `json:"requiredPermission,omitempty"`
}

type CustomPermissionDependencyRequired struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomPermissionDependencyRequired"`

	CustomPermission string `json:"customPermission,omitempty"`

	Dependency bool `json:"dependency,omitempty"`
}

type CustomSite struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomSite"`

	*Metadata

	Active bool `json:"active,omitempty"`

	AllowHomePage bool `json:"allowHomePage,omitempty"`

	AllowStandardAnswersPages bool `json:"allowStandardAnswersPages,omitempty"`

	AllowStandardIdeasPages bool `json:"allowStandardIdeasPages,omitempty"`

	AllowStandardLookups bool `json:"allowStandardLookups,omitempty"`

	AllowStandardSearch bool `json:"allowStandardSearch,omitempty"`

	AnalyticsTrackingCode string `json:"analyticsTrackingCode,omitempty"`

	AuthorizationRequiredPage string `json:"authorizationRequiredPage,omitempty"`

	BandwidthExceededPage string `json:"bandwidthExceededPage,omitempty"`

	ChangePasswordPage string `json:"changePasswordPage,omitempty"`

	ChatterAnswersForgotPasswordConfirmPage string `json:"chatterAnswersForgotPasswordConfirmPage,omitempty"`

	ChatterAnswersForgotPasswordPage string `json:"chatterAnswersForgotPasswordPage,omitempty"`

	ChatterAnswersHelpPage string `json:"chatterAnswersHelpPage,omitempty"`

	ChatterAnswersLoginPage string `json:"chatterAnswersLoginPage,omitempty"`

	ChatterAnswersRegistrationPage string `json:"chatterAnswersRegistrationPage,omitempty"`

	ClickjackProtectionLevel *SiteClickjackProtectionLevel `json:"clickjackProtectionLevel,omitempty"`

	CustomWebAddresses []*SiteWebAddress `json:"customWebAddresses,omitempty"`

	Description string `json:"description,omitempty"`

	FavoriteIcon string `json:"favoriteIcon,omitempty"`

	FileNotFoundPage string `json:"fileNotFoundPage,omitempty"`

	ForgotPasswordPage string `json:"forgotPasswordPage,omitempty"`

	GenericErrorPage string `json:"genericErrorPage,omitempty"`

	GuestProfile string `json:"guestProfile,omitempty"`

	InMaintenancePage string `json:"inMaintenancePage,omitempty"`

	InactiveIndexPage string `json:"inactiveIndexPage,omitempty"`

	IndexPage string `json:"indexPage,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	MyProfilePage string `json:"myProfilePage,omitempty"`

	Portal string `json:"portal,omitempty"`

	RequireHttps bool `json:"requireHttps,omitempty"`

	RequireInsecurePortalAccess bool `json:"requireInsecurePortalAccess,omitempty"`

	RobotsTxtPage string `json:"robotsTxtPage,omitempty"`

	RootComponent string `json:"rootComponent,omitempty"`

	SelfRegPage string `json:"selfRegPage,omitempty"`

	ServerIsDown string `json:"serverIsDown,omitempty"`

	SiteAdmin string `json:"siteAdmin,omitempty"`

	SiteRedirectMappings []*SiteRedirectMapping `json:"siteRedirectMappings,omitempty"`

	SiteTemplate string `json:"siteTemplate,omitempty"`

	SiteType *SiteType `json:"siteType,omitempty"`

	Subdomain string `json:"subdomain,omitempty"`

	UrlPathPrefix string `json:"urlPathPrefix,omitempty"`
}

type SiteWebAddress struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SiteWebAddress"`

	Certificate string `json:"certificate,omitempty"`

	DomainName string `json:"domainName,omitempty"`

	Primary bool `json:"primary,omitempty"`
}

type SiteRedirectMapping struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SiteRedirectMapping"`

	Action *SiteRedirect `json:"action,omitempty"`

	IsActive bool `json:"isActive,omitempty"`

	Source string `json:"source,omitempty"`

	Target string `json:"target,omitempty"`
}

type CustomTab struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomTab"`

	*Metadata

	ActionOverrides []*ActionOverride `json:"actionOverrides,omitempty"`

	AuraComponent string `json:"auraComponent,omitempty"`

	CustomObject bool `json:"customObject,omitempty"`

	Description string `json:"description,omitempty"`

	FlexiPage string `json:"flexiPage,omitempty"`

	FrameHeight int32 `json:"frameHeight,omitempty"`

	HasSidebar bool `json:"hasSidebar,omitempty"`

	Icon string `json:"icon,omitempty"`

	Label string `json:"label,omitempty"`

	MobileReady bool `json:"mobileReady,omitempty"`

	Motif string `json:"motif,omitempty"`

	Page string `json:"page,omitempty"`

	Scontrol string `json:"scontrol,omitempty"`

	SplashPageLink string `json:"splashPageLink,omitempty"`

	Url string `json:"url,omitempty"`

	UrlEncodingKey *Encoding `json:"urlEncodingKey,omitempty"`
}

type Dashboard struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Dashboard"`

	*Metadata

	BackgroundEndColor string `json:"backgroundEndColor,omitempty"`

	BackgroundFadeDirection *ChartBackgroundDirection `json:"backgroundFadeDirection,omitempty"`

	BackgroundStartColor string `json:"backgroundStartColor,omitempty"`

	DashboardFilters []*DashboardFilter `json:"dashboardFilters,omitempty"`

	DashboardGridLayout *DashboardGridLayout `json:"dashboardGridLayout,omitempty"`

	DashboardResultRefreshedDate string `json:"dashboardResultRefreshedDate,omitempty"`

	DashboardResultRunningUser string `json:"dashboardResultRunningUser,omitempty"`

	DashboardType *DashboardType `json:"dashboardType,omitempty"`

	Description string `json:"description,omitempty"`

	FolderName string `json:"folderName,omitempty"`

	IsGridLayout bool `json:"isGridLayout,omitempty"`

	LeftSection *DashboardComponentSection `json:"leftSection,omitempty"`

	MiddleSection *DashboardComponentSection `json:"middleSection,omitempty"`

	RightSection *DashboardComponentSection `json:"rightSection,omitempty"`

	RunningUser string `json:"runningUser,omitempty"`

	TextColor string `json:"textColor,omitempty"`

	Title string `json:"title,omitempty"`

	TitleColor string `json:"titleColor,omitempty"`

	TitleSize int32 `json:"titleSize,omitempty"`
}

type DashboardFilter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DashboardFilter"`

	DashboardFilterOptions []*DashboardFilterOption `json:"dashboardFilterOptions,omitempty"`

	Name string `json:"name,omitempty"`
}

type DashboardFilterOption struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DashboardFilterOption"`

	Operator *DashboardFilterOperation `json:"operator,omitempty"`

	Values []string `json:"values,omitempty"`
}

type DashboardGridLayout struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DashboardGridLayout"`

	DashboardGridComponents []*DashboardGridComponent `json:"dashboardGridComponents,omitempty"`

	NumberOfColumns int32 `json:"numberOfColumns,omitempty"`

	RowHeight int32 `json:"rowHeight,omitempty"`
}

type DashboardGridComponent struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DashboardGridComponent"`

	ColSpan int32 `json:"colSpan,omitempty"`

	ColumnIndex int32 `json:"columnIndex,omitempty"`

	DashboardComponent *DashboardComponent `json:"dashboardComponent,omitempty"`

	RowIndex int32 `json:"rowIndex,omitempty"`

	RowSpan int32 `json:"rowSpan,omitempty"`
}

type DashboardComponent struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DashboardComponent"`

	AutoselectColumnsFromReport bool `json:"autoselectColumnsFromReport,omitempty"`

	ChartAxisRange *ChartRangeType `json:"chartAxisRange,omitempty"`

	ChartAxisRangeMax float64 `json:"chartAxisRangeMax,omitempty"`

	ChartAxisRangeMin float64 `json:"chartAxisRangeMin,omitempty"`

	ChartSummary []*ChartSummary `json:"chartSummary,omitempty"`

	ComponentType *DashboardComponentType `json:"componentType,omitempty"`

	DashboardFilterColumns []*DashboardFilterColumn `json:"dashboardFilterColumns,omitempty"`

	DashboardTableColumn []*DashboardTableColumn `json:"dashboardTableColumn,omitempty"`

	DisplayUnits *ChartUnits `json:"displayUnits,omitempty"`

	DrillDownUrl string `json:"drillDownUrl,omitempty"`

	DrillEnabled bool `json:"drillEnabled,omitempty"`

	DrillToDetailEnabled bool `json:"drillToDetailEnabled,omitempty"`

	EnableHover bool `json:"enableHover,omitempty"`

	ExpandOthers bool `json:"expandOthers,omitempty"`

	Footer string `json:"footer,omitempty"`

	GaugeMax float64 `json:"gaugeMax,omitempty"`

	GaugeMin float64 `json:"gaugeMin,omitempty"`

	GroupingColumn []string `json:"groupingColumn,omitempty"`

	Header string `json:"header,omitempty"`

	IndicatorBreakpoint1 float64 `json:"indicatorBreakpoint1,omitempty"`

	IndicatorBreakpoint2 float64 `json:"indicatorBreakpoint2,omitempty"`

	IndicatorHighColor string `json:"indicatorHighColor,omitempty"`

	IndicatorLowColor string `json:"indicatorLowColor,omitempty"`

	IndicatorMiddleColor string `json:"indicatorMiddleColor,omitempty"`

	LegendPosition *ChartLegendPosition `json:"legendPosition,omitempty"`

	MaxValuesDisplayed int32 `json:"maxValuesDisplayed,omitempty"`

	MetricLabel string `json:"metricLabel,omitempty"`

	Page string `json:"page,omitempty"`

	PageHeightInPixels int32 `json:"pageHeightInPixels,omitempty"`

	Report string `json:"report,omitempty"`

	Scontrol string `json:"scontrol,omitempty"`

	ScontrolHeightInPixels int32 `json:"scontrolHeightInPixels,omitempty"`

	ShowPercentage bool `json:"showPercentage,omitempty"`

	ShowPicturesOnCharts bool `json:"showPicturesOnCharts,omitempty"`

	ShowPicturesOnTables bool `json:"showPicturesOnTables,omitempty"`

	ShowTotal bool `json:"showTotal,omitempty"`

	ShowValues bool `json:"showValues,omitempty"`

	SortBy *DashboardComponentFilter `json:"sortBy,omitempty"`

	Title string `json:"title,omitempty"`

	UseReportChart bool `json:"useReportChart,omitempty"`
}

type ChartSummary struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ChartSummary"`

	Aggregate *ReportSummaryType `json:"aggregate,omitempty"`

	AxisBinding *ChartAxis `json:"axisBinding,omitempty"`

	Column string `json:"column,omitempty"`
}

type DashboardFilterColumn struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DashboardFilterColumn"`

	Column string `json:"column,omitempty"`
}

type DashboardTableColumn struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DashboardTableColumn"`

	AggregateType *ReportSummaryType `json:"aggregateType,omitempty"`

	CalculatePercent bool `json:"calculatePercent,omitempty"`

	Column string `json:"column,omitempty"`

	DecimalPlaces int32 `json:"decimalPlaces,omitempty"`

	ShowTotal bool `json:"showTotal,omitempty"`

	SortBy *DashboardComponentFilter `json:"sortBy,omitempty"`
}

type DashboardComponentSection struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DashboardComponentSection"`

	ColumnSize *DashboardComponentSize `json:"columnSize,omitempty"`

	Components []*DashboardComponent `json:"components,omitempty"`
}

type DataCategoryGroup struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DataCategoryGroup"`

	*Metadata

	Active bool `json:"active,omitempty"`

	DataCategory *DataCategory `json:"dataCategory,omitempty"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	ObjectUsage *ObjectUsage `json:"objectUsage,omitempty"`
}

type DataCategory struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DataCategory"`

	DataCategory []*DataCategory `json:"dataCategory,omitempty"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type ObjectUsage struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ObjectUsage"`

	Object []string `json:"object,omitempty"`
}

type DelegateGroup struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DelegateGroup"`

	*Metadata

	CustomObjects []string `json:"customObjects,omitempty"`

	Groups []string `json:"groups,omitempty"`

	Label string `json:"label,omitempty"`

	LoginAccess bool `json:"loginAccess,omitempty"`

	PermissionSets []string `json:"permissionSets,omitempty"`

	Profiles []string `json:"profiles,omitempty"`

	Roles []string `json:"roles,omitempty"`
}

type DuplicateRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DuplicateRule"`

	*Metadata

	ActionOnInsert *DupeActionType `json:"actionOnInsert,omitempty"`

	ActionOnUpdate *DupeActionType `json:"actionOnUpdate,omitempty"`

	AlertText string `json:"alertText,omitempty"`

	Description string `json:"description,omitempty"`

	DuplicateRuleFilter *DuplicateRuleFilter `json:"duplicateRuleFilter,omitempty"`

	DuplicateRuleMatchRules []*DuplicateRuleMatchRule `json:"duplicateRuleMatchRules,omitempty"`

	IsActive bool `json:"isActive,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	OperationsOnInsert []string `json:"operationsOnInsert,omitempty"`

	OperationsOnUpdate []string `json:"operationsOnUpdate,omitempty"`

	SecurityOption *DupeSecurityOptionType `json:"securityOption,omitempty"`

	SortOrder int32 `json:"sortOrder,omitempty"`
}

type DuplicateRuleFilter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DuplicateRuleFilter"`

	BooleanFilter string `json:"booleanFilter,omitempty"`

	DuplicateRuleFilterItems []*DuplicateRuleFilterItem `json:"duplicateRuleFilterItems,omitempty"`
}

type DuplicateRuleMatchRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DuplicateRuleMatchRule"`

	MatchRuleSObjectType string `json:"matchRuleSObjectType,omitempty"`

	MatchingRule string `json:"matchingRule,omitempty"`

	ObjectMapping *ObjectMapping `json:"objectMapping,omitempty"`
}

type ObjectMapping struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ObjectMapping"`

	InputObject string `json:"inputObject,omitempty"`

	MappingFields []*ObjectMappingField `json:"mappingFields,omitempty"`

	OutputObject string `json:"outputObject,omitempty"`
}

type ObjectMappingField struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ObjectMappingField"`

	InputField string `json:"inputField,omitempty"`

	OutputField string `json:"outputField,omitempty"`
}

type EmbeddedServiceConfig struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EmbeddedServiceConfig"`

	*Metadata

	MasterLabel string `json:"masterLabel,omitempty"`

	Site string `json:"site,omitempty"`
}

type EmbeddedServiceLiveAgent struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EmbeddedServiceLiveAgent"`

	*Metadata

	EmbeddedServiceConfig string `json:"embeddedServiceConfig,omitempty"`

	LiveAgentChatUrl string `json:"liveAgentChatUrl,omitempty"`

	LiveAgentContentUrl string `json:"liveAgentContentUrl,omitempty"`

	LiveChatButton string `json:"liveChatButton,omitempty"`

	LiveChatDeployment string `json:"liveChatDeployment,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`
}

type EntitlementProcess struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EntitlementProcess"`

	*Metadata

	SObjectType string `json:"SObjectType,omitempty"`

	Active bool `json:"active,omitempty"`

	BusinessHours string `json:"businessHours,omitempty"`

	Description string `json:"description,omitempty"`

	EntryStartDateField string `json:"entryStartDateField,omitempty"`

	ExitCriteriaBooleanFilter string `json:"exitCriteriaBooleanFilter,omitempty"`

	ExitCriteriaFilterItems []*FilterItem `json:"exitCriteriaFilterItems,omitempty"`

	ExitCriteriaFormula string `json:"exitCriteriaFormula,omitempty"`

	IsVersionDefault bool `json:"isVersionDefault,omitempty"`

	Milestones []*EntitlementProcessMilestoneItem `json:"milestones,omitempty"`

	Name string `json:"name,omitempty"`

	VersionMaster string `json:"versionMaster,omitempty"`

	VersionNotes string `json:"versionNotes,omitempty"`

	VersionNumber int32 `json:"versionNumber,omitempty"`
}

type EntitlementProcessMilestoneItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EntitlementProcessMilestoneItem"`

	BusinessHours string `json:"businessHours,omitempty"`

	CriteriaBooleanFilter string `json:"criteriaBooleanFilter,omitempty"`

	MilestoneCriteriaFilterItems []*FilterItem `json:"milestoneCriteriaFilterItems,omitempty"`

	MilestoneCriteriaFormula string `json:"milestoneCriteriaFormula,omitempty"`

	MilestoneName string `json:"milestoneName,omitempty"`

	MinutesCustomClass string `json:"minutesCustomClass,omitempty"`

	MinutesToComplete int32 `json:"minutesToComplete,omitempty"`

	SuccessActions []*WorkflowActionReference `json:"successActions,omitempty"`

	TimeTriggers []*EntitlementProcessMilestoneTimeTrigger `json:"timeTriggers,omitempty"`

	UseCriteriaStartTime bool `json:"useCriteriaStartTime,omitempty"`
}

type EntitlementProcessMilestoneTimeTrigger struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EntitlementProcessMilestoneTimeTrigger"`

	Actions []*WorkflowActionReference `json:"actions,omitempty"`

	TimeLength int32 `json:"timeLength,omitempty"`

	WorkflowTimeTriggerUnit *MilestoneTimeUnits `json:"workflowTimeTriggerUnit,omitempty"`
}

type EntitlementSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EntitlementSettings"`

	*Metadata

	AssetLookupLimitedToActiveEntitlementsOnAccount bool `json:"assetLookupLimitedToActiveEntitlementsOnAccount,omitempty"`

	AssetLookupLimitedToActiveEntitlementsOnContact bool `json:"assetLookupLimitedToActiveEntitlementsOnContact,omitempty"`

	AssetLookupLimitedToSameAccount bool `json:"assetLookupLimitedToSameAccount,omitempty"`

	AssetLookupLimitedToSameContact bool `json:"assetLookupLimitedToSameContact,omitempty"`

	EnableEntitlementVersioning bool `json:"enableEntitlementVersioning,omitempty"`

	EnableEntitlements bool `json:"enableEntitlements,omitempty"`

	EntitlementLookupLimitedToActiveStatus bool `json:"entitlementLookupLimitedToActiveStatus,omitempty"`

	EntitlementLookupLimitedToSameAccount bool `json:"entitlementLookupLimitedToSameAccount,omitempty"`

	EntitlementLookupLimitedToSameAsset bool `json:"entitlementLookupLimitedToSameAsset,omitempty"`

	EntitlementLookupLimitedToSameContact bool `json:"entitlementLookupLimitedToSameContact,omitempty"`
}

type EntitlementTemplate struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EntitlementTemplate"`

	*Metadata

	BusinessHours string `json:"businessHours,omitempty"`

	CasesPerEntitlement int32 `json:"casesPerEntitlement,omitempty"`

	EntitlementProcess string `json:"entitlementProcess,omitempty"`

	IsPerIncident bool `json:"isPerIncident,omitempty"`

	Term int32 `json:"term,omitempty"`

	Type_ string `json:"type,omitempty"`
}

type EscalationRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EscalationRule"`

	*Metadata

	Active bool `json:"active,omitempty"`

	RuleEntry []*RuleEntry `json:"ruleEntry,omitempty"`
}

type EscalationRules struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EscalationRules"`

	*Metadata

	EscalationRule []*EscalationRule `json:"escalationRule,omitempty"`
}

type EventDelivery struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EventDelivery"`

	*Metadata

	EventParameters []*EventParameterMap `json:"eventParameters,omitempty"`

	EventSubscription string `json:"eventSubscription,omitempty"`

	ReferenceData string `json:"referenceData,omitempty"`

	Type_ *EventDeliveryType `json:"type,omitempty"`
}

type EventParameterMap struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EventParameterMap"`

	ParameterName string `json:"parameterName,omitempty"`

	ParameterValue string `json:"parameterValue,omitempty"`
}

type EventSubscription struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EventSubscription"`

	*Metadata

	Active bool `json:"active,omitempty"`

	EventParameters []*EventParameterMap `json:"eventParameters,omitempty"`

	EventType string `json:"eventType,omitempty"`

	ReferenceData string `json:"referenceData,omitempty"`
}

type EventType struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EventType"`

	*Metadata

	Description string `json:"description,omitempty"`

	Fields []*EventTypeParameter `json:"fields,omitempty"`

	Label string `json:"label,omitempty"`
}

type EventTypeParameter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EventTypeParameter"`

	DefaultValue string `json:"defaultValue,omitempty"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	MaxOccurs int32 `json:"maxOccurs,omitempty"`

	MinOccurs int32 `json:"minOccurs,omitempty"`

	Name string `json:"name,omitempty"`

	SObjectType string `json:"sObjectType,omitempty"`

	Type_ *FieldType `json:"type,omitempty"`
}

type ExternalDataSource struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ExternalDataSource"`

	*Metadata

	AuthProvider string `json:"authProvider,omitempty"`

	Certificate string `json:"certificate,omitempty"`

	CustomConfiguration string `json:"customConfiguration,omitempty"`

	Endpoint string `json:"endpoint,omitempty"`

	IsWritable bool `json:"isWritable,omitempty"`

	Label string `json:"label,omitempty"`

	OauthRefreshToken string `json:"oauthRefreshToken,omitempty"`

	OauthScope string `json:"oauthScope,omitempty"`

	OauthToken string `json:"oauthToken,omitempty"`

	Password string `json:"password,omitempty"`

	PrincipalType *ExternalPrincipalType `json:"principalType,omitempty"`

	Protocol *AuthenticationProtocol `json:"protocol,omitempty"`

	Repository string `json:"repository,omitempty"`

	Type_ *ExternalDataSourceType `json:"type,omitempty"`

	Username string `json:"username,omitempty"`

	Version string `json:"version,omitempty"`
}

type ExternalServiceRegistration struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ExternalServiceRegistration"`

	*Metadata

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	NamedCredential string `json:"namedCredential,omitempty"`

	Schema string `json:"schema,omitempty"`

	SchemaType string `json:"schemaType,omitempty"`

	SchemaUrl string `json:"schemaUrl,omitempty"`

	Status string `json:"status,omitempty"`
}

type FlexiPage struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlexiPage"`

	*Metadata

	Description string `json:"description,omitempty"`

	FlexiPageRegions []*FlexiPageRegion `json:"flexiPageRegions,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	PageTemplate string `json:"pageTemplate,omitempty"`

	ParentFlexiPage string `json:"parentFlexiPage,omitempty"`

	PlatformActionlist *PlatformActionList `json:"platformActionlist,omitempty"`

	QuickActionList *QuickActionList `json:"quickActionList,omitempty"`

	SobjectType string `json:"sobjectType,omitempty"`

	Type_ *FlexiPageType `json:"type,omitempty"`
}

type FlexiPageRegion struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlexiPageRegion"`

	Appendable *RegionFlagStatus `json:"appendable,omitempty"`

	ComponentInstances []*ComponentInstance `json:"componentInstances,omitempty"`

	Mode *FlexiPageRegionMode `json:"mode,omitempty"`

	Name string `json:"name,omitempty"`

	Prependable *RegionFlagStatus `json:"prependable,omitempty"`

	Replaceable *RegionFlagStatus `json:"replaceable,omitempty"`

	Type_ *FlexiPageRegionType `json:"type,omitempty"`
}

type ComponentInstance struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ComponentInstance"`

	ComponentInstanceProperties []*ComponentInstanceProperty `json:"componentInstanceProperties,omitempty"`

	ComponentName string `json:"componentName,omitempty"`
}

type ComponentInstanceProperty struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ComponentInstanceProperty"`

	Name string `json:"name,omitempty"`

	Type_ *ComponentInstancePropertyTypeEnum `json:"type,omitempty"`

	Value string `json:"value,omitempty"`
}

type QuickActionList struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata QuickActionList"`

	QuickActionListItems []*QuickActionListItem `json:"quickActionListItems,omitempty"`
}

type QuickActionListItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata QuickActionListItem"`

	QuickActionName string `json:"quickActionName,omitempty"`
}

type Flow struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Flow"`

	*Metadata

	ActionCalls []*FlowActionCall `json:"actionCalls,omitempty"`

	ApexPluginCalls []*FlowApexPluginCall `json:"apexPluginCalls,omitempty"`

	Assignments []*FlowAssignment `json:"assignments,omitempty"`

	Choices []*FlowChoice `json:"choices,omitempty"`

	Constants []*FlowConstant `json:"constants,omitempty"`

	Decisions []*FlowDecision `json:"decisions,omitempty"`

	Description string `json:"description,omitempty"`

	DynamicChoiceSets []*FlowDynamicChoiceSet `json:"dynamicChoiceSets,omitempty"`

	Formulas []*FlowFormula `json:"formulas,omitempty"`

	InterviewLabel string `json:"interviewLabel,omitempty"`

	Label string `json:"label,omitempty"`

	Loops []*FlowLoop `json:"loops,omitempty"`

	ProcessMetadataValues []*FlowMetadataValue `json:"processMetadataValues,omitempty"`

	ProcessType *FlowProcessType `json:"processType,omitempty"`

	RecordCreates []*FlowRecordCreate `json:"recordCreates,omitempty"`

	RecordDeletes []*FlowRecordDelete `json:"recordDeletes,omitempty"`

	RecordLookups []*FlowRecordLookup `json:"recordLookups,omitempty"`

	RecordUpdates []*FlowRecordUpdate `json:"recordUpdates,omitempty"`

	Screens []*FlowScreen `json:"screens,omitempty"`

	StartElementReference string `json:"startElementReference,omitempty"`

	Steps []*FlowStep `json:"steps,omitempty"`

	Subflows []*FlowSubflow `json:"subflows,omitempty"`

	TextTemplates []*FlowTextTemplate `json:"textTemplates,omitempty"`

	Variables []*FlowVariable `json:"variables,omitempty"`

	Waits []*FlowWait `json:"waits,omitempty"`
}

type FlowActionCall struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowActionCall"`

	*FlowNode

	ActionName string `json:"actionName,omitempty"`

	ActionType *InvocableActionType `json:"actionType,omitempty"`

	Connector *FlowConnector `json:"connector,omitempty"`

	FaultConnector *FlowConnector `json:"faultConnector,omitempty"`

	InputParameters []*FlowActionCallInputParameter `json:"inputParameters,omitempty"`

	OutputParameters []*FlowActionCallOutputParameter `json:"outputParameters,omitempty"`
}

type FlowNode struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowNode"`

	*FlowElement

	Label string `json:"label,omitempty"`

	LocationX int32 `json:"locationX,omitempty"`

	LocationY int32 `json:"locationY,omitempty"`
}

type FlowElement struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowElement"`

	*FlowBaseElement

	Description string `json:"description,omitempty"`

	Name string `json:"name,omitempty"`
}

type FlowBaseElement struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowBaseElement"`

	ProcessMetadataValues []*FlowMetadataValue `json:"processMetadataValues,omitempty"`
}

type FlowMetadataValue struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowMetadataValue"`

	Name string `json:"name,omitempty"`

	Value *FlowElementReferenceOrValue `json:"value,omitempty"`
}

type FlowElementReferenceOrValue struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowElementReferenceOrValue"`

	BooleanValue bool `json:"booleanValue,omitempty"`

	DateTimeValue time.Time `json:"dateTimeValue,omitempty"`

	DateValue time.Time `json:"dateValue,omitempty"`

	ElementReference string `json:"elementReference,omitempty"`

	NumberValue float64 `json:"numberValue,omitempty"`

	StringValue string `json:"stringValue,omitempty"`
}

type FlowActionCallInputParameter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowActionCallInputParameter"`

	*FlowBaseElement

	Name string `json:"name,omitempty"`

	Value *FlowElementReferenceOrValue `json:"value,omitempty"`
}

type FlowActionCallOutputParameter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowActionCallOutputParameter"`

	*FlowBaseElement

	AssignToReference string `json:"assignToReference,omitempty"`

	Name string `json:"name,omitempty"`
}

type FlowApexPluginCallInputParameter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowApexPluginCallInputParameter"`

	*FlowBaseElement

	Name string `json:"name,omitempty"`

	Value *FlowElementReferenceOrValue `json:"value,omitempty"`
}

type FlowApexPluginCallOutputParameter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowApexPluginCallOutputParameter"`

	*FlowBaseElement

	AssignToReference string `json:"assignToReference,omitempty"`

	Name string `json:"name,omitempty"`
}

type FlowAssignmentItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowAssignmentItem"`

	*FlowBaseElement

	AssignToReference string `json:"assignToReference,omitempty"`

	Operator *FlowAssignmentOperator `json:"operator,omitempty"`

	Value *FlowElementReferenceOrValue `json:"value,omitempty"`
}

type FlowChoiceUserInput struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowChoiceUserInput"`

	*FlowBaseElement

	IsRequired bool `json:"isRequired,omitempty"`

	PromptText string `json:"promptText,omitempty"`

	ValidationRule *FlowInputValidationRule `json:"validationRule,omitempty"`
}

type FlowInputValidationRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowInputValidationRule"`

	ErrorMessage string `json:"errorMessage,omitempty"`

	FormulaExpression string `json:"formulaExpression,omitempty"`
}

type FlowCondition struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowCondition"`

	*FlowBaseElement

	LeftValueReference string `json:"leftValueReference,omitempty"`

	Operator *FlowComparisonOperator `json:"operator,omitempty"`

	RightValue *FlowElementReferenceOrValue `json:"rightValue,omitempty"`
}

type FlowConnector struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowConnector"`

	*FlowBaseElement

	TargetReference string `json:"targetReference,omitempty"`
}

type FlowInputFieldAssignment struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowInputFieldAssignment"`

	*FlowBaseElement

	Field string `json:"field,omitempty"`

	Value *FlowElementReferenceOrValue `json:"value,omitempty"`
}

type FlowOutputFieldAssignment struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowOutputFieldAssignment"`

	*FlowBaseElement

	AssignToReference string `json:"assignToReference,omitempty"`

	Field string `json:"field,omitempty"`
}

type FlowRecordFilter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowRecordFilter"`

	*FlowBaseElement

	Field string `json:"field,omitempty"`

	Operator *FlowRecordFilterOperator `json:"operator,omitempty"`

	Value *FlowElementReferenceOrValue `json:"value,omitempty"`
}

type FlowScreenRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowScreenRule"`

	*FlowBaseElement

	ConditionLogic string `json:"conditionLogic,omitempty"`

	Conditions []*FlowCondition `json:"conditions,omitempty"`

	Label string `json:"label,omitempty"`

	RuleActions []*FlowScreenRuleAction `json:"ruleActions,omitempty"`
}

type FlowScreenRuleAction struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowScreenRuleAction"`

	*FlowBaseElement

	Attribute string `json:"attribute,omitempty"`

	FieldReference string `json:"fieldReference,omitempty"`

	Value *FlowElementReferenceOrValue `json:"value,omitempty"`
}

type FlowSubflowInputAssignment struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowSubflowInputAssignment"`

	*FlowBaseElement

	Name string `json:"name,omitempty"`

	Value *FlowElementReferenceOrValue `json:"value,omitempty"`
}

type FlowSubflowOutputAssignment struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowSubflowOutputAssignment"`

	*FlowBaseElement

	AssignToReference string `json:"assignToReference,omitempty"`

	Name string `json:"name,omitempty"`
}

type FlowWaitEventInputParameter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowWaitEventInputParameter"`

	*FlowBaseElement

	Name string `json:"name,omitempty"`

	Value *FlowElementReferenceOrValue `json:"value,omitempty"`
}

type FlowWaitEventOutputParameter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowWaitEventOutputParameter"`

	*FlowBaseElement

	AssignToReference string `json:"assignToReference,omitempty"`

	Name string `json:"name,omitempty"`
}

type FlowChoice struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowChoice"`

	*FlowElement

	ChoiceText string `json:"choiceText,omitempty"`

	DataType *FlowDataType `json:"dataType,omitempty"`

	UserInput *FlowChoiceUserInput `json:"userInput,omitempty"`

	Value *FlowElementReferenceOrValue `json:"value,omitempty"`
}

type FlowConstant struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowConstant"`

	*FlowElement

	DataType *FlowDataType `json:"dataType,omitempty"`

	Value *FlowElementReferenceOrValue `json:"value,omitempty"`
}

type FlowDynamicChoiceSet struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowDynamicChoiceSet"`

	*FlowElement

	DataType *FlowDataType `json:"dataType,omitempty"`

	DisplayField string `json:"displayField,omitempty"`

	Filters []*FlowRecordFilter `json:"filters,omitempty"`

	Limit int32 `json:"limit,omitempty"`

	Object string `json:"object,omitempty"`

	OutputAssignments []*FlowOutputFieldAssignment `json:"outputAssignments,omitempty"`

	PicklistField string `json:"picklistField,omitempty"`

	PicklistObject string `json:"picklistObject,omitempty"`

	SortField string `json:"sortField,omitempty"`

	SortOrder *SortOrder `json:"sortOrder,omitempty"`

	ValueField string `json:"valueField,omitempty"`
}

type FlowFormula struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowFormula"`

	*FlowElement

	DataType *FlowDataType `json:"dataType,omitempty"`

	Expression string `json:"expression,omitempty"`

	Scale int32 `json:"scale,omitempty"`
}

type FlowRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowRule"`

	*FlowElement

	ConditionLogic string `json:"conditionLogic,omitempty"`

	Conditions []*FlowCondition `json:"conditions,omitempty"`

	Connector *FlowConnector `json:"connector,omitempty"`

	Label string `json:"label,omitempty"`
}

type FlowScreenField struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowScreenField"`

	*FlowElement

	ChoiceReferences []string `json:"choiceReferences,omitempty"`

	DataType *FlowDataType `json:"dataType,omitempty"`

	DefaultSelectedChoiceReference string `json:"defaultSelectedChoiceReference,omitempty"`

	DefaultValue *FlowElementReferenceOrValue `json:"defaultValue,omitempty"`

	FieldText string `json:"fieldText,omitempty"`

	FieldType *FlowScreenFieldType `json:"fieldType,omitempty"`

	HelpText string `json:"helpText,omitempty"`

	IsRequired bool `json:"isRequired,omitempty"`

	Scale int32 `json:"scale,omitempty"`

	ValidationRule *FlowInputValidationRule `json:"validationRule,omitempty"`
}

type FlowTextTemplate struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowTextTemplate"`

	*FlowElement

	Text string `json:"text,omitempty"`
}

type FlowVariable struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowVariable"`

	*FlowElement

	DataType *FlowDataType `json:"dataType,omitempty"`

	IsCollection bool `json:"isCollection,omitempty"`

	IsInput bool `json:"isInput,omitempty"`

	IsOutput bool `json:"isOutput,omitempty"`

	ObjectType string `json:"objectType,omitempty"`

	Scale int32 `json:"scale,omitempty"`

	Value *FlowElementReferenceOrValue `json:"value,omitempty"`
}

type FlowWaitEvent struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowWaitEvent"`

	*FlowElement

	ConditionLogic string `json:"conditionLogic,omitempty"`

	Conditions []*FlowCondition `json:"conditions,omitempty"`

	Connector *FlowConnector `json:"connector,omitempty"`

	EventType string `json:"eventType,omitempty"`

	InputParameters []*FlowWaitEventInputParameter `json:"inputParameters,omitempty"`

	Label string `json:"label,omitempty"`

	OutputParameters []*FlowWaitEventOutputParameter `json:"outputParameters,omitempty"`
}

type FlowApexPluginCall struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowApexPluginCall"`

	*FlowNode

	ApexClass string `json:"apexClass,omitempty"`

	Connector *FlowConnector `json:"connector,omitempty"`

	FaultConnector *FlowConnector `json:"faultConnector,omitempty"`

	InputParameters []*FlowApexPluginCallInputParameter `json:"inputParameters,omitempty"`

	OutputParameters []*FlowApexPluginCallOutputParameter `json:"outputParameters,omitempty"`
}

type FlowAssignment struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowAssignment"`

	*FlowNode

	AssignmentItems []*FlowAssignmentItem `json:"assignmentItems,omitempty"`

	Connector *FlowConnector `json:"connector,omitempty"`
}

type FlowDecision struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowDecision"`

	*FlowNode

	DefaultConnector *FlowConnector `json:"defaultConnector,omitempty"`

	DefaultConnectorLabel string `json:"defaultConnectorLabel,omitempty"`

	Rules []*FlowRule `json:"rules,omitempty"`
}

type FlowLoop struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowLoop"`

	*FlowNode

	AssignNextValueToReference string `json:"assignNextValueToReference,omitempty"`

	CollectionReference string `json:"collectionReference,omitempty"`

	IterationOrder *IterationOrder `json:"iterationOrder,omitempty"`

	NextValueConnector *FlowConnector `json:"nextValueConnector,omitempty"`

	NoMoreValuesConnector *FlowConnector `json:"noMoreValuesConnector,omitempty"`
}

type FlowRecordCreate struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowRecordCreate"`

	*FlowNode

	AssignRecordIdToReference string `json:"assignRecordIdToReference,omitempty"`

	Connector *FlowConnector `json:"connector,omitempty"`

	FaultConnector *FlowConnector `json:"faultConnector,omitempty"`

	InputAssignments []*FlowInputFieldAssignment `json:"inputAssignments,omitempty"`

	InputReference string `json:"inputReference,omitempty"`

	Object string `json:"object,omitempty"`
}

type FlowRecordDelete struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowRecordDelete"`

	*FlowNode

	Connector *FlowConnector `json:"connector,omitempty"`

	FaultConnector *FlowConnector `json:"faultConnector,omitempty"`

	Filters []*FlowRecordFilter `json:"filters,omitempty"`

	InputReference string `json:"inputReference,omitempty"`

	Object string `json:"object,omitempty"`
}

type FlowRecordLookup struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowRecordLookup"`

	*FlowNode

	AssignNullValuesIfNoRecordsFound bool `json:"assignNullValuesIfNoRecordsFound,omitempty"`

	Connector *FlowConnector `json:"connector,omitempty"`

	FaultConnector *FlowConnector `json:"faultConnector,omitempty"`

	Filters []*FlowRecordFilter `json:"filters,omitempty"`

	Object string `json:"object,omitempty"`

	OutputAssignments []*FlowOutputFieldAssignment `json:"outputAssignments,omitempty"`

	OutputReference string `json:"outputReference,omitempty"`

	QueriedFields []string `json:"queriedFields,omitempty"`

	SortField string `json:"sortField,omitempty"`

	SortOrder *SortOrder `json:"sortOrder,omitempty"`
}

type FlowRecordUpdate struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowRecordUpdate"`

	*FlowNode

	Connector *FlowConnector `json:"connector,omitempty"`

	FaultConnector *FlowConnector `json:"faultConnector,omitempty"`

	Filters []*FlowRecordFilter `json:"filters,omitempty"`

	InputAssignments []*FlowInputFieldAssignment `json:"inputAssignments,omitempty"`

	InputReference string `json:"inputReference,omitempty"`

	Object string `json:"object,omitempty"`
}

type FlowScreen struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowScreen"`

	*FlowNode

	AllowBack bool `json:"allowBack,omitempty"`

	AllowFinish bool `json:"allowFinish,omitempty"`

	AllowPause bool `json:"allowPause,omitempty"`

	Connector *FlowConnector `json:"connector,omitempty"`

	Fields []*FlowScreenField `json:"fields,omitempty"`

	HelpText string `json:"helpText,omitempty"`

	PausedText string `json:"pausedText,omitempty"`

	Rules []*FlowScreenRule `json:"rules,omitempty"`
}

type FlowStep struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowStep"`

	*FlowNode

	Connectors []*FlowConnector `json:"connectors,omitempty"`
}

type FlowSubflow struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowSubflow"`

	*FlowNode

	Connector *FlowConnector `json:"connector,omitempty"`

	FlowName string `json:"flowName,omitempty"`

	InputAssignments []*FlowSubflowInputAssignment `json:"inputAssignments,omitempty"`

	OutputAssignments []*FlowSubflowOutputAssignment `json:"outputAssignments,omitempty"`
}

type FlowWait struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowWait"`

	*FlowNode

	DefaultConnector *FlowConnector `json:"defaultConnector,omitempty"`

	DefaultConnectorLabel string `json:"defaultConnectorLabel,omitempty"`

	FaultConnector *FlowConnector `json:"faultConnector,omitempty"`

	WaitEvents []*FlowWaitEvent `json:"waitEvents,omitempty"`
}

type FlowDefinition struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FlowDefinition"`

	*Metadata

	ActiveVersionNumber int32 `json:"activeVersionNumber,omitempty"`

	Description string `json:"description,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`
}

type Folder struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Folder"`

	*Metadata

	AccessType *FolderAccessTypes `json:"accessType,omitempty"`

	FolderShares []*FolderShare `json:"folderShares,omitempty"`

	Name string `json:"name,omitempty"`

	PublicFolderAccess *PublicFolderAccess `json:"publicFolderAccess,omitempty"`

	SharedTo *SharedTo `json:"sharedTo,omitempty"`
}

type FolderShare struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FolderShare"`

	AccessLevel *FolderShareAccessLevel `json:"accessLevel,omitempty"`

	SharedTo string `json:"sharedTo,omitempty"`

	SharedToType *FolderSharedToType `json:"sharedToType,omitempty"`
}

type DashboardFolder struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DashboardFolder"`

	*Folder
}

type DocumentFolder struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DocumentFolder"`

	*Folder
}

type EmailFolder struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EmailFolder"`

	*Folder
}

type ReportFolder struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportFolder"`

	*Folder
}

type ForecastingSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ForecastingSettings"`

	*Metadata

	DisplayCurrency *DisplayCurrency `json:"displayCurrency,omitempty"`

	EnableForecasts bool `json:"enableForecasts,omitempty"`

	ForecastingCategoryMappings []*ForecastingCategoryMapping `json:"forecastingCategoryMappings,omitempty"`

	ForecastingTypeSettings []*ForecastingTypeSettings `json:"forecastingTypeSettings,omitempty"`
}

type ForecastingCategoryMapping struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ForecastingCategoryMapping"`

	ForecastingItemCategoryApiName string `json:"forecastingItemCategoryApiName,omitempty"`

	WeightedSourceCategories []*WeightedSourceCategory `json:"weightedSourceCategories,omitempty"`
}

type WeightedSourceCategory struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WeightedSourceCategory"`

	SourceCategoryApiName string `json:"sourceCategoryApiName,omitempty"`

	Weight float64 `json:"weight,omitempty"`
}

type ForecastingTypeSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ForecastingTypeSettings"`

	Active bool `json:"active,omitempty"`

	AdjustmentsSettings *AdjustmentsSettings `json:"adjustmentsSettings,omitempty"`

	DisplayedCategoryApiNames []string `json:"displayedCategoryApiNames,omitempty"`

	ForecastRangeSettings *ForecastRangeSettings `json:"forecastRangeSettings,omitempty"`

	ForecastedCategoryApiNames []string `json:"forecastedCategoryApiNames,omitempty"`

	IsAmount bool `json:"isAmount,omitempty"`

	IsAvailable bool `json:"isAvailable,omitempty"`

	IsQuantity bool `json:"isQuantity,omitempty"`

	ManagerAdjustableCategoryApiNames []string `json:"managerAdjustableCategoryApiNames,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	Name string `json:"name,omitempty"`

	OpportunityListFieldsLabelMappings []*OpportunityListFieldsLabelMapping `json:"opportunityListFieldsLabelMappings,omitempty"`

	OpportunityListFieldsSelectedSettings *OpportunityListFieldsSelectedSettings `json:"opportunityListFieldsSelectedSettings,omitempty"`

	OpportunityListFieldsUnselectedSettings *OpportunityListFieldsUnselectedSettings `json:"opportunityListFieldsUnselectedSettings,omitempty"`

	OwnerAdjustableCategoryApiNames []string `json:"ownerAdjustableCategoryApiNames,omitempty"`

	QuotasSettings *QuotasSettings `json:"quotasSettings,omitempty"`
}

type AdjustmentsSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AdjustmentsSettings"`

	EnableAdjustments bool `json:"enableAdjustments,omitempty"`

	EnableOwnerAdjustments bool `json:"enableOwnerAdjustments,omitempty"`
}

type ForecastRangeSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ForecastRangeSettings"`

	Beginning int32 `json:"beginning,omitempty"`

	Displaying int32 `json:"displaying,omitempty"`

	PeriodType *PeriodTypes `json:"periodType,omitempty"`
}

type OpportunityListFieldsLabelMapping struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata OpportunityListFieldsLabelMapping"`

	Field string `json:"field,omitempty"`

	Label string `json:"label,omitempty"`
}

type OpportunityListFieldsSelectedSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata OpportunityListFieldsSelectedSettings"`

	Field []string `json:"field,omitempty"`
}

type OpportunityListFieldsUnselectedSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata OpportunityListFieldsUnselectedSettings"`

	Field []string `json:"field,omitempty"`
}

type QuotasSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata QuotasSettings"`

	ShowQuotas bool `json:"showQuotas,omitempty"`
}

type GlobalValueSet struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata GlobalValueSet"`

	*Metadata

	CustomValue []*CustomValue `json:"customValue,omitempty"`

	Description string `json:"description,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	Sorted bool `json:"sorted,omitempty"`
}

type GlobalValueSetTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata GlobalValueSetTranslation"`

	*Metadata

	ValueTranslation []*ValueTranslation `json:"valueTranslation,omitempty"`
}

type ValueTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ValueTranslation"`

	MasterLabel string `json:"masterLabel,omitempty"`

	Translation string `json:"translation,omitempty"`
}

type Group struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Group"`

	*Metadata

	DoesIncludeBosses bool `json:"doesIncludeBosses,omitempty"`

	Name string `json:"name,omitempty"`
}

type HomePageComponent struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata HomePageComponent"`

	*Metadata

	Body string `json:"body,omitempty"`

	Height int32 `json:"height,omitempty"`

	Links []string `json:"links,omitempty"`

	Page string `json:"page,omitempty"`

	PageComponentType *PageComponentType `json:"pageComponentType,omitempty"`

	ShowLabel bool `json:"showLabel,omitempty"`

	ShowScrollbars bool `json:"showScrollbars,omitempty"`

	Width *PageComponentWidth `json:"width,omitempty"`
}

type HomePageLayout struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata HomePageLayout"`

	*Metadata

	NarrowComponents []string `json:"narrowComponents,omitempty"`

	WideComponents []string `json:"wideComponents,omitempty"`
}

type IdeasSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata IdeasSettings"`

	*Metadata

	EnableChatterProfile bool `json:"enableChatterProfile,omitempty"`

	EnableIdeaThemes bool `json:"enableIdeaThemes,omitempty"`

	EnableIdeas bool `json:"enableIdeas,omitempty"`

	EnableIdeasReputation bool `json:"enableIdeasReputation,omitempty"`

	HalfLife float64 `json:"halfLife,omitempty"`

	IdeasProfilePage string `json:"ideasProfilePage,omitempty"`
}

type InstalledPackage struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata InstalledPackage"`

	*Metadata

	Password string `json:"password,omitempty"`

	VersionNumber string `json:"versionNumber,omitempty"`
}

type KeywordList struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata KeywordList"`

	*Metadata

	Description string `json:"description,omitempty"`

	Keywords []*Keyword `json:"keywords,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`
}

type Keyword struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Keyword"`

	Keyword string `json:"keyword,omitempty"`
}

type KnowledgeSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata KnowledgeSettings"`

	*Metadata

	Answers *KnowledgeAnswerSettings `json:"answers,omitempty"`

	Cases *KnowledgeCaseSettings `json:"cases,omitempty"`

	DefaultLanguage string `json:"defaultLanguage,omitempty"`

	EnableChatterQuestionKBDeflection bool `json:"enableChatterQuestionKBDeflection,omitempty"`

	EnableCreateEditOnArticlesTab bool `json:"enableCreateEditOnArticlesTab,omitempty"`

	EnableExternalMediaContent bool `json:"enableExternalMediaContent,omitempty"`

	EnableKnowledge bool `json:"enableKnowledge,omitempty"`

	Languages *KnowledgeLanguageSettings `json:"languages,omitempty"`

	ShowArticleSummariesCustomerPortal bool `json:"showArticleSummariesCustomerPortal,omitempty"`

	ShowArticleSummariesInternalApp bool `json:"showArticleSummariesInternalApp,omitempty"`

	ShowArticleSummariesPartnerPortal bool `json:"showArticleSummariesPartnerPortal,omitempty"`

	ShowValidationStatusField bool `json:"showValidationStatusField,omitempty"`

	SuggestedArticles *KnowledgeSuggestedArticlesSettings `json:"suggestedArticles,omitempty"`
}

type KnowledgeAnswerSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata KnowledgeAnswerSettings"`

	AssignTo string `json:"assignTo,omitempty"`

	DefaultArticleType string `json:"defaultArticleType,omitempty"`

	EnableArticleCreation bool `json:"enableArticleCreation,omitempty"`
}

type KnowledgeCaseSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata KnowledgeCaseSettings"`

	ArticlePDFCreationProfile string `json:"articlePDFCreationProfile,omitempty"`

	ArticlePublicSharingCommunities *KnowledgeCommunitiesSettings `json:"articlePublicSharingCommunities,omitempty"`

	ArticlePublicSharingSites *KnowledgeSitesSettings `json:"articlePublicSharingSites,omitempty"`

	ArticlePublicSharingSitesChatterAnswers *KnowledgeSitesSettings `json:"articlePublicSharingSitesChatterAnswers,omitempty"`

	AssignTo string `json:"assignTo,omitempty"`

	CustomizationClass string `json:"customizationClass,omitempty"`

	DefaultContributionArticleType string `json:"defaultContributionArticleType,omitempty"`

	Editor *KnowledgeCaseEditor `json:"editor,omitempty"`

	EnableArticleCreation bool `json:"enableArticleCreation,omitempty"`

	EnableArticlePublicSharingSites bool `json:"enableArticlePublicSharingSites,omitempty"`

	UseProfileForPDFCreation bool `json:"useProfileForPDFCreation,omitempty"`
}

type KnowledgeCommunitiesSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata KnowledgeCommunitiesSettings"`

	Community []string `json:"community,omitempty"`
}

type KnowledgeSitesSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata KnowledgeSitesSettings"`

	Site []string `json:"site,omitempty"`
}

type KnowledgeLanguageSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata KnowledgeLanguageSettings"`

	Language []*KnowledgeLanguage `json:"language,omitempty"`
}

type KnowledgeLanguage struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata KnowledgeLanguage"`

	Active bool `json:"active,omitempty"`

	DefaultAssignee string `json:"defaultAssignee,omitempty"`

	DefaultAssigneeType *KnowledgeLanguageLookupValueType `json:"defaultAssigneeType,omitempty"`

	DefaultReviewer string `json:"defaultReviewer,omitempty"`

	DefaultReviewerType *KnowledgeLanguageLookupValueType `json:"defaultReviewerType,omitempty"`

	Name string `json:"name,omitempty"`
}

type KnowledgeSuggestedArticlesSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata KnowledgeSuggestedArticlesSettings"`

	CaseFields *KnowledgeCaseFieldsSettings `json:"caseFields,omitempty"`

	UseSuggestedArticlesForCase bool `json:"useSuggestedArticlesForCase,omitempty"`
}

type KnowledgeCaseFieldsSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata KnowledgeCaseFieldsSettings"`

	Field []*KnowledgeCaseField `json:"field,omitempty"`
}

type KnowledgeCaseField struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata KnowledgeCaseField"`

	Name string `json:"name,omitempty"`
}

type Layout struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Layout"`

	*Metadata

	CustomButtons []string `json:"customButtons,omitempty"`

	CustomConsoleComponents *CustomConsoleComponents `json:"customConsoleComponents,omitempty"`

	EmailDefault bool `json:"emailDefault,omitempty"`

	ExcludeButtons []string `json:"excludeButtons,omitempty"`

	FeedLayout *FeedLayout `json:"feedLayout,omitempty"`

	Headers []*LayoutHeader `json:"headers,omitempty"`

	LayoutSections []*LayoutSection `json:"layoutSections,omitempty"`

	MiniLayout *MiniLayout `json:"miniLayout,omitempty"`

	MultilineLayoutFields []string `json:"multilineLayoutFields,omitempty"`

	PlatformActionList *PlatformActionList `json:"platformActionList,omitempty"`

	QuickActionList *QuickActionList `json:"quickActionList,omitempty"`

	RelatedContent *RelatedContent `json:"relatedContent,omitempty"`

	RelatedLists []*RelatedListItem `json:"relatedLists,omitempty"`

	RelatedObjects []string `json:"relatedObjects,omitempty"`

	RunAssignmentRulesDefault bool `json:"runAssignmentRulesDefault,omitempty"`

	ShowEmailCheckbox bool `json:"showEmailCheckbox,omitempty"`

	ShowHighlightsPanel bool `json:"showHighlightsPanel,omitempty"`

	ShowInteractionLogPanel bool `json:"showInteractionLogPanel,omitempty"`

	ShowKnowledgeComponent bool `json:"showKnowledgeComponent,omitempty"`

	ShowRunAssignmentRulesCheckbox bool `json:"showRunAssignmentRulesCheckbox,omitempty"`

	ShowSolutionSection bool `json:"showSolutionSection,omitempty"`

	ShowSubmitAndAttachButton bool `json:"showSubmitAndAttachButton,omitempty"`

	SummaryLayout *SummaryLayout `json:"summaryLayout,omitempty"`
}

type CustomConsoleComponents struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomConsoleComponents"`

	PrimaryTabComponents *PrimaryTabComponents `json:"primaryTabComponents,omitempty"`

	SubtabComponents *SubtabComponents `json:"subtabComponents,omitempty"`
}

type PrimaryTabComponents struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PrimaryTabComponents"`

	Containers []*Container `json:"containers,omitempty"`
}

type Container struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Container"`

	Height int32 `json:"height,omitempty"`

	IsContainerAutoSizeEnabled bool `json:"isContainerAutoSizeEnabled,omitempty"`

	Region string `json:"region,omitempty"`

	SidebarComponents []*SidebarComponent `json:"sidebarComponents,omitempty"`

	Style string `json:"style,omitempty"`

	Unit string `json:"unit,omitempty"`

	Width int32 `json:"width,omitempty"`
}

type SidebarComponent struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SidebarComponent"`

	ComponentType string `json:"componentType,omitempty"`

	Height int32 `json:"height,omitempty"`

	Label string `json:"label,omitempty"`

	Lookup string `json:"lookup,omitempty"`

	Page string `json:"page,omitempty"`

	RelatedLists []*RelatedList `json:"relatedLists,omitempty"`

	Unit string `json:"unit,omitempty"`

	Width int32 `json:"width,omitempty"`
}

type RelatedList struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RelatedList"`

	HideOnDetail bool `json:"hideOnDetail,omitempty"`

	Name string `json:"name,omitempty"`
}

type SubtabComponents struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SubtabComponents"`

	Containers []*Container `json:"containers,omitempty"`
}

type FeedLayout struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FeedLayout"`

	AutocollapsePublisher bool `json:"autocollapsePublisher,omitempty"`

	CompactFeed bool `json:"compactFeed,omitempty"`

	FeedFilterPosition *FeedLayoutFilterPosition `json:"feedFilterPosition,omitempty"`

	FeedFilters []*FeedLayoutFilter `json:"feedFilters,omitempty"`

	FullWidthFeed bool `json:"fullWidthFeed,omitempty"`

	HideSidebar bool `json:"hideSidebar,omitempty"`

	HighlightExternalFeedItems bool `json:"highlightExternalFeedItems,omitempty"`

	LeftComponents []*FeedLayoutComponent `json:"leftComponents,omitempty"`

	RightComponents []*FeedLayoutComponent `json:"rightComponents,omitempty"`

	UseInlineFiltersInConsole bool `json:"useInlineFiltersInConsole,omitempty"`
}

type FeedLayoutFilter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FeedLayoutFilter"`

	FeedFilterName string `json:"feedFilterName,omitempty"`

	FeedFilterType *FeedLayoutFilterType `json:"feedFilterType,omitempty"`

	FeedItemType *FeedItemType `json:"feedItemType,omitempty"`
}

type FeedLayoutComponent struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FeedLayoutComponent"`

	ComponentType *FeedLayoutComponentType `json:"componentType,omitempty"`

	Height int32 `json:"height,omitempty"`

	Page string `json:"page,omitempty"`
}

type LayoutSection struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LayoutSection"`

	CustomLabel bool `json:"customLabel,omitempty"`

	DetailHeading bool `json:"detailHeading,omitempty"`

	EditHeading bool `json:"editHeading,omitempty"`

	Label string `json:"label,omitempty"`

	LayoutColumns []*LayoutColumn `json:"layoutColumns,omitempty"`

	Style *LayoutSectionStyle `json:"style,omitempty"`
}

type LayoutColumn struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LayoutColumn"`

	LayoutItems []*LayoutItem `json:"layoutItems,omitempty"`

	Reserved string `json:"reserved,omitempty"`
}

type LayoutItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LayoutItem"`

	AnalyticsCloudComponent *AnalyticsCloudComponentLayoutItem `json:"analyticsCloudComponent,omitempty"`

	Behavior *UiBehavior `json:"behavior,omitempty"`

	Canvas string `json:"canvas,omitempty"`

	Component string `json:"component,omitempty"`

	CustomLink string `json:"customLink,omitempty"`

	EmptySpace bool `json:"emptySpace,omitempty"`

	Field string `json:"field,omitempty"`

	Height int32 `json:"height,omitempty"`

	Page string `json:"page,omitempty"`

	ReportChartComponent *ReportChartComponentLayoutItem `json:"reportChartComponent,omitempty"`

	Scontrol string `json:"scontrol,omitempty"`

	ShowLabel bool `json:"showLabel,omitempty"`

	ShowScrollbars bool `json:"showScrollbars,omitempty"`

	Width string `json:"width,omitempty"`
}

type AnalyticsCloudComponentLayoutItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AnalyticsCloudComponentLayoutItem"`

	AssetType string `json:"assetType,omitempty"`

	DevName string `json:"devName,omitempty"`

	Error string `json:"error,omitempty"`

	Filter string `json:"filter,omitempty"`

	Height int32 `json:"height,omitempty"`

	HideOnError bool `json:"hideOnError,omitempty"`

	ShowSharing bool `json:"showSharing,omitempty"`

	ShowTitle bool `json:"showTitle,omitempty"`

	Width string `json:"width,omitempty"`
}

type ReportChartComponentLayoutItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportChartComponentLayoutItem"`

	CacheData bool `json:"cacheData,omitempty"`

	ContextFilterableField string `json:"contextFilterableField,omitempty"`

	Error string `json:"error,omitempty"`

	HideOnError bool `json:"hideOnError,omitempty"`

	IncludeContext bool `json:"includeContext,omitempty"`

	ReportName string `json:"reportName,omitempty"`

	ShowTitle bool `json:"showTitle,omitempty"`

	Size *ReportChartComponentSize `json:"size,omitempty"`
}

type MiniLayout struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata MiniLayout"`

	Fields []string `json:"fields,omitempty"`

	RelatedLists []*RelatedListItem `json:"relatedLists,omitempty"`
}

type RelatedListItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RelatedListItem"`

	CustomButtons []string `json:"customButtons,omitempty"`

	ExcludeButtons []string `json:"excludeButtons,omitempty"`

	Fields []string `json:"fields,omitempty"`

	RelatedList string `json:"relatedList,omitempty"`

	SortField string `json:"sortField,omitempty"`

	SortOrder *SortOrder `json:"sortOrder,omitempty"`
}

type RelatedContent struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RelatedContent"`

	RelatedContentItems []*RelatedContentItem `json:"relatedContentItems,omitempty"`
}

type RelatedContentItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RelatedContentItem"`

	LayoutItem *LayoutItem `json:"layoutItem,omitempty"`
}

type SummaryLayout struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SummaryLayout"`

	MasterLabel string `json:"masterLabel,omitempty"`

	SizeX int32 `json:"sizeX,omitempty"`

	SizeY int32 `json:"sizeY,omitempty"`

	SizeZ int32 `json:"sizeZ,omitempty"`

	SummaryLayoutItems []*SummaryLayoutItem `json:"summaryLayoutItems,omitempty"`

	SummaryLayoutStyle *SummaryLayoutStyle `json:"summaryLayoutStyle,omitempty"`
}

type SummaryLayoutItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SummaryLayoutItem"`

	CustomLink string `json:"customLink,omitempty"`

	Field string `json:"field,omitempty"`

	PosX int32 `json:"posX,omitempty"`

	PosY int32 `json:"posY,omitempty"`

	PosZ int32 `json:"posZ,omitempty"`
}

type LeadConvertSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LeadConvertSettings"`

	*Metadata

	AllowOwnerChange bool `json:"allowOwnerChange,omitempty"`

	OpportunityCreationOptions *VisibleOrRequired `json:"opportunityCreationOptions,omitempty"`
}

type Letterhead struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Letterhead"`

	*Metadata

	Available bool `json:"available,omitempty"`

	BackgroundColor string `json:"backgroundColor,omitempty"`

	BodyColor string `json:"bodyColor,omitempty"`

	BottomLine *LetterheadLine `json:"bottomLine,omitempty"`

	Description string `json:"description,omitempty"`

	Footer *LetterheadHeaderFooter `json:"footer,omitempty"`

	Header *LetterheadHeaderFooter `json:"header,omitempty"`

	MiddleLine *LetterheadLine `json:"middleLine,omitempty"`

	Name string `json:"name,omitempty"`

	TopLine *LetterheadLine `json:"topLine,omitempty"`
}

type LetterheadLine struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LetterheadLine"`

	Color string `json:"color,omitempty"`

	Height int32 `json:"height,omitempty"`
}

type LetterheadHeaderFooter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LetterheadHeaderFooter"`

	BackgroundColor string `json:"backgroundColor,omitempty"`

	Height int32 `json:"height,omitempty"`

	HorizontalAlignment *LetterheadHorizontalAlignment `json:"horizontalAlignment,omitempty"`

	Logo string `json:"logo,omitempty"`

	VerticalAlignment *LetterheadVerticalAlignment `json:"verticalAlignment,omitempty"`
}

type LicenseDefinition struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LicenseDefinition"`

	*Metadata

	AggregationGroup string `json:"aggregationGroup,omitempty"`

	Description string `json:"description,omitempty"`

	IsPublished bool `json:"isPublished,omitempty"`

	Label string `json:"label,omitempty"`

	LicensedCustomPermissions []*LicensedCustomPermissions `json:"licensedCustomPermissions,omitempty"`

	LicensingAuthority string `json:"licensingAuthority,omitempty"`

	LicensingAuthorityProvider string `json:"licensingAuthorityProvider,omitempty"`

	MinPlatformVersion int32 `json:"minPlatformVersion,omitempty"`

	Origin string `json:"origin,omitempty"`

	Revision int32 `json:"revision,omitempty"`

	TrialLicenseDuration int32 `json:"trialLicenseDuration,omitempty"`

	TrialLicenseQuantity int32 `json:"trialLicenseQuantity,omitempty"`
}

type LicensedCustomPermissions struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LicensedCustomPermissions"`

	CustomPermission string `json:"customPermission,omitempty"`

	LicenseDefinition string `json:"licenseDefinition,omitempty"`
}

type LiveAgentSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LiveAgentSettings"`

	*Metadata

	EnableLiveAgent bool `json:"enableLiveAgent,omitempty"`
}

type LiveChatAgentConfig struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LiveChatAgentConfig"`

	*Metadata

	Assignments *AgentConfigAssignments `json:"assignments,omitempty"`

	AutoGreeting string `json:"autoGreeting,omitempty"`

	Capacity int32 `json:"capacity,omitempty"`

	CriticalWaitTime int32 `json:"criticalWaitTime,omitempty"`

	CustomAgentName string `json:"customAgentName,omitempty"`

	EnableAgentFileTransfer bool `json:"enableAgentFileTransfer,omitempty"`

	EnableAgentSneakPeek bool `json:"enableAgentSneakPeek,omitempty"`

	EnableAssistanceFlag bool `json:"enableAssistanceFlag,omitempty"`

	EnableAutoAwayOnDecline bool `json:"enableAutoAwayOnDecline,omitempty"`

	EnableAutoAwayOnPushTimeout bool `json:"enableAutoAwayOnPushTimeout,omitempty"`

	EnableChatConferencing bool `json:"enableChatConferencing,omitempty"`

	EnableChatMonitoring bool `json:"enableChatMonitoring,omitempty"`

	EnableChatTransferToAgent bool `json:"enableChatTransferToAgent,omitempty"`

	EnableChatTransferToButton bool `json:"enableChatTransferToButton,omitempty"`

	EnableChatTransferToSkill bool `json:"enableChatTransferToSkill,omitempty"`

	EnableLogoutSound bool `json:"enableLogoutSound,omitempty"`

	EnableNotifications bool `json:"enableNotifications,omitempty"`

	EnableRequestSound bool `json:"enableRequestSound,omitempty"`

	EnableSneakPeek bool `json:"enableSneakPeek,omitempty"`

	EnableVisitorBlocking bool `json:"enableVisitorBlocking,omitempty"`

	EnableWhisperMessage bool `json:"enableWhisperMessage,omitempty"`

	Label string `json:"label,omitempty"`

	SupervisorDefaultAgentStatusFilter *SupervisorAgentStatusFilter `json:"supervisorDefaultAgentStatusFilter,omitempty"`

	SupervisorDefaultButtonFilter string `json:"supervisorDefaultButtonFilter,omitempty"`

	SupervisorDefaultSkillFilter string `json:"supervisorDefaultSkillFilter,omitempty"`

	SupervisorSkills *SupervisorAgentConfigSkills `json:"supervisorSkills,omitempty"`

	TransferableButtons *AgentConfigButtons `json:"transferableButtons,omitempty"`

	TransferableSkills *AgentConfigSkills `json:"transferableSkills,omitempty"`
}

type AgentConfigAssignments struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AgentConfigAssignments"`

	Profiles *AgentConfigProfileAssignments `json:"profiles,omitempty"`

	Users *AgentConfigUserAssignments `json:"users,omitempty"`
}

type AgentConfigProfileAssignments struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AgentConfigProfileAssignments"`

	Profile []string `json:"profile,omitempty"`
}

type AgentConfigUserAssignments struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AgentConfigUserAssignments"`

	User []string `json:"user,omitempty"`
}

type SupervisorAgentConfigSkills struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SupervisorAgentConfigSkills"`

	Skill []string `json:"skill,omitempty"`
}

type AgentConfigButtons struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AgentConfigButtons"`

	Button []string `json:"button,omitempty"`
}

type AgentConfigSkills struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AgentConfigSkills"`

	Skill []string `json:"skill,omitempty"`
}

type LiveChatButton struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LiveChatButton"`

	*Metadata

	Animation *LiveChatButtonPresentation `json:"animation,omitempty"`

	AutoGreeting string `json:"autoGreeting,omitempty"`

	ChasitorIdleTimeout int32 `json:"chasitorIdleTimeout,omitempty"`

	ChasitorIdleTimeoutWarning int32 `json:"chasitorIdleTimeoutWarning,omitempty"`

	ChatPage string `json:"chatPage,omitempty"`

	CustomAgentName string `json:"customAgentName,omitempty"`

	Deployments *LiveChatButtonDeployments `json:"deployments,omitempty"`

	EnableQueue bool `json:"enableQueue,omitempty"`

	InviteEndPosition *LiveChatButtonInviteEndPosition `json:"inviteEndPosition,omitempty"`

	InviteImage string `json:"inviteImage,omitempty"`

	InviteStartPosition *LiveChatButtonInviteStartPosition `json:"inviteStartPosition,omitempty"`

	IsActive bool `json:"isActive,omitempty"`

	Label string `json:"label,omitempty"`

	NumberOfReroutingAttempts int32 `json:"numberOfReroutingAttempts,omitempty"`

	OfflineImage string `json:"offlineImage,omitempty"`

	OnlineImage string `json:"onlineImage,omitempty"`

	OptionsCustomRoutingIsEnabled bool `json:"optionsCustomRoutingIsEnabled,omitempty"`

	OptionsHasChasitorIdleTimeout bool `json:"optionsHasChasitorIdleTimeout,omitempty"`

	OptionsHasInviteAfterAccept bool `json:"optionsHasInviteAfterAccept,omitempty"`

	OptionsHasInviteAfterReject bool `json:"optionsHasInviteAfterReject,omitempty"`

	OptionsHasRerouteDeclinedRequest bool `json:"optionsHasRerouteDeclinedRequest,omitempty"`

	OptionsIsAutoAccept bool `json:"optionsIsAutoAccept,omitempty"`

	OptionsIsInviteAutoRemove bool `json:"optionsIsInviteAutoRemove,omitempty"`

	OverallQueueLength int32 `json:"overallQueueLength,omitempty"`

	PerAgentQueueLength int32 `json:"perAgentQueueLength,omitempty"`

	PostChatPage string `json:"postChatPage,omitempty"`

	PostChatUrl string `json:"postChatUrl,omitempty"`

	PreChatFormPage string `json:"preChatFormPage,omitempty"`

	PreChatFormUrl string `json:"preChatFormUrl,omitempty"`

	PushTimeOut int32 `json:"pushTimeOut,omitempty"`

	RoutingType *LiveChatButtonRoutingType `json:"routingType,omitempty"`

	Site string `json:"site,omitempty"`

	Skills *LiveChatButtonSkills `json:"skills,omitempty"`

	TimeToRemoveInvite int32 `json:"timeToRemoveInvite,omitempty"`

	Type_ *LiveChatButtonType `json:"type,omitempty"`

	WindowLanguage *Language `json:"windowLanguage,omitempty"`
}

type LiveChatButtonDeployments struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LiveChatButtonDeployments"`

	Deployment []string `json:"deployment,omitempty"`
}

type LiveChatButtonSkills struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LiveChatButtonSkills"`

	Skill []string `json:"skill,omitempty"`
}

type LiveChatDeployment struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LiveChatDeployment"`

	*Metadata

	BrandingImage string `json:"brandingImage,omitempty"`

	ConnectionTimeoutDuration int32 `json:"connectionTimeoutDuration,omitempty"`

	ConnectionWarningDuration int32 `json:"connectionWarningDuration,omitempty"`

	DisplayQueuePosition bool `json:"displayQueuePosition,omitempty"`

	DomainWhiteList *LiveChatDeploymentDomainWhitelist `json:"domainWhiteList,omitempty"`

	EnablePrechatApi bool `json:"enablePrechatApi,omitempty"`

	EnableTranscriptSave bool `json:"enableTranscriptSave,omitempty"`

	Label string `json:"label,omitempty"`

	MobileBrandingImage string `json:"mobileBrandingImage,omitempty"`

	Site string `json:"site,omitempty"`

	WindowTitle string `json:"windowTitle,omitempty"`
}

type LiveChatDeploymentDomainWhitelist struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LiveChatDeploymentDomainWhitelist"`

	Domain []string `json:"domain,omitempty"`
}

type LiveChatSensitiveDataRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LiveChatSensitiveDataRule"`

	*Metadata

	ActionType *SensitiveDataActionType `json:"actionType,omitempty"`

	Description string `json:"description,omitempty"`

	EnforceOn int32 `json:"enforceOn,omitempty"`

	IsEnabled bool `json:"isEnabled,omitempty"`

	Pattern string `json:"pattern,omitempty"`

	Replacement string `json:"replacement,omitempty"`
}

type ManagedTopic struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ManagedTopic"`

	*Metadata

	ManagedTopicType string `json:"managedTopicType,omitempty"`

	Name string `json:"name,omitempty"`

	ParentName string `json:"parentName,omitempty"`

	Position int32 `json:"position,omitempty"`

	TopicDescription string `json:"topicDescription,omitempty"`
}

type ManagedTopics struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ManagedTopics"`

	*Metadata

	ManagedTopic []*ManagedTopic `json:"managedTopic,omitempty"`
}

type MarketingActionSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata MarketingActionSettings"`

	*Metadata

	EnableMarketingAction bool `json:"enableMarketingAction,omitempty"`
}

type MarketingResourceType struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata MarketingResourceType"`

	*Metadata

	Description string `json:"description,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	Object string `json:"object,omitempty"`

	Provider string `json:"provider,omitempty"`
}

type MatchingRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata MatchingRule"`

	*Metadata

	BooleanFilter string `json:"booleanFilter,omitempty"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	MatchingRuleItems []*MatchingRuleItem `json:"matchingRuleItems,omitempty"`

	RuleStatus *MatchingRuleStatus `json:"ruleStatus,omitempty"`
}

type MatchingRuleItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata MatchingRuleItem"`

	BlankValueBehavior *BlankValueBehavior `json:"blankValueBehavior,omitempty"`

	FieldName string `json:"fieldName,omitempty"`

	MatchingMethod *MatchingMethod `json:"matchingMethod,omitempty"`
}

type MatchingRules struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata MatchingRules"`

	*Metadata

	MatchingRules []*MatchingRule `json:"matchingRules,omitempty"`
}

type MetadataWithContent struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata MetadataWithContent"`

	*Metadata

	Content []byte `json:"content,omitempty"`
}

type ApexClass struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ApexClass"`

	*MetadataWithContent

	ApiVersion float64 `json:"apiVersion,omitempty"`

	PackageVersions []*PackageVersion `json:"packageVersions,omitempty"`

	Status *ApexCodeUnitStatus `json:"status,omitempty"`
}

type ApexComponent struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ApexComponent"`

	*MetadataWithContent

	ApiVersion float64 `json:"apiVersion,omitempty"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	PackageVersions []*PackageVersion `json:"packageVersions,omitempty"`
}

type ApexPage struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ApexPage"`

	*MetadataWithContent

	ApiVersion float64 `json:"apiVersion,omitempty"`

	AvailableInTouch bool `json:"availableInTouch,omitempty"`

	ConfirmationTokenRequired bool `json:"confirmationTokenRequired,omitempty"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	PackageVersions []*PackageVersion `json:"packageVersions,omitempty"`
}

type ApexTrigger struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ApexTrigger"`

	*MetadataWithContent

	ApiVersion float64 `json:"apiVersion,omitempty"`

	PackageVersions []*PackageVersion `json:"packageVersions,omitempty"`

	Status *ApexCodeUnitStatus `json:"status,omitempty"`
}

type Certificate struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Certificate"`

	*MetadataWithContent

	CaSigned bool `json:"caSigned,omitempty"`

	EncryptedWithPlatformEncryption bool `json:"encryptedWithPlatformEncryption,omitempty"`

	ExpirationDate time.Time `json:"expirationDate,omitempty"`

	KeySize int32 `json:"keySize,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	PrivateKeyExportable bool `json:"privateKeyExportable,omitempty"`
}

type ContentAsset struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ContentAsset"`

	*MetadataWithContent

	Format *ContentAssetFormat `json:"format,omitempty"`

	Language string `json:"language,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	OriginNetwork string `json:"originNetwork,omitempty"`

	Relationships *ContentAssetRelationships `json:"relationships,omitempty"`

	Versions *ContentAssetVersions `json:"versions,omitempty"`
}

type ContentAssetRelationships struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ContentAssetRelationships"`

	Organization *ContentAssetLink `json:"organization,omitempty"`
}

type ContentAssetLink struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ContentAssetLink"`

	Access *ContentAssetAccess `json:"access,omitempty"`

	Name string `json:"name,omitempty"`
}

type ContentAssetVersions struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ContentAssetVersions"`

	Version []*ContentAssetVersion `json:"version,omitempty"`
}

type ContentAssetVersion struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ContentAssetVersion"`

	Number string `json:"number,omitempty"`

	PathOnClient string `json:"pathOnClient,omitempty"`

	ZipEntry string `json:"zipEntry,omitempty"`
}

type DataPipeline struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DataPipeline"`

	*MetadataWithContent

	ApiVersion float64 `json:"apiVersion,omitempty"`

	Label string `json:"label,omitempty"`

	ScriptType *DataPipelineType `json:"scriptType,omitempty"`
}

type Document struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Document"`

	*MetadataWithContent

	Description string `json:"description,omitempty"`

	InternalUseOnly bool `json:"internalUseOnly,omitempty"`

	Keywords string `json:"keywords,omitempty"`

	Name string `json:"name,omitempty"`

	Public bool `json:"public,omitempty"`
}

type EmailTemplate struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata EmailTemplate"`

	*MetadataWithContent

	ApiVersion float64 `json:"apiVersion,omitempty"`

	AttachedDocuments []string `json:"attachedDocuments,omitempty"`

	Attachments []*Attachment `json:"attachments,omitempty"`

	Available bool `json:"available,omitempty"`

	Description string `json:"description,omitempty"`

	EncodingKey *Encoding `json:"encodingKey,omitempty"`

	Letterhead string `json:"letterhead,omitempty"`

	Name string `json:"name,omitempty"`

	PackageVersions []*PackageVersion `json:"packageVersions,omitempty"`

	Style *EmailTemplateStyle `json:"style,omitempty"`

	Subject string `json:"subject,omitempty"`

	TextOnly string `json:"textOnly,omitempty"`

	Type_ *EmailTemplateType `json:"type,omitempty"`
}

type Attachment struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Attachment"`

	Content []byte `json:"content,omitempty"`

	Name string `json:"name,omitempty"`
}

type Scontrol struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Scontrol"`

	*MetadataWithContent

	ContentSource *SControlContentSource `json:"contentSource,omitempty"`

	Description string `json:"description,omitempty"`

	EncodingKey *Encoding `json:"encodingKey,omitempty"`

	FileContent []byte `json:"fileContent,omitempty"`

	FileName string `json:"fileName,omitempty"`

	Name string `json:"name,omitempty"`

	SupportsCaching bool `json:"supportsCaching,omitempty"`
}

type SiteDotCom struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SiteDotCom"`

	*MetadataWithContent

	Label string `json:"label,omitempty"`

	SiteType *SiteType `json:"siteType,omitempty"`
}

type StaticResource struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata StaticResource"`

	*MetadataWithContent

	CacheControl *StaticResourceCacheControl `json:"cacheControl,omitempty"`

	ContentType string `json:"contentType,omitempty"`

	Description string `json:"description,omitempty"`
}

type UiPlugin struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata UiPlugin"`

	*MetadataWithContent

	Description string `json:"description,omitempty"`

	ExtensionPointIdentifier string `json:"extensionPointIdentifier,omitempty"`

	IsEnabled bool `json:"isEnabled,omitempty"`

	Language string `json:"language,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`
}

type WaveDashboard struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WaveDashboard"`

	*MetadataWithContent

	Application string `json:"application,omitempty"`

	Description string `json:"description,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	TemplateAssetSourceName string `json:"templateAssetSourceName,omitempty"`
}

type WaveDataflow struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WaveDataflow"`

	*MetadataWithContent

	Description string `json:"description,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`
}

type WaveLens struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WaveLens"`

	*MetadataWithContent

	Application string `json:"application,omitempty"`

	Datasets []string `json:"datasets,omitempty"`

	Description string `json:"description,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	TemplateAssetSourceName string `json:"templateAssetSourceName,omitempty"`

	VisualizationType string `json:"visualizationType,omitempty"`
}

type MilestoneType struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata MilestoneType"`

	*Metadata

	Description string `json:"description,omitempty"`

	RecurrenceType *MilestoneTypeRecurrenceType `json:"recurrenceType,omitempty"`
}

type MobileSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata MobileSettings"`

	*Metadata

	ChatterMobile *ChatterMobileSettings `json:"chatterMobile,omitempty"`

	DashboardMobile *DashboardMobileSettings `json:"dashboardMobile,omitempty"`

	ObjforceMobile *SFDCMobileSettings `json:"objforceMobile,omitempty"`

	TouchMobile *TouchMobileSettings `json:"touchMobile,omitempty"`
}

type ChatterMobileSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ChatterMobileSettings"`

	EnablePushNotifications bool `json:"enablePushNotifications,omitempty"`
}

type DashboardMobileSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DashboardMobileSettings"`

	EnableDashboardIPadApp bool `json:"enableDashboardIPadApp,omitempty"`
}

type SFDCMobileSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SFDCMobileSettings"`

	EnableMobileLite bool `json:"enableMobileLite,omitempty"`

	EnableUserToDeviceLinking bool `json:"enableUserToDeviceLinking,omitempty"`
}

type TouchMobileSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata TouchMobileSettings"`

	EnableTouchAppIPad bool `json:"enableTouchAppIPad,omitempty"`

	EnableTouchAppIPhone bool `json:"enableTouchAppIPhone,omitempty"`

	EnableTouchBrowserIPad bool `json:"enableTouchBrowserIPad,omitempty"`

	EnableTouchIosPhone bool `json:"enableTouchIosPhone,omitempty"`

	EnableVisualforceInTouch bool `json:"enableVisualforceInTouch,omitempty"`
}

type ModerationRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ModerationRule"`

	*Metadata

	Action *ModerationRuleAction `json:"action,omitempty"`

	Active bool `json:"active,omitempty"`

	Description string `json:"description,omitempty"`

	EntitiesAndFields []*ModeratedEntityField `json:"entitiesAndFields,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	UserMessage string `json:"userMessage,omitempty"`
}

type ModeratedEntityField struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ModeratedEntityField"`

	EntityName string `json:"entityName,omitempty"`

	FieldName string `json:"fieldName,omitempty"`

	KeywordList string `json:"keywordList,omitempty"`
}

type NameSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata NameSettings"`

	*Metadata

	EnableMiddleName bool `json:"enableMiddleName,omitempty"`

	EnableNameSuffix bool `json:"enableNameSuffix,omitempty"`
}

type NamedCredential struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata NamedCredential"`

	*Metadata

	AuthProvider string `json:"authProvider,omitempty"`

	Certificate string `json:"certificate,omitempty"`

	Endpoint string `json:"endpoint,omitempty"`

	Label string `json:"label,omitempty"`

	OauthRefreshToken string `json:"oauthRefreshToken,omitempty"`

	OauthScope string `json:"oauthScope,omitempty"`

	OauthToken string `json:"oauthToken,omitempty"`

	Password string `json:"password,omitempty"`

	PrincipalType *ExternalPrincipalType `json:"principalType,omitempty"`

	Protocol *AuthenticationProtocol `json:"protocol,omitempty"`

	Username string `json:"username,omitempty"`
}

type Network struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Network"`

	*Metadata

	AllowMembersToFlag bool `json:"allowMembersToFlag,omitempty"`

	AllowedExtensions string `json:"allowedExtensions,omitempty"`

	Branding *Branding `json:"branding,omitempty"`

	CaseCommentEmailTemplate string `json:"caseCommentEmailTemplate,omitempty"`

	ChangePasswordTemplate string `json:"changePasswordTemplate,omitempty"`

	Description string `json:"description,omitempty"`

	EmailSenderAddress string `json:"emailSenderAddress,omitempty"`

	EmailSenderName string `json:"emailSenderName,omitempty"`

	EnableGuestChatter bool `json:"enableGuestChatter,omitempty"`

	EnableInvitation bool `json:"enableInvitation,omitempty"`

	EnableKnowledgeable bool `json:"enableKnowledgeable,omitempty"`

	EnableNicknameDisplay bool `json:"enableNicknameDisplay,omitempty"`

	EnablePrivateMessages bool `json:"enablePrivateMessages,omitempty"`

	EnableReputation bool `json:"enableReputation,omitempty"`

	EnableSiteAsContainer bool `json:"enableSiteAsContainer,omitempty"`

	FeedChannel string `json:"feedChannel,omitempty"`

	ForgotPasswordTemplate string `json:"forgotPasswordTemplate,omitempty"`

	LogoutUrl string `json:"logoutUrl,omitempty"`

	MaxFileSizeKb int32 `json:"maxFileSizeKb,omitempty"`

	NavigationLinkSet *NavigationLinkSet `json:"navigationLinkSet,omitempty"`

	NetworkMemberGroups *NetworkMemberGroup `json:"networkMemberGroups,omitempty"`

	NewSenderAddress string `json:"newSenderAddress,omitempty"`

	PicassoSite string `json:"picassoSite,omitempty"`

	ReputationLevels *ReputationLevelDefinitions `json:"reputationLevels,omitempty"`

	ReputationPointsRules *ReputationPointsRules `json:"reputationPointsRules,omitempty"`

	SelfRegProfile string `json:"selfRegProfile,omitempty"`

	SelfRegistration bool `json:"selfRegistration,omitempty"`

	SendWelcomeEmail bool `json:"sendWelcomeEmail,omitempty"`

	Site string `json:"site,omitempty"`

	Status *NetworkStatus `json:"status,omitempty"`

	Tabs *NetworkTabSet `json:"tabs,omitempty"`

	UrlPathPrefix string `json:"urlPathPrefix,omitempty"`

	WelcomeTemplate string `json:"welcomeTemplate,omitempty"`
}

type Branding struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Branding"`

	LoginFooterText string `json:"loginFooterText,omitempty"`

	LoginLogo string `json:"loginLogo,omitempty"`

	PageFooter string `json:"pageFooter,omitempty"`

	PageHeader string `json:"pageHeader,omitempty"`

	PrimaryColor string `json:"primaryColor,omitempty"`

	PrimaryComplementColor string `json:"primaryComplementColor,omitempty"`

	QuaternaryColor string `json:"quaternaryColor,omitempty"`

	QuaternaryComplementColor string `json:"quaternaryComplementColor,omitempty"`

	SecondaryColor string `json:"secondaryColor,omitempty"`

	TertiaryColor string `json:"tertiaryColor,omitempty"`

	TertiaryComplementColor string `json:"tertiaryComplementColor,omitempty"`

	ZeronaryColor string `json:"zeronaryColor,omitempty"`

	ZeronaryComplementColor string `json:"zeronaryComplementColor,omitempty"`
}

type NavigationLinkSet struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata NavigationLinkSet"`

	NavigationMenuItem []*NavigationMenuItem `json:"navigationMenuItem,omitempty"`
}

type NavigationMenuItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata NavigationMenuItem"`

	DefaultListViewId string `json:"defaultListViewId,omitempty"`

	Label string `json:"label,omitempty"`

	Position int32 `json:"position,omitempty"`

	PubliclyAvailable bool `json:"publiclyAvailable,omitempty"`

	Target string `json:"target,omitempty"`

	TargetPreference string `json:"targetPreference,omitempty"`

	Type_ string `json:"type,omitempty"`
}

type NetworkMemberGroup struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata NetworkMemberGroup"`

	PermissionSet []string `json:"permissionSet,omitempty"`

	Profile []string `json:"profile,omitempty"`
}

type ReputationLevelDefinitions struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReputationLevelDefinitions"`

	Level []*ReputationLevel `json:"level,omitempty"`
}

type ReputationLevel struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReputationLevel"`

	Branding *ReputationBranding `json:"branding,omitempty"`

	Label string `json:"label,omitempty"`

	LowerThreshold float64 `json:"lowerThreshold,omitempty"`
}

type ReputationBranding struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReputationBranding"`

	SmallImage string `json:"smallImage,omitempty"`
}

type ReputationPointsRules struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReputationPointsRules"`

	PointsRule []*ReputationPointsRule `json:"pointsRule,omitempty"`
}

type ReputationPointsRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReputationPointsRule"`

	EventType string `json:"eventType,omitempty"`

	Points int32 `json:"points,omitempty"`
}

type NetworkTabSet struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata NetworkTabSet"`

	CustomTab []string `json:"customTab,omitempty"`

	DefaultTab string `json:"defaultTab,omitempty"`

	StandardTab []string `json:"standardTab,omitempty"`
}

type OpportunitySettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata OpportunitySettings"`

	*Metadata

	AutoActivateNewReminders bool `json:"autoActivateNewReminders,omitempty"`

	EnableFindSimilarOpportunities bool `json:"enableFindSimilarOpportunities,omitempty"`

	EnableOpportunityTeam bool `json:"enableOpportunityTeam,omitempty"`

	EnableUpdateReminders bool `json:"enableUpdateReminders,omitempty"`

	FindSimilarOppFilter *FindSimilarOppFilter `json:"findSimilarOppFilter,omitempty"`

	PromptToAddProducts bool `json:"promptToAddProducts,omitempty"`
}

type FindSimilarOppFilter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FindSimilarOppFilter"`

	SimilarOpportunitiesDisplayColumns []string `json:"similarOpportunitiesDisplayColumns,omitempty"`

	SimilarOpportunitiesMatchFields []string `json:"similarOpportunitiesMatchFields,omitempty"`
}

type OrderSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata OrderSettings"`

	*Metadata

	EnableNegativeQuantity bool `json:"enableNegativeQuantity,omitempty"`

	EnableOrders bool `json:"enableOrders,omitempty"`

	EnableReductionOrders bool `json:"enableReductionOrders,omitempty"`
}

type OrgPreferenceSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata OrgPreferenceSettings"`

	*Metadata

	Preferences []*OrganizationSettingsDetail `json:"preferences,omitempty"`
}

type OrganizationSettingsDetail struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata OrganizationSettingsDetail"`

	SettingName string `json:"settingName,omitempty"`

	SettingValue bool `json:"settingValue,omitempty"`
}

type Package struct {
	*Metadata

	ApiAccessLevel *APIAccessLevel `json:"apiAccessLevel,omitempty"`

	Description string `json:"description,omitempty"`

	NamespacePrefix string `json:"namespacePrefix,omitempty"`

	ObjectPermissions []*ProfileObjectPermissions `json:"objectPermissions,omitempty"`

	PackageType string `json:"packageType,omitempty"`

	PostInstallClass string `json:"postInstallClass,omitempty"`

	SetupWeblink string `json:"setupWeblink,omitempty"`

	Types []*PackageTypeMembers `json:"types,omitempty"`

	UninstallClass string `json:"uninstallClass,omitempty"`

	Version string `json:"version,omitempty"`
}

type ProfileObjectPermissions struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileObjectPermissions"`

	AllowCreate bool `json:"allowCreate,omitempty"`

	AllowDelete bool `json:"allowDelete,omitempty"`

	AllowEdit bool `json:"allowEdit,omitempty"`

	AllowRead bool `json:"allowRead,omitempty"`

	ModifyAllRecords bool `json:"modifyAllRecords,omitempty"`

	Object string `json:"object,omitempty"`

	ViewAllRecords bool `json:"viewAllRecords,omitempty"`
}

type PackageTypeMembers struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PackageTypeMembers"`

	Members []string `json:"members,omitempty"`

	Name string `json:"name,omitempty"`
}

type PathAssistant struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PathAssistant"`

	*Metadata

	Active bool `json:"active,omitempty"`

	EntityName string `json:"entityName,omitempty"`

	FieldName string `json:"fieldName,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	PathAssistantSteps []*PathAssistantStep `json:"pathAssistantSteps,omitempty"`

	RecordTypeName string `json:"recordTypeName,omitempty"`
}

type PathAssistantStep struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PathAssistantStep"`

	FieldNames []string `json:"fieldNames,omitempty"`

	Info string `json:"info,omitempty"`

	PicklistValueName string `json:"picklistValueName,omitempty"`
}

type PathAssistantSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PathAssistantSettings"`

	*Metadata

	PathAssistantEnabled bool `json:"pathAssistantEnabled,omitempty"`
}

type PermissionSet struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PermissionSet"`

	*Metadata

	ApplicationVisibilities []*PermissionSetApplicationVisibility `json:"applicationVisibilities,omitempty"`

	ClassAccesses []*PermissionSetApexClassAccess `json:"classAccesses,omitempty"`

	CustomPermissions []*PermissionSetCustomPermissions `json:"customPermissions,omitempty"`

	Description string `json:"description,omitempty"`

	ExternalDataSourceAccesses []*PermissionSetExternalDataSourceAccess `json:"externalDataSourceAccesses,omitempty"`

	FieldPermissions []*PermissionSetFieldPermissions `json:"fieldPermissions,omitempty"`

	HasActivationRequired bool `json:"hasActivationRequired,omitempty"`

	Label string `json:"label,omitempty"`

	License string `json:"license,omitempty"`

	ObjectPermissions []*PermissionSetObjectPermissions `json:"objectPermissions,omitempty"`

	PageAccesses []*PermissionSetApexPageAccess `json:"pageAccesses,omitempty"`

	RecordTypeVisibilities []*PermissionSetRecordTypeVisibility `json:"recordTypeVisibilities,omitempty"`

	TabSettings []*PermissionSetTabSetting `json:"tabSettings,omitempty"`

	UserPermissions []*PermissionSetUserPermission `json:"userPermissions,omitempty"`
}

type PermissionSetApplicationVisibility struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PermissionSetApplicationVisibility"`

	Application string `json:"application,omitempty"`

	Visible bool `json:"visible,omitempty"`
}

type PermissionSetApexClassAccess struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PermissionSetApexClassAccess"`

	ApexClass string `json:"apexClass,omitempty"`

	Enabled bool `json:"enabled,omitempty"`
}

type PermissionSetCustomPermissions struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PermissionSetCustomPermissions"`

	Enabled bool `json:"enabled,omitempty"`

	Name string `json:"name,omitempty"`
}

type PermissionSetExternalDataSourceAccess struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PermissionSetExternalDataSourceAccess"`

	Enabled bool `json:"enabled,omitempty"`

	ExternalDataSource string `json:"externalDataSource,omitempty"`
}

type PermissionSetFieldPermissions struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PermissionSetFieldPermissions"`

	Editable bool `json:"editable,omitempty"`

	Field string `json:"field,omitempty"`

	Readable bool `json:"readable,omitempty"`
}

type PermissionSetObjectPermissions struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PermissionSetObjectPermissions"`

	AllowCreate bool `json:"allowCreate,omitempty"`

	AllowDelete bool `json:"allowDelete,omitempty"`

	AllowEdit bool `json:"allowEdit,omitempty"`

	AllowRead bool `json:"allowRead,omitempty"`

	ModifyAllRecords bool `json:"modifyAllRecords,omitempty"`

	Object string `json:"object,omitempty"`

	ViewAllRecords bool `json:"viewAllRecords,omitempty"`
}

type PermissionSetApexPageAccess struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PermissionSetApexPageAccess"`

	ApexPage string `json:"apexPage,omitempty"`

	Enabled bool `json:"enabled,omitempty"`
}

type PermissionSetRecordTypeVisibility struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PermissionSetRecordTypeVisibility"`

	RecordType string `json:"recordType,omitempty"`

	Visible bool `json:"visible,omitempty"`
}

type PermissionSetTabSetting struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PermissionSetTabSetting"`

	Tab string `json:"tab,omitempty"`

	Visibility *PermissionSetTabVisibility `json:"visibility,omitempty"`
}

type PermissionSetUserPermission struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PermissionSetUserPermission"`

	Enabled bool `json:"enabled,omitempty"`

	Name string `json:"name,omitempty"`
}

type PersonListSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PersonListSettings"`

	*Metadata

	EnablePersonList bool `json:"enablePersonList,omitempty"`
}

type PersonalJourneySettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PersonalJourneySettings"`

	*Metadata

	EnableExactTargetForObjforceApps bool `json:"enableExactTargetForObjforceApps,omitempty"`
}

type PlatformCachePartition struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PlatformCachePartition"`

	*Metadata

	Description string `json:"description,omitempty"`

	IsDefaultPartition bool `json:"isDefaultPartition,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	PlatformCachePartitionTypes []*PlatformCachePartitionType `json:"platformCachePartitionTypes,omitempty"`
}

type PlatformCachePartitionType struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PlatformCachePartitionType"`

	AllocatedCapacity int32 `json:"allocatedCapacity,omitempty"`

	AllocatedPurchasedCapacity int32 `json:"allocatedPurchasedCapacity,omitempty"`

	AllocatedTrialCapacity int32 `json:"allocatedTrialCapacity,omitempty"`

	CacheType *PlatformCacheType `json:"cacheType,omitempty"`
}

type Portal struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Portal"`

	*Metadata

	Active bool `json:"active,omitempty"`

	Admin string `json:"admin,omitempty"`

	DefaultLanguage string `json:"defaultLanguage,omitempty"`

	Description string `json:"description,omitempty"`

	EmailSenderAddress string `json:"emailSenderAddress,omitempty"`

	EmailSenderName string `json:"emailSenderName,omitempty"`

	EnableSelfCloseCase bool `json:"enableSelfCloseCase,omitempty"`

	FooterDocument string `json:"footerDocument,omitempty"`

	ForgotPassTemplate string `json:"forgotPassTemplate,omitempty"`

	HeaderDocument string `json:"headerDocument,omitempty"`

	IsSelfRegistrationActivated bool `json:"isSelfRegistrationActivated,omitempty"`

	LoginHeaderDocument string `json:"loginHeaderDocument,omitempty"`

	LogoDocument string `json:"logoDocument,omitempty"`

	LogoutUrl string `json:"logoutUrl,omitempty"`

	NewCommentTemplate string `json:"newCommentTemplate,omitempty"`

	NewPassTemplate string `json:"newPassTemplate,omitempty"`

	NewUserTemplate string `json:"newUserTemplate,omitempty"`

	OwnerNotifyTemplate string `json:"ownerNotifyTemplate,omitempty"`

	SelfRegNewUserUrl string `json:"selfRegNewUserUrl,omitempty"`

	SelfRegUserDefaultProfile string `json:"selfRegUserDefaultProfile,omitempty"`

	SelfRegUserDefaultRole *PortalRoles `json:"selfRegUserDefaultRole,omitempty"`

	SelfRegUserTemplate string `json:"selfRegUserTemplate,omitempty"`

	ShowActionConfirmation bool `json:"showActionConfirmation,omitempty"`

	StylesheetDocument string `json:"stylesheetDocument,omitempty"`

	Type_ *PortalType `json:"type,omitempty"`
}

type PostTemplate struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PostTemplate"`

	*Metadata

	Default_ bool `json:"default,omitempty"`

	Description string `json:"description,omitempty"`

	Fields []string `json:"fields,omitempty"`

	Label string `json:"label,omitempty"`
}

type ProductSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProductSettings"`

	*Metadata

	EnableCascadeActivateToRelatedPrices bool `json:"enableCascadeActivateToRelatedPrices,omitempty"`

	EnableQuantitySchedule bool `json:"enableQuantitySchedule,omitempty"`

	EnableRevenueSchedule bool `json:"enableRevenueSchedule,omitempty"`
}

type Profile struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Profile"`

	*Metadata

	ApplicationVisibilities []*ProfileApplicationVisibility `json:"applicationVisibilities,omitempty"`

	ClassAccesses []*ProfileApexClassAccess `json:"classAccesses,omitempty"`

	Custom bool `json:"custom,omitempty"`

	CustomPermissions []*ProfileCustomPermissions `json:"customPermissions,omitempty"`

	Description string `json:"description,omitempty"`

	ExternalDataSourceAccesses []*ProfileExternalDataSourceAccess `json:"externalDataSourceAccesses,omitempty"`

	FieldPermissions []*ProfileFieldLevelSecurity `json:"fieldPermissions,omitempty"`

	LayoutAssignments []*ProfileLayoutAssignment `json:"layoutAssignments,omitempty"`

	LoginHours *ProfileLoginHours `json:"loginHours,omitempty"`

	LoginIpRanges []*ProfileLoginIpRange `json:"loginIpRanges,omitempty"`

	ObjectPermissions []*ProfileObjectPermissions `json:"objectPermissions,omitempty"`

	PageAccesses []*ProfileApexPageAccess `json:"pageAccesses,omitempty"`

	ProfileActionOverrides []*ProfileActionOverride `json:"profileActionOverrides,omitempty"`

	RecordTypeVisibilities []*ProfileRecordTypeVisibility `json:"recordTypeVisibilities,omitempty"`

	TabVisibilities []*ProfileTabVisibility `json:"tabVisibilities,omitempty"`

	UserLicense string `json:"userLicense,omitempty"`

	UserPermissions []*ProfileUserPermission `json:"userPermissions,omitempty"`
}

type ProfileApplicationVisibility struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileApplicationVisibility"`

	Application string `json:"application,omitempty"`

	Default_ bool `json:"default,omitempty"`

	Visible bool `json:"visible,omitempty"`
}

type ProfileApexClassAccess struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileApexClassAccess"`

	ApexClass string `json:"apexClass,omitempty"`

	Enabled bool `json:"enabled,omitempty"`
}

type ProfileCustomPermissions struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileCustomPermissions"`

	Enabled bool `json:"enabled,omitempty"`

	Name string `json:"name,omitempty"`
}

type ProfileExternalDataSourceAccess struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileExternalDataSourceAccess"`

	Enabled bool `json:"enabled,omitempty"`

	ExternalDataSource string `json:"externalDataSource,omitempty"`
}

type ProfileFieldLevelSecurity struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileFieldLevelSecurity"`

	Editable bool `json:"editable,omitempty"`

	Field string `json:"field,omitempty"`

	Readable bool `json:"readable,omitempty"`
}

type ProfileLayoutAssignment struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileLayoutAssignment"`

	Layout string `json:"layout,omitempty"`

	RecordType string `json:"recordType,omitempty"`
}

type ProfileLoginHours struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileLoginHours"`

	FridayEnd string `json:"fridayEnd,omitempty"`

	FridayStart string `json:"fridayStart,omitempty"`

	MondayEnd string `json:"mondayEnd,omitempty"`

	MondayStart string `json:"mondayStart,omitempty"`

	SaturdayEnd string `json:"saturdayEnd,omitempty"`

	SaturdayStart string `json:"saturdayStart,omitempty"`

	SundayEnd string `json:"sundayEnd,omitempty"`

	SundayStart string `json:"sundayStart,omitempty"`

	ThursdayEnd string `json:"thursdayEnd,omitempty"`

	ThursdayStart string `json:"thursdayStart,omitempty"`

	TuesdayEnd string `json:"tuesdayEnd,omitempty"`

	TuesdayStart string `json:"tuesdayStart,omitempty"`

	WednesdayEnd string `json:"wednesdayEnd,omitempty"`

	WednesdayStart string `json:"wednesdayStart,omitempty"`
}

type ProfileLoginIpRange struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileLoginIpRange"`

	Description string `json:"description,omitempty"`

	EndAddress string `json:"endAddress,omitempty"`

	StartAddress string `json:"startAddress,omitempty"`
}

type ProfileApexPageAccess struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileApexPageAccess"`

	ApexPage string `json:"apexPage,omitempty"`

	Enabled bool `json:"enabled,omitempty"`
}

type ProfileActionOverride struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileActionOverride"`

	ActionName string `json:"actionName,omitempty"`

	Content string `json:"content,omitempty"`

	FormFactor *FormFactor `json:"formFactor,omitempty"`

	PageOrSobjectType string `json:"pageOrSobjectType,omitempty"`

	RecordType string `json:"recordType,omitempty"`

	Type_ *ActionOverrideType `json:"type,omitempty"`
}

type ProfileRecordTypeVisibility struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileRecordTypeVisibility"`

	Default_ bool `json:"default,omitempty"`

	PersonAccountDefault bool `json:"personAccountDefault,omitempty"`

	RecordType string `json:"recordType,omitempty"`

	Visible bool `json:"visible,omitempty"`
}

type ProfileTabVisibility struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileTabVisibility"`

	Tab string `json:"tab,omitempty"`

	Visibility *TabVisibility `json:"visibility,omitempty"`
}

type ProfileUserPermission struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ProfileUserPermission"`

	Enabled bool `json:"enabled,omitempty"`

	Name string `json:"name,omitempty"`
}

type Queue struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Queue"`

	*Metadata

	DoesSendEmailToMembers bool `json:"doesSendEmailToMembers,omitempty"`

	Email string `json:"email,omitempty"`

	Name string `json:"name,omitempty"`

	QueueSobject []*QueueSobject `json:"queueSobject,omitempty"`
}

type QueueSobject struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata QueueSobject"`

	SobjectType string `json:"sobjectType,omitempty"`
}

type QuickAction struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata QuickAction"`

	*Metadata

	Canvas string `json:"canvas,omitempty"`

	Description string `json:"description,omitempty"`

	FieldOverrides []*FieldOverride `json:"fieldOverrides,omitempty"`

	Height int32 `json:"height,omitempty"`

	Icon string `json:"icon,omitempty"`

	IsProtected bool `json:"isProtected,omitempty"`

	Label string `json:"label,omitempty"`

	LightningComponent string `json:"lightningComponent,omitempty"`

	OptionsCreateFeedItem bool `json:"optionsCreateFeedItem,omitempty"`

	Page string `json:"page,omitempty"`

	QuickActionLayout *QuickActionLayout `json:"quickActionLayout,omitempty"`

	QuickActionSendEmailOptions *QuickActionSendEmailOptions `json:"quickActionSendEmailOptions,omitempty"`

	StandardLabel *QuickActionLabel `json:"standardLabel,omitempty"`

	SuccessMessage string `json:"successMessage,omitempty"`

	TargetObject string `json:"targetObject,omitempty"`

	TargetParentField string `json:"targetParentField,omitempty"`

	TargetRecordType string `json:"targetRecordType,omitempty"`

	Type_ *QuickActionType `json:"type,omitempty"`

	Width int32 `json:"width,omitempty"`
}

type FieldOverride struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FieldOverride"`

	Field string `json:"field,omitempty"`

	Formula string `json:"formula,omitempty"`

	LiteralValue string `json:"literalValue,omitempty"`
}

type QuickActionLayout struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata QuickActionLayout"`

	LayoutSectionStyle *LayoutSectionStyle `json:"layoutSectionStyle,omitempty"`

	QuickActionLayoutColumns []*QuickActionLayoutColumn `json:"quickActionLayoutColumns,omitempty"`
}

type QuickActionLayoutColumn struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata QuickActionLayoutColumn"`

	QuickActionLayoutItems []*QuickActionLayoutItem `json:"quickActionLayoutItems,omitempty"`
}

type QuickActionLayoutItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata QuickActionLayoutItem"`

	EmptySpace bool `json:"emptySpace,omitempty"`

	Field string `json:"field,omitempty"`

	UiBehavior *UiBehavior `json:"uiBehavior,omitempty"`
}

type QuickActionSendEmailOptions struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata QuickActionSendEmailOptions"`

	DefaultEmailTemplateName string `json:"defaultEmailTemplateName,omitempty"`

	IgnoreDefaultEmailTemplateSubject bool `json:"ignoreDefaultEmailTemplateSubject,omitempty"`
}

type QuoteSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata QuoteSettings"`

	*Metadata

	EnableQuote bool `json:"enableQuote,omitempty"`
}

type RemoteSiteSetting struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RemoteSiteSetting"`

	*Metadata

	Description string `json:"description,omitempty"`

	DisableProtocolSecurity bool `json:"disableProtocolSecurity,omitempty"`

	IsActive bool `json:"isActive,omitempty"`

	Url string `json:"url,omitempty"`
}

type Report struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Report"`

	*Metadata

	Aggregates []*ReportAggregate `json:"aggregates,omitempty"`

	Block []*Report `json:"block,omitempty"`

	BlockInfo *ReportBlockInfo `json:"blockInfo,omitempty"`

	Buckets []*ReportBucketField `json:"buckets,omitempty"`

	Chart *ReportChart `json:"chart,omitempty"`

	ColorRanges []*ReportColorRange `json:"colorRanges,omitempty"`

	Columns []*ReportColumn `json:"columns,omitempty"`

	CrossFilters []*ReportCrossFilter `json:"crossFilters,omitempty"`

	Currency *CurrencyIsoCode `json:"currency,omitempty"`

	DataCategoryFilters []*ReportDataCategoryFilter `json:"dataCategoryFilters,omitempty"`

	Description string `json:"description,omitempty"`

	Division string `json:"division,omitempty"`

	Filter *ReportFilter `json:"filter,omitempty"`

	FolderName string `json:"folderName,omitempty"`

	Format *ReportFormat `json:"format,omitempty"`

	GroupingsAcross []*ReportGrouping `json:"groupingsAcross,omitempty"`

	GroupingsDown []*ReportGrouping `json:"groupingsDown,omitempty"`

	HistoricalSelector *ReportHistoricalSelector `json:"historicalSelector,omitempty"`

	Name string `json:"name,omitempty"`

	NumSubscriptions int32 `json:"numSubscriptions,omitempty"`

	Params []*ReportParam `json:"params,omitempty"`

	ReportType string `json:"reportType,omitempty"`

	RoleHierarchyFilter string `json:"roleHierarchyFilter,omitempty"`

	RowLimit int32 `json:"rowLimit,omitempty"`

	Scope string `json:"scope,omitempty"`

	ShowCurrentDate bool `json:"showCurrentDate,omitempty"`

	ShowDetails bool `json:"showDetails,omitempty"`

	SortColumn string `json:"sortColumn,omitempty"`

	SortOrder *SortOrder `json:"sortOrder,omitempty"`

	TerritoryHierarchyFilter string `json:"territoryHierarchyFilter,omitempty"`

	TimeFrameFilter *ReportTimeFrameFilter `json:"timeFrameFilter,omitempty"`

	UserFilter string `json:"userFilter,omitempty"`
}

type ReportAggregate struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportAggregate"`

	AcrossGroupingContext string `json:"acrossGroupingContext,omitempty"`

	CalculatedFormula string `json:"calculatedFormula,omitempty"`

	Datatype *ReportAggregateDatatype `json:"datatype,omitempty"`

	Description string `json:"description,omitempty"`

	DeveloperName string `json:"developerName,omitempty"`

	DownGroupingContext string `json:"downGroupingContext,omitempty"`

	IsActive bool `json:"isActive,omitempty"`

	IsCrossBlock bool `json:"isCrossBlock,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	ReportType string `json:"reportType,omitempty"`

	Scale int32 `json:"scale,omitempty"`
}

type ReportBlockInfo struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportBlockInfo"`

	AggregateReferences []*ReportAggregateReference `json:"aggregateReferences,omitempty"`

	BlockId string `json:"blockId,omitempty"`

	JoinTable string `json:"joinTable,omitempty"`
}

type ReportAggregateReference struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportAggregateReference"`

	Aggregate string `json:"aggregate,omitempty"`
}

type ReportBucketField struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportBucketField"`

	BucketType *ReportBucketFieldType `json:"bucketType,omitempty"`

	DeveloperName string `json:"developerName,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	NullTreatment *ReportFormulaNullTreatment `json:"nullTreatment,omitempty"`

	OtherBucketLabel string `json:"otherBucketLabel,omitempty"`

	SourceColumnName string `json:"sourceColumnName,omitempty"`

	UseOther bool `json:"useOther,omitempty"`

	Values []*ReportBucketFieldValue `json:"values,omitempty"`
}

type ReportBucketFieldValue struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportBucketFieldValue"`

	SourceValues []*ReportBucketFieldSourceValue `json:"sourceValues,omitempty"`

	Value string `json:"value,omitempty"`
}

type ReportBucketFieldSourceValue struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportBucketFieldSourceValue"`

	From string `json:"from,omitempty"`

	SourceValue string `json:"sourceValue,omitempty"`

	To string `json:"to,omitempty"`
}

type ReportChart struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportChart"`

	BackgroundColor1 string `json:"backgroundColor1,omitempty"`

	BackgroundColor2 string `json:"backgroundColor2,omitempty"`

	BackgroundFadeDir *ChartBackgroundDirection `json:"backgroundFadeDir,omitempty"`

	ChartSummaries []*ChartSummary `json:"chartSummaries,omitempty"`

	ChartType *ChartType `json:"chartType,omitempty"`

	EnableHoverLabels bool `json:"enableHoverLabels,omitempty"`

	ExpandOthers bool `json:"expandOthers,omitempty"`

	GroupingColumn string `json:"groupingColumn,omitempty"`

	LegendPosition *ChartLegendPosition `json:"legendPosition,omitempty"`

	Location *ChartPosition `json:"location,omitempty"`

	SecondaryGroupingColumn string `json:"secondaryGroupingColumn,omitempty"`

	ShowAxisLabels bool `json:"showAxisLabels,omitempty"`

	ShowPercentage bool `json:"showPercentage,omitempty"`

	ShowTotal bool `json:"showTotal,omitempty"`

	ShowValues bool `json:"showValues,omitempty"`

	Size *ReportChartSize `json:"size,omitempty"`

	SummaryAxisManualRangeEnd float64 `json:"summaryAxisManualRangeEnd,omitempty"`

	SummaryAxisManualRangeStart float64 `json:"summaryAxisManualRangeStart,omitempty"`

	SummaryAxisRange *ChartRangeType `json:"summaryAxisRange,omitempty"`

	TextColor string `json:"textColor,omitempty"`

	TextSize int32 `json:"textSize,omitempty"`

	Title string `json:"title,omitempty"`

	TitleColor string `json:"titleColor,omitempty"`

	TitleSize int32 `json:"titleSize,omitempty"`
}

type ReportColorRange struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportColorRange"`

	Aggregate *ReportSummaryType `json:"aggregate,omitempty"`

	ColumnName string `json:"columnName,omitempty"`

	HighBreakpoint float64 `json:"highBreakpoint,omitempty"`

	HighColor string `json:"highColor,omitempty"`

	LowBreakpoint float64 `json:"lowBreakpoint,omitempty"`

	LowColor string `json:"lowColor,omitempty"`

	MidColor string `json:"midColor,omitempty"`
}

type ReportColumn struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportColumn"`

	AggregateTypes []*ReportSummaryType `json:"aggregateTypes,omitempty"`

	Field string `json:"field,omitempty"`

	ReverseColors bool `json:"reverseColors,omitempty"`

	ShowChanges bool `json:"showChanges,omitempty"`
}

type ReportCrossFilter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportCrossFilter"`

	CriteriaItems []*ReportFilterItem `json:"criteriaItems,omitempty"`

	Operation *ObjectFilterOperator `json:"operation,omitempty"`

	PrimaryTableColumn string `json:"primaryTableColumn,omitempty"`

	RelatedTable string `json:"relatedTable,omitempty"`

	RelatedTableJoinColumn string `json:"relatedTableJoinColumn,omitempty"`
}

type ReportFilterItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportFilterItem"`

	Column string `json:"column,omitempty"`

	ColumnToColumn bool `json:"columnToColumn,omitempty"`

	IsUnlocked bool `json:"isUnlocked,omitempty"`

	Operator *FilterOperation `json:"operator,omitempty"`

	Snapshot string `json:"snapshot,omitempty"`

	Value string `json:"value,omitempty"`
}

type ReportDataCategoryFilter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportDataCategoryFilter"`

	DataCategory string `json:"dataCategory,omitempty"`

	DataCategoryGroup string `json:"dataCategoryGroup,omitempty"`

	Operator *DataCategoryFilterOperation `json:"operator,omitempty"`
}

type ReportFilter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportFilter"`

	BooleanFilter string `json:"booleanFilter,omitempty"`

	CriteriaItems []*ReportFilterItem `json:"criteriaItems,omitempty"`

	Language *Language `json:"language,omitempty"`
}

type ReportGrouping struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportGrouping"`

	AggregateType *ReportAggrType `json:"aggregateType,omitempty"`

	DateGranularity *UserDateGranularity `json:"dateGranularity,omitempty"`

	Field string `json:"field,omitempty"`

	SortByName string `json:"sortByName,omitempty"`

	SortOrder *SortOrder `json:"sortOrder,omitempty"`

	SortType *ReportSortType `json:"sortType,omitempty"`
}

type ReportHistoricalSelector struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportHistoricalSelector"`

	Snapshot []string `json:"snapshot,omitempty"`
}

type ReportParam struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportParam"`

	Name string `json:"name,omitempty"`

	Value string `json:"value,omitempty"`
}

type ReportTimeFrameFilter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportTimeFrameFilter"`

	DateColumn string `json:"dateColumn,omitempty"`

	EndDate time.Time `json:"endDate,omitempty"`

	Interval *UserDateInterval `json:"interval,omitempty"`

	StartDate time.Time `json:"startDate,omitempty"`
}

type ReportType struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportType"`

	*Metadata

	Autogenerated bool `json:"autogenerated,omitempty"`

	BaseObject string `json:"baseObject,omitempty"`

	Category *ReportTypeCategory `json:"category,omitempty"`

	Deployed bool `json:"deployed,omitempty"`

	Description string `json:"description,omitempty"`

	Join *ObjectRelationship `json:"join,omitempty"`

	Label string `json:"label,omitempty"`

	Sections []*ReportLayoutSection `json:"sections,omitempty"`
}

type ObjectRelationship struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ObjectRelationship"`

	Join *ObjectRelationship `json:"join,omitempty"`

	OuterJoin bool `json:"outerJoin,omitempty"`

	Relationship string `json:"relationship,omitempty"`
}

type ReportLayoutSection struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportLayoutSection"`

	Columns []*ReportTypeColumn `json:"columns,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`
}

type ReportTypeColumn struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportTypeColumn"`

	CheckedByDefault bool `json:"checkedByDefault,omitempty"`

	DisplayNameOverride string `json:"displayNameOverride,omitempty"`

	Field string `json:"field,omitempty"`

	Table string `json:"table,omitempty"`
}

type RoleOrTerritory struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata RoleOrTerritory"`

	*Metadata

	CaseAccessLevel string `json:"caseAccessLevel,omitempty"`

	ContactAccessLevel string `json:"contactAccessLevel,omitempty"`

	Description string `json:"description,omitempty"`

	MayForecastManagerShare bool `json:"mayForecastManagerShare,omitempty"`

	Name string `json:"name,omitempty"`

	OpportunityAccessLevel string `json:"opportunityAccessLevel,omitempty"`
}

type Role struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Role"`

	*RoleOrTerritory

	ParentRole string `json:"parentRole,omitempty"`
}

type Territory struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Territory"`

	*RoleOrTerritory

	AccountAccessLevel string `json:"accountAccessLevel,omitempty"`

	ParentTerritory string `json:"parentTerritory,omitempty"`
}

type SamlSsoConfig struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SamlSsoConfig"`

	*Metadata

	AttributeName string `json:"attributeName,omitempty"`

	AttributeNameIdFormat string `json:"attributeNameIdFormat,omitempty"`

	DecryptionCertificate string `json:"decryptionCertificate,omitempty"`

	ErrorUrl string `json:"errorUrl,omitempty"`

	ExecutionUserId string `json:"executionUserId,omitempty"`

	IdentityLocation *SamlIdentityLocationType `json:"identityLocation,omitempty"`

	IdentityMapping *SamlIdentityType `json:"identityMapping,omitempty"`

	Issuer string `json:"issuer,omitempty"`

	LoginUrl string `json:"loginUrl,omitempty"`

	LogoutUrl string `json:"logoutUrl,omitempty"`

	Name string `json:"name,omitempty"`

	OauthTokenEndpoint string `json:"oauthTokenEndpoint,omitempty"`

	RedirectBinding bool `json:"redirectBinding,omitempty"`

	RequestSignatureMethod string `json:"requestSignatureMethod,omitempty"`

	ObjforceLoginUrl string `json:"objforceLoginUrl,omitempty"`

	SamlEntityId string `json:"samlEntityId,omitempty"`

	SamlJitHandlerId string `json:"samlJitHandlerId,omitempty"`

	SamlVersion *SamlType `json:"samlVersion,omitempty"`

	UserProvisioning bool `json:"userProvisioning,omitempty"`

	ValidationCert string `json:"validationCert,omitempty"`
}

type SearchSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SearchSettings"`

	*Metadata

	DocumentContentSearchEnabled bool `json:"documentContentSearchEnabled,omitempty"`

	OptimizeSearchForCJKEnabled bool `json:"optimizeSearchForCJKEnabled,omitempty"`

	RecentlyViewedUsersForBlankLookupEnabled bool `json:"recentlyViewedUsersForBlankLookupEnabled,omitempty"`

	SearchSettingsByObject *SearchSettingsByObject `json:"searchSettingsByObject,omitempty"`

	SidebarAutoCompleteEnabled bool `json:"sidebarAutoCompleteEnabled,omitempty"`

	SidebarDropDownListEnabled bool `json:"sidebarDropDownListEnabled,omitempty"`

	SidebarLimitToItemsIOwnCheckboxEnabled bool `json:"sidebarLimitToItemsIOwnCheckboxEnabled,omitempty"`

	SingleSearchResultShortcutEnabled bool `json:"singleSearchResultShortcutEnabled,omitempty"`

	SpellCorrectKnowledgeSearchEnabled bool `json:"spellCorrectKnowledgeSearchEnabled,omitempty"`
}

type SearchSettingsByObject struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SearchSettingsByObject"`

	SearchSettingsByObject []*ObjectSearchSetting `json:"searchSettingsByObject,omitempty"`
}

type ObjectSearchSetting struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ObjectSearchSetting"`

	EnhancedLookupEnabled bool `json:"enhancedLookupEnabled,omitempty"`

	LookupAutoCompleteEnabled bool `json:"lookupAutoCompleteEnabled,omitempty"`

	Name string `json:"name,omitempty"`

	ResultsPerPageCount int32 `json:"resultsPerPageCount,omitempty"`
}

type SecuritySettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SecuritySettings"`

	*Metadata

	NetworkAccess *NetworkAccess `json:"networkAccess,omitempty"`

	PasswordPolicies *PasswordPolicies `json:"passwordPolicies,omitempty"`

	SessionSettings *SessionSettings `json:"sessionSettings,omitempty"`
}

type NetworkAccess struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata NetworkAccess"`

	IpRanges []*IpRange `json:"ipRanges,omitempty"`
}

type IpRange struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata IpRange"`

	Description string `json:"description,omitempty"`

	End string `json:"end,omitempty"`

	Start string `json:"start,omitempty"`
}

type PasswordPolicies struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata PasswordPolicies"`

	ApiOnlyUserHomePageURL string `json:"apiOnlyUserHomePageURL,omitempty"`

	Complexity *Complexity `json:"complexity,omitempty"`

	Expiration *Expiration `json:"expiration,omitempty"`

	HistoryRestriction string `json:"historyRestriction,omitempty"`

	LockoutInterval *LockoutInterval `json:"lockoutInterval,omitempty"`

	MaxLoginAttempts *MaxLoginAttempts `json:"maxLoginAttempts,omitempty"`

	MinimumPasswordLength string `json:"minimumPasswordLength,omitempty"`

	MinimumPasswordLifetime bool `json:"minimumPasswordLifetime,omitempty"`

	ObscureSecretAnswer bool `json:"obscureSecretAnswer,omitempty"`

	PasswordAssistanceMessage string `json:"passwordAssistanceMessage,omitempty"`

	PasswordAssistanceURL string `json:"passwordAssistanceURL,omitempty"`

	QuestionRestriction *QuestionRestriction `json:"questionRestriction,omitempty"`
}

type SessionSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SessionSettings"`

	DisableTimeoutWarning bool `json:"disableTimeoutWarning,omitempty"`

	EnableCSPOnEmail bool `json:"enableCSPOnEmail,omitempty"`

	EnableCSRFOnGet bool `json:"enableCSRFOnGet,omitempty"`

	EnableCSRFOnPost bool `json:"enableCSRFOnPost,omitempty"`

	EnableCacheAndAutocomplete bool `json:"enableCacheAndAutocomplete,omitempty"`

	EnableClickjackNonsetupSFDC bool `json:"enableClickjackNonsetupSFDC,omitempty"`

	EnableClickjackNonsetupUser bool `json:"enableClickjackNonsetupUser,omitempty"`

	EnableClickjackNonsetupUserHeaderless bool `json:"enableClickjackNonsetupUserHeaderless,omitempty"`

	EnableClickjackSetup bool `json:"enableClickjackSetup,omitempty"`

	EnablePostForSessions bool `json:"enablePostForSessions,omitempty"`

	EnableSMSIdentity bool `json:"enableSMSIdentity,omitempty"`

	EnforceIpRangesEveryRequest bool `json:"enforceIpRangesEveryRequest,omitempty"`

	ForceLogoutOnSessionTimeout bool `json:"forceLogoutOnSessionTimeout,omitempty"`

	ForceRelogin bool `json:"forceRelogin,omitempty"`

	LockSessionsToDomain bool `json:"lockSessionsToDomain,omitempty"`

	LockSessionsToIp bool `json:"lockSessionsToIp,omitempty"`

	LogoutURL string `json:"logoutURL,omitempty"`

	SecurityCentralKillSession bool `json:"securityCentralKillSession,omitempty"`

	SessionTimeout *SessionTimeout `json:"sessionTimeout,omitempty"`
}

type SharingBaseRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SharingBaseRule"`

	*Metadata

	AccessLevel string `json:"accessLevel,omitempty"`

	AccountSettings *AccountSharingRuleSettings `json:"accountSettings,omitempty"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	SharedTo *SharedTo `json:"sharedTo,omitempty"`
}

type AccountSharingRuleSettings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AccountSharingRuleSettings"`

	CaseAccessLevel string `json:"caseAccessLevel,omitempty"`

	ContactAccessLevel string `json:"contactAccessLevel,omitempty"`

	OpportunityAccessLevel string `json:"opportunityAccessLevel,omitempty"`
}

type SharingCriteriaRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SharingCriteriaRule"`

	*SharingBaseRule

	BooleanFilter string `json:"booleanFilter,omitempty"`

	CriteriaItems []*FilterItem `json:"criteriaItems,omitempty"`
}

type SharingOwnerRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SharingOwnerRule"`

	*SharingBaseRule

	SharedFrom *SharedTo `json:"sharedFrom,omitempty"`
}

type SharingTerritoryRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SharingTerritoryRule"`

	*SharingOwnerRule
}

type SharingRules struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SharingRules"`

	*Metadata

	SharingCriteriaRules []*SharingCriteriaRule `json:"sharingCriteriaRules,omitempty"`

	SharingOwnerRules []*SharingOwnerRule `json:"sharingOwnerRules,omitempty"`

	SharingTerritoryRules []*SharingTerritoryRule `json:"sharingTerritoryRules,omitempty"`
}

type SharingSet struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SharingSet"`

	*Metadata

	AccessMappings []*AccessMapping `json:"accessMappings,omitempty"`

	Description string `json:"description,omitempty"`

	Name string `json:"name,omitempty"`

	Profiles []string `json:"profiles,omitempty"`
}

type AccessMapping struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata AccessMapping"`

	AccessLevel string `json:"accessLevel,omitempty"`

	Object string `json:"object,omitempty"`

	ObjectField string `json:"objectField,omitempty"`

	UserField string `json:"userField,omitempty"`
}

type Skill struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Skill"`

	*Metadata

	Assignments *SkillAssignments `json:"assignments,omitempty"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`
}

type SkillAssignments struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SkillAssignments"`

	Profiles *SkillProfileAssignments `json:"profiles,omitempty"`

	Users *SkillUserAssignments `json:"users,omitempty"`
}

type SkillProfileAssignments struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SkillProfileAssignments"`

	Profile []string `json:"profile,omitempty"`
}

type SkillUserAssignments struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SkillUserAssignments"`

	User []string `json:"user,omitempty"`
}

type StandardValueSet struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata StandardValueSet"`

	*Metadata

	Sorted bool `json:"sorted,omitempty"`

	StandardValue []*StandardValue `json:"standardValue,omitempty"`
}

type StandardValueSetTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata StandardValueSetTranslation"`

	*Metadata

	ValueTranslation []*ValueTranslation `json:"valueTranslation,omitempty"`
}

type SynonymDictionary struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SynonymDictionary"`

	*Metadata

	Groups []*SynonymGroup `json:"groups,omitempty"`

	IsProtected bool `json:"isProtected,omitempty"`

	Label string `json:"label,omitempty"`
}

type SynonymGroup struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata SynonymGroup"`

	Languages []*Language `json:"languages,omitempty"`

	Terms []string `json:"terms,omitempty"`
}

type Territory2 struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Territory2"`

	*Metadata

	AccountAccessLevel string `json:"accountAccessLevel,omitempty"`

	CaseAccessLevel string `json:"caseAccessLevel,omitempty"`

	ContactAccessLevel string `json:"contactAccessLevel,omitempty"`

	CustomFields []*FieldValue `json:"customFields,omitempty"`

	Description string `json:"description,omitempty"`

	Name string `json:"name,omitempty"`

	OpportunityAccessLevel string `json:"opportunityAccessLevel,omitempty"`

	ParentTerritory string `json:"parentTerritory,omitempty"`

	RuleAssociations []*Territory2RuleAssociation `json:"ruleAssociations,omitempty"`

	Territory2Type string `json:"territory2Type,omitempty"`
}

type FieldValue struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata FieldValue"`

	Name string `json:"name,omitempty"`

	Value interface{} `json:"value,omitempty"`
}

type Territory2RuleAssociation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Territory2RuleAssociation"`

	Inherited bool `json:"inherited,omitempty"`

	RuleName string `json:"ruleName,omitempty"`
}

type Territory2Model struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Territory2Model"`

	*Metadata

	CustomFields []*FieldValue `json:"customFields,omitempty"`

	Description string `json:"description,omitempty"`

	Name string `json:"name,omitempty"`
}

type Territory2Rule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Territory2Rule"`

	*Metadata

	Active bool `json:"active,omitempty"`

	BooleanFilter string `json:"booleanFilter,omitempty"`

	Name string `json:"name,omitempty"`

	ObjectType string `json:"objectType,omitempty"`

	RuleItems []*Territory2RuleItem `json:"ruleItems,omitempty"`
}

type Territory2RuleItem struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Territory2RuleItem"`

	Field string `json:"field,omitempty"`

	Operation *FilterOperation `json:"operation,omitempty"`

	Value string `json:"value,omitempty"`
}

type Territory2Settings struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Territory2Settings"`

	*Metadata

	DefaultAccountAccessLevel string `json:"defaultAccountAccessLevel,omitempty"`

	DefaultCaseAccessLevel string `json:"defaultCaseAccessLevel,omitempty"`

	DefaultContactAccessLevel string `json:"defaultContactAccessLevel,omitempty"`

	DefaultOpportunityAccessLevel string `json:"defaultOpportunityAccessLevel,omitempty"`

	OpportunityFilterSettings *Territory2SettingsOpportunityFilter `json:"opportunityFilterSettings,omitempty"`
}

type Territory2SettingsOpportunityFilter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Territory2SettingsOpportunityFilter"`

	ApexClassName string `json:"apexClassName,omitempty"`

	EnableFilter bool `json:"enableFilter,omitempty"`

	RunOnCreate bool `json:"runOnCreate,omitempty"`
}

type Territory2Type struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Territory2Type"`

	*Metadata

	Description string `json:"description,omitempty"`

	Name string `json:"name,omitempty"`

	Priority int32 `json:"priority,omitempty"`
}

type TransactionSecurityPolicy struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata TransactionSecurityPolicy"`

	*Metadata

	Action *TransactionSecurityAction `json:"action,omitempty"`

	Active bool `json:"active,omitempty"`

	ApexClass string `json:"apexClass,omitempty"`

	EventType *MonitoredEvents `json:"eventType,omitempty"`

	ExecutionUser string `json:"executionUser,omitempty"`

	ResourceName string `json:"resourceName,omitempty"`
}

type TransactionSecurityAction struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata TransactionSecurityAction"`

	Block bool `json:"block,omitempty"`

	EndSession bool `json:"endSession,omitempty"`

	Notifications []*TransactionSecurityNotification `json:"notifications,omitempty"`

	TwoFactorAuthentication bool `json:"twoFactorAuthentication,omitempty"`
}

type TransactionSecurityNotification struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata TransactionSecurityNotification"`

	InApp bool `json:"inApp,omitempty"`

	SendEmail bool `json:"sendEmail,omitempty"`

	User string `json:"user,omitempty"`
}

type Translations struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Translations"`

	*Metadata

	CustomApplications []*CustomApplicationTranslation `json:"customApplications,omitempty"`

	CustomDataTypeTranslations []*CustomDataTypeTranslation `json:"customDataTypeTranslations,omitempty"`

	CustomLabels []*CustomLabelTranslation `json:"customLabels,omitempty"`

	CustomPageWebLinks []*CustomPageWebLinkTranslation `json:"customPageWebLinks,omitempty"`

	CustomTabs []*CustomTabTranslation `json:"customTabs,omitempty"`

	QuickActions []*GlobalQuickActionTranslation `json:"quickActions,omitempty"`

	ReportTypes []*ReportTypeTranslation `json:"reportTypes,omitempty"`

	Scontrols []*ScontrolTranslation `json:"scontrols,omitempty"`
}

type CustomApplicationTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomApplicationTranslation"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type CustomDataTypeTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomDataTypeTranslation"`

	Components []*CustomDataTypeComponentTranslation `json:"components,omitempty"`

	CustomDataTypeName string `json:"customDataTypeName,omitempty"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`
}

type CustomDataTypeComponentTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomDataTypeComponentTranslation"`

	DeveloperSuffix string `json:"developerSuffix,omitempty"`

	Label string `json:"label,omitempty"`
}

type CustomLabelTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomLabelTranslation"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type CustomPageWebLinkTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomPageWebLinkTranslation"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type CustomTabTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata CustomTabTranslation"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type GlobalQuickActionTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata GlobalQuickActionTranslation"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type ReportTypeTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportTypeTranslation"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`

	Sections []*ReportTypeSectionTranslation `json:"sections,omitempty"`
}

type ReportTypeSectionTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportTypeSectionTranslation"`

	Columns []*ReportTypeColumnTranslation `json:"columns,omitempty"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type ReportTypeColumnTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ReportTypeColumnTranslation"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type ScontrolTranslation struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ScontrolTranslation"`

	Label string `json:"label,omitempty"`

	Name string `json:"name,omitempty"`
}

type VisualizationPlugin struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata VisualizationPlugin"`

	*Metadata

	Description string `json:"description,omitempty"`

	DeveloperName string `json:"developerName,omitempty"`

	Icon string `json:"icon,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	VisualizationResources []*VisualizationResource `json:"visualizationResources,omitempty"`

	VisualizationTypes []*VisualizationType `json:"visualizationTypes,omitempty"`
}

type VisualizationResource struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata VisualizationResource"`

	Description string `json:"description,omitempty"`

	File string `json:"file,omitempty"`

	Rank int32 `json:"rank,omitempty"`

	Type_ *VisualizationResourceType `json:"type,omitempty"`
}

type VisualizationType struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata VisualizationType"`

	Description string `json:"description,omitempty"`

	DeveloperName string `json:"developerName,omitempty"`

	Icon string `json:"icon,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	ScriptBootstrapMethod string `json:"scriptBootstrapMethod,omitempty"`
}

type WaveApplication struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WaveApplication"`

	*Metadata

	AssetIcon string `json:"assetIcon,omitempty"`

	Description string `json:"description,omitempty"`

	Folder string `json:"folder,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	Shares []*FolderShare `json:"shares,omitempty"`

	TemplateOrigin string `json:"templateOrigin,omitempty"`

	TemplateVersion string `json:"templateVersion,omitempty"`
}

type WaveDataset struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WaveDataset"`

	*Metadata

	Application string `json:"application,omitempty"`

	Description string `json:"description,omitempty"`

	MasterLabel string `json:"masterLabel,omitempty"`

	TemplateAssetSourceName string `json:"templateAssetSourceName,omitempty"`
}

type WaveTemplateBundle struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WaveTemplateBundle"`

	*Metadata

	AssetIcon string `json:"assetIcon,omitempty"`

	AssetVersion float64 `json:"assetVersion,omitempty"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	TemplateType string `json:"templateType,omitempty"`
}

type Workflow struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata Workflow"`

	*Metadata

	Alerts []*WorkflowAlert `json:"alerts,omitempty"`

	FieldUpdates []*WorkflowFieldUpdate `json:"fieldUpdates,omitempty"`

	FlowActions []*WorkflowFlowAction `json:"flowActions,omitempty"`

	KnowledgePublishes []*WorkflowKnowledgePublish `json:"knowledgePublishes,omitempty"`

	OutboundMessages []*WorkflowOutboundMessage `json:"outboundMessages,omitempty"`

	Rules []*WorkflowRule `json:"rules,omitempty"`

	Send []*WorkflowSend `json:"send,omitempty"`

	Tasks []*WorkflowTask `json:"tasks,omitempty"`
}

type WorkflowAlert struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowAlert"`

	*WorkflowAction

	CcEmails []string `json:"ccEmails,omitempty"`

	Description string `json:"description,omitempty"`

	Protected bool `json:"protected,omitempty"`

	Recipients []*WorkflowEmailRecipient `json:"recipients,omitempty"`

	SenderAddress string `json:"senderAddress,omitempty"`

	SenderType *ActionEmailSenderType `json:"senderType,omitempty"`

	Template string `json:"template,omitempty"`
}

type WorkflowAction struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowAction"`

	*Metadata
}

type WorkflowFieldUpdate struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowFieldUpdate"`

	*WorkflowAction

	Description string `json:"description,omitempty"`

	Field string `json:"field,omitempty"`

	Formula string `json:"formula,omitempty"`

	LiteralValue string `json:"literalValue,omitempty"`

	LookupValue string `json:"lookupValue,omitempty"`

	LookupValueType *LookupValueType `json:"lookupValueType,omitempty"`

	Name string `json:"name,omitempty"`

	NotifyAssignee bool `json:"notifyAssignee,omitempty"`

	Operation *FieldUpdateOperation `json:"operation,omitempty"`

	Protected bool `json:"protected,omitempty"`

	ReevaluateOnChange bool `json:"reevaluateOnChange,omitempty"`

	TargetObject string `json:"targetObject,omitempty"`
}

type WorkflowFlowAction struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowFlowAction"`

	*WorkflowAction

	Description string `json:"description,omitempty"`

	Flow string `json:"flow,omitempty"`

	FlowInputs []*WorkflowFlowActionParameter `json:"flowInputs,omitempty"`

	Label string `json:"label,omitempty"`

	Language string `json:"language,omitempty"`

	Protected bool `json:"protected,omitempty"`
}

type WorkflowFlowActionParameter struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowFlowActionParameter"`

	Name string `json:"name,omitempty"`

	Value string `json:"value,omitempty"`
}

type WorkflowKnowledgePublish struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowKnowledgePublish"`

	*WorkflowAction

	Action *KnowledgeWorkflowAction `json:"action,omitempty"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	Language string `json:"language,omitempty"`

	Protected bool `json:"protected,omitempty"`
}

type WorkflowOutboundMessage struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowOutboundMessage"`

	*WorkflowAction

	ApiVersion float64 `json:"apiVersion,omitempty"`

	Description string `json:"description,omitempty"`

	EndpointUrl string `json:"endpointUrl,omitempty"`

	Fields []string `json:"fields,omitempty"`

	IncludeSessionId bool `json:"includeSessionId,omitempty"`

	IntegrationUser string `json:"integrationUser,omitempty"`

	Name string `json:"name,omitempty"`

	Protected bool `json:"protected,omitempty"`

	UseDeadLetterQueue bool `json:"useDeadLetterQueue,omitempty"`
}

type WorkflowSend struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowSend"`

	*WorkflowAction

	Action *SendAction `json:"action,omitempty"`

	Description string `json:"description,omitempty"`

	Label string `json:"label,omitempty"`

	Language string `json:"language,omitempty"`

	Protected bool `json:"protected,omitempty"`
}

type WorkflowTask struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowTask"`

	*WorkflowAction

	AssignedTo string `json:"assignedTo,omitempty"`

	AssignedToType *ActionTaskAssignedToTypes `json:"assignedToType,omitempty"`

	Description string `json:"description,omitempty"`

	DueDateOffset int32 `json:"dueDateOffset,omitempty"`

	NotifyAssignee bool `json:"notifyAssignee,omitempty"`

	OffsetFromField string `json:"offsetFromField,omitempty"`

	Priority string `json:"priority,omitempty"`

	Protected bool `json:"protected,omitempty"`

	Status string `json:"status,omitempty"`

	Subject string `json:"subject,omitempty"`
}

type WorkflowEmailRecipient struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowEmailRecipient"`

	Field string `json:"field,omitempty"`

	Recipient string `json:"recipient,omitempty"`

	Type_ *ActionEmailRecipientTypes `json:"type,omitempty"`
}

type WorkflowRule struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowRule"`

	*Metadata

	Actions []*WorkflowActionReference `json:"actions,omitempty"`

	Active bool `json:"active,omitempty"`

	BooleanFilter string `json:"booleanFilter,omitempty"`

	CriteriaItems []*FilterItem `json:"criteriaItems,omitempty"`

	Description string `json:"description,omitempty"`

	Formula string `json:"formula,omitempty"`

	TriggerType *WorkflowTriggerTypes `json:"triggerType,omitempty"`

	WorkflowTimeTriggers []*WorkflowTimeTrigger `json:"workflowTimeTriggers,omitempty"`
}

type WorkflowTimeTrigger struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata WorkflowTimeTrigger"`

	Actions []*WorkflowActionReference `json:"actions,omitempty"`

	OffsetFromField string `json:"offsetFromField,omitempty"`

	TimeLength string `json:"timeLength,omitempty"`

	WorkflowTimeTriggerUnit *WorkflowTimeUnits `json:"workflowTimeTriggerUnit,omitempty"`
}

type XOrgHub struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata XOrgHub"`

	*Metadata

	Label string `json:"label,omitempty"`

	LockSharedObjects bool `json:"lockSharedObjects,omitempty"`

	PermissionSets []string `json:"permissionSets,omitempty"`

	SharedObjects []*XOrgHubSharedObject `json:"sharedObjects,omitempty"`
}

type XOrgHubSharedObject struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata XOrgHubSharedObject"`

	Fields []string `json:"fields,omitempty"`

	Name string `json:"name,omitempty"`
}

type SaveResult struct {
	Errors []*Error `json:"errors,omitempty"`

	FullName string `json:"fullName,omitempty"`

	Success bool `json:"success,omitempty"`
}

type Error struct {
	ExtendedErrorDetails []*ExtendedErrorDetails `json:"extendedErrorDetails,omitempty"`

	Fields []string `json:"fields,omitempty"`

	Message string `json:"message,omitempty"`

	StatusCode StatusCode `json:"statusCode,omitempty"`
}

type ExtendedErrorDetails struct {
	ExtendedErrorCode *ExtendedErrorCode `json:"extendedErrorCode,omitempty"`
}

type DeleteResult struct {
	Errors []*Error `json:"errors,omitempty"`

	FullName string `json:"fullName,omitempty"`

	Success bool `json:"success,omitempty"`
}

type DeployOptions struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata DeployOptions"`

	AllowMissingFiles bool `json:"allowMissingFiles,omitempty"`

	AutoUpdatePackage bool `json:"autoUpdatePackage,omitempty"`

	CheckOnly bool `json:"checkOnly,omitempty"`

	IgnoreWarnings bool `json:"ignoreWarnings,omitempty"`

	PerformRetrieve bool `json:"performRetrieve,omitempty"`

	PurgeOnDelete bool `json:"purgeOnDelete,omitempty"`

	RollbackOnError bool `json:"rollbackOnError,omitempty"`

	RunTests []string `json:"runTests,omitempty"`

	SinglePackage bool `json:"singlePackage,omitempty"`

	TestLevel *TestLevel `json:"testLevel,omitempty"`
}

type AsyncResult struct {
	Done bool `json:"done,omitempty"`

	Id ID `json:"id,omitempty"`

	Message string `json:"message,omitempty"`

	State AsyncRequestState `json:"state,omitempty"`

	StatusCode StatusCode `json:"statusCode,omitempty"`
}

type DescribeMetadataResult struct {
	MetadataObjects []*DescribeMetadataObject `json:"metadataObjects,omitempty"`

	OrganizationNamespace string `json:"organizationNamespace,omitempty"`

	PartialSaveAllowed bool `json:"partialSaveAllowed,omitempty"`

	TestRequired bool `json:"testRequired,omitempty"`
}

type DescribeMetadataObject struct {
	ChildXmlNames []string `json:"childXmlNames,omitempty"`

	DirectoryName string `json:"directoryName,omitempty"`

	InFolder bool `json:"inFolder,omitempty"`

	MetaFile bool `json:"metaFile,omitempty"`

	Suffix string `json:"suffix,omitempty"`

	XmlName string `json:"xmlName,omitempty"`
}

type DescribeValueTypeResult struct {
	ApiCreatable bool `json:"apiCreatable,omitempty"`

	ApiDeletable bool `json:"apiDeletable,omitempty"`

	ApiReadable bool `json:"apiReadable,omitempty"`

	ApiUpdatable bool `json:"apiUpdatable,omitempty"`

	ParentField *ValueTypeField `json:"parentField,omitempty"`

	ValueTypeFields []*ValueTypeField `json:"valueTypeFields,omitempty"`
}

type ValueTypeField struct {
	Fields []*ValueTypeField `json:"fields,omitempty"`

	ForeignKeyDomain []string `json:"foreignKeyDomain,omitempty"`

	IsForeignKey bool `json:"isForeignKey,omitempty"`

	IsNameField bool `json:"isNameField,omitempty"`

	MinOccurs int32 `json:"minOccurs,omitempty"`

	Name string `json:"name,omitempty"`

	PicklistValues []*PicklistEntry `json:"picklistValues,omitempty"`

	SoapType string `json:"soapType,omitempty"`

	ValueRequired bool `json:"valueRequired,omitempty"`
}

type PicklistEntry struct {
	Active bool `json:"active,omitempty"`

	DefaultValue bool `json:"defaultValue,omitempty"`

	Label string `json:"label,omitempty"`

	ValidFor string `json:"validFor,omitempty"`

	Value string `json:"value,omitempty"`
}

type ListMetadataQuery struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata ListMetadataQuery"`

	Folder string `json:"folder,omitempty"`

	Type string `json:"type,omitempty"`
}

type ReadResult struct {
	Records []*Metadata `json:"records,omitempty"`
}

type RetrieveRequest struct {
	ApiVersion float64 `json:"apiVersion,omitempty"`

	PackageNames []string `json:"packageNames,omitempty"`

	SinglePackage bool `json:"singlePackage,omitempty"`

	SpecificFiles []string `json:"specificFiles,omitempty"`

	Unpackaged *Package `json:"unpackaged,omitempty"`
}

type UpsertResult struct {
	Created bool `json:"created,omitempty"`

	Errors []*Error `json:"errors,omitempty"`

	FullName string `json:"fullName,omitempty"`

	Success bool `json:"success,omitempty"`
}

type LogInfo struct {
	XMLName xml.Name `json:"http://soap.sforce.com/2006/04/metadata LogInfo"`

	Category *LogCategory `json:"category,omitempty"`

	Level *LogCategoryLevel `json:"level,omitempty"`
}

type LoginRequest struct {
	XMLName  xml.Name `json:"urn:partner.soap.sforce.com login"`
	Username string   `json:"username"`
	Password string   `json:"password"`
}

type LoginResponse struct {
	XMLName     xml.Name    `json:"urn:partner.soap.sforce.com loginResponse"`
	LoginResult LoginResult `json:"result"`
}

type LoginResult struct {
	MetadataServerUrl string `json:"metadataServerUrl"`
	PasswordExpired   bool   `json:"passwordExpired"`
	Sandbox           bool   `json:"sandbox`
	ServerUrl         string `json:"serverUrl"`
	SessionId         string `json:"sessionId"`
	UserId            ID    `json:"userId"`
	//	UserInfo *UserInfo `json:"userInfo"`
}

type MetadataPortType struct {
	client *soapforce.SOAPClient
}

func (service *MetadataPortType) SetDebug(debug bool) {
	service.client.SetDebug(debug)
}

func (service *MetadataPortType) SetServerUrl(url string) {
	service.client.SetServerUrl(url)
}

func (service *MetadataPortType) SetLogger(logger io.Writer) {
	service.client.SetLogger(logger)
}

func (service *MetadataPortType) SetGzip(gz bool) {
	service.client.SetGzip(gz)
}

func NewMetadataPortType(url string, tls bool, auth *soapforce.BasicAuth) *MetadataPortType {
	if url == "" {
		url = "https://login.objforce.com/v1"
	}
	client := soapforce.NewSOAPClient(url, tls, auth)

	return &MetadataPortType{
		client: client,
	}
}

func (service *MetadataPortType) SetHeader(header interface{}) {
	service.client.AddHeader(header)
}

/* Cancels a metadata deploy. */
func (service *MetadataPortType) CancelDeploy(request *CancelDeploy) (*CancelDeployResponse, error) {
	response := new(CancelDeployResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Check the current status of an asyncronous deploy call. */
func (service *MetadataPortType) CheckDeployStatus(request *CheckDeployStatus) (*CheckDeployStatusResponse, error) {
	response := new(CheckDeployStatusResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Check the current status of an asyncronous deploy call. */
func (service *MetadataPortType) CheckRetrieveStatus(request *CheckRetrieveStatus) (*CheckRetrieveStatusResponse, error) {
	response := new(CheckRetrieveStatusResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Creates metadata entries synchronously. */
func (service *MetadataPortType) CreateMetadata(request *CreateMetadata) (*CreateMetadataResponse, error) {
	response := new(CreateMetadataResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Deletes metadata entries synchronously. */
func (service *MetadataPortType) DeleteMetadata(request *DeleteMetadata) (*DeleteMetadataResponse, error) {
	response := new(DeleteMetadataResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Deploys a zipfile full of metadata entries asynchronously. */
func (service *MetadataPortType) Deploy(request *Deploy) (*DeployResponse, error) {
	response := new(DeployResponse)
	// modify
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Deploys a previously validated payload without running tests. */
func (service *MetadataPortType) DeployRecentValidation(request *DeployRecentValidation) (*DeployRecentValidationResponse, error) {
	response := new(DeployRecentValidationResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describes features of the metadata API. */
func (service *MetadataPortType) DescribeMetadata(request *DescribeMetadata) (*DescribeMetadataResponse, error) {
	response := new(DescribeMetadataResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Describe a complex value type */
func (service *MetadataPortType) DescribeValueType(request *DescribeValueType) (*DescribeValueTypeResponse, error) {
	response := new(DescribeValueTypeResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Lists the available metadata components. */
func (service *MetadataPortType) ListMetadata(request *ListMetadata) (*ListMetadataResponse, error) {
	response := new(ListMetadataResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Reads metadata entries synchronously. */
func (service *MetadataPortType) ReadMetadata(request *ReadMetadata) (*ReadMetadataResponse, error) {
	response := new(ReadMetadataResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Renames a metadata entry synchronously. */
func (service *MetadataPortType) RenameMetadata(request *RenameMetadata) (*RenameMetadataResponse, error) {
	response := new(RenameMetadataResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Retrieves a set of individually specified metadata entries. */
func (service *MetadataPortType) Retrieve(request *Retrieve) (*RetrieveResponse, error) {
	response := new(RetrieveResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Updates metadata entries synchronously. */
func (service *MetadataPortType) UpdateMetadata(request *UpdateMetadata) (*UpdateMetadataResponse, error) {
	response := new(UpdateMetadataResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Upserts metadata entries synchronously. */
func (service *MetadataPortType) UpsertMetadata(request *UpsertMetadata) (*UpsertMetadataResponse, error) {
	response := new(UpsertMetadataResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* Upserts metadata entries synchronously. */
func (service *MetadataPortType) Login(request *LoginRequest) (*LoginResponse, error) {
	response := new(LoginResponse)
	err := service.client.Call(request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

var timeout = time.Duration(30 * time.Second)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, timeout)
}

type BasicAuth struct {
	Login    string
	Password string
}
