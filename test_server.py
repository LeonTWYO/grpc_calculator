import requests

base_url = "http://localhost:8080"

def add(operand1, operand2):
    url = base_url + "/v1/calculator/add"
    payload = {"operand1": operand1, "operand2": operand2}
    response = requests.post(url, json=payload)
    return response.json()

def multiply(factor1, factor2):
    url = base_url + "/v1/calculator/multiply"
    payload = {"factor1": factor1, "factor2": factor2}
    response = requests.post(url, json=payload)
    return response.json()

def divide(dividend, divisor):
    url = base_url + "/v1/calculator/divide"
    payload = {"dividend": dividend, "divisor": divisor}
    response = requests.post(url, json=payload)
    return response.json()

# Test the add endpoint
add_response = add(5, 3)
print("Add Response:", add_response)

# Test the multiply endpoint
multiply_response = multiply(4, 7)
print("Multiply Response:", multiply_response)

# Test the divide endpoint
divide_response = divide(10, 2)
print("Divide Response:", divide_response)
