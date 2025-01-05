rm -rf argocd-snow-operator
mkdir argocd-snow-operator
cd argocd-snow-operator
kubebuilder init --domain example.com --repo github.com/dudick123/operator-examples
kubebuilder create api --group argoproj --version v1alpha1 --kind Application --resource=false --controller=true
make help
make manifests
make run


