// Copyright © 2018 Julien SENON <julien.senon@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dynamo

import (
	"os"

	"github.com/jsenon/vpncentralmanager/config"
	"github.com/rs/zerolog/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// ConnectDynamo will create session to dynamodb
func ConnectDynamo() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Endpoint: aws.String(os.Getenv("urldynamo")),
		Region:   aws.String("eu-central-1")},
	)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", config.Service).
			Msgf("Cannot Connect to Dynamo for %s", config.Service)
	}
	// fmt.Println("Session to DB")
	return sess, err
}

// SearchDynamo will find item in dynamodb
func SearchDynamo(svc *dynamodb.DynamoDB, table string, searchfield string, attr string) (*dynamodb.GetItemOutput, error) {
	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			attr: {
				S: aws.String(searchfield),
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return result, err
}

// UpdateStatusDynamo update status col of an item
func UpdateStatusDynamo(svc *dynamodb.DynamoDB, table string, key map[string]*dynamodb.AttributeValue, newvalue map[string]*dynamodb.AttributeValue) error {
	input := &dynamodb.UpdateItemInput{
		Key:              key,
		TableName:        aws.String(table),
		UpdateExpression: aws.String("set #S = :s"),
		ExpressionAttributeNames: map[string]*string{
			"#S": aws.String("Status"),
		},
		ExpressionAttributeValues: newvalue,
		ReturnValues:              aws.String("UPDATED_NEW"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		log.Error().Msgf("Update item error:  %s ", err)
		return err
	}
	return nil
}

// UpdateipvpnDynamo update status col of an item
func UpdateipvpnDynamo(svc *dynamodb.DynamoDB, table string, key map[string]*dynamodb.AttributeValue, newvalue map[string]*dynamodb.AttributeValue) error {
	input := &dynamodb.UpdateItemInput{
		Key:              key,
		TableName:        aws.String(table),
		UpdateExpression: aws.String("set #V = :v"),
		ExpressionAttributeNames: map[string]*string{
			"#V": aws.String("AddressVpn"),
		},
		ExpressionAttributeValues: newvalue,
		ReturnValues:              aws.String("UPDATED_NEW"),
	}

	_, err := svc.UpdateItem(input)
	if err != nil {
		log.Error().Msgf("Update item error:  %s ", err)
		return err
	}
	log.Info().Msg("Successfully updated")
	return nil
}
