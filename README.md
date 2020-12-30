# Quickstart
## Google Kubernetes Engine
* Enable GKE API: `gcloud services enable container.googleapis.com`
* Create GKE cluster: `gcloud container clusters create fantasy --project erichaase-fantasy --zone us-central1-c --num-nodes 1 --machine-type g1-small --release-channel regular --disk-size 32GB`

## Google Cloud Build
* Enable GCB and GCR APIs: `gcloud services enable cloudbuild.googleapis.com containerregistry.googleapis.com`
* [Create GCB trigger(s) via UI](https://console.cloud.google.com/cloud-build/triggers/add)