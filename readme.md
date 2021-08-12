# Go Challenge 
Solução de cadastro de parceiros com integração ViaCep e Google GeoCode, projeto baseado no desafio do [Zé Delivery](https://github.com/ZXVentures/ze-code-challenges/blob/master/backend_pt.md).

### Tecnologias Utilizadas
1. [Golang](https://golang.org/)
2. [gRPC](https://grpc.io/)
3. [GraphQL](https://graphql.org/)
4. [Microsoft Azure](https://azure.microsoft.com/pt-br/features/azure-portal/)
5. [Azure Cosmos DB](https://azure.microsoft.com/pt-br/services/cosmos-db/)
6. [Azure Container Registry](https://azure.microsoft.com/en-us/services/container-registry/)
7. [Azure Kubernetes Service](https://azure.microsoft.com/pt-br/services/kubernetes-service/)
8. [Github Actions](https://docs.github.com/pt/actions)
9. [Docker](https://www.docker.com/)
10. [Terraform](https://www.terraform.io/)

### Desenho da Solução
<p align="center">
  <img src=".docs/Go-Challenge.png" width="800" title="Main">
</p>

## Funcionalidades
1. Criar um parceiro
2. Parceiros
3. Parceiros por ID
4. Parceiros por localização (distância por kilômetro)

## Executando local com docker
1. Executar o seguinte comando.
   ```
   docker-compose up -d --build
   ```
   
## Infra as code (terraform)
1. Iniciando o terraform. na pasta terraform executar:
   ```
   terraform init
   ```
2. Aplicar a formatação
   ```
   terraform fmt
   ```
3. Validação do que foi criado
   ```
   terraform validate
   ```
4. Aplicar o plano (planejamento dos recursos que serão criados)
   ```
   terraform plan
   ```
5. Executar a criação da infra
   ```
   terraform apply -auto-approve
   ```

## Microsoft Azure
1. Autenticação no azure (Precisamos instalar o ([Azure CLI](https://docs.microsoft.com/pt-br/cli/azure/install-azure-cli))
   ```
   az login
   ```
2. Obtendo credenciais do cluster AKS 
   ```
   az aks get-credentials --resource-group $RESOURCE_GROUP --name $NAME
   ```
3. Build da imagem docker
   ```
   docker image build -t partnersregistry.azurecr.io/address:latest -f address/Dockerfile .

   docker image build -t partnersregistry.azurecr.io/partner:latest -f partner/Dockerfile .
   ```
4. Autenticação no ACR
   ```
   docker login partnersregistry.azurecr.io -u $USERNAME -p $PASSWORD
   ```
5. Upload da imagem no ACR
   ```
   docker image push partnersregistry.azurecr.io/address:latest

   docker image push partnersregistry.azurecr.io/partner:latest
   ```
## Kubernetes
1. Aplicar os manifestos do k8s
   ```
   kubectl apply -f .\.k8s\namespaces\ -R
   kubectl apply -f .\.k8s\deployments\ -R -n partner
   kubectl apply -f .\.k8s\services\ -R -n partner
   ```

## gRPC e GraphQL   
1. gRPC
    ```
    protoc --go_out=address/pb --go_opt=paths=source_relative --go-grpc_out=address/pb --go-grpc_opt=paths=source_relative --proto_path=address/protofiles address/protofiles/*.proto
    ```

2. GraphQL
   ```
   go get github.com/99designs/gqlgen
   ```
   ```
   go run github.com/99designs/gqlgen init
   ```
   ```
   go run github.com/99designs/gqlgen generate
   ```
   