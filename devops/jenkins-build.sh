echo "create namespace jenkins "
kubectl create ns jenkins
echo  "labeling one of nodes"
nodeName=$(kubectl get node  | grep "<none>" | awk 'NR==1' | awk '{print $1 }')
kubectl label node $nodeName openldap: "true"
#echo "create secret openldap-config -- openldap-config.txt  "
#kubectl create secret generic openldap-config --from-file=./openldap-config.txt -n open-ldap
#echo "create configmap openldap-init -- init.ldif "
#kubectl create configmap openldap-init --from-file=./init.ldif -n open-ldap
echo "create jenkins deployment"
kubectl apply -f ./jenkins/jenkins-deployment.yaml




