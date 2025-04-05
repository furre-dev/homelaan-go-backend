package interview

import "github.com/furre-dev/homelaan-go-backend/graph/model"

var Questions = []model.Question{
	{
		QuestionTitle: "Where are you currently living (city and country)?",
		ProfileField:  "geo_info.location",
		QuestionIndex: 0,
	},
	{
		QuestionTitle: "How long have you lived abroad?",
		ProfileField:  "geo_info.years_abroad",
	},
	{
		QuestionTitle: "Do you plan to stay long-term in your current country, or do you have plans to relocate?",
		ProfileField:  "geo_info.will_stay_long_term_in_current_location",
	},
	{
		QuestionTitle: "Do you have family living with you abroad (spouse, children)?",
		ProfileField:  "geo_info.has_family_abroad",
	},
	{
		QuestionTitle: "Are you a citizen/permanent resident in your current country?",
		ProfileField:  "geo_info.residency_status",
	},
	{
		QuestionTitle: "What is your current role and in which organization?",
		ProfileField:  "[professional_background.current_role, professional_background.organization, professional_background.industry]",
	},
	{
		QuestionTitle: "Briefly describe your professional background (roles, companies, career milestones).",
		ProfileField:  "professional_background.career_overview",
	},
	{
		QuestionTitle: "What industries/sectors do you have strong experience in?",
		ProfileField:  "professional_background.industry",
	},
	{
		QuestionTitle: "What functional areas do you have deep domain knowledge in (e.g., legal, finance, HR, engineering, operations, sales, marketing, product, strategy, etc.)?",
		ProfileField:  "professional_background.expertise_areas",
	},
	{
		QuestionTitle: "What is your seniority level and decision-making scope in your current role?",
		ProfileField:  "[professional_background.seniority_level, professional_background.decision_making_scope]",
	},
	{
		QuestionTitle: "In which industries do you have a strong professional network?",
		ProfileField:  "network_market_access.industries_with_strong_network",
	},
	{
		QuestionTitle: "Do you have access to decision-makers, buyers, investors, or strategic partners in your current country or industry?",
		ProfileField:  "network_market_access.professional_contacts",
	},
	{
		QuestionTitle: "Are you connected to specific ethnic, business, or professional diaspora groups locally?",
		ProfileField:  "network_market_access.local_business_diaspora_groups",
	},
	{
		QuestionTitle: `Are you connected to specific ethnic, business, or professional diaspora groups locally? Please write all of them that apply! 
										Advisory roles
										Non-executive / Executive board positions
										Business development / Opening doors
										Commercial partnerships / Sales representation
										Investment in startups
										Investment in small/mid-sized businesses (SMBs)
										Mentoring entrepreneurs
										Operational involvement in ventures (interim leadership, turnaround, model
										Co-founding new ventures
	`,
		ProfileField: "engagement_preferences.business_opportunities",
	},
	{
		QuestionTitle: "What sectors or opportunity types are you particularly interested in supporting?",
		ProfileField:  "engagement_preferences.sectors_of_interest",
	},
	{
		QuestionTitle: "Are you interested in opportunities based in your home country, your current country, or cross-border plays?",
		ProfileField:  "engagement_preferences.opportunity_locations",
	},
	{
		QuestionTitle: "Are you an active investor in startups or SMBs?",
		ProfileField:  "investment_appetite.active_investor_in_startups_or_small_mid_businesses",
	},
	{
		QuestionTitle: "What is your typical investment range?",
		ProfileField:  "investment_appetite.investment_range",
	},
	{
		QuestionTitle: "Are you part of any angel groups or investment syndicates?",
		ProfileField:  "investment_appetite.angel_group_membership",
	},
	{
		QuestionTitle: "Are you open to passive investment, active involvement, or both?",
		ProfileField:  "investment_appetite.involvement_type",
	},
	{
		QuestionTitle: "How much time are you willing to commit monthly for advisory or board-related work?",
		ProfileField:  "availability_and_commitment.monthly_commitment_hours",
	},
	{
		QuestionTitle: "Are you open to short-term projects, or only long-term engagements?",
		ProfileField:  "availability_and_commitment.project_duration_preference",
	},
	{
		QuestionTitle: "Would you consider travel for business engagement if needed?",
		ProfileField:  "availability_and_commitment.willingness_to_travel",
	},
	{
		QuestionTitle: "Is your spouse or partner interested in also being profiled for business opportunities?",
		ProfileField:  "spousal_involvement.spouse_interested",
	},
	{
		QuestionTitle: "If yes, what is their professional background and area of expertise?",
		ProfileField:  "[spousal_involvement.spouse_professional_background, spousal_involvement.spouse_expertise_areas]",
	},
}
