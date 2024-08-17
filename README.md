Dynamic Rule Engine
Overview
The Dynamic Rule Engine is a 3-tier application designed to determine user eligibility based on various attributes using Abstract Syntax Trees (ASTs) to represent conditional rules. The system supports dynamic creation, combination, and modification of rules.

Features
Create Rules: Converts rule strings into AST representations.
Combine Rules: Merges multiple rule ASTs into a single combined AST.
Evaluate Rules: Evaluates combined rules against provided data to determine eligibility.
Components
Backend: Implements the logic for rule creation, combination, and evaluation.
API: Provides endpoints for interacting with the rule engine.
UI: A simple interface for users to interact with the API.
API Endpoints

1. Create Rule
Endpoint: /api/v1/create_rule
Method: POST
Request Body:
json
Copy code
{
  "rule_string": "((age > 30 AND department = 'Sales') OR (age < 25 AND department = 'Marketing')) AND (salary > 50000 OR experience > 5)"
}
Response:
json
Copy code
{
  "AST": {
    "Type": "Operator",
    "Value": "AND",
    "Left": {
      "Type": "Operator",
      "Value": "OR",
      "Left": { /* Sub-node */ },
      "Right": { /* Sub-node */ }
    },
    "Right": {
      "Type": "Operator",
      "Value": "OR",
      "Left": { /* Sub-node */ },
      "Right": { /* Sub-node */ }
    }
  }
}

2. Combine Rules
Endpoint: /api/v1/combine_rules
Method: POST
Request Body:
json
Copy code
[
  "((age > 30 AND department = 'Sales') OR (age < 25 AND department = 'Marketing')) AND (salary > 50000 OR experience > 5)",
  "((age > 30 AND department = 'Marketing')) AND (salary > 20000 OR experience > 5)"
]
Response:
json
Copy code
{
  "AST": {
    "Type": "Operator",
    "Value": "AND",
    "Left": { /* Combined AST from rules */ },
    "Right": { /* Combined AST from rules */ }
  }
}

3. Evaluate Rule
Endpoint: /api/v1/evaluate_rule
Method: POST
Request Body:
json
Copy code
{
  "ast": {
    "Type": "Operator",
    "Value": "AND",
    "Left": { /* AST nodes */ },
    "Right": { /* AST nodes */ }
  },
  "data": {
    "age": 35,
    "department": "Sales",
    "salary": 60000,
    "experience": 3
  }
}
Response:
json
Copy code
{
  "result": true
}


Running the Application
Build and Run: Compile the Go application and run it.

insall go 1.19 or above
clone repo :- git clone https://github.com/dushyant1435/Dynamic-Rule-Engine.git

run following commands on terimal :-
1.go mod tidy
2.go run main.go

<!-- go server started -->

3.open index.html in live server it


Access the API: The API will be available at http://localhost:8080.

Use the UI: Open the provided HTML file in a web browser to interact with the API through a simple interface.

Dependencies
Go 1.20+
Required Go modules