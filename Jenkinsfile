pipeline {
    
    agent any

    tools {
        go 'go1.14'
    }
    
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        DH_PAT = credentials('JaySAP_DH_PAT')
        GH_WEB_APP_REPO = 'git@github.com:jaythamke/web-app.git'
        WEB_APP_DIR = "${JENKINS_HOME}/web-app"
    }

    stages {
        stage('Compile') {
            steps {
                sh """
                echo 'Cloning repository...'
                git clone $GH_WEB_APP_REPO $WEB_APP_DIR
                cd $WEB_APP_DIR
                echo 'Compiling golang application...'
                go build ./bin/go-web-app
                """
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
                sh """
                #! /bin/bash
                sleep 2
                echo "This is a simulation of testing!"
                """
            }
        }
        stage('Build Image') {
            steps {
                sh """
                echo 'Building Docker image...'
                cd $WEB_APP_DIR
                export APP_REVISION=$(cat REVISION)
                docker build . --tag go-web-app:$APP_REVISION
                docker images
                """
            }
        }
        stage('Push') {
            steps {
                echo 'Push image to docker hub....'
            }
        }
    }
}