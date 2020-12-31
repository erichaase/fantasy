# Quickstart
### Google Kubernetes Engine
* Enable GKE API: `gcloud services enable container.googleapis.com`
* Create GKE cluster: `gcloud container clusters create fantasy --project erichaase-fantasy --zone us-central1-c --num-nodes 1 --machine-type g1-small --release-channel regular --disk-size 32GB`

### Google Cloud Build
* Enable GCB and GCR APIs: `gcloud services enable cloudbuild.googleapis.com containerregistry.googleapis.com`
* Grant GCB SA k8s access: `gcloud projects add-iam-policy-binding erichaase-fantasy --member=serviceAccount:822045923302@cloudbuild.gserviceaccount.com --role=roles/container.developer`
* [Create GCB trigger(s) via UI](https://console.cloud.google.com/cloud-build/triggers/add)