echo "create envoy configmap "
kubectl create configmap envoy-config --from-file=envoy.yaml
echo "create envoy deployment "
kubectl apply -f https://github.com/cncamp/101/blob/master/module4/envoy-deploy.yaml
kubectl expose deploy envoy --selector run=envoy --port=10000 --type=NodePort

echo "list process envoy and config "
containerId=$(kubectl get pod -n default | grep "Running" | awk '{print $1 }')
kubectl exec -it $containerId -n default bash -- ps -ef
kubectl exec -it $containerId -n default bash -- cat /etc/envoy/envoy.yaml

echo "show envoy net config and mock request "
kubectl get svc -owide | grep "envoy"
ip=$(kubectl get svc -owide | grep "envoy"| awk '{print $3 }')
tcp_ip=$(kubectl get svc -owide | grep "envoy" | awk '{print $5 }')
port=${tcp_ip%:*}
echo $ip:$port
curl $ip:$port -w '\n'

echo "update deployment envoy set port 23456 and show net config "
kubectl delete configmaps envoy-config
cd update
kubectl create configmap envoy-config --from-file=envoy.yaml
cd ../
kubectl rollout restart deployment envoy -n default
kubectl delete svc envoy
kubectl expose deploy envoy --selector run=envoy --port=23456 --type=NodePort
kubectl get svc -owide | grep "envoy"
ip=$(kubectl get svc -owide | grep "envoy" | awk '{print $3 }')
port=$(kubectl get svc -owide | grep "envoy" | awk '{print $5 }')
port=${port%:*}
echo $ip:$port
curl $ip:$port -w '\n'

echo "delete deploy "
deploy=$(kubectl get deployments.apps -n default -owide | grep "run=envoy" | awk '{print $3 }')
kubectl delete deployments.apps $deploy -n default --cascade=orphan







