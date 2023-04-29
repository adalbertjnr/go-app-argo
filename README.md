# Projeto CI/CD

Este projeto utiliza diversas ferramentas para implementar uma pipeline de CI/CD.

- Pipeline CI/CD utilizando o GitHub Actions.
- Argo para etapa de CD.
- Utilizado Docker para construção da imagem da aplicação. Docker Compose para testes da aplicação com o banco de dados postgres.
- Cluster Kubernetes para gerenciar as imagens construídas da aplicação na etapa anterior. 
- Template do helm criado do zero com Deployment, Service, HPA e IngressRoutes -- utilizado o Traefik como proxy reverso.
- Monitoramento da aplicação utilizando o Prometheus e o Grafana.

Além disso, este repositório também contém arquivos de configuração do cluster e testes com o Docker Compose.

E também há outro repositório de configuração que é monitorado pelo Argo.

<br>

**Repositório com o template do helm que está sendo monitorado pelo ArgoCD:** 
```
https://github.com/souzagmu/go-app-argo-cfg
```
<br>

> Ainda faltam outras ferramentas que irei implementar
> - Sonarqube na pipeline para análise do código
> - ELK Stack

<br>
<br>

![ARGO](https://user-images.githubusercontent.com/123520913/235320688-bab309be-1967-4ad0-9d62-a648944c06ff.png)

