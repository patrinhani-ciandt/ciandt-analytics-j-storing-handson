#CI&T Analytics Journey - Big Table Hands On

## Pre req.
* Set up a [service account](https://cloud.google.com/storGoogle Cloud JSON key to the directory agePutda/authentication#service_accounts), and download a JSON key file for the account:

    1- In the [Developers Console](https://console.developers.google.com), open your project by clicking on the project name.

    2- In the left sidebar, click APIs & auth, then Credentials.

    3- Under OAuth, click Create new Client ID.

    4- Select Service account, then click Create Client ID.

    5- Read the confirmation dialog, then click Okay, got it. A key is downloaded automatically in JSON format. Keep the JSON key in a safe place.

## Initializing the container

    $ docker run -t -i --name=gcloud-golang-dev patrinhani/gcloud-golang bash

## Configuring your Google Cloud Account into the container

Run the commands bellow on your container

    $ gcloud auth login
    $ gcloud config set project [PROJECT_ID]

## Be fun !
