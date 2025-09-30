package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Tech-Preta/aws-resources/pkg/services"
)

// Exemplo básico de como usar os serviços AWS programaticamente
func main() {
	ctx := context.Background()

	// Exemplo 1: Criar bucket S3
	fmt.Println("=== Exemplo 1: Criando bucket S3 ===")
	
	s3Service, err := services.NewS3Service("us-east-1")
	if err != nil {
		log.Fatalf("Erro ao criar serviço S3: %v", err)
	}

	bucketParams := map[string]interface{}{
		"bucket_name": "meu-bucket-exemplo-" + fmt.Sprintf("%d", 12345),
		"region":      "us-east-1",
	}

	result, err := s3Service.CreateResource(ctx, bucketParams)
	if err != nil {
		log.Fatalf("Erro ao criar bucket: %v", err)
	}

	fmt.Printf("Resultado S3: Success=%t, Message=%s\n", result.Success, result.Message)
	if result.Data != nil {
		for key, value := range result.Data {
			fmt.Printf("  %s: %v\n", key, value)
		}
	}

	// Exemplo 2: Criar instância EC2
	fmt.Println("\n=== Exemplo 2: Criando instância EC2 ===")
	
	ec2Service, err := services.NewEC2Service("us-east-1")
	if err != nil {
		log.Fatalf("Erro ao criar serviço EC2: %v", err)
	}

	ec2Params := map[string]interface{}{
		"image_id":      "ami-0c94855ba95b798c7", // Amazon Linux 2023
		"instance_type": "t2.micro",
		"key_name":      "", // Deixe vazio se não tiver key pair
		"count":         "1",
		"region":        "us-east-1",
	}

	result, err = ec2Service.CreateResource(ctx, ec2Params)
	if err != nil {
		log.Fatalf("Erro ao criar instância EC2: %v", err)
	}

	fmt.Printf("Resultado EC2: Success=%t, Message=%s\n", result.Success, result.Message)
	if result.Data != nil {
		for key, value := range result.Data {
			fmt.Printf("  %s: %v\n", key, value)
		}
	}
}