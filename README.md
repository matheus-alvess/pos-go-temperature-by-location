## Documentação da API de Clima

### Visão Geral
Esta API permite obter informações sobre a temperatura com base no CEP fornecido. A API utiliza o framework Gin para lidar com as requisições HTTP e um serviço de backend para recuperar os dados de temperatura.

### Endpoints

#### Obter Clima pelo CEP

**Descrição:** Este endpoint retorna as informações de temperatura para um CEP específico.

- **URL:** `/weather/:cep`
- **Método:** `GET`
- **Parâmetros da URL:**
    - `cep` (string): O CEP deve ser composto por 8 dígitos.

- **Respostas:**
    - **200 OK:** Retorna as informações de temperatura.
    - **422 Unprocessable Entity:** O CEP fornecido é inválido.
    - **404 Not Found:** O CEP não foi encontrado.
    - **500 Internal Server Error:** Ocorreu um erro ao tentar obter a temperatura.

### Exemplo de Uso

#### Requisição

```bash
curl -X GET "http://localhost:8080/weather/12345678"
```

Respostas

- 200 OK

```json
{
  "TempC": 25.5,
  "TempF": 50.6,
  "TempK": 77.9
}
```

- 422 Unprocessable Entity

```json
{
  "message": "invalid zipcode"
}
```

- 404 Not Found

```json
{
  "message": "can not find zipcode"
}
```

- 500 Internal Server Error
```json
{
  "message": "failed to get temperature"
}
```

## Como Executar Testes Unitários

Para rodar os testes unitários da aplicação, use o comando:

```sh
go test ./...
```

### Como Executar a Aplicação com Docker Compose

Se você preferir usar o Docker Compose para gerenciar os contêineres da aplicação, siga estas instruções:

1. **Construa e Execute os Contêineres:**

   Certifique-se de que você está no mesmo diretório que o arquivo `docker-compose.yml` e execute o seguinte comando:

```bash
   docker-compose up --build
```

Isso irá construir as imagens Docker necessárias e iniciar os contêineres da aplicação. O sinalizador --build garante que as imagens sejam construídas novamente se houver alterações nos arquivos de código-fonte.

### Como Executar a Aplicação no Docker

Para executar a aplicação no Docker, primeiro você precisa ter o Docker instalado em sua máquina.

1. **Construa a Imagem Docker:**

   Abra o terminal e navegue até o diretório onde se encontra o arquivo `Dockerfile` e execute o seguinte comando:

```bash
   docker build -t pos-go-temperature-by-location .
```

Isso irá criar uma imagem Docker com o nome pos-go-temperature-by-location com base no conteúdo do diretório atual.

2. **Execute o Contêiner Docker:**

Depois de construir a imagem Docker, você pode executar um contêiner a partir dela usando o seguinte comando:

```bash
docker run -p 8080:8080 -e WEATHER_API_KEY=0221a1d62222490882322259242305 pos-go-temperature-by-location
```

Isso irá iniciar um contêiner Docker com a aplicação em execução na porta 8080.

**Nota**: Certifique-se de substituir **`0221a1d62222490882322259242305`** pelo seu próprio API key, se necessário.

