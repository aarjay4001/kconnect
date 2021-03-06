#!/bin/bash  

set -e 

echo "creating directory kconnect"
mkdir -p kconnect
cd kconnect

latest_kconnect_release_tag=$(curl --silent "https://api.github.com/repos/fidelity/kconnect/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
latest_helm_release_tag=$(curl --silent "https://api.github.com/repos/helm/helm/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
echo "kconnect version: $latest_kconnect_release_tag"

if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # linux
    arch=$(dpkg --print-architecture)

    kconnect_url=$(echo "https://github.com/fidelity/kconnect/releases/download/TAG/kconnect_linux_ARCH.tar.gz" | sed "s/TAG/$latest_kconnect_release_tag/" | sed "s/ARCH/$arch/" )
    kubectl_url=$(echo "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/ARCH/kubectl" | sed "s/ARCH/$arch/" )
    helm_url=$(echo "https://get.helm.sh/helm-TAG-linux-ARCH.tar.gz" | sed "s/TAG/$latest_helm_release_tag/" | sed "s/ARCH/$arch/" )
    aws_iam_authenticator_url=$(echo "https://amazon-eks.s3.us-west-2.amazonaws.com/1.18.9/2020-11-02/bin/linux/ARCH/aws-iam-authenticator" | sed "s/ARCH/$arch/" )

    # download 
    curl -s -L $kconnect_url -o kconnect.tar.gz
    curl -s -LO $kubectl_url
    curl -s -L $helm_url -o helm.tar.gz
    curl -s -L $aws_iam_authenticator_url -o aws-iam-authenticator

    # unzip
    tar -xf kconnect.tar.gz
    tar -xf helm.tar.gz
    mv linux-*/helm .

    # cleanup
    rm -f kconnect.tar.gz
    rm -f helm.tar.gz
    rm -rf linux-*

    # permissions
    chmod +x kubectl
    chmod +x aws-iam-authenticator

elif [[ "$OSTYPE" == "darwin"* ]]; then
    # Mac OSX
    kconnect_url=$(echo "https://github.com/fidelity/kconnect/releases/download/TAG/kconnect_macos_amd64.tar.gz" | sed "s/TAG/$latest_kconnect_release_tag/" )
    kubectl_url="https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/darwin/amd64/kubectl"
    helm_url=$(echo "https://get.helm.sh/helm-TAG-darwin-amd64.tar.gz" | sed "s/TAG/$latest_helm_release_tag/" )
    aws_iam_authenticator_url="https://amazon-eks.s3.us-west-2.amazonaws.com/1.18.9/2020-11-02/bin/darwin/amd64/aws-iam-authenticator"

    # download 
    curl -s -L $kconnect_url -o kconnect.tar.gz
    curl -s -LO $kubectl_url
    curl -s -L $helm_url -o helm.tar.gz
    curl -s -L $aws_iam_authenticator_url -o aws-iam-authenticator

    # unzip
    tar -xf kconnect.tar.gz
    tar -xf helm.tar.gz
    mv darwin-*/helm .

    # cleanup
    rm -f kconnect.tar.gz
    rm -f helm.tar.gz
    rm -rf darwin-*

    # permissions
    chmod +x kubectl
    chmod +x aws-iam-authenticator

elif [[ "$OSTYPE" == "msys" ]]; then
    # Win git bash
   
    kconnect_url=$(echo "https://github.com/fidelity/kconnect/releases/download/TAG/kconnect_windows_amd64.zip" | sed "s/TAG/$latest_kconnect_release_tag/" )
    kubectl_url=$(echo "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/windows/amd64/kubectl.exe" )
    helm_url=$(echo "https://get.helm.sh/helm-TAG-windows-amd64.zip" | sed "s/TAG/$latest_helm_release_tag/" )
    aws_iam_authenticator_url="https://amazon-eks.s3.us-west-2.amazonaws.com/1.18.9/2020-11-02/bin/windows/amd64/aws-iam-authenticator.exe"

    # download 
    curl -s -L $kconnect_url -o kconnect.zip
    curl -s -LO $kubectl_url
    curl -s -L $helm_url -o helm.zip
    curl -s -L $aws_iam_authenticator_url -o aws-iam-authenticator.exe

    # unzip
    unzip kconnect.zip
    unzip helm.zip
    mv windows-amd64/helm.exe .

    # cleanup
    rm -f kconnect.zip
    rm -f helm.zip
    rm -rf windows-amd64

fi
