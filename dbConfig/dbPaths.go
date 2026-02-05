package dbConfig

import (
	mongo "go.mongodb.org/mongo-driver/mongo"
)

var DATABASE *mongo.Database



// const DATABASE_URL = "mongodb+srv://pmchamoth_db_user:t0rgjOG2h89qexQD@imomgt.y8j6wpi.mongodb.net/?appName=IMOMgt"

// const DATABASE_NAME ="Demand-Forecast"
const DATABASE_URL = "mongodb+srv://sachinayeshmantha_db_user:hL459yVavsSB24h5@k8s-marls-cluster1.gkm1ebb.mongodb.net/?appName=k8s-marls-cluster1"

const DATABASE_NAME = "User-Mgt-Service"