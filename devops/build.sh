echo "create namespace open-ldap "
kubectl create ns open-ldap
echo  "labeling one of nodes"
nodeName=$(kubectl get node  | grep "<none>" | awk 'NR==1' | awk '{print $1 }')
kubectl label node $nodeName openldap: "true"
#echo "create secret openldap-config -- openldap-config.txt  "
#kubectl create secret generic openldap-config --from-file=./openldap-config.txt -n open-ldap
#echo "create configmap openldap-init -- init.ldif "
#kubectl create configmap openldap-init --from-file=./init.ldif -n open-ldap
echo "create openldap deployment"
kubectl apply -f openldap-deployment.yaml
echo "create openldap-phpadmin deployment"
kubectl apply -f openldap-phpadmin-deployment.yaml



