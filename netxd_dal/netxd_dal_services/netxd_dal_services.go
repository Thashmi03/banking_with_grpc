package netxddalservices

import (
	netxddalinterface "banking_with_grpc/netxd_dal/netxd_dal_interface"
	netxddalmodels "banking_with_grpc/netxd_dal/netxd_dal_models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerService struct {
	CustomerCollection *mongo.Collection
	ctx                context.Context
}


func InitCustomerService(collection *mongo.Collection, ctx context.Context) netxddalinterface.ICustomer {
	return &CustomerService{collection, ctx}
}

func (c * CustomerService)CreateCustomer(detail * netxddalmodels.Customer)(*netxddalmodels.DbResponse,error){
	detail.IsActive = true
	detail.CreatedAt = time.Now()
	detail.UpdatedAt = detail.CreatedAt

	res,err:=c.CustomerCollection.InsertOne(c.ctx,&detail)
	if err!=nil{
		return nil,err
	}
	var newUser *netxddalmodels.DbResponse
	query:=bson.M{"_id":res.InsertedID}

	err = c.CustomerCollection.FindOne(c.ctx,query).Decode(&newUser)
	if err != nil{
		return nil,err
	}
	return newUser,nil
}