import requests


if __name__ == "__main__":
    url = "http://localhost:8080/convert"
    value = float(input("Digite o valor da temperatura: "))
    unit = input("Digite a unidade de temperatura: ")
        
    params = {
        "value": value,
        "unit": unit
    }
        
    response = requests.post(url, json=params)
        
    if response.ok:
        print("Resposta do servidor:", response.json())
            
    else:
        print("Erro na requisição:", response.status_code)
    
    