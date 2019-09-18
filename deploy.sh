docker build -t rubendario/multi-server:latest -t rubendario/multi-server:$SHA .
docker push rubendario/multi-server:latest
docker push rubendario/multi-server:$SHA
kubectl apply -f k8s
kubectl set image deployments/server-deployment server=rubendario/multi-server:$SHA