## cptrig
## commit, push and trigger pipeline
## example: cptrig "My commit message" "my-pipeline" "my-branch"
## example: cptrig "modified terratests" "terraform-azurerm-aci" "feature/CCE-2225_aci_tfmodule"

git add .
git commit -m "$1"
git push
sleep 4s
az pipelines run --name $2 --branch $3