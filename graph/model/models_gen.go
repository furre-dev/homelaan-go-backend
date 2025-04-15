// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AvailabilityCommitment struct {
	MonthlyCommitmentHours    *int32                       `json:"monthly_commitment_hours"`
	ProjectDurationPreference []*ProjectDurationPreference `json:"project_duration_preference"`
	WillingnessToTravel       *bool                        `json:"willingness_to_travel"`
}

type AvailabilityCommitmentInput struct {
	MonthlyCommitmentHours    *int32                       `json:"monthly_commitment_hours"`
	ProjectDurationPreference []*ProjectDurationPreference `json:"project_duration_preference"`
	WillingnessToTravel       *bool                        `json:"willingness_to_travel"`
}

type EngagementPreferences struct {
	BusinessOpportunities []*BusinessOpportunities `json:"business_opportunities"`
	SectorsOfInterest     []*string                `json:"sectors_of_interest"`
	OpportunityLocations  []*OpportunityLocations  `json:"opportunity_locations"`
}

type EngagementPreferencesInput struct {
	BusinessOpportunities []*BusinessOpportunities `json:"business_opportunities"`
	SectorsOfInterest     []*string                `json:"sectors_of_interest"`
	OpportunityLocations  []*OpportunityLocations  `json:"opportunity_locations"`
}

type InvestmentAppetite struct {
	ActiveInvestorInStartupsOrSmallMidBusinesses *bool              `json:"active_investor_in_startups_or_small_mid_businesses"`
	InvestmentRange                              *InvestmentRange   `json:"investment_range"`
	AngelGroupMembership                         *bool              `json:"angel_group_membership"`
	InvolvementType                              []*InvolvementType `json:"involvement_type"`
}

type InvestmentAppetiteInput struct {
	ActiveInvestorInStartupsOrSmallMidBusinesses *bool                 `json:"active_investor_in_startups_or_small_mid_businesses"`
	InvestmentRange                              *InvestmentRangeInput `json:"investment_range"`
	AngelGroupMembership                         *bool                 `json:"angel_group_membership"`
	InvolvementType                              []*InvolvementType    `json:"involvement_type"`
}

type InvestmentRange struct {
	FromUsd *int32 `json:"from_usd"`
	ToUsd   *int32 `json:"to_usd"`
}

type InvestmentRangeInput struct {
	FromUsd *int32 `json:"from_usd"`
	ToUsd   *int32 `json:"to_usd"`
}

type InvestorProfile struct {
	GeoInfo                   *PersonalGeographicProfile `json:"geo_info"`
	ProfessionalBackground    *ProfessionalBackground    `json:"professional_background"`
	NetworkMarketAccess       *NetworkMarketAccess       `json:"network_market_access"`
	EngagementPreferences     *EngagementPreferences     `json:"engagement_preferences"`
	InvestmentAppetite        *InvestmentAppetite        `json:"investment_appetite"`
	AvailabilityAndCommitment *AvailabilityCommitment    `json:"availability_and_commitment"`
	SpousalInvolvement        *SpousalInvolvement        `json:"spousal_involvement"`
	AdditionalInfo            *string                    `json:"additional_info"`
}

type InvestorProfileInput struct {
	GeoInfo                   *PersonalGeographicProfileInput `json:"geo_info"`
	ProfessionalBackground    *ProfessionalBackgroundInput    `json:"professional_background"`
	NetworkMarketAccess       *NetworkMarketAccessInput       `json:"network_market_access"`
	EngagementPreferences     *EngagementPreferencesInput     `json:"engagement_preferences"`
	InvestmentAppetite        *InvestmentAppetiteInput        `json:"investment_appetite"`
	AvailabilityAndCommitment *AvailabilityCommitmentInput    `json:"availability_and_commitment"`
	SpousalInvolvement        *SpousalInvolvementInput        `json:"spousal_involvement"`
	AdditionalInfo            *string                         `json:"additional_info"`
}

type Location struct {
	City    *string `json:"city"`
	Country *string `json:"country"`
}

type LocationInput struct {
	City    *string `json:"city"`
	Country *string `json:"country"`
}

type Mutation struct {
}

type NetworkMarketAccess struct {
	IndustriesWithStrongNetwork []*string             `json:"industries_with_strong_network"`
	ProfessionalContacts        *ProfessionalContacts `json:"professional_contacts"`
	LocalBusinessDiasporaGroups []*string             `json:"local_business_diaspora_groups"`
}

type NetworkMarketAccessInput struct {
	IndustriesWithStrongNetwork []*string             `json:"industries_with_strong_network"`
	ProfessionalContacts        *ProfessionalContacts `json:"professional_contacts"`
	LocalBusinessDiasporaGroups []*string             `json:"local_business_diaspora_groups"`
}

type PersonalGeographicProfile struct {
	Location                          *Location        `json:"location"`
	YearsAbroad                       *int32           `json:"years_abroad"`
	WillStayLongTermInCurrentLocation *bool            `json:"will_stay_long_term_in_current_location"`
	HasFamilyAbroad                   *bool            `json:"has_family_abroad"`
	ResidencyStatus                   *ResidencyStatus `json:"residency_status"`
}

type PersonalGeographicProfileInput struct {
	Location                          *LocationInput   `json:"location"`
	YearsAbroad                       *int32           `json:"years_abroad"`
	WillStayLongTermInCurrentLocation *bool            `json:"will_stay_long_term_in_current_location"`
	HasFamilyAbroad                   *bool            `json:"has_family_abroad"`
	ResidencyStatus                   *ResidencyStatus `json:"residency_status"`
}

type ProfessionalBackground struct {
	CurrentRole         *string   `json:"current_role"`
	Organization        *string   `json:"organization"`
	Industry            *string   `json:"industry"`
	CareerOverview      *string   `json:"career_overview"`
	ExpertiseAreas      []*string `json:"expertise_areas"`
	SeniorityLevel      *string   `json:"seniority_level"`
	DecisionMakingScope *string   `json:"decision_making_scope"`
}

type ProfessionalBackgroundInput struct {
	CurrentRole         *string   `json:"current_role"`
	Organization        *string   `json:"organization"`
	Industry            *string   `json:"industry"`
	CareerOverview      *string   `json:"career_overview"`
	ExpertiseAreas      []*string `json:"expertise_areas"`
	SeniorityLevel      *string   `json:"seniority_level"`
	DecisionMakingScope *string   `json:"decision_making_scope"`
}

type Query struct {
}

type Question struct {
	QuestionTitle  string `json:"question_title"`
	ProfileField   string `json:"profile_field"`
	IsLastQuestion bool   `json:"is_last_question"`
	QuestionIndex  int32  `json:"question_index"`
}

type QuestionInput struct {
	QuestionTitle   string `json:"question_title"`
	ProfileField    string `json:"profile_field"`
	UserInputAnswer string `json:"user_input_answer"`
}

type SpousalInvolvement struct {
	SpouseInterested             *bool     `json:"spouse_interested"`
	SpouseProfessionalBackground *string   `json:"spouse_professional_background"`
	SpouseExpertiseAreas         []*string `json:"spouse_expertise_areas"`
}

type SpousalInvolvementInput struct {
	SpouseInterested             *bool     `json:"spouse_interested"`
	SpouseProfessionalBackground *string   `json:"spouse_professional_background"`
	SpouseExpertiseAreas         []*string `json:"spouse_expertise_areas"`
}

type User struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	UserType  *UserType `json:"user_type"`
}

type UserInput struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	UserType  *UserType `json:"user_type"`
}

type UserTypeInput struct {
	UserType UserType `json:"user_type"`
	ID       string   `json:"id"`
}

type BusinessOpportunities string

const (
	BusinessOpportunitiesAdvisoryRoles                         BusinessOpportunities = "ADVISORY_ROLES"
	BusinessOpportunitiesNonExecutiveOrExecutiveBoardPositions BusinessOpportunities = "NON_EXECUTIVE_OR_EXECUTIVE_BOARD_POSITIONS"
	BusinessOpportunitiesBusinessDevelopment                   BusinessOpportunities = "BUSINESS_DEVELOPMENT"
	BusinessOpportunitiesSalesOrCommercial                     BusinessOpportunities = "SALES_OR_COMMERCIAL"
	BusinessOpportunitiesInvestmentInStartups                  BusinessOpportunities = "INVESTMENT_IN_STARTUPS"
	BusinessOpportunitiesInvestmentInSmallMidBusinesses        BusinessOpportunities = "INVESTMENT_IN_SMALL_MID_BUSINESSES"
	BusinessOpportunitiesMentoring                             BusinessOpportunities = "MENTORING"
	BusinessOpportunitiesOperationalInvolvement                BusinessOpportunities = "OPERATIONAL_INVOLVEMENT"
	BusinessOpportunitiesFoundingNewOperations                 BusinessOpportunities = "FOUNDING_NEW_OPERATIONS"
)

var AllBusinessOpportunities = []BusinessOpportunities{
	BusinessOpportunitiesAdvisoryRoles,
	BusinessOpportunitiesNonExecutiveOrExecutiveBoardPositions,
	BusinessOpportunitiesBusinessDevelopment,
	BusinessOpportunitiesSalesOrCommercial,
	BusinessOpportunitiesInvestmentInStartups,
	BusinessOpportunitiesInvestmentInSmallMidBusinesses,
	BusinessOpportunitiesMentoring,
	BusinessOpportunitiesOperationalInvolvement,
	BusinessOpportunitiesFoundingNewOperations,
}

func (e BusinessOpportunities) IsValid() bool {
	switch e {
	case BusinessOpportunitiesAdvisoryRoles, BusinessOpportunitiesNonExecutiveOrExecutiveBoardPositions, BusinessOpportunitiesBusinessDevelopment, BusinessOpportunitiesSalesOrCommercial, BusinessOpportunitiesInvestmentInStartups, BusinessOpportunitiesInvestmentInSmallMidBusinesses, BusinessOpportunitiesMentoring, BusinessOpportunitiesOperationalInvolvement, BusinessOpportunitiesFoundingNewOperations:
		return true
	}
	return false
}

func (e BusinessOpportunities) String() string {
	return string(e)
}

func (e *BusinessOpportunities) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BusinessOpportunities(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BusinessOpportunities", str)
	}
	return nil
}

func (e BusinessOpportunities) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type InvolvementType string

const (
	InvolvementTypePassiveInvestment InvolvementType = "PASSIVE_INVESTMENT"
	InvolvementTypeActiveInvolvement InvolvementType = "ACTIVE_INVOLVEMENT"
)

var AllInvolvementType = []InvolvementType{
	InvolvementTypePassiveInvestment,
	InvolvementTypeActiveInvolvement,
}

func (e InvolvementType) IsValid() bool {
	switch e {
	case InvolvementTypePassiveInvestment, InvolvementTypeActiveInvolvement:
		return true
	}
	return false
}

func (e InvolvementType) String() string {
	return string(e)
}

func (e *InvolvementType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = InvolvementType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid InvolvementType", str)
	}
	return nil
}

func (e InvolvementType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OpportunityLocations string

const (
	OpportunityLocationsHomeCountry    OpportunityLocations = "HOME_COUNTRY"
	OpportunityLocationsCurrentCountry OpportunityLocations = "CURRENT_COUNTRY"
	OpportunityLocationsCrossBorder    OpportunityLocations = "CROSS_BORDER"
)

var AllOpportunityLocations = []OpportunityLocations{
	OpportunityLocationsHomeCountry,
	OpportunityLocationsCurrentCountry,
	OpportunityLocationsCrossBorder,
}

func (e OpportunityLocations) IsValid() bool {
	switch e {
	case OpportunityLocationsHomeCountry, OpportunityLocationsCurrentCountry, OpportunityLocationsCrossBorder:
		return true
	}
	return false
}

func (e OpportunityLocations) String() string {
	return string(e)
}

func (e *OpportunityLocations) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OpportunityLocations(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OpportunityLocations", str)
	}
	return nil
}

func (e OpportunityLocations) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProfessionalContacts string

const (
	ProfessionalContactsDecisionMakers    ProfessionalContacts = "DECISION_MAKERS"
	ProfessionalContactsBuyers            ProfessionalContacts = "BUYERS"
	ProfessionalContactsInvestors         ProfessionalContacts = "INVESTORS"
	ProfessionalContactsStrategicPartners ProfessionalContacts = "STRATEGIC_PARTNERS"
)

var AllProfessionalContacts = []ProfessionalContacts{
	ProfessionalContactsDecisionMakers,
	ProfessionalContactsBuyers,
	ProfessionalContactsInvestors,
	ProfessionalContactsStrategicPartners,
}

func (e ProfessionalContacts) IsValid() bool {
	switch e {
	case ProfessionalContactsDecisionMakers, ProfessionalContactsBuyers, ProfessionalContactsInvestors, ProfessionalContactsStrategicPartners:
		return true
	}
	return false
}

func (e ProfessionalContacts) String() string {
	return string(e)
}

func (e *ProfessionalContacts) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProfessionalContacts(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProfessionalContacts", str)
	}
	return nil
}

func (e ProfessionalContacts) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProjectDurationPreference string

const (
	ProjectDurationPreferenceLongTerm  ProjectDurationPreference = "LONG_TERM"
	ProjectDurationPreferenceShortTerm ProjectDurationPreference = "SHORT_TERM"
)

var AllProjectDurationPreference = []ProjectDurationPreference{
	ProjectDurationPreferenceLongTerm,
	ProjectDurationPreferenceShortTerm,
}

func (e ProjectDurationPreference) IsValid() bool {
	switch e {
	case ProjectDurationPreferenceLongTerm, ProjectDurationPreferenceShortTerm:
		return true
	}
	return false
}

func (e ProjectDurationPreference) String() string {
	return string(e)
}

func (e *ProjectDurationPreference) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProjectDurationPreference(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ProjectDurationPreference", str)
	}
	return nil
}

func (e ProjectDurationPreference) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ResidencyStatus string

const (
	ResidencyStatusCitizen             ResidencyStatus = "CITIZEN"
	ResidencyStatusPermanentResident   ResidencyStatus = "PERMANENT_RESIDENT"
	ResidencyStatusTemporaryResident   ResidencyStatus = "TEMPORARY_RESIDENT"
	ResidencyStatusWorkVisaHolder      ResidencyStatus = "WORK_VISA_HOLDER"
	ResidencyStatusStudentVisaHolder   ResidencyStatus = "STUDENT_VISA_HOLDER"
	ResidencyStatusTouristVisaHolder   ResidencyStatus = "TOURIST_VISA_HOLDER"
	ResidencyStatusRefugee             ResidencyStatus = "REFUGEE"
	ResidencyStatusAsylumSeeker        ResidencyStatus = "ASYLUM_SEEKER"
	ResidencyStatusUndocumented        ResidencyStatus = "UNDOCUMENTED"
	ResidencyStatusDiplomaticImmunity  ResidencyStatus = "DIPLOMATIC_IMMUNITY"
	ResidencyStatusInvestorVisaHolder  ResidencyStatus = "INVESTOR_VISA_HOLDER"
	ResidencyStatusRetireeVisaHolder   ResidencyStatus = "RETIREE_VISA_HOLDER"
	ResidencyStatusFamilySponsored     ResidencyStatus = "FAMILY_SPONSORED"
	ResidencyStatusEmploymentSponsored ResidencyStatus = "EMPLOYMENT_SPONSORED"
	ResidencyStatusSpecialStatus       ResidencyStatus = "SPECIAL_STATUS"
	ResidencyStatusStatelessPerson     ResidencyStatus = "STATELESS_PERSON"
	ResidencyStatusExile               ResidencyStatus = "EXILE"
	ResidencyStatusLongTermResident    ResidencyStatus = "LONG_TERM_RESIDENT"
	ResidencyStatusShortTermResident   ResidencyStatus = "SHORT_TERM_RESIDENT"
)

var AllResidencyStatus = []ResidencyStatus{
	ResidencyStatusCitizen,
	ResidencyStatusPermanentResident,
	ResidencyStatusTemporaryResident,
	ResidencyStatusWorkVisaHolder,
	ResidencyStatusStudentVisaHolder,
	ResidencyStatusTouristVisaHolder,
	ResidencyStatusRefugee,
	ResidencyStatusAsylumSeeker,
	ResidencyStatusUndocumented,
	ResidencyStatusDiplomaticImmunity,
	ResidencyStatusInvestorVisaHolder,
	ResidencyStatusRetireeVisaHolder,
	ResidencyStatusFamilySponsored,
	ResidencyStatusEmploymentSponsored,
	ResidencyStatusSpecialStatus,
	ResidencyStatusStatelessPerson,
	ResidencyStatusExile,
	ResidencyStatusLongTermResident,
	ResidencyStatusShortTermResident,
}

func (e ResidencyStatus) IsValid() bool {
	switch e {
	case ResidencyStatusCitizen, ResidencyStatusPermanentResident, ResidencyStatusTemporaryResident, ResidencyStatusWorkVisaHolder, ResidencyStatusStudentVisaHolder, ResidencyStatusTouristVisaHolder, ResidencyStatusRefugee, ResidencyStatusAsylumSeeker, ResidencyStatusUndocumented, ResidencyStatusDiplomaticImmunity, ResidencyStatusInvestorVisaHolder, ResidencyStatusRetireeVisaHolder, ResidencyStatusFamilySponsored, ResidencyStatusEmploymentSponsored, ResidencyStatusSpecialStatus, ResidencyStatusStatelessPerson, ResidencyStatusExile, ResidencyStatusLongTermResident, ResidencyStatusShortTermResident:
		return true
	}
	return false
}

func (e ResidencyStatus) String() string {
	return string(e)
}

func (e *ResidencyStatus) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ResidencyStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ResidencyStatus", str)
	}
	return nil
}

func (e ResidencyStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserType string

const (
	UserTypeFounder  UserType = "FOUNDER"
	UserTypeInvestor UserType = "INVESTOR"
)

var AllUserType = []UserType{
	UserTypeFounder,
	UserTypeInvestor,
}

func (e UserType) IsValid() bool {
	switch e {
	case UserTypeFounder, UserTypeInvestor:
		return true
	}
	return false
}

func (e UserType) String() string {
	return string(e)
}

func (e *UserType) UnmarshalGQL(v any) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserType", str)
	}
	return nil
}

func (e UserType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
