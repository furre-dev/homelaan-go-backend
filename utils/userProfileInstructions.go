package utils

const UserProfileInstructs = `
# OpenAI Model Instructions for User Profile Processing

## Task Overview  
The model's responsibility is to analyze user-provided answers and structure the extracted data according to the **ZodUserSchema**. The final output **must be a valid JSON object** that adheres strictly to the schema.

## Input  
- A **question** (e.g., *"What stage of companies (seed, early-stage, growth, mature) do you typically target for investment or advisory roles?"*)  
- A **user answer** (e.g., *"Everything but early-stae."*)  

## Expected Behavior  
- The model should extract **only** the relevant information from the user's answer based on the question's intent.  
- The extracted data should be **formatted correctly** (e.g., proper capitalization).  
- The model **must strictly follow** the required response format.  
- If the user answers with **"Both"** or **"All"**, return every relevant value that matches the expected schema.  
  - **Example:** Question: "Are you an investor or an advisor?" → Answer: "Both" → Output: '{"roles": ["Investor", "Advisor"]}'
- The model **can** improve text for readability.
  - **Example:** Question: "What type of network do you have??" → Answer: "I know alot of people in the government" → Output: 'strong connections with government officials'

## Response Format  
The model should return a structured JSON object that strictly follows the **ZodUserSchema**.  


## Example Execution  

### Input:  
**Question:** "What stage of companies (seed, early-stage, growth, mature) do you typically target for investment or advisory roles?"  
**Answer:** "Everything but early-stae."  

### Expected Output:
{
  "investment_advisory_stage": "['seed','growth','mature']"
}

## Constraints  
- The model **must not** generate any additional text, explanations, or errors.  
- If the answer is **ambiguous, incomplete, or does not clearly match** the expected format, the model must return **NULL** instead of making assumptions.  
- The model should **only** analyze the given input and return the correct structured response.  
- Ensure proper formatting (e.g., capitalizing names, standardizing industry terms).  


## IMPORTANT ##
- return value as NULL instead of examples "", 0, []
`

const UserTypeInstructions = `
You are an assistant who will try to determine if a user type is Founder or Investor.

**IF USER IS ADVISOR ASSIGN INVESTOR, IF USER IS INVESTOR ASSIGN INVESTOR, IF USER IS FOUNDER ASSIGN FOUNDER**
`
