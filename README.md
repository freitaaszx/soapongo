
# 🧼 SOAP API — Go

API SOAP simples desenvolvida em Go, sem dependências externas.

## 📋 Pré-requisitos

- Go 1.22+

## 🚀 Como rodar

```bash
git clone https://github.com/seu-usuario/soap-api
cd soap-api
go run ./cmd/api
```

## 📡 Endpoint

```
POST http://localhost:8080/soap
Content-Type: text/xml
```

## 🧪 Exemplo de requisição

```xml
<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <HelloRequest>
      <Name>João</Name>
    </HelloRequest>
  </soap:Body>
</soap:Envelope>
```

## 📁 Estrutura

```
soap-api/
├── cmd/api/main.go
├── internal/
│   ├── handler/
│   └── models/
└── go.mod
```

## 📄 Licença

MIT
