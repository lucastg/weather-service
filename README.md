# Weather Service
Este projeto é um serviço web desenvolvido em Go que recebe um CEP válido, identifica a cidade correspondente e retorna o clima atual com as temperaturas formatadas em Celsius, Fahrenheit e Kelvin.

## Requisitos do Sistema

Receber um CEP válido de 8 dígitos.
Consultar a localização usando a API ViaCEP.

Consultar as temperaturas da localização usando a API WeatherAPI.

Retornar as temperaturas nos seguintes formatos:

- Celsius
- Fahrenheit
- Kelvin

## Requisitos

- Go 1.23 ou superior
- Docker (se preferir executar dentro de um contêiner)

Se você deseja executar a ferramenta diretamente no seu ambiente local (sem Docker), siga os passos abaixo.

#### **Baixe o repositório**
Clone o repositório para sua máquina local:

```bash
https://github.com/lucastg/weather-service.git
```

### **Instale as dependências (opcional)**
```bash
go mod tidy
```

### **Configure suas credenciais de API:**
- WeatherAPI: Obtenha uma chave em [WeatherAPI](https://www.weatherapi.com/).
- Copie o arquivo .env.exemplo com o nome .env
```
WEATHERAPI_KEY=your_weather_api_key
```

### **Execute o programa**
```bash
go run main.go
```

### **Teste a API localmente**
- Endpoint: http://localhost:8080/weather/{cep}
- Exemplo: http://localhost:8080/weather/00000001


# **Executar com Docker**
Você também pode executar o sistema dentro de um contêiner Docker. Para isso, siga os passos abaixo.

### **Configure suas credenciais de API:**
- WeatherAPI: Obtenha uma chave em [WeatherAPI](https://www.weatherapi.com/).
- Copie o arquivo **.env.exemplo** com o nome **.env**
```
WEATHERAPI_KEY=your_weather_api_key
```

## **Construa a imagem Docker**
```bash
docker build -t weather-service .
```

## **Execute o container**
```bash
docker run -p 8080:8080 weather-service
```
### **Acesse o endpoint**
- Exemplo: http://localhost:8080/weather/00000001

# Google Cloud Run
```bash
curl --location https://weather-service-zplysilqxa-uc.a.run.app/weather/{cep}
```

# **Licença**
Este projeto está licenciado sob a [Licença MIT](LICENSE).
