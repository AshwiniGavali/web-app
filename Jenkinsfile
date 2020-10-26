pipeline {
    tools {
        go 'go1.14'
    }
    
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }

    stages {
        stage('Compile') {
            steps {
                sh """
                echo 'Building..'

                """
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
                sh """
                #! /bin/bash
                sleep 10
                echo "This is a simulation of testing!"
                """
            }
        }
        stage('Build Image') {
            steps {
                echo 'Building Docker image....'
            }
        }
        stage('Push') {
            steps {
                echo 'Push image to docker hub....'
            }
        }
    }
}