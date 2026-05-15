# StartTech Cloud DevOps Infrastructure Project

## Project Overview

The StartTech Cloud DevOps Infrastructure Project is a production-style cloud-native application deployment built on Amazon Web Services (AWS) using modern DevOps tools and best practices.

This project demonstrates the implementation of:

* Infrastructure as Code (IaC) using Terraform
* CI/CD automation using GitHub Actions
* Frontend hosting using Amazon S3 and CloudFront
* Backend deployment using Docker and Amazon EC2
* Auto Scaling and Load Balancing
* Redis caching using ElastiCache
* MongoDB Atlas database integration
* CloudWatch monitoring and logging
* Secure deployment practices using IAM Roles and GitHub Secrets

The goal of this project is to simulate a real-world DevOps production environment that is scalable, automated, monitored, and cloud-native.

---

# Technologies Used

## Cloud Platform

* Amazon Web Services (AWS)

## Infrastructure as Code

* Terraform

## Backend Technologies

* Golang
* Docker
* Amazon ECR
* EC2

## Frontend Technologies

* React
* S3 Static Website Hosting
* CloudFront CDN

## Databases & Caching

* MongoDB Atlas
* Redis ElastiCache

## CI/CD & Automation

* GitHub Actions

## Monitoring & Logging

* AWS CloudWatch

## Version Control

* Git
* GitHub

---

# Project Architecture

## High-Level Architecture

```text
Users
   ↓
CloudFront CDN
   ↓
S3 Frontend Hosting
   ↓
Application Load Balancer
   ↓
Auto Scaling Group
   ↓
EC2 Docker Containers
   ↓
Redis + MongoDB Atlas
```

---

# Infrastructure Components

## 1. Networking Layer

### VPC

A custom VPC was created to isolate and secure the infrastructure environment.

### Public Subnets

Used for:

* Application Load Balancer
* Public-facing resources

### Private Subnets

Used for:

* Redis ElastiCache
* Backend services

### Internet Gateway

Provides internet access to public-facing resources.

### Route Tables

Used to route internet traffic appropriately between public and private subnets.

---

## 2. Compute Layer

### Launch Template

Defines EC2 instance configuration including:

* AMI
* Security groups
* User data
* Instance type

### Auto Scaling Group

Automatically manages EC2 instances by:

* Replacing unhealthy instances
* Scaling instances dynamically
* Ensuring high availability

### EC2 Backend Servers

Backend services are deployed inside Docker containers running on Ubuntu EC2 instances.

---

## 3. Load Balancing Layer

### Application Load Balancer (ALB)

Distributes incoming traffic across backend EC2 instances.

Features:

* Health checks
* Traffic distribution
* High availability

---

## 4. Frontend Hosting Layer

### Amazon S3

Used to host static React frontend build files.

### CloudFront

Provides:

* Global CDN delivery
* HTTPS support
* Faster frontend access
* Content caching

---

## 5. Backend Layer

### Golang API

The backend API was developed using Go.

The API listens on:

```text
Port 8080
```

### Docker

Backend services are containerized using Docker.

---

## 6. Database Layer

### MongoDB Atlas

Used as the primary database for persistent storage.

### Redis ElastiCache

Used for:

* Caching
* Temporary storage
* Performance optimization

---

## 7. Monitoring Layer

### CloudWatch

Used for:

* Log collection
* CPU monitoring
* Metrics visualization
* Alarms

---

# Terraform Infrastructure Setup

## Terraform Files Structure

```text
starttech-infra/
│
├── terraform/
│   ├── main.tf
│   ├── variables.tf
│   ├── outputs.tf
│   ├── provider.tf
│   └── terraform.tfvars
│
├── scripts/
│   └── user-data.sh
│
├── monitoring/
│   ├── cloudwatch-dashboard.json
│   ├── alarm-definitions.json
│   └── log-insights-queries.txt
│
└── .github/
    └── workflows/
        └── infrastructure-deploy.yml
```

---

# Backend Application Structure

```text
starttech-application/
│
├── backend/
│   ├── main.go
│   ├── go.mod
│   ├── Dockerfile
│   └── .env
│
├── frontend/
│   ├── src/
│   ├── public/
│   └── package.json
│
└── .github/
    └── workflows/
        ├── frontend-ci-cd.yml
        └── backend-ci-cd.yml
```

---

# CI/CD Pipelines

## 1. Infrastructure Pipeline

### Purpose

Automates Terraform deployment.

### Workflow Actions

* Terraform Init
* Terraform Format Check
* Terraform Validate
* Terraform Plan
* Terraform Apply

### Trigger

Runs automatically on:

```text
git push origin main
```

---

## 2. Frontend CI/CD Pipeline

### Purpose

Automates React frontend deployment.

### Workflow Actions

* Install dependencies
* Run tests
* Build React application
* Upload build files to S3
* Invalidate CloudFront cache

---

## 3. Backend CI/CD Pipeline

### Purpose

Automates backend Docker deployment.

### Workflow Actions

* Run Go tests
* Build Docker image
* Push Docker image to ECR
* SSH into EC2
* Pull latest image
* Run Docker container

---

# Security Implementation

## IAM Roles

IAM roles were configured to:

* Allow EC2 instances access to CloudWatch
* Allow EC2 instances pull access from ECR

## Security Groups

Security groups were configured for:

* SSH access
* HTTP traffic
* Backend API communication
* Redis access

## GitHub Secrets

Sensitive credentials are securely stored using GitHub Secrets.

Secrets used:

```text
AWS_ACCESS_KEY_ID
AWS_SECRET_ACCESS_KEY
AWS_REGION
EC2_HOST
EC2_USER
EC2_SSH_KEY
ECR_REPOSITORY
MONGO_URI
S3_BUCKET_NAME
CLOUDFRONT_DISTRIBUTION_ID
```

## Private Subnets

Sensitive backend services were placed inside private subnets.

---

# Monitoring and Logging

## CloudWatch Log Group

Used to centralize backend logs.

## CloudWatch Alarms

Configured for:

* High CPU usage
* Infrastructure monitoring

## CloudWatch Agent

Installed automatically on EC2 instances using user-data scripts.

---

# Backend Docker Deployment Process

## Docker Build

```bash
docker build -t starttech-backend .
```

## Docker Run

```bash
docker run -d -p 8080:8080 starttech-backend
```

## Docker Push to ECR

```bash
docker push <ecr-repository>
```

---

# GitHub Actions Backend Workflow

```yaml
name: Backend CI/CD Pipeline

on:
  push:
    branches:
      - main

jobs:
  backend-deploy:
    runs-on: ubuntu-latest

    defaults:
      run:
        working-directory: backend

    steps:
      - name: Checkout Code
        uses: actions/checkout@v4

      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.24'

      - name: Run Go Tests
        run: go test ./...

      - name: Configure AWS Credentials
        uses: aws-actions/configure-aws-credentials@v4
        with:
          aws-access-key-id: ${{ secrets.AWS_ACCESS_KEY_ID }}
          aws-secret-access-key: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          aws-region: ${{ secrets.AWS_REGION }}
```

---

# Runbook

# Infrastructure Deployment

## Terraform Initialization

```bash
terraform init
```

## Terraform Validation

```bash
terraform validate
```

## Terraform Plan

```bash
terraform plan
```

## Terraform Apply

```bash
terraform apply
```

---

# Frontend Deployment

Push frontend changes:

```bash
git push origin main
```

GitHub Actions automatically:

* Builds React application
* Deploys to S3
* Invalidates CloudFront cache

---

# Backend Deployment

Push backend changes:

```bash
git push origin main
```

GitHub Actions automatically:

* Builds Docker image
* Pushes image to ECR
* Deploys container to EC2

---

# Common Troubleshooting Guide

## 502 Bad Gateway

### Cause

Backend container not running.

### Solution

Check:

* Docker container status
* Target group health
* Backend port configuration

Commands:

```bash
sudo docker ps
```

---

## SSH Connection Refused

### Cause

* Port 22 blocked
* Wrong EC2 public IP
* Security group misconfiguration

### Solution

Verify:

* EC2 public IP
* Security group inbound rules
* SSH access manually

---

## ECR Pull Failure

### Cause

EC2 missing ECR permissions.

### Solution

Attach ECR IAM policy to EC2 role.

---

## CloudFront 404 Error

### Cause

Missing frontend build files.

### Solution

Upload React build files to S3.

---

# Deployment Verification

## Backend Verification

Open:

```text
http://<load-balancer-dns>
```

Expected response:

```text
Backend API is running 🚀
```

---

## Frontend Verification

Open:

```text
https://<cloudfront-url>
```

Expected:

* React frontend loads successfully

---

# Lessons Learned

This project provided practical experience in:

* AWS Infrastructure Design
* Infrastructure as Code
* CI/CD Automation
* Docker Containerization
* Cloud Security
* Monitoring and Logging
* DevOps Best Practices
* Cloud-Native Deployment

---

# Future Improvements

Possible future enhancements include:

* HTTPS with ACM Certificates
* Route53 custom domains
* ECS or Kubernetes migration
* Blue/Green deployments
* Prometheus and Grafana monitoring
* Remote Terraform State
* Terraform Modules
* WAF Integration

---

# Conclusion

The StartTech Cloud DevOps Infrastructure Project successfully demonstrates the deployment of a scalable, automated, monitored, and cloud-native application infrastructure using modern DevOps engineering practices.

This project reflects real-world production deployment concepts and showcases practical cloud engineering, automation, and infrastructure management skills.
