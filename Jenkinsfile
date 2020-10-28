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
        GH_WEB_APP_REPO = 'https://github.com/jaythamke/web-app.git'
        WEB_APP_DIR = "${JENKINS_HOME}/web-app"
        DH_REPO = "jayeshthamkesap"
        REVISION_DATA = "/tmp/REVISION"
    }

    stages {
        stage('Compile') {
            steps {
                sh '''
                #! /bin/bash
                echo 'Cloning repository...'
                rm -rf $WEB_APP_DIR
                git clone $GH_WEB_APP_REPO $WEB_APP_DIR
                cd $WEB_APP_DIR
                echo 'Compiling golang application...'
                go build -o ${WEB_APP_DIR}/bin/go-web-app
                '''
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
                sh '''
                #! /bin/bash
                sleep 2
                echo "This is a simulation of testing!"
                '''
            }
        }
        stage('Build Image') {
            steps {
                sh '''
                #! /bin/bash
                echo 'Building Docker image...'
                cd $WEB_APP_DIR
                APP_REVISION=$(cat REVISION)
                cat REVISION > $REVISION_DATA
                docker build . --tag go-web-app:$APP_REVISION
                docker images
                '''
            }
        }
        stage('Push') {
            steps {
                sh '''
                #! /bin/bash
                echo 'Preparing for publish image to docker hub...'
                cd $WEB_APP_DIR
                echo 'Retagging the image...'
                APP_REVISION=$(cat $REVISION_DATA)
                docker tag go-web-app:$APP_REVISION $DH_REPO/web-app:$APP_REVISION
                echo 'Pushing image to docker hub...'
                docker login --username $DH_REPO --password $DH_PAT
                docker push $DH_REPO/web-app:$APP_REVISION
                '''
            }
        }
    }
}