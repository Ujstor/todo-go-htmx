pipeline {
    agent any

    environment {
        GITHUB_USER = 'ujstor'
        GITHUB_REPO = 'todo-go'
        DOCKER_HUB_USERNAME = 'ujstor'
        DOCKER_REPO_NAME = 'todo-go'
        BRANCH = 'master'
        VERSION_PART = 'Patch' // Patch, Minor, Major
        DOCKER_JENKINS_CERDIDENTALS_ID = 'be9636c4-b828-41af-ad0b-46d4182dfb06'
        TAG = '' // Generated automatically
    }

    stages {
        stage('Checkout Code') {
            steps {
                script {
                    git(url: "https://github.com/${GITHUB_USER}/${GITHUB_REPO}/", branch: env.BRANCH_NAME)
                }
            }
        }

        stage('chmod Tag sh') {
            steps {
                script {
                    sh "chmod 777 ${WORKSPACE}/docker_tag.sh"
                }
            }
        }

        stage('Generate Docker Image Tag') {
            when {
                expression { env.BRANCH_NAME == env.BRANCH}
            }
            steps {
                script {
                    TAG = sh(script: "${WORKSPACE}/docker_tag.sh $DOCKER_HUB_USERNAME $DOCKER_REPO_NAME $VERSION_PART", returnStdout: true).trim()

                    if (TAG) {
                        echo "Docker image tag generated successfully: $TAG"
                    } else {
                        error "Failed to generate Docker image tag"
                    }

                    env.TAG = TAG
                }
            }
        }

        stage('Docker Login') {
            when {
                expression { env.BRANCH_NAME == env.BRANCH }
            }
            steps {
                script {

                    withCredentials([usernamePassword(credentialsId: env.DOCKER_JENKINS_CERDIDENTALS_ID, passwordVariable: 'DOCKER_PASSWORD', usernameVariable: 'DOCKER_USERNAME')]) {
                        sh "docker login -u $DOCKER_USERNAME -p $DOCKER_PASSWORD"
                    }
                }
            }
        }

        stage('Build') {
            when {
                expression { env.BRANCH_NAME == env.BRANCH }
            }
            steps {
                script {
                    sh "docker build --no-cache -t ${DOCKER_HUB_USERNAME}/${DOCKER_REPO_NAME}:${TAG} ."
                }
            }
        }

        stage('Deploy') {
            when {
                expression { env.BRANCH_NAME == env.BRANCH }
            }
            steps {
                script {
                    sh "docker push ${DOCKER_HUB_USERNAME}/${DOCKER_REPO_NAME}:${TAG}"
                }
            }
        }

        stage('Environment Cleanup') {
            when {
                expression { env.BRANCH_NAME == env.BRANCH }
            }
            steps {
                script {
                    sh "docker rmi ${DOCKER_HUB_USERNAME}/${DOCKER_REPO_NAME}:${TAG}"
                }
            }
        }
    }

    post {
        success {
            echo "Pipeline completed successfully"
        }
    }
}