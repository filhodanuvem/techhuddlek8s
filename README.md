# techhuddlek8s

# Preparando o ambiente 


Instalando cluster kubernetes local (melhor rodar isso antes da call)
```
brew install socket_vmnet
brew tap homebrew/services
brew install qemu
brew install minikube
```

Criando rede local, cluster com 3 nodes e habilitando o uso de um registry local 
```
sudo brew services start socket_vmnet
minikube start --nodes 3 -p local --memory="2GB" --cpus=2 --vm=true 
```

Para deletar o cluster 

```
minikube delete --all
```

💡 Nesse ponto temos um cluster 3 máquinas de 2Gb e 2 cores cada.

1) falar do control-plane e do kubeapi
2) falar para que servem os contextos e mostrar a url da api do contexto local
3) curl --insecure -X GET https://127.0.0.1:60621   /livez
4) falar do kubectl config get-contexts, kubectx  e `current-context`. 

# Cap 1 criação de pods 

0) kubectl apply -f ./1/pod.yaml
1) kubectl delete -f ./1/pod.yaml

# Cap 2 request de recursos 

0) kubectl apply -k ./2
1) 1 pod pending, mostrar events
2) mudar pra 800Mi um pod
3) mensagem muda, 3 nodes com insuficient cpu mas so 1 com insuficient memory
4) mudar para 0.1 cpu
5) recursos imutaveis

# Cap 3 deployments 

0) porque existe deploy
1) replicaset 

# cap 4 aplicacao que consome muita memoria

0) docker build . -t cloudson/overload 
1) docker push cloudson/overload 
2) kubectl apply -k ./5
3) kubectl logs -f -l app=overload

# cap 5 taints e tolerations

0) kubectl taint nodes local-m02 time=credito:NoSchedule
1) aumentar pods do cap3 
2) aplicar 5