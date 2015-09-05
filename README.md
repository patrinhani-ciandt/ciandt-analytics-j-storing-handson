#CI&T Analytics Journey - Big Table Hands On

## Pre req.
* Set up a [service account](https://cloud.google.com/storGoogle Cloud JSON key to the directory agePutda/authentication#service_accounts), and download a JSON key file for the account:

    1- In the [Developers Console](https://console.developers.google.com), open your project by clicking on the project name.

    2- In the left sidebar, click APIs & auth, then Credentials.

    3- Under OAuth, click Create new Client ID.

    4- Select Service account, then click Create Client ID.

    5- Read the confirmation dialog, then click Okay, got it. A key is downloaded automatically in JSON format. Keep the JSON key in a safe place.

## Initializing the container

    $ docker run -t -i --name=analytics-j-storing-handson patrinhani/ciandt-analytics-j-storing-handson bash

    for Windows:
    - Create a folder called "DockerHome" on  c:\Users\[username]
    $ docker run --rm -t -i -e WHOAMI=$(whoami) --name=analytics-j-img-test -v //c/Users/[username]/DockerHome:/home  patrinhani/ciandt-analytics-j-storing-handson

    for Linux / MacOS:
    - Create a folder called "DockerHome" on  ~/
    $ docker run --rm -t -i -e WHOAMI=$(whoami) --name=analytics-j-img-test -v ~/DockerHome:/home  patrinhani/ciandt-analytics-j-storing-handson

## Configuring your Google Cloud Account into the container

Run the commands bellow on your container

    $ gcloud auth login
    $ gcloud config set project [PROJECT_ID]

## Be fun !
