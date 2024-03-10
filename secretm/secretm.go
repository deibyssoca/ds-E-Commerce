package secretm

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/deibyssoca/ds-E-Commerce/awsgo"
	"github.com/deibyssoca/ds-E-Commerce/models"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson
	fmt.Println(" > Get Secret " + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Conf)

	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})
	if err != nil {
		fmt.Println(err.Error())
		return secretData, err
	}
	/* conversión implisita del puntero de Key.SecretString(Json con toda la estructura)  */
	json.Unmarshal([]byte(*key.SecretString), &secretData)
	fmt.Println(" > read Secret OK " + secretName)
	return secretData, nil
}
